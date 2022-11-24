// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import "./NamespaceNode.sol";

/// @notice Namespace Merkle Tree Multiproof structure. Proves multiple leaves.
struct NamespaceMerkleMultiproof {
    // The beginning key of the leaves to verify.
    uint256 beginKey;
    // The ending key of the leaves to verify.
    uint256 endKey;
    // List of side nodes to verify and calculate tree.
    NamespaceNode[] sideNodes;
    // Empty if the namespace is present in the tree. In case the namespace to
    // be proved is in the min/max range of the tree but absent, this will
    // contain the leaf node necessary to verify the proof of absence.
    NamespaceNode[] leafNode;
    // The number of leaves in the tree.
    uint256 numLeaves;
}
