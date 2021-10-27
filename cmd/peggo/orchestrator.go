package peggo

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	ethcmn "github.com/ethereum/go-ethereum/common"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/cobra"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/umee-network/peggo/cmd/peggo/client"
	"github.com/umee-network/peggo/orchestrator"
	"github.com/umee-network/peggo/orchestrator/coingecko"
	"github.com/umee-network/peggo/orchestrator/cosmos"
	"github.com/umee-network/peggo/orchestrator/cosmos/tmclient"
	"github.com/umee-network/peggo/orchestrator/ethereum/committer"
	"github.com/umee-network/peggo/orchestrator/ethereum/peggy"
	"github.com/umee-network/peggo/orchestrator/ethereum/provider"
	"github.com/umee-network/peggo/orchestrator/relayer"
	peggytypes "github.com/umee-network/umee/x/peggy/types"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func getOrchestratorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orchestrator",
		Short: "Starts the orchestrator",
		Long: `Starts the orchestrator's main relayer loop. Only start the orchestrator
if the Peggy (Gravity Bridge) contract has been deployed and initialized, which
requires all validators to set their delegate keys.

Inputs in the CLI commands can be provided via flags or environment variables. If
using the later, prefix the environment variable with PEGGO_ and the named of the
flag (e.g. PEGGO_COSMOS_PK).`,
		RunE: func(cmd *cobra.Command, args []string) error {
			konfig, err := parseServerConfig(cmd)
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
			ethKeyFromAddress, signerFn, personalSignFn, err := initEthereumAccountsManager(uint64(ethChainID), konfig)
			if err != nil {
				return fmt.Errorf("failed to initialize Ethereum account: %w", err)
			}

			fmt.Fprintf(os.Stderr, "Using Cosmos validator address: %s\n", valAddress)
			fmt.Fprintf(os.Stderr, "Using Ethereum address: %s\n", ethKeyFromAddress)

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

			daemonClient, err := client.NewCosmosClient(clientCtx, cosmosGRPC, client.OptionGasPrices(cosmosGasPrices))
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

			erc20ContractMapping := make(map[ethcmn.Address]string)
			// TODO: Figure out what this is and if we need it???
			// erc20Addr := ethcmn.HexToAddress(peggyParams.CosmosCoinErc20Contract)
			// TODO: add this as a config to peggyParams
			// erc20ContractMapping[erc20Addr] = "umee" //ctypes.InjectiveCoin

			ethRPCEndpoint := konfig.String(flagEthRPC)
			ethRPC, err := ethrpc.Dial(ethRPCEndpoint)
			if err != nil {
				return fmt.Errorf("failed to dial Ethereum RPC node: %w", err)
			}

			fmt.Fprintf(os.Stderr, "Connected to Ethereum RPC: %s\n", ethRPCEndpoint)
			ethProvider := provider.NewEVMProvider(ethRPC)

			ethGasPriceAdjustment := konfig.Float64(flagEthGasAdjustment)
			ethCommitter, err := committer.NewEthCommitter(ethKeyFromAddress, ethGasPriceAdjustment, signerFn, ethProvider)
			if err != nil && err != grpc.ErrServerStopped {
				return fmt.Errorf("failed to create Ethereum committer: %w", err)
			}

			peggyAddress := ethcmn.HexToAddress(peggyParams.BridgeEthereumAddress)
			peggyContract, err := peggy.NewPeggyContract(ethCommitter, peggyAddress)
			if err != nil {
				return fmt.Errorf("failed to create Ethereum committer: %w", err)
			}

			relayValSets := konfig.Bool(flagRelayValsets)
			relayBatches := konfig.Bool(flagRelayBatches)
			relayer := relayer.NewPeggyRelayer(peggyQueryClient, peggyContract, relayValSets, relayBatches)

			coingeckoAPI := konfig.String(flagCoinGeckoAPI)
			coingeckoFeed := coingecko.NewCoingeckoPriceFeed(100, &coingecko.Config{
				BaseURL: coingeckoAPI,
			})

			// TODO: Pass logger to NewPeggyOrchestrator
			// logger, err := getLogger(cmd)
			// if err != nil {
			// 	return err
			// }

			minBatchFeeUSD := konfig.Float64(flagMinBatchFeeUSD)
			orch := orchestrator.NewPeggyOrchestrator(
				peggyQueryClient,
				peggyBroadcaster,
				tmclient.NewRPCClient(tmRPCEndpoint),
				peggyContract,
				ethKeyFromAddress,
				signerFn,
				personalSignFn,
				erc20ContractMapping,
				relayer,
				minBatchFeeUSD,
				coingeckoFeed,
			)

			ctx, cancel = context.WithCancel(context.Background())
			g, errCtx := errgroup.WithContext(ctx)

			g.Go(func() error {
				return startOrchestrator(errCtx, orch)
			})

			// listen for and trap any OS signal to gracefully shutdown and exit
			trapSignal(cancel)

			return g.Wait()
		},
	}

	cmd.Flags().Bool(flagRelayValsets, false, "Relay validator set updates to Ethereum")
	cmd.Flags().Bool(flagRelayBatches, false, "Relay transaction batches to Ethereum")
	cmd.Flags().Float64(flagMinBatchFeeUSD, float64(23.3), "If non-zero, batch requests will only be made if fee threshold criteria is met")
	cmd.Flags().String(flagCoinGeckoAPI, "https://api.coingecko.com/api/v3", "Specify the coingecko API endpoint")
	cmd.Flags().AddFlagSet(cosmosFlagSet())
	cmd.Flags().AddFlagSet(cosmosKeyringFlagSet())
	cmd.Flags().AddFlagSet(ethereumKeyOptsFlagSet())
	cmd.Flags().AddFlagSet(ethereumOptsFlagSet())

	return cmd
}

func trapSignal(cancel context.CancelFunc) {
	var sigCh = make(chan os.Signal)

	signal.Notify(sigCh, syscall.SIGTERM)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		sig := <-sigCh
		fmt.Fprintf(os.Stderr, "Caught signal (%s); shutting down...\n", sig)
		cancel()
	}()
}

func startOrchestrator(ctx context.Context, orch orchestrator.PeggyOrchestrator) error {
	srvErrCh := make(chan error, 1)
	go func() {
		fmt.Fprintln(os.Stderr, "Starting orchestrator...")
		srvErrCh <- orch.Start(ctx)
	}()

	for {
		select {
		case <-ctx.Done():
			return nil

		case err := <-srvErrCh:
			fmt.Fprintln(os.Stderr, "Failed to start orchestrator")
			return err
		}
	}

}
