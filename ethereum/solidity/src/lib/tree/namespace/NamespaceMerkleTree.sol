// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import "../Constants.sol";
import "../Utils.sol";
import "./NamespaceMerkleProof.sol";
import "./NamespaceNode.sol";
import "./TreeHasher.sol";

/// @title Namespace Merkle Tree.
library NamespaceMerkleTree {
    /// @notice Verify if element exists in Merkle tree, given data, proof, and root.
    /// @param root The root of the tree in which verify the given leaf.
    /// @param proof Namespace Merkle proof for the leaf.
    /// @param minmaxNID Namespace ID of the leaf.
    /// @param data The data of the leaf to verify.
    /// @return `true` if the proof is valid, `false` otherwise.
    /// @dev proof.numLeaves is necessary to determine height of subtree containing the data to prove.
    function verify(
        NamespaceNode memory root,
        NamespaceMerkleProof memory proof,
        bytes8 minmaxNID,
        bytes memory data
    ) internal pure returns (bool) {
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

        // A sibling at height 1 is created by getting the leafDigest of the original data.
        NamespaceNode memory node = leafDigest(minmaxNID, data);

        // Handle case where proof is empty: i.e, only one leaf exists, so verify hash(data) is root
        if (proof.sideNodes.length == 0) {
            if (proof.numLeaves == 1) {
                return namespaceNodeEquals(root, node);
            } else {
                return false;
            }
        }

        uint256 height = 1;
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
            if (proof.sideNodes.length <= height - 1) {
                return false;
            }
            if (proof.key - subTreeStartIndex < (1 << (height - 1))) {
                node = nodeDigest(node, proof.sideNodes[height - 1]);
            } else {
                node = nodeDigest(proof.sideNodes[height - 1], node);
            }

            height += 1;
        }

        // Determine if the next hash belongs to an orphan that was elevated. This
        // is the case IFF 'stableEnd' (the last index of the largest full subtree)
        // is equal to the number of leaves in the Merkle tree.
        if (stableEnd != proof.numLeaves - 1) {
            if (proof.sideNodes.length <= height - 1) {
                return false;
            }
            node = nodeDigest(node, proof.sideNodes[height - 1]);
            height += 1;
        }

        // All remaining elements in the proof set will belong to a left sibling.
        while (height - 1 < proof.sideNodes.length) {
            node = nodeDigest(proof.sideNodes[height - 1], node);
            height += 1;
        }

        return namespaceNodeEquals(root, node);
    }
}
