package peggo

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"os"
	"time"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/knadh/koanf"
	"github.com/spf13/cobra"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/umee-network/peggo/cmd/peggo/client"
	"github.com/umee-network/peggo/orchestrator/cosmos"
	wrappers "github.com/umee-network/peggo/solidity/wrappers/Peggy.sol"
	peggytypes "github.com/umee-network/umee/x/peggy/types"
	"google.golang.org/grpc"
)

func getBridgeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge",
		Short: "Commands to interface with Peggy (Gravity Bridge) Ethereum contract",
		Long: `Commands to interface with Peggy (Gravity Bridge) Ethereum contract.
		
Inputs in the CLI commands can be provided via flags or environment variables. If
using the later, prefix the environment variable with PEGGO_ and the named of the
flag (e.g. PEGGO_COSMOS_PK).`,
	}

	cmd.PersistentFlags().AddFlagSet(cosmosFlagSet())
	cmd.PersistentFlags().AddFlagSet(bridgeFlagSet())

	cmd.AddCommand(
		deployPeggyCmd(),
		initPeggyCmd(),
		deployERC20Cmd(),
	)

	return cmd
}

// TODO: Support --admin capabilities.
func deployPeggyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy-peggy",
		Short: "Deploy the Peggy (Gravity Bridge) smart contract on Ethereum",
		RunE: func(cmd *cobra.Command, args []string) error {
			konfig, err := parseServerConfig(cmd)
			if err != nil {
				return err
			}

			ethRPCEndpoint := konfig.String(flagEthRPC)
			ethRPC, err := ethclient.Dial(ethRPCEndpoint)
			if err != nil {
				return fmt.Errorf("failed to dial Ethereum RPC node: %w", err)
			}

			auth, err := buildTransactOpts(konfig, ethRPC)
			if err != nil {
				return err
			}

			address, tx, _, err := wrappers.DeployPeggy(auth, ethRPC)
			if err != nil {
				return fmt.Errorf("failed deploy Peggy (Gravity Bridge) contract: %w", err)
			}

			_, _ = fmt.Fprintf(os.Stderr, `Peggy (Gravity Bridge) contract successfully deployed!
Address: %s
Transaction: %s
`,
				address.Hex(),
				tx.Hash().Hex(),
			)

			return nil
		},
	}

	return cmd
}

func initPeggyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init-peggy",
		Short: "Initialize the Peggy (Gravity Bridge) smart contract on Ethereum",
		Long: `Initialize the Peggy (Gravity Bridge) smart contract on Ethereum using
the current validator set and their respective powers.

Note, each validator must have their Ethereum delegate keys registered on chain
prior to initializing.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			konfig, err := parseServerConfig(cmd)
			if err != nil {
				return err
			}

			cosmosChainID := konfig.String(flagCosmosChainID)
			clientCtx, err := client.NewClientContext(cosmosChainID, "", nil)
			if err != nil {
				return err
			}

			tmRPCEndpoint := konfig.String(flagTendermintRPC)
			cosmosGRPC := konfig.String(flagCosmosGRPC)

			tmRPC, err := rpchttp.New(tmRPCEndpoint, "/websocket")
			if err != nil {
				return fmt.Errorf("failed to create Tendermint RPC client: %w", err)
			}

			fmt.Fprintf(os.Stderr, "Connected to Tendermint RPC: %s\n", tmRPCEndpoint)
			clientCtx = clientCtx.WithClient(tmRPC).WithNodeURI(tmRPCEndpoint)

			daemonClient, err := client.NewCosmosClient(clientCtx, cosmosGRPC)
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

			peggyParams, err := getPeggyParams(gRPCConn)
			if err != nil {
				return err
			}

			ethRPCEndpoint := konfig.String(flagEthRPC)
			ethRPC, err := ethclient.Dial(ethRPCEndpoint)
			if err != nil {
				return fmt.Errorf("failed to dial Ethereum RPC node: %w", err)
			}

			peggyContract, err := getPeggyContract(konfig, peggyParams, gRPCConn)
			if err != nil {
				return err
			}

			auth, err := buildTransactOpts(konfig, ethRPC)
			if err != nil {
				return err
			}

			powerThresholdInt := konfig.Int64(flagPowerThreshold)
			if powerThresholdInt < 0 {
				return fmt.Errorf("invalid power threshold: %d", powerThresholdInt)
			}

			powerThreshold := big.NewInt(powerThresholdInt)

			var peggyID [32]byte
			copy(peggyID[:], peggyParams.PeggyId)

			peggyQueryClient := cosmos.NewPeggyQueryClient(peggytypes.NewQueryClient(gRPCConn))
			currValSet, err := peggyQueryClient.CurrentValset(cmd.Context())
			if err != nil {
				return err
			}

			var (
				validators = make([]ethcmn.Address, len(currValSet.Members))
				powers     = make([]*big.Int, len(currValSet.Members))

				totalPower uint64
			)

			for i, member := range currValSet.Members {
				validators[i] = ethcmn.HexToAddress(member.EthereumAddress)
				powers[i] = new(big.Int).SetUint64(member.Power)
				totalPower += member.Power
			}

			if totalPower < uint64(powerThresholdInt) {
				return fmt.Errorf(
					"refusing to deploy; total power (%d) < power threshold (%d)",
					totalPower, powerThresholdInt,
				)
			}

			tx, err := peggyContract.Initialize(auth, peggyID, powerThreshold, validators, powers)
			if err != nil {
				return fmt.Errorf("failed to initialize Peggy (Gravity Bridge): %w", err)
			}

			_, _ = fmt.Fprintf(os.Stderr, `Peggy (Gravity Bridge) contract successfully initialized!
Gravity Addres: %s
PeggyID: %s
Init Params:
  Peggy ID: 0x%X
  Power Threshold: %d
  Validator Set Size: %d
  Validator Total Power: %d
Transaction: %s
			`,
				peggyParams.BridgeEthereumAddress,
				peggyParams.PeggyId,
				peggyID,
				powerThresholdInt,
				len(validators),
				totalPower,
				tx.Hash().Hex(),
			)

			return nil
		},
	}

	cmd.Flags().Uint64(flagPowerThreshold, 2834678415, "The validator power threshold to initialize Peggy with")

	return cmd
}

func deployERC20Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy-erc20 [base-denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Deploy a Cosmos native asset on Ethereum as an ERC20 token",
		RunE: func(cmd *cobra.Command, args []string) error {
			konfig, err := parseServerConfig(cmd)
			if err != nil {
				return err
			}

			ethRPCEndpoint := konfig.String(flagEthRPC)
			ethRPC, err := ethclient.Dial(ethRPCEndpoint)
			if err != nil {
				return fmt.Errorf("failed to dial Ethereum RPC node: %w", err)
			}

			auth, err := buildTransactOpts(konfig, ethRPC)
			if err != nil {
				return err
			}

			// query for the name and symbol on-chain via the token's metadata
			cosmosChainID := konfig.String(flagCosmosChainID)
			clientCtx, err := client.NewClientContext(cosmosChainID, "", nil)
			if err != nil {
				return err
			}

			tmRPCEndpoint := konfig.String(flagTendermintRPC)
			cosmosGRPC := konfig.String(flagCosmosGRPC)

			tmRPC, err := rpchttp.New(tmRPCEndpoint, "/websocket")
			if err != nil {
				return fmt.Errorf("failed to create Tendermint RPC client: %w", err)
			}

			fmt.Fprintf(os.Stderr, "Connected to Tendermint RPC: %s\n", tmRPCEndpoint)
			clientCtx = clientCtx.WithClient(tmRPC).WithNodeURI(tmRPCEndpoint)

			daemonClient, err := client.NewCosmosClient(clientCtx, cosmosGRPC)
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

			peggyParams, err := getPeggyParams(gRPCConn)
			if err != nil {
				return err
			}

			peggyContract, err := getPeggyContract(konfig, peggyParams, gRPCConn)
			if err != nil {
				return err
			}

			baseDenom := args[0]
			bankQuerier := banktypes.NewQueryClient(gRPCConn)

			ctx, cancel = context.WithTimeout(context.Background(), time.Minute)
			defer cancel()

			resp, err := bankQuerier.DenomMetadata(ctx, &banktypes.QueryDenomMetadataRequest{Denom: baseDenom})
			if err != nil {
				return fmt.Errorf("failed to query for bank metadata: %w", err)
			}

			var decimals uint8
			for _, unit := range resp.Metadata.DenomUnits {
				if unit.Denom == baseDenom {
					if unit.Exponent > math.MaxUint8 {
						return fmt.Errorf("token exponent too large; %d > %d", unit.Exponent, math.MaxInt8)
					}

					decimals = uint8(unit.Exponent)
					break
				}
			}

			tx, err := peggyContract.DeployERC20(auth, baseDenom, resp.Metadata.Name, resp.Metadata.Symbol, decimals)
			if err != nil {
				return fmt.Errorf("failed deploy Cosmos native ERC20 token: %w", err)
			}

			_, _ = fmt.Fprintf(os.Stderr, `Cosmos native token deployed as an ERC20 on Ethereum!
Base Denom: %s
Name: %s
Symbol: %s
Transaction: %s
			`,
				baseDenom,
				resp.Metadata.Name,
				resp.Metadata.Symbol,
				tx.Hash().Hex(),
			)

			return nil
		},
	}

	return cmd
}

func buildTransactOpts(konfig *koanf.Koanf, ethClient *ethclient.Client) (*bind.TransactOpts, error) {
	ethPrivKeyHexStr := konfig.String(flagEthPK)

	privKey, err := ethcrypto.HexToECDSA(ethPrivKeyHexStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("invalid public key; expected: %T, got: %T", &ecdsa.PublicKey{}, publicKey)
	}

	goCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	fromAddress := ethcrypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := ethClient.PendingNonceAt(goCtx, fromAddress)
	if err != nil {
		return nil, err
	}

	goCtx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ethChainID, err := ethClient.ChainID(goCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to get Ethereum chain ID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, ethChainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create Ethereum transactor: %w", err)
	}

	var gasPrice *big.Int

	gasPriceInt := konfig.Int64(flagEthGasPrice)
	switch {
	case gasPriceInt < 0:
		return nil, fmt.Errorf("invalid Ethereum gas price: %d", gasPriceInt)

	case gasPriceInt > 0:
		gasPrice = big.NewInt(gasPriceInt)

	default:
		gasPrice, err = ethClient.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get Ethereum gas estimate: %w", err)
		}
	}

	gasLimit := konfig.Int64(flagEthGasLimit)
	if gasLimit < 0 {
		return nil, fmt.Errorf("invalid Ethereum gas limit: %d", gasLimit)
	}

	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(gasLimit) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

func getPeggyParams(gRPCConn *grpc.ClientConn) (*peggytypes.Params, error) {
	peggyQueryClient := cosmos.NewPeggyQueryClient(peggytypes.NewQueryClient(gRPCConn))

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	peggyParams, err := peggyQueryClient.PeggyParams(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query for Peggy params: %w", err)
	}

	return peggyParams, nil
}

func getPeggyContract(konfig *koanf.Koanf, peggyParams *peggytypes.Params, gRPCConn *grpc.ClientConn) (*wrappers.Peggy, error) {
	ethRPCEndpoint := konfig.String(flagEthRPC)
	ethRPC, err := ethclient.Dial(ethRPCEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to dial Ethereum RPC node: %w", err)
	}

	contract, err := wrappers.NewPeggy(ethcmn.HexToAddress(peggyParams.BridgeEthereumAddress), ethRPC)
	if err != nil {
		return nil, fmt.Errorf("failed to create Peggy contract instance: %w", err)
	}

	return contract, nil
}
