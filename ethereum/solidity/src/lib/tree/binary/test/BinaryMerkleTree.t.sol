// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "ds-test/test.sol";

import "../BinaryMerkleProof.sol";
import "../BinaryMerkleTree.sol";

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
}
