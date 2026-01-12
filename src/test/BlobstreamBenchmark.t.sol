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
    function deriveKey(string calldata, string calldata, uint32) external returns (uint256);
}

/// @notice Example command to run the benchmark:
/// `forge test --match-test testBenchmarkSubmitDataRootTupleRoot -vvvvvv --gas-report`.
/// To change the validator set size, change the `numberOfValidators` constant.
/// To make custom calculations of the gas, you can use the `gasleft()` solidity
/// built-in function.
/// The following answer has some insights on using that:
/// https://ethereum.stackexchange.com/a/132325/65649
/// The gas estimations might not be accurate to the real cost in a real network,
/// and that's because foundry doesn't track calldata cost. source:
/// https://github.com/foundry-rs/foundry/issues/3475#issuecomment-1469940917
/// To have accurate results, make sure to add the following costs:
/// A byte of calldata costs either 4 gas (if it is zero) or 16 gas (if it is any other value).
contract Benchmark is DSTest {
    uint256 private constant numberOfValidators = 100;
    uint256 private constant numberOfSigners = 30;

    // Private keys used for test signatures.
    uint256[] private privateKeys;

    Blobstream private bridge;

    Validator[] private validators;
    uint256 private totalValidatorPower = 1000000;
    uint256 private dataTupleRootNonce = 0;

    // Set up Foundry cheatcodes.
    CheatCodes cheats = CheatCodes(HEVM_ADDRESS);

    function setUp() public {
        uint256 initialValsetNonce = 0;
        privateKeys = derivePrivateKeys(numberOfValidators);
        validators = initializeValidators(privateKeys);

        bytes32 hash = computeValidatorSetHash(validators);
        bytes32 checkpoint = domainSeparateValidatorSetHash(initialValsetNonce, (2 * totalValidatorPower) / 3, hash);
        bridge = new Blobstream();
        bridge.initialize(initialValsetNonce, (2 * totalValidatorPower) / 3, checkpoint);
    }

    function testBenchmarkSubmitDataRootTupleRoot() public {
        uint256 initialValsetNonce = 0;
        uint256 nonce = 1;

        // 32 bytes, chosen at random.
        bytes32 newTupleRoot = 0x0de92bac0b356560d821f8e7b6f5c9fe4f3f88f6c822283efd7ab51ad56a640e;
        bytes32 newDataRootTupleRoot = domainSeparateDataRootTupleRoot(nonce, newTupleRoot);

        // Signature for the update.
        Signature[] memory sigs = new Signature[](numberOfValidators);
        bytes32 digest_eip191 = ECDSA.toEthSignedMessageHash(newDataRootTupleRoot);
        uint256 threshold = 2 * totalValidatorPower / 3;
        uint256 cumulatedPower = 0;
        for (uint256 i = 0; i < numberOfValidators; i++) {
            if (cumulatedPower > threshold) {
                break;
            }
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(privateKeys[i], digest_eip191);
            sigs[i] = Signature(v, r, s);
            cumulatedPower += validators[i].power;
        }

        // these are called here so that they're part of the gas report.
        //        uint256 currentPowerThreshold = (2 * votingPower * numberOfValidators) / 3;
        //        bytes32 currentValidatorSetHash = bridge.computeValidatorSetHash(validators);
        //        bridge.domainSeparateValidatorSetHash(nonce, currentPowerThreshold, currentValidatorSetHash);
        //        bridge.checkValidatorSignatures(validators, sigs, newDataRootTupleRoot, currentPowerThreshold);

        bridge.submitDataRootTupleRoot(nonce, initialValsetNonce, newTupleRoot, validators, sigs);
    }

    function computeValidatorSetHash(Validator[] memory _validators) private pure returns (bytes32) {
        return keccak256(abi.encode(_validators));
    }

    function domainSeparateValidatorSetHash(uint256 _nonce, uint256 _powerThreshold, bytes32 _validatorSetHash)
        private
        pure
        returns (bytes32)
    {
        bytes32 c = keccak256(
            abi.encode(VALIDATOR_SET_HASH_DOMAIN_SEPARATOR, _nonce, _powerThreshold, _validatorSetHash)
        );

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

    function derivePrivateKeys(uint256 count) private returns (uint256[] memory) {
        string memory mnemonic = "test test test test test test test test test test test junk";
        uint256[] memory keys = new uint256[](count);
        for (uint32 i = 0; i < count; i++) {
            keys[i] = cheats.deriveKey(mnemonic, "m/44'/60'/0'/0", i);
        }
        return keys;
    }

    function initializeValidators(uint256[] memory keys) private returns (Validator[] memory) {
        Validator[] memory vs = new Validator[](keys.length);
        uint256 threshold = 2 * totalValidatorPower / 3;
        uint256 primaryPower = threshold / (numberOfSigners - 1);
        uint256 secondaryPower = (totalValidatorPower - threshold) / (numberOfValidators - numberOfSigners + 1);
        for (uint256 i = 0; i < keys.length; i++) {
            if (i < numberOfSigners) {
                vs[i] = Validator(cheats.addr(keys[i]), primaryPower);
            } else {
                vs[i] = Validator(cheats.addr(keys[i]), secondaryPower);
            }
        }
        return vs;
    }
}
