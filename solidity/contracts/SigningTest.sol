// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

contract SigningTest {
	function checkSignature(
		address _signer,
		bytes32 _theHash,
		uint8 _v,
		bytes32 _r,
		bytes32 _s
	) public view {
		bytes32 messageDigest = keccak256(abi.encode("\x19Ethereum Signed Message:\n32", _theHash));
		require(_signer == ecrecover(messageDigest, _v, _r, _s), "Signature does not match.");
	}
}
