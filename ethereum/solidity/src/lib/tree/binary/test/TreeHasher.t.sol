// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

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

    function testNodeDigest() external {}
}
