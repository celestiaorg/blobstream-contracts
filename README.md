# Blobstream-contracts

<!-- markdownlint-disable MD013 MD041 -->

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://godoc.org/github.com/celestiaorg/blobstream-contracts)
[![Version](https://img.shields.io/github/tag/celestiaorg/blobstream-contracts.svg?style=flat-square)](https://github.com/celestiaorg/blobstream-contracts/releases/latest)
[![License: Apache-2.0](https://img.shields.io/github/license/celestiaorg/blobstream-contracts.svg?style=flat-square)](https://github.com/celestiaorg/blobstream-contracts/blob/master/LICENSE)

Blobstream is a Celestia -> EVM message relay.
It is based on Umee's Gravity Bridge implementation, [Peggo](https://github.com/umee-network/peggo).

## ⚠️ DEPRECATED

**The [Blobstream contract](https://github.com/celestiaorg/blobstream-contracts/blob/master/src/Blobstream.sol) is deprecated and not maintained. It may contain bugs and should not be used in production.**

**For production use, please use one of the following maintained alternatives:**

- **[sp1-blobstream](https://github.com/succinctlabs/sp1-blobstream/)** - The official implementation used in production deployments
- **[blobstream0](https://github.com/risc0/blobstream0)** - Alternative implementation using RISC Zero

Only the interface and verification libraries in this repository are still relevant for integration purposes.

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
Building requires [Go 1.19+](https://golang.org/dl/).
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

### Format

Format Solidity with:

```sh
forge fmt
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

Blobstream allows Celestia block header data roots to be relayed in one direction, from Celestia to an EVM chain.
It does not support bridging assets such as fungible or non-fungible tokens directly, and cannot send messages from the EVM chain back to Celestia.

It works by relying on a set of signers to attest to some event on Celestia: the Celestia validator set.
Blobstream contract keeps track of the Celestia validator set by updating its view of the validator set with `updateValidatorSet()`.
More than 2/3 of the voting power of the current view of the validator set must sign off on new relayed events, submitted with `submitDataRootTupleRoot()`.
Each event is a batch of `DataRootTuple`s, with each tuple representing a single [data root (i.e. block header)](https://celestiaorg.github.io/celestia-app/specs/data_structures.html#header).
Relayed tuples are in the same order as Celestia block headers.

### Events and messages relayed

 **Validator sets**:
 The relayer informs the Blobstream contract who are the current validators and their power.
 This results in an execution of the `updateValidatorSet` function.

 **Batches**:
 The relayer informs the Blobstream contract of new data root tuple roots.
 This results in an execution of the `submitDataRootTupleRoot` function.

## Audits

| Date       | Auditor                                       | celestia-app                                                                        | blobstream-contracts                                                                                           | Report                                                                                                                                         |
|------------|-----------------------------------------------|-------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| 2023/10/17 | [Binary Builders](https://binary.builders/)   | [v1.0.0-rc10](https://github.com/celestiaorg/celestia-app/releases/tag/v1.0.0-rc10) | [eb7a4e7](https://github.com/celestiaorg/blobstream-contracts/commit/eb7a4e74718b80277ad9dde116ead67383f5fe15) | [binary-builders.pdf](https://github.com/celestiaorg/blobstream-contracts/files/13961809/2023-10-17_Celestia_Audit_Report_Binary_Builders.pdf) |
| 2023/10/26 | [Informal Systems](https://informal.systems/) | [v1.0.0](https://github.com/celestiaorg/celestia-app/tree/v1.0.0)                   | [cf301adf](https://github.com/celestiaorg/blobstream-contracts/blob/cf301adfbfdae138526199fab805822400dcfd5d)  | [informal-systems.pdf](https://github.com/celestiaorg/blobstream-contracts/files/13961767/Celestia_.Q4.2023.QGB-v2-20231026_182304.pdf)        |
| 2023/11/16 | [Ottersec](https://osec.io/)                  | [v1.3.0](https://github.com/celestiaorg/celestia-app/releases/tag/v1.3.0)           | [v3.1.0](https://github.com/celestiaorg/blobstream-contracts/releases/tag/v3.1.0)                              | [ottersec.pdf](https://github.com/celestiaorg/blobstream-contracts/files/14383577/celestia_blobstream_audit_final.pdf)                         |
