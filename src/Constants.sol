// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

/// @dev bytes32 encoding of the string "checkpoint"
bytes32 constant VALIDATOR_SET_HASH_DOMAIN_SEPARATOR =
    0x636865636b706f696e7400000000000000000000000000000000000000000000;

/// @dev bytes32 encoding of the string "transactionBatch"
bytes32 constant DATA_ROOT_TUPLE_ROOT_DOMAIN_SEPARATOR =
    0x7472616e73616374696f6e426174636800000000000000000000000000000000;
