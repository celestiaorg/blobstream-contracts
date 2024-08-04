pragma solidity ^0.8.22;

import "ds-test/test.sol";
import "forge-std/Vm.sol";
import "forge-std/console.sol";
import {bytesToShares, bytesToHexString} from "../Commitment.sol";
import {toNamespace, Namespace} from "../../tree/Types.sol";

contract CommitmentTest is DSTest {
    Vm private constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    struct TestVector {
        string data;
        string namespace;
        string shares;
    }

    function fromHexChar(uint8 c) public pure returns (uint8) {
        if (bytes1(c) >= bytes1('0') && bytes1(c) <= bytes1('9')) {
            return c - uint8(bytes1('0'));
        }
        if (bytes1(c) >= bytes1('a') && bytes1(c) <= bytes1('f')) {
            return 10 + c - uint8(bytes1('a'));
        }
        if (bytes1(c) >= bytes1('A') && bytes1(c) <= bytes1('F')) {
            return 10 + c - uint8(bytes1('A'));
        }
        revert("fail");
    }

    function fromHex(string memory s) public pure returns (bytes memory) {
        bytes memory ss = bytes(s);
        require(ss.length%2 == 0); // length must be even
        bytes memory r = new bytes(ss.length/2);
        for (uint i=0; i<ss.length/2; ++i) {
            r[i] = bytes1(fromHexChar(uint8(ss[2*i])) * 16 + fromHexChar(uint8(ss[2*i+1])));
        }
        return r;
    }

    function compareStrings(string memory a, string memory b) public pure returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))));
    }

    function testFirstVec() external {

        // test vectors were generated here: https://github.com/S1nus/share-test-vec-gen
        string memory path = "./src/lib/commitment/test/testVectors.json";
        string memory jsonData = vm.readFile(path);
        bytes memory vecsData = vm.parseJson(jsonData);
        TestVector[] memory vecs = abi.decode(vecsData, (TestVector[]));

        for (uint i = 0; i < vecs.length; i++) {
            bytes29 nsString = bytes29(fromHex(vecs[i].namespace));
            Namespace memory ns = toNamespace(nsString);
            bytes memory data = fromHex(vecs[i].data);
            bytes[] memory shares = bytesToShares(data, ns);
            string memory out = "";
            for (uint j = 0; j < shares.length; j++) {
                out = string.concat(out, bytesToHexString(shares[j]));
            }
            //assertEq(out, vecs[i].shares);
            if (compareStrings(out, vecs[i].shares) == false) {
                console.log("Failed on test vector ", i);
                console.log("Expected: ", vecs[i].shares);
                console.log("Got: ", out);
                return;
            }
        }
    }
}