pragma solidity ^0.8.22;

import "ds-test/test.sol";
import "forge-std/Vm.sol";
import "forge-std/console.sol";
import {_bytesToSharesV0, _createCommitment, _bytesToHexString} from "../Commitment.sol";
import {toNamespace, Namespace} from "../../tree/Types.sol";
import "../../tree/namespace/NamespaceMerkleTree.sol";
import "../../tree/namespace/NamespaceMerkleMultiproof.sol";

contract CommitmentTest is DSTest {
    Vm private constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    struct TestVector {
        string commitment;
        string data;
        string namespace;
        string shares;
    }

    function fromHexChar(uint8 c) public pure returns (uint8) {
        if (bytes1(c) >= bytes1("0") && bytes1(c) <= bytes1("9")) {
            return c - uint8(bytes1("0"));
        }
        if (bytes1(c) >= bytes1("a") && bytes1(c) <= bytes1("f")) {
            return 10 + c - uint8(bytes1("a"));
        }
        if (bytes1(c) >= bytes1("A") && bytes1(c) <= bytes1("F")) {
            return 10 + c - uint8(bytes1("A"));
        }
        revert("fail");
    }

    function fromHex(string memory s) public pure returns (bytes memory) {
        bytes memory ss = bytes(s);
        require(ss.length % 2 == 0); // length must be even
        bytes memory r = new bytes(ss.length / 2);
        for (uint256 i = 0; i < ss.length / 2; ++i) {
            r[i] = bytes1(fromHexChar(uint8(ss[2 * i])) * 16 + fromHexChar(uint8(ss[2 * i + 1])));
        }
        return r;
    }

    function compareStrings(string memory a, string memory b) public pure returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))));
    }

    function testBytesToSharesV0() external view {
        // test vectors were generated here: https://github.com/S1nus/share-test-vec-gen
        string memory path = "./src/lib/commitment/test/testVectors.json";
        string memory jsonData = vm.readFile(path);
        bytes memory vecsData = vm.parseJson(jsonData);
        TestVector[] memory vecs = abi.decode(vecsData, (TestVector[]));

        for (uint256 i = 0; i < vecs.length; i++) {
            bytes29 nsString = bytes29(fromHex(vecs[i].namespace));
            Namespace memory ns = toNamespace(nsString);
            bytes memory data = fromHex(vecs[i].data);
            (bytes[] memory shares, bool err) = _bytesToSharesV0(data, ns);
            string memory out = "";
            for (uint256 j = 0; j < shares.length; j++) {
                out = string.concat(out, _bytesToHexString(shares[j]));
            }
            // none of the test vectors should cause an error
            if (!compareStrings(out, vecs[i].shares)) {
                console.log("expected: ", vecs[i].shares);
                console.log("got: ", out);
                revert();
            }
        }
    }

    function testCreateCommitmentV0() external {
        string memory path = "./src/lib/commitment/test/testVectors2.json";
        string memory jsonData = vm.readFile(path);
        bytes memory vecsData = vm.parseJson(jsonData);
        TestVector[] memory vecs = abi.decode(vecsData, (TestVector[]));

        for (uint256 i = 0; i < vecs.length; i++) {
            bytes29 nsString = bytes29(fromHex(vecs[i].namespace));
            Namespace memory ns = toNamespace(nsString);
            bytes memory data = fromHex(vecs[i].data);
            (bytes[] memory shares, bool err) = _bytesToSharesV0(data, ns);
            bytes32 commitment = _createCommitment(shares, ns);
            if (!compareStrings(_bytesToHexString(abi.encodePacked(commitment)), vecs[i].commitment)) {
                console.log("Commitment mismatch for vector ", i);
                console.log("expected: ", vecs[i].commitment);
                console.log("got: ", _bytesToHexString(abi.encodePacked(commitment)));
                revert();
            }
        }
    }
}
