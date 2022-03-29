// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import "../lib/openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "../DataRootTuple.sol";
import "../QuantumGravityBridge.sol";

import "ds-test/test.sol";

contract RelayerTest is DSTest {
    // private keys used for hardcoded signatures
    // testAddr  = "0x9c2B12b5a07FC6D719Ed7646e5041A7E85758329"
    // testPriv  = "64a1d6f0e760a8d62b4afdde4096f16f51b401eaaecc915740f71770ea76a8ad"
    // testAddr2 = "0xe650B084f05C6194f6e552e3b9f08718Bc8a9d56"
    // testPriv2 = "6e8bdfa979ab645b41c4d17cb1329b2a44684c82b61b1b060ea9b6e1c927a4f4"

    bytes32 constant VALIDATOR_SET_HASH_DOMAIN_SEPARATOR =
        0x636865636b706f696e7400000000000000000000000000000000000000000000;

    bytes32 constant DATA_ROOT_TUPLE_ROOT_DOMAIN_SEPARATOR =
        0x7472616e73616374696f6e426174636800000000000000000000000000000000;

    bytes32 constant BRIDGE_ID =
        VALIDATOR_SET_HASH_DOMAIN_SEPARATOR;

    QuantumGravityBridge bridge;
    
    Validator[] private validators;
    uint256 private votingPower = 3334;
    uint256 private valsetNonce = 0;
    uint256 private dataTupleRootNonce = 0;
    function setUp() public {
        validators.push(Validator(0x9c2B12b5a07FC6D719Ed7646e5041A7E85758329, 5000));
        bytes32 hash = computeValidatorSetHash(validators);
        bridge = new QuantumGravityBridge(BRIDGE_ID, votingPower, hash);
    }

    function testUpdateValidatorSet() public {
        // save the old test validator set before we add to it
        Validator[] memory oldVS = new Validator[](1);
        oldVS[0] = Validator(0x9c2B12b5a07FC6D719Ed7646e5041A7E85758329, 5000);

        // hardcoded signature for the first validator set update
        Signature[] memory sigs = new Signature[](1);
        sigs[0] = Signature(27, 0xbe1d908f5f4f307230f7f6489253ea3051096f0de57377bf1ac218cf7c560a08, 0x4eed1044b69cebcaf45aabc2363b55a7047ce7dd0dd219baec4c4ac3d097c6f8);

        // change test validator set
        validators.push(Validator(0xe650B084f05C6194f6e552e3b9f08718Bc8a9d56, 5000));
        bytes32 newVSHash = keccak256(abi.encode(validators));

        uint256 newNonce = 1;
        uint256 newPowerThreshhold = 6668; 
        bridge.updateValidatorSet(newNonce, newPowerThreshhold, newVSHash, oldVS, sigs);
        bytes32 newCheckpoint = domainSeparateValidatorSetHash(BRIDGE_ID, newNonce, newPowerThreshhold, newVSHash);


        assertEq(bridge.state_lastValidatorSetNonce(), newNonce);
        assertEq(bridge.state_powerThreshold(), newPowerThreshhold);
        assertEq(
            bridge.state_lastValidatorSetCheckpoint(), 
            newCheckpoint
        );
    }

    function testSubmitDataRootTupleRoot() public {
        // hardcoded signature for the first validator set update
        Signature[] memory sigs = new Signature[](1);
        sigs[0] = Signature(28, 0x8a71b11dfc2f5bf6ac3a9e7f8e74a3e6d58aa24957c4a6a8fb6021b6b3c9a57e, 0x796de792753d60cdd79a922516c8345db6661fe1762a59c0aff7698ac4cb7eea);
        
        Validator[] memory vals = new Validator[](1);
        vals[0] = Validator(0x9c2B12b5a07FC6D719Ed7646e5041A7E85758329, 5000);

        uint256 newNonce = 1;
        bytes32 newTupleRoot = VALIDATOR_SET_HASH_DOMAIN_SEPARATOR;
        bridge.submitDataRootTupleRoot(newNonce, newTupleRoot, vals, sigs);

        assertEq(bridge.state_lastDataRootTupleRootNonce(), newNonce);
        assertEq(bridge.state_dataRootTupleRoots(newNonce), newTupleRoot);
    }

    function computeValidatorSetHash(Validator[] memory _validators) private pure returns (bytes32) {
        return keccak256(abi.encode(_validators));
    }

    function domainSeparateValidatorSetHash(
        bytes32 _bridge_id,
        uint256 _nonce,
        uint256 _powerThreshold,
        bytes32 _validatorSetHash
    ) private pure returns (bytes32) {
        bytes32 c = keccak256(
            abi.encode(_bridge_id, VALIDATOR_SET_HASH_DOMAIN_SEPARATOR, _nonce, _powerThreshold, _validatorSetHash)
        );

        return c;
    }
}
