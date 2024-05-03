// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";

import "../../../Constants.sol";
import "../../../DataRootTuple.sol";
import "../DAVerifier.sol";
import "../../../Blobstream.sol";
import "../../tree/binary/BinaryMerkleProof.sol";
import "../../tree/namespace/NamespaceMerkleMultiproof.sol";
import "../../tree/Types.sol";

import "ds-test/test.sol";

interface CheatCodes {
    function addr(uint256 privateKey) external returns (address);

    function sign(uint256 privateKey, bytes32 digest) external returns (uint8 v, bytes32 r, bytes32 s);
}

/*
The data used to generate the proof:

The block used contains a single share:
0x0000000000000000000000000000000000000000000000000000000001010000014500000026c3020a95010a92010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412720a2f63656c657374696131746b376c776a77336676616578657770687237687833333472766b67646b736d636537666b66122f63656c65737469613167616b61646d63386a73667873646c676e6d64643867773736346739796165776e32726d386d1a0e0a0475746961120631303030303012670a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2103f3e16481ff7c9c2a677f08a30a887e5f9c14313cb624b8c5f7f955d143c81d9212040a020801180112130a0d0a04757469611205323230303010d0e80c1a4068f074601f1bb923f6d6e69d2e3fc3af145c9252eceeb0ac4fba9f661ca0428326f0080478cc969129c0074c3d97ae925de34c5f9d98a458cd47a565a2bb08cc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000

The extended block is:
0x0000000000000000000000000000000000000000000000000000000001010000014500000026c3020a95010a92010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412720a2f63656c657374696131746b376c776a77336676616578657770687237687833333472766b67646b736d636537666b66122f63656c65737469613167616b61646d63386a73667873646c676e6d64643867773736346739796165776e32726d386d1a0e0a0475746961120631303030303012670a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2103f3e16481ff7c9c2a677f08a30a887e5f9c14313cb624b8c5f7f955d143c81d9212040a020801180112130a0d0a04757469611205323230303010d0e80c1a4068f074601f1bb923f6d6e69d2e3fc3af145c9252eceeb0ac4fba9f661ca0428326f0080478cc969129c0074c3d97ae925de34c5f9d98a458cd47a565a2bb08cc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
0x0000000000000000000000000000000000000000000000000000000001010000014500000026c3020a95010a92010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412720a2f63656c657374696131746b376c776a77336676616578657770687237687833333472766b67646b736d636537666b66122f63656c65737469613167616b61646d63386a73667873646c676e6d64643867773736346739796165776e32726d386d1a0e0a0475746961120631303030303012670a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2103f3e16481ff7c9c2a677f08a30a887e5f9c14313cb624b8c5f7f955d143c81d9212040a020801180112130a0d0a04757469611205323230303010d0e80c1a4068f074601f1bb923f6d6e69d2e3fc3af145c9252eceeb0ac4fba9f661ca0428326f0080478cc969129c0074c3d97ae925de34c5f9d98a458cd47a565a2bb08cc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
0x0000000000000000000000000000000000000000000000000000000001010000014500000026c3020a95010a92010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412720a2f63656c657374696131746b376c776a77336676616578657770687237687833333472766b67646b736d636537666b66122f63656c65737469613167616b61646d63386a73667873646c676e6d64643867773736346739796165776e32726d386d1a0e0a0475746961120631303030303012670a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2103f3e16481ff7c9c2a677f08a30a887e5f9c14313cb624b8c5f7f955d143c81d9212040a020801180112130a0d0a04757469611205323230303010d0e80c1a4068f074601f1bb923f6d6e69d2e3fc3af145c9252eceeb0ac4fba9f661ca0428326f0080478cc969129c0074c3d97ae925de34c5f9d98a458cd47a565a2bb08cc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
0x0000000000000000000000000000000000000000000000000000000001010000014500000026c3020a95010a92010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412720a2f63656c657374696131746b376c776a77336676616578657770687237687833333472766b67646b736d636537666b66122f63656c65737469613167616b61646d63386a73667873646c676e6d64643867773736346739796165776e32726d386d1a0e0a0475746961120631303030303012670a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2103f3e16481ff7c9c2a677f08a30a887e5f9c14313cb624b8c5f7f955d143c81d9212040a020801180112130a0d0a04757469611205323230303010d0e80c1a4068f074601f1bb923f6d6e69d2e3fc3af145c9252eceeb0ac4fba9f661ca0428326f0080478cc969129c0074c3d97ae925de34c5f9d98a458cd47a565a2bb08cc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000

The row roots:
0x00000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000001787bf77b567506b6e1d0048bfd89edd352a4fbc102e62f07cc9fe6b4cbe5ee69
0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7329c7d336d0140840837fc0d8eafa2403f4f6b019b602581cd9f04e28026eae

The column roots:
0x00000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000001787bf77b567506b6e1d0048bfd89edd352a4fbc102e62f07cc9fe6b4cbe5ee69
0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7329c7d336d0140840837fc0d8eafa2403f4f6b019b602581cd9f04e28026eae

The data root: 0x55cfc29fc0cd263906122d5cb859091224495b141fc0c51529612d7ab8962950

The height: 3

The blocks data roots used to create the commitment:
1. 0x3d96b7d238e7e0456f6af8e7cdf0a67bd6cf9c2089ecb559c659dcaa1f880353
2. 0x3d96b7d238e7e0456f6af8e7cdf0a67bd6cf9c2089ecb559c659dcaa1f880353
3. 0x55cfc29fc0cd263906122d5cb859091224495b141fc0c51529612d7ab8962950
4. 0x3d96b7d238e7e0456f6af8e7cdf0a67bd6cf9c2089ecb559c659dcaa1f880353

The nonce: 2

The data root tuple root: 0xf89859a09c0f2b1bbb039618d0fe60432b8c247f7ccde97814655f2acffb3434
*/

contract DAVerifierTest is DSTest {
    // Private keys used for test signatures.
    uint256 constant testPriv1 = 0x64a1d6f0e760a8d62b4afdde4096f16f51b401eaaecc915740f71770ea76a8ad;

    Blobstream bridge;
    TestFixture fixture;

    Validator[] private validators;
    uint256 private votingPower = 5000;

    // Set up Foundry cheatcodes.
    CheatCodes cheats = CheatCodes(HEVM_ADDRESS);

    // deploy a Blobstream contract and submit the following:
    // - initial valset.
    // - data root tuple root that commits to the proofs tested below.
    function setUp() public {
        fixture = new TestFixture();

        uint256 initialValsetNonce = 1;

        validators.push(Validator(cheats.addr(testPriv1), votingPower));
        bytes32 hash = computeValidatorSetHash(validators);
        bytes32 checkpoint = domainSeparateValidatorSetHash(initialValsetNonce, (2 * votingPower) / 3, hash);
        bridge = new Blobstream();
        bridge.initialize(initialValsetNonce, (2 * votingPower) / 3, checkpoint);

        bytes32 newDataRootTupleRoot =
            domainSeparateDataRootTupleRoot(fixture.dataRootTupleRootNonce(), fixture.dataRootTupleRoot());

        // Signature for the update.
        Signature[] memory sigs = new Signature[](1);
        bytes32 digest_eip191 = ECDSA.toEthSignedMessageHash(newDataRootTupleRoot);
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(testPriv1, digest_eip191);
        sigs[0] = Signature(v, r, s);

        Validator[] memory valSet = new Validator[](1);
        valSet[0] = Validator(cheats.addr(testPriv1), votingPower);

        bridge.submitDataRootTupleRoot(
            fixture.dataRootTupleRootNonce(), initialValsetNonce, fixture.dataRootTupleRoot(), valSet, sigs
        );

        assertEq(bridge.state_eventNonce(), fixture.dataRootTupleRootNonce());
        assertEq(bridge.state_dataRootTupleRoots(fixture.dataRootTupleRootNonce()), fixture.dataRootTupleRoot());

        assertTrue(
            bridge.verifyAttestation(
                fixture.dataRootTupleRootNonce(), fixture.getDataRootTuple(), fixture.getDataRootTupleProof()
            )
        );
    }

    function testVerifySharesToDataRootTupleRoot() public {
        bytes[] memory _data = new bytes[](1);
        _data[0] = fixture.shareData();

        NamespaceMerkleMultiproof[] memory _shareProofs = new NamespaceMerkleMultiproof[](1);
        _shareProofs[0] = fixture.getShareToRowRootProof();

        NamespaceNode[] memory _rowRoots = new NamespaceNode[](1);
        _rowRoots[0] = fixture.getFirstRowRootNode();

        BinaryMerkleProof[] memory _rowProofs = new BinaryMerkleProof[](1);
        _rowProofs[0] = fixture.getRowRootToDataRootProof();

        AttestationProof memory attestationProof = AttestationProof(
            fixture.dataRootTupleRootNonce(), fixture.getDataRootTuple(), fixture.getDataRootTupleProof()
        );
        SharesProof memory sharesProof =
            SharesProof(_data, _shareProofs, fixture.getNamespace(), _rowRoots, _rowProofs, attestationProof);

        (bool valid, DAVerifier.ErrorCodes errorCode) = DAVerifier.verifySharesToDataRootTupleRoot(bridge, sharesProof);
        assertTrue(valid);
        assertEq(uint8(errorCode), uint8(DAVerifier.ErrorCodes.NoError));
    }

    function testVerifyRowRootToDataRootTupleRoot() public {
        AttestationProof memory attestationProof = AttestationProof(
            fixture.dataRootTupleRootNonce(), fixture.getDataRootTuple(), fixture.getDataRootTupleProof()
        );

        (bool valid, DAVerifier.ErrorCodes errorCode) = DAVerifier.verifyRowRootToDataRootTupleRoot(
            bridge, fixture.getFirstRowRootNode(), fixture.getRowRootToDataRootProof(), attestationProof
        );
        assertTrue(valid);
        assertEq(uint8(errorCode), uint8(DAVerifier.ErrorCodes.NoError));
    }

    function testVerifyMultiRowRootsToDataRootTupleRoot() public {
        NamespaceNode[] memory _rowRoots = new NamespaceNode[](1);
        _rowRoots[0] = fixture.getFirstRowRootNode();

        BinaryMerkleProof[] memory _rowProofs = new BinaryMerkleProof[](1);
        _rowProofs[0] = fixture.getRowRootToDataRootProof();

        AttestationProof memory attestationProof = AttestationProof(
            fixture.dataRootTupleRootNonce(), fixture.getDataRootTuple(), fixture.getDataRootTupleProof()
        );

        (bool valid, DAVerifier.ErrorCodes errorCode) =
            DAVerifier.verifyMultiRowRootsToDataRootTupleRoot(bridge, _rowRoots, _rowProofs, attestationProof);
        assertTrue(valid);
        assertEq(uint8(errorCode), uint8(DAVerifier.ErrorCodes.NoError));
    }

    function testComputeSquareSizeFromRowProof() public {
        (bool validMerkleProof, BinaryMerkleTree.ErrorCodes error) =
            BinaryMerkleTree.verify(fixture.dataRoot(), fixture.getRowRootToDataRootProof(), fixture.firstRowRoot());
        assertEq(uint256(error), uint256(BinaryMerkleTree.ErrorCodes.NoError));
        assertTrue(validMerkleProof);

        // check that the computed square size is correct
        uint256 expectedSquareSize = 1;
        (uint256 actualSquareSize, DAVerifier.ErrorCodes errorCode) =
            DAVerifier.computeSquareSizeFromRowProof(fixture.getRowRootToDataRootProof());
        assertEq(actualSquareSize, expectedSquareSize);
        assertEq(uint8(errorCode), uint8(DAVerifier.ErrorCodes.NoError));
    }

    function testComputeSquareSizeFromShareProof() public {
        bytes[] memory _data = new bytes[](1);
        _data[0] = fixture.shareData();

        // check that the merkle proof is valid
        bool validMerkleProof = NamespaceMerkleTree.verifyMulti(
            fixture.getFirstRowRootNode(), fixture.getShareToRowRootProof(), fixture.getNamespace(), _data
        );
        assertTrue(validMerkleProof);

        // check that the computed square size is correct
        uint256 expectedSquareSize = 1;
        uint256 actualSquareSize = DAVerifier.computeSquareSizeFromShareProof(fixture.getShareToRowRootProof());
        assertEq(actualSquareSize, expectedSquareSize);
    }

    function testValidSlice() public {
        bytes[] memory data = new bytes[](4);
        data[0] = "a";
        data[1] = "b";
        data[2] = "c";
        data[3] = "d";

        (bytes[] memory result, DAVerifier.ErrorCodes error) = DAVerifier.slice(data, 1, 3);

        assertEq(uint256(error), uint256(DAVerifier.ErrorCodes.NoError));
        assertEq(string(result[0]), string(data[1]));
        assertEq(string(result[1]), string(data[2]));
    }

    function testInvalidSliceBeginEnd() public {
        bytes[] memory data = new bytes[](4);
        data[0] = "a";
        data[1] = "b";
        data[2] = "c";
        data[3] = "d";

        (bytes[] memory result, DAVerifier.ErrorCodes error) = DAVerifier.slice(data, 2, 1);

        assertEq(uint256(error), uint256(DAVerifier.ErrorCodes.InvalidRange));
    }

    function testOutOfBoundsSlice() public {
        bytes[] memory data = new bytes[](4);
        data[0] = "a";
        data[1] = "b";
        data[2] = "c";
        data[3] = "d";

        (bytes[] memory result, DAVerifier.ErrorCodes error) = DAVerifier.slice(data, 2, 5);
        assertEq(uint256(error), uint256(DAVerifier.ErrorCodes.OutOfBoundsRange));

        (result, error) = DAVerifier.slice(data, 6, 8);
        assertEq(uint256(error), uint256(DAVerifier.ErrorCodes.OutOfBoundsRange));
    }

    function computeValidatorSetHash(Validator[] memory _validators) private pure returns (bytes32) {
        return keccak256(abi.encode(_validators));
    }

    function domainSeparateValidatorSetHash(uint256 _nonce, uint256 _powerThreshold, bytes32 _validatorSetHash)
        private
        pure
        returns (bytes32)
    {
        bytes32 c =
            keccak256(abi.encode(VALIDATOR_SET_HASH_DOMAIN_SEPARATOR, _nonce, _powerThreshold, _validatorSetHash));

        return c;
    }

    function domainSeparateDataRootTupleRoot(uint256 _nonce, bytes32 _dataRootTupleRoot)
        private
        pure
        returns (bytes32)
    {
        bytes32 c = keccak256(abi.encode(DATA_ROOT_TUPLE_ROOT_DOMAIN_SEPARATOR, _nonce, _dataRootTupleRoot));

        return c;
    }
}

/// @title TestFixture contains the necessary information to create proofs for the token
/// transfer transaction that happened on Celestia. It represents the data mentioned in
/// the comment at the beginning of this file.
contract TestFixture {
    /// @notice the share containing the token transfer transaction on Celestia.
    bytes public shareData = abi.encodePacked(
        hex"0000000000000000000000000000000000000000000000000000000001010000",
        hex"014500000026c3020a95010a92010a1c2f636f736d6f732e62616e6b2e763162",
        hex"657461312e4d736753656e6412720a2f63656c657374696131746b376c776a77",
        hex"336676616578657770687237687833333472766b67646b736d636537666b6612",
        hex"2f63656c65737469613167616b61646d63386a73667873646c676e6d64643867",
        hex"773736346739796165776e32726d386d1a0e0a04757469611206313030303030",
        hex"12670a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b31",
        hex"2e5075624b657912230a2103f3e16481ff7c9c2a677f08a30a887e5f9c14313c",
        hex"b624b8c5f7f955d143c81d9212040a020801180112130a0d0a04757469611205",
        hex"323230303010d0e80c1a4068f074601f1bb923f6d6e69d2e3fc3af145c9252ec",
        hex"eeb0ac4fba9f661ca0428326f0080478cc969129c0074c3d97ae925de34c5f9d",
        hex"98a458cd47a565a2bb08cc000000000000000000000000000000000000000000",
        hex"0000000000000000000000000000000000000000000000000000000000000000",
        hex"0000000000000000000000000000000000000000000000000000000000000000",
        hex"0000000000000000000000000000000000000000000000000000000000000000",
        hex"0000000000000000000000000000000000000000000000000000000000000000"
    );

    /// @notice the first EDS row root.
    bytes public firstRowRoot = abi.encodePacked(
        hex"00000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000001787bf77b567506b6e1d0048bfd89edd352a4fbc102e62f07cc9fe6b4cbe5ee69"
    );

    /// @notice the second EDS row root.
    bytes public secondRowRoot = abi.encodePacked(
        hex"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7329c7d336d0140840837fc0d8eafa2403f4f6b019b602581cd9f04e28026eae"
    );

    /// @notice the first EDS column root.
    bytes public firstColumnRoot = abi.encodePacked(
        hex"00000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000001787bf77b567506b6e1d0048bfd89edd352a4fbc102e62f07cc9fe6b4cbe5ee69"
    );

    /// @notice the second EDS column root.
    bytes public secondColumnRoot = abi.encodePacked(
        hex"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7329c7d336d0140840837fc0d8eafa2403f4f6b019b602581cd9f04e28026eae"
    );

    /// @notice the data root of the block containing the token transfer transaction.
    bytes32 public dataRoot = 0x55cfc29fc0cd263906122d5cb859091224495b141fc0c51529612d7ab8962950;

    /// @notice the height of the block containing the submitted token transfer transaction.
    uint256 public height = 3;

    /// @notice the data root tuple root committing to the Celestia block.
    bytes32 public dataRootTupleRoot = 0xf89859a09c0f2b1bbb039618d0fe60432b8c247f7ccde97814655f2acffb3434;

    /// @notice the data root tuple root nonce in the Blobstream contract.
    uint256 public dataRootTupleRootNonce = 2;

    /// @notice the data root tuple to data root tuple root proof side nodes.
    bytes32[] public dataRootProofSideNodes = [
        bytes32(0xb5d4d27ec6b206a205bf09dde3371ffba62e5b53d27bbec4255b7f4f27ef5d90),
        bytes32(0x406e22ba94989ca721453057a1391fc531edb342c86a0ab4cc722276b54036ec)
    ];

    /// @notice shares to data root proof side nodes.
    NamespaceNode[] public shareToDataRootProofSideNodes = [
        NamespaceNode(
            Namespace(0xff, 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff),
            Namespace(0xff, 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff),
            0x0ec8148c743a4a4db384f40f487cae2fd1ca0d18442d1f162916bdf1cc61b679
        )
    ];

    /// @notice row root to data root proof side nodes.
    bytes32[] public rowRootToDataRootProofSideNodes = [
        bytes32(0x5bc0cf3322dd5c9141a2dcd76947882351690c9aec61015802efc6742992643f),
        bytes32(0xff576381b02abadc50e414f6b4efcae31091cd40a5aba75f56be52d1bb2efcae)
    ];

    /// @notice the share's namespace.
    function getNamespace() public pure returns (Namespace memory) {
        return Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000001);
    }

    /// @notice the data root tuple of the block containing the token transfer transaction.
    function getDataRootTuple() public view returns (DataRootTuple memory) {
        return DataRootTuple(height, dataRoot);
    }

    /// @notice the data root tuple to data root tuple root proof.
    function getDataRootTupleProof() public view returns (BinaryMerkleProof memory) {
        return BinaryMerkleProof(dataRootProofSideNodes, 2, 4);
    }

    /// @notice the first EDS row root.
    function getFirstRowRootNode() public pure returns (NamespaceNode memory) {
        return NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000001),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000001),
            0x787bf77b567506b6e1d0048bfd89edd352a4fbc102e62f07cc9fe6b4cbe5ee69
        );
    }

    /// @notice the second EDS row root.
    function getSecondRowRootNode() public pure returns (NamespaceNode memory) {
        return NamespaceNode(
            Namespace(0xff, 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff),
            Namespace(0xff, 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff),
            0x7329c7d336d0140840837fc0d8eafa2403f4f6b019b602581cd9f04e28026eae
        );
    }

    /// @notice the first EDS column root.
    function getFirstColumnRootNode() public pure returns (NamespaceNode memory) {
        return NamespaceNode(
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000001),
            Namespace(0x00, 0x00000000000000000000000000000000000000000000000000000001),
            0x787bf77b567506b6e1d0048bfd89edd352a4fbc102e62f07cc9fe6b4cbe5ee69
        );
    }

    /// @notice the second EDS column root.
    function getSecondColumnRootNode() public pure returns (NamespaceNode memory) {
        return NamespaceNode(
            Namespace(0xff, 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff),
            Namespace(0xff, 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff),
            0x7329c7d336d0140840837fc0d8eafa2403f4f6b019b602581cd9f04e28026eae
        );
    }

    /// @notice shares to row root proof.
    function getShareToRowRootProof() public view returns (NamespaceMerkleMultiproof memory) {
        return NamespaceMerkleMultiproof(0, 1, shareToDataRootProofSideNodes);
    }

    /// @notice row root to data root proof.
    function getRowRootToDataRootProof() public view returns (BinaryMerkleProof memory) {
        return BinaryMerkleProof(rowRootToDataRootProofSideNodes, 0, 4);
    }
}
