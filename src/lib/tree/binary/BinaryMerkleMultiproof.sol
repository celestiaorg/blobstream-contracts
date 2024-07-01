// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

/// @notice Merkle Tree Proof structure.
struct BinaryMerkleMultiproof {
    // List of side nodes to verify and calculate tree.
    bytes32[] sideNodes;
    // The beginning key of the leaves to verify.
    uint256 beginKey;
    // The ending key of the leaves to verify.
    uint256 endKey;
}
