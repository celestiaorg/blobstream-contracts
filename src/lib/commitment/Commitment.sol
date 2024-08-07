// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Namespace, isReservedNamespace} from "../tree/Types.sol";
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

function _merkleMountainRangeSizes(uint256 totalSize, maxTreeSize) pure returns (uint256[] memory) {
    uint256[] treeSizes = new uint256[](0);
    while (totalSize != 0) {
        if (totalSize >= maxTreeSize) {
            treeSizes.push(maxTreeSize);
            totalSize -= maxTreeSize;
        }
        else {
            uint256 treeSize = _roundDownPowerOfTwo(totalSize);
            treeSizes.push(treeSize);
            totalSize -= treeSize;
        }
    }
    return treeSizes;
}

function _createCommitment(bytes[] memory shares) pure returns (bytes32 commitment) {
    uint256 subtreeWidth = _subtreeWidth(shares.length, SUBTREE_ROOT_THRESHOLD);
    uint256[] treeSizes = _merkleMountainRangeSizes(shares.length, subtreeWidth);
    bytes[][][] leafSets = new bytes[][][](treeSizes.length);
    uint256 cursor = 0;
    for (uint256 i = 0; i < treeSizes.length; i++) {
        leafSets[i] = new bytes[][](treeSizes[i]);

    }
}