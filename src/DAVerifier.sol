// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import "./lib/openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "./Constants.sol";
import "./DataRootTuple.sol";
import "./QuantumGravityBridge.sol";
import "./lib/tree/binary/BinaryMerkleProof.sol";
import "./lib/tree/namespace/NamespaceMerkleTree.sol";

/// @notice Contains the necessary parameters to prove that some shares, which were posted to
/// the Celestia network, were committed to by the QGB smart contract.
struct SharesProof {
    // The shares that were committed to.
    bytes[] data;
    // The shares proof to the row roots. If the shares span multiple rows, we will have multiple nmt proofs.
    NamespaceMerkleMultiproof[] shareProofs;
    // The namespace ID of the shares.
    bytes8 namespaceID;
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
contract DAVerifier {
    /////////////
    // Storage //
    /////////////

    /// @notice the QGB bridge contract that the proofs will be verified against.
    QuantumGravityBridge public bridge;

    ////////////
    // Errors //
    ////////////

    /// @notice The shares to the rows proof is invalid.
    /// @param i The index of the invalid proof.
    error InvalidSharesToRowsProof(uint256 i);

    /// @notice The rows to the data root proof is invalid.
    /// @param i The index of the invalid proof.
    error InvalidRowsToDataRootProof(uint256 i);

    /// @notice The data root tuple to the data root tuple roof proof is invalid.
    error InvalidDataRootTupleToDataRootTupleRootProof();

    /// @notice The number of share proofs isn't equal to the number of rows roots.
    error UnequalShareProofsAndRowsRootsNumber();

    /// @notice The number of rows proofs isn't equal to the number of rows roots.
    error UnequalRowsProofsAndRowsRootsNumber();

    /// @notice The verifier data length isn't equal to the number of shares in the shares proofs.
    error UnequalDataLengthAndNumberOfSharesProofs();

    ///////////////
    // Functions //
    ///////////////

    /// @param _bridge The QGB smart contract address that the proofs will be verified against.
    constructor(address _bridge) {
        bridge = QuantumGravityBridge(_bridge);
    }

    /// @notice Verifies that the shares, which were posted to Celestia, were committed to by the QGB smart contract.
    /// @param _sharesProof The proof of the shares to the data root tuple root.
    /// @param _root The data root of the block that contains the shares.
    /// @return `true` if the proof is valid, `false` otherwise.
    function verify(SharesProof memory _sharesProof, bytes32 _root) external view returns (bool) {
        // checking that the data root was committed to by the QGB smart contract
        if (
            !bridge.verifyAttestation(
                _sharesProof.attestationProof.tupleRootNonce,
                _sharesProof.attestationProof.tuple,
                _sharesProof.attestationProof.proof
            )
        ) {
            revert InvalidDataRootTupleToDataRootTupleRootProof();
        }

        // checking that the rows roots commit to the data root.
        if (_sharesProof.rowsProofs.length != _sharesProof.rowsRoots.length) {
            revert UnequalRowsProofsAndRowsRootsNumber();
        }

        for (uint256 i = 0; i < _sharesProof.rowsProofs.length; i++) {
            bytes memory rowRoot = abi.encodePacked(
                _sharesProof.rowsRoots[i].min, _sharesProof.rowsRoots[i].max, _sharesProof.rowsRoots[i].digest
            );
            if (!BinaryMerkleTree.verify(_root, _sharesProof.rowsProofs[i], rowRoot)) {
                revert InvalidRowsToDataRootProof(i);
            }
        }

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
