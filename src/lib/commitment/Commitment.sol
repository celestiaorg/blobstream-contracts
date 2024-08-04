// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Namespace} from "../tree/Types.sol";
import "openzeppelin-contracts/contracts/utils/math/Math.sol";
import "forge-std/console.sol";

// Turn little-endian into big-endian
// Source: https://ethereum.stackexchange.com/a/83627
function reverse(uint32 input) pure returns (uint32 v) {
    v = input;

    // swap bytes
    v = ((v & 0xFF00FF00) >> 8) |
        ((v & 0x00FF00FF) << 8);

    // swap 2-byte long pairs
    v = (v >> 16) | (v << 16);
}

function div_ceil(uint256 x, uint256 y) pure returns (uint256 z) {
    z = x / y + (x % y == 0 ? 0 : 1);
}

function num_shares(uint256 blobSize) pure returns (uint256) {
    return div_ceil((Math.max(blobSize, 478) - 478), 482) + 1;
}

function copyNamespace(bytes memory share, bytes29 namespaceBytes) {
    for (uint256 i = 0; i < namespaceBytes.length; i++) {
        share[i] = namespaceBytes[i];
    }
}

function writeInfoByteV0(bytes memory share, bool startingSequence) pure {
    share[29] = bytes1(startingSequence ? 1 : 0);
}

function writeSequenceLength(bytes memory share, uint32 sequenceLength) pure {
    // Removed the "reverse", because it didn't work- apparently it's already big-endian?
    //bytes4 sequenceLengthBigEndianBytes = bytes4(abi.encodePacked(reverse(sequenceLength)));
    bytes4 sequenceLengthBigEndianBytes = bytes4(abi.encodePacked(sequenceLength));
    share[30] = sequenceLengthBigEndianBytes[0];
    share[31] = sequenceLengthBigEndianBytes[1];
    share[32] = sequenceLengthBigEndianBytes[2];
    share[33] = sequenceLengthBigEndianBytes[3];
}

function copyBytes(bytes memory buffer, uint32 cursor, bytes memory data, uint32 length) returns (uint32) {

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

function bytesToHexString(bytes memory buffer) pure returns (string memory) {

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
function bytesToShares(bytes memory blobData, Namespace memory namespace) returns (bytes[] memory shares) {
    // Allocate memory for the shares
    uint256 numShares = num_shares(blobData.length); 
    bytes[] memory _shares = new bytes[](numShares);
    for (uint256 i = 0; i < _shares.length; i++) {
        _shares[i] = new bytes(512);
    }
    // Get the namespace bytes
    bytes29 namespaceBytes = namespace.toBytes();

    // The first share is special, because it has startingSequence set to true and the 4-byte sequence length
    copyNamespace(_shares[0], namespaceBytes);
    writeInfoByteV0(_shares[0], true);
    writeSequenceLength(_shares[0], uint32(blobData.length));
    uint32 cursor = 0;
    cursor = copyBytes(_shares[0], cursor, blobData, uint32(478)); //only 478 bytes, because 4 bytes are used for the sequence length

    if (shares.length != 1) {
        // The remaining shares are all the same
        for (uint256 i = 1; i < _shares.length; i++) {
            // Copy the namespace
            copyNamespace(_shares[i], namespaceBytes);
            // Write the info byte
            writeInfoByteV0(_shares[i], false);
            // Copy the data
            cursor = copyBytes(_shares[i], cursor, blobData, uint32(482)); // copy the full 482 bytes
        }
    }

    shares = _shares;
}