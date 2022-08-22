# Quantum Gravity Bridge

<!-- markdownlint-disable MD013 MD041 -->

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://godoc.org/github.com/celestiaorg/quantum-gravity-bridge)
[![Go Report Card](https://goreportcard.com/badge/github.com/celestiaorg/quantum-gravity-bridge?style=flat-square)](https://goreportcard.com/report/github.com/celestiaorg/quantum-gravity-bridge)
[![Version](https://img.shields.io/github/tag/celestiaorg/quantum-gravity-bridge.svg?style=flat-square)](https://github.com/celestiaorg/quantum-gravity-bridge/releases/latest)
[![License: Apache-2.0](https://img.shields.io/github/license/celestiaorg/quantum-gravity-bridge.svg?style=flat-square)](https://github.com/celestiaorg/quantum-gravity-bridge/blob/main/LICENSE)

The Quantum Gravity Bridge (QGB) is a Celestia -> EVM message relay.
It is based on Umee's Gravity Bridge implementation, [Peggo](https://github.com/umee-network/peggo).
**This project is under active development and should not be used in production**.

## Table of Contents

- [Building From Source](#building-from-source)
- [Send a message from Celestia to an EVM chain](#send-a-message-from-celestia-to-an-evm-chain)
- [How it works](#how-it-works)

## Building From Source

### Dependencies

Initialize git submodules, needed for Forge dependencies:

```sh
git submodule init
git submodule update
```

To regenerate the Go ABI wrappers with `make gen`, you need the `abigen` tool.
Building requires [Go 1.17+](https://golang.org/dl/).
Install `abigen` with:

```sh
git clone https://github.com/ethereum/go-ethereum.git
cd go-ethereum
make devtools
```

### Build and Test Contracts

Build with:

```sh
forge build
```

Test with:

```sh
forge test
```

### Regenerate Go Wrappers

Go wrappers can be regenerated with:

```sh
make
```

## Send a message from Celestia to an EVM chain

A message can be included on Celestia by using the Celestia app.
Instructions [here](https://github.com/celestiaorg/celestia-app).

## How it works

The QGB allows Celestia block header data roots to be relayed in one direction, from Celestia to an EVM chain.
It does not support bridging assets such as fungible or non-fungible tokens directly, and cannot send messages from the EVM chain back to Celestia.

It works by relying on a set of signers to attest to some event on Celestia: the Celestia validator set.
The QGB contract keeps track of the Celestia validator set by updating its view of the validator set with `updateValidatorSet()`.
More than 2/3 of the voting power of the current view of the validator set must sign off on new relayed events, submitted with `submitDataRootTupleRoot()`.
Each event is a batch of `DataRootTuple`s, with each tuple representing a single [data root (i.e. block header)](https://celestiaorg.github.io/celestia-specs/latest/specs/data_structures.html#header).
Relayed tuples are in the same order as Celestia block headers.

### Events and messages relayed

 **Validator sets**:
 The relayer informs the QGB contract who are the current validators and their power.
 This results in an execution of the `updateValidatorSet` function.

 **Batches**:
 The relayer informs the QGB contract of new data root tuple roots.
 This results in an execution of the `submitDataRootTupleRoot` function.
