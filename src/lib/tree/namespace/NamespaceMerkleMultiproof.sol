// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "./NamespaceNode.sol";

/// @notice Namespace Merkle Tree Multiproof structure. Proves multiple leaves.
struct NamespaceMerkleMultiproof {
    // The (included) beginning key of the leaves to verify.
    uint256 beginKey;
    // The (excluded) ending key of the leaves to verify.
    uint256 endKey;
    // List of side nodes to verify and calculate tree.
    NamespaceNode[] sideNodes;
}
