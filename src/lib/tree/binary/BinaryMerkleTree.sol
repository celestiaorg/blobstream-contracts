// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "../Constants.sol";
import "../Utils.sol";
import "./TreeHasher.sol";
import "./BinaryMerkleProof.sol";
import "./BinaryMerkleMultiproof.sol";

/// @title Binary Merkle Tree.
library BinaryMerkleTree {
    /////////////////
    // Error codes //
    /////////////////

    enum ErrorCodes {
        NoError,
        /// @notice The provided side nodes count is invalid for the proof.
        InvalidNumberOfSideNodes,
        /// @notice The provided proof key is not part of the tree.
        KeyNotInTree,
        /// @notice Invalid number of leaves in proof.
        InvalidNumberOfLeavesInProof,
        /// @notice The proof contains unexpected side nodes.
        UnexpectedInnerHashes,
        /// @notice The proof verification expected at least one inner hash.
        ExpectedAtLeastOneInnerHash
    }

    ///////////////
    // Functions //
    ///////////////

    /// @notice Verify if element exists in Merkle tree, given data, proof, and root.
    /// @param root The root of the tree in which verify the given leaf.
    /// @param proof Binary Merkle proof for the leaf.
    /// @param data The data of the leaf to verify.
    /// @return `true` is proof is valid, `false` otherwise.
    /// @dev proof.numLeaves is necessary to determine height of subtree containing the data to prove.
    function verify(bytes32 root, BinaryMerkleProof memory proof, bytes memory data)
        internal
        pure
        returns (bool, ErrorCodes)
    {
        // Check proof is correct length for the key it is proving
        if (proof.numLeaves <= 1) {
            if (proof.sideNodes.length != 0) {
                return (false, ErrorCodes.InvalidNumberOfSideNodes);
            }
        } else if (proof.sideNodes.length != pathLengthFromKey(proof.key, proof.numLeaves)) {
            return (false, ErrorCodes.InvalidNumberOfSideNodes);
        }

        // Check key is in tree
        if (proof.key >= proof.numLeaves) {
            return (false, ErrorCodes.KeyNotInTree);
        }

        // A sibling at height 1 is created by getting the hash of the data to prove.
        bytes32 digest = leafDigest(data);

        // Null proof is only valid if numLeaves = 1
        // If so, just verify hash(data) is root
        if (proof.sideNodes.length == 0) {
            if (proof.numLeaves == 1) {
                return (root == digest, ErrorCodes.NoError);
            } else {
                return (false, ErrorCodes.NoError);
            }
        }

        (bytes32 computedHash, ErrorCodes error) = computeRootHash(proof.key, proof.numLeaves, digest, proof.sideNodes);

        if (error != ErrorCodes.NoError) {
            return (false, error);
        }

        return (computedHash == root, ErrorCodes.NoError);
    }

    function verifyMulti(bytes32 root, BinaryMerkleMultiproof memory proof, bytes[] memory data)
        internal
        pure
        returns (bool)
    {
        bytes32[] memory nodes = new bytes32[](data.length);
        for (uint256 i = 0; i < data.length; i++) {
            nodes[i] = leafDigest(data[i]);
        }

        return verifyMultiHashes(root, proof, nodes);
    }

    function verifyMultiHashes(bytes32 root, BinaryMerkleMultiproof memory proof, bytes32[] memory leafNodes)
        internal
        pure
        returns (bool)
    {
        uint256 leafIndex = 0;
        bytes32[] memory leftSubtrees = new bytes32[](proof.sideNodes.length);

        for (uint256 i = 0; leafIndex != proof.beginKey && i < proof.sideNodes.length; ++i) {
            uint256 subtreeSize = _nextSubtreeSize(leafIndex, proof.beginKey);
            leftSubtrees[i] = proof.sideNodes[i];
            leafIndex += subtreeSize;
        }

        uint256 proofRangeSubtreeEstimate = _getSplitPoint(proof.endKey) * 2;
        if (proofRangeSubtreeEstimate < 1) {
            proofRangeSubtreeEstimate = 1;
        }

        (bytes32 rootHash, uint256 proofHead,,) =
            _computeRootMulti(proof, leafNodes, 0, proofRangeSubtreeEstimate, 0, 0);
        for (uint256 i = proofHead; i < proof.sideNodes.length; ++i) {
            rootHash = nodeDigest(rootHash, proof.sideNodes[i]);
        }

        return (rootHash == root);
    }

    function _computeRootMulti(
        BinaryMerkleMultiproof memory proof,
        bytes32[] memory leafNodes,
        uint256 begin,
        uint256 end,
        uint256 headProof,
        uint256 headLeaves
    ) private pure returns (bytes32, uint256, uint256, bool) {
        // reached a leaf
        if (end - begin == 1) {
            // if current range overlaps with proof range, pop and return a leaf
            if (proof.beginKey <= begin && begin < proof.endKey) {
                // Note: second return value is guaranteed to be `false` by
                // construction.
                return _popLeavesIfNonEmpty(leafNodes, headLeaves, leafNodes.length, headProof);
            }

            // if current range does not overlap with proof range,
            // pop and return a proof node (leaf) if present,
            // else return nil because leaf doesn't exist
            return _popProofIfNonEmpty(proof.sideNodes, headProof, end, headLeaves);
        }

        // if current range does not overlap with proof range,
        // pop and return a proof node if present,
        // else return nil because subtree doesn't exist
        if (end <= proof.beginKey || begin >= proof.endKey) {
            return _popProofIfNonEmpty(proof.sideNodes, headProof, end, headLeaves);
        }

        // Recursively get left and right subtree
        uint256 k = _getSplitPoint(end - begin);
        (bytes32 left, uint256 newHeadProofLeft, uint256 newHeadLeavesLeft,) =
            _computeRootMulti(proof, leafNodes, begin, begin + k, headProof, headLeaves);
        (bytes32 right, uint256 newHeadProof, uint256 newHeadLeaves, bool rightIsNil) =
            _computeRootMulti(proof, leafNodes, begin + k, end, newHeadProofLeft, newHeadLeavesLeft);

        // only right leaf/subtree can be non-existent
        if (rightIsNil == true) {
            return (left, newHeadProof, newHeadLeaves, false);
        }
        bytes32 hash = nodeDigest(left, right);
        return (hash, newHeadProof, newHeadLeaves, false);
    }

    function _popProofIfNonEmpty(bytes32[] memory nodes, uint256 headProof, uint256 end, uint256 headLeaves)
        private
        pure
        returns (bytes32, uint256, uint256, bool)
    {
        (bytes32 node, uint256 newHead, bool isNil) = _popIfNonEmpty(nodes, headProof, end);
        return (node, newHead, headLeaves, isNil);
    }

    function _popLeavesIfNonEmpty(bytes32[] memory nodes, uint256 headLeaves, uint256 end, uint256 headProof)
        private
        pure
        returns (bytes32, uint256, uint256, bool)
    {
        (bytes32 node, uint256 newHead, bool isNil) = _popIfNonEmpty(nodes, headLeaves, end);
        return (node, headProof, newHead, isNil);
    }

    function _popIfNonEmpty(bytes32[] memory nodes, uint256 head, uint256 end)
        private
        pure
        returns (bytes32, uint256, bool)
    {
        if (nodes.length == 0 || head >= nodes.length || head >= end) {
            bytes32 node;
            return (node, head, true);
        }
        return (nodes[head], head + 1, false);
    }

    /// @notice Use the leafHash and innerHashes to get the root merkle hash.
    /// If the length of the innerHashes slice isn't exactly correct, the result is nil.
    /// Recursive impl.
    function computeRootHash(uint256 key, uint256 numLeaves, bytes32 leafHash, bytes32[] memory sideNodes)
        private
        pure
        returns (bytes32, ErrorCodes)
    {
        if (numLeaves == 0) {
            return (leafHash, ErrorCodes.InvalidNumberOfLeavesInProof);
        }
        if (numLeaves == 1) {
            if (sideNodes.length != 0) {
                return (leafHash, ErrorCodes.UnexpectedInnerHashes);
            }
            return (leafHash, ErrorCodes.NoError);
        }
        if (sideNodes.length == 0) {
            return (leafHash, ErrorCodes.ExpectedAtLeastOneInnerHash);
        }
        uint256 numLeft = _getSplitPoint(numLeaves);
        bytes32[] memory sideNodesLeft = slice(sideNodes, 0, sideNodes.length - 1);
        ErrorCodes error;
        if (key < numLeft) {
            bytes32 leftHash;
            (leftHash, error) = computeRootHash(key, numLeft, leafHash, sideNodesLeft);
            if (error != ErrorCodes.NoError) {
                return (leafHash, error);
            }
            return (nodeDigest(leftHash, sideNodes[sideNodes.length - 1]), ErrorCodes.NoError);
        }
        bytes32 rightHash;
        (rightHash, error) = computeRootHash(key - numLeft, numLeaves - numLeft, leafHash, sideNodesLeft);
        if (error != ErrorCodes.NoError) {
            return (leafHash, error);
        }
        return (nodeDigest(sideNodes[sideNodes.length - 1], rightHash), ErrorCodes.NoError);
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
