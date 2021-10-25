# Injective's Peggo [![Peggy.sol MythX](https://badgen.net/https/api.mythx.io/v1/projects/82ca9468-f86d-4550-a0ae-bc120eeb055f/badge/data?cache=300&icon=https://raw.githubusercontent.com/ConsenSys/mythx-github-badge/main/logo_white.svg)](https://docs.mythx.io/dashboard/github-badges)

Peggo is a Go implementation of the Peggy Orchestrator for the Injective Chain.

`orchestrator` package provides all required components, while `orchestrator/cmd` has exectables that run.

List of executables:

* `peggo_orchestrator` is the main Validator companion binary for Peggy.
* `register_eth_key` is a special purpose binary for bootstrapping Peggy chains.

## Installation

Get yourself `Go 1.15+` at https://golang.org/dl/ first, then:

```
$ go get github.com/umee-network/peggo/orchestrator/cmd/...
```

## peggo_orchestrator

### Configuration

Use CLI args, flags or create `.env` with environment variables

### Usage

```
$ peggo_orchestrator -h

Usage: peggo_orchestrator [OPTIONS] COMMAND [arg...]

The Validator companion binary for Peggy.

Options:
      --env                 Application environment (env $APP_ENV) (default "local")
  -l, --log-level           Available levels: error, warn, info, debug. (env $APP_LOG_LEVEL) (default "info")
      --svc-wait-timeout    Standard wait timeout for all service dependencies (e.g. injectived). (env $SERVICE_WAIT_TIMEOUT) (default "1m")
      --cosmos-privkey      The Cosmos private key of the validator. (env $PEGGY_COSMOS_PRIVKEY)
      --cosmos-grpc         Cosmos GRPC querying endpoint (env $PEGGY_COSMOS_GRPC) (default "tcp://localhost:9900")
      --tendermint-rpc      Tednermint RPC endpoint (env $PEGGY_TENDERMINT_RPC) (default "http://localhost:26657")
      --fees                The Cosmos Denom in which to pay Cosmos chain fees (env $PEGGY_FEE_DENOM) (default "inj")
      --chain-id            Specify Chain ID of the injectived service. (env $INJECTIVED_CHAIN_ID) (default "888")
      --eth-node-http       Specify HTTP endpoint for an Ethereum node. (env $PEGGY_ETH_RPC) (default "http://localhost:1317")
      --eth-privkey         The Ethereum private key of the validator(Ex: 5D862464FE95...) (env $PEGGY_ETH_PRIVATE_KEY)
      --contract-address    The Ethereum contract address of Peggy (env $PEGGY_CONTRACT_ADDRESS)
      --statsd-prefix       Specify StatsD compatible metrics prefix. (env $STATSD_PREFIX) (default "relayer_api")
      --statsd-addr         UDP address of a StatsD compatible metrics aggregator. (env $STATSD_ADDR) (default "localhost:8125")
      --statsd-stuck-func   Sets a duration to consider a function to be stuck (e.g. in deadlock). (env $STATSD_STUCK_DUR) (default "5m")
      --statsd-mocking      If enabled replaces statsd client with a mock one that simply logs values. (env $STATSD_MOCKING) (default "false")
      --statsd-disabled     Force disabling statsd reporting completely. (env $STATSD_DISABLED) (default "false")

Commands:
  version                   Print the version information and exit.

Run 'peggo_orchestrator COMMAND --help' for more information on a command.
```

## register_eth_key

### Configuration

Use CLI args, flags or create `.env` with environment variables

### Usage

```
$ register_eth_key -h

Usage: register_eth_key [OPTIONS]

Special purpose binary for bootstrapping Peggy chains.

Options:
      --cosmos-privkey   The Cosmos private key of the validator. Must be saved when you generate your key (env $PEGGY_COSMOS_PRIVKEY)
      --cosmos-grpc      Cosmos GRPC querying endpoint (env $PEGGY_COSMOS_GRPC) (default "tcp://localhost:9900")
      --fees             The Cosmos Denom in which to pay Cosmos chain fees (env $PEGGY_FEE_DENOM) (default "inj")
      --chain-id         Specify Chain ID of the injectived service. (env $INJECTIVED_CHAIN_ID) (default "888")
```

## License

Apache 2.0
