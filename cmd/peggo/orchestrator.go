package peggo

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/celestiaorg/quantum-gravity-bridge/cmd/peggo/client"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/coingecko"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/cosmos"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/cosmos/tmclient"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/committer"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/peggy"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/provider"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/relayer"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	peggytypes "github.com/umee-network/umee/x/peggy/types"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func getOrchestratorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orchestrator",
		Short: "Starts the orchestrator",
		RunE: func(cmd *cobra.Command, args []string) error {
			konfig, err := parseServerConfig(cmd)
			if err != nil {
				return err
			}

			logger, err := getLogger(cmd)
			if err != nil {
				return err
			}

			cosmosUseLedger := konfig.Bool(flagCosmosUseLedger)
			ethUseLedger := konfig.Bool(flagEthUseLedger)
			if cosmosUseLedger || ethUseLedger {
				return fmt.Errorf("cannot use Ledger for orchestrator")
			}

			valAddress, cosmosKeyring, err := initCosmosKeyring(konfig)
			if err != nil {
				return fmt.Errorf("failed to initialize Cosmos keyring: %w", err)
			}

			ethChainID := konfig.Int64(flagEthChainID)
			ethKeyFromAddress, signerFn, personalSignFn, err := initEthereumAccountsManager(logger, uint64(ethChainID), konfig)
			if err != nil {
				return fmt.Errorf("failed to initialize Ethereum account: %w", err)
			}

			cosmosChainID := konfig.String(flagCosmosChainID)
			clientCtx, err := client.NewClientContext(cosmosChainID, valAddress.String(), cosmosKeyring)
			if err != nil {
				return err
			}

			tmRPCEndpoint := konfig.String(flagTendermintRPC)
			cosmosGRPC := konfig.String(flagCosmosGRPC)
			cosmosGasPrices := konfig.String(flagCosmosGasPrices)

			tmRPC, err := rpchttp.New(tmRPCEndpoint, "/websocket")
			if err != nil {
				return fmt.Errorf("failed to create Tendermint RPC client: %w", err)
			}

			fmt.Fprintf(os.Stderr, "Connected to Tendermint RPC: %s\n", tmRPCEndpoint)
			clientCtx = clientCtx.WithClient(tmRPC).WithNodeURI(tmRPCEndpoint)

			daemonClient, err := client.NewCosmosClient(clientCtx, logger, cosmosGRPC, client.OptionGasPrices(cosmosGasPrices))
			if err != nil {
				return err
			}

			// TODO: Clean this up to be more ergonomic and clean. We can probably
			// encapsulate all of this into a single utility function that gracefully
			// checks for the gRPC status/health.
			//
			// Ref: https://github.com/umee-network/peggo/issues/2
			fmt.Fprintln(os.Stderr, "Waiting for cosmos gRPC service...")
			time.Sleep(time.Second)

			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			defer cancel()

			gRPCConn := daemonClient.QueryClient()
			waitForService(ctx, gRPCConn)

			peggyQuerier := peggytypes.NewQueryClient(gRPCConn)
			peggyBroadcaster := cosmos.NewPeggyBroadcastClient(
				logger,
				peggyQuerier,
				daemonClient,
				signerFn,
				personalSignFn,
			)

			// query peggy params
			peggyQueryClient := cosmos.NewPeggyQueryClient(peggyQuerier)
			ctx, cancel = context.WithTimeout(context.Background(), time.Minute)
			defer cancel()

			peggyParams, err := peggyQueryClient.PeggyParams(ctx)
			if err != nil {
				return fmt.Errorf("failed to query for Peggy params: %w", err)
			}

			ethRPCEndpoint := konfig.String(flagEthRPC)
			ethRPC, err := ethrpc.Dial(ethRPCEndpoint)
			if err != nil {
				return fmt.Errorf("failed to dial Ethereum RPC node: %w", err)
			}

			fmt.Fprintf(os.Stderr, "Connected to Ethereum RPC: %s\n", ethRPCEndpoint)
			ethProvider := provider.NewEVMProvider(ethRPC)

			ethGasPriceAdjustment := konfig.Float64(flagEthGasAdjustment)
			ethCommitter, err := committer.NewEthCommitter(
				logger,
				ethKeyFromAddress,
				ethGasPriceAdjustment,
				signerFn,
				ethProvider,
			)
			if err != nil && err != grpc.ErrServerStopped {
				return fmt.Errorf("failed to create Ethereum committer: %w", err)
			}

			peggyAddress := ethcmn.HexToAddress(peggyParams.BridgeEthereumAddress)
			peggyContract, err := peggy.NewPeggyContract(logger, ethCommitter, peggyAddress)
			if err != nil {
				return fmt.Errorf("failed to create Ethereum committer: %w", err)
			}

			relayer := relayer.NewPeggyRelayer(
				logger,
				peggyQueryClient,
				peggyContract,
				konfig.Bool(flagRelayValsets),
				konfig.Bool(flagRelayBatches),
				konfig.Duration(flagRelayerLoopDuration),
			)

			coingeckoAPI := konfig.String(flagCoinGeckoAPI)
			coingeckoFeed := coingecko.NewCoingeckoPriceFeed(logger, 100, &coingecko.Config{
				BaseURL: coingeckoAPI,
			})

			logger = logger.With().
				Str("relayer_validator_addr", sdk.ValAddress(valAddress).String()).
				Str("relayer_ethereum_addr", ethKeyFromAddress.String()).
				Logger()

			orch := orchestrator.NewPeggyOrchestrator(
				logger,
				peggyQueryClient,
				peggyBroadcaster,
				tmclient.NewRPCClient(logger, tmRPCEndpoint),
				peggyContract,
				ethKeyFromAddress,
				signerFn,
				personalSignFn,
				relayer,
				konfig.Duration(flagOrchLoopDuration),
				konfig.Duration(flagCosmosBlockTime),
				orchestrator.SetMinBatchFee(konfig.Float64(flagMinBatchFeeUSD)),
				orchestrator.SetPriceFeeder(coingeckoFeed),
			)

			ctx, cancel = context.WithCancel(context.Background())
			g, errCtx := errgroup.WithContext(ctx)

			g.Go(func() error {
				return startOrchestrator(errCtx, logger, orch)
			})

			// listen for and trap any OS signal to gracefully shutdown and exit
			trapSignal(cancel)

			return g.Wait()
		},
	}

	cmd.Flags().Bool(flagRelayValsets, false, "Relay validator set updates to Ethereum")
	cmd.Flags().Bool(flagRelayBatches, false, "Relay transaction batches to Ethereum")
	cmd.Flags().Duration(flagRelayerLoopDuration, 5*time.Minute, "Duration between relayer loops")
	cmd.Flags().Duration(flagOrchLoopDuration, 1*time.Minute, "Duration between orchestrator loops")
	cmd.Flags().Duration(flagCosmosBlockTime, 5*time.Second, "Average block time of the cosmos chain")
	cmd.Flags().Float64(
		flagMinBatchFeeUSD,
		float64(0.0),
		"If non-zero, batch requests will only be made if fee threshold criteria is met",
	)
	cmd.Flags().String(flagCoinGeckoAPI, "https://api.coingecko.com/api/v3", "Specify the coingecko API endpoint")
	cmd.Flags().AddFlagSet(cosmosFlagSet())
	cmd.Flags().AddFlagSet(cosmosKeyringFlagSet())
	cmd.Flags().AddFlagSet(ethereumKeyOptsFlagSet())
	cmd.Flags().AddFlagSet(ethereumOptsFlagSet())

	return cmd
}

func trapSignal(cancel context.CancelFunc) {
	var sigCh = make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		sig := <-sigCh
		fmt.Fprintf(os.Stderr, "Caught signal (%s); shutting down...\n", sig)
		cancel()
	}()
}

func startOrchestrator(ctx context.Context, logger zerolog.Logger, orch orchestrator.PeggyOrchestrator) error {
	srvErrCh := make(chan error, 1)
	go func() {
		logger.Info().Msg("starting orchestrator...")
		srvErrCh <- orch.Start(ctx)
	}()

	for {
		select {
		case <-ctx.Done():
			return nil

		case err := <-srvErrCh:
			logger.Error().Err(err).Msg("failed to start orchestrator")
			return err
		}
	}
}
