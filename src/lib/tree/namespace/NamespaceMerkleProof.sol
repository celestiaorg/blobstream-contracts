// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "./NamespaceNode.sol";

/// @notice Namespace Merkle Tree Proof structure.
struct NamespaceMerkleProof {
    // List of side nodes to verify and calculate tree.
    NamespaceNode[] sideNodes;
    // The key of the leaf to verify.
    uint256 key;
    // The number of leaves in the tree
    uint256 numLeaves;
}
