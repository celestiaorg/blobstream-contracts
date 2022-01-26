# Quantum Gravity Bridge

<!-- markdownlint-disable MD013 MD041 -->

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://godoc.org/github.com/celestiaorg/quantum-gravity-bridge)
[![Go Report Card](https://goreportcard.com/badge/github.com/celestiaorg/quantum-gravity-bridge?style=flat-square)](https://goreportcard.com/report/github.com/celestiaorg/quantum-gravity-bridge)
[![Version](https://img.shields.io/github/tag/celestiaorg/quantum-gravity-bridge.svg?style=flat-square)](https://github.com/celestiaorg/quantum-gravity-bridge/releases/latest)
[![License: Apache-2.0](https://img.shields.io/github/license/celestiaorg/quantum-gravity-bridge.svg?style=flat-square)](https://github.com/celestiaorg/quantum-gravity-bridge/blob/main/LICENSE)
[![GitHub Super-Linter](https://img.shields.io/github/workflow/status/celestiaorg/quantum-gravity-bridge/Lint?style=flat-square&label=Lint)](https://github.com/marketplace/actions/super-linter)

The Quantum Gravity Bridge (QGB) is a Celestia -> EVM message relay.
It is based on Umee's Gravity Bridge implementation, [Peggo](https://github.com/umee-network/peggo).
**This project is under active development and should not be used in production**.

## Table of Contents

- [Dependencies](#dependencies)
- [How to run](#how-to-run)
- [How it works](#how-it-works)

## Dependencies

- [Go 1.17+](https://golang.org/dl/)

<!--
## Installation

To install the `qgb` binary:

```console
$ make install
```
-->

## How to run

<!--
### Setup

First we must register the validator's Ethereum key.
This key will be used to sign events relayed to the EVM chain (message tuples or validator set updates).

```console
$ qgb tx register-eth-key \
  --cosmos-chain-id="..." \
  --cosmos-grpc="tcp://..." \
  --tendermint-rpc="http://..." \
  --cosmos-keyring=... \
  --cosmos-keyring-dir=... \
  --cosmos-from=... \
  --eth-pk=$ETH_PK
```

### Run the orchestrator

```console
$ qgb orchestrator \
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
-->

### Send a message from Celestia to an EVM chain

A message can be included on Celestia by using the Celestia app.
Instructions [here](https://github.com/celestiaorg/celestia-app).

## How it works

The QGB allows messages to be relayed in one direction, from Celestia to an EVM chain.
It does not support bridging assets such as fungible or non-fungible tokens directly, and cannot send messages from the EVM chain back to Celestia.

It works by relying on a set of signers to attest to some event on Celestia: the Celestia validator set.
The QGB contract keeps track of the Celestia validator set by updating its view of the validator set with `updateValidatorSet()`.
At least 2/3 of the voting power of the current view of the validator set must sign off on new relayed events, submitted with `submitMessageTupleRoot()`.
Each event is a batch of `MessageTuple`s, with each tuple representing a single message posted to Celestia.
Relayed tuples are in the same order as the messages they represent are paid for on Celestia.

### Events and messages relayed

 **Validator sets**:
 The relayer informs the QGB contract who are the current validators and their power.
 This results in an execution of the `updateValidatorSet` function.

 **Batches**:
 The relayer informs the QGB contract of new tuple roots.
 This results in an execution of the `submitMessageTupleRoot` function.
