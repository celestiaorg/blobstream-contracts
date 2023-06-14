// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import "ds-test/test.sol";

import "../BinaryMerkleProof.sol";
import "../BinaryMerkleTree.sol";

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
    function setUp() external {}

    function testVerifyNone() external {
        bytes32 root = sha256("");
        bytes32[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 0;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data;
        bool isValid = BinaryMerkleTree.verify(root, proof, data);
        assertTrue(!isValid);
    }

    function testVerifyOneLeafEmpty() external {
        bytes32 root = 0x6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d;
        bytes32[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data;
        bool isValid = BinaryMerkleTree.verify(root, proof, data);
        assertTrue(isValid);
    }

    function testVerifyOneLeafSome() external {
        bytes32 root = 0x48c90c8ae24688d6bef5d48a30c2cc8b6754335a8db21793cc0a8e3bed321729;
        bytes32[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"deadbeef";
        bool isValid = BinaryMerkleTree.verify(root, proof, data);
        assertTrue(isValid);
    }

    function testVerifyOneLeaf01() external {
        bytes32 root = 0xb413f47d13ee2fe6c845b2ee141af81de858df4ec549a58b7970bb96645bc8d2;
        bytes32[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        BinaryMerkleProof memory proof = BinaryMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        bool isValid = BinaryMerkleTree.verify(root, proof, data);
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
        bool isValid = BinaryMerkleTree.verify(root, proof, data);
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
        bool isValid = BinaryMerkleTree.verify(root, proof, data);
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
        bool isValid = BinaryMerkleTree.verify(root, proof, data);
        assertTrue(isValid);
    }
}
