// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "../Constants.sol";
import "../Types.sol";
import "./NamespaceNode.sol";

/// @notice Get the minimum namespace.
// solhint-disable-next-line func-visibility
function namespaceMin(Namespace memory l, Namespace memory r) pure returns (Namespace memory) {
    if (l.lessThan(r)) {
        return l;
    } else {
        return r;
    }
}

/// @notice Get the maximum namespace.
// solhint-disable-next-line func-visibility
function namespaceMax(Namespace memory l, Namespace memory r) pure returns (Namespace memory) {
    if (l.greaterThan(r)) {
        return l;
    } else {
        return r;
    }
}

/// @notice Hash a leaf node.
/// @param namespace Namespace of the leaf.
/// @param data Raw data of the leaf.
/// @dev More details in https://github.com/celestiaorg/celestia-specs/blob/master/src/specs/data_structures.md#namespace-merkle-tree
// solhint-disable-next-line func-visibility
function leafDigest(Namespace memory namespace, bytes memory data) pure returns (NamespaceNode memory) {
    bytes32 digest = sha256(abi.encodePacked(Constants.LEAF_PREFIX, namespace.toBytes(), data));
    NamespaceNode memory node = NamespaceNode(namespace, namespace, digest);
    return node;
}

/// @notice Hash an internal node.
/// @param l Left child.
/// @param r Right child.
/// @dev More details in https://github.com/celestiaorg/celestia-specs/blob/master/src/specs/data_structures.md#namespace-merkle-tree
// solhint-disable-next-line func-visibility
function nodeDigest(NamespaceNode memory l, NamespaceNode memory r) pure returns (NamespaceNode memory) {
    Namespace memory min = namespaceMin(l.min, r.min);
    Namespace memory max;
    if (l.min.equalTo(PARITY_SHARE_NAMESPACE())) {
        max = PARITY_SHARE_NAMESPACE();
    } else if (r.min.equalTo(PARITY_SHARE_NAMESPACE())) {
        max = l.max;
    } else {
        max = namespaceMax(l.max, r.max);
    }

    bytes32 digest = sha256(
        abi.encodePacked(
            Constants.NODE_PREFIX,
            l.min.toBytes(),
            l.max.toBytes(),
            l.digest,
            r.min.toBytes(),
            r.max.toBytes(),
            r.digest
        )
    );

    NamespaceNode memory node = NamespaceNode(min, max, digest);
    return node;
}
