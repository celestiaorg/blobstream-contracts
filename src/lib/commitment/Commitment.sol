// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Namespace, isReservedNamespace} from "../tree/Types.sol";
import "../tree/namespace/NamespaceMerkleTree.sol";
import "../tree/binary/BinaryMerkleTree.sol";
import "../tree/binary/BinaryMerkleMultiproof.sol";
import "../tree/namespace/NamespaceNode.sol";
import "../tree/namespace/NamespaceMerkleMultiproof.sol";
import {leafDigest} from "../tree/namespace/TreeHasher.sol";
import {leafDigest as bLeafDigest} from "../tree/binary/TreeHasher.sol";
import "../../../lib/openzeppelin-contracts-upgradeable/contracts/utils/math/MathUpgradeable.sol";
import "forge-std/console.sol";

uint256 constant SUBTREE_ROOT_THRESHOLD = 64;

function _divCeil(uint256 x, uint256 y) pure returns (uint256 z) {
    z = x / y + (x % y == 0 ? 0 : 1);
}

function _numShares(uint256 blobSize) pure returns (uint256) {
    return _divCeil((MathUpgradeable.max(blobSize, 478) - 478), 482) + 1;
}

function _copyNamespace(bytes memory share, bytes29 namespaceBytes) pure {
    for (uint256 i = 0; i < namespaceBytes.length; i++) {
        share[i] = namespaceBytes[i];
    }
}

function _writeInfoByteV0(bytes memory share, bool startingSequence) pure {
    share[29] = bytes1(startingSequence ? 1 : 0);
}

function _writeSequenceLength(bytes memory share, uint32 sequenceLength) pure {
    // Removed the "reverse", because it didn't work- maybe it's already big-endian?
    //bytes4 sequenceLengthBigEndianBytes = bytes4(abi.encodePacked(reverse(sequenceLength)));
    bytes4 sequenceLengthBigEndianBytes = bytes4(abi.encodePacked(sequenceLength));
    share[30] = sequenceLengthBigEndianBytes[0];
    share[31] = sequenceLengthBigEndianBytes[1];
    share[32] = sequenceLengthBigEndianBytes[2];
    share[33] = sequenceLengthBigEndianBytes[3];
}

function _copyBytes(bytes memory buffer, uint32 cursor, bytes memory data, uint32 length) pure returns (uint32) {

    uint256 start = buffer.length - length;
    for (uint256 i = start; i < buffer.length; i++) {
        if (cursor < data.length) {
            buffer[i] = data[cursor];
            cursor++;
        }
        else {
            buffer[i] = 0;
        }
    }
    return cursor;
}

function _bytesToHexString(bytes memory buffer) pure returns (string memory) {

    // Fixed buffer size for hexadecimal convertion
    bytes memory converted = new bytes(buffer.length * 2);

    bytes memory _base = "0123456789abcdef";

    for (uint256 i = 0; i < buffer.length; i++) {
        converted[i * 2] = _base[uint8(buffer[i]) / _base.length];
        converted[i * 2 + 1] = _base[uint8(buffer[i]) % _base.length];
    }

    return string(abi.encodePacked(converted));
}

// Share Version 0
function _bytesToSharesV0(bytes memory blobData, Namespace memory namespace) pure returns (bytes[] memory shares, bool err) {
    if (namespace.version != 0) {
        return (new bytes[](0), true);
    }
    if (isReservedNamespace(namespace)) {
        return (new bytes[](0), true);
    }
    // Allocate memory for the shares
    uint256 numShares = _numShares(blobData.length); 
    bytes[] memory _shares = new bytes[](numShares);
    for (uint256 i = 0; i < _shares.length; i++) {
        _shares[i] = new bytes(512);
    }
    // Get the namespace bytes
    bytes29 namespaceBytes = namespace.toBytes();

    // The first share is special, because it has startingSequence set to true and the 4-byte sequence length
    _copyNamespace(_shares[0], namespaceBytes);
    _writeInfoByteV0(_shares[0], true);
    _writeSequenceLength(_shares[0], uint32(blobData.length));
    uint32 cursor = 0;
    cursor = _copyBytes(_shares[0], cursor, blobData, uint32(478)); //only 478 bytes, because 4 bytes are used for the sequence length

    if (shares.length != 1) {
        // The remaining shares are all the same
        for (uint256 i = 1; i < _shares.length; i++) {
            // Copy the namespace
            _copyNamespace(_shares[i], namespaceBytes);
            // Write the info byte
            _writeInfoByteV0(_shares[i], false);
            // Copy the data
            cursor = _copyBytes(_shares[i], cursor, blobData, uint32(482)); // copy the full 482 bytes
        }
    }

    shares = _shares;
    err = false;
}

function _roundDownPowerOfTwo(uint256 x) pure returns (uint256) {
    uint256 result = 1;
    while (result * 2 <= x) {
        result *= 2;
    }
    return result;
}

function _roundUpPowerOfTwo(uint256 x) pure returns (uint256) {
    uint256 result = 1;
    while (result < x) {
        result *= 2;
    }
    return result;
}

function _blobMinSquareSize(uint256 shareCount) pure returns (uint256) {
    return _roundUpPowerOfTwo(MathUpgradeable.sqrt(shareCount, MathUpgradeable.Rounding.Ceil));
}

function _subtreeWidth(uint256 shareCount, uint256 subtreeRootThreshold) pure returns (uint256) {
    uint256 s = shareCount / subtreeRootThreshold;
    if (s != 0) {
        s++;
    }
    s = _roundUpPowerOfTwo(s);
    return MathUpgradeable.min(s, _blobMinSquareSize(shareCount));
}

function _merkleMountainRangeSizes(uint256 totalSize, uint256 maxTreeSize) pure returns (uint256[] memory) {
    // Overestimate size of array
    // This is a workaround because Solidity doesn't support dynamic memory arrays like Go or Rust
    uint256 bigTrees = totalSize / maxTreeSize;
    uint256 leftovers = totalSize % maxTreeSize;
    uint256 numTrees;
    if (leftovers == 0) {
        numTrees = bigTrees;
    }
    else {
        numTrees = bigTrees + MathUpgradeable.log2(leftovers) + (leftovers%2);
    }
    uint256[] memory treeSizes = new uint256[](numTrees);
    uint256 count = 0;
    while (totalSize != 0) {
        if (totalSize >= maxTreeSize) {
            treeSizes[count] = maxTreeSize;
            totalSize -= maxTreeSize;
        }
        else {
            uint256 treeSize = _roundDownPowerOfTwo(totalSize);
            treeSizes[count] = treeSize;
            totalSize -= treeSize;
        }
        count++;
    }
    return treeSizes;
}

function _createCommitment(bytes[] memory shares, Namespace memory namespace) view returns (bytes32 commitment) {
    uint256 subtreeWidth = _subtreeWidth(shares.length, SUBTREE_ROOT_THRESHOLD);
    uint256[] memory treeSizes = _merkleMountainRangeSizes(shares.length, subtreeWidth);
    bytes[][] memory leafSets = new bytes[][](treeSizes.length);
    uint256 cursor = 0;
    // So much copying...
    // This could likely be optimized, but I'm not an EVM expert
    // Let's see if the gas is too high, optimize later.
    // Stop when we hit 0, the delimeter indicating the end of the array
    for (uint256 i = 0; i < treeSizes.length; i++) {
        leafSets[i] = new bytes[](treeSizes[i]);
        for (uint256 j = 0; j < treeSizes[i]; j++) {
            //leafSets[i][j] = new bytes(512);
            // Try with 512 + 29, prefixing with the 29 byte namespace
            leafSets[i][j] = new bytes(541);
            // copy the share
            for (uint256 k = 0; k < 512; k++) {
                leafSets[i][j][k] = shares[cursor][k];
            }
            // copy the namespace bytes
            for (uint256 k = 512; k < 541; k++) {
                console.log("namespace byte %d: %s", k, _bytesToHexString(abi.encodePacked(namespace.toBytes()[k-512])));
                leafSets[i][j][k] = namespace.toBytes()[k-512];
            }
            cursor += treeSizes[i];
        }
    }

    //NamespaceNode[] memory subtreeRoots = new NamespaceNode[](leafSets.length);
    bytes32[] memory subtreeRoots = new bytes32[](leafSets.length);
    // Fore each leafSet, compute the root using _computeRoot. Pass a null value for the "proof" parameter
    //NamespaceMerkleMultiproof memory nullproof = NamespaceMerkleMultiproof(0, 0, new NamespaceNode[](0));
    for (uint256 i = 0; i < leafSets.length; i++) {
        NamespaceNode[] memory leafNamespaceNodes = new NamespaceNode[](leafSets[i].length);
        for (uint256 j = 0; j < leafSets[i].length; j++) {
            leafNamespaceNodes[j] = leafDigest(namespace, leafSets[i][j]);
        }
        console.log("first node: %s", _bytesToHexString(abi.encodePacked(leafNamespaceNodes[0].digest)));
        NamespaceMerkleMultiproof memory emptyProof = NamespaceMerkleMultiproof(0, leafSets[i].length, new NamespaceNode[](0));
        (NamespaceNode memory root,,,) = NamespaceMerkleTree._computeRoot(emptyProof, leafNamespaceNodes, 0, leafNamespaceNodes.length, 0, 0);
        subtreeRoots[i] = bLeafDigest(bytes(abi.encodePacked(root.digest)));
        console.log("root digest ", _bytesToHexString(abi.encodePacked(root.digest)));
        console.log("subtree root ", _bytesToHexString(abi.encodePacked(subtreeRoots[i])));
    }
    //BinaryMerkleMultiproof memory nullBinaryProof = BinaryMerkleMultiproof(new bytes32[](0), 0, 0);
    BinaryMerkleMultiproof memory emptyBinaryProof = BinaryMerkleMultiproof(new bytes32[](0), 0, subtreeRoots.length);
    (bytes32 binaryTreeRoot,,,) = BinaryMerkleTree._computeRootMulti(emptyBinaryProof, subtreeRoots, 0, subtreeRoots.length, 0, 0);
    commitment = binaryTreeRoot;
}