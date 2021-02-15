package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/ecdsa"
	"io/ioutil"
	"os"
	"strings"
	"time"

	chainclient "github.com/InjectiveLabs/sdk-go/chain/client"
	"github.com/InjectiveLabs/sdk-go/chain/crypto/ethsecp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	cli "github.com/jawher/mow.cli"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"
	"google.golang.org/grpc/connectivity"

	"github.com/InjectiveLabs/peggo/orchestrator/sidechain"
	"github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

var app = cli.App("register_eth_key", "Special purpose binary for bootstrapping Peggy chains.")

func main() {
	readEnv()
	initFlags()
	app.Before = prepareApp
	app.Action = runApp

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

func prepareApp() {
	app.Spec = "[OPTIONS] [ETH_PRIVKEY]"

	app.LongDesc = `Special purpose binary for bootstrapping Peggy chains. This will submit and optionally
            generate an Ethereum key that will be used to sign messages on behalf of your Validator
            on the Cosmos blockchain running the Peggy module. Be aware this Ethereum key must be kept
            safe as you can be slashed for losing it.`
}

func runApp() {
	defer closer.Close()

	var err error
	var ethPrivkeyVar *ecdsa.PrivateKey
	if len(*ethPrivKey) == 0 {
		log.Infoln("Generatig new Ethereum privkey, please save it")
		pk, err := ethcrypto.GenerateKey()
		if err != nil {
			log.Fatalln(err)
		}

		ethPrivkeyVar = pk
		log.Infoln("Generate privkey with address", ethPrivkeyAddress(ethPrivkeyVar))
	} else {
		ethPrivkeyVar, err = ethcrypto.HexToECDSA(*ethPrivKey)
		if err != nil {
			log.Fatalln(err)
		}

		log.Infoln("Loaded provided privkey for address", ethPrivkeyAddress(ethPrivkeyVar))
	}

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

	broadcastCtx, cancelFn := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFn()

	if err = peggyBroadcaster.UpdatePeggyEthAddress(broadcastCtx, ethPrivkeyVar); err != nil {
		log.WithError(err).Errorln("failed to broadcast Tx")
		time.Sleep(time.Second)
		return
	}

	log.Infof("Registered Ethereum address %s for validator address %s",
		ethPrivkeyAddress(ethPrivkeyVar), cryptotypes.PrivKey(cosmosPk).PubKey().Address().String())
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
