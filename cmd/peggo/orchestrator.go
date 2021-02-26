package main

import (
	"bytes"
	"context"
	"os"
	"time"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	cli "github.com/jawher/mow.cli"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"

	"github.com/InjectiveLabs/peggo/modules/peggy/types"
	"github.com/InjectiveLabs/peggo/orchestrator"
	"github.com/InjectiveLabs/peggo/orchestrator/cosmos"
	"github.com/InjectiveLabs/peggo/orchestrator/cosmos/client"
	"github.com/InjectiveLabs/peggo/orchestrator/cosmos/tmclient"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/committer"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/peggy"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/provider"
)

// startOrchestrator action runs an infinite loop,
// listening for events and performing hooks.
//
// $ peggo orchestrator
func orchestratorCmd(cmd *cli.Cmd) {
	// orchestrator-specific CLI options
	var (
		// Cosmos params
		cosmosChainID  *string
		cosmosGRPC     *string
		tendermintRPC  *string
		cosmosFeeDenom *string

		// Cosmos Key Management
		cosmosKeyringDir     *string
		cosmosKeyringAppName *string
		cosmosKeyringBackend *string

		cosmosKeyFrom       *string
		cosmosKeyPassphrase *string
		cosmosPrivKey       *string
		cosmosUseLedger     *bool

		// Ethereum params
		ethChainID       *int
		ethNodeRPC       *string
		ethPeggyContract *string

		// Ethereum Key Management
		ethKeystoreDir *string
		ethKeyFrom     *string
		ethPassphrase  *string
		ethPrivKey     *string
		ethUseLedger   *bool

		// ERC20 Contract Mapping
		erc20ContractMapping *[]string

		// Metrics
		statsdPrefix   *string
		statsdAddr     *string
		statsdStuckDur *string
		statsdMocking  *string
		statsdDisabled *string
	)

	initCosmosOptions(
		cmd,
		&cosmosChainID,
		&cosmosGRPC,
		&tendermintRPC,
		&cosmosFeeDenom,
	)

	initCosmosKeyOptions(
		cmd,
		&cosmosKeyringDir,
		&cosmosKeyringAppName,
		&cosmosKeyringBackend,
		&cosmosKeyFrom,
		&cosmosKeyPassphrase,
		&cosmosPrivKey,
		&cosmosUseLedger,
	)

	initEthereumOptions(
		cmd,
		&ethChainID,
		&ethNodeRPC,
		&ethPeggyContract,
	)

	initEthereumKeyOptions(
		cmd,
		&ethKeystoreDir,
		&ethKeyFrom,
		&ethPassphrase,
		&ethPrivKey,
		&ethUseLedger,
	)

	initStatsdOptions(
		cmd,
		&statsdPrefix,
		&statsdAddr,
		&statsdStuckDur,
		&statsdMocking,
		&statsdDisabled,
	)

	erc20ContractMapping = cmd.StringsArg("ERC20_MAPPING", []string{}, "Mapping between cosmos_denom:contract_address for ERC20 tokens")

	cmd.Action = func() {
		// ensure a clean exit
		defer closer.Close()

		startMetricsGathering(
			statsdPrefix,
			statsdAddr,
			statsdStuckDur,
			statsdMocking,
			statsdDisabled,
		)

		if *cosmosUseLedger || *ethUseLedger {
			log.Fatalln("cannot really use Ledger for orchestrator, since signatures msut be realtime")
		}

		peggyAddress := ethcmn.HexToAddress(*ethPeggyContract)
		if bytes.Equal(peggyAddress.Bytes(), ethcmn.Address{}.Bytes()) {
			log.Fatalln("no Peggy contract address specified, use --contract-address or PEGGY_CONTRACT_ADDRESS env")
		}

		evmRPC, err := rpc.Dial(*ethNodeRPC)
		if err != nil {
			log.WithField("endpoint", *ethNodeRPC).WithError(err).Fatalln("Failed to connect to Ethereum RPC")
			return
		}
		ethProvider := provider.NewEVMProvider(evmRPC)
		log.Infoln("Connected to Ethereum RPC at", *ethNodeRPC)

		valAddress, cosmosKeyring, err := initCosmosKeyring(
			cosmosKeyringDir,
			cosmosKeyringAppName,
			cosmosKeyringBackend,
			cosmosKeyFrom,
			cosmosKeyPassphrase,
			cosmosPrivKey,
			cosmosUseLedger,
		)
		if err != nil {
			log.WithError(err).Fatalln("failed to init Cosmos keyring")
		}

		ethKeyFromAddress, signerFn, personalSignFn, err := initEthereumAccountsManager(
			uint64(*ethChainID),
			ethKeystoreDir,
			ethKeyFrom,
			ethPassphrase,
			ethPrivKey,
			ethUseLedger,
		)
		if err != nil {
			log.WithError(err).Fatalln("failed to init Ethereum account")
		}

		log.Infoln("Using Cosmos ValAddress", valAddress.String())
		log.Infoln("Using Ethereum address", ethKeyFromAddress.String())

		clientCtx, err := client.NewClientContext(*cosmosChainID, valAddress.String(), cosmosKeyring)
		if err != nil {
			log.WithError(err).Fatalln("failed to initialize cosmos client context")
		}
		clientCtx = clientCtx.WithNodeURI(*tendermintRPC)
		tmRPC, err := rpchttp.New(*tendermintRPC, "/websocket")
		if err != nil {
			log.WithError(err)
		}
		clientCtx = clientCtx.WithClient(tmRPC)

		daemonClient, err := client.NewCosmosClient(clientCtx, *cosmosGRPC)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"endpoint": *cosmosGRPC,
			}).Fatalln("failed to connect to daemon, is injectived running?")
		}

		log.Infoln("Waiting for injectived GRPC")
		time.Sleep(1 * time.Second)

		daemonWaitCtx, cancelWait := context.WithTimeout(context.Background(), time.Minute)
		grpcConn := daemonClient.QueryClient()
		waitForService(daemonWaitCtx, grpcConn)
		peggyQuerier := types.NewQueryClient(grpcConn)
		peggyBroadcaster := cosmos.NewPeggyBroadcastClient(
			peggyQuerier,
			daemonClient,
			signerFn,
			personalSignFn,
		)
		cancelWait()

		ethCommitter, err := committer.NewEthCommitter(ethcmn.Address{}, nil, ethProvider)
		orShutdown(err)

		peggyContract, err := peggy.NewPeggyContract(ethCommitter, peggyAddress)
		orShutdown(err)

		cosmosQueryClient := cosmos.NewPeggyQueryClient(peggyQuerier)

		svc := orchestrator.NewPeggyOrchestrator(
			cosmosQueryClient,
			peggyBroadcaster,
			tmclient.NewRPCClient(*tendermintRPC),
			peggyContract,
			ethKeyFromAddress,
			signerFn,
			personalSignFn,
			parseERC20ContractMapping(*erc20ContractMapping),
		)

		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)

		go func() {
			if err := svc.Start(ctx); err != nil {
				log.Errorln(err)

				// signal there that the app failed
				os.Exit(1)
			}
		}()

		closer.Hold()
	}
}
