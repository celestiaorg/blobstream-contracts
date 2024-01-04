// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "../Constants.sol";
import "../Utils.sol";
import "./TreeHasher.sol";
import "./BinaryMerkleProof.sol";

/// @title Binary Merkle Tree.
library BinaryMerkleTree {
    /// @notice Verify if element exists in Merkle tree, given data, proof, and root.
    /// @param root The root of the tree in which verify the given leaf.
    /// @param proof Binary Merkle proof for the leaf.
    /// @param data The data of the leaf to verify.
    /// @return `true` is proof is valid, `false` otherwise.
    /// @dev proof.numLeaves is necessary to determine height of subtree containing the data to prove.
    function verify(bytes32 root, BinaryMerkleProof memory proof, bytes memory data) internal pure returns (bool) {
        // Check proof is correct length for the key it is proving
        if (proof.numLeaves <= 1) {
            if (proof.sideNodes.length != 0) {
                return false;
            }
        } else if (proof.sideNodes.length != pathLengthFromKey(proof.key, proof.numLeaves)) {
            return false;
        }

        // Check key is in tree
        if (proof.key >= proof.numLeaves) {
            return false;
        }

        // A sibling at height 1 is created by getting the hash of the data to prove.
        bytes32 digest = leafDigest(data);

        // Null proof is only valid if numLeaves = 1
        // If so, just verify hash(data) is root
        if (proof.sideNodes.length == 0) {
            if (proof.numLeaves == 1) {
                return (root == digest);
            } else {
                return false;
            }
        }

        bytes32 computedHash = computeRootHash(proof.key, proof.numLeaves, digest, proof.sideNodes);

        return (computedHash == root);
    }

    /// @notice Use the leafHash and innerHashes to get the root merkle hash.
    /// If the length of the innerHashes slice isn't exactly correct, the result is nil.
    /// Recursive impl.
    function computeRootHash(uint256 key, uint256 numLeaves, bytes32 leafHash, bytes32[] memory sideNodes)
        private
        pure
        returns (bytes32)
    {
        if (numLeaves == 0) {
            revert("cannot call computeRootHash with 0 number of leaves");
        }
        if (numLeaves == 1) {
            if (sideNodes.length != 0) {
                revert("unexpected inner hashes");
            }
            return leafHash;
        }
        if (sideNodes.length == 0) {
            revert("expected at least one inner hash");
        }
        uint256 numLeft = _getSplitPoint(numLeaves);
        bytes32[] memory sideNodesLeft = slice(sideNodes, 0, sideNodes.length - 1);
        if (key < numLeft) {
            bytes32 leftHash = computeRootHash(key, numLeft, leafHash, sideNodesLeft);
            return nodeDigest(leftHash, sideNodes[sideNodes.length - 1]);
        }
        bytes32 rightHash = computeRootHash(key - numLeft, numLeaves - numLeft, leafHash, sideNodesLeft);
        return nodeDigest(sideNodes[sideNodes.length - 1], rightHash);
    }

    /// @notice creates a slice of bytes32 from the data slice of bytes32 containing the elements
    /// that correspond to the provided range.
    /// It selects a half-open range which includes the begin element, but excludes the end one.
    /// @param _data The slice that we want to select data from.
    /// @param _begin The beginning of the range (inclusive).
    /// @param _end The ending of the range (exclusive).
    /// @return _ the sliced data.
    function slice(bytes32[] memory _data, uint256 _begin, uint256 _end) internal pure returns (bytes32[] memory) {
        if (_begin > _end) {
            revert("Invalid range: _begin is greater than _end");
        }
        if (_begin > _data.length || _end > _data.length) {
            revert("Invalid range: _begin or _end are out of bounds");
        }
        bytes32[] memory out = new bytes32[](_end - _begin);
        for (uint256 i = _begin; i < _end; i++) {
            out[i - _begin] = _data[i];
        }
        return out;
    }
}
