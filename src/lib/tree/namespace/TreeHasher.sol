// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.19;

import "../Constants.sol";
import "../Types.sol";
import "./NamespaceNode.sol";

/// @notice Get the minimum namespace.
// solhint-disable-next-line func-visibility
function namespaceMin(NamespaceID l, NamespaceID r) pure returns (NamespaceID) {
    if (l < r) {
        return l;
    } else {
        return r;
    }
}

/// @notice Get the maximum namespace.
// solhint-disable-next-line func-visibility
function namespaceMax(NamespaceID l, NamespaceID r) pure returns (NamespaceID) {
    if (l > r) {
        return l;
    } else {
        return r;
    }
}

/// @notice Hash a leaf node.
/// @param minmaxNID Namespace ID.
/// @param data Raw data of the leaf.
/// @dev More details in https://github.com/celestiaorg/celestia-specs/blob/master/src/specs/data_structures.md#namespace-merkle-tree
// solhint-disable-next-line func-visibility
function leafDigest(NamespaceID minmaxNID, bytes memory data) pure returns (NamespaceNode memory) {
    bytes32 digest = sha256(abi.encodePacked(Constants.LEAF_PREFIX, minmaxNID, data));
    NamespaceNode memory node = NamespaceNode(minmaxNID, minmaxNID, digest);
    return node;
}

/// @notice Hash an internal node.
/// @param l Left child.
/// @param r Right child.
/// @dev More details in https://github.com/celestiaorg/celestia-specs/blob/master/src/specs/data_structures.md#namespace-merkle-tree
// solhint-disable-next-line func-visibility
function nodeDigest(NamespaceNode memory l, NamespaceNode memory r) pure returns (NamespaceNode memory) {
    NamespaceID min = namespaceMin(l.min, r.min);
    NamespaceID max;
    if (l.min == Constants.PARITY_SHARE_NAMESPACE_ID) {
        max = Constants.PARITY_SHARE_NAMESPACE_ID;
    } else if (r.min == Constants.PARITY_SHARE_NAMESPACE_ID) {
        max = l.max;
    } else {
        max = namespaceMax(l.max, r.max);
    }

    bytes32 digest = sha256(abi.encodePacked(Constants.NODE_PREFIX, l.min, l.max, l.digest, r.min, r.max, r.digest));

    NamespaceNode memory node = NamespaceNode(min, max, digest);
    return node;
}
