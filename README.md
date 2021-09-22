# Injective's Peggo [![Peggy.sol MythX](https://badgen.net/https/api.mythx.io/v1/projects/82ca9468-f86d-4550-a0ae-bc120eeb055f/badge/data?cache=300&icon=https://raw.githubusercontent.com/ConsenSys/mythx-github-badge/main/logo_white.svg)](https://docs.mythx.io/dashboard/github-badges)

Peggo is a Go implementation of the Peggy Orchestrator for the Injective Chain. 

Important Commands:

* `peggo orchestrator` starts the orchestrator main loop.
* `peggo tx register-eth-key` is a special command to submit an Ethereum key that will be used to sign messages on behalf of your Validator


## Installation

Get yourself `Go 1.15+` at https://golang.org/dl/ first, then:

```
$ go get github.com/InjectiveLabs/peggo/orchestrator/cmd/...
```

## peggo

Peggo is a companion executable for orchestrating a Peggy validator.

### Configuration

Use CLI args, flags or create `.env` with environment variables

### Usage

```
$ peggo --help

Usage: peggo [OPTIONS] COMMAND [arg...]

Peggo is a companion executable for orchestrating a Peggy validator.

Options:
  -e, --env                The environment name this app runs in. Used for metrics and error reporting. (env $PEGGO_ENV) (default "local")
  -l, --log-level          Available levels: error, warn, info, debug. (env $PEGGO_LOG_LEVEL) (default "info")
      --svc-wait-timeout   Standard wait timeout for external services (e.g. Cosmos daemon GRPC connection) (env $PEGGO_SERVICE_WAIT_TIMEOUT) (default "1m")

Commands:
  orchestrator             Starts the orchestrator main loop.
  q, query                 Query commands that can get state info from Peggy.
  tx                       Transactions for Peggy governance and maintenance.
  version                  Print the version information and exit.

Run 'peggo COMMAND --help' for more information on a command.      
```

## Commands

### peggo orchestrator

```
$ peggo orchestrator -h

Usage: peggo orchestrator [OPTIONS]

Starts the orchestrator main loop.

Options:
      --cosmos-chain-id                  Specify Chain ID of the Cosmos network. (env $PEGGO_COSMOS_CHAIN_ID) (default "888")
      --cosmos-grpc                      Cosmos GRPC querying endpoint (env $PEGGO_COSMOS_GRPC) (default "tcp://localhost:9900")
      --tendermint-rpc                   Tendermint RPC endpoint (env $PEGGO_TENDERMINT_RPC) (default "http://localhost:26657")
      --cosmos-gas-prices                Specify Cosmos chain transaction fees as DecCoins gas prices (env $PEGGO_COSMOS_GAS_PRICES)
      --cosmos-keyring                   Specify Cosmos keyring backend (os|file|kwallet|pass|test) (env $PEGGO_COSMOS_KEYRING) (default "file")
      --cosmos-keyring-dir               Specify Cosmos keyring dir, if using file keyring. (env $PEGGO_COSMOS_KEYRING_DIR)
      --cosmos-keyring-app               Specify Cosmos keyring app name. (env $PEGGO_COSMOS_KEYRING_APP) (default "peggo")
      --cosmos-from                      Specify the Cosmos validator key name or address. If specified, must exist in keyring, ledger or match the privkey. (env $PEGGO_COSMOS_FROM)
      --cosmos-from-passphrase           Specify keyring passphrase, otherwise Stdin will be used. (env $PEGGO_COSMOS_FROM_PASSPHRASE) (default "peggo")
      --cosmos-pk                        Provide a raw Cosmos account private key of the validator in hex. USE FOR TESTING ONLY! (env $PEGGO_COSMOS_PK)
      --cosmos-use-ledger                Use the Cosmos app on hardware ledger to sign transactions. (env $PEGGO_COSMOS_USE_LEDGER)
      --eth-chain-id                     Specify Chain ID of the Ethereum network. (env $PEGGO_ETH_CHAIN_ID) (default 42)
      --eth-node-http                    Specify HTTP endpoint for an Ethereum node. (env $PEGGO_ETH_RPC) (default "http://localhost:1317")
      --eth-node-alchemy-ws              Specify websocket url for an Alchemy ethereum node. (env $PEGGO_ETH_ALCHEMY_WS)
      --eth_gas_price_adjustment         gas price adjustment for Ethereum transactions (env $PEGGO_ETH_GAS_PRICE_ADJUSTMENT) (default 1.3)
      --eth-keystore-dir                 Specify Ethereum keystore dir (Geth-format) prefix. (env $PEGGO_ETH_KEYSTORE_DIR)
      --eth-from                         Specify the from address. If specified, must exist in keystore, ledger or match the privkey. (env $PEGGO_ETH_FROM)
      --eth-passphrase                   Passphrase to unlock the private key from armor, if empty then stdin is used. (env $PEGGO_ETH_PASSPHRASE)
      --eth-pk                           Provide a raw Ethereum private key of the validator in hex. USE FOR TESTING ONLY! (env $PEGGO_ETH_PK)
      --eth-use-ledger                   Use the Ethereum app on hardware ledger to sign transactions. (env $PEGGO_ETH_USE_LEDGER)
      --relay_valsets                    If enabled, relayer will relay valsets to ethereum (env $PEGGO_RELAY_VALSETS)
      --relay_valset_offset_dur          If set, relayer will broadcast valsetUpdate only after relayValsetOffsetDur has passed from time of valsetUpdate creation (env $PEGGO_RELAY_VALSET_OFFSET_DUR) (default "5m")
      --relay_batches                    If enabled, relayer will relay batches to ethereum (env $PEGGO_RELAY_BATCHES)
      --relay_batch_offset_dur           If set, relayer will broadcast batches only after relayBatchOffsetDur has passed from time of batch creation (env $PEGGO_RELAY_BATCH_OFFSET_DUR) (default "5m")
      --relay_pending_tx_wait_duration   If set, relayer will broadcast pending batches/valsetupdate only after pendingTxWaitDuration has passed (env $PEGGO_RELAY_PENDING_TX_WAIT_DURATION) (default "20m")
      --min_batch_fee_usd                If set, batch request will create batches only if fee threshold exceeds (env $PEGGO_MIN_BATCH_FEE_USD) (default 23.3)
      --coingecko_api                    Specify HTTP endpoint for coingecko api. (env $PEGGO_COINGECKO_API) (default "https://api.coingecko.com/api/v3")

```

### peggo tx register-eth-key

```
 peggo tx register-eth-key --help

Usage: peggo tx register-eth-key [OPTIONS]

Submits an Ethereum key that will be used to sign messages on behalf of your Validator

Options:
      --cosmos-chain-id          Specify Chain ID of the Cosmos network. (env $PEGGO_COSMOS_CHAIN_ID) (default "888")
      --cosmos-grpc              Cosmos GRPC querying endpoint (env $PEGGO_COSMOS_GRPC) (default "tcp://localhost:9900")
      --tendermint-rpc           Tendermint RPC endpoint (env $PEGGO_TENDERMINT_RPC) (default "http://localhost:26657")
      --cosmos-gas-prices        Specify Cosmos chain transaction fees as DecCoins gas prices (env $PEGGO_COSMOS_GAS_PRICES)
      --cosmos-keyring           Specify Cosmos keyring backend (os|file|kwallet|pass|test) (env $PEGGO_COSMOS_KEYRING) (default "file")
      --cosmos-keyring-dir       Specify Cosmos keyring dir, if using file keyring. (env $PEGGO_COSMOS_KEYRING_DIR)
      --cosmos-keyring-app       Specify Cosmos keyring app name. (env $PEGGO_COSMOS_KEYRING_APP) (default "peggo")
      --cosmos-from              Specify the Cosmos validator key name or address. If specified, must exist in keyring, ledger or match the privkey. (env $PEGGO_COSMOS_FROM)
      --cosmos-from-passphrase   Specify keyring passphrase, otherwise Stdin will be used. (env $PEGGO_COSMOS_FROM_PASSPHRASE) (default "peggo")
      --cosmos-pk                Provide a raw Cosmos account private key of the validator in hex. USE FOR TESTING ONLY! (env $PEGGO_COSMOS_PK)
      --cosmos-use-ledger        Use the Cosmos app on hardware ledger to sign transactions. (env $PEGGO_COSMOS_USE_LEDGER)
      --eth-keystore-dir         Specify Ethereum keystore dir (Geth-format) prefix. (env $PEGGO_ETH_KEYSTORE_DIR)
      --eth-from                 Specify the from address. If specified, must exist in keystore, ledger or match the privkey. (env $PEGGO_ETH_FROM)
      --eth-passphrase           Passphrase to unlock the private key from armor, if empty then stdin is used. (env $PEGGO_ETH_PASSPHRASE)
      --eth-pk                   Provide a raw Ethereum private key of the validator in hex. USE FOR TESTING ONLY! (env $PEGGO_ETH_PK)
      --eth-use-ledger           Use the Ethereum app on hardware ledger to sign transactions. (env $PEGGO_ETH_USE_LEDGER)
  -y, --yes                      Always auto-confirm actions, such as transaction sending. (env $PEGGO_ALWAYS_AUTO_CONFIRM)
```

## License

Apache 2.0
