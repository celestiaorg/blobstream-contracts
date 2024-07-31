// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "./Constants.sol";

/// @notice Calculate the starting bit of the path to a leaf
/// @param numLeaves : The total number of leaves in the tree
/// @return startingBit : The starting bit of the path
// solhint-disable-next-line func-visibility
function getStartingBit(uint256 numLeaves) pure returns (uint256 startingBit) {
    // Determine height of the left subtree. This is the maximum path length, so all paths start at this offset from the right-most bit
    startingBit = 0;
    while ((1 << startingBit) < numLeaves) {
        startingBit += 1;
    }
    return Constants.MAX_HEIGHT - startingBit;
}

/// @notice Calculate the length of the path to a leaf
/// @param key: The key of the leaf
/// @param numLeaves: The total number of leaves in the tree
/// @return pathLength : The length of the path to the leaf
// solhint-disable-next-line func-visibility
function pathLengthFromKey(uint256 key, uint256 numLeaves) pure returns (uint256 pathLength) {
    if (numLeaves <= 1) {
        // if the number of leaves of the tree is 1 or 0, the path always is 0.
        return 0;
    }
    // Get the height of the left subtree. This is equal to the offset of the starting bit of the path
    pathLength = Constants.MAX_HEIGHT - getStartingBit(numLeaves);

    // Determine the number of leaves in the left subtree
    uint256 numLeavesLeftSubTree = (1 << (pathLength - 1));

    // If leaf is in left subtree, path length is full height of left subtree
    if (key <= numLeavesLeftSubTree - 1) {
        return pathLength;
    }
    // If left sub tree has only one leaf but key is not there, path has one additional step
    else if (numLeavesLeftSubTree == 1) {
        return 1;
    }
    // Otherwise, add 1 to height and recurse into right subtree
    else {
        return 1 + pathLengthFromKey(key - numLeavesLeftSubTree, numLeaves - numLeavesLeftSubTree);
    }
}

/// @notice Returns the minimum number of bits required to represent `x`; the
/// result is 0 for `x` == 0.
/// @param x Number.
function _bitsLen(uint256 x) pure returns (uint256) {
    uint256 count = 0;

    while (x != 0) {
        count++;
        x >>= 1;
    }

    return count;
}

/// @notice Returns the largest power of 2 less than `x`.
/// @param x Number.
function _getSplitPoint(uint256 x) pure returns (uint256) {
    // Note: since `x` is always an unsigned int * 2, the only way for this
    // to be violated is if the input == 0. Since the input is the end
    // index exclusive, an input of 0 is guaranteed to be invalid (it would
    // be a proof of inclusion of nothing, which is vacuous).
    require(x >= 1);

    uint256 bitLen = _bitsLen(x);
    uint256 k = 1 << (bitLen - 1);
    if (k == x) {
        k >>= 1;
    }
    return k;
}

/// @notice Returns the size of the subtree adjacent to `begin` that does
/// not overlap `end`.
/// @param begin Begin index, inclusive.
/// @param end End index, exclusive.
function _nextSubtreeSize(uint256 begin, uint256 end) pure returns (uint256) {
    uint256 ideal = _bitsTrailingZeroes(begin);
    uint256 max = _bitsLen(end - begin) - 1;
    if (ideal > max) {
        return 1 << max;
    }
    return 1 << ideal;
}

/// @notice Returns the number of trailing zero bits in `x`; the result is
/// 256 for `x` == 0.
/// @param x Number.
function _bitsTrailingZeroes(uint256 x) pure returns (uint256) {
    uint256 mask = 1;
    uint256 count = 0;

    while (x != 0 && mask & x == 0) {
        count++;
        x >>= 1;
    }

    return count;
}
