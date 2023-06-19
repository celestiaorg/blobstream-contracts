// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.19;

type NamespaceID is bytes8;

using {equality as ==} for NamespaceID global;
using {lessthan as <} for NamespaceID global;
using {greaterthan as >} for NamespaceID global;

function equality(NamespaceID l, NamespaceID r) pure returns (bool) {
    return NamespaceID.unwrap(l) == NamespaceID.unwrap(r);
}

function lessthan(NamespaceID l, NamespaceID r) pure returns (bool) {
    return NamespaceID.unwrap(l) < NamespaceID.unwrap(r);
}

function greaterthan(NamespaceID l, NamespaceID r) pure returns (bool) {
    return NamespaceID.unwrap(l) > NamespaceID.unwrap(r);
}
