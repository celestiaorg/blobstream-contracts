// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "./@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract CosmosERC20 is ERC20 {
	uint256 MAX_UINT = 2**256 - 1;

	constructor(
		address _peggyAddress,
		string memory _name,
		string memory _symbol,
		uint8 _decimals
	) ERC20(_name, _symbol) {
		_setupDecimals(_decimals);
		_mint(_peggyAddress, MAX_UINT);
	}
}
