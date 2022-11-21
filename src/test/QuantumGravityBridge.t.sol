// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import "../lib/openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "../Constants.sol";
import "../DataRootTuple.sol";
import "../QuantumGravityBridge.sol";

import "ds-test/test.sol";

interface CheatCodes {
    function addr(uint256 privateKey) external returns (address);

    function sign(uint256 privateKey, bytes32 digest)
        external
        returns (
            uint8 v,
            bytes32 r,
            bytes32 s
        );
}

contract RelayerTest is DSTest {
    // Private keys used for test signatures.
    uint256 constant testPriv1 = 0x64a1d6f0e760a8d62b4afdde4096f16f51b401eaaecc915740f71770ea76a8ad;
    uint256 constant testPriv2 = 0x6e8bdfa979ab645b41c4d17cb1329b2a44684c82b61b1b060ea9b6e1c927a4f4;

    QuantumGravityBridge bridge;

    Validator[] private validators;
    uint256 private votingPower = 5000;
    uint256 private dataTupleRootNonce = 0;

    // Set up Foundry cheatcodes.
    CheatCodes cheats = CheatCodes(HEVM_ADDRESS);

    function setUp() public {
        uint256 initialVelsetNonce = 0;

        validators.push(Validator(cheats.addr(testPriv1), votingPower));
        bytes32 hash = computeValidatorSetHash(validators);
        bridge = new QuantumGravityBridge(initialVelsetNonce, (2 * votingPower) / 3, hash);
    }

    function testUpdateValidatorSet() public {
        uint256 initialVelsetNonce = 0;

        // Save the old test validator set before we add to it.
        Validator[] memory oldVS = new Validator[](1);
        oldVS[0] = Validator(cheats.addr(testPriv1), votingPower);

        uint256 newNonce = 1;
        validators.push(Validator(cheats.addr(testPriv2), votingPower));
        votingPower += votingPower;
        uint256 newPowerThreshold = (2 * votingPower) / 3;
        bytes32 newVSHash = keccak256(abi.encode(validators));
        bytes32 newCheckpoint = domainSeparateValidatorSetHash(newNonce, newPowerThreshold, newVSHash);

        // Signature for the first validator set update.
        Signature[] memory sigs = new Signature[](1);
        bytes32 digest_eip191 = ECDSA.toEthSignedMessageHash(newCheckpoint);
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(testPriv1, digest_eip191);
        sigs[0] = Signature(v, r, s);

        bridge.updateValidatorSet(newNonce, initialVelsetNonce, newPowerThreshold, newVSHash, oldVS, sigs);

        assertEq(bridge.state_eventNonce(), newNonce);
        assertEq(bridge.state_powerThreshold(), newPowerThreshold);
        assertEq(bridge.state_lastValidatorSetCheckpoint(), newCheckpoint);
    }

    function testSubmitDataRootTupleRoot() public {
        uint256 initialVelsetNonce = 0;
        uint256 nonce = 1;
        // 32 bytes, chosen at random.
        bytes32 newTupleRoot = 0x0de92bac0b356560d821f8e7b6f5c9fe4f3f88f6c822283efd7ab51ad56a640e;

        bytes32 newDataRootTupleRoot = domainSeparateDataRootTupleRoot(nonce, newTupleRoot);

        // Signature for the update.
        Signature[] memory sigs = new Signature[](1);
        bytes32 digest_eip191 = ECDSA.toEthSignedMessageHash(newDataRootTupleRoot);
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(testPriv1, digest_eip191);
        sigs[0] = Signature(v, r, s);

        Validator[] memory valSet = new Validator[](1);
        valSet[0] = Validator(cheats.addr(testPriv1), votingPower);

        bridge.submitDataRootTupleRoot(nonce, initialVelsetNonce, newTupleRoot, valSet, sigs);

        assertEq(bridge.state_eventNonce(), nonce);
        assertEq(bridge.state_dataRootTupleRoots(nonce), newTupleRoot);
    }

    function computeValidatorSetHash(Validator[] memory _validators) private pure returns (bytes32) {
        return keccak256(abi.encode(_validators));
    }

    function domainSeparateValidatorSetHash(
        uint256 _nonce,
        uint256 _powerThreshold,
        bytes32 _validatorSetHash
    ) private pure returns (bytes32) {
        bytes32 c = keccak256(
            abi.encode(VALIDATOR_SET_HASH_DOMAIN_SEPARATOR, _nonce, _powerThreshold, _validatorSetHash)
        );

        return c;
    }

    function domainSeparateDataRootTupleRoot(
        uint256 _nonce,
        bytes32 _dataRootTupleRoot
    ) private pure returns (bytes32) {
        bytes32 c = keccak256(
            abi.encode(DATA_ROOT_TUPLE_ROOT_DOMAIN_SEPARATOR, _nonce, _dataRootTupleRoot)
        );

        return c;
    }
}
