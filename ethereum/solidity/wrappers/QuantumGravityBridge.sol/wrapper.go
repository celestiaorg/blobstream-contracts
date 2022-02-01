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

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextUpgradeableMetaData.ABI instead.
var ContextUpgradeableABI = ContextUpgradeableMetaData.ABI

// ContextUpgradeable is an auto generated Go binding around an Ethereum contract.
type ContextUpgradeable struct {
	ContextUpgradeableCaller     // Read-only binding to the contract
	ContextUpgradeableTransactor // Write-only binding to the contract
	ContextUpgradeableFilterer   // Log filterer for contract events
}

// ContextUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextUpgradeableSession struct {
	Contract     *ContextUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContextUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextUpgradeableCallerSession struct {
	Contract *ContextUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContextUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextUpgradeableTransactorSession struct {
	Contract     *ContextUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContextUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextUpgradeableRaw struct {
	Contract *ContextUpgradeable // Generic contract binding to access the raw methods on
}

// ContextUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextUpgradeableCallerRaw struct {
	Contract *ContextUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactorRaw struct {
	Contract *ContextUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextUpgradeable creates a new instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextUpgradeable, error) {
	contract, err := bindContextUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeable{ContextUpgradeableCaller: ContextUpgradeableCaller{contract: contract}, ContextUpgradeableTransactor: ContextUpgradeableTransactor{contract: contract}, ContextUpgradeableFilterer: ContextUpgradeableFilterer{contract: contract}}, nil
}

// NewContextUpgradeableCaller creates a new read-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextUpgradeableCaller, error) {
	contract, err := bindContextUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableCaller{contract: contract}, nil
}

// NewContextUpgradeableTransactor creates a new write-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextUpgradeableTransactor, error) {
	contract, err := bindContextUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableTransactor{contract: contract}, nil
}

// NewContextUpgradeableFilterer creates a new log filterer instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextUpgradeableFilterer, error) {
	contract, err := bindContextUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableFilterer{contract: contract}, nil
}

// bindContextUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.ContextUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a069cfd6a827c489a6a0784f0b962b0d334b54b46aa42da33bea4f6b9adb115764736f6c634300080b0033",
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

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// OwnableUpgradeableWithExpiryMetaData contains all meta data concerning the OwnableUpgradeableWithExpiry contract.
var OwnableUpgradeableWithExpiryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getOwnershipExpiryTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwnershipExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnershipAfterExpiry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1ee7a108": "getOwnershipExpiryTimestamp()",
		"5afe97bb": "isOwnershipExpired()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"8c64865f": "renounceOwnershipAfterExpiry()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableUpgradeableWithExpiryABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableUpgradeableWithExpiryMetaData.ABI instead.
var OwnableUpgradeableWithExpiryABI = OwnableUpgradeableWithExpiryMetaData.ABI

// Deprecated: Use OwnableUpgradeableWithExpiryMetaData.Sigs instead.
// OwnableUpgradeableWithExpiryFuncSigs maps the 4-byte function signature to its string representation.
var OwnableUpgradeableWithExpiryFuncSigs = OwnableUpgradeableWithExpiryMetaData.Sigs

// OwnableUpgradeableWithExpiry is an auto generated Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiry struct {
	OwnableUpgradeableWithExpiryCaller     // Read-only binding to the contract
	OwnableUpgradeableWithExpiryTransactor // Write-only binding to the contract
	OwnableUpgradeableWithExpiryFilterer   // Log filterer for contract events
}

// OwnableUpgradeableWithExpiryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableWithExpiryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableWithExpiryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableUpgradeableWithExpiryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableWithExpirySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableUpgradeableWithExpirySession struct {
	Contract     *OwnableUpgradeableWithExpiry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OwnableUpgradeableWithExpiryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableUpgradeableWithExpiryCallerSession struct {
	Contract *OwnableUpgradeableWithExpiryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// OwnableUpgradeableWithExpiryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableUpgradeableWithExpiryTransactorSession struct {
	Contract     *OwnableUpgradeableWithExpiryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// OwnableUpgradeableWithExpiryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiryRaw struct {
	Contract *OwnableUpgradeableWithExpiry // Generic contract binding to access the raw methods on
}

// OwnableUpgradeableWithExpiryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiryCallerRaw struct {
	Contract *OwnableUpgradeableWithExpiryCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableUpgradeableWithExpiryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiryTransactorRaw struct {
	Contract *OwnableUpgradeableWithExpiryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableUpgradeableWithExpiry creates a new instance of OwnableUpgradeableWithExpiry, bound to a specific deployed contract.
func NewOwnableUpgradeableWithExpiry(address common.Address, backend bind.ContractBackend) (*OwnableUpgradeableWithExpiry, error) {
	contract, err := bindOwnableUpgradeableWithExpiry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableWithExpiry{OwnableUpgradeableWithExpiryCaller: OwnableUpgradeableWithExpiryCaller{contract: contract}, OwnableUpgradeableWithExpiryTransactor: OwnableUpgradeableWithExpiryTransactor{contract: contract}, OwnableUpgradeableWithExpiryFilterer: OwnableUpgradeableWithExpiryFilterer{contract: contract}}, nil
}

// NewOwnableUpgradeableWithExpiryCaller creates a new read-only instance of OwnableUpgradeableWithExpiry, bound to a specific deployed contract.
func NewOwnableUpgradeableWithExpiryCaller(address common.Address, caller bind.ContractCaller) (*OwnableUpgradeableWithExpiryCaller, error) {
	contract, err := bindOwnableUpgradeableWithExpiry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableWithExpiryCaller{contract: contract}, nil
}

// NewOwnableUpgradeableWithExpiryTransactor creates a new write-only instance of OwnableUpgradeableWithExpiry, bound to a specific deployed contract.
func NewOwnableUpgradeableWithExpiryTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableUpgradeableWithExpiryTransactor, error) {
	contract, err := bindOwnableUpgradeableWithExpiry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableWithExpiryTransactor{contract: contract}, nil
}

// NewOwnableUpgradeableWithExpiryFilterer creates a new log filterer instance of OwnableUpgradeableWithExpiry, bound to a specific deployed contract.
func NewOwnableUpgradeableWithExpiryFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableUpgradeableWithExpiryFilterer, error) {
	contract, err := bindOwnableUpgradeableWithExpiry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableWithExpiryFilterer{contract: contract}, nil
}

// bindOwnableUpgradeableWithExpiry binds a generic wrapper to an already deployed contract.
func bindOwnableUpgradeableWithExpiry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableUpgradeableWithExpiryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeableWithExpiry.Contract.OwnableUpgradeableWithExpiryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.OwnableUpgradeableWithExpiryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.OwnableUpgradeableWithExpiryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeableWithExpiry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.contract.Transact(opts, method, params...)
}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCaller) GetOwnershipExpiryTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OwnableUpgradeableWithExpiry.contract.Call(opts, &out, "getOwnershipExpiryTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) GetOwnershipExpiryTimestamp() (*big.Int, error) {
	return _OwnableUpgradeableWithExpiry.Contract.GetOwnershipExpiryTimestamp(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCallerSession) GetOwnershipExpiryTimestamp() (*big.Int, error) {
	return _OwnableUpgradeableWithExpiry.Contract.GetOwnershipExpiryTimestamp(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCaller) IsOwnershipExpired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OwnableUpgradeableWithExpiry.contract.Call(opts, &out, "isOwnershipExpired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) IsOwnershipExpired() (bool, error) {
	return _OwnableUpgradeableWithExpiry.Contract.IsOwnershipExpired(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCallerSession) IsOwnershipExpired() (bool, error) {
	return _OwnableUpgradeableWithExpiry.Contract.IsOwnershipExpired(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableUpgradeableWithExpiry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) Owner() (common.Address, error) {
	return _OwnableUpgradeableWithExpiry.Contract.Owner(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCallerSession) Owner() (common.Address, error) {
	return _OwnableUpgradeableWithExpiry.Contract.Owner(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.RenounceOwnership(&_OwnableUpgradeableWithExpiry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.RenounceOwnership(&_OwnableUpgradeableWithExpiry.TransactOpts)
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactor) RenounceOwnershipAfterExpiry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.contract.Transact(opts, "renounceOwnershipAfterExpiry")
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) RenounceOwnershipAfterExpiry() (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.RenounceOwnershipAfterExpiry(&_OwnableUpgradeableWithExpiry.TransactOpts)
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactorSession) RenounceOwnershipAfterExpiry() (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.RenounceOwnershipAfterExpiry(&_OwnableUpgradeableWithExpiry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.TransferOwnership(&_OwnableUpgradeableWithExpiry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.TransferOwnership(&_OwnableUpgradeableWithExpiry.TransactOpts, newOwner)
}

// OwnableUpgradeableWithExpiryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnableUpgradeableWithExpiry contract.
type OwnableUpgradeableWithExpiryOwnershipTransferredIterator struct {
	Event *OwnableUpgradeableWithExpiryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableUpgradeableWithExpiryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableWithExpiryOwnershipTransferred)
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
		it.Event = new(OwnableUpgradeableWithExpiryOwnershipTransferred)
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
func (it *OwnableUpgradeableWithExpiryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableWithExpiryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableWithExpiryOwnershipTransferred represents a OwnershipTransferred event raised by the OwnableUpgradeableWithExpiry contract.
type OwnableUpgradeableWithExpiryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableUpgradeableWithExpiryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeableWithExpiry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableWithExpiryOwnershipTransferredIterator{contract: _OwnableUpgradeableWithExpiry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableWithExpiryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeableWithExpiry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableWithExpiryOwnershipTransferred)
				if err := _OwnableUpgradeableWithExpiry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableUpgradeableWithExpiryOwnershipTransferred, error) {
	event := new(OwnableUpgradeableWithExpiryOwnershipTransferred)
	if err := _OwnableUpgradeableWithExpiry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuantumGravityBridgeMetaData contains all meta data concerning the QuantumGravityBridge contract.
var QuantumGravityBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bridge_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_validatorSetHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageTupleRootNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidatorSetNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MalformedCurrentValidatorSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SuppliedValidatorSetInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageTupleRoot\",\"type\":\"bytes32\"}],\"name\":\"MessageTupleRootEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"powerThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"validatorSetHash\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnershipExpiryTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwnershipExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnershipAfterExpiry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastMessageTupleRootNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValidatorSetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValidatorSetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_messageTupleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_messageTupleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"submitMessageTupleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newPowerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_newValidatorSetHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"af4482af": "BRIDGE_ID()",
		"1ee7a108": "getOwnershipExpiryTimestamp()",
		"5afe97bb": "isOwnershipExpired()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"8c64865f": "renounceOwnershipAfterExpiry()",
		"9a1cfb83": "state_lastMessageTupleRootNonce()",
		"5433218c": "state_lastValidatorSetCheckpoint()",
		"0f0cdf25": "state_lastValidatorSetNonce()",
		"d8340ecf": "state_messageTupleRoots(uint256)",
		"e5a2b5d2": "state_powerThreshold()",
		"030643d7": "submitMessageTupleRoot(uint256,bytes32,(address,uint256)[],(uint8,bytes32,bytes32)[])",
		"f2fde38b": "transferOwnership(address)",
		"4693c0df": "updateValidatorSet(uint256,uint256,bytes32,(address,uint256)[],(uint8,bytes32,bytes32)[])",
	},
	Bin: "0x60a060405234801561001057600080fd5b5060405161111b38038061111b83398101604081905261002f916101d6565b6100376100d8565b60808381526040805160208082018790526918da1958dadc1bda5b9d60b21b8284015260006060830181905293820186905260a08083018690528351808403909101815260c0830180855281519190920120606681905560678790559086905260e08201859052915183917fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c91908190036101000190a25050505050610204565b600054610100900460ff16806100f1575060005460ff16155b6101585760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff1615801561017a576000805461ffff19166101011790555b603380546001600160a01b031916339081179091554260345560405181906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35080156101d3576000805461ff00191690555b50565b6000806000606084860312156101eb57600080fd5b8351925060208401519150604084015190509250925092565b608051610ee061023b600039600081816101950152818161024e0152818161029b015281816103e1015261042d0152610ee06000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638c64865f1161008c578063af4482af11610066578063af4482af14610190578063d8340ecf146101b7578063e5a2b5d2146101d7578063f2fde38b146101e057600080fd5b80638c64865f146101645780638da5cb5b1461016c5780639a1cfb831461018757600080fd5b80634693c0df116100c85780634693c0df146101285780635433218c1461013b5780635afe97bb14610144578063715018a61461015c57600080fd5b8063030643d7146100ef5780630f0cdf25146101045780631ee7a10814610120575b600080fd5b6101026100fd366004610c54565b6101f3565b005b61010d60685481565b6040519081526020015b60405180910390f35b61010d61036c565b610102610136366004610cd7565b610385565b61010d60665481565b61014c6104bb565b6040519015158152602001610117565b6101026104cc565b610102610535565b6033546040516001600160a01b039091168152602001610117565b61010d60695481565b61010d7f000000000000000000000000000000000000000000000000000000000000000081565b61010d6101c5366004610d64565b606a6020526000908152604090205481565b61010d60675481565b6101026101ee366004610d99565b610589565b606754606954871161021857604051632d37469760e01b815260040160405180910390fd5b8382146102385760405163c6617b7b60e01b815260040160405180910390fd5b600061024486866106a4565b90506066546102777f000000000000000000000000000000000000000000000000000000000000000060685485856106d7565b1461029557604051630bbdaec960e11b815260040160405180910390fd5b604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091526f0e8e4c2dce6c2c6e8d2dedc84c2e8c6d60831b82840152606082018b905260808083018b90528351808403909101815260a0909201909252805191012061030f878787878588610728565b60698990556000898152606a6020526040908190208990555189907f6f71145a387d176c99498fbbb78a98e85a46472eb5a4de9e06242e4edf4527e490610359908b815260200190565b60405180910390a2505050505050505050565b60006034546301dfe2006103809190610dd1565b905090565b6068546067548189116103ab576040516368a35ffd60e11b815260040160405180910390fd5b8483146103cb5760405163c6617b7b60e01b815260040160405180910390fd5b60006103d787876106a4565b90506066546104087f00000000000000000000000000000000000000000000000000000000000000008585856106d7565b1461042657604051630bbdaec960e11b815260040160405180910390fd5b60006104547f0000000000000000000000000000000000000000000000000000000000000000858c8c6106d7565b9050610464888888888588610728565b606681905560678a905560688b9055604080518b8152602081018b90528c917fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c910160405180910390a25050505050505050505050565b60006104c561036c565b4211905090565b6033546001600160a01b0316331461052b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b61053361080c565b565b61053d6104bb565b61052b5760405162461bcd60e51b815260206004820152601960248201527f4f776e657273686970206e6f74207965742065787069726564000000000000006044820152606401610522565b6033546001600160a01b031633146105e35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610522565b6001600160a01b0381166106485760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610522565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b600082826040516020016106b9929190610de9565b60405160208183030381529060405280519060200120905092915050565b604080516020808201969096526918da1958dadc1bda5b9d60b21b818301526060810194909452608084019290925260a0808401919091528151808403909101815260c09092019052805191012090565b6000805b868110156107e15761077d88888381811061074957610749610e40565b61075f9260206040909202019081019150610d99565b8588888581811061077257610772610e40565b905060600201610856565b61079a57604051638baa579f60e01b815260040160405180910390fd5b8787828181106107ac576107ac610e40565b90506040020160200135826107c19190610dd1565b91508282106107cf576107e1565b806107d981610e56565b91505061072c565b50818110156108035760405163cabeb65560e01b815260040160405180910390fd5b50505050505050565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380546001600160a01b0319169055565b6000806108b0846040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b90506108d2816108c36020860186610e71565b856020013586604001356108f0565b6001600160a01b0316856001600160a01b0316149150509392505050565b600080600061090187878787610918565b9150915061090e81610a05565b5095945050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561094f57506000905060036109fc565b8460ff16601b1415801561096757508460ff16601c14155b1561097857506000905060046109fc565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156109cc573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166109f5576000600192509250506109fc565b9150600090505b94509492505050565b6000816004811115610a1957610a19610e94565b1415610a225750565b6001816004811115610a3657610a36610e94565b1415610a845760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610522565b6002816004811115610a9857610a98610e94565b1415610ae65760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610522565b6003816004811115610afa57610afa610e94565b1415610b535760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610522565b6004816004811115610b6757610b67610e94565b1415610bc05760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610522565b50565b60008083601f840112610bd557600080fd5b50813567ffffffffffffffff811115610bed57600080fd5b6020830191508360208260061b8501011115610c0857600080fd5b9250929050565b60008083601f840112610c2157600080fd5b50813567ffffffffffffffff811115610c3957600080fd5b602083019150836020606083028501011115610c0857600080fd5b60008060008060008060808789031215610c6d57600080fd5b8635955060208701359450604087013567ffffffffffffffff80821115610c9357600080fd5b610c9f8a838b01610bc3565b90965094506060890135915080821115610cb857600080fd5b50610cc589828a01610c0f565b979a9699509497509295939492505050565b600080600080600080600060a0888a031215610cf257600080fd5b873596506020880135955060408801359450606088013567ffffffffffffffff80821115610d1f57600080fd5b610d2b8b838c01610bc3565b909650945060808a0135915080821115610d4457600080fd5b50610d518a828b01610c0f565b989b979a50959850939692959293505050565b600060208284031215610d7657600080fd5b5035919050565b80356001600160a01b0381168114610d9457600080fd5b919050565b600060208284031215610dab57600080fd5b610db482610d7d565b9392505050565b634e487b7160e01b600052601160045260246000fd5b60008219821115610de457610de4610dbb565b500190565b6020808252818101839052600090604080840186845b87811015610e33576001600160a01b03610e1883610d7d565b16835281850135858401529183019190830190600101610dff565b5090979650505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600019821415610e6a57610e6a610dbb565b5060010190565b600060208284031215610e8357600080fd5b813560ff81168114610db457600080fd5b634e487b7160e01b600052602160045260246000fdfea2646970667358221220720c4e23c19cd6cf187193ba63bb7310301264b4190f3d24db912656e9be17c064736f6c634300080b0033",
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

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) GetOwnershipExpiryTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "getOwnershipExpiryTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) GetOwnershipExpiryTimestamp() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.GetOwnershipExpiryTimestamp(&_QuantumGravityBridge.CallOpts)
}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) GetOwnershipExpiryTimestamp() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.GetOwnershipExpiryTimestamp(&_QuantumGravityBridge.CallOpts)
}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) IsOwnershipExpired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "isOwnershipExpired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) IsOwnershipExpired() (bool, error) {
	return _QuantumGravityBridge.Contract.IsOwnershipExpired(&_QuantumGravityBridge.CallOpts)
}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) IsOwnershipExpired() (bool, error) {
	return _QuantumGravityBridge.Contract.IsOwnershipExpired(&_QuantumGravityBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) Owner() (common.Address, error) {
	return _QuantumGravityBridge.Contract.Owner(&_QuantumGravityBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) Owner() (common.Address, error) {
	return _QuantumGravityBridge.Contract.Owner(&_QuantumGravityBridge.CallOpts)
}

// StateLastMessageTupleRootNonce is a free data retrieval call binding the contract method 0x9a1cfb83.
//
// Solidity: function state_lastMessageTupleRootNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StateLastMessageTupleRootNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_lastMessageTupleRootNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastMessageTupleRootNonce is a free data retrieval call binding the contract method 0x9a1cfb83.
//
// Solidity: function state_lastMessageTupleRootNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StateLastMessageTupleRootNonce() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StateLastMessageTupleRootNonce(&_QuantumGravityBridge.CallOpts)
}

// StateLastMessageTupleRootNonce is a free data retrieval call binding the contract method 0x9a1cfb83.
//
// Solidity: function state_lastMessageTupleRootNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StateLastMessageTupleRootNonce() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StateLastMessageTupleRootNonce(&_QuantumGravityBridge.CallOpts)
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

// StateMessageTupleRoots is a free data retrieval call binding the contract method 0xd8340ecf.
//
// Solidity: function state_messageTupleRoots(uint256 ) view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StateMessageTupleRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_messageTupleRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateMessageTupleRoots is a free data retrieval call binding the contract method 0xd8340ecf.
//
// Solidity: function state_messageTupleRoots(uint256 ) view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StateMessageTupleRoots(arg0 *big.Int) ([32]byte, error) {
	return _QuantumGravityBridge.Contract.StateMessageTupleRoots(&_QuantumGravityBridge.CallOpts, arg0)
}

// StateMessageTupleRoots is a free data retrieval call binding the contract method 0xd8340ecf.
//
// Solidity: function state_messageTupleRoots(uint256 ) view returns(bytes32)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StateMessageTupleRoots(arg0 *big.Int) ([32]byte, error) {
	return _QuantumGravityBridge.Contract.StateMessageTupleRoots(&_QuantumGravityBridge.CallOpts, arg0)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuantumGravityBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuantumGravityBridge *QuantumGravityBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.RenounceOwnership(&_QuantumGravityBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.RenounceOwnership(&_QuantumGravityBridge.TransactOpts)
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactor) RenounceOwnershipAfterExpiry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuantumGravityBridge.contract.Transact(opts, "renounceOwnershipAfterExpiry")
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_QuantumGravityBridge *QuantumGravityBridgeSession) RenounceOwnershipAfterExpiry() (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.RenounceOwnershipAfterExpiry(&_QuantumGravityBridge.TransactOpts)
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorSession) RenounceOwnershipAfterExpiry() (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.RenounceOwnershipAfterExpiry(&_QuantumGravityBridge.TransactOpts)
}

// SubmitMessageTupleRoot is a paid mutator transaction binding the contract method 0x030643d7.
//
// Solidity: function submitMessageTupleRoot(uint256 _nonce, bytes32 _messageTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactor) SubmitMessageTupleRoot(opts *bind.TransactOpts, _nonce *big.Int, _messageTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.contract.Transact(opts, "submitMessageTupleRoot", _nonce, _messageTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitMessageTupleRoot is a paid mutator transaction binding the contract method 0x030643d7.
//
// Solidity: function submitMessageTupleRoot(uint256 _nonce, bytes32 _messageTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeSession) SubmitMessageTupleRoot(_nonce *big.Int, _messageTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.SubmitMessageTupleRoot(&_QuantumGravityBridge.TransactOpts, _nonce, _messageTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitMessageTupleRoot is a paid mutator transaction binding the contract method 0x030643d7.
//
// Solidity: function submitMessageTupleRoot(uint256 _nonce, bytes32 _messageTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorSession) SubmitMessageTupleRoot(_nonce *big.Int, _messageTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.SubmitMessageTupleRoot(&_QuantumGravityBridge.TransactOpts, _nonce, _messageTupleRoot, _currentValidatorSet, _sigs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QuantumGravityBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.TransferOwnership(&_QuantumGravityBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.TransferOwnership(&_QuantumGravityBridge.TransactOpts, newOwner)
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

// QuantumGravityBridgeMessageTupleRootEventIterator is returned from FilterMessageTupleRootEvent and is used to iterate over the raw logs and unpacked data for MessageTupleRootEvent events raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeMessageTupleRootEventIterator struct {
	Event *QuantumGravityBridgeMessageTupleRootEvent // Event containing the contract specifics and raw log

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
func (it *QuantumGravityBridgeMessageTupleRootEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuantumGravityBridgeMessageTupleRootEvent)
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
		it.Event = new(QuantumGravityBridgeMessageTupleRootEvent)
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
func (it *QuantumGravityBridgeMessageTupleRootEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuantumGravityBridgeMessageTupleRootEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuantumGravityBridgeMessageTupleRootEvent represents a MessageTupleRootEvent event raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeMessageTupleRootEvent struct {
	Nonce            *big.Int
	MessageTupleRoot [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMessageTupleRootEvent is a free log retrieval operation binding the contract event 0x6f71145a387d176c99498fbbb78a98e85a46472eb5a4de9e06242e4edf4527e4.
//
// Solidity: event MessageTupleRootEvent(uint256 indexed nonce, bytes32 messageTupleRoot)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) FilterMessageTupleRootEvent(opts *bind.FilterOpts, nonce []*big.Int) (*QuantumGravityBridgeMessageTupleRootEventIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.FilterLogs(opts, "MessageTupleRootEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeMessageTupleRootEventIterator{contract: _QuantumGravityBridge.contract, event: "MessageTupleRootEvent", logs: logs, sub: sub}, nil
}

// WatchMessageTupleRootEvent is a free log subscription operation binding the contract event 0x6f71145a387d176c99498fbbb78a98e85a46472eb5a4de9e06242e4edf4527e4.
//
// Solidity: event MessageTupleRootEvent(uint256 indexed nonce, bytes32 messageTupleRoot)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) WatchMessageTupleRootEvent(opts *bind.WatchOpts, sink chan<- *QuantumGravityBridgeMessageTupleRootEvent, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.WatchLogs(opts, "MessageTupleRootEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuantumGravityBridgeMessageTupleRootEvent)
				if err := _QuantumGravityBridge.contract.UnpackLog(event, "MessageTupleRootEvent", log); err != nil {
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

// ParseMessageTupleRootEvent is a log parse operation binding the contract event 0x6f71145a387d176c99498fbbb78a98e85a46472eb5a4de9e06242e4edf4527e4.
//
// Solidity: event MessageTupleRootEvent(uint256 indexed nonce, bytes32 messageTupleRoot)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) ParseMessageTupleRootEvent(log types.Log) (*QuantumGravityBridgeMessageTupleRootEvent, error) {
	event := new(QuantumGravityBridgeMessageTupleRootEvent)
	if err := _QuantumGravityBridge.contract.UnpackLog(event, "MessageTupleRootEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuantumGravityBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeOwnershipTransferredIterator struct {
	Event *QuantumGravityBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QuantumGravityBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuantumGravityBridgeOwnershipTransferred)
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
		it.Event = new(QuantumGravityBridgeOwnershipTransferred)
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
func (it *QuantumGravityBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuantumGravityBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuantumGravityBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the QuantumGravityBridge contract.
type QuantumGravityBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QuantumGravityBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QuantumGravityBridgeOwnershipTransferredIterator{contract: _QuantumGravityBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QuantumGravityBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuantumGravityBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuantumGravityBridgeOwnershipTransferred)
				if err := _QuantumGravityBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuantumGravityBridge *QuantumGravityBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*QuantumGravityBridgeOwnershipTransferred, error) {
	event := new(QuantumGravityBridgeOwnershipTransferred)
	if err := _QuantumGravityBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
