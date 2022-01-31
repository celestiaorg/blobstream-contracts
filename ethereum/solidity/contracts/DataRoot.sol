// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

/// @notice A tuple of message with metadata. Each message commitment is
///  associated with a Celestia block height and a namespace ID.
/// @dev https://github.com/celestiaorg/celestia-specs/blob/master/src/specs/data_structures.md#signedtransactiondatapayformessage
struct MessageTuple {
    // Celestia block height the message was included in.
    // Genesis block is height = 1.
    uint256 height;
    // Namespace ID of the message.
    bytes8 namespaceID;
    // Commitment to the message.
    bytes32 messageCommitment;
}
