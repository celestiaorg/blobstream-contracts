// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "./OwnableUpgradeableWithExpiry.sol";

struct Validator {
    address addr;
    uint256 power;
}

struct Signature {
    uint8 v;
    bytes32 r;
    bytes32 s;
}

/// @title Quantum Gravity Bridge: Celestia -> Ethereum, Data Availability relay.
contract QuantumGravityBridge is OwnableUpgradeableWithExpiry {
    // Don't change the order of state for working upgrades.
    // AND BE AWARE OF INHERITANCE VARIABLES!
    // Inherited contracts contain storage slots and must be accounted for in any upgrades
    // always test an exact upgrade on testnet and localhost before mainnet upgrades.

    ///////////////
    // Constants //
    ///////////////

    // bytes32 encoding of the string "checkpoint"
    bytes32 constant VALIDATOR_SET_HASH_DOMAIN_SEPARATOR =
        0x636865636b706f696e7400000000000000000000000000000000000000000000;

    // bytes32 encoding of the string "transactionBatch"
    bytes32 constant MESSAGE_ROOT_DOMAIN_SEPARATOR = 0x7472616e73616374696f6e426174636800000000000000000000000000000000;

    ////////////////
    // Immutables //
    ////////////////

    bytes32 public immutable BRIDGE_ID;

    /////////////
    // Storage //
    /////////////

    bytes32 public s_lastValidatorSetCheckpoint;
    uint256 public s_powerThreshold;
    uint256 public s_lastValidatorSetNonce;
    uint256 public s_lastMessageRootNonce;
    mapping(uint256 => bytes32) public s_messageRoots;

    ////////////
    // Events //
    ////////////

    /// @notice Emitted when a new root of messages is relayed.
    /// @param nonce Nonce.
    /// @param messagesRoot Merkle root of relayed messages.
    event MessageRootEvent(uint256 indexed nonce, bytes32 messagesRoot);

    /// @notice Emitted when the validator set is updated.
    /// @param nonce Nonce.
    /// @param powerThreshold New voting power threshold.
    /// @param validatorSetHash Hash of new validator set.
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

    /// @notice Message root nonce nonce must be greater than the current nonce.
    error InvalidMessageRootNonce();

    ///////////////
    // Functions //
    ///////////////

    /// @param _bridge_id Identifier of the bridge, used in signatures for
    /// domain separation.
    /// @param _powerThreshold Initial voting power that is needed to approve
    /// operations.
    /// @param _validatorSetHash Initial validator set hash. This does not need
    /// to be the genesis validator set of the bridged chain, only the initial
    /// validator set of the bridge.
    constructor(
        bytes32 _bridge_id,
        uint256 _powerThreshold,
        bytes32 _validatorSetHash
    ) {
        __Ownable_init_unchained();

        BRIDGE_ID = _bridge_id;

        // CHECKS

        uint256 nonce = 0;
        bytes32 newCheckpoint = domainSeparateValidatorSetHash(_bridge_id, nonce, _powerThreshold, _validatorSetHash);

        // EFFECTS

        s_lastValidatorSetCheckpoint = newCheckpoint;
        s_powerThreshold = _powerThreshold;

        // LOGS

        emit ValidatorSetUpdatedEvent(nonce, _powerThreshold, _validatorSetHash);
    }

    /// @notice Utility function to verify EIP-191 signatures
    function verifySig(
        address _signer,
        bytes32 _digest,
        Signature calldata _sig
    ) private pure returns (bool) {
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
    ///     keccak256(bridge_id, VALIDATOR_SET_HASH_DOMAIN_SEPARATOR, nonce, power_threshold, validator_set_hash)
    /// The elements in the validator set should be monotonically decreasing by power.
    /// @param _bridge_id Bridge ID.
    /// @param _nonce Nonce.
    /// @param _powerThreshold The voting power threshold.
    /// @param _validatorSetHash Validator set hash.
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

    /// @dev Make a domain-separated commitment to a message root.
    /// A hash of all relevant information about a message root.
    /// The format of the hash is:
    ///     keccak256(bridge_id, MESSAGE_ROOT_DOMAIN_SEPARATOR, nonce, message_root)
    /// @param _bridge_id Bridge ID.
    /// @param _nonce Nonce.
    /// @param _messageRoot Message root.
    function domainSeparateMessageRoot(
        bytes32 _bridge_id,
        uint256 _nonce,
        bytes32 _messageRoot
    ) private pure returns (bytes32) {
        bytes32 c = keccak256(abi.encode(_bridge_id, MESSAGE_ROOT_DOMAIN_SEPARATOR, _nonce, _messageRoot));

        return c;
    }

    /// @dev Checks that enough voting power signed over a digest.
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

        // Check that there was enough power
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
    /// @param _newValidatorSetHash The hash of the new validator set
    /// @param _newNonce The new nonce.
    /// @param _currentValidatorSet The current validator set.
    /// @param _sigs Signatures.
    function updateValidatorSet(
        uint256 _newNonce,
        uint256 _newPowerThreshold,
        bytes32 _newValidatorSetHash,
        Validator[] calldata _currentValidatorSet,
        Signature[] calldata _sigs
    ) external {
        // CHECKS

        uint256 currentNonce = s_lastValidatorSetNonce;
        uint256 currentPowerThreshold = s_powerThreshold;

        // Check that the valset nonce is greater than the old one.
        if (_newNonce <= currentNonce) {
            revert InvalidValidatorSetNonce();
        }

        // Check that current validators and signatures are well-formed.
        if (_currentValidatorSet.length != _sigs.length) {
            revert MalformedCurrentValidatorSet();
        }

        // Check that the supplied current validator set matches the saved checkpoint.
        bytes32 currentValidatorSetHash = computeValidatorSetHash(_currentValidatorSet);
        if (
            domainSeparateValidatorSetHash(BRIDGE_ID, currentNonce, currentPowerThreshold, currentValidatorSetHash) !=
            s_lastValidatorSetCheckpoint
        ) {
            revert SuppliedValidatorSetInvalid();
        }

        // Check that enough current validators have signed off on the new validator set.
        bytes32 newCheckpoint = domainSeparateValidatorSetHash(
            BRIDGE_ID,
            currentNonce,
            _newPowerThreshold,
            _newValidatorSetHash
        );
        checkValidatorSignatures(_currentValidatorSet, _sigs, newCheckpoint, currentPowerThreshold);

        // EFFECTS

        s_lastValidatorSetCheckpoint = newCheckpoint;
        s_powerThreshold = _newPowerThreshold;
        s_lastValidatorSetNonce = _newNonce;

        // LOGS

        emit ValidatorSetUpdatedEvent(_newNonce, _newPowerThreshold, _newValidatorSetHash);
    }

    /// @notice Relays a root of Celestia -> Ethereum messages. Anyone can
    // call this function, but they must supply valid signatures of the current
    // validator set over the message root.
    /// @param _nonce The message root nonce.
    /// @param _messageRoot The Merkle root of messages.
    /// @param _currentValidatorSet The current validator set.
    /// @param _sigs Signatures.
    function submitMessageRoot(
        uint256 _nonce,
        bytes32 _messageRoot,
        Validator[] calldata _currentValidatorSet,
        Signature[] calldata _sigs
    ) external {
        // CHECKS

        uint256 currentPowerThreshold = s_powerThreshold;

        // Check that the message root nonce is higher than the last nonce.
        if (_nonce <= s_lastMessageRootNonce) {
            revert InvalidMessageRootNonce();
        }

        // Check that current validators and signatures are well-formed.
        if (_currentValidatorSet.length != _sigs.length) {
            revert MalformedCurrentValidatorSet();
        }

        // Check that the supplied current validator set matches the saved checkpoint.
        bytes32 currentValidatorSetHash = computeValidatorSetHash(_currentValidatorSet);
        if (
            domainSeparateValidatorSetHash(
                BRIDGE_ID,
                s_lastValidatorSetNonce,
                currentPowerThreshold,
                currentValidatorSetHash
            ) != s_lastValidatorSetCheckpoint
        ) {
            revert SuppliedValidatorSetInvalid();
        }

        // Check that enough current validators have signed off on the message root and nonce.
        bytes32 c = domainSeparateMessageRoot(BRIDGE_ID, _nonce, _messageRoot);
        checkValidatorSignatures(_currentValidatorSet, _sigs, c, currentPowerThreshold);

        // EFFECTS

        s_lastMessageRootNonce = _nonce;
        s_messageRoots[_nonce] = _messageRoot;

        // LOGS

        emit MessageRootEvent(_nonce, _messageRoot);
    }
}
