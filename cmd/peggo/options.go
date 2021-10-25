package main

import cli "github.com/jawher/mow.cli"

// initGlobalOptions defines some global CLI options, that are useful for most parts of the app.
// Before adding option to there, consider moving it into the actual Cmd.
func initGlobalOptions(
	envName **string,
	appLogLevel **string,
	svcWaitTimeout **string,
) {
	*appLogLevel = app.String(cli.StringOpt{
		Name:   "l log-level",
		Desc:   "Available levels: error, warn, info, debug.",
		EnvVar: "PEGGO_LOG_LEVEL",
		Value:  "info",
	})

	*svcWaitTimeout = app.String(cli.StringOpt{
		Name:   "svc-wait-timeout",
		Desc:   "Standard wait timeout for external services (e.g. Cosmos daemon GRPC connection)",
		EnvVar: "PEGGO_SERVICE_WAIT_TIMEOUT",
		Value:  "1m",
	})
}

func initInteractiveOptions(
	cmd *cli.Cmd,
	alwaysAutoConfirm **bool,
) {
	*alwaysAutoConfirm = cmd.Bool(cli.BoolOpt{
		Name:   "y yes",
		Desc:   "Always auto-confirm actions, such as transaction sending.",
		EnvVar: "PEGGO_ALWAYS_AUTO_CONFIRM",
		Value:  false,
	})
}

func initCosmosOptions(
	cmd *cli.Cmd,
	cosmosChainID **string,
	cosmosGRPC **string,
	tendermintRPC **string,
	cosmosGasPrices **string,
) {
	*cosmosChainID = cmd.String(cli.StringOpt{
		Name:   "cosmos-chain-id",
		Desc:   "Specify Chain ID of the Cosmos network.",
		EnvVar: "PEGGO_COSMOS_CHAIN_ID",
		Value:  "888",
	})

	*cosmosGRPC = cmd.String(cli.StringOpt{
		Name:   "cosmos-grpc",
		Desc:   "Cosmos GRPC querying endpoint",
		EnvVar: "PEGGO_COSMOS_GRPC",
		Value:  "tcp://localhost:9900",
	})

	*tendermintRPC = cmd.String(cli.StringOpt{
		Name:   "tendermint-rpc",
		Desc:   "Tendermint RPC endpoint",
		EnvVar: "PEGGO_TENDERMINT_RPC",
		Value:  "http://localhost:26657",
	})

	*cosmosGasPrices = cmd.String(cli.StringOpt{
		Name:   "cosmos-gas-prices",
		Desc:   "Specify Cosmos chain transaction fees as DecCoins gas prices",
		EnvVar: "PEGGO_COSMOS_GAS_PRICES",
		Value:  "", // example: 500000000inj
	})
}

func initCosmosKeyOptions(
	cmd *cli.Cmd,
	cosmosKeyringDir **string,
	cosmosKeyringAppName **string,
	cosmosKeyringBackend **string,
	cosmosKeyFrom **string,
	cosmosKeyPassphrase **string,
	cosmosPrivKey **string,
	cosmosUseLedger **bool,
) {
	*cosmosKeyringBackend = cmd.String(cli.StringOpt{
		Name:   "cosmos-keyring",
		Desc:   "Specify Cosmos keyring backend (os|file|kwallet|pass|test)",
		EnvVar: "PEGGO_COSMOS_KEYRING",
		Value:  "file",
	})

	*cosmosKeyringDir = cmd.String(cli.StringOpt{
		Name:   "cosmos-keyring-dir",
		Desc:   "Specify Cosmos keyring dir, if using file keyring.",
		EnvVar: "PEGGO_COSMOS_KEYRING_DIR",
		Value:  "",
	})

	*cosmosKeyringAppName = cmd.String(cli.StringOpt{
		Name:   "cosmos-keyring-app",
		Desc:   "Specify Cosmos keyring app name.",
		EnvVar: "PEGGO_COSMOS_KEYRING_APP",
		Value:  "peggo",
	})

	*cosmosKeyFrom = cmd.String(cli.StringOpt{
		Name:   "cosmos-from",
		Desc:   "Specify the Cosmos validator key name or address. If specified, must exist in keyring, ledger or match the privkey.",
		EnvVar: "PEGGO_COSMOS_FROM",
	})

	*cosmosKeyPassphrase = cmd.String(cli.StringOpt{
		Name:   "cosmos-from-passphrase",
		Desc:   "Specify keyring passphrase, otherwise Stdin will be used.",
		EnvVar: "PEGGO_COSMOS_FROM_PASSPHRASE",
		Value:  "peggo",
	})

	*cosmosPrivKey = cmd.String(cli.StringOpt{
		Name:   "cosmos-pk",
		Desc:   "Provide a raw Cosmos account private key of the validator in hex. USE FOR TESTING ONLY!",
		EnvVar: "PEGGO_COSMOS_PK",
	})

	*cosmosUseLedger = cmd.Bool(cli.BoolOpt{
		Name:   "cosmos-use-ledger",
		Desc:   "Use the Cosmos app on hardware ledger to sign transactions.",
		EnvVar: "PEGGO_COSMOS_USE_LEDGER",
		Value:  false,
	})
}

func initEthereumOptions(
	cmd *cli.Cmd,
	ethChainID **int,
	ethNodeRPC **string,
	ethGasPriceAdjustment **float64,
) {
	*ethChainID = cmd.Int(cli.IntOpt{
		Name:   "eth-chain-id",
		Desc:   "Specify Chain ID of the Ethereum network.",
		EnvVar: "PEGGO_ETH_CHAIN_ID",
		Value:  42,
	})

	*ethNodeRPC = cmd.String(cli.StringOpt{
		Name:   "eth-node-http",
		Desc:   "Specify HTTP endpoint for an Ethereum node.",
		EnvVar: "PEGGO_ETH_RPC",
		Value:  "http://localhost:1317",
	})

	*ethGasPriceAdjustment = cmd.Float64(cli.Float64Opt{
		Name:   "eth_gas_price_adjustment",
		Desc:   "gas price adjustment for Ethereum transactions",
		EnvVar: "PEGGO_ETH_GAS_PRICE_ADJUSTMENT",
		Value:  float64(1.3),
	})
}

func initEthereumKeyOptions(
	cmd *cli.Cmd,
	ethKeystoreDir **string,
	ethKeyFrom **string,
	ethPassphrase **string,
	ethPrivKey **string,
	ethUseLedger **bool,
) {
	*ethKeystoreDir = cmd.String(cli.StringOpt{
		Name:   "eth-keystore-dir",
		Desc:   "Specify Ethereum keystore dir (Geth-format) prefix.",
		EnvVar: "PEGGO_ETH_KEYSTORE_DIR",
	})

	*ethKeyFrom = cmd.String(cli.StringOpt{
		Name:   "eth-from",
		Desc:   "Specify the from address. If specified, must exist in keystore, ledger or match the privkey.",
		EnvVar: "PEGGO_ETH_FROM",
	})

	*ethPassphrase = cmd.String(cli.StringOpt{
		Name:   "eth-passphrase",
		Desc:   "Passphrase to unlock the private key from armor, if empty then stdin is used.",
		EnvVar: "PEGGO_ETH_PASSPHRASE",
	})

	*ethPrivKey = cmd.String(cli.StringOpt{
		Name:   "eth-pk",
		Desc:   "Provide a raw Ethereum private key of the validator in hex. USE FOR TESTING ONLY!",
		EnvVar: "PEGGO_ETH_PK",
	})

	*ethUseLedger = cmd.Bool(cli.BoolOpt{
		Name:   "eth-use-ledger",
		Desc:   "Use the Ethereum app on hardware ledger to sign transactions.",
		EnvVar: "PEGGO_ETH_USE_LEDGER",
		Value:  false,
	})
}

// initStatsdOptions sets options for StatsD metrics.

// initRelayerOption sets options for relayer.
func initRelayerOptions(
	cmd *cli.Cmd,
	relayValsets **bool,
	relayBatches **bool,
) {
	*relayValsets = cmd.Bool(cli.BoolOpt{
		Name:   "relay_valsets",
		Desc:   "If enabled, relayer will relay valsets to ethereum",
		EnvVar: "PEGGO_RELAY_VALSETS",
		Value:  false,
	})

	*relayBatches = cmd.Bool(cli.BoolOpt{
		Name:   "relay_batches",
		Desc:   "If enabled, relayer will relay batches to ethereum",
		EnvVar: "PEGGO_RELAY_BATCHES",
		Value:  false,
	})
}

// initBatchRequesterOptions sets options for batch requester.
func initBatchRequesterOptions(
	cmd *cli.Cmd,
	minBatchFeeUSD **float64,
) {
	*minBatchFeeUSD = cmd.Float64(cli.Float64Opt{
		Name:   "min_batch_fee_usd",
		Desc:   "If set, batch request will create batches only if fee threshold exceeds",
		EnvVar: "PEGGO_MIN_BATCH_FEE_USD",
		Value:  float64(23.3),
	})
}

// initCoingeckoOptions sets options for coingecko.
func initCoingeckoOptions(
	cmd *cli.Cmd,
	baseUrl **string,
) {
	*baseUrl = cmd.String(cli.StringOpt{
		Name:   "coingecko_api",
		Desc:   "Specify HTTP endpoint for coingecko api.",
		EnvVar: "PEGGO_COINGECKO_API",
		Value:  "https://api.coingecko.com/api/v3",
	})
}
