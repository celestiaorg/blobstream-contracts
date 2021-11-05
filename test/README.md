## Peggo Testsuite

Welcome to the PegGo testing framework. The goal of this suite is aligned with the overall project goal - to move stuff onto common ground and iterate faster.
By using the same lang for module, orchestrator and test we can achieve the full test coverage of all logical branches.

This is a special place where we don't care about things like:
* Node version
* Myriads of JS packages complaining about versions and API inconsistencies
* Different locations of ERC20 contract artifacts
* Stuff being deployed slowly
* Debugging ganache bugs

We care about:
* Speed of the full run
* Go code coverage reports
* Zero issues coming from tooling or dev env
* Cross-platformity (macOS youKnow)
* Supporting any target EVM that implements Ethereum JSON-RPC
* 100% compatibility with real network

## Prerequisites

You can specify any remote EVM endpoint to run the test against, but the best and most stable way to test the stuff is to run a Ganache or Hardhat instance. Hardhat is used solely as a JSON-RPC node provider.

Preferred Solc compiler toolkit:
* https://github.com/crytic/solc-select

Run `solc-select use 0.8.2` before starting any tests.

### Running with Hardhat

Hardhat is a newer alternative to Ganache that has convenient initialization via the config file.

Running the init script will install node_modules inside `./test/ethereum` dir.

```
$ ./test/ethereum/hardhat-init.sh
```

After init is done, the following command can be used to launch a Hardhat server instance:

```
$ ./test/ethereum/hardhat.sh
```

The only option that can be set via ENV variable:

* `HARDHAT_PORT` - specify the port for server to listen on. Defaults to `8545`.

The rest of the options can be tweaked via `./test/ethereum/hardhat.config.js`

### Pre-prod testing with Geth

In order to get maximum compatibility with the real blockchain environment and avoid any bugs in the EVM runtime of Hardhat/Ganache,
also check different blocktime conditions, one might want to run Geth itself.

Running this script will init a persistent data storage for the private network.

```
$ ./test/ethereum/geth-init.sh
```

Init options can be set via ENV variables:

* `GETH_NETWORK_ID` - specify Ethereum Network ID, defaults to `50`.
* `GETH_ALGO` - specify the consensus algorith for block producing. Defaults to `clique` (PoA), but `ethash` (PoW) is supported too. Make sure you adjust difficulty by patching your Geth (see at the bottom of this page)
* `CHAIN_DIR` - specify the data dir, a prefix for all data dirs and logs. Defaults to `./data`

Chain options can be tweaked in `./test/ethereum/geth/genesis.json`

After init is done, the following command can be used to launch a full Geth node instance:

```
$ ./test/ethereum/geth.sh
```

Running options can be set via ENV variables:

* `GETH_NETWORK_ID` - specify Ethereum Network ID, defaults to `50`.
* `GETH_ALGO` - specify the consensus algorith for block producing. Defaults to `ethash` (PoW), but `clique` (PoA) supported.
* `GETH_PORT` - specify the port for server to listen on. Defaults to `8545`.
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

### Misc: Patching Geth

Geth by default scales difficulty of the blocks to hit the target block pace. So even if your network starts with `difficulty=1` in genesis, the difficulty will be higher in the next blocks and waiting times would be very high. Especially that DAG regeneration phases. A solution to this in local setup would be to either use `clique` consensus for PoA-style block producing, or just patch the Geth code, so the difficulty won't grow.

Just clone the `go-ethereum` repo, apply this patch:

```diff
diff --git a/consensus/ethash/consensus.go b/consensus/ethash/consensus.go
index bdc02098a..c17ea5b76 100644
--- a/consensus/ethash/consensus.go
+++ b/consensus/ethash/consensus.go
@@ -315,19 +315,7 @@ func (ethash *Ethash) CalcDifficulty(chain consensus.ChainHeaderReader, time uin
 // the difficulty that a new block should have when created at time
 // given the parent block's time and difficulty.
 func CalcDifficulty(config *params.ChainConfig, time uint64, parent *types.Header) *big.Int {
-       next := new(big.Int).Add(parent.Number, big1)
-       switch {
-       case config.IsMuirGlacier(next):
-               return calcDifficultyEip2384(time, parent)
-       case config.IsConstantinople(next):
-               return calcDifficultyConstantinople(time, parent)
-       case config.IsByzantium(next):
-               return calcDifficultyByzantium(time, parent)
-       case config.IsHomestead(next):
-               return calcDifficultyHomestead(time, parent)
-       default:
-               return calcDifficultyFrontier(time, parent)
-       }
+       return big1
 }
```

And install it with `go install ./cmd/geth`. Welcome to the Geth forking!

## Contributing

Patches and suggestions are welcome. We're looking for better coverage and maybe some isolated benchmarks.

🍻