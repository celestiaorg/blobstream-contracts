# QGB fraud proofs

A Celestium is a rollup that uses Celestia for data availability but settles on Ethereum or any other EVM chain. In simpler terms, it's a layer 2 solution that leverages the Celestia blockchain security to ensure that data is available, while the actual transactions and smart contract interactions are settled on the Ethereum blockchain or any other EVM-compatible blockchain.

Celestiums use the Quantum Gravity Bridge (QGB) contract, deployed on an EVM chain, to verify that their transaction data has been published to Celestia correctly. Alternatively, they can use an optimistic scheme, only verifying transaction data inclusion during fraud proof execution.

# Fraud proofs

Fraud proofs is the mechanism used to inform light clients in the case of an invalid rollup state transition or unavailable rollup block data. They rely on rollup full nodes getting the data that was published to Celestia, and executing all the state transitions to verify the rollup state. If they discover an invalid state transition or unavailable rollup data, they emit a fraud proof with the necessary information to convince light clients that fraud happened. This allows for trust-minimized light clients, as the network only needs one honest full node to create the fraud proof and propagate it.

If the Celestium is settlement smart contract based, then the contract would only need to receive a fraud proof to decide whether data was published correctly or not.

## Rollup header

Rollups can adopt many approaches to prove that fraud happened. One of which could be having the following fields in the rollup header:

- Rollup block state root
- A sequence of spans in Celestia: which references where the rollup data was published in the Celestia chain.

> [!NOTE]  
> The sequence of spans can be defined using the following: `Height`, `start index`, and `length` in the Celestia block, in the case of a single Celestia block. However, it could be generalized to span over multiple blocks.

For the rest of the document, we will suppose that the sequence of spans only references one Celestia block.

## Unavailable data fraud proof

By construction, the rollup block data is the sequence of spans defined in the header. Thus, proving that the data is unavailable goes back to proving that the sequence of spans doesn't belong to the Celestia block, i.e. the span is out of bounds.

We could prove that via creating a binary [Merkle proof](https://github.com/celestiaorg/celestia-core/blob/c3ab251659f6fe0f36d10e0dbd14c29a78a85352/crypto/merkle/proof.go#L19-L31) of any row/column to the Celestia data root. This proof will provide the `Total` which is the number of rows/columns in the extended data square. And, we can use that to calculate the square size.

Then, we will use that information to check if the provided transaction index, in the header, is out of the square size bounds.

For the data root, we will use a binary Merkle proof to prove its inclusion in a data root tuple root that was committed to by the QGB smart contract. More on this in [here](#1-data-root-inclusion-proof).

## Invalid transaction fraud proof

In order to prove an invalid transaction in the rollup, we need to prove the following:

- Prove that that transaction was posted to Celestia
- Prove that the transaction is invalid: left to the rollup to define.

The first part, proving that the transaction was posted to Celestia, can be done in three steps:

1. Prove that the data root is committed to by the QGB smart contract
2. Inclusion proof of the transaction to Celestia data root
3. Prove that that transaction is in the rollup sequence spans

### 1. Data root inclusion proof

To prove the data root is committed to by the QGB smart contract, we will need to provide a Merkle proof of the data root tuple to a data root tuple root. This can be created using the [`data_root_inclusion_proof`](https://github.com/celestiaorg/celestia-core/blob/c3ab251659f6fe0f36d10e0dbd14c29a78a85352/rpc/client/http/http.go#L492-L511) query.

### 2. Transaction inclusion proof

To prove that a rollup transaction is part of the data root, we will need to provide two proofs: a namespace merkle proof of the transaction to a row root. This could be done via proving the shares that contain the transaction to the row root using a namespace merkle proof. And, a binary merkle proof of the row root to the data root.

These proofs can be generated using the [`ProveShares`](https://github.com/celestiaorg/celestia-core/blob/c3ab251659f6fe0f36d10e0dbd14c29a78a85352/rpc/client/http/http.go#L526-L543) query.

### 3. Transaction part of the rollup sequence

The Celestia block is a 2d matrix of rows and columns, where the row roots and column roots are committed to in the data root. We can use this property to point which rows and columns does the transaction shares belong to. Then, we will use those coordinates to calculate the row major index of the transaction, and verify if it is part of the sequence span.

The coordinates can be gotten using a namespace merkle proof of the shares to the row root, and also to the column root. Then, we will create merkle proofs of that row root and column root to the data root. These [proofs](https://github.com/celestiaorg/celestia-core/blob/c3ab251659f6fe0f36d10e0dbd14c29a78a85352/crypto/merkle/proof.go#L19-L31) will contain the [`Index`](https://github.com/celestiaorg/celestia-core/blob/c3ab251659f6fe0f36d10e0dbd14c29a78a85352/crypto/merkle/proof.go#L28) of the row/column, and the [`Total`](https://github.com/celestiaorg/celestia-core/blob/c3ab251659f6fe0f36d10e0dbd14c29a78a85352/crypto/merkle/proof.go#L27) number of rows/columns in the proofs which would allow us to calculate the row major index of the transaction.

Note: the `Total` is the total number of rows/columns of the extended data square, and not the original one.
