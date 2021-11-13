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

    ////////////////
    // Immutables //
    ////////////////

    bytes32 public immutable BRIDGE_ID;

    /////////////
    // Storage //
    /////////////

    bytes32 public s_lastValidatorSetCheckpoint;
    uint256 public s_powerThreshold;
    mapping(address => uint256) public s_lastBatchNonces;
    mapping(bytes32 => uint256) public s_invalidationMapping;
    uint256 public s_lastValidatorSetNonce;
    uint256 public s_lastMessageRootNonce;

    ////////////
    // Events //
    ////////////

    /// @notice Emitted when a new root of messages is relayed.
    /// @param nonce Nonce.
    /// @param messagesRoot Merkle root of relayed messages.
    event MessageRootEvent(uint256 indexed nonce, bytes32 messagesRoot);

    /// @notice Emitted when the validator set is updated.
    /// @param nonce Nonce.
    /// @param validatorSetRoot Merkle root of new validator set.
    event ValidatorSetUpdatedEvent(uint256 indexed nonce, bytes32 validatorSetRoot);

    ////////////
    // Errors //
    ////////////

    error MalformedCurrentValidatorSet();

    error InvalidSignature();

    error InsufficientVotingPower();

    ///////////////
    // Functions //
    ///////////////

    /// @param _bridge_id Identifier of the bridge, used in signatures.
    /// @param _powerThreshold Initial voting power that is needed to approve operations.
    /// @param _validatorSetRoot Initial validator set root.
    constructor(
        bytes32 _bridge_id,
        uint256 _powerThreshold,
        bytes32 _validatorSetRoot
    ) {
        __Ownable_init_unchained();

        BRIDGE_ID = _bridge_id;

        // CHECKS

        uint256 nonce = 0;
        bytes32 newCheckpoint = makeCheckpoint(nonce, _validatorSetRoot);

        // EFFECTS

        s_lastValidatorSetCheckpoint = newCheckpoint;
        s_powerThreshold = _powerThreshold;

        // LOGS

        emit ValidatorSetUpdatedEvent(nonce, _validatorSetRoot);
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

    /// @dev Make a new checkpoint from the supplied validator set.
    /// @dev A checkpoint is a hash of all relevant information about the valset.
    /// @dev The format of the checkpoint is:
    /// @dev     keccak256(bridge_id, bytes32("checkpoint"), nonce, validator_set_root)
    /// @dev The leaves in the validator set tree should be monotonically decreasing by power.
    /// @param _nonce Nonce.
    /// @param _validatorSetRoot Validator set root.
    function makeCheckpoint(uint256 _nonce, bytes32 _validatorSetRoot) private pure returns (bytes32) {
        // bytes32 encoding of the string "checkpoint"
        bytes32 methodName = 0x636865636b706f696e7400000000000000000000000000000000000000000000;

        bytes32 checkpoint = keccak256(abi.encode(BRIDGE_ID, methodName, _nonce, _validatorSetRoot));

        return checkpoint;
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
            // If v is set to 0, this signifies that it was not possible to get a signature from this validator and we skip evaluation.
            // (In a valid signature, it is either 27 or 28).
            if (_sigs[i].v != 0) {
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
        }

        // Check that there was enough power
        if (cumulativePower < _powerThreshold) {
            revert InsufficientVotingPower();
        }
    }

    // This updates the valset by checking that the validators in the current valset have signed off on the
    // new valset. The signatures supplied are the signatures of the current valset over the checkpoint hash
    // generated from the new valset.
    // Anyone can call this function, but they must supply valid signatures of state_powerThreshold of the current valset over
    // the new valset.
    function updateValset(
        // The new version of the validator set
        bytes32 _newValidatorSetRoot,
        uint256 _newNonce,
        // The current validators that approve the change
        Validator[] calldata _currentValidatorSet,
        uint256 _currentNonce,
        // These are arrays of the parts of the current validator's signatures
        Signature[] calldata _sigs
    ) external {
        // CHECKS

        // Check that the valset nonce is greater than the old one
        require(
            _newValset.valsetNonce > _currentValset.valsetNonce,
            "New valset nonce must be greater than the current nonce"
        );

        // Check that new validators and powers set is well-formed
        require(_newValset.validators.length == _newValset.powers.length, "Malformed new validator set");

        // Check that current validators, powers, and signatures (v,r,s) set is well-formed
        require(
            _currentValset.validators.length == _currentValset.powers.length &&
                _currentValset.validators.length == _v.length &&
                _currentValset.validators.length == _r.length &&
                _currentValset.validators.length == _s.length,
            "Malformed current validator set"
        );

        // Check that the supplied current validator set matches the saved checkpoint
        require(
            makeCheckpoint(_currentValset, state_peggyId) == state_lastValsetCheckpoint,
            "Supplied current validators and powers do not match checkpoint."
        );

        // Check that enough current validators have signed off on the new validator set
        bytes32 newCheckpoint = makeCheckpoint(_newValset, state_peggyId);
        checkValidatorSignatures(
            _currentValset.validators,
            _currentValset.powers,
            _v,
            _r,
            _s,
            newCheckpoint,
            state_powerThreshold
        );

        // EFFECTS

        // Stored to be used next time to validate that the valset
        // supplied by the caller is correct.
        state_lastValsetCheckpoint = newCheckpoint;

        // Store new nonce
        state_lastValsetNonce = _newValset.valsetNonce;

        // Send submission reward to msg.sender if reward token is a valid value
        if (_newValset.rewardToken != address(0) && _newValset.rewardAmount != 0) {
            IERC20(_newValset.rewardToken).safeTransfer(msg.sender, _newValset.rewardAmount);
        }

        // LOGS

        state_lastEventNonce = state_lastEventNonce + 1;
        emit ValsetUpdatedEvent(
            _newValset.valsetNonce,
            state_lastEventNonce,
            _newValset.rewardAmount,
            _newValset.rewardToken,
            _newValset.validators,
            _newValset.powers
        );
    }

    // submitBatch processes a batch of Cosmos -> Ethereum transactions by sending the tokens in the transactions
    // to the destination addresses. It is approved by the current Cosmos validator set.
    // Anyone can call this function, but they must supply valid signatures of state_powerThreshold of the current valset over
    // the batch.
    function submitBatch(
        // The validators that approve the batch
        ValsetArgs memory _currentValset,
        // These are arrays of the parts of the validators signatures
        uint8[] memory _v,
        bytes32[] memory _r,
        bytes32[] memory _s,
        // The batch of transactions
        uint256[] memory _amounts,
        address[] memory _destinations,
        uint256[] memory _fees,
        uint256 _batchNonce,
        address _tokenContract,
        // a block height beyond which this batch is not valid
        // used to provide a fee-free timeout
        uint256 _batchTimeout
    ) external {
        // CHECKS scoped to reduce stack depth
        {
            // Check that the batch nonce is higher than the last nonce for this token
            require(
                state_lastBatchNonces[_tokenContract] < _batchNonce,
                "New batch nonce must be greater than the current nonce"
            );

            // Check that the block height is less than the timeout height
            require(block.number < _batchTimeout, "Batch timeout must be greater than the current block height");

            // Check that current validators, powers, and signatures (v,r,s) set is well-formed
            require(
                _currentValset.validators.length == _currentValset.powers.length &&
                    _currentValset.validators.length == _v.length &&
                    _currentValset.validators.length == _r.length &&
                    _currentValset.validators.length == _s.length,
                "Malformed current validator set"
            );

            // Check that the supplied current validator set matches the saved checkpoint
            require(
                makeCheckpoint(_currentValset, state_peggyId) == state_lastValsetCheckpoint,
                "Supplied current validators and powers do not match checkpoint."
            );

            // Check that the transaction batch is well-formed
            require(
                _amounts.length == _destinations.length && _amounts.length == _fees.length,
                "Malformed batch of transactions"
            );

            // Check that enough current validators have signed off on the transaction batch and valset
            checkValidatorSignatures(
                _currentValset.validators,
                _currentValset.powers,
                _v,
                _r,
                _s,
                // Get hash of the transaction batch and checkpoint
                keccak256(
                    abi.encode(
                        state_peggyId,
                        // bytes32 encoding of "transactionBatch"
                        0x7472616e73616374696f6e426174636800000000000000000000000000000000,
                        _amounts,
                        _destinations,
                        _fees,
                        _batchNonce,
                        _tokenContract,
                        _batchTimeout
                    )
                ),
                state_powerThreshold
            );

            // EFFECTS

            // Store batch nonce
            state_lastBatchNonces[_tokenContract] = _batchNonce;

            {
                // Send transaction amounts to destinations
                uint256 totalFee;
                for (uint256 i = 0; i < _amounts.length; i++) {
                    IERC20(_tokenContract).safeTransfer(_destinations[i], _amounts[i]);
                    totalFee = totalFee + _fees[i];
                }

                if (totalFee > 0) {
                    // Send transaction fees to msg.sender
                    IERC20(_tokenContract).safeTransfer(msg.sender, totalFee);
                }
            }
        }

        // LOGS scoped to reduce stack depth
        {
            state_lastEventNonce = state_lastEventNonce + 1;
            emit TransactionBatchExecutedEvent(_batchNonce, _tokenContract, state_lastEventNonce);
        }
    }
}
