## Peggo Testsuite

Welcome to the PegGo testing framework. The goal of this suite is aligned with the overall project goal - to move stuff onto common ground and iterate faster.
By using the same lang for module, orchestrator and test we can achieve the full test coverage of all logical branches.

This is a special place where we don't care about things like:
* Node version
* Myriads of JS packages complaining about versions and API inconsistencies
* Different locations of ERC20 contract artifacts
* Stuff being deployed slowly
* Forking the Eth mainnet

We care about:
* Speed of the full run
* Code coverage reports
* Zero issues coming from tooling or dev env
* Cross-platformity (macOS youKnow)
* Supporting any target EVM that implements Ethereum JSON-RPC

## Prerequisites

You can specify any remote EVM endpoint to run the test against, but the best and most stable way to trest the stuff is to run a Ganache instance. Or hardhat, if you prefer, but it's harder to setup and gives no benefits in this case.

To setup Ganache with UI just go to:
* https://www.trufflesuite.com/ganache

Or, if you prefer CLI approach:
```
$ yarn global add ganache-cli

[1/4] 🔍  Resolving packages...
[2/4] 🚚  Fetching packages...
[3/4] 🔗  Linking dependencies...
[4/4] 🔨  Building fresh packages...
success Installed "ganache-cli@6.12.2" with binaries:
      - ganache-cli
```

Preferred Solc compiler toolkit:
* https://github.com/crytic/solc-select

Run `solc select 0.6.6` before starting any tests.

### Running Ethereum

When having Ganache CLI installed, run the following command to download a snapshot and init the data dir:

```
$ ./test/ethereum/ganache-init.sh
```

Options can be set via ENV variables:

* `GANACHE_NETWORK_ID` - specify Ethereum Network ID, defaults to `50`.
* `CHAIN_DIR` - specify the data dir, a prefix for all data dirs and logs. Defaults to `./data`

After init is done, the following command can be used to launch a Ganache instance:

```
$ ./test/ethereum/ganache.sh
```

Options can be set via ENV variables:

* `GANACHE_NETWORK_ID` - specify Ethereum Network ID, defaults to `50`.
* `CHAIN_DIR` - specify the data dir, a prefix for all data dirs and logs. Defaults to `./data`

### Cosmos Daemon

This testsuite supports different Cosmos backends, basically any app that has `peggy` module built-in will do. We expect that the generic app is Cosmos-SDK compatible and has very similar CLI interface to `gaiad`. There is a script that would launch an isolated full 3-node network, running natively on the host machine. Just make sure to provide the target executable as an argument.

```
$ CHAIN_ID=888 DENOM=inj ./test/cosmos/multinode.sh injectived
$ CHAIN_ID=somm DENOM=samoleans STAKE_DENOM=stake SCALE_FACTOR=000000 ./test/cosmos/multinode.sh sommelier
```

Full list of the supported ENV variables:
* `CHAIN_ID` - specifies Cosmos Chain ID, like `peggy-1`
* `CHAIN_DIR` - is a prefix for all data dirs and logs, will be removed if `CLEANUP=1`
* `DENOM` - Cosmos coin denom, the default coin of the network. Examples: `uatom`, `aphoton`, `samoleans` etc
* `STAKE_DENOM` - Cosmos coin denom that is used for staking and governance. On the Cosmos Hub it's `stake`. Defaults to value of `DENOM` in the script.
* `SCALE_FACTOR` - Scale factor for the Cosmos coin. Defaults to 1e18 to reflect Ethereum token balances. Use `000000` to follow Cosmos uatom (1e6) style.
* `CLEANUP` - if this option set to `1`, then the `CHAIN_DIR` will be removed in the most unsafe manner.
* `LOG_LEVEL` - sets the log level of the Cosmos node configuration. Defaults to Cosmos' default (`main:info,state:info,statesync:info,*:error`).

**Important**: it is safe to run the script multiple times, it will stop nodes upon running, and optionally cleanup the state. If the state is not empty, the script will start nodes without running initialization again. So it could be used for manually retriable tests.

### Cosmos Accounts

The script imports 3 validator accounts and 1 user account, specified by mnemonics in the script itself. Each validator account accessible as `val` on the corresponding nodes, and user account is shared across all three nodes as `user`.

## Contributing

Patches and suggestions are welcome. We're looking for better coverage and maybe some isolated benchmarks.

🍻