// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.19;

import "ds-test/test.sol";

import "../../Constants.sol";
import "../../Types.sol";
import "../NamespaceNode.sol";
import "../TreeHasher.sol";

contract TreeHasherTest is DSTest {
    function setUp() external {}

    function assertEqNamespaceNode(NamespaceNode memory first, NamespaceNode memory second) internal {
        assertEq(NamespaceID.unwrap(first.min), NamespaceID.unwrap(second.min));
        assertEq(NamespaceID.unwrap(first.max), NamespaceID.unwrap(second.max));
        assertEq(first.digest, second.digest);
    }

    function testLeafDigestEmpty() external {
        NamespaceID nid = NamespaceID.wrap(0x0000000000000000);
        NamespaceNode memory expected =
            NamespaceNode(nid, nid, 0x3e7077fd2f66d689e0cee6a7cf5b37bf2dca7c979af356d0a31cbc5c85605c7d);
        bytes memory data;
        NamespaceNode memory node = leafDigest(nid, data);
        assertEqNamespaceNode(node, expected);
    }

    function testLeafDigestSome() external {
        NamespaceID nid = NamespaceID.wrap(0xdeadbeefdeadbeef);
        NamespaceNode memory expected =
            NamespaceNode(nid, nid, 0x7c5146e5a2fe11d16375bfebe907722d77fad468411a2704f3863e41993186bb);
        bytes memory data = hex"69";
        NamespaceNode memory node = leafDigest(nid, data);
        assertEqNamespaceNode(node, expected);
    }

    function testNodeDigest() external {
        NamespaceID nidLeft = NamespaceID.wrap(0x0000000000000000);
        NamespaceID nidRight = NamespaceID.wrap(0xdeadbeefdeadbeef);
        NamespaceNode memory expected =
            NamespaceNode(nidLeft, nidRight, 0xc09cccb48cbc3a3ce4b19b9f25da11325d4fdf823ba56e990006fbc1eb8ddaf2);
        NamespaceNode memory left =
            NamespaceNode(nidLeft, nidLeft, 0xdb55da3fc3098e9c42311c6013304ff36b19ef73d12ea932054b5ad51df4f49d);
        NamespaceNode memory right =
            NamespaceNode(nidRight, nidRight, 0xc75cb66ae28d8ebc6eded002c28a8ba0d06d3a78c6b5cbf9b2ade051f0775ac4);
        NamespaceNode memory node = nodeDigest(left, right);
        assertEqNamespaceNode(node, expected);
    }

    function testNodeParity() external {
        NamespaceID nidMin = NamespaceID.wrap(0x0000000000000000);
        NamespaceID nidMax = NamespaceID.wrap(0xdeadbeefdeadbeef);
        NamespaceNode memory expected =
            NamespaceNode(nidMin, nidMax, 0xb16c8e95fa3655fa06d2ccf09f8351443c5a838a1f1b8d5cf2cb1ec00adf2662);
        NamespaceNode memory left =
            NamespaceNode(nidMin, nidMax, 0xdb55da3fc3098e9c42311c6013304ff36b19ef73d12ea932054b5ad51df4f49d);
        NamespaceNode memory right = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0xc75cb66ae28d8ebc6eded002c28a8ba0d06d3a78c6b5cbf9b2ade051f0775ac4
        );
        NamespaceNode memory node = nodeDigest(left, right);
        assertEqNamespaceNode(node, expected);
    }
}
