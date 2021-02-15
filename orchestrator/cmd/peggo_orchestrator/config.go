package main

import (
	"strings"
	"time"

	cli "github.com/jawher/mow.cli"
	log "github.com/xlab/suplog"
)

var (
	envName            *string
	appLogLevel        *string
	svcWaitTimeout     *string
	cosmosPrivkey      *string
	cosmosGRPC         *string
	tendermintRPC      *string
	feeDenom           *string
	chainId            *string
	ethNodeRPC         *string
	ethPrivKey         *string
	contractAddrHex    *string
	injContractAddrHex *string
	statsdPrefix       *string
	statsdAddr         *string
	statsdStuckDur     *string
	statsdMocking      *string
	statsdDisabled     *string
)

func initFlags() {
	envName = app.String(cli.StringOpt{
		Name:   "env",
		Desc:   "Application environment",
		EnvVar: "APP_ENV",
		Value:  "local",
	})

	appLogLevel = app.String(cli.StringOpt{
		Name:   "l log-level",
		Desc:   "Available levels: error, warn, info, debug.",
		EnvVar: "APP_LOG_LEVEL",
		Value:  "info",
	})

	svcWaitTimeout = app.String(cli.StringOpt{
		Name:   "svc-wait-timeout",
		Desc:   "Standard wait timeout for all service dependencies (e.g. injectived).",
		EnvVar: "SERVICE_WAIT_TIMEOUT",
		Value:  "1m",
	})

	cosmosPrivkey = app.String(cli.StringOpt{
		Name:   "cosmos-privkey",
		Desc:   "The Cosmos private key of the validator.",
		EnvVar: "PEGGY_COSMOS_PRIVKEY",
	})

	cosmosGRPC = app.String(cli.StringOpt{
		Name:   "cosmos-grpc",
		Desc:   "Cosmos GRPC querying endpoint",
		EnvVar: "PEGGY_COSMOS_GRPC",
		Value:  "tcp://localhost:9900",
	})

	tendermintRPC = app.String(cli.StringOpt{
		Name:   "tendermint-rpc",
		Desc:   "Tednermint RPC endpoint",
		EnvVar: "PEGGY_TENDERMINT_RPC",
		Value:  "http://localhost:26657",
	})

	feeDenom = app.String(cli.StringOpt{
		Name:   "fees",
		Desc:   "The Cosmos Denom in which to pay Cosmos chain fees",
		EnvVar: "PEGGY_FEE_DENOM",
		Value:  "inj",
	})

	chainId = app.String(cli.StringOpt{
		Name:   "chain-id",
		Desc:   "Specify Chain ID of the injectived service.",
		EnvVar: "INJECTIVED_CHAIN_ID",
		Value:  "888",
	})

	ethNodeRPC = app.String(cli.StringOpt{
		Name:   "eth-node-http",
		Desc:   "Specify HTTP endpoint for an Ethereum node.",
		EnvVar: "PEGGY_ETH_RPC",
		Value:  "http://localhost:1317",
	})

	ethPrivKey = app.String(cli.StringOpt{
		Name:   "eth-privkey",
		Desc:   "The Ethereum private key of the validator(Ex: 5D862464FE95...)",
		EnvVar: "PEGGY_ETH_PRIVATE_KEY",
	})

	contractAddrHex = app.String(cli.StringOpt{
		Name:   "contract-address",
		Desc:   "The Ethereum contract address of Peggy",
		EnvVar: "PEGGY_CONTRACT_ADDRESS",
	})

	injContractAddrHex = app.String(cli.StringOpt{
		Name:   "inj-contract-address",
		Desc:   "The Ethereum contract address of INJ Erc20",
		EnvVar: "INJ_CONTRACT_ADDRESS",
	})

	statsdPrefix = app.String(cli.StringOpt{
		Name:   "statsd-prefix",
		Desc:   "Specify StatsD compatible metrics prefix.",
		EnvVar: "STATSD_PREFIX",
		Value:  "relayer_api",
	})

	statsdAddr = app.String(cli.StringOpt{
		Name:   "statsd-addr",
		Desc:   "UDP address of a StatsD compatible metrics aggregator.",
		EnvVar: "STATSD_ADDR",
		Value:  "localhost:8125",
	})

	statsdStuckDur = app.String(cli.StringOpt{
		Name:   "statsd-stuck-func",
		Desc:   "Sets a duration to consider a function to be stuck (e.g. in deadlock).",
		EnvVar: "STATSD_STUCK_DUR",
		Value:  "5m",
	})

	statsdMocking = app.String(cli.StringOpt{
		Name:   "statsd-mocking",
		Desc:   "If enabled replaces statsd client with a mock one that simply logs values.",
		EnvVar: "STATSD_MOCKING",
		Value:  "false",
	})

	statsdDisabled = app.String(cli.StringOpt{
		Name:   "statsd-disabled",
		Desc:   "Force disabling statsd reporting completely.",
		EnvVar: "STATSD_DISABLED",
		Value:  "false",
	})
}

func Level(s string) log.Level {
	switch s {
	case "1", "error":
		return log.ErrorLevel
	case "2", "warn":
		return log.WarnLevel
	case "3", "info":
		return log.InfoLevel
	case "4", "debug":
		return log.DebugLevel
	default:
		return log.FatalLevel
	}
}

func toBool(s string) bool {
	switch strings.ToLower(s) {
	case "true", "1", "t", "yes":
		return true
	default:
		return false
	}
}

func duration(s string, defaults time.Duration) time.Duration {
	dur, err := time.ParseDuration(s)
	if err != nil {
		dur = defaults
	}
	return dur
}

func checkStatsdPrefix(s string) string {
	if !strings.HasSuffix(s, ".") {
		return s + "."
	}
	return s
}
