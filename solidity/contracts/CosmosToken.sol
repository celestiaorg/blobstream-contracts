// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;

import "./@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract CosmosERC20 is ERC20 {
	uint256 MAX_UINT = 2**256 - 1;
	uint8 immutable private _decimals;

	constructor(
		address peggyAddress_,
		string memory name_,
		string memory symbol_,
		uint8 decimals_
	) ERC20(name_, symbol_) {
		_decimals = decimals_;
		_mint(peggyAddress_, MAX_UINT);
	}

	function decimals() public view virtual override returns (uint8) {
		return _decimals;
	}
}
