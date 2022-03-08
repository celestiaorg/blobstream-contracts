// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.4;


contract HashingTest {
	bytes32 public lastCheckpoint;
	address[] public state_validators;
	uint256[] public state_powers;
	uint256 public state_nonce;

	// CheckpointHash copies how makeCheckpoint in Peggy.sol works
	function CheckpointHash(
		address[] memory _validators,
		uint256[] memory _powers,
		uint256 _valsetNonce,
		uint256 _rewardAmount,
		address _rewardToken,
		bytes32 _peggyId
	) public {
		// bytes32 encoding of the string "checkpoint"
		bytes32 methodName = 0x636865636b706f696e7400000000000000000000000000000000000000000000;

		bytes32 checkpoint =
			keccak256(
				abi.encode(
				_peggyId,
				methodName,
				_valsetNonce,
				_validators,
				_powers,
				_rewardAmount,
				_rewardToken
				)
			);

		lastCheckpoint = checkpoint;
	}

	function JustSaveEverything(
		address[] memory _validators,
		uint256[] memory _powers,
		uint256 _valsetNonce
	) public {
		state_validators = _validators;
		state_powers = _powers;
		state_nonce = _valsetNonce;
	}

	function JustSaveEverythingAgain(
		address[] memory _validators,
		uint256[] memory _powers,
		uint256 _valsetNonce
	) public {
		state_validators = _validators;
		state_powers = _powers;
		state_nonce = _valsetNonce;
	}
}
