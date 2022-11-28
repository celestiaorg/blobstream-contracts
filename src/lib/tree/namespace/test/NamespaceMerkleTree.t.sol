// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import "ds-test/test.sol";

import "../NamespaceNode.sol";
import "../NamespaceMerkleProof.sol";
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
0x0000000000000010 0x0000000000000010 0x531d57c729081d903721f7584b2fa031c8308918779e47d9ef68991b7a30eadf
0x0000000000000020 0x0000000000000020 0xd3e46dc7795fef402ede7504a037e43af169f19b76cbdb2c7abb12252b6b2ecc
0x0000000000000030 0x0000000000000030 0x78c554db7e421f683df27a171146ca2aa4659e1ec01e6c61fc7291b28f8da6dd
0x0000000000000040 0x0000000000000040 0xa8bcefebd33001489fc678d3891c6fe71ce7ec0b5cbd2fc37fb5178c41a23ac3
0xffffffffffffffff 0xffffffffffffffff 0x19628f21c9871b13b730e4f0c3f1eb0a033e5ea36e2d928e5580dce8276f3a1f
0xffffffffffffffff 0xffffffffffffffff 0x2041097752c4838ed7649383808636798c3bbd9dcb7c70054a833536eca57509
0xffffffffffffffff 0xffffffffffffffff 0x6d9dfdf16675fe8a327bbf048a8496686ceb1444268965477c00a73720ec743a
0xffffffffffffffff 0xffffffffffffffff 0x22efe732017c70f7ef831b3c0b841d11fdf2230cfabf178d9560e0e4beb5adcd

Inner nodes(depth = 2): min namespace, max namespace, data
0x0000000000000010 0x0000000000000020 0x8985f8bdc931cbae27e6ce4851ffe91f30f3c5c54c785aa9beac6e1b8494e63c
0x0000000000000030 0x0000000000000040 0xbbd409114c569aa80cbc35146b183ac85f6e218f6345d59b7f6822ae440a7f9c
0xffffffffffffffff 0xffffffffffffffff 0xf11ca80de48fa801927a8061e234a29b8ab63b8239a9ea1efecf92688999602d
0xffffffffffffffff 0xffffffffffffffff 0xca4e971ee703d46a64ff78d8abf98618e79ce7d3c95e08f41806d3fb96c2bf0a

Inner nodes(depth = 1): min namespace, max namespace, data
0x0000000000000010 0x0000000000000040 0x31cb53b761143d0a1b6b7f096b64c6c0543266fda00654070a2d485d0a66b281
0xffffffffffffffff 0xffffffffffffffff 0xb3da10c55a205c40528dd8a65e5be607e8a08d5a02198fdd6407419ae3c373c9

Root node: min namespace, max namespace, data
0x0000000000000010 0x0000000000000040 0x135fc2adb4f8569783f67b463d2245d95ea98046523a02ce015edaa292a92085
**/

contract NamespaceMerkleTreeTest is DSTest {
    function setUp() external {}

    function assertEqNamespaceNode(NamespaceNode memory first, NamespaceNode memory second) internal {
        assertEq(first.min, second.min);
        assertEq(first.max, second.max);
        assertEq(first.digest, second.digest);
    }

    function testVerifyNone() external {
        bytes8 nid = 0x0000000000000000;
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
        bytes8 nid = 0x0000000000000000;
        NamespaceNode memory root = NamespaceNode(
            nid,
            nid,
            0x3e7077fd2f66d689e0cee6a7cf5b37bf2dca7c979af356d0a31cbc5c85605c7d
        );
        NamespaceNode[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data;
        bool isValid = NamespaceMerkleTree.verify(root, proof, nid, data);
        assertTrue(isValid);
    }

    function testVerifyOneLeafSome() external {
        bytes8 nid = 0x0000000000000000;
        NamespaceNode memory root = NamespaceNode(
            nid,
            nid,
            0xaa96931c3a623dc18aa70c476c74ef95fc8a828c5c0d664f5f64963a2f02be13
        );
        NamespaceNode[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"deadbeef";
        bool isValid = NamespaceMerkleTree.verify(root, proof, nid, data);
        assertTrue(isValid);
    }

    function testVerifyOneLeaf01() external {
        bytes8 nid = 0x0000000000000000;
        NamespaceNode memory root = NamespaceNode(
            nid,
            nid,
            0x52807a4607ea5debf0b7d4ccb452f4af03e16b06a8e0aa0dfe177db1ff02123d
        );
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
            0x0000000000000010,
            0x0000000000000020,
            0x8985f8bdc931cbae27e6ce4851ffe91f30f3c5c54c785aa9beac6e1b8494e63c
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](1);
        sideNodes[0] = NamespaceNode(
            0x0000000000000020,
            0x0000000000000020,
            0xd3e46dc7795fef402ede7504a037e43af169f19b76cbdb2c7abb12252b6b2ecc
        );

        uint256 key = 0;
        uint256 numLeaves = 2;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        bool isValid = NamespaceMerkleTree.verify(root, proof, 0x0000000000000010, data);
        assertTrue(isValid);
    }

    function testVerifyLeafOneOfFour() external {
        NamespaceNode memory root = NamespaceNode(
            0x0000000000000010,
            0x0000000000000040,
            0x31cb53b761143d0a1b6b7f096b64c6c0543266fda00654070a2d485d0a66b281
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](2);
        sideNodes[0] = NamespaceNode(
            0x0000000000000020,
            0x0000000000000020,
            0xd3e46dc7795fef402ede7504a037e43af169f19b76cbdb2c7abb12252b6b2ecc
        );
        sideNodes[1] = NamespaceNode(
            0x0000000000000030,
            0x0000000000000040,
            0xbbd409114c569aa80cbc35146b183ac85f6e218f6345d59b7f6822ae440a7f9c
        );

        uint256 key = 0;
        uint256 numLeaves = 4;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        bool isValid = NamespaceMerkleTree.verify(root, proof, 0x0000000000000010, data);
        assertTrue(isValid);
    }

    function testVerifyLeafOneOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            0x0000000000000010,
            0x0000000000000040,
            0x135fc2adb4f8569783f67b463d2245d95ea98046523a02ce015edaa292a92085
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            0x0000000000000020,
            0x0000000000000020,
            0xd3e46dc7795fef402ede7504a037e43af169f19b76cbdb2c7abb12252b6b2ecc
        );
        sideNodes[1] = NamespaceNode(
            0x0000000000000030,
            0x0000000000000040,
            0xbbd409114c569aa80cbc35146b183ac85f6e218f6345d59b7f6822ae440a7f9c
        );
        sideNodes[2] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0xb3da10c55a205c40528dd8a65e5be607e8a08d5a02198fdd6407419ae3c373c9
        );

        uint256 key = 0;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"01";
        bool isValid = NamespaceMerkleTree.verify(root, proof, 0x0000000000000010, data);
        assertTrue(isValid);
    }

    function testVerifyLeafSevenOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            0x0000000000000010,
            0x0000000000000040,
            0x135fc2adb4f8569783f67b463d2245d95ea98046523a02ce015edaa292a92085
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x22efe732017c70f7ef831b3c0b841d11fdf2230cfabf178d9560e0e4beb5adcd
        );
        sideNodes[1] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0xf11ca80de48fa801927a8061e234a29b8ab63b8239a9ea1efecf92688999602d
        );
        sideNodes[2] = NamespaceNode(
            0x0000000000000010,
            0x0000000000000040,
            0x31cb53b761143d0a1b6b7f096b64c6c0543266fda00654070a2d485d0a66b281
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
            0x0000000000000010,
            0x0000000000000040,
            0x135fc2adb4f8569783f67b463d2245d95ea98046523a02ce015edaa292a92085
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x6d9dfdf16675fe8a327bbf048a8496686ceb1444268965477c00a73720ec743a
        );
        sideNodes[1] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0xf11ca80de48fa801927a8061e234a29b8ab63b8239a9ea1efecf92688999602d
        );
        sideNodes[2] = NamespaceNode(
            0x0000000000000010,
            0x0000000000000040,
            0x31cb53b761143d0a1b6b7f096b64c6c0543266fda00654070a2d485d0a66b281
        );

        uint256 key = 7;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bytes memory data = hex"08";
        bool isValid = NamespaceMerkleTree.verify(root, proof, Constants.PARITY_SHARE_NAMESPACE_ID, data);
        assertTrue(isValid);
    }

    function testVerifyInnerLeafIsRoot() external {
        bytes8 nid = 0x0000000000000000;
        NamespaceNode memory root = NamespaceNode(
            nid,
            nid,
            0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70
        );
        NamespaceNode[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        NamespaceNode memory node = NamespaceNode(
            nid,
            nid,
            0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70
        );
        uint256 startingHeight = 1;
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, startingHeight);
        assertTrue(isValid);
    }

    function testVerifyInnerFalseForStartingHeightZero() external {
        bytes8 nid = 0x0000000000000000;
        NamespaceNode memory root = NamespaceNode(
            nid,
            nid,
            0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70
        );
        NamespaceNode[] memory sideNodes;
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        NamespaceNode memory node = NamespaceNode(
            nid,
            nid,
            0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70
        );
        uint256 startingHeight = 0;
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, startingHeight);
        assertTrue(!isValid);
    }

    function testVerifyInnerFalseForTooLargeKey() external {
        bytes8 nid = 0x0000000000000000;
        NamespaceNode memory root = NamespaceNode(
            nid,
            nid,
            0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70
        );
        NamespaceNode[] memory sideNodes;
        uint256 key = 3; // key is larger than num leaves
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        NamespaceNode memory node = NamespaceNode(
            nid,
            nid,
            0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70
        );
        uint256 startingHeight = 0;
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, startingHeight);
        assertTrue(!isValid);
    }

    function testVerifyInnerFalseForIncorrectProofLength() external {
        bytes8 nid = 0x0000000000000000;
        NamespaceNode memory root = NamespaceNode(
            nid,
            nid,
            0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](1);
        sideNodes[0] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x24ddc56b10cebbf760b3a744ad3a0e91093db34b4d22995f6de6dac918e38ae5
        );
        uint256 key = 0;
        uint256 numLeaves = 1;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        NamespaceNode memory node = NamespaceNode(
            nid,
            nid,
            0xc59fa9c4ec515726c2b342544433f844c7b930cf7a5e7abab593332453ceaf70
        );
        uint256 startingHeight = 0;
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, startingHeight);
        assertTrue(!isValid);
    }

    function testVerifyInnerOneOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            0x0000000000000010,
            0x0000000000000040,
            0x16c760661bc5ed683d27dc2f045a81a67e837928527e0de209a195b6db60f462
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](2);
        sideNodes[0] = NamespaceNode(
            0x0000000000000030,
            0x0000000000000040,
            0x3f8ded56b6a8d4e1e36832e8be93234e2e3a18c1a42edfb505ecc09f0039a10f
        );
        sideNodes[1] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x27209d167edf7ea1463f462b850471ce31b124b0b3405c33f9c39e692c9170da
        );
        NamespaceNode memory node = NamespaceNode(
            0x0000000000000010,
            0x0000000000000020,
            0xe0a6f55a5c2d86e0057b89d79bf5c6c3fdc5a40061566c39e93077556e2e3482
        );

        uint256 key = 0;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, 2);
        assertTrue(isValid);
    }

    function testVerifyInnerSevenOfEight() external {
        NamespaceNode memory root = NamespaceNode(
            0x0000000000000010,
            0x0000000000000040,
            0x16c760661bc5ed683d27dc2f045a81a67e837928527e0de209a195b6db60f462
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](2);
        sideNodes[0] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x61d6762ff063c2008a412246bc6bb370885c4e1a935ca28ed8699dc5c68ff28a
        );
        sideNodes[1] = NamespaceNode(
            0x0000000000000010,
            0x0000000000000040,
            0xed6a82bfecd113f693065e3b1f271f21150b6d793917402f6c05a01feb6b3eb8
        );
        NamespaceNode memory node = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0x9086b06cbc327959e3c34546aadc886300aff3e5c8f96a328267abf64ca5d25b
        );

        uint256 key = 6;
        uint256 numLeaves = 8;
        NamespaceMerkleProof memory proof = NamespaceMerkleProof(sideNodes, key, numLeaves);
        bool isValid = NamespaceMerkleTree.verifyInner(root, proof, node, 2);
        assertTrue(isValid);
    }
}
