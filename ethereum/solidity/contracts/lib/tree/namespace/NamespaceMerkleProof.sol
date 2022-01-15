// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "./NamespaceNode.sol";

/// @notice Namespace Merkle Tree Proof structure.
struct NamespaceMerkleProof {
    // List of side nodes to verify and calculate tree.
    NamespaceNode[] sideNodes;
}
