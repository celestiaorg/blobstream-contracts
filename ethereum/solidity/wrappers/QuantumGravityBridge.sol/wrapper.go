// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BinaryMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type BinaryMerkleProof struct {
	SideNodes [][32]byte
	Key       *big.Int
	NumLeaves *big.Int
}

// DataRootTuple is an auto generated low-level Go binding around an user-defined struct.
type DataRootTuple struct {
	Height   *big.Int
	DataRoot [32]byte
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Addr  common.Address
	Power *big.Int
}

// BinaryMerkleTreeMetaData contains all meta data concerning the BinaryMerkleTree contract.
var BinaryMerkleTreeMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122080c9e017991ec01c87b659c80c4b851d2bfbc72ff84ee1d8d004ede6351bb63f64736f6c634300080c0033",
}

// BinaryMerkleTreeABI is the input ABI used to generate the binding from.
// Deprecated: Use BinaryMerkleTreeMetaData.ABI instead.
var BinaryMerkleTreeABI = BinaryMerkleTreeMetaData.ABI

// BinaryMerkleTreeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BinaryMerkleTreeMetaData.Bin instead.
var BinaryMerkleTreeBin = BinaryMerkleTreeMetaData.Bin

// DeployBinaryMerkleTree deploys a new Ethereum contract, binding an instance of BinaryMerkleTree to it.
func DeployBinaryMerkleTree(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BinaryMerkleTree, error) {
	parsed, err := BinaryMerkleTreeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BinaryMerkleTreeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BinaryMerkleTree{BinaryMerkleTreeCaller: BinaryMerkleTreeCaller{contract: contract}, BinaryMerkleTreeTransactor: BinaryMerkleTreeTransactor{contract: contract}, BinaryMerkleTreeFilterer: BinaryMerkleTreeFilterer{contract: contract}}, nil
}

// BinaryMerkleTree is an auto generated Go binding around an Ethereum contract.
type BinaryMerkleTree struct {
	BinaryMerkleTreeCaller     // Read-only binding to the contract
	BinaryMerkleTreeTransactor // Write-only binding to the contract
	BinaryMerkleTreeFilterer   // Log filterer for contract events
}

// BinaryMerkleTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BinaryMerkleTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BinaryMerkleTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BinaryMerkleTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BinaryMerkleTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BinaryMerkleTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BinaryMerkleTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BinaryMerkleTreeSession struct {
	Contract     *BinaryMerkleTree // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BinaryMerkleTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BinaryMerkleTreeCallerSession struct {
	Contract *BinaryMerkleTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BinaryMerkleTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BinaryMerkleTreeTransactorSession struct {
	Contract     *BinaryMerkleTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BinaryMerkleTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BinaryMerkleTreeRaw struct {
	Contract *BinaryMerkleTree // Generic contract binding to access the raw methods on
}

// BinaryMerkleTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BinaryMerkleTreeCallerRaw struct {
	Contract *BinaryMerkleTreeCaller // Generic read-only contract binding to access the raw methods on
}

// BinaryMerkleTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BinaryMerkleTreeTransactorRaw struct {
	Contract *BinaryMerkleTreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBinaryMerkleTree creates a new instance of BinaryMerkleTree, bound to a specific deployed contract.
func NewBinaryMerkleTree(address common.Address, backend bind.ContractBackend) (*BinaryMerkleTree, error) {
	contract, err := bindBinaryMerkleTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BinaryMerkleTree{BinaryMerkleTreeCaller: BinaryMerkleTreeCaller{contract: contract}, BinaryMerkleTreeTransactor: BinaryMerkleTreeTransactor{contract: contract}, BinaryMerkleTreeFilterer: BinaryMerkleTreeFilterer{contract: contract}}, nil
}

// NewBinaryMerkleTreeCaller creates a new read-only instance of BinaryMerkleTree, bound to a specific deployed contract.
func NewBinaryMerkleTreeCaller(address common.Address, caller bind.ContractCaller) (*BinaryMerkleTreeCaller, error) {
	contract, err := bindBinaryMerkleTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BinaryMerkleTreeCaller{contract: contract}, nil
}

// NewBinaryMerkleTreeTransactor creates a new write-only instance of BinaryMerkleTree, bound to a specific deployed contract.
func NewBinaryMerkleTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*BinaryMerkleTreeTransactor, error) {
	contract, err := bindBinaryMerkleTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BinaryMerkleTreeTransactor{contract: contract}, nil
}

// NewBinaryMerkleTreeFilterer creates a new log filterer instance of BinaryMerkleTree, bound to a specific deployed contract.
func NewBinaryMerkleTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*BinaryMerkleTreeFilterer, error) {
	contract, err := bindBinaryMerkleTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BinaryMerkleTreeFilterer{contract: contract}, nil
}

// bindBinaryMerkleTree binds a generic wrapper to an already deployed contract.
func bindBinaryMerkleTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BinaryMerkleTreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BinaryMerkleTree *BinaryMerkleTreeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BinaryMerkleTree.Contract.BinaryMerkleTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BinaryMerkleTree *BinaryMerkleTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BinaryMerkleTree.Contract.BinaryMerkleTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BinaryMerkleTree *BinaryMerkleTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BinaryMerkleTree.Contract.BinaryMerkleTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BinaryMerkleTree *BinaryMerkleTreeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BinaryMerkleTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BinaryMerkleTree *BinaryMerkleTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BinaryMerkleTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BinaryMerkleTree *BinaryMerkleTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BinaryMerkleTree.Contract.contract.Transact(opts, method, params...)
}

// ConstantsMetaData contains all meta data concerning the Constants contract.
var ConstantsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122012539861920e3186959e4a0f034710a70fe4af456d0b0566dc49bb8ae3ab132664736f6c634300080c0033",
}

// ConstantsABI is the input ABI used to generate the binding from.
// Deprecated: Use ConstantsMetaData.ABI instead.
var ConstantsABI = ConstantsMetaData.ABI

// ConstantsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConstantsMetaData.Bin instead.
var ConstantsBin = ConstantsMetaData.Bin

// DeployConstants deploys a new Ethereum contract, binding an instance of Constants to it.
func DeployConstants(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Constants, error) {
	parsed, err := ConstantsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConstantsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Constants{ConstantsCaller: ConstantsCaller{contract: contract}, ConstantsTransactor: ConstantsTransactor{contract: contract}, ConstantsFilterer: ConstantsFilterer{contract: contract}}, nil
}

// Constants is an auto generated Go binding around an Ethereum contract.
type Constants struct {
	ConstantsCaller     // Read-only binding to the contract
	ConstantsTransactor // Write-only binding to the contract
	ConstantsFilterer   // Log filterer for contract events
}

// ConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstantsSession struct {
	Contract     *Constants        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstantsCallerSession struct {
	Contract *ConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstantsTransactorSession struct {
	Contract     *ConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstantsRaw struct {
	Contract *Constants // Generic contract binding to access the raw methods on
}

// ConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstantsCallerRaw struct {
	Contract *ConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// ConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstantsTransactorRaw struct {
	Contract *ConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstants creates a new instance of Constants, bound to a specific deployed contract.
func NewConstants(address common.Address, backend bind.ContractBackend) (*Constants, error) {
	contract, err := bindConstants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Constants{ConstantsCaller: ConstantsCaller{contract: contract}, ConstantsTransactor: ConstantsTransactor{contract: contract}, ConstantsFilterer: ConstantsFilterer{contract: contract}}, nil
}

// NewConstantsCaller creates a new read-only instance of Constants, bound to a specific deployed contract.
func NewConstantsCaller(address common.Address, caller bind.ContractCaller) (*ConstantsCaller, error) {
	contract, err := bindConstants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantsCaller{contract: contract}, nil
}

// NewConstantsTransactor creates a new write-only instance of Constants, bound to a specific deployed contract.
func NewConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstantsTransactor, error) {
	contract, err := bindConstants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantsTransactor{contract: contract}, nil
}

// NewConstantsFilterer creates a new log filterer instance of Constants, bound to a specific deployed contract.
func NewConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstantsFilterer, error) {
	contract, err := bindConstants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstantsFilterer{contract: contract}, nil
}

// bindConstants binds a generic wrapper to an already deployed contract.
func bindConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstantsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constants *ConstantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constants.Contract.ConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constants *ConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constants.Contract.ConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constants *ConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constants.Contract.ConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constants *ConstantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constants *ConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constants *ConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constants.Contract.contract.Transact(opts, method, params...)
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204d294be19df097f845e135c46ca541f26c9b64a484e10dc675363c2aa9d39e0f64736f6c634300080c0033",
}

// ECDSAABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAMetaData.ABI instead.
var ECDSAABI = ECDSAMetaData.ABI

// ECDSABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAMetaData.Bin instead.
var ECDSABin = ECDSAMetaData.Bin

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// IDAOracleMetaData contains all meta data concerning the IDAOracle contract.
var IDAOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tupleRootIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"_tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1f3302a9": "verifyAttestation(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))",
	},
}

// IDAOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IDAOracleMetaData.ABI instead.
var IDAOracleABI = IDAOracleMetaData.ABI

// Deprecated: Use IDAOracleMetaData.Sigs instead.
// IDAOracleFuncSigs maps the 4-byte function signature to its string representation.
var IDAOracleFuncSigs = IDAOracleMetaData.Sigs

// IDAOracle is an auto generated Go binding around an Ethereum contract.
type IDAOracle struct {
	IDAOracleCaller     // Read-only binding to the contract
	IDAOracleTransactor // Write-only binding to the contract
	IDAOracleFilterer   // Log filterer for contract events
}

// IDAOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDAOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDAOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDAOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDAOracleSession struct {
	Contract     *IDAOracle        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDAOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDAOracleCallerSession struct {
	Contract *IDAOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IDAOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDAOracleTransactorSession struct {
	Contract     *IDAOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IDAOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDAOracleRaw struct {
	Contract *IDAOracle // Generic contract binding to access the raw methods on
}

// IDAOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDAOracleCallerRaw struct {
	Contract *IDAOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IDAOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDAOracleTransactorRaw struct {
	Contract *IDAOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDAOracle creates a new instance of IDAOracle, bound to a specific deployed contract.
func NewIDAOracle(address common.Address, backend bind.ContractBackend) (*IDAOracle, error) {
	contract, err := bindIDAOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDAOracle{IDAOracleCaller: IDAOracleCaller{contract: contract}, IDAOracleTransactor: IDAOracleTransactor{contract: contract}, IDAOracleFilterer: IDAOracleFilterer{contract: contract}}, nil
}

// NewIDAOracleCaller creates a new read-only instance of IDAOracle, bound to a specific deployed contract.
func NewIDAOracleCaller(address common.Address, caller bind.ContractCaller) (*IDAOracleCaller, error) {
	contract, err := bindIDAOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDAOracleCaller{contract: contract}, nil
}

// NewIDAOracleTransactor creates a new write-only instance of IDAOracle, bound to a specific deployed contract.
func NewIDAOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IDAOracleTransactor, error) {
	contract, err := bindIDAOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDAOracleTransactor{contract: contract}, nil
}

// NewIDAOracleFilterer creates a new log filterer instance of IDAOracle, bound to a specific deployed contract.
func NewIDAOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IDAOracleFilterer, error) {
	contract, err := bindIDAOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDAOracleFilterer{contract: contract}, nil
}

// bindIDAOracle binds a generic wrapper to an already deployed contract.
func bindIDAOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDAOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDAOracle *IDAOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDAOracle.Contract.IDAOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDAOracle *IDAOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDAOracle.Contract.IDAOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDAOracle *IDAOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDAOracle.Contract.IDAOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDAOracle *IDAOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDAOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDAOracle *IDAOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDAOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDAOracle *IDAOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDAOracle.Contract.contract.Transact(opts, method, params...)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootIndex, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_IDAOracle *IDAOracleCaller) VerifyAttestation(opts *bind.CallOpts, _tupleRootIndex *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	var out []interface{}
	err := _IDAOracle.contract.Call(opts, &out, "verifyAttestation", _tupleRootIndex, _tuple, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootIndex, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_IDAOracle *IDAOracleSession) VerifyAttestation(_tupleRootIndex *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _IDAOracle.Contract.VerifyAttestation(&_IDAOracle.CallOpts, _tupleRootIndex, _tuple, _proof)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootIndex, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_IDAOracle *IDAOracleCallerSession) VerifyAttestation(_tupleRootIndex *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _IDAOracle.Contract.VerifyAttestation(&_IDAOracle.CallOpts, _tupleRootIndex, _tuple, _proof)
}

// QuantumGravityBridgeMetaData contains all meta data concerning the QuantumGravityBridge contract.
var QuantumGravityBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bridge_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_validatorSetHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataRootTupleRootNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidatorSetNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MalformedCurrentValidatorSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SuppliedValidatorSetInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRootTupleRoot\",\"type\":\"bytes32\"}],\"name\":\"DataRootTupleRootEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"powerThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"validatorSetHash\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_dataRootTupleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastDataRootTupleRootNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValidatorSetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValidatorSetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_dataRootTupleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"submitDataRootTupleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newPowerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_newValidatorSetHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tupleRootIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"_tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"af4482af": "BRIDGE_ID()",
		"817f985b": "state_dataRootTupleRoots(uint256)",
		"179e64c5": "state_lastDataRootTupleRootNonce()",
		"5433218c": "state_lastValidatorSetCheckpoint()",
		"0f0cdf25": "state_lastValidatorSetNonce()",
		"e5a2b5d2": "state_powerThreshold()",
		"e3e51e19": "submitDataRootTupleRoot(uint256,bytes32,(address,uint256)[],(uint8,bytes32,bytes32)[])",
		"4693c0df": "updateValidatorSet(uint256,uint256,bytes32,(address,uint256)[],(uint8,bytes32,bytes32)[])",
		"1f3302a9": "verifyAttestation(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))",
	},
	Bin: "0x60a060405234801561001057600080fd5b5060405161149a38038061149a83398101604081905261002f916100ce565b60808381526040805160208082018790526918da1958dadc1bda5b9d60b21b8284015260006060830181905293820186905260a08083018690528351808403909101815260c083018085528151919092012080855560018790559086905260e08201859052915183917fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c91908190036101000190a250505050506100fc565b6000806000606084860312156100e357600080fd5b8351925060208401519150604084015190509250925092565b608051611367610133600039600081816101230152818161021f0152818161026b0152818161035401526103a101526113676000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80635433218c116100665780635433218c146100f5578063817f985b146100fe578063af4482af1461011e578063e3e51e1914610145578063e5a2b5d21461015857600080fd5b80630f0cdf2514610098578063179e64c5146100b45780631f3302a9146100bd5780634693c0df146100e0575b600080fd5b6100a160025481565b6040519081526020015b60405180910390f35b6100a160035481565b6100d06100cb366004610e53565b610161565b60405190151581526020016100ab565b6100f36100ee366004611006565b6101c3565b005b6100a160005481565b6100a161010c366004611093565b60046020526000908152604090205481565b6100a17f000000000000000000000000000000000000000000000000000000000000000081565b6100f36101533660046110ac565b6102f9565b6100a160015481565b6000600354841115610175575060006101bc565b60008481526004602090815260408083205481518751818501529287015183830152815180840383018152606090930190915291906101b79083908690610472565b925050505b9392505050565b6002546001548189116101e9576040516368a35ffd60e11b815260040160405180910390fd5b8483146102095760405163c6617b7b60e01b815260040160405180910390fd5b600061021587876106e0565b90506000546102467f0000000000000000000000000000000000000000000000000000000000000000858585610714565b1461026457604051630bbdaec960e11b815260040160405180910390fd5b60006102927f0000000000000000000000000000000000000000000000000000000000000000858c8c610714565b90506102a2888888888588610765565b600081905560018a905560028b9055604080518b8152602081018b90528c917fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c910160405180910390a25050505050505050505050565b600154600354871161031e5760405163e869766d60e01b815260040160405180910390fd5b83821461033e5760405163c6617b7b60e01b815260040160405180910390fd5b600061034a86866106e0565b905060005461037d7f00000000000000000000000000000000000000000000000000000000000000006002548585610714565b1461039b57604051630bbdaec960e11b815260040160405180910390fd5b604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091526f0e8e4c2dce6c2c6e8d2dedc84c2e8c6d60831b82840152606082018b905260808083018b90528351808403909101815260a09092019092528051910120610415878787878588610765565b6003899055600089815260046020526040908190208990555189907f6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f9061045f908b815260200190565b60405180910390a2505050505050505050565b600060018360400151116104955782515115610490575060006101bc565b6104b7565b6104a783602001518460400151610873565b835151146104b7575060006101bc565b82604001518360200151106104ce575060006101bc565b60006104d9836108fa565b845151909150610503578360400151600114156104f957841490506101bc565b60009150506101bc565b60208401516001905b60208601516000906001841b90610524908290611145565b61052e9190611167565b90506000600161054081861b84611186565b61054a919061119e565b90508760400151811061055e575050610624565b91508161056c60018561119e565b88515111610582576000955050505050506101bc565b61058d60018561119e565b6001901b8289602001516105a1919061119e565b10156105de5787516105d79086906105ba60018861119e565b815181106105ca576105ca6111b5565b602002602001015161096f565b9450610610565b875161060d906105ef60018761119e565b815181106105ff576105ff6111b5565b60200260200101518661096f565b94505b61061b600185611186565b9350505061050c565b60018660400151610635919061119e565b811461067d5761064660018361119e565b8651511161065a57600093505050506101bc565b855161066d9084906105ba60018661119e565b925061067a600183611186565b91505b85515161068b60018461119e565b10156106d45785516106c0906106a260018561119e565b815181106106b2576106b26111b5565b60200260200101518461096f565b92506106cd600183611186565b915061067d565b50509093149392505050565b600082826040516020016106f59291906111e7565b6040516020818303038152906040528051906020012090505b92915050565b604080516020808201969096526918da1958dadc1bda5b9d60b21b818301526060810194909452608084019290925260a0808401919091528151808403909101815260c09092019052805191012090565b6000805b8681101561084857610791868683818110610786576107866111b5565b9050606002016109ed565b1561079b57610836565b6107e48888838181106107b0576107b06111b5565b6107c6926020604090920201908101915061123e565b858888858181106107d9576107d96111b5565b905060600201610a21565b61080157604051638baa579f60e01b815260040160405180910390fd5b878782818110610813576108136111b5565b90506040020160200135826108289190611186565b915082821061083657610848565b8061084081611259565b915050610769565b508181101561086a5760405163cabeb65560e01b815260040160405180910390fd5b50505050505050565b600061087e82610abb565b61088a9061010061119e565b9050600061089960018361119e565b6001901b90506001816108ac919061119e565b84116108b8575061070e565b80600114156108cb57600191505061070e565b6108e76108d8828661119e565b6108e2838661119e565b610873565b6108f2906001611186565b91505061070e565b60006002600060f81b836040516020016109159291906112af565b60408051601f198184030181529082905261092f916112d3565b602060405180830381855afa15801561094c573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061070e91906112df565b604051600160f81b6020820152602181018390526041810182905260009060029060610160408051601f19818403018152908290526109ad916112d3565b602060405180830381855afa1580156109ca573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906101bc91906112df565b60006020820135158015610a0357506040820135155b801561070e5750610a1760208301836112f8565b60ff161592915050565b600080610a7b846040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050610a9d81610a8e60208601866112f8565b85602001358660400135610ae8565b6001600160a01b0316856001600160a01b0316149150509392505050565b60005b81816001901b1015610adc57610ad5600182611186565b9050610abe565b61070e8161010061119e565b6000806000610af987878787610b10565b91509150610b0681610bfd565b5095945050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610b475750600090506003610bf4565b8460ff16601b14158015610b5f57508460ff16601c14155b15610b705750600090506004610bf4565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610bc4573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610bed57600060019250925050610bf4565b9150600090505b94509492505050565b6000816004811115610c1157610c1161131b565b1415610c1a5750565b6001816004811115610c2e57610c2e61131b565b1415610c815760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064015b60405180910390fd5b6002816004811115610c9557610c9561131b565b1415610ce35760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610c78565b6003816004811115610cf757610cf761131b565b1415610d505760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610c78565b6004816004811115610d6457610d6461131b565b1415610dbd5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610c78565b50565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610df957610df9610dc0565b60405290565b6040516060810167ffffffffffffffff81118282101715610df957610df9610dc0565b604051601f8201601f1916810167ffffffffffffffff81118282101715610e4b57610e4b610dc0565b604052919050565b60008060008385036080811215610e6957600080fd5b8435935060206040601f1983011215610e8157600080fd5b610e89610dd6565b86820135815260408701358282015293506060860135915067ffffffffffffffff80831115610eb757600080fd5b9186019160608389031215610ecb57600080fd5b610ed3610dff565b833582811115610ee257600080fd5b8401601f81018a13610ef357600080fd5b803583811115610f0557610f05610dc0565b8060051b9350610f16858501610e22565b818152938201850193858101908c861115610f3057600080fd5b928601925b85841015610f4e57833582529286019290860190610f35565b80855250505050828401358382015260408401356040820152809450505050509250925092565b60008083601f840112610f8757600080fd5b50813567ffffffffffffffff811115610f9f57600080fd5b6020830191508360208260061b8501011115610fba57600080fd5b9250929050565b60008083601f840112610fd357600080fd5b50813567ffffffffffffffff811115610feb57600080fd5b602083019150836020606083028501011115610fba57600080fd5b600080600080600080600060a0888a03121561102157600080fd5b873596506020880135955060408801359450606088013567ffffffffffffffff8082111561104e57600080fd5b61105a8b838c01610f75565b909650945060808a013591508082111561107357600080fd5b506110808a828b01610fc1565b989b979a50959850939692959293505050565b6000602082840312156110a557600080fd5b5035919050565b600080600080600080608087890312156110c557600080fd5b8635955060208701359450604087013567ffffffffffffffff808211156110eb57600080fd5b6110f78a838b01610f75565b9096509450606089013591508082111561111057600080fd5b5061111d89828a01610fc1565b979a9699509497509295939492505050565b634e487b7160e01b600052601160045260246000fd5b60008261116257634e487b7160e01b600052601260045260246000fd5b500490565b60008160001904831182151516156111815761118161112f565b500290565b600082198211156111995761119961112f565b500190565b6000828210156111b0576111b061112f565b500390565b634e487b7160e01b600052603260045260246000fd5b80356001600160a01b03811681146111e257600080fd5b919050565b6020808252818101839052600090604080840186845b87811015611231576001600160a01b03611216836111cb565b168352818501358584015291830191908301906001016111fd565b5090979650505050505050565b60006020828403121561125057600080fd5b6101bc826111cb565b600060001982141561126d5761126d61112f565b5060010190565b6000815160005b81811015611295576020818501810151868301520161127b565b818111156112a4576000828601525b509290920192915050565b6001600160f81b03198316815260006112cb6001830184611274565b949350505050565b60006101bc8284611274565b6000602082840312156112f157600080fd5b5051919050565b60006020828403121561130a57600080fd5b813560ff811681146101bc57600080fd5b634e487b7160e01b600052602160045260246000fdfea2646970667358221220f91c02b452291bac00bc82a80a18edfe0af5154e54bac31c41cf4236ff90e6bd64736f6c634300080c0033",
}

// QuantumGravityBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use QuantumGravityBridgeMetaData.ABI instead.
var QuantumGravityBridgeABI = QuantumGravityBridgeMetaData.ABI

// Deprecated: Use QuantumGravityBridgeMetaData.Sigs instead.
// QuantumGravityBridgeFuncSigs maps the 4-byte function signature to its string representation.
var QuantumGravityBridgeFuncSigs = QuantumGravityBridgeMetaData.Sigs

// QuantumGravityBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QuantumGravityBridgeMetaData.Bin instead.
var QuantumGravityBridgeBin = QuantumGravityBridgeMetaData.Bin

// DeployQuantumGravityBridge deploys a new Ethereum contract, binding an instance of QuantumGravityBridge to it.
func DeployQuantumGravityBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge_id [32]byte, _powerThreshold *big.Int, _validatorSetHash [32]byte) (common.Address, *types.Transaction, *QuantumGravityBridge, error) {
	parsed, err := QuantumGravityBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QuantumGravityBridgeBin), backend, _bridge_id, _powerThreshold, _validatorSetHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QuantumGravityBridge{QuantumGravityBridgeCaller: QuantumGravityBridgeCaller{contract: contract}, QuantumGravityBridgeTransactor: QuantumGravityBridgeTransactor{contract: contract}, QuantumGravityBridgeFilterer: QuantumGravityBridgeFilterer{contract: contract}}, nil
}

// QuantumGravityBridge is an auto generated Go binding around an Ethereum contract.
type QuantumGravityBridge struct {
	QuantumGravityBridgeCaller     // Read-only binding to the contract
	QuantumGravityBridgeTransactor // Write-only binding to the contract
	QuantumGravityBridgeFilterer   // Log filterer for contract events
}

// QuantumGravityBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuantumGravityBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuantumGravityBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuantumGravityBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuantumGravityBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuantumGravityBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuantumGravityBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuantumGravityBridgeSession struct {
	Contract     *QuantumGravityBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// QuantumGravityBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuantumGravityBridgeCallerSession struct {
	Contract *QuantumGravityBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// QuantumGravityBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuantumGravityBridgeTransactorSession struct {
	Contract     *QuantumGravityBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// QuantumGravityBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuantumGravityBridgeRaw struct {
	Contract *QuantumGravityBridge // Generic contract binding to access the raw methods on
}

// QuantumGravityBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuantumGravityBridgeCallerRaw struct {
	Contract *QuantumGravityBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// QuantumGravityBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuantumGravityBridgeTransactorRaw struct {
	Contract *QuantumGravityBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuantumGravityBridge creates a new instance of QuantumGravityBridge, bound to a specific deployed contract.
func NewQuantumGravityBridge(address common.Address, backend bind.ContractBackend) (*QuantumGravityBridge, error) {
	contract, err := bindQuantumGravityBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridge{QuantumGravityBridgeCaller: QuantumGravityBridgeCaller{contract: contract}, QuantumGravityBridgeTransactor: QuantumGravityBridgeTransactor{contract: contract}, QuantumGravityBridgeFilterer: QuantumGravityBridgeFilterer{contract: contract}}, nil
}

// NewQuantumGravityBridgeCaller creates a new read-only instance of QuantumGravityBridge, bound to a specific deployed contract.
func NewQuantumGravityBridgeCaller(address common.Address, caller bind.ContractCaller) (*QuantumGravityBridgeCaller, error) {
	contract, err := bindQuantumGravityBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeCaller{contract: contract}, nil
}

// NewQuantumGravityBridgeTransactor creates a new write-only instance of QuantumGravityBridge, bound to a specific deployed contract.
func NewQuantumGravityBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*QuantumGravityBridgeTransactor, error) {
	contract, err := bindQuantumGravityBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeTransactor{contract: contract}, nil
}

// NewQuantumGravityBridgeFilterer creates a new log filterer instance of QuantumGravityBridge, bound to a specific deployed contract.
func NewQuantumGravityBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*QuantumGravityBridgeFilterer, error) {
	contract, err := bindQuantumGravityBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeFilterer{contract: contract}, nil
}

// bindQuantumGravityBridge binds a generic wrapper to an already deployed contract.
func bindQuantumGravityBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuantumGravityBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuantumGravityBridge *QuantumGravityBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuantumGravityBridge.Contract.QuantumGravityBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuantumGravityBridge *QuantumGravityBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.QuantumGravityBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuantumGravityBridge *QuantumGravityBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.QuantumGravityBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuantumGravityBridge *QuantumGravityBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuantumGravityBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEID is a free data retrieval call binding the contract method 0xaf4482af.
//
// Solidity: function BRIDGE_ID() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) BRIDGEID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "BRIDGE_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BRIDGEID is a free data retrieval call binding the contract method 0xaf4482af.
//
// Solidity: function BRIDGE_ID() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) BRIDGEID() ([32]byte, error) {
	return _QuantumGravityBridge.Contract.BRIDGEID(&_QuantumGravityBridge.CallOpts)
}

// BRIDGEID is a free data retrieval call binding the contract method 0xaf4482af.
//
// Solidity: function BRIDGE_ID() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) BRIDGEID() ([32]byte, error) {
	return _QuantumGravityBridge.Contract.BRIDGEID(&_QuantumGravityBridge.CallOpts)
}

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StateDataRootTupleRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_dataRootTupleRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StateDataRootTupleRoots(arg0 *big.Int) ([32]byte, error) {
	return _QuantumGravityBridge.Contract.StateDataRootTupleRoots(&_QuantumGravityBridge.CallOpts, arg0)
}

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StateDataRootTupleRoots(arg0 *big.Int) ([32]byte, error) {
	return _QuantumGravityBridge.Contract.StateDataRootTupleRoots(&_QuantumGravityBridge.CallOpts, arg0)
}

// StateLastDataRootTupleRootNonce is a free data retrieval call binding the contract method 0x179e64c5.
//
// Solidity: function state_lastDataRootTupleRootNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StateLastDataRootTupleRootNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_lastDataRootTupleRootNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastDataRootTupleRootNonce is a free data retrieval call binding the contract method 0x179e64c5.
//
// Solidity: function state_lastDataRootTupleRootNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StateLastDataRootTupleRootNonce() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StateLastDataRootTupleRootNonce(&_QuantumGravityBridge.CallOpts)
}

// StateLastDataRootTupleRootNonce is a free data retrieval call binding the contract method 0x179e64c5.
//
// Solidity: function state_lastDataRootTupleRootNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StateLastDataRootTupleRootNonce() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StateLastDataRootTupleRootNonce(&_QuantumGravityBridge.CallOpts)
}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StateLastValidatorSetCheckpoint(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_lastValidatorSetCheckpoint")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StateLastValidatorSetCheckpoint() ([32]byte, error) {
	return _QuantumGravityBridge.Contract.StateLastValidatorSetCheckpoint(&_QuantumGravityBridge.CallOpts)
}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StateLastValidatorSetCheckpoint() ([32]byte, error) {
	return _QuantumGravityBridge.Contract.StateLastValidatorSetCheckpoint(&_QuantumGravityBridge.CallOpts)
}

// StateLastValidatorSetNonce is a free data retrieval call binding the contract method 0x0f0cdf25.
//
// Solidity: function state_lastValidatorSetNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StateLastValidatorSetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_lastValidatorSetNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastValidatorSetNonce is a free data retrieval call binding the contract method 0x0f0cdf25.
//
// Solidity: function state_lastValidatorSetNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StateLastValidatorSetNonce() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StateLastValidatorSetNonce(&_QuantumGravityBridge.CallOpts)
}

// StateLastValidatorSetNonce is a free data retrieval call binding the contract method 0x0f0cdf25.
//
// Solidity: function state_lastValidatorSetNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StateLastValidatorSetNonce() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StateLastValidatorSetNonce(&_QuantumGravityBridge.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StatePowerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_powerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StatePowerThreshold() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StatePowerThreshold(&_QuantumGravityBridge.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StatePowerThreshold() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StatePowerThreshold(&_QuantumGravityBridge.CallOpts)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootIndex, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) VerifyAttestation(opts *bind.CallOpts, _tupleRootIndex *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "verifyAttestation", _tupleRootIndex, _tuple, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootIndex, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) VerifyAttestation(_tupleRootIndex *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _QuantumGravityBridge.Contract.VerifyAttestation(&_QuantumGravityBridge.CallOpts, _tupleRootIndex, _tuple, _proof)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootIndex, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) VerifyAttestation(_tupleRootIndex *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _QuantumGravityBridge.Contract.VerifyAttestation(&_QuantumGravityBridge.CallOpts, _tupleRootIndex, _tuple, _proof)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe3e51e19.
//
// Solidity: function submitDataRootTupleRoot(uint256 _nonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactor) SubmitDataRootTupleRoot(opts *bind.TransactOpts, _nonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.contract.Transact(opts, "submitDataRootTupleRoot", _nonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe3e51e19.
//
// Solidity: function submitDataRootTupleRoot(uint256 _nonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeSession) SubmitDataRootTupleRoot(_nonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.SubmitDataRootTupleRoot(&_QuantumGravityBridge.TransactOpts, _nonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe3e51e19.
//
// Solidity: function submitDataRootTupleRoot(uint256 _nonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorSession) SubmitDataRootTupleRoot(_nonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.SubmitDataRootTupleRoot(&_QuantumGravityBridge.TransactOpts, _nonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x4693c0df.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactor) UpdateValidatorSet(opts *bind.TransactOpts, _newNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.contract.Transact(opts, "updateValidatorSet", _newNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x4693c0df.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeSession) UpdateValidatorSet(_newNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.UpdateValidatorSet(&_QuantumGravityBridge.TransactOpts, _newNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x4693c0df.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorSession) UpdateValidatorSet(_newNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.UpdateValidatorSet(&_QuantumGravityBridge.TransactOpts, _newNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// QuantumGravityBridgeDataRootTupleRootEventIterator is returned from FilterDataRootTupleRootEvent and is used to iterate over the raw logs and unpacked data for DataRootTupleRootEvent events raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeDataRootTupleRootEventIterator struct {
	Event *QuantumGravityBridgeDataRootTupleRootEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuantumGravityBridgeDataRootTupleRootEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuantumGravityBridgeDataRootTupleRootEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuantumGravityBridgeDataRootTupleRootEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuantumGravityBridgeDataRootTupleRootEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuantumGravityBridgeDataRootTupleRootEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuantumGravityBridgeDataRootTupleRootEvent represents a DataRootTupleRootEvent event raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeDataRootTupleRootEvent struct {
	Nonce             *big.Int
	DataRootTupleRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDataRootTupleRootEvent is a free log retrieval operation binding the contract event 0x6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f.
//
// Solidity: event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) FilterDataRootTupleRootEvent(opts *bind.FilterOpts, nonce []*big.Int) (*QuantumGravityBridgeDataRootTupleRootEventIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.FilterLogs(opts, "DataRootTupleRootEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeDataRootTupleRootEventIterator{contract: _QuantumGravityBridge.contract, event: "DataRootTupleRootEvent", logs: logs, sub: sub}, nil
}

// WatchDataRootTupleRootEvent is a free log subscription operation binding the contract event 0x6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f.
//
// Solidity: event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) WatchDataRootTupleRootEvent(opts *bind.WatchOpts, sink chan<- *QuantumGravityBridgeDataRootTupleRootEvent, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.WatchLogs(opts, "DataRootTupleRootEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuantumGravityBridgeDataRootTupleRootEvent)
				if err := _QuantumGravityBridge.contract.UnpackLog(event, "DataRootTupleRootEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDataRootTupleRootEvent is a log parse operation binding the contract event 0x6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f.
//
// Solidity: event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) ParseDataRootTupleRootEvent(log types.Log) (*QuantumGravityBridgeDataRootTupleRootEvent, error) {
	event := new(QuantumGravityBridgeDataRootTupleRootEvent)
	if err := _QuantumGravityBridge.contract.UnpackLog(event, "DataRootTupleRootEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuantumGravityBridgeValidatorSetUpdatedEventIterator is returned from FilterValidatorSetUpdatedEvent and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdatedEvent events raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeValidatorSetUpdatedEventIterator struct {
	Event *QuantumGravityBridgeValidatorSetUpdatedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *QuantumGravityBridgeValidatorSetUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuantumGravityBridgeValidatorSetUpdatedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(QuantumGravityBridgeValidatorSetUpdatedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *QuantumGravityBridgeValidatorSetUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuantumGravityBridgeValidatorSetUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuantumGravityBridgeValidatorSetUpdatedEvent represents a ValidatorSetUpdatedEvent event raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeValidatorSetUpdatedEvent struct {
	Nonce            *big.Int
	PowerThreshold   *big.Int
	ValidatorSetHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdatedEvent is a free log retrieval operation binding the contract event 0xe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c.
//
// Solidity: event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) FilterValidatorSetUpdatedEvent(opts *bind.FilterOpts, nonce []*big.Int) (*QuantumGravityBridgeValidatorSetUpdatedEventIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.FilterLogs(opts, "ValidatorSetUpdatedEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeValidatorSetUpdatedEventIterator{contract: _QuantumGravityBridge.contract, event: "ValidatorSetUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdatedEvent is a free log subscription operation binding the contract event 0xe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c.
//
// Solidity: event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) WatchValidatorSetUpdatedEvent(opts *bind.WatchOpts, sink chan<- *QuantumGravityBridgeValidatorSetUpdatedEvent, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.WatchLogs(opts, "ValidatorSetUpdatedEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuantumGravityBridgeValidatorSetUpdatedEvent)
				if err := _QuantumGravityBridge.contract.UnpackLog(event, "ValidatorSetUpdatedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorSetUpdatedEvent is a log parse operation binding the contract event 0xe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c.
//
// Solidity: event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) ParseValidatorSetUpdatedEvent(log types.Log) (*QuantumGravityBridgeValidatorSetUpdatedEvent, error) {
	event := new(QuantumGravityBridgeValidatorSetUpdatedEvent)
	if err := _QuantumGravityBridge.contract.UnpackLog(event, "ValidatorSetUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
