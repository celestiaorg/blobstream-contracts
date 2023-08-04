// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.19;

import "openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";

import "../../Constants.sol";
import "../../DataRootTuple.sol";
import "../../IDAOracle.sol";
import "../tree/binary/BinaryMerkleProof.sol";
import "../tree/binary/BinaryMerkleTree.sol";
import "../tree/namespace/NamespaceMerkleTree.sol";
import "../tree/Types.sol";

/// @notice Contains the necessary parameters to prove that some shares, which were posted to
/// the Celestia network, were committed to by the QGB smart contract.
struct SharesProof {
    // The shares that were committed to.
    bytes[] data;
    // The shares proof to the row roots. If the shares span multiple rows, we will have multiple nmt proofs.
    NamespaceMerkleMultiproof[] shareProofs;
    // The namespace ID of the shares.
    NamespaceID namespaceID;
    // The rows where the shares belong. If the shares span multiple rows, we will have multiple rows.
    NamespaceNode[] rowsRoots;
    // The proofs of the rowsRoots to the data root.
    BinaryMerkleProof[] rowsProofs;
    // The proof of the data root tuple to the data root tuple root that was posted to the QGB contract.
    AttestationProof attestationProof;
}

/// @notice Contains the necessary parameters needed to verify that a data root tuple
/// was committed to, by the QGB smart contract, at some specif nonce.
struct AttestationProof {
    // the attestation nonce that commits to the data root tuple.
    uint256 tupleRootNonce;
    // the data root tuple that was committed to.
    DataRootTuple tuple;
    // the binary merkle proof of the tuple to the commitment.
    BinaryMerkleProof proof;
}

/// @title DAVerifier: Celestia -> EVM, Data Availability verifier.
/// @dev The DAVerifier verifies that some shares, which were posted on Celestia, were committed to
/// by the QGB smart contract.
library DAVerifier {
    ////////////
    // Errors //
    ////////////

    /// @notice The shares to the rows proof is invalid.
    /// @param i The index of the invalid proof.
    error InvalidSharesToRowsProof(uint256 i);

    /// @notice The rows to the data root proof is invalid.
    /// @param i The index of the invalid proof.
    error InvalidRowsToDataRootProof(uint256 i);

    /// @notice The row to the data root proof is invalid.
    error InvalidRowToDataRootProof();

    /// @notice The data root tuple to the data root tuple roof proof is invalid.
    error InvalidDataRootTupleToDataRootTupleRootProof();

    /// @notice The number of share proofs isn't equal to the number of rows roots.
    error UnequalShareProofsAndRowsRootsNumber();

    /// @notice The number of rows proofs isn't equal to the number of rows roots.
    error UnequalRowsProofsAndRowsRootsNumber();

    /// @notice The verifier data length isn't equal to the number of shares in the shares proofs.
    error UnequalDataLengthAndNumberOfSharesProofs();

    /// @notice The number of leaves in the binary merkle proof is not divisible by 4.
    /// @param i The provided proof number of leaves.
    error InvalidNumberOfLeavesInProof(uint256 i);

    ///////////////
    // Functions //
    ///////////////

    /// @notice Verifies that the shares, which were posted to Celestia, were committed to by the QGB smart contract.
    /// @param _bridge The QGB smart contract instance.
    /// @param _sharesProof The proof of the shares to the data root tuple root.
    /// @param _root The data root of the block that contains the shares.
    /// @return `true` if the proof is valid, `false` otherwise.
    function verifySharesToDataRootTupleRoot(IDAOracle _bridge, SharesProof memory _sharesProof, bytes32 _root)
        external
        view
        returns (bool)
    {
        // checking that the data root was committed to by the QGB smart contract
        // this will revert if the proof is not valid
        verifyMultiRowRootsToDataRootTupleRoot(
            _bridge, _sharesProof.rowsRoots, _sharesProof.rowsProofs, _sharesProof.attestationProof, _root
        );

        // checking that the shares were committed to by the rows roots.
        if (_sharesProof.shareProofs.length != _sharesProof.rowsRoots.length) {
            revert UnequalShareProofsAndRowsRootsNumber();
        }

        uint256 numberOfSharesInProofs = 0;
        for (uint256 i = 0; i < _sharesProof.shareProofs.length; i++) {
            numberOfSharesInProofs += _sharesProof.shareProofs[i].endKey - _sharesProof.shareProofs[i].beginKey;
        }

        if (_sharesProof.data.length != numberOfSharesInProofs) {
            revert UnequalDataLengthAndNumberOfSharesProofs();
        }

        uint256 cursor = 0;
        for (uint256 i = 0; i < _sharesProof.shareProofs.length; i++) {
            uint256 sharesUsed = _sharesProof.shareProofs[i].endKey - _sharesProof.shareProofs[i].beginKey;
            NamespaceNode memory rowRoot =
                NamespaceNode(_sharesProof.namespaceID, _sharesProof.namespaceID, _sharesProof.rowsRoots[i].digest);
            if (
                !NamespaceMerkleTree.verifyMulti(
                    rowRoot,
                    _sharesProof.shareProofs[i],
                    _sharesProof.namespaceID,
                    slice(_sharesProof.data, cursor, cursor + sharesUsed)
                )
            ) {
                revert InvalidSharesToRowsProof(i);
            }
            cursor += sharesUsed;
        }

        return true;
    }

    /// @notice Verifies that a row/column root, from a Celestia block, was committed to by the QGB smart contract.
    /// @param _bridge The QGB smart contract instance.
    /// @param _rowRoot The row/column root to be proven.
    /// @param _rowProof The proof of the row/column root to the data root.
    /// @param _root The data root of the block that contains the row.
    /// @return `true` if the proof is valid, `false` otherwise.
    function verifyRowRootToDataRootTupleRoot(
        IDAOracle _bridge,
        NamespaceNode memory _rowRoot,
        BinaryMerkleProof memory _rowProof,
        AttestationProof memory _attestationProof,
        bytes32 _root
    ) public view returns (bool) {
        // checking that the data root was committed to by the QGB smart contract
        if (
            !_bridge.verifyAttestation(
                _attestationProof.tupleRootNonce, _attestationProof.tuple, _attestationProof.proof
            )
        ) {
            revert InvalidDataRootTupleToDataRootTupleRootProof();
        }

        bytes memory rowRoot = abi.encodePacked(_rowRoot.min, _rowRoot.max, _rowRoot.digest);
        if (!BinaryMerkleTree.verify(_root, _rowProof, rowRoot)) {
            revert InvalidRowToDataRootProof();
        }

        return true;
    }

    /// @notice Verifies that a set of rows/columns, from a Celestia block, were committed to by the QGB smart contract.
    /// @param _bridge The QGB smart contract instance.
    /// @param _rowsRoots The set of row/column roots to be proved.
    /// @param _rowsProofs The set of proofs of the _rowsRoots in the same order.
    /// @param _root The data root of the block that contains the rows.
    /// @return `true` if the proof is valid, `false` otherwise.
    function verifyMultiRowRootsToDataRootTupleRoot(
        IDAOracle _bridge,
        NamespaceNode[] memory _rowsRoots,
        BinaryMerkleProof[] memory _rowsProofs,
        AttestationProof memory _attestationProof,
        bytes32 _root
    ) public view returns (bool) {
        // checking that the data root was committed to by the QGB smart contract
        if (
            !_bridge.verifyAttestation(
                _attestationProof.tupleRootNonce, _attestationProof.tuple, _attestationProof.proof
            )
        ) {
            revert InvalidDataRootTupleToDataRootTupleRootProof();
        }

        // checking that the rows roots commit to the data root.
        if (_rowsProofs.length != _rowsRoots.length) {
            revert UnequalRowsProofsAndRowsRootsNumber();
        }

        for (uint256 i = 0; i < _rowsProofs.length; i++) {
            bytes memory rowRoot = abi.encodePacked(_rowsRoots[i].min, _rowsRoots[i].max, _rowsRoots[i].digest);
            if (!BinaryMerkleTree.verify(_root, _rowsProofs[i], rowRoot)) {
                revert InvalidRowsToDataRootProof(i);
            }
        }

        return true;
    }

    /// @notice computes the Celestia block square size from a row/column root to data root binary merkle proof.
    /// Note: the provided proof is not authenticated to the QGB smart contract. It is the user's responsibility
    /// to verify that the proof is valid and was successfully committed to using
    //  the `DAVerifier.verifyRowRootToDataRootTupleRoot()` method
    /// Note: the minimum square size is 1. Thus, we don't expect the proof to have number of leaves equal to 0.
    /// @param _proof The proof of the row/column root to the data root.
    /// @return The square size of the corresponding block.
    function computeSquareSizeFromRowProof(BinaryMerkleProof memory _proof) external pure returns (uint256) {
        if (_proof.numLeaves % 4 != 0) {
            revert InvalidNumberOfLeavesInProof(_proof.numLeaves);
        }
        // we divide the number of leaves of the proof by 4 because the rows/columns tree is constructed
        // from the extended block row roots and column roots.
        return _proof.numLeaves / 4;
    }

    /// @notice computes the Celestia block square size from a shares to row/column root proof.
    /// Note: the provided proof is not authenticated to the QGB smart contract. It is the user's responsibility
    /// to verify that the proof is valid and that the shares were successfully committed to using
    /// the `DAVerifier.verifySharesToDataRootTupleRoot()` method.
    /// Note: the minimum square size is 1. Thus, we don't expect the proof not to contain any side node.
    /// @param _proof The proof of the shares to the row/column root.
    /// @return The square size of the corresponding block.
    function computeSquareSizeFromShareProof(NamespaceMerkleMultiproof memory _proof) external pure returns (uint256) {
        uint256 extendedSquareRowSize = 2 ** _proof.sideNodes.length;
        // we divide the extended square row size by 2 because the square size is the
        // the size of the row of the original square size.
        return extendedSquareRowSize / 2;
    }

    /// @notice creates a slice of bytes from the data slice of bytes containing the elements
    /// that correspond to the provided range.
    /// It selects a half-open range which includes the begin element, but excludes the end one.
    /// @param _data The slice that we want to select data from.
    /// @param _begin The beginning of the range (inclusive).
    /// @param _end The ending of the range (exclusive).
    /// @return _ the sliced data.
    function slice(bytes[] memory _data, uint256 _begin, uint256 _end) internal pure returns (bytes[] memory) {
        bytes[] memory out = new bytes[](_end-_begin);
        for (uint256 i = _begin; i < _end; i++) {
            out[i] = _data[i];
        }
        return out;
    }
}
