// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "ds-test/test.sol";

import "../../Constants.sol";
import "../../Types.sol";
import "../NamespaceNode.sol";
import "../TreeHasher.sol";

contract TreeHasherTest is DSTest {
    function setUp() external {}

    function assertEqNamespaceNode(NamespaceNode memory first, NamespaceNode memory second) internal {
        assertTrue(first.min.equalTo(second.min));
        assertTrue(first.max.equalTo(second.max));
        assertEq(first.digest, second.digest);
    }

    function testLeafDigestEmpty() external {
        Namespace memory nid = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000000);
        NamespaceNode memory expected =
            NamespaceNode(nid, nid, 0x0679246d6c4216de0daa08e5523fb2674db2b6599c3b72ff946b488a15290b62);
        bytes memory data;
        NamespaceNode memory node = leafDigest(nid, data);
        assertEqNamespaceNode(node, expected);
    }

    function testLeafDigestSome() external {
        Namespace memory nid = Namespace(0xde, 0xadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefde);
        NamespaceNode memory expected =
            NamespaceNode(nid, nid, 0x3624c7f7169cb5bbd0d010b851ebd0edca10b2a1b126f5fb1a6d5e0d98356e63);
        bytes memory data = hex"69";
        NamespaceNode memory node = leafDigest(nid, data);
        assertEqNamespaceNode(node, expected);
    }

    function testNodeDigest() external {
        Namespace memory nidLeft = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000000);
        Namespace memory nidRight = Namespace(0xde, 0xadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefde);
        NamespaceNode memory expected =
            NamespaceNode(nidLeft, nidRight, 0x95cad48bc181484c851004cf772abe767391e19549d3b8192b55b1d654a71bcd);
        NamespaceNode memory left =
            NamespaceNode(nidLeft, nidLeft, 0xdb55da3fc3098e9c42311c6013304ff36b19ef73d12ea932054b5ad51df4f49d);
        NamespaceNode memory right =
            NamespaceNode(nidRight, nidRight, 0xc75cb66ae28d8ebc6eded002c28a8ba0d06d3a78c6b5cbf9b2ade051f0775ac4);
        NamespaceNode memory node = nodeDigest(left, right);
        assertEqNamespaceNode(node, expected);
    }

    function testNodeParity() external {
        Namespace memory nidMin = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000000);
        Namespace memory nidMax = Namespace(0xde, 0xadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefde);
        NamespaceNode memory expected =
            NamespaceNode(nidMin, nidMax, 0xc6960f535d4ab0aed075aed34a116725e8035012ceffe5405ae72abe3bcaa28f);
        NamespaceNode memory left =
            NamespaceNode(nidMin, nidMax, 0xdb55da3fc3098e9c42311c6013304ff36b19ef73d12ea932054b5ad51df4f49d);
        NamespaceNode memory right = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0xc75cb66ae28d8ebc6eded002c28a8ba0d06d3a78c6b5cbf9b2ade051f0775ac4
        );
        NamespaceNode memory node = nodeDigest(left, right);
        assertEqNamespaceNode(node, expected);
    }
}
