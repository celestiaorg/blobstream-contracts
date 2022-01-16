// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "./MessageTuple.sol";
import "./lib/tree/binary/BinaryMerkleProof.sol";

/// @notice Data Availability Oracle interface.
interface IDAOracle {
    /// @notice Verify a Data Availability attestation.
    function verifyAttestation(
        bytes32 tupleRoot,
        MessageTuple memory tuple,
        BinaryMerkleProof memory proof
    ) external view returns (bool);
}
