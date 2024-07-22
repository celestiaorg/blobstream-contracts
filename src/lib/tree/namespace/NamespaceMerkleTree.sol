// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "../Constants.sol";
import "../Types.sol";
import "../Utils.sol";
import "./NamespaceMerkleProof.sol";
import "./NamespaceMerkleMultiproof.sol";
import "./NamespaceNode.sol";
import "./TreeHasher.sol";

/// @title Namespace Merkle Tree.
library NamespaceMerkleTree {
    /// @notice Verify if element exists in Merkle tree, given data, proof, and root.
    /// @param root The root of the tree in which the given leaf is verified.
    /// @param proof Namespace Merkle proof for the leaf.
    /// @param namespace Namespace of the leaf.
    /// @param data The data of the leaf to verify.
    /// @return `true` if the proof is valid, `false` otherwise.
    /// @dev proof.numLeaves is necessary to determine height of subtree containing the data to prove.
    function verify(
        NamespaceNode memory root,
        NamespaceMerkleProof memory proof,
        Namespace memory namespace,
        bytes memory data
    ) internal pure returns (bool) {
        // A sibling at height 1 is created by getting the leafDigest of the original data.
        NamespaceNode memory node = leafDigest(namespace, data);

        // Since we're verifying a leaf, height parameter is 1.
        return verifyInner(root, proof, node, 1);
    }

    /// @notice Verify if inner node exists in Merkle tree, given node, proof, and root.
    /// @param root The root of the tree in which the given leaf is verified.
    /// @param proof Namespace Merkle proof for the leaf.
    /// proof.key is any key in the subtree rooted at the inner node.
    /// @param node The inner node to verify.
    /// @param startingHeight Starting height of the proof.
    /// @return `true` if the proof is valid, `false` otherwise.
    /// @dev proof.numLeaves is necessary to determine height of subtree containing the data to prove.
    function verifyInner(
        NamespaceNode memory root,
        NamespaceMerkleProof memory proof,
        NamespaceNode memory node,
        uint256 startingHeight
    ) internal pure returns (bool) {
        // Check starting height is at least 1
        if (startingHeight < 1) {
            return false;
        }
        uint256 heightOffset = startingHeight - 1;

        // Check proof is correct length for the key it is proving
        if (proof.numLeaves <= 1) {
            if (proof.sideNodes.length != 0) {
                return false;
            }
        } else if (proof.sideNodes.length + heightOffset != pathLengthFromKey(proof.key, proof.numLeaves)) {
            return false;
        }

        // Check key is in tree
        if (proof.key >= proof.numLeaves) {
            return false;
        }
        // Handle case where proof is empty: i.e, only one leaf exists, so verify hash(data) is root
        if (proof.sideNodes.length == 0) {
            if (proof.numLeaves == 1) {
                return namespaceNodeEquals(root, node);
            } else {
                return false;
            }
        }

        // The case where inner node is actually the root of a tree with more than one node is not relevant
        // to our use case, since the only case where an inner node is the root of the tree is when the tree
        // has only one inner node. So, there is no need to handle that case.

        uint256 height = startingHeight;
        uint256 stableEnd = proof.key;

        // While the current subtree (of height 'height') is complete, determine
        // the position of the next sibling using the complete subtree algorithm.
        // 'stableEnd' tells us the ending index of the last full subtree. It gets
        // initialized to 'key' because the first full subtree was the
        // subtree of height 1, created above (and had an ending index of
        // 'key').

        while (true) {
            // Determine if the subtree is complete. This is accomplished by
            // rounding down the key to the nearest 1 << 'height', adding 1
            // << 'height', and comparing the result to the number of leaves in the
            // Merkle tree.

            uint256 subTreeStartIndex = (proof.key / (1 << height)) * (1 << height);
            uint256 subTreeEndIndex = subTreeStartIndex + (1 << height) - 1;

            // If the Merkle tree does not have a leaf at index
            // 'subTreeEndIndex', then the subtree of the current height is not
            // a complete subtree.
            if (subTreeEndIndex >= proof.numLeaves) {
                break;
            }
            stableEnd = subTreeEndIndex;

            // Determine if the key is in the first or the second half of
            // the subtree.
            if (proof.sideNodes.length + heightOffset <= height - 1) {
                return false;
            }
            if (proof.key - subTreeStartIndex < (1 << (height - 1))) {
                node = nodeDigest(node, proof.sideNodes[height - heightOffset - 1]);
            } else {
                node = nodeDigest(proof.sideNodes[height - heightOffset - 1], node);
            }

            height += 1;
        }

        // Determine if the next hash belongs to an orphan that was elevated. This
        // is the case IFF 'stableEnd' (the last index of the largest full subtree)
        // is equal to the number of leaves in the Merkle tree.
        if (stableEnd != proof.numLeaves - 1) {
            if (proof.sideNodes.length <= height - heightOffset - 1) {
                return false;
            }
            node = nodeDigest(node, proof.sideNodes[height - heightOffset - 1]);
            height += 1;
        }
        // All remaining elements in the proof set will belong to a left sibling.
        while (height - heightOffset - 1 < proof.sideNodes.length) {
            node = nodeDigest(proof.sideNodes[height - heightOffset - 1], node);
            height += 1;
        }

        return namespaceNodeEquals(root, node);
    }

    /// @notice Verify if contiguous elements exists in Merkle tree, given leaves, mutliproof, and root.
    /// @param root The root of the tree in which the given leaves are verified.
    /// @param proof Namespace Merkle multiproof for the leaves.
    /// @param namespace Namespace of the leaves. All leaves must have the same namespace.
    /// @param data The leaves to verify. Note: leaf data must be the _entire_ share (including namespace prefixing).
    /// @return `true` if the proof is valid, `false` otherwise.
    function verifyMulti(
        NamespaceNode memory root,
        NamespaceMerkleMultiproof memory proof,
        Namespace memory namespace,
        bytes[] memory data
    ) internal pure returns (bool) {
        // Hash all the leaves to get leaf nodes.
        NamespaceNode[] memory nodes = new NamespaceNode[](data.length);
        for (uint256 i = 0; i < data.length; ++i) {
            nodes[i] = leafDigest(namespace, data[i]);
        }

        // Verify inclusion of leaf nodes.
        return verifyMultiHashes(root, proof, nodes);
    }

    /// @notice Verify if contiguous leaf hashes exists in Merkle tree, given leaf nodes, multiproof, and root.
    /// @param root The root of the tree in which the given leaf nodes are verified.
    /// @param proof Namespace Merkle multiproof for the leaves.
    /// @param leafNodes The leaf nodes to verify.
    /// @return `true` if the proof is valid, `false` otherwise.
    function verifyMultiHashes(
        NamespaceNode memory root,
        NamespaceMerkleMultiproof memory proof,
        NamespaceNode[] memory leafNodes
    ) internal pure returns (bool) {
        uint256 leafIndex = 0;
        NamespaceNode[] memory leftSubtrees = new NamespaceNode[](proof.sideNodes.length);

        for (uint256 i = 0; leafIndex != proof.beginKey && i < proof.sideNodes.length; ++i) {
            uint256 subtreeSize = _nextSubtreeSize(leafIndex, proof.beginKey);
            leftSubtrees[i] = proof.sideNodes[i];
            leafIndex += subtreeSize;
        }

        // estimate the leaf size of the subtree containing the proof range
        uint256 proofRangeSubtreeEstimate = _getSplitPoint(proof.endKey) * 2;
        if (proofRangeSubtreeEstimate < 1) {
            proofRangeSubtreeEstimate = 1;
        }

        (NamespaceNode memory rootHash, uint256 proofHead,,) =
            _computeRoot(proof, leafNodes, 0, proofRangeSubtreeEstimate, 0, 0);
        for (uint256 i = proofHead; i < proof.sideNodes.length; ++i) {
            rootHash = nodeDigest(rootHash, proof.sideNodes[i]);
        }

        return namespaceNodeEquals(rootHash, root);
    }

    /// @notice Computes the NMT root recursively.
    /// @param proof Namespace Merkle multiproof for the leaves.
    /// @param leafNodes Leaf nodes for which inclusion is proven.
    /// @param begin Begin index, inclusive.
    /// @param end End index, exclusive.
    /// @param headProof Internal detail: head of proof sidenodes array. Used for recursion. Set to `0` on first call.
    /// @param headLeaves Internal detail: head of leaves array. Used for recursion. Set to `0` on first call.
    /// @return _ Subtree root.
    /// @return _ New proof sidenodes array head. Used for recursion.
    /// @return _ New leaves array head. Used for recursion.
    /// @return _ If the subtree root is "nil."
    function _computeRoot(
        NamespaceMerkleMultiproof memory proof,
        NamespaceNode[] memory leafNodes,
        uint256 begin,
        uint256 end,
        uint256 headProof,
        uint256 headLeaves
    ) private pure returns (NamespaceNode memory, uint256, uint256, bool) {
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
        (NamespaceNode memory left, uint256 newHeadProofLeft, uint256 newHeadLeavesLeft,) =
            _computeRoot(proof, leafNodes, begin, begin + k, headProof, headLeaves);
        (NamespaceNode memory right, uint256 newHeadProof, uint256 newHeadLeaves, bool rightIsNil) =
            _computeRoot(proof, leafNodes, begin + k, end, newHeadProofLeft, newHeadLeavesLeft);

        // only right leaf/subtree can be non-existent
        if (rightIsNil == true) {
            return (left, newHeadProof, newHeadLeaves, false);
        }
        NamespaceNode memory hash = nodeDigest(left, right);
        return (hash, newHeadProof, newHeadLeaves, false);
    }

    /// @notice Pop from the leaf nodes array slice if it's not empty.
    /// @param nodes Entire leaf nodes array.
    /// @param headLeaves Head of leaf nodes array slice.
    /// @param end End of leaf nodes array slice.
    /// @param headProof Used only to return for recursion.
    /// @return _ Popped node.
    /// @return _ Head of proof sidenodes array slice (unchanged).
    /// @return _ New head of leaf nodes array slice.
    /// @return _ If the popped node is "nil."
    function _popLeavesIfNonEmpty(NamespaceNode[] memory nodes, uint256 headLeaves, uint256 end, uint256 headProof)
        private
        pure
        returns (NamespaceNode memory, uint256, uint256, bool)
    {
        (NamespaceNode memory node, uint256 newHead, bool isNil) = _popIfNonEmpty(nodes, headLeaves, end);
        return (node, headProof, newHead, isNil);
    }

    /// @notice Pop from the proof sidenodes array slice if it's not empty.
    /// @param nodes Entire proof sidenodes array.
    /// @param headLeaves Head of proof sidenodes array slice.
    /// @param end End of proof sidenodes array slice.
    /// @param headProof Used only to return for recursion.
    /// @return _ Popped node.
    /// @return _ New head of proof sidenodes array slice.
    /// @return _ Head of proof sidenodes array slice (unchanged).
    /// @return _ If the popped node is "nil."
    function _popProofIfNonEmpty(NamespaceNode[] memory nodes, uint256 headProof, uint256 end, uint256 headLeaves)
        private
        pure
        returns (NamespaceNode memory, uint256, uint256, bool)
    {
        (NamespaceNode memory node, uint256 newHead, bool isNil) = _popIfNonEmpty(nodes, headProof, end);
        return (node, newHead, headLeaves, isNil);
    }

    /// @notice Pop from an array slice if it's not empty.
    /// @param nodes Entire array.
    /// @param head Head of array slice.
    /// @param end End of array slice.
    /// @return _ Popped node.
    /// @return _ New head of array slice.
    /// @return _ If the popped node is "nil."
    function _popIfNonEmpty(NamespaceNode[] memory nodes, uint256 head, uint256 end)
        private
        pure
        returns (NamespaceNode memory, uint256, bool)
    {
        if (nodes.length == 0 || head >= nodes.length || head >= end) {
            NamespaceNode memory node;
            return (node, head, true);
        }
        return (nodes[head], head + 1, false);
    }
}
