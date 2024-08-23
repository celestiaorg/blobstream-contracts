// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "../Constants.sol";

/// @notice Calculate the digest of a node.
/// @param left The left child.
/// @param right The right child.
/// @return digest The node digest.
/// @dev More details in https://github.com/celestiaorg/celestia-specs/blob/master/src/specs/data_structures.md#binary-merkle-tree
// solhint-disable-next-line func-visibility
function nodeDigest(bytes32 left, bytes32 right) pure returns (bytes32 digest) {
    digest = sha256(abi.encodePacked(Constants.NODE_PREFIX, left, right));
}

/// @notice Calculate the digest of a leaf.
/// @param data The data of the leaf.
/// @return digest The leaf digest.
/// @dev More details in https://github.com/celestiaorg/celestia-specs/blob/master/src/specs/data_structures.md#binary-merkle-tree
// solhint-disable-next-line func-visibility
function leafDigest(bytes memory data) pure returns (bytes32 digest) {
    digest = sha256(abi.encodePacked(Constants.LEAF_PREFIX, data));
}
