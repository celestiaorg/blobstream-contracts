// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";

import "../Constants.sol";
import "../DataRootTuple.sol";
import "../Blobstream.sol";
import "../lib/tree/binary/BinaryMerkleProof.sol";

import "ds-test/test.sol";

interface CheatCodes {
    function addr(uint256 privateKey) external returns (address);

    function sign(uint256 privateKey, bytes32 digest) external returns (uint8 v, bytes32 r, bytes32 s);
}

contract RelayerTest is DSTest {
    // Private keys used for test signatures.
    uint256 constant testPriv1 = 0x64a1d6f0e760a8d62b4afdde4096f16f51b401eaaecc915740f71770ea76a8ad;
    uint256 constant testPriv2 = 0x6e8bdfa979ab645b41c4d17cb1329b2a44684c82b61b1b060ea9b6e1c927a4f4;

    Blobstream bridge;

    Validator[] private validators;
    uint256 private votingPower = 5000;

    // Set up Foundry cheatcodes.
    CheatCodes cheats = CheatCodes(HEVM_ADDRESS);

    function setUp() public {
        uint256 initialValsetNonce = 1;

        validators.push(Validator(cheats.addr(testPriv1), votingPower));
        bytes32 hash = computeValidatorSetHash(validators);
        bytes32 validatorSetCheckpoint = domainSeparateValidatorSetHash(initialValsetNonce, (2 * votingPower) / 3, hash);
        bridge = new Blobstream();
        bridge.initialize(initialValsetNonce, (2 * votingPower) / 3, validatorSetCheckpoint);
    }

    function testUpdateValidatorSet() public {
        uint256 initialValsetNonce = 1;

        // Save the old test validator set before we add to it.
        Validator[] memory oldVS = new Validator[](1);
        oldVS[0] = Validator(cheats.addr(testPriv1), votingPower);

        uint256 newNonce = 2;
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

        bridge.updateValidatorSet(newNonce, initialValsetNonce, newPowerThreshold, newVSHash, oldVS, sigs);

        assertEq(bridge.state_eventNonce(), newNonce);
        assertEq(bridge.state_powerThreshold(), newPowerThreshold);
        assertEq(bridge.state_lastValidatorSetCheckpoint(), newCheckpoint);
    }

    function testSubmitDataRootTupleRoot() public {
        uint256 initialValsetNonce = 1;
        uint256 nonce = 2;
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

        bridge.submitDataRootTupleRoot(nonce, initialValsetNonce, newTupleRoot, valSet, sigs);

        assertEq(bridge.state_eventNonce(), nonce);
        assertEq(bridge.state_dataRootTupleRoots(nonce), newTupleRoot);
    }

    function testDeployContractAtCustomNonce() public {
        uint256 initialValsetNonce = 1;
        uint256 targetNonce = 200;

        Validator[] memory valSet = new Validator[](1);
        valSet[0] = Validator(cheats.addr(testPriv1), votingPower);

        bytes32 hash = computeValidatorSetHash(valSet);
        bytes32 validatorSetCheckpoint = domainSeparateValidatorSetHash(initialValsetNonce, (2 * votingPower) / 3, hash);
        Blobstream newBridge = new Blobstream();
        newBridge.initialize(targetNonce, (2 * votingPower) / 3, validatorSetCheckpoint);

        // 32 bytes, chosen at random.
        bytes32 newTupleRoot = 0x0de92bac0b356560d821f8e7b6f5c9fe4f3f88f6c822283efd7ab51ad56a640e;

        bytes32 newDataRootTupleRoot = domainSeparateDataRootTupleRoot(targetNonce + 1, newTupleRoot);

        // Signature for the update.
        Signature[] memory sigs = new Signature[](1);
        bytes32 digest_eip191 = ECDSA.toEthSignedMessageHash(newDataRootTupleRoot);
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(testPriv1, digest_eip191);
        sigs[0] = Signature(v, r, s);

        newBridge.submitDataRootTupleRoot(targetNonce + 1, initialValsetNonce, newTupleRoot, valSet, sigs);

        assertEq(newBridge.state_eventNonce(), targetNonce + 1);
        assertEq(newBridge.state_dataRootTupleRoots(targetNonce + 1), newTupleRoot);
    }

    /*
    the values used in the verify attestation test are in the format `<height padded to 32 bytes || data root>`, which
    represent an encoded `abi.encode(DataRootTuple)`:

    0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
    0x00000000000000000000000000000000000000000000000000000000000000010101010101010101010101010101010101010101010101010101010101010101
    0x00000000000000000000000000000000000000000000000000000000000000020202020202020202020202020202020202020202020202020202020202020202
    0x00000000000000000000000000000000000000000000000000000000000000030303030303030303030303030303030303030303030303030303030303030303
    */
    function testVerifyAttestation() public {
        uint256 initialValsetNonce = 1;
        // data root tuple root nonce.
        uint256 nonce = 2;
        // commitment to a set of roots.
        // these values were generated using the tendermint implementation of binary merkle trees:
        // https://github.com/celestiaorg/celestia-core/blob/60310e7aa554bb76b735a010847a6613bcfe06e8/crypto/merkle/proof.go#L33-L48
        bytes32 newTupleRoot = 0x82dc1607d84557d3579ce602a45f5872e821c36dbda7ec926dfa17ebc8d5c013;
        // a data root committed to by the above tuple root.
        bytes32 newTuple = 0x0101010101010101010101010101010101010101010101010101010101010101;
        // the height of the data root.
        uint256 height = 1;
        // the binary merkle proof of the data root to the data root tuple root.
        bytes32[] memory sideNodes = new bytes32[](2);
        sideNodes[0] = 0x98ce42deef51d40269d542f5314bef2c7468d401ad5d85168bfab4c0108f75f7;
        sideNodes[1] = 0x575664048c9e64260eca2304d177b11d1566d0c954f1417fc76a4f9f27350063;
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

        bridge.submitDataRootTupleRoot(nonce, initialValsetNonce, newTupleRoot, valSet, sigs);

        assertEq(bridge.state_eventNonce(), nonce);
        assertEq(bridge.state_dataRootTupleRoots(nonce), newTupleRoot);

        DataRootTuple memory t;
        t.height = height;
        t.dataRoot = newTuple;

        // verify that the tuple was committed to
        bool committedTo = bridge.verifyAttestation(nonce, t, newTupleProof);
        assertTrue(committedTo);
    }

    function computeValidatorSetHash(Validator[] memory _validators) private pure returns (bytes32) {
        return keccak256(abi.encode(_validators));
    }

    function domainSeparateValidatorSetHash(uint256 _nonce, uint256 _powerThreshold, bytes32 _validatorSetHash)
        private
        pure
        returns (bytes32)
    {
        bytes32 c =
            keccak256(abi.encode(VALIDATOR_SET_HASH_DOMAIN_SEPARATOR, _nonce, _powerThreshold, _validatorSetHash));

        return c;
    }

    function domainSeparateDataRootTupleRoot(uint256 _nonce, bytes32 _dataRootTupleRoot)
        private
        pure
        returns (bytes32)
    {
        bytes32 c = keccak256(abi.encode(DATA_ROOT_TUPLE_ROOT_DOMAIN_SEPARATOR, _nonce, _dataRootTupleRoot));

        return c;
    }
}
