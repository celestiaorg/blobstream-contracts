// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.19;

import "./Types.sol";

library Constants {
    ///////////////
    // Constants //
    ///////////////

    /// @dev Maximum tree height
    uint256 internal constant MAX_HEIGHT = 256;

    /// @dev The prefixes of leaves and nodes
    bytes1 internal constant LEAF_PREFIX = 0x00;
    bytes1 internal constant NODE_PREFIX = 0x01;

    /// @dev Parity share namespace ID
    NamespaceID internal constant PARITY_SHARE_NAMESPACE_ID =
        NamespaceID.wrap(0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF);
}
