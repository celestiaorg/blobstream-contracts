// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import "ds-test/test.sol";

import "../NamespaceNode.sol";
import "../NamespaceMerkleMultiproof.sol";
import "../NamespaceMerkleTree.sol";

/**
TEST VECTORS

Data blocks: namespace id, data
0x0000000000000010 0x01
0x0000000000000020 0x02
0x0000000000000030 0x03
0x0000000000000040 0x04
0xffffffffffffffff 0x05
0xffffffffffffffff 0x06
0xffffffffffffffff 0x07
0xffffffffffffffff 0x08

Leaf nodes: min namespace, max namespace, data
0x0000000000000010 0x0000000000000010 0xcb9b006518aa5b6e8f62dcda719f42a17033573e2cde97fe2748944f81638514
0x0000000000000020 0x0000000000000020 0xf4653e02dfeff8eddbcf1c7230dfea1dd45b7bcc2fb1ce6d04c33f2229e10f6b
0x0000000000000030 0x0000000000000030 0x1f7e7711dd732649f2599fa0a47330c48ad64e460c1fb1287ba531797702e5fd
0x0000000000000040 0x0000000000000040 0x32706b95e3c3e7b4dd285fd4f73ad33dfb2d37e2dd11b3e355749d218ec2e00d
0xffffffffffffffff 0xffffffffffffffff 0x41206f8a19e9497538158cee344eae117bb0a7ba396561c4e1c3b245fced4f7f
0xffffffffffffffff 0xffffffffffffffff 0x84dd5d21f95db8c01adb5c742191da892b01eaafe8dafc6b19a560331e5d5912
0xffffffffffffffff 0xffffffffffffffff 0x24ddc56b10cebbf760b3a744ad3a0e91093db34b4d22995f6de6dac918e38ae5
0xffffffffffffffff 0xffffffffffffffff 0xf5a80844a112828c28da280019cb6e97765f81e1e003cc78a198901494db2641

Inner nodes(depth = 2): min namespace, max namespace, data
0x0000000000000010 0x0000000000000020 0xe0a6f55a5c2d86e0057b89d79bf5c6c3fdc5a40061566c39e93077556e2e3482
0x0000000000000030 0x0000000000000040 0x3f8ded56b6a8d4e1e36832e8be93234e2e3a18c1a42edfb505ecc09f0039a10f
0xffffffffffffffff 0xffffffffffffffff 0x61d6762ff063c2008a412246bc6bb370885c4e1a935ca28ed8699dc5c68ff28a
0xffffffffffffffff 0xffffffffffffffff 0x9086b06cbc327959e3c34546aadc886300aff3e5c8f96a328267abf64ca5d25b

Inner nodes(depth = 1): min namespace, max namespace, data
0x0000000000000010 0x0000000000000040 0xed6a82bfecd113f693065e3b1f271f21150b6d793917402f6c05a01feb6b3eb8
0xffffffffffffffff 0xffffffffffffffff 0x27209d167edf7ea1463f462b850471ce31b124b0b3405c33f9c39e692c9170da

Root node: min namespace, max namespace, data
0x0000000000000010 0x0000000000000040 0x16c760661bc5ed683d27dc2f045a81a67e837928527e0de209a195b6db60f462
**/

contract NamespaceMerkleMultiproofTest is DSTest {
    function setUp() external {}

    function assertEqNamespaceNode(NamespaceNode memory first, NamespaceNode memory second) internal {
        assertEq(first.min, second.min);
        assertEq(first.max, second.max);
        assertEq(first.digest, second.digest);
    }

    /// @notice Verify inclusion of leaves 0 and 1.
    function testVerifyMulti01() external {
        bytes8 nid = 0x0000000000000000;
        NamespaceNode memory root = NamespaceNode(nid, nid, sha256(""));
        NamespaceNode[] memory sideNodes;
        uint256 keyBegin = 0;
        uint256 keyEnd = 0;
        uint256 numLeaves = 0;
        NamespaceMerkleMultiproof memory proof = NamespaceMerkleMultiproof(keyBegin, keyEnd, sideNodes, numLeaves);
        bytes[] memory data;
        bool isValid = NamespaceMerkleTree.verifyMulti(root, proof, nid, data);
        assertTrue(!isValid);
    }

    /// @notice Verify inclusion of leaves 0 and 1.
    function testVerifyMulti01() external {
        bytes8 nid = 0x0000000000000000;
        NamespaceNode memory root = NamespaceNode(
            0x0000000000000010,
            0x0000000000000040,
            0x16c760661bc5ed683d27dc2f045a81a67e837928527e0de209a195b6db60f462
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            0x0000000000000020,
            0x0000000000000020,
            0xf4653e02dfeff8eddbcf1c7230dfea1dd45b7bcc2fb1ce6d04c33f2229e10f6b
        );
        sideNodes[1] = NamespaceNode(
            0x0000000000000030,
            0x0000000000000040,
            0x3f8ded56b6a8d4e1e36832e8be93234e2e3a18c1a42edfb505ecc09f0039a10f
        );
        sideNodes[2] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x27209d167edf7ea1463f462b850471ce31b124b0b3405c33f9c39e692c9170da
        );

        uint256 beginKey = 0;
        uint256 endKey = 1;
        uint256 numLeaves = 8;
        NamespaceMerkleMultiproof memory proof = NamespaceMerkleMultiproof(beginKey, endKey, sideNodes, numLeaves);
        bytes memory data = hex"01";
        bool isValid = NamespaceMerkleTree.verifyMulti(root, proof, nid, data);
        assertTrue(isValid);
    }
}
