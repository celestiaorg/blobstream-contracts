// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";
import "openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import "openzeppelin-contracts-upgradeable/contracts/proxy/utils/UUPSUpgradeable.sol";
import "openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";

import "./Constants.sol";
import "./DataRootTuple.sol";
import "./IDAOracle.sol";
import "./lib/tree/binary/BinaryMerkleProof.sol";
import "./lib/tree/binary/BinaryMerkleTree.sol";

struct Validator {
    address addr;
    uint256 power;
}

struct Signature {
    uint8 v;
    bytes32 r;
    bytes32 s;
}

/// @title Blobstream: Celestia -> EVM, Data Availability relay.
/// @dev The relay relies on a set of signers to attest to some event on
/// Celestia. These signers are the Celestia validator set, who sign over every
/// Celestia block. Keeping track of the Celestia validator set is accomplished
/// by updating this contract's view of the validator set with
/// `updateValidatorSet()`. At least 2/3 of the voting power of the current
/// view of the validator set must sign off on new relayed events, submitted
/// with `submitDataRootTupleRoot()`. Each event is a batch of `DataRootTuple`s
/// (see ./DataRootTuple.sol), with each tuple representing a single data root
/// in a Celestia block header. Relayed tuples are in the same order as the
/// block headers.
/// @dev DO NOT REMOVE INHERITANCE OF THE FOLLOWING CONTRACTS: Initializable, UUPSUpgradeable and
/// OwnableUpgradeable! They're essential for upgradability.
contract Blobstream is IDAOracle, Initializable, UUPSUpgradeable, OwnableUpgradeable {
    // Don't change the order of state for working upgrades AND BE AWARE OF
    // INHERITANCE VARIABLES! Inherited contracts contain storage slots and must
    // be accounted for in any upgrades. Always test an exact upgrade on testnet
    // and localhost before mainnet upgrades.

    /////////////
    // Storage //
    /////////////

    /// @notice Domain-separated commitment to the latest validator set.
    bytes32 public state_lastValidatorSetCheckpoint;
    /// @notice Voting power required to submit a new update.
    uint256 public state_powerThreshold;
    /// @notice Nonce for bridge events. Must be incremented sequentially.
    uint256 public state_eventNonce;
    /// @notice Mapping of data root tuple root nonces to data root tuple roots.
    mapping(uint256 => bytes32) public state_dataRootTupleRoots;

    ////////////
    // Events //
    ////////////

    /// @notice Emitted when a new root of data root tuples is relayed.
    /// @param nonce Event nonce.
    /// @param dataRootTupleRoot Merkle root of relayed data root tuples.
    /// See `submitDataRootTupleRoot`.
    event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot);

    /// @notice Emitted when the validator set is updated.
    /// @param nonce Event nonce.
    /// @param powerThreshold New voting power threshold.
    /// @param validatorSetHash Hash of new validator set.
    /// See `updateValidatorSet`.
    event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash);

    ////////////
    // Errors //
    ////////////

    /// @notice Malformed current validator set.
    error MalformedCurrentValidatorSet();

    /// @notice Validator signature does not match.
    error InvalidSignature();

    /// @notice Submitted validator set signatures do not have enough power.
    error InsufficientVotingPower();

    /// @notice New validator set nonce must be greater than the current nonce.
    error InvalidValidatorSetNonce();

    /// @notice Supplied current validators and powers do not match checkpoint.
    error SuppliedValidatorSetInvalid();

    /// @notice Data root tuple root nonce must be greater than the current nonce.
    error InvalidDataRootTupleRootNonce();

    ///////////////
    // Functions //
    ///////////////

    /// @param _nonce Initial event nonce.
    /// @param _powerThreshold Initial voting power that is needed to approve
    /// operations.
    /// @param _validatorSetCheckpoint Initial checkpoint of the validator set. This does not need
    /// to be the genesis validator set of the bridged chain, only the initial
    /// validator set of the bridge.
    /// @dev DO NOT REMOVE THE INITIALIZER! It is mandatory for upgradability.
    function initialize(uint256 _nonce, uint256 _powerThreshold, bytes32 _validatorSetCheckpoint) public initializer {
        // EFFECTS

        state_eventNonce = _nonce;
        state_lastValidatorSetCheckpoint = _validatorSetCheckpoint;
        state_powerThreshold = _powerThreshold;

        /// @dev Initialize the OwnableUpgradeable explicitly.
        /// DO NOT REMOVE! It is mandatory for allowing the owner to upgrade.
        __Ownable_init(_msgSender());
    }

    /// @dev only authorize the upgrade for the owner of the contract.
    /// Additional access control logic can be added to allow multiple actors to upgrade.
    /// @dev DO NOT REMOVE! It is mandatory for upgradability.
    function _authorizeUpgrade(address) internal override onlyOwner {}

    /// @notice Utility function to check if a signature is nil.
    /// If all bytes of the 65-byte signature are zero, then it's a nil signature.
    function isSigNil(Signature calldata _sig) private pure returns (bool) {
        return (_sig.r == 0 && _sig.s == 0 && _sig.v == 0);
    }

    /// @notice Utility function to verify EIP-191 signatures.
    function verifySig(address _signer, bytes32 _digest, Signature calldata _sig) private pure returns (bool) {
        bytes32 digest_eip191 = ECDSA.toEthSignedMessageHash(_digest);

        return _signer == ECDSA.recover(digest_eip191, _sig.v, _sig.r, _sig.s);
    }

    /// @dev Computes the hash of a validator set.
    /// @param _validators The validator set to hash.
    function computeValidatorSetHash(Validator[] calldata _validators) private pure returns (bytes32) {
        return keccak256(abi.encode(_validators));
    }

    /// @dev Make a domain-separated commitment to the validator set.
    /// A hash of all relevant information about the validator set.
    /// The format of the hash is:
    ///     keccak256(VALIDATOR_SET_HASH_DOMAIN_SEPARATOR, nonce, power_threshold, validator_set_hash)
    /// The elements in the validator set should be monotonically decreasing by power.
    /// @param _nonce Nonce.
    /// @param _powerThreshold The voting power threshold.
    /// @param _validatorSetHash Validator set hash.
    function domainSeparateValidatorSetHash(uint256 _nonce, uint256 _powerThreshold, bytes32 _validatorSetHash)
        private
        pure
        returns (bytes32)
    {
        bytes32 c =
            keccak256(abi.encode(VALIDATOR_SET_HASH_DOMAIN_SEPARATOR, _nonce, _powerThreshold, _validatorSetHash));

        return c;
    }

    /// @dev Make a domain-separated commitment to a data root tuple root.
    /// A hash of all relevant information about a data root tuple root.
    /// The format of the hash is:
    ///     keccak256(DATA_ROOT_TUPLE_ROOT_DOMAIN_SEPARATOR, nonce, dataRootTupleRoot)
    /// @param _nonce Event nonce.
    /// @param _dataRootTupleRoot Data root tuple root.
    function domainSeparateDataRootTupleRoot(uint256 _nonce, bytes32 _dataRootTupleRoot)
        private
        pure
        returns (bytes32)
    {
        bytes32 c = keccak256(abi.encode(DATA_ROOT_TUPLE_ROOT_DOMAIN_SEPARATOR, _nonce, _dataRootTupleRoot));

        return c;
    }

    /// @dev Checks that enough voting power signed over a digest.
    /// It expects the signatures to be in the same order as the _currentValidators.
    /// @param _currentValidators The current validators.
    /// @param _sigs The current validators' signatures.
    /// @param _digest This is what we are checking they have signed.
    /// @param _powerThreshold At least this much power must have signed.
    function checkValidatorSignatures(
        // The current validator set and their powers
        Validator[] calldata _currentValidators,
        Signature[] calldata _sigs,
        bytes32 _digest,
        uint256 _powerThreshold
    ) private pure {
        uint256 cumulativePower = 0;
        for (uint256 i = 0; i < _currentValidators.length; i++) {
            // If the signature is nil, then it's not present so continue.
            if (isSigNil(_sigs[i])) {
                continue;
            }

            // Check that the current validator has signed off on the hash.
            if (!verifySig(_currentValidators[i].addr, _digest, _sigs[i])) {
                revert InvalidSignature();
            }

            // Sum up cumulative power.
            cumulativePower += _currentValidators[i].power;

            // Break early to avoid wasting gas.
            if (cumulativePower >= _powerThreshold) {
                break;
            }
        }
        // Check that there was enough power.
        if (cumulativePower < _powerThreshold) {
            revert InsufficientVotingPower();
        }
    }

    /// @notice This updates the validator set by checking that the validators
    /// in the current validator set have signed off on the new validator set.
    /// The signatures supplied are the signatures of the current validator set
    /// over the checkpoint hash generated from the new validator set. Anyone
    /// can call this function, but they must supply valid signatures of the
    /// current validator set over the new validator set.
    ///
    /// The validator set hash that is signed over is domain separated as per
    /// `domainSeparateValidatorSetHash`.
    /// @param _newNonce The new event nonce.
    /// @param _oldNonce The nonce of the latest update to the validator set.
    /// @param _newPowerThreshold At least this much power must have signed.
    /// @param _newValidatorSetHash The hash of the new validator set.
    /// @param _currentValidatorSet The current validator set.
    /// @param _sigs Signatures.
    function updateValidatorSet(
        uint256 _newNonce,
        uint256 _oldNonce,
        uint256 _newPowerThreshold,
        bytes32 _newValidatorSetHash,
        Validator[] calldata _currentValidatorSet,
        Signature[] calldata _sigs
    ) external {
        // CHECKS

        uint256 currentNonce = state_eventNonce;
        uint256 currentPowerThreshold = state_powerThreshold;
        bytes32 lastValidatorSetCheckpoint = state_lastValidatorSetCheckpoint;

        // Check that the new nonce is one more than the current one.
        if (_newNonce != currentNonce + 1) {
            revert InvalidValidatorSetNonce();
        }

        // Check that current validators and signatures are well-formed.
        if (_currentValidatorSet.length != _sigs.length) {
            revert MalformedCurrentValidatorSet();
        }

        // Check that the supplied current validator set matches the saved checkpoint.
        bytes32 currentValidatorSetHash = computeValidatorSetHash(_currentValidatorSet);
        if (
            domainSeparateValidatorSetHash(_oldNonce, currentPowerThreshold, currentValidatorSetHash)
                != lastValidatorSetCheckpoint
        ) {
            revert SuppliedValidatorSetInvalid();
        }

        // Check that enough current validators have signed off on the new validator set.
        bytes32 newCheckpoint = domainSeparateValidatorSetHash(_newNonce, _newPowerThreshold, _newValidatorSetHash);
        checkValidatorSignatures(_currentValidatorSet, _sigs, newCheckpoint, currentPowerThreshold);

        // EFFECTS

        state_lastValidatorSetCheckpoint = newCheckpoint;
        state_powerThreshold = _newPowerThreshold;
        state_eventNonce = _newNonce;

        // LOGS

        emit ValidatorSetUpdatedEvent(_newNonce, _newPowerThreshold, _newValidatorSetHash);
    }

    /// @notice Relays a root of Celestia data root tuples to an EVM chain. Anyone
    /// can call this function, but they must supply valid signatures of the
    /// current validator set over the data root tuple root.
    ///
    /// The data root root is the Merkle root of the binary Merkle tree
    /// (https://github.com/celestiaorg/celestia-specs/blob/master/src/specs/data_structures.md#binary-merkle-tree)
    /// where each leaf is in an ABI-encoded `DataRootTuple`. Each relayed data
    /// root tuple will 1:1 mirror data roots as they are included in headers
    /// on Celestia, _in order of inclusion_.
    ///
    /// The data tuple root that is signed over is domain separated as per
    /// `domainSeparateDataRootTupleRoot`.
    /// @param _newNonce The new event nonce.
    /// @param _validatorSetNonce The nonce of the latest update to the
    /// validator set.
    /// @param _dataRootTupleRoot The Merkle root of data root tuples.
    /// @param _currentValidatorSet The current validator set.
    /// @param _sigs Signatures.
    function submitDataRootTupleRoot(
        uint256 _newNonce,
        uint256 _validatorSetNonce,
        bytes32 _dataRootTupleRoot,
        Validator[] calldata _currentValidatorSet,
        Signature[] calldata _sigs
    ) external {
        // CHECKS

        uint256 currentNonce = state_eventNonce;
        uint256 currentPowerThreshold = state_powerThreshold;
        bytes32 lastValidatorSetCheckpoint = state_lastValidatorSetCheckpoint;

        // Check that the new nonce is one more than the current one.
        if (_newNonce != currentNonce + 1) {
            revert InvalidDataRootTupleRootNonce();
        }

        // Check that current validators and signatures are well-formed.
        if (_currentValidatorSet.length != _sigs.length) {
            revert MalformedCurrentValidatorSet();
        }

        // Check that the supplied current validator set matches the saved checkpoint.
        bytes32 currentValidatorSetHash = computeValidatorSetHash(_currentValidatorSet);
        if (
            domainSeparateValidatorSetHash(_validatorSetNonce, currentPowerThreshold, currentValidatorSetHash)
                != lastValidatorSetCheckpoint
        ) {
            revert SuppliedValidatorSetInvalid();
        }

        // Check that enough current validators have signed off on the data
        // root tuple root and nonce.
        bytes32 c = domainSeparateDataRootTupleRoot(_newNonce, _dataRootTupleRoot);
        checkValidatorSignatures(_currentValidatorSet, _sigs, c, currentPowerThreshold);

        // EFFECTS

        state_eventNonce = _newNonce;
        state_dataRootTupleRoots[_newNonce] = _dataRootTupleRoot;

        // LOGS

        emit DataRootTupleRootEvent(_newNonce, _dataRootTupleRoot);
    }

    /// @dev see "./IDAOracle.sol"
    function verifyAttestation(uint256 _tupleRootNonce, DataRootTuple memory _tuple, BinaryMerkleProof memory _proof)
        external
        view
        override
        returns (bool)
    {
        // Tuple must have been committed before.
        if (_tupleRootNonce > state_eventNonce) {
            return false;
        }

        // Load the tuple root at the given index from storage.
        bytes32 root = state_dataRootTupleRoots[_tupleRootNonce];

        // Verify the proof.
        (bool isProofValid,) = BinaryMerkleTree.verify(root, _proof, abi.encode(_tuple));

        return isProofValid;
    }
}
