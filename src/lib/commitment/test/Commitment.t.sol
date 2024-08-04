pragma solidity ^0.8.22;

import "ds-test/test.sol";
import "forge-std/Vm.sol";
import "forge-std/console.sol";
import {bytesToShares} from "../Commitment.sol";
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

    function bytesToHexString(bytes memory buffer) public pure returns (string memory) {

        // Fixed buffer size for hexadecimal convertion
        bytes memory converted = new bytes(buffer.length * 2);

        bytes memory _base = "0123456789abcdef";

        for (uint256 i = 0; i < buffer.length; i++) {
            converted[i * 2] = _base[uint8(buffer[i]) / _base.length];
            converted[i * 2 + 1] = _base[uint8(buffer[i]) % _base.length];
        }

        return string(abi.encodePacked(converted));
    }


    function testFirstVec() external {
        string memory path = "./src/lib/commitment/test/testVectors.json";
        string memory jsonData = vm.readFile(path);
        bytes memory vecsData = vm.parseJson(jsonData);
        TestVector[] memory vecs = abi.decode(vecsData, (TestVector[]));
        console.log(vecs[0].namespace);
        Namespace memory namespace = toNamespace(bytes29(fromHex(vecs[0].namespace)));
        console.log(bytesToHexString(abi.encodePacked(namespace.toBytes())));
        bytes memory data = fromHex(vecs[0].data);
        bytes memory shares = bytesToShares(data, namespace)[0];
        console.log(bytesToHexString(shares));
    }
}