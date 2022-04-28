// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import "./lib/openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "./Constants.sol";
import "./DataRootTuple.sol";
import "./IDAOracle.sol";
import "./lib/tree/binary/BinaryMerkleProof.sol";
import "./lib/tree/binary/BinaryMerkleTree.sol";

struct Validator {
    address addr;
    uint256 power;
}

struct Signature {
    uint8 v;
    bytes32 r;
    bytes32 s;
}

/// 量子引力桥：Celestia -> EVM，数据可用性中继。
/// 中继依靠一组签名者来证明 Celestia 上的某些事件。 这些签名者是 Celestia 验证者集，他们签署每个 Celestia 区块。
/// 跟踪 Celestia 验证器集是通过使用 `updateValidatorSet()` 更新此合约的验证器集视图来完成的。
/// 验证者集当前视图的至少 2/3 的投票权必须签署新的中继事件，并使用 `submitDataRootTupleRoot()` 提交。
/// 每个事件都是一批 `DataRootTuple`（参见 ./DataRootTuple.sol），每个元组代表 Celestia 块头中的单个数据根。 中继元组与块头的顺序相同。

contract QuantumGravityBridge is IDAOracle {
    // 不要更改工作升级的状态顺序并注意继承变量！
    // 继承的合约包含存储槽，并且必须在任何升级中加以考虑。 在主网升级之前，始终在测试网和本地主机上测试精确的升级。

    ////////////////
    // Immutables //可变变量 BRIDGE_ID
    ////////////////

    bytes32 public immutable BRIDGE_ID;

    /////////////
    // Storage //存储
    /////////////

    ///对最新区块的验证
    bytes32 public state_lastValidatorSetCheckpoint;
    ///2/3验证通过才能更新区块
    uint256 public state_powerThreshold;
    ///验证者的唯一随机数
    uint256 public state_lastValidatorSetNonce;
    ///数据根元组根更新的唯一随机数。
    uint256 public state_lastDataRootTupleRootNonce;
    /// 数据根随机数到数据根的映射
    mapping(uint256 => bytes32) public state_dataRootTupleRoots;

    ////////////
    // Events //
    ////////////

    /// 当中继数据根元组的新根时发出。
    /// dataRootTupleRoot 中继数据根元组的默克尔根。
    event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot);

    /// @notice 新验证者进入时执行的事件
    /// @param nonce Nonce.
    /// @param powerThreshold 新的投票门槛（2/3）
    /// @param validatorSetHash ///新验证器集的哈希。
    /// See `updateValidatorSet`.
    event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash);

    ////////////
    // Errors //定义特殊要求
    ////////////

    /// 格式错误
    error MalformedCurrentValidatorSet();

    ///验证者签名不匹配
    error InvalidSignature();

    /// 提交的验证者集权利不足
    error InsufficientVotingPower();

    /// 新的验证器集随机数必须大于当前随机数
    error InvalidValidatorSetNonce();

    /// 提供的当前验证器和权力与检查点不匹配（需要深入研究）
    error SuppliedValidatorSetInvalid();

    /// 数据根元组根随机数必须大于当前随机数。
    error InvalidDataRootTupleRootNonce();

    ///////////////
    // Functions //函数
    ///////////////

    /// @param _bridge_id 桥的标识符，用于域分离的签名。
    /// @param _nonce Celestia 初始化桥的高度。
    /// @param _powerThreshold 批准操作所需的初始投票权。
    /// @param _validatorSetHash 初始验证器集哈希。 这不需要是桥接链的创世验证者集，只需桥的初始验证者集。

    constructor(
        bytes32 _bridge_id,
        uint256 _nonce,
        uint256 _powerThreshold,
        bytes32 _validatorSetHash
    ) {
        BRIDGE_ID = _bridge_id;

        // CHECKS

        bytes32 newCheckpoint = domainSeparateValidatorSetHash(_bridge_id, _nonce, _powerThreshold, _validatorSetHash);

        // EFFECTS

        state_lastValidatorSetNonce = _nonce;///验证者状态
        state_lastValidatorSetCheckpoint = newCheckpoint;///检查数据可用
        state_powerThreshold = _powerThreshold;///检查数据投票权

        // LOGS

        emit ValidatorSetUpdatedEvent(_nonce, _powerThreshold, _validatorSetHash);///连接到以太坊，输入三个验证结果，确定数据的可用性
    }

    /// 检查签名是否为 nil 的实用函数。如果 65 字节签名的所有字节都为零，则它是 nil 签名。

    function isSigNil(Signature calldata _sig) private pure returns (bool) {
        return (_sig.r == 0 && _sig.s == 0 && _sig.v == 0);
    }

    ////// @notice 用于验证 EIP-191 签名的实用程序函数。
    function verifySig(
        address _signer,
        bytes32 _digest,        ///哈希值
        Signature calldata _sig
    ) private pure returns (bool) {
        bytes32 digest_eip191 = ECDSA.toEthSignedMessageHash(_digest);

        return _signer == ECDSA.recover(digest_eip191, _sig.v, _sig.r, _sig.s); ///检查所有数据是否符合
    }

    /// 计算验证器集的哈希。 @param _validators 设置为散列的验证器。
    
    function computeValidatorSetHash(Validator[] calldata _validators) private pure returns (bytes32) {
        return keccak256(abi.encode(_validators));
    }

    /// 将验证者分开的实现
    /// 关于验证者的哈希值：
    /// 格式:
    ///     keccak256(bridge_id, VALIDATOR_SET_HASH_DOMAIN_SEPARATOR, nonce, power_threshold, validator_set_hash)
    /// 验证器集合中的元素应该按幂单调递减。（没看明白）
    function domainSeparateValidatorSetHash(
        bytes32 _bridge_id,
        uint256 _nonce,
        uint256 _powerThreshold,
        bytes32 _validatorSetHash
    ) private pure returns (bytes32) {
        bytes32 c = keccak256(
            abi.encode(_bridge_id, VALIDATOR_SET_HASH_DOMAIN_SEPARATOR, _nonce, _powerThreshold, _validatorSetHash)
        );

        return c;
    }


    /// 对数据根元组根进行域分隔的承诺。
    /// 有关数据根元组根的所有相关信息的哈希。
    /// 格式：
    /// keccak256(bridge_id, DATA_ROOT_TUPLE_ROOT_DOMAIN_SEPARATOR, oldNonce, newNonce, dataRootTupleRoot)
    /// @param _bridge_id Bridge ID.
    /// @param _oldNonce 开始的 Celestia 区块高度。
    /// @param _newNonce 结束的 Celestia 区块高度。
    /// @param _dataRootTupleRoot 数据根
    function domainSeparateDataRootTupleRoot( ///区分不同的数据根
        bytes32 _bridge_id,
        uint256 _oldNonce,
        uint256 _newNonce,
        bytes32 _dataRootTupleRoot
    ) private pure returns (bytes32) {
        bytes32 c = keccak256(
            abi.encode(_bridge_id, DATA_ROOT_TUPLE_ROOT_DOMAIN_SEPARATOR, _oldNonce, _newNonce, _dataRootTupleRoot)
        );

        return c;
    }

////////////////验证环节：

    /// @dev Checks that enough voting power signed over a digest.
    /// @param _currentValidators The current validators.
    /// @param _sigs The current validators' signatures.
    /// @param _digest This is what we are checking they have signed.
    /// @param _powerThreshold At least this much power must have signed.
    function checkValidatorSignatures(
        // The current validator set and their powers
        Validator[] calldata _currentValidators,
        Signature[] calldata _sigs,
        bytes32 _digest,
        uint256 _powerThreshold
    ) private pure {
        uint256 cumulativePower = 0;

        for (uint256 i = 0; i < _currentValidators.length; i++) {
            // If the signature is nil, then it's not present so continue.
            if (isSigNil(_sigs[i])) {
                continue;
            }

            // Check that the current validator has signed off on the hash.
            if (!verifySig(_currentValidators[i].addr, _digest, _sigs[i])) {
                revert InvalidSignature();
            }

            // Sum up cumulative power.
            cumulativePower += _currentValidators[i].power;

            // Break early to avoid wasting gas.
            if (cumulativePower >= _powerThreshold) {
                break;         ///////停止循环，验证通过
            }
        }

        // Check that there was enough power.
        if (cumulativePower < _powerThreshold) {
            revert InsufficientVotingPower();
        }
    }

    /// @notice 这通过检查验证器来更新验证器集
     /// 在当前验证者集中已经签署了新的验证者集。
     /// 提供的签名是当前验证器集在新验证器集生成的检查点哈希上的签名。 任何人都可以调用此函数，但他们必须提供当前验证者集的有效签名，而不是新的验证者集。
    ///
    /// 签名的验证器集哈希是域分隔的 `domainSeparateValidatorSetHash`.
    /// @param _newValidatorSetHash The hash of the new validator set.
    /// @param _newNonce The new Celestia block height.
    /// @param _currentValidatorSet The current validator set.
    /// @param _sigs Signatures.
    function updateValidatorSet(
        uint256 _newNonce,
        uint256 _newPowerThreshold,
        bytes32 _newValidatorSetHash,
        Validator[] calldata _currentValidatorSet,
        Signature[] calldata _sigs
    ) external {
        // CHECKS

        uint256 currentNonce = state_lastValidatorSetNonce;
        uint256 currentPowerThreshold = state_powerThreshold;

        // Check that the new validator set nonce is greater than the old one.
        if (_newNonce <= currentNonce) {
            revert InvalidValidatorSetNonce();
        }

        // Check that current validators and signatures are well-formed.
        if (_currentValidatorSet.length != _sigs.length) {
            revert MalformedCurrentValidatorSet();
        }

        // Check that the supplied current validator set matches the saved checkpoint.
        bytes32 currentValidatorSetHash = computeValidatorSetHash(_currentValidatorSet);
        if (
            domainSeparateValidatorSetHash(BRIDGE_ID, currentNonce, currentPowerThreshold, currentValidatorSetHash) !=
            state_lastValidatorSetCheckpoint
        ) {
            revert SuppliedValidatorSetInvalid();
        }

        // Check that enough current validators have signed off on the new validator set.
        bytes32 newCheckpoint = domainSeparateValidatorSetHash(
            BRIDGE_ID,
            _newNonce,
            _newPowerThreshold,
            _newValidatorSetHash
        );
        checkValidatorSignatures(_currentValidatorSet, _sigs, newCheckpoint, currentPowerThreshold);

        // EFFECTS

        state_lastValidatorSetCheckpoint = newCheckpoint;
        state_powerThreshold = _newPowerThreshold;
        state_lastValidatorSetNonce = _newNonce;

        // LOGS

        emit ValidatorSetUpdatedEvent(_newNonce, _newPowerThreshold, _newValidatorSetHash);
    }

    /// @notice 将 Celestia 数据根元组的根中继到 EVM 链。 任何人都可以调用此函数，但他们必须提供当前验证器集在数据根元组根上的有效签名。
     ///数据根根是二叉Merkle树的Merkle根
     /// (https://github.com/celestiaorg/celestia-specs/blob/master/src/specs/data_structures.md#binary-merkle-tree)
     /// 其中每个叶子都在 ABI 编码的 `DataRootTuple` 中。 每个中继的数据根元组将 1:1 镜像数据根，因为它们包含在 Celestia 的标头中，_按包含顺序_。
     /// 被签名的数据元组根是按照 `domainSeparateDataRootTupleRoot` 进行域分隔的。
     
    /// @param _nonce The Celestia block height up to which the data root tuple
    /// root commits to.
    /// @param _dataRootTupleRoot The Merkle root of data root tuples.
    /// @param _currentValidatorSet The current validator set.
    /// @param _sigs Signatures.
    function submitDataRootTupleRoot(
        uint256 _nonce,
        bytes32 _dataRootTupleRoot,
        Validator[] calldata _currentValidatorSet,
        Signature[] calldata _sigs
    ) external {
        // CHECKS

        uint256 currentNonce = state_lastDataRootTupleRootNonce;
        uint256 currentPowerThreshold = state_powerThreshold;

        // Check that the data root tuple root nonce is higher than the last nonce.
        if (_nonce <= currentNonce) {
            revert InvalidDataRootTupleRootNonce();
        }

        // Check that current validators and signatures are well-formed.
        if (_currentValidatorSet.length != _sigs.length) {
            revert MalformedCurrentValidatorSet();
        }

        // Check that the supplied current validator set matches the saved checkpoint.
        bytes32 currentValidatorSetHash = computeValidatorSetHash(_currentValidatorSet);
        if (
            domainSeparateValidatorSetHash(
                BRIDGE_ID,
                state_lastValidatorSetNonce,
                currentPowerThreshold,
                currentValidatorSetHash
            ) != state_lastValidatorSetCheckpoint
        ) {
            revert SuppliedValidatorSetInvalid();
        }

        // Check that enough current validators have signed off on the data
        // root tuple root and nonce.
        bytes32 c = domainSeparateDataRootTupleRoot(BRIDGE_ID, currentNonce, _nonce, _dataRootTupleRoot);
        checkValidatorSignatures(_currentValidatorSet, _sigs, c, currentPowerThreshold);

        // EFFECTS

        state_lastDataRootTupleRootNonce = _nonce;
        state_dataRootTupleRoots[_nonce] = _dataRootTupleRoot;

        // LOGS

        emit DataRootTupleRootEvent(_nonce, _dataRootTupleRoot);
    }

    /// @dev see "./IDAOracle.sol"
    function verifyAttestation(
        uint256 _tupleRootIndex,
        DataRootTuple memory _tuple,
        BinaryMerkleProof memory _proof
    ) external view override returns (bool) {
        // Tuple must have been committed before.
        if (_tupleRootIndex > state_lastDataRootTupleRootNonce) {
            return false;
        }

        // Load the tuple root at the given index from storage.
        bytes32 root = state_dataRootTupleRoots[_tupleRootIndex];

        // Verify the proof.
        bool isProofValid = BinaryMerkleTree.verify(root, _proof, abi.encode(_tuple));

        return isProofValid;
    }
}
