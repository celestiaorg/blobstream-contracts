// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "ds-test/test.sol";

import "../NamespaceNode.sol";
import "../TreeHasher.sol";

contract TreeHasherTest is DSTest {
    function setUp() external {}

    function assertEqNamespaceNode(NamespaceNode memory first, NamespaceNode memory second) internal {
        assertEq(first.min, second.min);
        assertEq(first.max, second.max);
        assertEq(first.digest, second.digest);
    }

    function testLeafDigestEmpty() external {
        bytes8 nid = 0x0000000000000000;
        NamespaceNode memory expected = NamespaceNode(
            nid,
            nid,
            0x0a88111852095cae045340ea1f0b279944b2a756a213d9b50107d7489771e159
        );
        bytes memory data;
        NamespaceNode memory node = leafDigest(nid, data);
        assertEqNamespaceNode(node, expected);
    }

    function testLeafDigestSome() external {
        bytes8 nid = 0xdeadbeefdeadbeef;
        NamespaceNode memory expected = NamespaceNode(
            nid,
            nid,
            0x2f8203f6673f9dffe69ca0b64e530656eb7445b062f69c32e2163931e637a659
        );
        bytes memory data = hex"69";
        NamespaceNode memory node = leafDigest(nid, data);
        assertEqNamespaceNode(node, expected);
    }

    function testNodeDigest() external {
        bytes8 nidLeft = 0x0000000000000000;
        bytes8 nidRight = 0xdeadbeefdeadbeef;
        NamespaceNode memory expected = NamespaceNode(
            nidLeft,
            nidRight,
            0xc09cccb48cbc3a3ce4b19b9f25da11325d4fdf823ba56e990006fbc1eb8ddaf2
        );
        NamespaceNode memory left = NamespaceNode(
            nidLeft,
            nidLeft,
            0xdb55da3fc3098e9c42311c6013304ff36b19ef73d12ea932054b5ad51df4f49d
        );
        NamespaceNode memory right = NamespaceNode(
            nidRight,
            nidRight,
            0xc75cb66ae28d8ebc6eded002c28a8ba0d06d3a78c6b5cbf9b2ade051f0775ac4
        );
        NamespaceNode memory node = nodeDigest(left, right);
        assertEqNamespaceNode(node, expected);
    }
}
