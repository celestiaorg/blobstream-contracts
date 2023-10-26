// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "ds-test/test.sol";

import "../Utils.sol";

contract UtilsTest is DSTest {
    function testPathLengthFromKey0_2() external {
        assertEq(pathLengthFromKey(0, 2), 1);
    }

    function testPathLengthFromKey1_2() external {
        assertEq(pathLengthFromKey(1, 2), 1);
    }

    function testPathLengthFromKey0_8() external {
        assertEq(pathLengthFromKey(0, 8), 3);
    }

    function testPathLengthFromKey1_8() external {
        assertEq(pathLengthFromKey(1, 8), 3);
    }

    function testPathLengthFromKey2_8() external {
        assertEq(pathLengthFromKey(2, 8), 3);
    }

    function testPathLengthFromKey3_8() external {
        assertEq(pathLengthFromKey(3, 8), 3);
    }

    function testPathLengthFromKey4_8() external {
        assertEq(pathLengthFromKey(4, 8), 3);
    }

    function testPathLengthFromKey5_8() external {
        assertEq(pathLengthFromKey(5, 8), 3);
    }

    function testPathLengthFromKey6_8() external {
        assertEq(pathLengthFromKey(6, 8), 3);
    }

    function testPathLengthFromKey7_8() external {
        assertEq(pathLengthFromKey(7, 8), 3);
    }
}
