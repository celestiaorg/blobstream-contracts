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
