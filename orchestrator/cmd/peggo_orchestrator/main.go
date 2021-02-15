package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	_ "net/http/pprof"

	chainclient "github.com/InjectiveLabs/sdk-go/chain/client"
	"github.com/InjectiveLabs/sdk-go/chain/crypto/ethsecp256k1"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	cli "github.com/jawher/mow.cli"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	"github.com/InjectiveLabs/peggo/orchestrator"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/committer"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/peggy"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/provider"
	"github.com/InjectiveLabs/peggo/orchestrator/metrics"
	"github.com/InjectiveLabs/peggo/orchestrator/sidechain"
	"github.com/InjectiveLabs/peggo/orchestrator/sidechain/tmclient"
	"github.com/InjectiveLabs/peggo/orchestrator/version"
	"github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

var app = cli.App("peggo_orchestrator", "The Validator companion binary for Peggy.")

func main() {

	readEnv()
	initFlags()
	app.Before = prepareApp
	app.Action = runApp

	app.Command("version", "Print the version information and exit.", versionCmd)

	_ = app.Run(os.Args)
}

func readEnv() {
	if envdata, _ := ioutil.ReadFile(".env"); len(envdata) > 0 {
		s := bufio.NewScanner(bytes.NewReader(envdata))
		for s.Scan() {
			parts := strings.Split(s.Text(), "=")
			if len(parts) != 2 {
				continue
			}
			if err := os.Setenv(parts[0], parts[1]); err != nil {
				log.WithField("name", parts[0]).WithError(err).Warningln("failed to override ENV variable")
			}
		}
	}
}

func versionCmd(c *cli.Cmd) {
	c.Action = func() {
		fmt.Println(version.Version())
	}
}

func prepareApp() {
	app.LongDesc = `The Validator companion binary for Peggy. This must be run by all Peggy chain validators
            and is a mix of a relayer + oracle + ethereum signing infrastructure.`

	log.DefaultLogger.SetLevel(Level(*appLogLevel))
}

func runApp() {
	defer closer.Close()

	if toBool(*statsdDisabled) {
		// initializes statsd client with a mock one with no-op enabled
		metrics.Disable()
	} else {
		go func() {
			for {
				hostname, _ := os.Hostname()
				err := metrics.Init(*statsdAddr, checkStatsdPrefix(*statsdPrefix), &metrics.StatterConfig{
					EnvName:              *envName,
					HostName:             hostname,
					StuckFunctionTimeout: duration(*statsdStuckDur, 30*time.Minute),
					MockingEnabled:       toBool(*statsdMocking) || *envName == "local",
				})
				if err != nil {
					log.WithError(err).Warningln("metrics init failed, will retry in 1 min")
					time.Sleep(time.Minute)
					continue
				}
				break
			}
			closer.Bind(func() {
				metrics.Close()
			})
		}()
	}

	var err error
	var ethSignerPk *ecdsa.PrivateKey
	if len(*ethPrivKey) > 0 {
		ethSignerPk, err = ethcrypto.HexToECDSA(*ethPrivKey)
		orShutdown(err)

		fromAddr := ethcrypto.PubkeyToAddress(ethSignerPk.PublicKey)
		log.WithField("from", fromAddr.Hex()).Info("Loaded sender private key")
	} else {
		log.Fatalln("No Ethereum account credentials provided")
	}

	peggyAddress := common.HexToAddress(*contractAddrHex)
	if bytes.Equal(peggyAddress.Bytes(), common.Address{}.Bytes()) {
		log.Fatalln("no Peggy contract address specified, use --contract-address or PEGGY_CONTRACT_ADDRESS env")
	}

	evmRPC, err := rpc.Dial(*ethNodeRPC)
	if err != nil {
		log.WithField("endpoint", *ethNodeRPC).WithError(err).Fatalln("Failed to connect to Ethereum RPC")
		return
	}
	ethProvider := provider.NewEVMProvider(evmRPC)
	log.Infoln("Connected to Ethereum RPC at", *ethNodeRPC)

	cosmosPk := &ethsecp256k1.PrivKey{
		Key: common.FromHex(*cosmosPrivkey),
	}

	clientCtx, err := chainclient.NewClientContext(*chainId, cosmosPk)
	if err != nil {
		log.WithError(err).Fatalln("failed to initialize sidechain client context")
	}
	clientCtx = clientCtx.WithNodeURI(*tendermintRPC)
	tmRPC, err := rpchttp.New(*tendermintRPC, "/websocket")
	if err != nil {
		log.WithError(err)
	}
	clientCtx = clientCtx.WithClient(tmRPC)

	daemonClient, err := chainclient.NewCosmosClient(clientCtx, *cosmosGRPC)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"endpoint": *cosmosGRPC,
		}).Fatalln("failed to connect to daemon, is injectived running?")
	}

	log.Infoln("Waiting for injectived GRPC")
	time.Sleep(1 * time.Second)

	daemonWaitCtx, cancelWait := context.WithTimeout(context.Background(), time.Minute)
	waitForService(daemonWaitCtx, daemonClient)
	peggyQuerier := types.NewQueryClient(daemonClient.QueryClient())
	peggyBroadcaster := sidechain.NewPeggyBroadcastClient(peggyQuerier, daemonClient)
	cancelWait()

	ethCommitter, err := committer.NewEthCommitter(ethSignerPk, ethProvider)
	orShutdown(err)

	peggyContract, err := peggy.NewPeggyContract(ethCommitter, peggyAddress)
	orShutdown(err)

	cosmosQueryClient := sidechain.NewPeggyQueryClient(peggyQuerier)

	svc := orchestrator.NewPeggyOrchestrator(
		cosmosQueryClient,
		peggyBroadcaster,
		tmclient.NewRPCClient(*tendermintRPC),
		peggyContract,
		ethSignerPk,
		*injContractAddrHex,
	)

	ctx, cancelFn := context.WithCancel(context.Background())
	closer.Bind(cancelFn)
	svc.RunLoop(ctx)
}

func ethPrivkeyAddress(ethPrivkey *ecdsa.PrivateKey) string {
	return ethcrypto.PubkeyToAddress(ethPrivkey.PublicKey).Hex()
}

func waitForService(ctx context.Context, daemon chainclient.CosmosClient) {
	for {
		select {
		case <-ctx.Done():
			log.Fatalln("service wait timed out")
		default:
			state := daemon.QueryClient().GetState()

			if state != connectivity.Ready {
				log.WithField("state", state.String()).Warningln("state of grpc connection not ready")
				time.Sleep(5 * time.Second)
				continue
			}

			return
		}
	}
}

func orShutdown(err error) {
	if err != nil && err != grpc.ErrServerStopped {
		log.WithError(err).Fatalln("unable to start index-price-oracle")
	}
}
