// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "ds-test/test.sol";

import "../../Types.sol";
import "../NamespaceNode.sol";
import "../NamespaceMerkleMultiproof.sol";
import "../NamespaceMerkleTree.sol";

/**
 * TEST VECTORS
 *
 * Data blocks: namespace, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x01
 * 0x0000000000000000000000000000000000000000000000000000000010 0x02
 * 0x0000000000000000000000000000000000000000000000000000000010 0x03
 * 0x0000000000000000000000000000000000000000000000000000000010 0x04
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x05
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x06
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x07
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x08
 *
 * Leaf nodes: min namespace, max namespace, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000010 0xfdb4e3c872666aa9869a1d46c8a5a0e735becdf17c62b9c3ccf4258449475bda
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000010 0x01a346b5c14a1b37e6c019eaff190f7a49718fb3036ec51360ee31de6ef58771
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000010 0x80cb31e074d15b09950610d26b9447d82a4c9beb04499fb51be9549c1a67f09f
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000010 0xc350aeddd5ada629057034f15d4545065213a7a28f9f9b77bdc71c4225145920
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x1617cc7010feae70f9ff07028da463c65ec19b1d6bafde31c7543718025e5efb
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x671157a4e268f7060abbdc4b48f091589555a0775a2694e6899833ec98fdb296
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x2669e36b48e95bd9903300e50c27c53984fc439f6235fade08e3f14e78a42aac
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x655790e24d376e9556a3cba9908a5d97f27faa050806ecfcb481861a83240bd5
 *
 * Inner nodes(depth = 2): min namespace, max namespace, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000010 0x0ba8a1c0dcf8798d617eeed351a350d4d68792b6c42e9beaf54dd30136ca7e38
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000010 0x6d43651bd68866cb3fc8d00512fa2ab570da16c2c5254a6a7671c0400b96441a
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x055a3ea75c438d752aeabbba94ed8fac93e0b32321256a65fde176dba14f5186
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x1b79ffd74644e8c287fe5f1dd70bc8ea02738697cebf2810ffb2dc5157485c40
 *
 * Inner nodes(depth = 1): min namespace, max namespace, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000010 0x23fcbabf97fa3bbef73038559ca480d0de5237762e42cac08090c48713eef910
 * 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffff 0x5aa3e7ea31995fdd38f41015275229b290a8ee4810521db766ad457b9a8373d6
 *
 * Root node: min namespace, max namespace, data
 * 0x0000000000000000000000000000000000000000000000000000000010 0x0000000000000000000000000000000000000000000000000000000010 0x5b3328b03a538d627db78668034089cb395f63d05b24fdf99558d36fe991d268
 *
 */
contract NamespaceMerkleMultiproofTest is DSTest {
    function setUp() external {}

    function assertEqNamespaceNode(NamespaceNode memory first, NamespaceNode memory second) internal {
        assertTrue(first.min.equalTo(second.min));
        assertTrue(first.max.equalTo(second.max));
        assertEq(first.digest, second.digest);
    }

    /// @notice Verify inclusion of leaves 0 and 1.
    function testVerifyMulti01() external {
        Namespace memory nid = Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010);
        NamespaceNode memory root = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            0x5b3328b03a538d627db78668034089cb395f63d05b24fdf99558d36fe991d268
        );
        NamespaceNode[] memory sideNodes = new NamespaceNode[](3);
        sideNodes[0] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            0xfdb4e3c872666aa9869a1d46c8a5a0e735becdf17c62b9c3ccf4258449475bda
        );
        sideNodes[1] = NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000010),
            0xc350aeddd5ada629057034f15d4545065213a7a28f9f9b77bdc71c4225145920
        );
        sideNodes[2] = NamespaceNode(
            PARITY_SHARE_NAMESPACE(),
            PARITY_SHARE_NAMESPACE(),
            0x5aa3e7ea31995fdd38f41015275229b290a8ee4810521db766ad457b9a8373d6
        );

        uint256 beginKey = 1;
        uint256 endKey = 3;
        NamespaceMerkleMultiproof memory proof = NamespaceMerkleMultiproof(beginKey, endKey, sideNodes);
        bytes[] memory data = new bytes[](2);
        data[0] = hex"02";
        data[1] = hex"03";
        bool isValid = NamespaceMerkleTree.verifyMulti(root, proof, nid, data);
        assertTrue(isValid);
    }
}
