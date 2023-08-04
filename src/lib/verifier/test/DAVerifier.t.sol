// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.19;

import "openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";

import "../../../Constants.sol";
import "../../../DataRootTuple.sol";
import "../DAVerifier.sol";
import "../../../QuantumGravityBridge.sol";
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
0x0000000000000001010000014300000011c1020a95010a92010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412720a2f63656c65737469613165383064747a75387a38786739676d7333716d346c34336639757a6c306174767473766a3564122f63656c65737469613167616b61646d63386a73667873646c676e6d64643867773736346739796165776e32726d386d1a0e0a0475746961120631303030303012650a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102207a8037a3a1dac112f77d982feaca3d8930e468b835a11ff176a159588334f312040a020801180112110a0b0a0475746961120335303010d0e80c1a40c19753445b3de4d70d6c25707d082968e8fd8c8b8fb4e135a570c8d291e90a7b30219bf5ab4840081c1479d8295a5d73ef1d635faf40467bbe7658398d24f1d600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000

The extended block is:
0x0000000000000001010000014300000011c1020a95010a92010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412720a2f63656c65737469613165383064747a75387a38786739676d7333716d346c34336639757a6c306174767473766a3564122f63656c65737469613167616b61646d63386a73667873646c676e6d64643867773736346739796165776e32726d386d1a0e0a0475746961120631303030303012650a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102207a8037a3a1dac112f77d982feaca3d8930e468b835a11ff176a159588334f312040a020801180112110a0b0a0475746961120335303010d0e80c1a40c19753445b3de4d70d6c25707d082968e8fd8c8b8fb4e135a570c8d291e90a7b30219bf5ab4840081c1479d8295a5d73ef1d635faf40467bbe7658398d24f1d600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
0x0000000000000001010000014300000011c1020a95010a92010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412720a2f63656c65737469613165383064747a75387a38786739676d7333716d346c34336639757a6c306174767473766a3564122f63656c65737469613167616b61646d63386a73667873646c676e6d64643867773736346739796165776e32726d386d1a0e0a0475746961120631303030303012650a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102207a8037a3a1dac112f77d982feaca3d8930e468b835a11ff176a159588334f312040a020801180112110a0b0a0475746961120335303010d0e80c1a40c19753445b3de4d70d6c25707d082968e8fd8c8b8fb4e135a570c8d291e90a7b30219bf5ab4840081c1479d8295a5d73ef1d635faf40467bbe7658398d24f1d600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
0x0000000000000001010000014300000011c1020a95010a92010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412720a2f63656c65737469613165383064747a75387a38786739676d7333716d346c34336639757a6c306174767473766a3564122f63656c65737469613167616b61646d63386a73667873646c676e6d64643867773736346739796165776e32726d386d1a0e0a0475746961120631303030303012650a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102207a8037a3a1dac112f77d982feaca3d8930e468b835a11ff176a159588334f312040a020801180112110a0b0a0475746961120335303010d0e80c1a40c19753445b3de4d70d6c25707d082968e8fd8c8b8fb4e135a570c8d291e90a7b30219bf5ab4840081c1479d8295a5d73ef1d635faf40467bbe7658398d24f1d600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
0x0000000000000001010000014300000011c1020a95010a92010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412720a2f63656c65737469613165383064747a75387a38786739676d7333716d346c34336639757a6c306174767473766a3564122f63656c65737469613167616b61646d63386a73667873646c676e6d64643867773736346739796165776e32726d386d1a0e0a0475746961120631303030303012650a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102207a8037a3a1dac112f77d982feaca3d8930e468b835a11ff176a159588334f312040a020801180112110a0b0a0475746961120335303010d0e80c1a40c19753445b3de4d70d6c25707d082968e8fd8c8b8fb4e135a570c8d291e90a7b30219bf5ab4840081c1479d8295a5d73ef1d635faf40467bbe7658398d24f1d600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000

The row roots:
0x000000000000000100000000000000018c8732952e0c3e3f0adf0a43665e30bc554cfad53635caccb52c7d38cc078af8
0xffffffffffffffffffffffffffffffff8a27b3798dc3f14c183597cdfa738c94758cbf0665fa96242672d65cf72881a9

The column roots:
0x000000000000000100000000000000018c8732952e0c3e3f0adf0a43665e30bc554cfad53635caccb52c7d38cc078af8
0xffffffffffffffffffffffffffffffff8a27b3798dc3f14c183597cdfa738c94758cbf0665fa96242672d65cf72881a9

The data root: 0x1108C0D8079563116167A66BE596DBE222E438C273ECC3B48E290465FC6093B2

The height: 2

The blocks data roots used to create the commitment:
1. 0x257760461993F8F197B421EC7435F3C36C3734923E3DA9A42DC73B05F07B3D08
2. 0x1108C0D8079563116167A66BE596DBE222E438C273ECC3B48E290465FC6093B2
3. 0x257760461993F8F197B421EC7435F3C36C3734923E3DA9A42DC73B05F07B3D08

The nonce: 2

The data root tuple root: 0x81A5323C06C5CF0EE22752CC01597F16E93A1C6CCA71625AAEE9D918D09345ED
*/

contract DAVerifierTest is DSTest {
    // Private keys used for test signatures.
    uint256 constant testPriv1 = 0x64a1d6f0e760a8d62b4afdde4096f16f51b401eaaecc915740f71770ea76a8ad;

    QuantumGravityBridge bridge;
    TestFixture fixture;

    Validator[] private validators;
    uint256 private votingPower = 5000;

    // Set up Foundry cheatcodes.
    CheatCodes cheats = CheatCodes(HEVM_ADDRESS);

    // deploy a QGB contract and submit the following:
    // - initial valset.
    // - data root tuple root that commits to the proofs tested below.
    function setUp() public {
        fixture = new TestFixture();

        uint256 initialVelsetNonce = 1;

        validators.push(Validator(cheats.addr(testPriv1), votingPower));
        bytes32 hash = computeValidatorSetHash(validators);
        bridge = new QuantumGravityBridge(initialVelsetNonce, (2 * votingPower) / 3, hash);

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
            fixture.dataRootTupleRootNonce(), initialVelsetNonce, fixture.dataRootTupleRoot(), valSet, sigs
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
            SharesProof(_data, _shareProofs, fixture.minimaxNID(), _rowRoots, _rowProofs, attestationProof);

        bool valid = DAVerifier.verifySharesToDataRootTupleRoot(bridge, sharesProof, fixture.dataRoot());
        assertTrue(valid);
    }

    function testVerifyRowRootToDataRootTupleRoot() public {
        AttestationProof memory attestationProof = AttestationProof(
            fixture.dataRootTupleRootNonce(), fixture.getDataRootTuple(), fixture.getDataRootTupleProof()
        );

        bool valid = DAVerifier.verifyRowRootToDataRootTupleRoot(
            bridge,
            fixture.getFirstRowRootNode(),
            fixture.getRowRootToDataRootProof(),
            attestationProof,
            fixture.dataRoot()
        );
        assertTrue(valid);
    }

    function testVerifyMultiRowRootsToDataRootTupleRoot() public {
        NamespaceNode[] memory _rowRoots = new NamespaceNode[](1);
        _rowRoots[0] = fixture.getFirstRowRootNode();

        BinaryMerkleProof[] memory _rowProofs = new BinaryMerkleProof[](1);
        _rowProofs[0] = fixture.getRowRootToDataRootProof();

        AttestationProof memory attestationProof = AttestationProof(
            fixture.dataRootTupleRootNonce(), fixture.getDataRootTuple(), fixture.getDataRootTupleProof()
        );

        bool valid = DAVerifier.verifyMultiRowRootsToDataRootTupleRoot(
            bridge, _rowRoots, _rowProofs, attestationProof, fixture.dataRoot()
        );
        assertTrue(valid);
    }

    function testComputeSquareSizeFromRowProof() public {
        bool validMerkleProof =
            BinaryMerkleTree.verify(fixture.dataRoot(), fixture.getRowRootToDataRootProof(), fixture.firstRowRoot());
        assertTrue(validMerkleProof);

        // check that the computed square size is correct
        uint256 expectedSquareSize = 1;
        uint256 actualSquareSize = DAVerifier.computeSquareSizeFromRowProof(fixture.getRowRootToDataRootProof());
        assertEq(actualSquareSize, expectedSquareSize);
    }

    function testComputeSquareSizeFromShareProof() public {
        bytes[] memory _data = new bytes[](1);
        _data[0] = fixture.shareData();

        // check that the merkle proof is valid
        bool validMerkleProof = NamespaceMerkleTree.verifyMulti(
            fixture.getFirstRowRootNode(), fixture.getShareToRowRootProof(), fixture.minimaxNID(), _data
        );
        assertTrue(validMerkleProof);

        // check that the computed square size is correct
        uint256 expectedSquareSize = 1;
        uint256 actualSquareSize = DAVerifier.computeSquareSizeFromShareProof(fixture.getShareToRowRootProof());
        assertEq(actualSquareSize, expectedSquareSize);
    }

    function computeValidatorSetHash(Validator[] memory _validators) private pure returns (bytes32) {
        return keccak256(abi.encode(_validators));
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

/// @title TestFixture contains the necessary information to create proofs for the blob
/// that was posted to Celestia. It represents the data mentioned in the comment at
/// the beginning of this file.
contract TestFixture {
    /// @notice the share containing the blob that was published to Celestia.
    bytes public shareData = abi.encodePacked(
        hex"0000000000000001010000014300000011c1020a95010a92010a1c2f636f736d",
        hex"6f732e62616e6b2e763162657461312e4d736753656e6412720a2f63656c6573",
        hex"7469613165383064747a75387a38786739676d7333716d346c34336639757a6c",
        hex"306174767473766a3564122f63656c65737469613167616b61646d63386a7366",
        hex"7873646c676e6d64643867773736346739796165776e32726d386d1a0e0a0475",
        hex"746961120631303030303012650a500a460a1f2f636f736d6f732e6372797074",
        hex"6f2e736563703235366b312e5075624b657912230a2102207a8037a3a1dac112",
        hex"f77d982feaca3d8930e468b835a11ff176a159588334f312040a020801180112",
        hex"110a0b0a0475746961120335303010d0e80c1a40c19753445b3de4d70d6c2570",
        hex"7d082968e8fd8c8b8fb4e135a570c8d291e90a7b30219bf5ab4840081c1479d8",
        hex"295a5d73ef1d635faf40467bbe7658398d24f1d6000000000000000000000000",
        hex"0000000000000000000000000000000000000000000000000000000000000000",
        hex"0000000000000000000000000000000000000000000000000000000000000000",
        hex"0000000000000000000000000000000000000000000000000000000000000000",
        hex"0000000000000000000000000000000000000000000000000000000000000000",
        hex"0000000000000000000000000000000000000000000000000000000000000000"
    );

    /// @notice the share's namespace ID.
    NamespaceID public minimaxNID = NamespaceID.wrap(0x0000000000000001);

    /// @notice the first EDS row root.
    bytes public firstRowRoot = abi.encodePacked(
        hex"000000000000000100000000000000018c8732952e0c3e3f0adf0a43665e30bc554cfad53635caccb52c7d38cc078af8"
    );

    /// @notice the second EDS row root.
    bytes public secondRowRoot = abi.encodePacked(
        hex"ffffffffffffffffffffffffffffffff8a27b3798dc3f14c183597cdfa738c94758cbf0665fa96242672d65cf72881a9"
    );

    /// @notice the first EDS column root.
    bytes public firstColumnRoot = abi.encodePacked(
        hex"000000000000000100000000000000018c8732952e0c3e3f0adf0a43665e30bc554cfad53635caccb52c7d38cc078af8"
    );

    /// @notice the second EDS column root.
    bytes public secondColumnRoot = abi.encodePacked(
        hex"ffffffffffffffffffffffffffffffff8a27b3798dc3f14c183597cdfa738c94758cbf0665fa96242672d65cf72881a9"
    );

    /// @notice the data root of the block containing the submitted blob.
    bytes32 public dataRoot = 0x1108C0D8079563116167A66BE596DBE222E438C273ECC3B48E290465FC6093B2;

    /// @notice the height of the block containing the submitted blob.
    uint256 public height = 2;

    /// @notice the data root tuple root committing to the Celestia block.
    bytes32 public dataRootTupleRoot = 0x81A5323C06C5CF0EE22752CC01597F16E93A1C6CCA71625AAEE9D918D09345ED;

    /// @notice the data root tuple root nonce in the QGB contract.
    uint256 public dataRootTupleRootNonce = 2;

    /// @notice the data root tuple to data root tuple root proof side nodes.
    bytes32[] public dataRootProofSideNodes = [
        bytes32(0xD380873912E163B240C72D2AED926CCED511A34467BE9E697F49465A7DF8F3BE),
        bytes32(0x055B7998D838C5846E1751A6C8BA8822459C492549AC7EA33ADDA48E4861C78F)
    ];

    /// @notice shares to data root proof side nodes.
    NamespaceNode[] public shareToDataRootProofSideNodes = [
        NamespaceNode(
            NamespaceID.wrap(0xffffffffffffffff),
            NamespaceID.wrap(0xffffffffffffffff),
            0x99ff60ce3818df2d1601a5a6a7d7bac82aa79d1726bca4e05b94e4ce38f06ffe
        )
    ];

    /// @notice row root to data root proof side nodes.
    bytes32[] public rowRootToDataRootProofSideNodes = [
        bytes32(0x3d9568eda3d860f78af0b605066eb6f90495b87d2081bf875e6a88ede0a1f6bc),
        bytes32(0xbf6d9e948bc2e4bc32a791135385bbc2a3b8f426d33d212c9f755e8f2dd964ad)
    ];

    /// @notice the data root tuple of the block containing the submitted blob.
    function getDataRootTuple() public view returns (DataRootTuple memory) {
        return DataRootTuple(height, dataRoot);
    }

    /// @notice the data root tuple to data root tuple root proof.
    function getDataRootTupleProof() public view returns (BinaryMerkleProof memory) {
        return BinaryMerkleProof(dataRootProofSideNodes, 1, 4);
    }

    /// @notice the first EDS row root.
    function getFirstRowRootNode() public view returns (NamespaceNode memory) {
        return NamespaceNode(
            NamespaceID.wrap(0x0000000000000001),
            NamespaceID.wrap(0x0000000000000001),
            0x8C8732952E0C3E3F0ADF0A43665E30BC554CFAD53635CACCB52C7D38CC078AF8
        );
    }

    /// @notice the second EDS row root.
    function getSecondRowRootNode() public view returns (NamespaceNode memory) {
        return NamespaceNode(
            NamespaceID.wrap(0x0000000000000001),
            NamespaceID.wrap(0xffffffffffffffff),
            0x8C8732952E0C3E3F0ADF0A43665E30BC554CFAD53635CACCB52C7D38CC078AF8
        );
    }

    /// @notice the first EDS column root.
    function getFirstColumnRootNode() public view returns (NamespaceNode memory) {
        return NamespaceNode(
            NamespaceID.wrap(0x0000000000000001),
            NamespaceID.wrap(0xffffffffffffffff),
            0x8C8732952E0C3E3F0ADF0A43665E30BC554CFAD53635CACCB52C7D38CC078AF8
        );
    }

    /// @notice the second EDS column root.
    function getSecondColumnRootNode() public view returns (NamespaceNode memory) {
        return NamespaceNode(
            NamespaceID.wrap(0x0000000000000001),
            NamespaceID.wrap(0xffffffffffffffff),
            0x8C8732952E0C3E3F0ADF0A43665E30BC554CFAD53635CACCB52C7D38CC078AF8
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
