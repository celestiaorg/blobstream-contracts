# Peggo

<!-- markdownlint-disable MD013 MD041 -->

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://godoc.org/github.com/celestiaorg/quantum-gravity-bridge)
[![Go Report Card](https://goreportcard.com/badge/github.com/celestiaorg/quantum-gravity-bridge?style=flat-square)](https://goreportcard.com/report/github.com/celestiaorg/quantum-gravity-bridge)
[![Version](https://img.shields.io/github/tag/celestiaorg/quantum-gravity-bridge.svg?style=flat-square)](https://github.com/celestiaorg/quantum-gravity-bridge/releases/latest)
[![License: Apache-2.0](https://img.shields.io/github/license/celestiaorg/quantum-gravity-bridge.svg?style=flat-square)](https://github.com/celestiaorg/quantum-gravity-bridge/blob/main/LICENSE)
[![GitHub Super-Linter](https://img.shields.io/github/workflow/status/celestiaorg/quantum-gravity-bridge/Lint?style=flat-square&label=Lint)](https://github.com/marketplace/actions/super-linter)

Peggo is a Go implementation of the Peggy (Gravity Bridge) Orchestrator originally
implemented by [Injective Labs](https://github.com/InjectiveLabs/). Peggo itself
is a fork of the original Gravity Bridge Orchestrator implemented by [Althea](https://github.com/althea-net).

## Table of Contents

- [Dependencies](#dependencies)
- [Installation](#installation)
- [How to run](#how-to-run)
- [How it works](#how-it-works)

## Dependencies

- [Go 1.17+](https://golang.org/dl/)

## Installation

To install the `peggo` binary:

```shell
$ make install
```

## How to run

### Setup

First we must register the validator's Ethereum key. This key will be used to
sign claims going from Ethereum to Umee and to sign any transactions sent to
Ethereum (batches or validator set updates).

```shell
$ peggo tx register-eth-key \
  --cosmos-chain-id="..." \
  --cosmos-grpc="tcp://..." \
  --tendermint-rpc="http://..." \
  --cosmos-keyring=... \
  --cosmos-keyring-dir=... \
  --cosmos-from=... \
  --eth-pk=$ETH_PK
```

### Run the orchestrator

```shell
$ peggo orchestrator \
  --eth-pk=$ETH_PK \
  --eth-rpc=$ETH_RPC \
  --relay-batches=true \
  --relay-valsets=true \
  --eth-chain-id=... \
  --cosmos-chain-id=... \
  --cosmos-grpc="tcp://..." \
  --tendermint-rpc="http://..." \
  --cosmos-keyring=... \
  --cosmos-keyring-dir=... \
  --cosmos-from=...
```

### Send a transfer from Umee to Ethereum

This is done using the command `umeed tx peggy send-to-eth`, use the `--help`
flag for more information.

If the coin doesn't have a corresponding ERC20 equivalent on the Ethereum
network, the transaction will fail. This is only required for Cosmos originated
coins and anyone can call the `deployERC20` function on the Peggy contract to
fix this (Peggo has a helper command for this, see
`peggo bridge deploy-erc20 --help` for more details).

This process takes longer than transfers the other way around because they get
relayed in batches rather than individually. It primarily depends on the amount
of transfers of the same token and the fees the senders are paying.

Important notice: if an "unlisted" (with no monetary value) ERC20 token gets
sent into Umee it won't be possible to transfer it back to Ethereum, unless a
validator is configured to batch and relay transactions of this token.

### Send a transfer from Ethereum to Umee

Any ERC20 token can be sent to Umee and it's done using the command
`peggo bridge send-to-cosmos`, use the `--help` flag for more information. It
can also be done by calling the `sendToCosmos` method on the Peggy contract.

The ERC20 tokens will be locked in the Peggy contract and new coins will be
minted on Umee with the denomination `peggy{token_address}`. This process takes
around 3 minutes or 12 Ethereum blocks.

## How it works

Peggo allows transfers of assets back and forth between Ethereum and Umee.
It supports both assets originating on Umee and assets originating on Ethereum
(any ERC20 token).

It works by scanning the events of the contract deployed on Ethereum (Peggy) and
relaying them as messages to the Umee chain; and relaying transaction batches and
validator sets from Umee to Ethereum.

### Events and messages observed/relayed

#### Ethereum

**Deposits** (`SendToCosmosEvent`): emitted when sending tokens from Ethereum to
Umee using the `sendToCosmos` function on Peggy.

**Withdraw** (`TransactionBatchExecutedEvent`): emitted when a batch of
transactions is sent from Umee to Ethereum using the `submitBatch` function on
the Peggy contract by a validator. This serves as a confirmation to Umee that
the batch was sent successfully.

**Valset update** (`ValsetUpdatedEvent`): emitted on init of the Peggy contract
and on every execution of the `updateValset` function.

**Deployed ERC 20** (`ERC20DeployedEvent`): emitted when executing the function
`deployERC20`. This event signals Umee that there's a new ERC20 deployed from
Peggy, so Umee can map the token contract address to the corresponding native
coin. This enables transfers from Umee to Ethereum.

#### Umee

 **Validator sets**: Umee informs the Peggy contract who are the current
 validators and their power. This results in an execution of the `updateValset`
 function.

 **Request batch**: Peggo will check for new transactions in the Outgoing TX Pool
 and if the transactions' fees are greater than the set minimum batch fee, it
 will send a message to Umee requesting a new batch.

 **Batches**: Peggo queries Umee for any batches ready to be relayed and relays
 them over to Ethereum using the `submitBatch` function on the Peggy contract.
