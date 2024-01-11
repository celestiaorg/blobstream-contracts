// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "ds-test/test.sol";

import "../NamespaceNode.sol";
import "../NamespaceMerkleProof.sol";
import "../NamespaceMerkleTree.sol";
import "../../Constants.sol";

/**
 * TEST VECTORS
 *
 * Data blocks: Namespace, data
 * Data blocks: Namespace, data
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
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000020 0x1dae5c3d39a8bf31ea33ba368238a52f816cd50485c580565609554cf360c91f
 * 0x0000000000000000000000000000000000000000000000000000000030 0x0000000000000000000000000000000000000000000000000000000040 0x2aa20c7587b009772a9a88402b7cc8fcb82edc9e31754e95544a670a696f55a7
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x055a3ea75c438d752aeabbba94ed8fac93e0b32321256a65fde176dba14f5186
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x1b79ffd74644e8c287fe5f1dd70bc8ea02738697cebf2810ffb2dc5157485c40
 *
 * Inner nodes(depth = 1): min namespace, max namespace, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000040 0xa8dcd9f365fb64aa6d72b5027fe74db0fc7d009c2d75c7b9b9656927281cb35e
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x5aa3e7ea31995fdd38f41015275229b290a8ee4810521db766ad457b9a8373d6
 *
 * Root node: min namespace, max namespace, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000040 0x34e6541306dc4e57a5a2a9ef57a46d5705ed09efb8c6a02580d3a972922b6862
 *
 */
contract NamespaceMerkleTreeTest is DSTest {
    function setUp() external {}

    function assertEqNamespaceNode(NamespaceNode memory first, NamespaceNode memory second) internal {
        assertTrue(first.min.equalTo(second.min));
        assertTrue(first.max.equalTo(second.max));
        assertEq(first.digest, second.digest);
    }

    function testVerifyNone() external {
        Namespace memory nid = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000000);
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
        Namespace memory nid = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000000);
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
        Namespace memory nid = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000000);
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
        Namespace memory nid = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000000);
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
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020),
            0x1dae5c3d39a8bf31ea33ba368238a52f816cd50485c580565609554cf360c91f
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](1);
        sideNodes[0] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020),
            0xc5fd5617b70207108c8d9bcf624b1eedf39b763af86f660255947674e043cd2c
        );

        uint256 key = 0;
        uint256 numLeaves = 2;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        bool isValid = NamespaceMerkleTree.verify(
            root, proof, Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010), data
        );
        assertTrue(isValid);
    }

    function testVerifyLeafOneOfFour() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0xa8dcd9f365fb64aa6d72b5027fe74db0fc7d009c2d75c7b9b9656927281cb35e
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](2);
        sideNodes[0] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020),
            0xc5fd5617b70207108c8d9bcf624b1eedf39b763af86f660255947674e043cd2c
        );
        sideNodes[1] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000030),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0x2aa20c7587b009772a9a88402b7cc8fcb82edc9e31754e95544a670a696f55a7
        );

        uint256 key = 0;
        uint256 numLeaves = 4;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        bool isValid = NamespaceMerkleTree.verify(
            root, proof, Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010), data
        );
        assertTrue(isValid);
    }

    function testVerifyLeafOneOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0x34e6541306dc4e57a5a2a9ef57a46d5705ed09efb8c6a02580d3a972922b6862
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020),
            0xc5fd5617b70207108c8d9bcf624b1eedf39b763af86f660255947674e043cd2c
        );
        sideNodes[1] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000030),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0x2aa20c7587b009772a9a88402b7cc8fcb82edc9e31754e95544a670a696f55a7
        );
        sideNodes[2] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x5aa3e7ea31995fdd38f41015275229b290a8ee4810521db766ad457b9a8373d6
        );

        uint256 key = 0;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        bool isValid = NamespaceMerkleTree.verify(
            root, proof, Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010), data
        );
        assertTrue(isValid);
    }

    function testVerifyLeafSevenOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0x34e6541306dc4e57a5a2a9ef57a46d5705ed09efb8c6a02580d3a972922b6862
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x655790e24d376e9556a3cba9908a5d97f27faa050806ecfcb481861a83240bd5
        );
        sideNodes[1] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x055a3ea75c438d752aeabbba94ed8fac93e0b32321256a65fde176dba14f5186
        );
        sideNodes[2] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0xa8dcd9f365fb64aa6d72b5027fe74db0fc7d009c2d75c7b9b9656927281cb35e
        );

        uint256 key = 6;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"07";
        bool isValid = NamespaceMerkleTree.verify(root, proof, PARITY_SHARE_NAMESPACE(), data);
        assertTrue(isValid);
    }

    function testVerifyLeafEightOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0x34e6541306dc4e57a5a2a9ef57a46d5705ed09efb8c6a02580d3a972922b6862
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x2669e36b48e95bd9903300e50c27c53984fc439f6235fade08e3f14e78a42aac
        );
        sideNodes[1] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x055a3ea75c438d752aeabbba94ed8fac93e0b32321256a65fde176dba14f5186
        );
        sideNodes[2] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0xa8dcd9f365fb64aa6d72b5027fe74db0fc7d009c2d75c7b9b9656927281cb35e
        );

        uint256 key = 7;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"08";
        bool isValid = NamespaceMerkleTree.verify(root, proof, PARITY_SHARE_NAMESPACE(), data);
        assertTrue(isValid);
    }

    function testVerifyLeafFiveOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0x34e6541306dc4e57a5a2a9ef57a46d5705ed09efb8c6a02580d3a972922b6862
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x671157a4e268f7060abbdc4b48f091589555a0775a2694e6899833ec98fdb296
        );
        sideNodes[1] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x1b79ffd74644e8c287fe5f1dd70bc8ea02738697cebf2810ffb2dc5157485c40
        );
        sideNodes[2] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0xa8dcd9f365fb64aa6d72b5027fe74db0fc7d009c2d75c7b9b9656927281cb35e
        );

        uint256 key = 4;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"05";
        bool isValid = NamespaceMerkleTree.verify(root, proof, PARITY_SHARE_NAMESPACE(), data);
        assertTrue(isValid);
    }

    function testVerifyLeafFourOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0x34e6541306dc4e57a5a2a9ef57a46d5705ed09efb8c6a02580d3a972922b6862
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000030),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000030),
            0x35e864d3e196ef0986fcf18eea2782c7e68794c7106dacc2a4f7e40d6d7c7069
        );
        sideNodes[1] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020),
            0x1dae5c3d39a8bf31ea33ba368238a52f816cd50485c580565609554cf360c91f
        );
        sideNodes[2] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x5aa3e7ea31995fdd38f41015275229b290a8ee4810521db766ad457b9a8373d6
        );

        uint256 key = 3;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"04";
        bool isValid = NamespaceMerkleTree.verify(
            root, proof, Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040), data
        );
        assertTrue(isValid);
    }

    function testVerifyLeafThreeOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0x34e6541306dc4e57a5a2a9ef57a46d5705ed09efb8c6a02580d3a972922b6862
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0xecdeb08b04dd92a17fec560e20c53269f65beff5a2626fa64f61bfa45b09119d
        );
        sideNodes[1] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020),
            0x1dae5c3d39a8bf31ea33ba368238a52f816cd50485c580565609554cf360c91f
        );
        sideNodes[2] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x5aa3e7ea31995fdd38f41015275229b290a8ee4810521db766ad457b9a8373d6
        );

        uint256 key = 2;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"03";
        bool isValid = NamespaceMerkleTree.verify(
            root, proof, Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000030), data
        );
        assertTrue(isValid);
    }

    function testVerifyLeafFiveOfSeven() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0xfe7100a7170cba2065c48e01cb18772ed93865100bb7610aed3f614829c87a48
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x671157a4e268f7060abbdc4b48f091589555a0775a2694e6899833ec98fdb296
        );
        sideNodes[1] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x2669e36b48e95bd9903300e50c27c53984fc439f6235fade08e3f14e78a42aac
        );
        sideNodes[2] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0xa8dcd9f365fb64aa6d72b5027fe74db0fc7d009c2d75c7b9b9656927281cb35e
        );

        uint256 key = 4;
        uint256 numLeaves = 7;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"05";
        bool isValid = NamespaceMerkleTree.verify(root, proof, PARITY_SHARE_NAMESPACE(), data);
        assertTrue(isValid);
    }

    function testVerifyLeafNineOfTen() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000060),
            0x21013157ca1c0d454c988665e05894f5cf9422928552349ac3fd359bd1d39ac1
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](2);
        sideNodes[0] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x8ecd4167595d96b6caf19871584b07f255a4d80037b122c9f1f71acb1366a1ae
        );
        sideNodes[1] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000060),
            0xee695202b2d3090a2319e7491483cf50e71a5907cebcf1fed4d02daa02f39827
        );

        uint256 key = 8;
        uint256 numLeaves = 10;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"09";
        bool isValid = NamespaceMerkleTree.verify(root, proof, PARITY_SHARE_NAMESPACE(), data);
        assertTrue(isValid);
    }

    function testVerifyLeafTwelveOfThirteen() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000060),
            0xcdf9d9d4b408a7bf1ec5653dcb5f8cda23a329754890b63344e706302ef70e43
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](4);
        sideNodes[0] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x311733a16ba3f14dca59dcd88e6b64276613cac5a9e20a4b228c520722851b3a
        );
        sideNodes[1] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x8137f8ca69ccd4d39d47836ace7aa22b010222eaa904e67a6ff9bf05542f7124
        );
        sideNodes[2] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x3666000822ff8e0e5bf01c170264fe39dc38d887a5ec5e87b4f72b328a323ec5
        );
        sideNodes[3] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000060),
            0xee695202b2d3090a2319e7491483cf50e71a5907cebcf1fed4d02daa02f39827
        );

        uint256 key = 11;
        uint256 numLeaves = 13;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"0c";
        bool isValid = NamespaceMerkleTree.verify(root, proof, PARITY_SHARE_NAMESPACE(), data);
        assertTrue(isValid);
    }

    function testVerifyLeafThirteenOfThirteen() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000060),
            0xcdf9d9d4b408a7bf1ec5653dcb5f8cda23a329754890b63344e706302ef70e43
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](2);
        sideNodes[0] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x31501dc8a114b0aa3cde0f4f99f0643760b3b11303ab1ee568538f3e5769fbfe
        );
        sideNodes[1] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000060),
            0xee695202b2d3090a2319e7491483cf50e71a5907cebcf1fed4d02daa02f39827
        );

        uint256 key = 12;
        uint256 numLeaves = 13;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"0d";
        bool isValid = NamespaceMerkleTree.verify(root, proof, PARITY_SHARE_NAMESPACE(), data);
        assertTrue(isValid);
    }

    function testVerifyInternalNodeOneAndTwoOfFour() external pure {
        uint256 key = 1;
        uint256 numLeaves = 4;
        uint256 startingHeight = 2;
        NamespaceNode[] memory sideNodes = new NamespaceNode[](1);

        bytes memory data = hex"01";
        NamespaceNode memory node1 =
            leafDigest(Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010), data);
        NamespaceNode memory node2 =
            leafDigest(Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020), data);
        NamespaceNode memory node3 =
            leafDigest(Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000030), data);
        NamespaceNode memory node4 =
            leafDigest(Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040), data);
        NamespaceNode memory node1_2 = nodeDigest(node1, node2);
        NamespaceNode memory node3_4 = nodeDigest(node3, node4);
        NamespaceNode memory root = nodeDigest(node1_2, node3_4);
        NamespaceNode memory startingNode = node1_2;

        sideNodes[0] = node3_4;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, startingNode, startingHeight);
        assert(isValid);
    }

    function testVerifyInternalNodeOneAndTwoOfThree() external pure {
        uint256 key = 0;
        uint256 numLeaves = 3;
        uint256 startingHeight = 2;
        NamespaceNode[] memory sideNodes = new NamespaceNode[](1);

        bytes memory data = hex"01";
        NamespaceNode memory node1 =
            leafDigest(Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010), data);
        NamespaceNode memory node2 =
            leafDigest(Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020), data);
        NamespaceNode memory node3 =
            leafDigest(Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000030), data);
        NamespaceNode memory node1_2 = nodeDigest(node1, node2);
        NamespaceNode memory root = nodeDigest(node1_2, node3);
        NamespaceNode memory startingNode = node1_2;

        sideNodes[0] = node3;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, startingNode, startingHeight);
        assert(isValid);
    }

    function testVerifyInnerLeafIsRoot() external {
        Namespace memory nid = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000000);
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
        Namespace memory nid = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020);
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
        Namespace memory nid = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020);
        NamespaceNode memory root =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        NamespaceNode[] memory sideNodes;
        uint256 key = 3; // key is larger than num leaves
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        NamespaceNode memory node =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        uint256 startingHeight = 1;
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, startingHeight);
        assertTrue(!isValid);
    }

    function testVerifyInnerFalseForIncorrectProofLength() external {
        Namespace memory nid = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020);
        NamespaceNode memory root =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        NamespaceNode[] memory sideNodes = new NamespaceNode[](1);
        sideNodes[0] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x24ddc56b10cebbf760b3a744ad3a0e91093db34b4d22995f6de6dac918e38ae5
        );
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        NamespaceNode memory node =
            NamespaceNode(nid, nid, 0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70);
        uint256 startingHeight = 1;
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, startingHeight);
        assertTrue(!isValid);
    }

    function testVerifyInnerOneOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0x34e6541306dc4e57a5a2a9ef57a46d5705ed09efb8c6a02580d3a972922b6862
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](2);
        sideNodes[0] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000030),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0x2aa20c7587b009772a9a88402b7cc8fcb82edc9e31754e95544a670a696f55a7
        );
        sideNodes[1] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x5aa3e7ea31995fdd38f41015275229b290a8ee4810521db766ad457b9a8373d6
        );
        NamespaceNode memory node = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000020),
            0x1dae5c3d39a8bf31ea33ba368238a52f816cd50485c580565609554cf360c91f
        );

        uint256 key = 0;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, 2);
        assertTrue(isValid);
    }

    function testVerifyInnerSevenOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0x34e6541306dc4e57a5a2a9ef57a46d5705ed09efb8c6a02580d3a972922b6862
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](2);
        sideNodes[0] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x055a3ea75c438d752aeabbba94ed8fac93e0b32321256a65fde176dba14f5186
        );
        sideNodes[1] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000040),
            0xa8dcd9f365fb64aa6d72b5027fe74db0fc7d009c2d75c7b9b9656927281cb35e
        );
        NamespaceNode memory node = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x1b79ffd74644e8c287fe5f1dd70bc8ea02738697cebf2810ffb2dc5157485c40
        );

        uint256 key = 6;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, 2);
        assertTrue(isValid);
    }
}
