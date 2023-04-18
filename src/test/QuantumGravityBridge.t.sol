// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import "../lib/openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "../Constants.sol";
import "../DataRootTuple.sol";
import "../QuantumGravityBridge.sol";
import "../lib/tree/binary/BinaryMerkleProof.sol";

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

    /*
    the values used in the verify attestation test are in the format
    `<height padded to 32 bytes || data root || original block size padded to 32 bytes>`,
    which represent an encoded `abi.encode(DataRootTuple)`:

    0x000000000000000000000000000000000000000000000000000000000000000101010101010101010101010101010101010101010101010101010101010101010000000000000000000000000000000000000000000000000000000000000020
    0x000000000000000000000000000000000000000000000000000000000000000202020202020202020202020202020202020202020202020202020202020202020000000000000000000000000000000000000000000000000000000000000020
    0x000000000000000000000000000000000000000000000000000000000000000303030303030303030303030303030303030303030303030303030303030303030000000000000000000000000000000000000000000000000000000000000040
    0x000000000000000000000000000000000000000000000000000000000000000404040404040404040404040404040404040404040404040404040404040404040000000000000000000000000000000000000000000000000000000000000040
    */
    function testVerifyAttestation() public {
        uint256 initialVelsetNonce = 0;
        // data root tuple root nonce.
        uint256 nonce = 1;
        // commitment to a set of roots.
        // these values were generated using the tendermint implementation of binary merkle trees:
        // https://github.com/celestiaorg/celestia-core/blob/60310e7aa554bb76b735a010847a6613bcfe06e8/crypto/merkle/proof.go#L33-L48
        bytes32 newTupleRoot = 0x3db8fb64a10316e447733af826165d0543b94b9de3b4146dcad70611f5089018;
        // a data root committed to by the above tuple root.
        bytes32 newTuple = 0x0202020202020202020202020202020202020202020202020202020202020202;
        // the height of the data root.
        uint256 height = 2;
        uint256 size = 32;
        // the binary merkle proof of the data root to the data root tuple root.
        bytes32[] memory sideNodes = new bytes32[](2);
        sideNodes[0] = 0x1ce3b03bea51b24a60e147396ced1b7eec566129aa43ed98c5b5060b937b4ea8;
        sideNodes[1] = 0x9e7df00a432e3c8c7ec9177f9581d43a6d837a5a05467f6ce5d730fe3c4467d7;
        BinaryMerkleProof memory newTupleProof;
        newTupleProof.sideNodes = sideNodes;
        newTupleProof.key = 1;
        newTupleProof.numLeaves = 4;

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

        DataRootTuple memory t;
        t.height = height;
        t.dataRoot = newTuple;
        t.squareSize = size;

        // verify that the tuple was committed to
        bool committedTo = bridge.verifyAttestation(nonce, t, newTupleProof);
        assertTrue(committedTo);
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
