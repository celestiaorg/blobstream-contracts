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
0x0000000000000010 0x02
0x0000000000000010 0x03
0x0000000000000010 0x04
0xffffffffffffffff 0x05
0xffffffffffffffff 0x06
0xffffffffffffffff 0x07
0xffffffffffffffff 0x08

Leaf nodes: min namespace, max namespace, data
0x0000000000000010 0x0000000000000010 0x531d57c729081d903721f7584b2fa031c8308918779e47d9ef68991b7a30eadf
0x0000000000000010 0x0000000000000010 0x95fd61ffb03e598ac50a5e203026c357595efc47e45d27334269a5e1d68212ed
0x0000000000000010 0x0000000000000010 0xf052f0755203a9d6c1714cc99e1d4a06433bb2427ae55a2a213b9b5f9f7b36ff
0x0000000000000010 0x0000000000000010 0x47ea271b50de032f8b021550f19350a09dc8d0e2372d1c2d876794487517f16e
0xffffffffffffffff 0xffffffffffffffff 0x19628f21c9871b13b730e4f0c3f1eb0a033e5ea36e2d928e5580dce8276f3a1f
0xffffffffffffffff 0xffffffffffffffff 0x2041097752c4838ed7649383808636798c3bbd9dcb7c70054a833536eca57509
0xffffffffffffffff 0xffffffffffffffff 0x6d9dfdf16675fe8a327bbf048a8496686ceb1444268965477c00a73720ec743a
0xffffffffffffffff 0xffffffffffffffff 0x22efe732017c70f7ef831b3c0b841d11fdf2230cfabf178d9560e0e4beb5adcd

Inner nodes(depth = 2): min namespace, max namespace, data
0x0000000000000010 0x0000000000000010 0x0f9f4a39c2b64182463d0f3165d0ab513afcc9cbcb4ac561dffaf4c9ef54170f
0x0000000000000010 0x0000000000000010 0xd23ad50046b8fe83a22580c52ffe259d257863196574fd69d3e27968401eb99a
0xffffffffffffffff 0xffffffffffffffff 0xf11ca80de48fa801927a8061e234a29b8ab63b8239a9ea1efecf92688999602d
0xffffffffffffffff 0xffffffffffffffff 0xca4e971ee703d46a64ff78d8abf98618e79ce7d3c95e08f41806d3fb96c2bf0a

Inner nodes(depth = 1): min namespace, max namespace, data
0x0000000000000010 0x0000000000000010 0x31cb53b761143d0a1b6b7f096b64c6c0543266fda00654070a2d485d0a66b281
0xffffffffffffffff 0xffffffffffffffff 0xb3da10c55a205c40528dd8a65e5be607e8a08d5a02198fdd6407419ae3c373c9

Root node: min namespace, max namespace, data
0x0000000000000010 0x0000000000000010 0x02f1d195cf45f96f9bf0875cb3a8aedff5df35605fb3f50ce52a272c30822466
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
        bytes8 nid = 0x0000000000000010;
        NamespaceNode memory root = NamespaceNode(
            0x0000000000000010,
            0x0000000000000040,
            0x02f1d195cf45f96f9bf0875cb3a8aedff5df35605fb3f50ce52a272c30822466
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            0x0000000000000010,
            0x0000000000000010,
            0x531d57c729081d903721f7584b2fa031c8308918779e47d9ef68991b7a30eadf
        );
        sideNodes[1] = NamespaceNode(
            0x0000000000000010,
            0x0000000000000010,
            0x47ea271b50de032f8b021550f19350a09dc8d0e2372d1c2d876794487517f16e
        );
        sideNodes[2] = NamespaceNode(
            Constants.PARITY_SHARE_NAMESPACE_ID,
            Constants.PARITY_SHARE_NAMESPACE_ID,
            0xb3da10c55a205c40528dd8a65e5be607e8a08d5a02198fdd6407419ae3c373c9
        );

        uint256 beginKey = 0;
        uint256 endKey = 1;
        uint256 numLeaves = 8;
        NamespaceMerkleMultiproof memory proof = NamespaceMerkleMultiproof(beginKey, endKey, sideNodes, numLeaves);
        bytes[] memory data = new bytes[](2);
        data[0] = hex"02";
        data[1] = hex"03";
        bool isValid = NamespaceMerkleTree.verifyMulti(root, proof, nid, data);
        assertTrue(isValid);
    }
}
