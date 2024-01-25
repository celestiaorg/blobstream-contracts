// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

/// @notice A representation of the Celestia-app namespace ID and its version.
/// See: https://celestiaorg.github.io/celestia-app/specs/namespace.html
struct Namespace {
    // The namespace version.
    bytes1 version;
    // The namespace ID.
    bytes28 id;
}

using {equalTo, lessThan, greaterThan, toBytes} for Namespace global;

function equalTo(Namespace memory l, Namespace memory r) pure returns (bool) {
    return l.toBytes() == r.toBytes();
}

function lessThan(Namespace memory l, Namespace memory r) pure returns (bool) {
    return l.toBytes() < r.toBytes();
}

function greaterThan(Namespace memory l, Namespace memory r) pure returns (bool) {
    return l.toBytes() > r.toBytes();
}

function toBytes(Namespace memory n) pure returns (bytes29) {
    return bytes29(abi.encodePacked(n.version, n.id));
}

function toNamespace(bytes29 n) pure returns (Namespace memory) {
    bytes memory id = new bytes(28);
    for (uint256 i = 1; i < 29; i++) {
        id[i - 1] = n[i];
    }
    return Namespace(n[0], bytes28(id));
}
