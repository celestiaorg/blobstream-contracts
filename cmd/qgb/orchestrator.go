package main

import (
	"context"
	"os"
	"time"

	ethcmn "github.com/ethereum/go-ethereum/common"
	cli "github.com/jawher/mow.cli"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"

	"github.com/InjectiveLabs/sdk-go/chain/client"
	"github.com/InjectiveLabs/sdk-go/chain/peggy/types"

	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/coingecko"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/cosmos"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/cosmos/tmclient"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/committer"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/peggy"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/provider"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/relayer"

	ctypes "github.com/InjectiveLabs/sdk-go/chain/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// startOrchestrator action runs an infinite loop,
// listening for events and performing hooks.
//
// $ qgb orchestrator
func orchestratorCmd(cmd *cli.Cmd) {
	// orchestrator-specific CLI options
	var (
		// Cosmos params
		cosmosChainID   *string
		cosmosGRPC      *string
		tendermintRPC   *string
		cosmosGasPrices *string

		// Cosmos Key Management
		cosmosKeyringDir     *string
		cosmosKeyringAppName *string
		cosmosKeyringBackend *string

		cosmosKeyFrom       *string
		cosmosKeyPassphrase *string
		cosmosPrivKey       *string
		cosmosUseLedger     *bool

		// Ethereum params
		ethChainID            *int
		ethNodeRPC            *string
		ethGasPriceAdjustment *float64

		// Ethereum Key Management
		ethKeystoreDir *string
		ethKeyFrom     *string
		ethPassphrase  *string
		ethPrivKey     *string
		ethUseLedger   *bool

		// Relayer config
		relayValsets *bool
		relayBatches *bool

		// Batch requester config
		minBatchFeeUSD *float64

		coingeckoApi *string
	)

	initCosmosOptions(
		cmd,
		&cosmosChainID,
		&cosmosGRPC,
		&tendermintRPC,
		&cosmosGasPrices,
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
		&ethGasPriceAdjustment,
	)

	initEthereumKeyOptions(
		cmd,
		&ethKeystoreDir,
		&ethKeyFrom,
		&ethPassphrase,
		&ethPrivKey,
		&ethUseLedger,
	)

	initRelayerOptions(
		cmd,
		&relayValsets,
		&relayBatches,
	)

	initBatchRequesterOptions(
		cmd,
		&minBatchFeeUSD,
	)

	initCoingeckoOptions(
		cmd,
		&coingeckoApi,
	)

	cmd.Before = func() {
		initMetrics(cmd)
	}

	cmd.Action = func() {
		// ensure a clean exit
		defer closer.Close()

		if *cosmosUseLedger || *ethUseLedger {
			log.Fatalln("cannot really use Ledger for orchestrator, since signatures msut be realtime")
		}

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

		daemonClient, err := client.NewCosmosClient(clientCtx, *cosmosGRPC, client.OptionGasPrices(*cosmosGasPrices))
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"endpoint": *cosmosGRPC,
			}).Fatalln("failed to connect to daemon, is celestiad running?")
		}

		log.Infoln("Waiting for celestiad GRPC")
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

		// Query peggy params
		cosmosQueryClient := cosmos.NewPeggyQueryClient(peggyQuerier)
		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)

		peggyParams, err := cosmosQueryClient.PeggyParams(ctx)
		if err != nil {
			log.WithError(err).Fatalln("failed to query peggy params, is celestiad running?")
		}

		peggyAddress := ethcmn.HexToAddress(peggyParams.BridgeEthereumAddress)
		injAddress := ethcmn.HexToAddress(peggyParams.CosmosCoinErc20Contract)

		erc20ContractMapping := make(map[ethcmn.Address]string)
		erc20ContractMapping[injAddress] = ctypes.InjectiveCoin

		evmRPC, err := rpc.Dial(*ethNodeRPC)
		if err != nil {
			log.WithField("endpoint", *ethNodeRPC).WithError(err).Fatalln("Failed to connect to Ethereum RPC")
			return
		}
		ethProvider := provider.NewEVMProvider(evmRPC)
		log.Infoln("Connected to Ethereum RPC at", *ethNodeRPC)

		ethCommitter, err := committer.NewEthCommitter(ethKeyFromAddress, *ethGasPriceAdjustment, signerFn, ethProvider)
		orShutdown(err)

		peggyContract, err := peggy.NewPeggyContract(ethCommitter, peggyAddress)
		orShutdown(err)

		relayer := relayer.NewPeggyRelayer(cosmosQueryClient, peggyContract, *relayValsets, *relayBatches)

		coingeckoConfig := coingecko.Config{
			BaseURL: *coingeckoApi,
		}
		coingeckoFeed := coingecko.NewCoingeckoPriceFeed(100, &coingeckoConfig)

		svc := orchestrator.NewPeggyOrchestrator(
			cosmosQueryClient,
			peggyBroadcaster,
			tmclient.NewRPCClient(*tendermintRPC),
			peggyContract,
			ethKeyFromAddress,
			signerFn,
			personalSignFn,
			erc20ContractMapping,
			relayer,
			*minBatchFeeUSD,
			coingeckoFeed,
		)

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
