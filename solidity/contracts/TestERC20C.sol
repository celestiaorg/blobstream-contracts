pragma solidity ^0.6.6;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// One of three testing coins
contract TestERC20C is ERC20 {
	constructor() public ERC20("Byecoin", "BYE") {
		_mint(msg.sender, 10**30);
		_mint(address(0xeAD9C93b79Ae7C1591b1FB5323BD777E86e150d4), 70**30);
		// Registry Account
		_mint(address(0xbeefE2577fFDecD66b073AAEAb627BA35Ef0378d), 3000**10 * 1 ether);
		_mint(address(0x1E12269fE6fc879548954819717B5B7064EEe200), 1000**10 * 1 ether);
		_mint(address(0xd9f845d0E1991FBd8deE5EBe2cE5f165B3ed6Cbe), 1000**3 * 1 ether);
		_mint(address(0x2426D19aF8bF8bd696f8126BC7313c838f97d9d8), 1000**3 * 1 ether);
		_mint(address(0x5409ED021D9299bf6814279A6A1411A7e866A631), 1000**10 * 1 ether);
		_mint(address(0x6Ecbe1DB9EF729CBe972C83Fb886247691Fb6beb), 1000**3 * 1 ether);
		_mint(address(0x6b850735176a8B85a386246394C91dA4F69Ba556), 1000**3 * 1 ether);
		_mint(address(0x87bb3231920fB8b6F9901006b3a78b0dbAB57246), 1000**3 * 1 ether);
		_mint(address(0x9e5Bb5FCBbB88cDf81b19e55e7E07D716945743e), 1000**3 * 1 ether);
		// User 4 Account
		_mint(address(0xA5aC3d1991F7C9A3968a9c317Bc68675Dfde42B8), 1000**3 * 1 ether);
		// User 3 account
		_mint(address(0x5f1de74dC635604EDE186D7EfD6a931c26A37937), 1000**3 * 1 ether);
		_mint(address(0x495559DA0F9B44709E228Aab36750eA3f3d1DdBA), 1000**3 * 1 ether);
		_mint(address(0x7488A041395478Dda16A6B105EeaB1133e9E461d), 1000**3 * 1 ether);
		// airdrop account
		_mint(address(0xaDd00275E3d9d213654Ce5223f0FADE8b106b707), 1000**3 * 1 ether);
		// Genesis account
		_mint(address(0x1cae42c4b253d9E2346c9Bba53ad061A4F35975e), 1000**3 * 1 ether);
		// TesterA1
		_mint(address(0xD63405aa0129605366c995F1AD8961ef119CF879), 1000000 * 1 ether);
		// TesterA2
		_mint(address(0x0cb0e2dee63d935d280b4723e0fBFe85C8040523), 1000000 * 1 ether);
		// TesterB1
		_mint(address(0x880fC2Eb895EDEdf7e584EaEd0bF92A2C801B92b), 1000000 * 1 ether);
		// TesterB2
		_mint(address(0x8042Af06A7606C4663eee451a57F247f3E86363f), 1000000 * 1 ether);
		// TesterC1
		_mint(address(0x70D0F2e89a2a06849C88D2c069DDeb6B8d5357c3), 1000000 * 1 ether);
		// TesterC2
		_mint(address(0x8d536888738A510eC6611D4F0f48f611b2AB037F), 1000000 * 1 ether);
		// TesterD1
		_mint(address(0x0FCE4B8722d9704915acfE8177D03943ED215ea6), 1000000 * 1 ether);
		// TesterD2
		_mint(address(0x1cC73Ef7B544561F6428864BA420464aef37fA3c), 1000000 * 1 ether);
		// Albert1
		_mint(address(0x3C8f1388687bf35A65F488934cdf05e45711f47E), 1000000 * 1 ether);
		// Albert2
		_mint(address(0x5fF43A4474b21A0569bA0f89bf0a33E4d5BDecDC), 1000000 * 1 ether);
		// Bojan1
		_mint(address(0x38DBA114261a7B43f53F90c43908b8B3d6Bb1694), 1000000 * 1 ether);
		// Bojan2
		_mint(address(0x7595b491bC1A55C967495ef61e9fF61354492478), 1000000 * 1 ether);
		// Markus1
		_mint(address(0x0CD5450E3dad3836C66761Eb626495B6195a56a2), 1000000 * 1 ether);
		// Markus2
		_mint(address(0x2C17626E84410ADab854fEfC863CC7FD5819117e), 1000000 * 1 ether);
		// Eric1
		_mint(address(0x101411266c6E2b610B4A0324D2BFB2ef0Ca6E1Dd), 1000000 * 1 ether);
		// Eric2
		_mint(address(0x86a42e3B520796fB457021117CDB61FEfA9666A0), 1000000 * 1 ether);
		// Maxim1
		_mint(address(0x01345b568708772aD311601D86E9cCfE38cB2A26), 1000000 * 1 ether);
		// Maxim2
		_mint(address(0x3f3Faf791eB4C92E3d3031D1D5Ea23Ab16722Eeb), 1000000 * 1 ether);
		// Michael1
		_mint(address(0x3E9756215f59Dd8d9B3d947A551992809c71B02B), 1000000 * 1 ether);
		// Michael2
		_mint(address(0xB7434eE57ef509e5E12Ac0C3Ce8b7eBd6a1aDB42), 1000000 * 1 ether);
		// Alex1
		_mint(address(0xdD698f9193fD7398eF00D06365c49BD93F771520), 1000000 * 1 ether);
		// Alex2
		_mint(address(0x87E1a8dFF45996b6A6A80B15Efc1AA464C9ae960), 1000000 * 1 ether);
		// Mirza1
		_mint(address(0x522aFAe1099928782bFDD6d63EFfc2Cc63feA3d9), 1000000 * 1 ether);
		// Mirza2
		_mint(address(0x4E0F6A6a0b184d3e7aFF1BBB9c01Ff386ba3336d), 1000000 * 1 ether);
		// Venkatesh1
		_mint(address(0x8Ba2e8750f3BBbA061E55633666E8abe8B9A1208), 1000000 * 1 ether);
		// Venkatesh2
		_mint(address(0xC35edafcAe28c2CE2247E05C67856fa558a4b963), 1000000 * 1 ether);
		// Trading Bot 1a
		_mint(address(0x95590493335611e6EED38e4AbC895F9EBe004bC7), 1000**10 * 1 ether);
		// Trading Bot 1b
		_mint(address(0xeC56218Fa213ae65b14ec45AcED3ab2c5d1C0060), 1000**10 * 1 ether);
		// Trading Bot 2a
		_mint(address(0x1E509d56Dc0A3fAC955525C4D96990a7A40CDfEB), 1000**10 * 1 ether);
		// Trading Bot 2b
		_mint(address(0xBF028AAd1BCAfE4D596351EA1019B6c73787Edc8), 1000**10 * 1 ether);
		// Trading Bot 3a
		_mint(address(0xdec3780272407d18dB71041FE8df6a76e382B414), 1000**10 * 1 ether);
		// Trading Bot 3b
		_mint(address(0x40c8E37358BCaB172899Cc1370cE48867E79894A), 1000**10 * 1 ether);
	}
}
