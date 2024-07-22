// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "ds-test/test.sol";
import "forge-std/Vm.sol";

import "../BinaryMerkleProof.sol";
import "../BinaryMerkleTree.sol";
import "../BinaryMerkleMultiproof.sol";

/**
 * TEST VECTORS
 *
 * 0x01
 * 0x02
 * 0x03
 * 0x04
 * 0x05
 * 0x06
 * 0x07
 * 0x08
 *
 * 0xb413f47d13ee2fe6c845b2ee141af81de858df4ec549a58b7970bb96645bc8d2
 * 0xfcf0a6c700dd13e274b6fba8deea8dd9b26e4eedde3495717cac8408c9c5177f
 * 0x583c7dfb7b3055d99465544032a571e10a134b1b6f769422bbb71fd7fa167a5d
 * 0x4f35212d12f9ad2036492c95f1fe79baf4ec7bd9bef3dffa7579f2293ff546a4
 * 0x9f1afa4dc124cba73134e82ff50f17c8f7164257c79fed9a13f5943a6acb8e3d
 * 0x40d88127d4d31a3891f41598eeed41174e5bc89b1eb9bbd66a8cbfc09956a3fd
 * 0x2ecd8a6b7d2845546659ad4cf443533cf921b19dc81fa83934e83821b4dfdcb7
 * 0xb4c43b50bf245bd727623e3c775a8fcfb8d823d00b57dd65f7f79dd33f126315
 *
 * 0x6bcf0e2e93e0a18e22789aee965e6553f4fbe93f0acfc4a705d691c8311c4965
 * 0x78850a5ab36238b076dd99fd258c70d523168704247988a94caa8c9ccd056b8d
 * 0x90eeb2c4a04ec33ee4dd2677593331910e4203db4fcc120a6cdb95b13cfe83f0
 * 0x28c01722dd8dd05b63bcdeb6878bc2c083118cc2b170646d6b842d0bdbdc9d29
 *
 * 0xfa02d31a63cc11cc624881e52af14af7a1c6ab745efa71021cb24086b9b1793f
 * 0x4301a067262bbb18b4919742326f6f6d706099f9c0e8b0f2db7b88f204b2cf09
 *
 * 0xc1ad6548cb4c7663110df219ec8b36ca63b01158956f4be31a38a88d0c7f7071
 *
 */
contract BinaryMerkleProofTest is DSTest {
    Vm private constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    function setUp() external {}

    function testVerifyNone() external {
        bytes32 root = sha256("");
        bytes32[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 0;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data;
        (bool isValid,) = BinaryMerkleTree.verify(root, proof, data);
        assertTrue(!isValid);
    }

    function testVerifyOneLeafEmpty() external {
        bytes32 root = 0x6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d;
        bytes32[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data;
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(isValid);
    }

    function testVerifyOneLeafSome() external {
        bytes32 root = 0x48c90c8ae24688d6bef5d48a30c2cc8b6754335a8db21793cc0a8e3bed321729;
        bytes32[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"deadbeef";
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(isValid);
    }

    function testVerifyOneLeaf01() external {
        bytes32 root = 0xb413f47d13ee2fe6c845b2ee141af81de858df4ec549a58b7970bb96645bc8d2;
        bytes32[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(isValid);
    }

    function testVerifyLeafOneOfEight() external {
        bytes32 root = 0xc1ad6548cb4c7663110df219ec8b36ca63b01158956f4be31a38a88d0c7f7071;
        bytes32[] memory sideNodes = new bytes32[](3);
        sideNodes[0] = 0xfcf0a6c700dd13e274b6fba8deea8dd9b26e4eedde3495717cac8408c9c5177f;
        sideNodes[1] = 0x78850a5ab36238b076dd99fd258c70d523168704247988a94caa8c9ccd056b8d;
        sideNodes[2] = 0x4301a067262bbb18b4919742326f6f6d706099f9c0e8b0f2db7b88f204b2cf09;

        uint256 key = 0;
        uint256 numLeaves = 8;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(isValid);
    }

    function testVerifyLeafTwoOfEight() external {
        bytes32 root = 0xc1ad6548cb4c7663110df219ec8b36ca63b01158956f4be31a38a88d0c7f7071;
        bytes32[] memory sideNodes = new bytes32[](3);
        sideNodes[0] = 0xb413f47d13ee2fe6c845b2ee141af81de858df4ec549a58b7970bb96645bc8d2;
        sideNodes[1] = 0x78850a5ab36238b076dd99fd258c70d523168704247988a94caa8c9ccd056b8d;
        sideNodes[2] = 0x4301a067262bbb18b4919742326f6f6d706099f9c0e8b0f2db7b88f204b2cf09;

        uint256 key = 1;
        uint256 numLeaves = 8;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"02";
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(isValid);
    }

    function testVerifyLeafThreeOfEight() external {
        bytes32 root = 0xc1ad6548cb4c7663110df219ec8b36ca63b01158956f4be31a38a88d0c7f7071;
        bytes32[] memory sideNodes = new bytes32[](3);
        sideNodes[0] = 0x4f35212d12f9ad2036492c95f1fe79baf4ec7bd9bef3dffa7579f2293ff546a4;
        sideNodes[1] = 0x6bcf0e2e93e0a18e22789aee965e6553f4fbe93f0acfc4a705d691c8311c4965;
        sideNodes[2] = 0x4301a067262bbb18b4919742326f6f6d706099f9c0e8b0f2db7b88f204b2cf09;

        uint256 key = 2;
        uint256 numLeaves = 8;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"03";
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(isValid);
    }

    function testVerifyLeafSevenOfEight() external {
        bytes32 root = 0xc1ad6548cb4c7663110df219ec8b36ca63b01158956f4be31a38a88d0c7f7071;
        bytes32[] memory sideNodes = new bytes32[](3);
        sideNodes[0] = 0xb4c43b50bf245bd727623e3c775a8fcfb8d823d00b57dd65f7f79dd33f126315;
        sideNodes[1] = 0x90eeb2c4a04ec33ee4dd2677593331910e4203db4fcc120a6cdb95b13cfe83f0;
        sideNodes[2] = 0xfa02d31a63cc11cc624881e52af14af7a1c6ab745efa71021cb24086b9b1793f;

        uint256 key = 6;
        uint256 numLeaves = 8;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"07";
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(isValid);
    }

    function testVerifyLeafEightOfEight() external {
        bytes32 root = 0xc1ad6548cb4c7663110df219ec8b36ca63b01158956f4be31a38a88d0c7f7071;
        bytes32[] memory sideNodes = new bytes32[](3);
        sideNodes[0] = 0x2ecd8a6b7d2845546659ad4cf443533cf921b19dc81fa83934e83821b4dfdcb7;
        sideNodes[1] = 0x90eeb2c4a04ec33ee4dd2677593331910e4203db4fcc120a6cdb95b13cfe83f0;
        sideNodes[2] = 0xfa02d31a63cc11cc624881e52af14af7a1c6ab745efa71021cb24086b9b1793f;

        uint256 key = 7;
        uint256 numLeaves = 8;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"08";
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(isValid);
    }

    // Test vectors:
    // 0x00
    // 0x01
    // 0x02
    // 0x03
    // 0x04
    function testVerifyProofOfFiveLeaves() external {
        bytes32 root = 0xb855b42d6c30f5b087e05266783fbd6e394f7b926013ccaa67700a8b0c5a596f;
        bytes32[] memory sideNodes = new bytes32[](3);
        sideNodes[0] = 0x96a296d224f285c67bee93c30f8a309157f0daa35dc5b87e410b78630a09cfc7;
        sideNodes[1] = 0x52c56b473e5246933e7852989cd9feba3b38f078742b93afff1e65ed46797825;
        sideNodes[2] = 0x4f35212d12f9ad2036492c95f1fe79baf4ec7bd9bef3dffa7579f2293ff546a4;

        uint256 key = 1;
        uint256 numLeaves = 5;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = bytes(hex"01");
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(isValid);
    }

    function testVerifyInvalidProofRoot() external {
        // correct root: 0xb855b42d6c30f5b087e05266783fbd6e394f7b926013ccaa67700a8b0c5a596f;
        bytes32 root = 0xc855b42d6c30f5b087e05266783fbd6e394f7b926013ccaa67700a8b0c5a596f;
        bytes32[] memory sideNodes = new bytes32[](3);
        sideNodes[0] = 0x96a296d224f285c67bee93c30f8a309157f0daa35dc5b87e410b78630a09cfc7;
        sideNodes[1] = 0x52c56b473e5246933e7852989cd9feba3b38f078742b93afff1e65ed46797825;
        sideNodes[2] = 0x4f35212d12f9ad2036492c95f1fe79baf4ec7bd9bef3dffa7579f2293ff546a4;

        uint256 key = 1;
        uint256 numLeaves = 5;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = bytes(hex"01");
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(!isValid);
    }

    function testVerifyInvalidProofKey() external {
        bytes32 root = 0xb855b42d6c30f5b087e05266783fbd6e394f7b926013ccaa67700a8b0c5a596f;
        bytes32[] memory sideNodes = new bytes32[](3);
        sideNodes[0] = 0x96a296d224f285c67bee93c30f8a309157f0daa35dc5b87e410b78630a09cfc7;
        sideNodes[1] = 0x52c56b473e5246933e7852989cd9feba3b38f078742b93afff1e65ed46797825;
        sideNodes[2] = 0x4f35212d12f9ad2036492c95f1fe79baf4ec7bd9bef3dffa7579f2293ff546a4;

        // correct key: 1
        uint256 key = 2;
        uint256 numLeaves = 5;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = bytes(hex"01");
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(!isValid);
    }

    function testVerifyInvalidProofNumberOfLeaves() external {
        bytes32 root = 0xb855b42d6c30f5b087e05266783fbd6e394f7b926013ccaa67700a8b0c5a596f;
        bytes32[] memory sideNodes = new bytes32[](3);
        sideNodes[0] = 0x96a296d224f285c67bee93c30f8a309157f0daa35dc5b87e410b78630a09cfc7;
        sideNodes[1] = 0x52c56b473e5246933e7852989cd9feba3b38f078742b93afff1e65ed46797825;
        sideNodes[2] = 0x4f35212d12f9ad2036492c95f1fe79baf4ec7bd9bef3dffa7579f2293ff546a4;

        uint256 key = 1;
        // correct numLeaves: 5
        uint256 numLeaves = 200;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = bytes(hex"01");
        (bool isValid,) = BinaryMerkleTree.verify(root, proof, data);
        assertTrue(!isValid);
    }

    function testVerifyInvalidProofSideNodes() external {
        bytes32 root = 0xb855b42d6c30f5b087e05266783fbd6e394f7b926013ccaa67700a8b0c5a596f;
        bytes32[] memory sideNodes = new bytes32[](3);
        sideNodes[0] = 0x96a296d224f285c67bee93c30f8a309157f0daa35dc5b87e410b78630a09cfc7;
        sideNodes[1] = 0x52c56b473e5246933e7852989cd9feba3b38f078742b93afff1e65ed46797825;
        // correct side node: 0x4f35212d12f9ad2036492c95f1fe79baf4ec7bd9bef3dffa7579f2293ff546a4;
        sideNodes[2] = 0x5f35212d12f9ad2036492c95f1fe79baf4ec7bd9bef3dffa7579f2293ff546a4;

        uint256 key = 1;
        uint256 numLeaves = 5;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = bytes(hex"01");
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(!isValid);
    }

    function testVerifyInvalidProofData() external {
        bytes32 root = 0xb855b42d6c30f5b087e05266783fbd6e394f7b926013ccaa67700a8b0c5a596f;
        bytes32[] memory sideNodes = new bytes32[](3);
        sideNodes[0] = 0x96a296d224f285c67bee93c30f8a309157f0daa35dc5b87e410b78630a09cfc7;
        sideNodes[1] = 0x52c56b473e5246933e7852989cd9feba3b38f078742b93afff1e65ed46797825;
        sideNodes[2] = 0x4f35212d12f9ad2036492c95f1fe79baf4ec7bd9bef3dffa7579f2293ff546a4;

        uint256 key = 1;
        uint256 numLeaves = 5;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        // correct data: 01
        bytes memory data = bytes(hex"012345");
        (bool isValid, BinaryMerkleTree.ErrorCodes error) = BinaryMerkleTree.verify(root, proof, data);
        assertEq(uint256(BinaryMerkleTree.ErrorCodes.NoError), uint256(error));
        assertTrue(!isValid);
    }

    function testValidSlice() public {
        bytes32[] memory data = new bytes32[](4);
        data[0] = "a";
        data[1] = "b";
        data[2] = "c";
        data[3] = "d";

        bytes32[] memory result = BinaryMerkleTree.slice(data, 1, 3);

        assertEq(result[0], data[1]);
        assertEq(result[1], data[2]);
    }

    function testSameKeyAndLeavesNumber() external {
        bytes32 root = 0xb855b42d6c30f5b087e05266783fbd6e394f7b926013ccaa67700a8b0c5a596f;
        bytes32[] memory sideNodes = new bytes32[](0);
        uint256 key = 3;
        uint256 numLeaves = 3;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = bytes(hex"01");
        (bool isValid,) = BinaryMerkleTree.verify(root, proof, data);
        assertTrue(!isValid);
    }

    function testConsecutiveKeyAndNumberOfLeaves() external {
        bytes32 root = 0xb855b42d6c30f5b087e05266783fbd6e394f7b926013ccaa67700a8b0c5a596f;
        bytes32[] memory sideNodes = new bytes32[](0);
        uint256 key = 6;
        uint256 numLeaves = 7;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = bytes(hex"01");
        (bool isValid,) = BinaryMerkleTree.verify(root, proof, data);
        assertTrue(!isValid);
    }

    function testInvalidSliceBeginEnd() public {
        bytes32[] memory data = new bytes32[](4);
        data[0] = "a";
        data[1] = "b";
        data[2] = "c";
        data[3] = "d";

        vm.expectRevert("Invalid range: _begin is greater than _end");
        BinaryMerkleTree.slice(data, 2, 1);
    }

    function testOutOfBoundsSlice() public {
        bytes32[] memory data = new bytes32[](4);
        data[0] = "a";
        data[1] = "b";
        data[2] = "c";
        data[3] = "d";

        vm.expectRevert("Invalid range: _begin or _end are out of bounds");
        BinaryMerkleTree.slice(data, 2, 5);
    }

    // The hard-coded serialized proofs and data were generated in Rust, with this code
    // https://github.com/S1nus/hyperchain-da/blob/main/src/clients/celestia/evm_types.rs#L132
    function testMultiproof() public {
        bytes memory proofData =
            hex"00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000090000000000000000000000000000000000000000000000000000000000000006ce29bcde696f84e35c5626904542a549b080e92603243b34794242473940706917519bf954f5b30495af5c8cdb9983e6319104badc1ea811ed2c421018a3ad7821ea268d3540deab8f9b2024464618610c9a7083620badcf505bda647cc8e9f82bfc87d990d8344f6efd44fcb09b46b87f9a92230d41329452efee8656c6760a9ad9f3a95af971e89e2a80b255bb56d5aae15de69803b52aa5079b33374b16e16178fc62a2f2ce6bf21909c0a0edea9525486e0ece65bff23499342cca38dd62";
        BinaryMerkleMultiproof memory multiproof = abi.decode(proofData, (BinaryMerkleMultiproof));
        bytes32 dataroot = hex"ef8920d86519bd5f8ce3c802b84fc9b9512483e4d4a5c9608b44af4d6639f7d1";
        bytes memory leafData =
            hex"00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000009000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000002a0000000000000000000000000000000000000000000000000000000000000032000000000000000000000000000000000000000000000000000000000000003a0000000000000000000000000000000000000000000000000000000000000042000000000000000000000000000000000000000000000000000000000000004a00000000000000000000000000000000000000000000000000000000000000520000000000000000000000000000000000000000000000000000000000000005a00000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000102030405746e218305fe3dbbef65feceed939fe8dd93c88b06c95473fbe344fb864060f3000000000000000000000000000000000000000000000000000000000000000000000000005a0000000000000000000000000000000000000000000000000102030405000000000000000000000000000000000000000000000000010203040555cd7fb524ae792c9d4bc8946d07209728c533a3e14d4e7c0c95c0b150d0c284000000000000000000000000000000000000000000000000000000000000000000000000005a00000000000000000000000000000000000000000000000001020304050000000000000000000000000000000000000000000000000102030405505c1e7c897461a152e152f1ff3ecc358fefdf1f69448ab1165b6ca76836933b000000000000000000000000000000000000000000000000000000000000000000000000005a00000000000000000000000000000000000000000000000001020304050000000000000000000000000000000000000000000000000102030405100a0548893d8eab0322f34f45ac84785cdf50dfab5102a12d958e6031bacebe000000000000000000000000000000000000000000000000000000000000000000000000005a0000000000000000000000000000000000000000000000000102030405000000000000000000000000000000000000000000000000010203040566e5eb1da67430f204a3c5615591f71316695c7ec1f1f713cde7e936d4a43ec1000000000000000000000000000000000000000000000000000000000000000000000000005a00000000000000000000000000000000000000000000000001020304050000000000000000000000000000000000000000000000000102030405d2a5de6299e28c2fec359a2718599f5ac22c2948a71d26a438295e531b6f4cb5000000000000000000000000000000000000000000000000000000000000000000000000005a00000000000000000000000000000000000000000000000001020304050000000000000000000000000000000000000000000000000102030405688c5238e50c0a8a556bfabff31bef1fa9cdd812c9fd4dcee5c2a0836f687fbf000000000000000000000000000000000000000000000000000000000000000000000000005a00000000000000000000000000000000000000000000000001020304050000000000000000000000000000000000000000000000000102030405b55a5b1efc2a22cdbfa21d050bd67147ff2b936c68354eb1a83bcdf14eb57e38000000000000000000000000000000000000000000000000000000000000000000000000005a000000000000000000000000000000000000000000000000010203040500000000000000000000000000000000000000000067480c4a88c4d129947e11c33fa811daa791771e591dd933498d1212d46b8cde9c34c28831b0b532000000000000";
        bytes[] memory leaves = abi.decode(leafData, (bytes[]));
        assertTrue(BinaryMerkleTree.verifyMulti(dataroot, multiproof, leaves));
    }
}
