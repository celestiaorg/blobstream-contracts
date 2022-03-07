// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

/// @notice Namespace Merkle Tree node.
struct NamespaceNode {
    // Minimum namespace ID.
    bytes8 min;
    // Maximum namespace ID.
    bytes8 max;
    // Node value.
    bytes32 digest;
}

/// @notice Compares two `NamespaceNode`s.
/// @param first First node.
/// @param second Second node.
/// @return `true` is equal, `false otherwise.
// solhint-disable-next-line func-visibility
function namespaceNodeEquals(NamespaceNode memory first, NamespaceNode memory second) pure returns (bool) {
    return (first.min == second.min) && (first.max == second.max) && (first.digest == second.digest);
}
