// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.19;

import "ds-test/test.sol";

import "../NamespaceNode.sol";
import "../NamespaceMerkleProof.sol";
import "../NamespaceMerkleTree.sol";

/**
 * TEST VECTORS
 *
 * Data blocks: namespace id, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x01
 * 0x0000000000000000000000000000000000000000000000000000000020 0x02
 * 0x0000000000000000000000000000000000000000000000000000000030 0x03
 * 0x0000000000000000000000000000000000000000000000000000000040 0x04
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x05
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x06
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x07
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x08
 *
 * Leaf nodes: min namespace, max namespace, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000010 0xfdb4e3c872666aa9869a1d46c8a5a0e735becdf17c62b9c3ccf4258449475bda
 * 0x0000000000000000000000000000000000000000000000000000000020 0x0000000000000000000000000000000000000000000000000000000020 0xc5fd5617b70207108c8d9bcf624b1eedf39b763af86f660255947674e043cd2c
 * 0x0000000000000000000000000000000000000000000000000000000030 0x0000000000000000000000000000000000000000000000000000000030 0x35e864d3e196ef0986fcf18eea2782c7e68794c7106dacc2a4f7e40d6d7c7069
 * 0x0000000000000000000000000000000000000000000000000000000040 0x0000000000000000000000000000000000000000000000000000000040 0xecdeb08b04dd92a17fec560e20c53269f65beff5a2626fa64f61bfa45b09119d
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x1617cc7010feae70f9ff07028da463c65ec19b1d6bafde31c7543718025e5efb
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x671157a4e268f7060abbdc4b48f091589555a0775a2694e6899833ec98fdb296
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x2669e36b48e95bd9903300e50c27c53984fc439f6235fade08e3f14e78a42aac
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x655790e24d376e9556a3cba9908a5d97f27faa050806ecfcb481861a83240bd5
 *
 * Inner nodes(depth = 2): min namespace, max namespace, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000020 0xd9654a9a8a4196b92358b3d94f1b7a21b5d8bddf57ab974ca9b869196c7c3cf1
 * 0x0000000000000000000000000000000000000000000000000000000030 0x0000000000000000000000000000000000000000000000000000000040 0x2aa20c7587b009772a9a88402b7cc8fcb82edc9e31754e95544a670a696f55a7
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x055a3ea75c438d752aeabbba94ed8fac93e0b32321256a65fde176dba14f5186
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x1b79ffd74644e8c287fe5f1dd70bc8ea02738697cebf2810ffb2dc5157485c40
 *
 * Inner nodes(depth = 1): min namespace, max namespace, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000040 0x203de1ab3986ac9302811f46d6e528fd66b4fb1b484a0de898a9af0f18e4403f
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x5aa3e7ea31995fdd38f41015275229b290a8ee4810521db766ad457b9a8373d6
 *
 * Root node: min namespace, max namespace, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000040 0xe6a9119463f4809bd9cd0e5df15fa4699d04dbf3f69d858f43a51014777d907b
 *
 */

contract NamespaceMerkleTreeTest is DSTest {
    function setUp() external {}

    function assertEqNamespaceNode(NamespaceNode memory first, NamespaceNode memory second) internal {
        assertEq(NamespaceID.unwrap(first.min), NamespaceID.unwrap(second.min));
        assertEq(NamespaceID.unwrap(first.max), NamespaceID.unwrap(second.max));
        assertEq(first.digest, second.digest);
    }

    function testVerifyNone() external {
        NamespaceID nid = NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000000);
        NamespaceNode memory root = NamespaceNode(nid, nid, sha256(""));
        NamespaceNode[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 0;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data;
        bool isValid = NamespaceMerkleTree.verify(root, proof, nid, data);
        assertTrue(!isValid);
    }

    function testVerifyOneLeafEmpty() external {
        NamespaceID nid = NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000000);
        NamespaceNode memory root =
            NamespaceNode(nid, nid, 0x0679246d6c4216de0daa08e5523fb2674db2b6599c3b72ff946b488a15290b62);
        NamespaceNode[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data;
        bool isValid = NamespaceMerkleTree.verify(root, proof, nid, data);
        assertTrue(isValid);
    }

    function testVerifyOneLeafSome() external {
        NamespaceID nid = NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000000);
        NamespaceNode memory root =
            NamespaceNode(nid, nid, 0x56d8381cfe28e8eb21da620145b7b977a74837720da5147b00bfab6f1b4af24d);
        NamespaceNode[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"deadbeef";
        bool isValid = NamespaceMerkleTree.verify(root, proof, nid, data);
        assertTrue(isValid);
    }

    function testVerifyOneLeaf01() external {
        NamespaceID nid = NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000000);
        NamespaceNode memory root =
            NamespaceNode(nid, nid, 0x353857cdb4c745eb9fdebbd8ec44093fabb9f08d437e2298d9e6afa1a409b30c);
        NamespaceNode[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        bool isValid = NamespaceMerkleTree.verify(root, proof, nid, data);
        assertTrue(isValid);
    }

    function testVerifyLeafOneOfTwo() external {
        NamespaceNode memory root = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000020),
            0xd9654a9a8a4196b92358b3d94f1b7a21b5d8bddf57ab974ca9b869196c7c3cf1
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](1);
        sideNodes[0] = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000020),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000020),
            0xc5fd5617b70207108c8d9bcf624b1eedf39b763af86f660255947674e043cd2c
        );

        uint256 key = 0;
        uint256 numLeaves = 2;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        bool isValid = NamespaceMerkleTree.verify(
            root, proof, NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010), data
        );
        assertTrue(isValid);
    }

    function testVerifyLeafOneOfFour() external {
        NamespaceNode memory root = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0x203de1ab3986ac9302811f46d6e528fd66b4fb1b484a0de898a9af0f18e4403f
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](2);
        sideNodes[0] = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000020),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000020),
            0xc5fd5617b70207108c8d9bcf624b1eedf39b763af86f660255947674e043cd2c
        );
        sideNodes[1] = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000030),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0x2aa20c7587b009772a9a88402b7cc8fcb82edc9e31754e95544a670a696f55a7
        );

        uint256 key = 0;
        uint256 numLeaves = 4;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        bool isValid = NamespaceMerkleTree.verify(
            root, proof, NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010), data
        );
        assertTrue(isValid);
    }

    function testVerifyLeafOneOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0xe6a9119463f4809bd9cd0e5df15fa4699d04dbf3f69d858f43a51014777d907b
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000020),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000020),
            0xc5fd5617b70207108c8d9bcf624b1eedf39b763af86f660255947674e043cd2c
        );
        sideNodes[1] = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000030),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0x2aa20c7587b009772a9a88402b7cc8fcb82edc9e31754e95544a670a696f55a7
        );
        sideNodes[2] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x5aa3e7ea31995fdd38f41015275229b290a8ee4810521db766ad457b9a8373d6
        );

        uint256 key = 0;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        bool isValid = NamespaceMerkleTree.verify(
            root, proof, NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010), data
        );
        assertTrue(isValid);
    }

    function testVerifyLeafSevenOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0xe6a9119463f4809bd9cd0e5df15fa4699d04dbf3f69d858f43a51014777d907b
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x655790e24d376e9556a3cba9908a5d97f27faa050806ecfcb481861a83240bd5
        );
        sideNodes[1] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x055a3ea75c438d752aeabbba94ed8fac93e0b32321256a65fde176dba14f5186
        );
        sideNodes[2] = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0x203de1ab3986ac9302811f46d6e528fd66b4fb1b484a0de898a9af0f18e4403f
        );

        uint256 key = 6;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"07";
        bool isValid = NamespaceMerkleTree.verify(root, proof, Constants.PARITY_SHARE_NAMESPACE_ID, data);
        assertTrue(isValid);
    }

    function testVerifyLeafEightOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0xe6a9119463f4809bd9cd0e5df15fa4699d04dbf3f69d858f43a51014777d907b
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x2669e36b48e95bd9903300e50c27c53984fc439f6235fade08e3f14e78a42aac
        );
        sideNodes[1] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x055a3ea75c438d752aeabbba94ed8fac93e0b32321256a65fde176dba14f5186
        );
        sideNodes[2] = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0x203de1ab3986ac9302811f46d6e528fd66b4fb1b484a0de898a9af0f18e4403f
        );

        uint256 key = 7;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"08";
        bool isValid = NamespaceMerkleTree.verify(root, proof, Constants.PARITY_SHARE_NAMESPACE_ID, data);
        assertTrue(isValid);
    }

    function testVerifyInnerLeafIsRoot() external {
        NamespaceID nid = NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000000);
        NamespaceNode memory root =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        NamespaceNode[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        NamespaceNode memory node =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        uint256 startingHeight = 1;
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, startingHeight);
        assertTrue(isValid);
    }

    function testVerifyInnerFalseForStartingHeightZero() external {
        NamespaceID nid = NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000020);
        NamespaceNode memory root =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        NamespaceNode[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        NamespaceNode memory node =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        uint256 startingHeight = 0;
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, startingHeight);
        assertTrue(!isValid);
    }

    function testVerifyInnerFalseForTooLargeKey() external {
        NamespaceID nid = NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000020);
        NamespaceNode memory root =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        NamespaceNode[] memory sideNodes;
        uint256 key = 3; // key is larger than num leaves
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        NamespaceNode memory node =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        uint256 startingHeight = 0;
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, startingHeight);
        assertTrue(!isValid);
    }

    function testVerifyInnerFalseForIncorrectProofLength() external {
        NamespaceID nid = NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000020);
        NamespaceNode memory root =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        NamespaceNode[] memory sideNodes = new NamespaceNode[](1);
        sideNodes[0] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x24ddc56b10cebbf760b3a744ad3a0e91093db34b4d22995f6de6dac918e38ae5
        );
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        NamespaceNode memory node =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        uint256 startingHeight = 0;
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, startingHeight);
        assertTrue(!isValid);
    }

    function testVerifyInnerOneOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0xe6a9119463f4809bd9cd0e5df15fa4699d04dbf3f69d858f43a51014777d907b
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](2);
        sideNodes[0] = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000030),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0x2aa20c7587b009772a9a88402b7cc8fcb82edc9e31754e95544a670a696f55a7
        );
        sideNodes[1] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x5aa3e7ea31995fdd38f41015275229b290a8ee4810521db766ad457b9a8373d6
        );
        NamespaceNode memory node = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000020),
            0xd9654a9a8a4196b92358b3d94f1b7a21b5d8bddf57ab974ca9b869196c7c3cf1
        );

        uint256 key = 0;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, 2);
        assertTrue(isValid);
    }

    function testVerifyInnerSevenOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0xe6a9119463f4809bd9cd0e5df15fa4699d04dbf3f69d858f43a51014777d907b
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](2);
        sideNodes[0] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x055a3ea75c438d752aeabbba94ed8fac93e0b32321256a65fde176dba14f5186
        );
        sideNodes[1] = NamespaceNode(
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000010),
            NamespaceID.wrap(0x0000000000000000000000000000000000000000000000000000000040),
            0x203de1ab3986ac9302811f46d6e528fd66b4fb1b484a0de898a9af0f18e4403f
        );
        NamespaceNode memory node = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x1b79ffd74644e8c287fe5f1dd70bc8ea02738697cebf2810ffb2dc5157485c40
        );

        uint256 key = 6;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, 2);
        assertTrue(isValid);
    }
}
