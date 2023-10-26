// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "ds-test/test.sol";

import "../TreeHasher.sol";

contract TreeHasherTest is DSTest {
    function setUp() external {}

    function testLeafDigestEmpty() external {
        bytes32 expected = 0x6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d;
        bytes memory data;
        bytes32 digest = leafDigest(data);
        assertEq(digest, expected);
    }

    function testLeafDigestSome() external {
        bytes32 expected = 0x48c90c8ae24688d6bef5d48a30c2cc8b6754335a8db21793cc0a8e3bed321729;
        bytes memory data = hex"deadbeef";
        bytes32 digest = leafDigest(data);
        assertEq(digest, expected);
    }

    function testNodeDigestEmptyChildren() external {
        bytes32 expected = 0xfe43d66afa4a9a5c4f9c9da89f4ffb52635c8f342e7ffb731d68e36c5982072a;
        bytes32 left = 0x6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d;
        bytes32 right = 0x6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d;
        bytes32 digest = nodeDigest(left, right);
        assertEq(digest, expected);
    }

    function testNodeDigestSomeChildren() external {
        bytes32 expected = 0x62343bba7c4d6259f0d4863cdf476f1c0ac1b9fbe9244723a9b8b5c8aae72c38;
        bytes32 left = 0xdb55da3fc3098e9c42311c6013304ff36b19ef73d12ea932054b5ad51df4f49d;
        bytes32 right = 0xc75cb66ae28d8ebc6eded002c28a8ba0d06d3a78c6b5cbf9b2ade051f0775ac4;
        bytes32 digest = nodeDigest(left, right);
        assertEq(digest, expected);
    }
}
