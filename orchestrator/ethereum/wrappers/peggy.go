// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PeggyABI is the input ABI used to generate the binding from.
const PeggyABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_peggyId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"SendToCosmosEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"TransactionBatchExecutedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_newValsetNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"ValsetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Address\",\"type\":\"address\"}],\"name\":\"lastBatchNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"seenTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendToCosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"state_lastBatchNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastEventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_peggyId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_currentValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_currentPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_currentValsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_destinations\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"submitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_currentValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_currentPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_theHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"}],\"name\":\"testCheckValidatorSignatures\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_peggyId\",\"type\":\"bytes32\"}],\"name\":\"testMakeCheckpoint\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenSymbols\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_newPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_newValsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_currentValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_currentPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_currentValsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"updateValset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Peggy is an auto generated Go binding around an Ethereum contract.
type Peggy struct {
	PeggyCaller     // Read-only binding to the contract
	PeggyTransactor // Write-only binding to the contract
	PeggyFilterer   // Log filterer for contract events
}

// PeggyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeggyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeggyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeggyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeggyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeggyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeggySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeggySession struct {
	Contract     *Peggy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PeggyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeggyCallerSession struct {
	Contract *PeggyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PeggyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeggyTransactorSession struct {
	Contract     *PeggyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PeggyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeggyRaw struct {
	Contract *Peggy // Generic contract binding to access the raw methods on
}

// PeggyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeggyCallerRaw struct {
	Contract *PeggyCaller // Generic read-only contract binding to access the raw methods on
}

// PeggyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeggyTransactorRaw struct {
	Contract *PeggyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeggy creates a new instance of Peggy, bound to a specific deployed contract.
func NewPeggy(address common.Address, backend bind.ContractBackend) (*Peggy, error) {
	contract, err := bindPeggy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Peggy{PeggyCaller: PeggyCaller{contract: contract}, PeggyTransactor: PeggyTransactor{contract: contract}, PeggyFilterer: PeggyFilterer{contract: contract}}, nil
}

// NewPeggyCaller creates a new read-only instance of Peggy, bound to a specific deployed contract.
func NewPeggyCaller(address common.Address, caller bind.ContractCaller) (*PeggyCaller, error) {
	contract, err := bindPeggy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeggyCaller{contract: contract}, nil
}

// NewPeggyTransactor creates a new write-only instance of Peggy, bound to a specific deployed contract.
func NewPeggyTransactor(address common.Address, transactor bind.ContractTransactor) (*PeggyTransactor, error) {
	contract, err := bindPeggy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeggyTransactor{contract: contract}, nil
}

// NewPeggyFilterer creates a new log filterer instance of Peggy, bound to a specific deployed contract.
func NewPeggyFilterer(address common.Address, filterer bind.ContractFilterer) (*PeggyFilterer, error) {
	contract, err := bindPeggy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeggyFilterer{contract: contract}, nil
}

// bindPeggy binds a generic wrapper to an already deployed contract.
func bindPeggy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PeggyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Peggy *PeggyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Peggy.Contract.PeggyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Peggy *PeggyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peggy.Contract.PeggyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Peggy *PeggyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Peggy.Contract.PeggyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Peggy *PeggyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Peggy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Peggy *PeggyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peggy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Peggy *PeggyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Peggy.Contract.contract.Transact(opts, method, params...)
}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Peggy *PeggyCaller) LastBatchNonce(opts *bind.CallOpts, _erc20Address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "lastBatchNonce", _erc20Address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Peggy *PeggySession) LastBatchNonce(_erc20Address common.Address) (*big.Int, error) {
	return _Peggy.Contract.LastBatchNonce(&_Peggy.CallOpts, _erc20Address)
}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Peggy *PeggyCallerSession) LastBatchNonce(_erc20Address common.Address) (*big.Int, error) {
	return _Peggy.Contract.LastBatchNonce(&_Peggy.CallOpts, _erc20Address)
}

// SeenTokens is a free data retrieval call binding the contract method 0x13100091.
//
// Solidity: function seenTokens(address ) view returns(bool)
func (_Peggy *PeggyCaller) SeenTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "seenTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeenTokens is a free data retrieval call binding the contract method 0x13100091.
//
// Solidity: function seenTokens(address ) view returns(bool)
func (_Peggy *PeggySession) SeenTokens(arg0 common.Address) (bool, error) {
	return _Peggy.Contract.SeenTokens(&_Peggy.CallOpts, arg0)
}

// SeenTokens is a free data retrieval call binding the contract method 0x13100091.
//
// Solidity: function seenTokens(address ) view returns(bool)
func (_Peggy *PeggyCallerSession) SeenTokens(arg0 common.Address) (bool, error) {
	return _Peggy.Contract.SeenTokens(&_Peggy.CallOpts, arg0)
}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Peggy *PeggyCaller) StateLastBatchNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "state_lastBatchNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Peggy *PeggySession) StateLastBatchNonces(arg0 common.Address) (*big.Int, error) {
	return _Peggy.Contract.StateLastBatchNonces(&_Peggy.CallOpts, arg0)
}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Peggy *PeggyCallerSession) StateLastBatchNonces(arg0 common.Address) (*big.Int, error) {
	return _Peggy.Contract.StateLastBatchNonces(&_Peggy.CallOpts, arg0)
}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Peggy *PeggyCaller) StateLastEventNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "state_lastEventNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Peggy *PeggySession) StateLastEventNonce() (*big.Int, error) {
	return _Peggy.Contract.StateLastEventNonce(&_Peggy.CallOpts)
}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Peggy *PeggyCallerSession) StateLastEventNonce() (*big.Int, error) {
	return _Peggy.Contract.StateLastEventNonce(&_Peggy.CallOpts)
}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Peggy *PeggyCaller) StateLastValsetCheckpoint(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "state_lastValsetCheckpoint")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Peggy *PeggySession) StateLastValsetCheckpoint() ([32]byte, error) {
	return _Peggy.Contract.StateLastValsetCheckpoint(&_Peggy.CallOpts)
}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Peggy *PeggyCallerSession) StateLastValsetCheckpoint() ([32]byte, error) {
	return _Peggy.Contract.StateLastValsetCheckpoint(&_Peggy.CallOpts)
}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Peggy *PeggyCaller) StateLastValsetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "state_lastValsetNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Peggy *PeggySession) StateLastValsetNonce() (*big.Int, error) {
	return _Peggy.Contract.StateLastValsetNonce(&_Peggy.CallOpts)
}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Peggy *PeggyCallerSession) StateLastValsetNonce() (*big.Int, error) {
	return _Peggy.Contract.StateLastValsetNonce(&_Peggy.CallOpts)
}

// StatePeggyId is a free data retrieval call binding the contract method 0x69dd3908.
//
// Solidity: function state_peggyId() view returns(bytes32)
func (_Peggy *PeggyCaller) StatePeggyId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "state_peggyId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StatePeggyId is a free data retrieval call binding the contract method 0x69dd3908.
//
// Solidity: function state_peggyId() view returns(bytes32)
func (_Peggy *PeggySession) StatePeggyId() ([32]byte, error) {
	return _Peggy.Contract.StatePeggyId(&_Peggy.CallOpts)
}

// StatePeggyId is a free data retrieval call binding the contract method 0x69dd3908.
//
// Solidity: function state_peggyId() view returns(bytes32)
func (_Peggy *PeggyCallerSession) StatePeggyId() ([32]byte, error) {
	return _Peggy.Contract.StatePeggyId(&_Peggy.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Peggy *PeggyCaller) StatePowerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "state_powerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Peggy *PeggySession) StatePowerThreshold() (*big.Int, error) {
	return _Peggy.Contract.StatePowerThreshold(&_Peggy.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Peggy *PeggyCallerSession) StatePowerThreshold() (*big.Int, error) {
	return _Peggy.Contract.StatePowerThreshold(&_Peggy.CallOpts)
}

// TestCheckValidatorSignatures is a free data retrieval call binding the contract method 0xdb7c4e57.
//
// Solidity: function testCheckValidatorSignatures(address[] _currentValidators, uint256[] _currentPowers, uint8[] _v, bytes32[] _r, bytes32[] _s, bytes32 _theHash, uint256 _powerThreshold) pure returns()
func (_Peggy *PeggyCaller) TestCheckValidatorSignatures(opts *bind.CallOpts, _currentValidators []common.Address, _currentPowers []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _theHash [32]byte, _powerThreshold *big.Int) error {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "testCheckValidatorSignatures", _currentValidators, _currentPowers, _v, _r, _s, _theHash, _powerThreshold)

	if err != nil {
		return err
	}

	return err

}

// TestCheckValidatorSignatures is a free data retrieval call binding the contract method 0xdb7c4e57.
//
// Solidity: function testCheckValidatorSignatures(address[] _currentValidators, uint256[] _currentPowers, uint8[] _v, bytes32[] _r, bytes32[] _s, bytes32 _theHash, uint256 _powerThreshold) pure returns()
func (_Peggy *PeggySession) TestCheckValidatorSignatures(_currentValidators []common.Address, _currentPowers []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _theHash [32]byte, _powerThreshold *big.Int) error {
	return _Peggy.Contract.TestCheckValidatorSignatures(&_Peggy.CallOpts, _currentValidators, _currentPowers, _v, _r, _s, _theHash, _powerThreshold)
}

// TestCheckValidatorSignatures is a free data retrieval call binding the contract method 0xdb7c4e57.
//
// Solidity: function testCheckValidatorSignatures(address[] _currentValidators, uint256[] _currentPowers, uint8[] _v, bytes32[] _r, bytes32[] _s, bytes32 _theHash, uint256 _powerThreshold) pure returns()
func (_Peggy *PeggyCallerSession) TestCheckValidatorSignatures(_currentValidators []common.Address, _currentPowers []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _theHash [32]byte, _powerThreshold *big.Int) error {
	return _Peggy.Contract.TestCheckValidatorSignatures(&_Peggy.CallOpts, _currentValidators, _currentPowers, _v, _r, _s, _theHash, _powerThreshold)
}

// TestMakeCheckpoint is a free data retrieval call binding the contract method 0xc227c30b.
//
// Solidity: function testMakeCheckpoint(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) pure returns()
func (_Peggy *PeggyCaller) TestMakeCheckpoint(opts *bind.CallOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) error {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "testMakeCheckpoint", _validators, _powers, _valsetNonce, _peggyId)

	if err != nil {
		return err
	}

	return err

}

// TestMakeCheckpoint is a free data retrieval call binding the contract method 0xc227c30b.
//
// Solidity: function testMakeCheckpoint(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) pure returns()
func (_Peggy *PeggySession) TestMakeCheckpoint(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) error {
	return _Peggy.Contract.TestMakeCheckpoint(&_Peggy.CallOpts, _validators, _powers, _valsetNonce, _peggyId)
}

// TestMakeCheckpoint is a free data retrieval call binding the contract method 0xc227c30b.
//
// Solidity: function testMakeCheckpoint(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) pure returns()
func (_Peggy *PeggyCallerSession) TestMakeCheckpoint(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) error {
	return _Peggy.Contract.TestMakeCheckpoint(&_Peggy.CallOpts, _validators, _powers, _valsetNonce, _peggyId)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint8)
func (_Peggy *PeggyCaller) TokenDecimals(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "tokenDecimals", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint8)
func (_Peggy *PeggySession) TokenDecimals(arg0 common.Address) (uint8, error) {
	return _Peggy.Contract.TokenDecimals(&_Peggy.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint8)
func (_Peggy *PeggyCallerSession) TokenDecimals(arg0 common.Address) (uint8, error) {
	return _Peggy.Contract.TokenDecimals(&_Peggy.CallOpts, arg0)
}

// TokenNames is a free data retrieval call binding the contract method 0xa8fc75e1.
//
// Solidity: function tokenNames(address ) view returns(string)
func (_Peggy *PeggyCaller) TokenNames(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "tokenNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenNames is a free data retrieval call binding the contract method 0xa8fc75e1.
//
// Solidity: function tokenNames(address ) view returns(string)
func (_Peggy *PeggySession) TokenNames(arg0 common.Address) (string, error) {
	return _Peggy.Contract.TokenNames(&_Peggy.CallOpts, arg0)
}

// TokenNames is a free data retrieval call binding the contract method 0xa8fc75e1.
//
// Solidity: function tokenNames(address ) view returns(string)
func (_Peggy *PeggyCallerSession) TokenNames(arg0 common.Address) (string, error) {
	return _Peggy.Contract.TokenNames(&_Peggy.CallOpts, arg0)
}

// TokenSymbols is a free data retrieval call binding the contract method 0xfb0b2b36.
//
// Solidity: function tokenSymbols(address ) view returns(string)
func (_Peggy *PeggyCaller) TokenSymbols(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "tokenSymbols", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenSymbols is a free data retrieval call binding the contract method 0xfb0b2b36.
//
// Solidity: function tokenSymbols(address ) view returns(string)
func (_Peggy *PeggySession) TokenSymbols(arg0 common.Address) (string, error) {
	return _Peggy.Contract.TokenSymbols(&_Peggy.CallOpts, arg0)
}

// TokenSymbols is a free data retrieval call binding the contract method 0xfb0b2b36.
//
// Solidity: function tokenSymbols(address ) view returns(string)
func (_Peggy *PeggyCallerSession) TokenSymbols(arg0 common.Address) (string, error) {
	return _Peggy.Contract.TokenSymbols(&_Peggy.CallOpts, arg0)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x5d2428fa.
//
// Solidity: function sendToCosmos(address _tokenContract, address _destination, uint256 _amount) returns()
func (_Peggy *PeggyTransactor) SendToCosmos(opts *bind.TransactOpts, _tokenContract common.Address, _destination common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Peggy.contract.Transact(opts, "sendToCosmos", _tokenContract, _destination, _amount)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x5d2428fa.
//
// Solidity: function sendToCosmos(address _tokenContract, address _destination, uint256 _amount) returns()
func (_Peggy *PeggySession) SendToCosmos(_tokenContract common.Address, _destination common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Peggy.Contract.SendToCosmos(&_Peggy.TransactOpts, _tokenContract, _destination, _amount)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x5d2428fa.
//
// Solidity: function sendToCosmos(address _tokenContract, address _destination, uint256 _amount) returns()
func (_Peggy *PeggyTransactorSession) SendToCosmos(_tokenContract common.Address, _destination common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Peggy.Contract.SendToCosmos(&_Peggy.TransactOpts, _tokenContract, _destination, _amount)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0xad131ec5.
//
// Solidity: function submitBatch(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract) returns()
func (_Peggy *PeggyTransactor) SubmitBatch(opts *bind.TransactOpts, _currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address) (*types.Transaction, error) {
	return _Peggy.contract.Transact(opts, "submitBatch", _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0xad131ec5.
//
// Solidity: function submitBatch(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract) returns()
func (_Peggy *PeggySession) SubmitBatch(_currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address) (*types.Transaction, error) {
	return _Peggy.Contract.SubmitBatch(&_Peggy.TransactOpts, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0xad131ec5.
//
// Solidity: function submitBatch(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract) returns()
func (_Peggy *PeggyTransactorSession) SubmitBatch(_currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address) (*types.Transaction, error) {
	return _Peggy.Contract.SubmitBatch(&_Peggy.TransactOpts, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xe3cb9f62.
//
// Solidity: function updateValset(address[] _newValidators, uint256[] _newPowers, uint256 _newValsetNonce, address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Peggy *PeggyTransactor) UpdateValset(opts *bind.TransactOpts, _newValidators []common.Address, _newPowers []*big.Int, _newValsetNonce *big.Int, _currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Peggy.contract.Transact(opts, "updateValset", _newValidators, _newPowers, _newValsetNonce, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xe3cb9f62.
//
// Solidity: function updateValset(address[] _newValidators, uint256[] _newPowers, uint256 _newValsetNonce, address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Peggy *PeggySession) UpdateValset(_newValidators []common.Address, _newPowers []*big.Int, _newValsetNonce *big.Int, _currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Peggy.Contract.UpdateValset(&_Peggy.TransactOpts, _newValidators, _newPowers, _newValsetNonce, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xe3cb9f62.
//
// Solidity: function updateValset(address[] _newValidators, uint256[] _newPowers, uint256 _newValsetNonce, address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Peggy *PeggyTransactorSession) UpdateValset(_newValidators []common.Address, _newPowers []*big.Int, _newValsetNonce *big.Int, _currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Peggy.Contract.UpdateValset(&_Peggy.TransactOpts, _newValidators, _newPowers, _newValsetNonce, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s)
}

// PeggySendToCosmosEventIterator is returned from FilterSendToCosmosEvent and is used to iterate over the raw logs and unpacked data for SendToCosmosEvent events raised by the Peggy contract.
type PeggySendToCosmosEventIterator struct {
	Event *PeggySendToCosmosEvent // Event containing the contract specifics and raw log

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
func (it *PeggySendToCosmosEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggySendToCosmosEvent)
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
		it.Event = new(PeggySendToCosmosEvent)
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
func (it *PeggySendToCosmosEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggySendToCosmosEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggySendToCosmosEvent represents a SendToCosmosEvent event raised by the Peggy contract.
type PeggySendToCosmosEvent struct {
	TokenContract common.Address
	Sender        common.Address
	Destination   common.Address
	Amount        *big.Int
	EventNonce    *big.Int
	Name          string
	Symbol        string
	Decimals      uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSendToCosmosEvent is a free log retrieval operation binding the contract event 0x5b0cfcd9629f2864d66b907e3625f822058529e928617605883190ad34ecb96d.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, address indexed _destination, uint256 _amount, uint256 _eventNonce, string _name, string _symbol, uint8 _decimals)
func (_Peggy *PeggyFilterer) FilterSendToCosmosEvent(opts *bind.FilterOpts, _tokenContract []common.Address, _sender []common.Address, _destination []common.Address) (*PeggySendToCosmosEventIterator, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _destinationRule []interface{}
	for _, _destinationItem := range _destination {
		_destinationRule = append(_destinationRule, _destinationItem)
	}

	logs, sub, err := _Peggy.contract.FilterLogs(opts, "SendToCosmosEvent", _tokenContractRule, _senderRule, _destinationRule)
	if err != nil {
		return nil, err
	}
	return &PeggySendToCosmosEventIterator{contract: _Peggy.contract, event: "SendToCosmosEvent", logs: logs, sub: sub}, nil
}

// WatchSendToCosmosEvent is a free log subscription operation binding the contract event 0x5b0cfcd9629f2864d66b907e3625f822058529e928617605883190ad34ecb96d.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, address indexed _destination, uint256 _amount, uint256 _eventNonce, string _name, string _symbol, uint8 _decimals)
func (_Peggy *PeggyFilterer) WatchSendToCosmosEvent(opts *bind.WatchOpts, sink chan<- *PeggySendToCosmosEvent, _tokenContract []common.Address, _sender []common.Address, _destination []common.Address) (event.Subscription, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _destinationRule []interface{}
	for _, _destinationItem := range _destination {
		_destinationRule = append(_destinationRule, _destinationItem)
	}

	logs, sub, err := _Peggy.contract.WatchLogs(opts, "SendToCosmosEvent", _tokenContractRule, _senderRule, _destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggySendToCosmosEvent)
				if err := _Peggy.contract.UnpackLog(event, "SendToCosmosEvent", log); err != nil {
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

// ParseSendToCosmosEvent is a log parse operation binding the contract event 0x5b0cfcd9629f2864d66b907e3625f822058529e928617605883190ad34ecb96d.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, address indexed _destination, uint256 _amount, uint256 _eventNonce, string _name, string _symbol, uint8 _decimals)
func (_Peggy *PeggyFilterer) ParseSendToCosmosEvent(log types.Log) (*PeggySendToCosmosEvent, error) {
	event := new(PeggySendToCosmosEvent)
	if err := _Peggy.contract.UnpackLog(event, "SendToCosmosEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggyTransactionBatchExecutedEventIterator is returned from FilterTransactionBatchExecutedEvent and is used to iterate over the raw logs and unpacked data for TransactionBatchExecutedEvent events raised by the Peggy contract.
type PeggyTransactionBatchExecutedEventIterator struct {
	Event *PeggyTransactionBatchExecutedEvent // Event containing the contract specifics and raw log

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
func (it *PeggyTransactionBatchExecutedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggyTransactionBatchExecutedEvent)
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
		it.Event = new(PeggyTransactionBatchExecutedEvent)
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
func (it *PeggyTransactionBatchExecutedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggyTransactionBatchExecutedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggyTransactionBatchExecutedEvent represents a TransactionBatchExecutedEvent event raised by the Peggy contract.
type PeggyTransactionBatchExecutedEvent struct {
	BatchNonce *big.Int
	Token      common.Address
	EventNonce *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionBatchExecutedEvent is a free log retrieval operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Peggy *PeggyFilterer) FilterTransactionBatchExecutedEvent(opts *bind.FilterOpts, _batchNonce []*big.Int, _token []common.Address) (*PeggyTransactionBatchExecutedEventIterator, error) {

	var _batchNonceRule []interface{}
	for _, _batchNonceItem := range _batchNonce {
		_batchNonceRule = append(_batchNonceRule, _batchNonceItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Peggy.contract.FilterLogs(opts, "TransactionBatchExecutedEvent", _batchNonceRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &PeggyTransactionBatchExecutedEventIterator{contract: _Peggy.contract, event: "TransactionBatchExecutedEvent", logs: logs, sub: sub}, nil
}

// WatchTransactionBatchExecutedEvent is a free log subscription operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Peggy *PeggyFilterer) WatchTransactionBatchExecutedEvent(opts *bind.WatchOpts, sink chan<- *PeggyTransactionBatchExecutedEvent, _batchNonce []*big.Int, _token []common.Address) (event.Subscription, error) {

	var _batchNonceRule []interface{}
	for _, _batchNonceItem := range _batchNonce {
		_batchNonceRule = append(_batchNonceRule, _batchNonceItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Peggy.contract.WatchLogs(opts, "TransactionBatchExecutedEvent", _batchNonceRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggyTransactionBatchExecutedEvent)
				if err := _Peggy.contract.UnpackLog(event, "TransactionBatchExecutedEvent", log); err != nil {
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

// ParseTransactionBatchExecutedEvent is a log parse operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Peggy *PeggyFilterer) ParseTransactionBatchExecutedEvent(log types.Log) (*PeggyTransactionBatchExecutedEvent, error) {
	event := new(PeggyTransactionBatchExecutedEvent)
	if err := _Peggy.contract.UnpackLog(event, "TransactionBatchExecutedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggyValsetUpdatedEventIterator is returned from FilterValsetUpdatedEvent and is used to iterate over the raw logs and unpacked data for ValsetUpdatedEvent events raised by the Peggy contract.
type PeggyValsetUpdatedEventIterator struct {
	Event *PeggyValsetUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *PeggyValsetUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggyValsetUpdatedEvent)
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
		it.Event = new(PeggyValsetUpdatedEvent)
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
func (it *PeggyValsetUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggyValsetUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggyValsetUpdatedEvent represents a ValsetUpdatedEvent event raised by the Peggy contract.
type PeggyValsetUpdatedEvent struct {
	NewValsetNonce *big.Int
	Validators     []common.Address
	Powers         []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValsetUpdatedEvent is a free log retrieval operation binding the contract event 0xc6d025c076bafcdd040f00632d5e280b3a5188963f110f8c70c4f810184b30f3.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, address[] _validators, uint256[] _powers)
func (_Peggy *PeggyFilterer) FilterValsetUpdatedEvent(opts *bind.FilterOpts, _newValsetNonce []*big.Int) (*PeggyValsetUpdatedEventIterator, error) {

	var _newValsetNonceRule []interface{}
	for _, _newValsetNonceItem := range _newValsetNonce {
		_newValsetNonceRule = append(_newValsetNonceRule, _newValsetNonceItem)
	}

	logs, sub, err := _Peggy.contract.FilterLogs(opts, "ValsetUpdatedEvent", _newValsetNonceRule)
	if err != nil {
		return nil, err
	}
	return &PeggyValsetUpdatedEventIterator{contract: _Peggy.contract, event: "ValsetUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchValsetUpdatedEvent is a free log subscription operation binding the contract event 0xc6d025c076bafcdd040f00632d5e280b3a5188963f110f8c70c4f810184b30f3.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, address[] _validators, uint256[] _powers)
func (_Peggy *PeggyFilterer) WatchValsetUpdatedEvent(opts *bind.WatchOpts, sink chan<- *PeggyValsetUpdatedEvent, _newValsetNonce []*big.Int) (event.Subscription, error) {

	var _newValsetNonceRule []interface{}
	for _, _newValsetNonceItem := range _newValsetNonce {
		_newValsetNonceRule = append(_newValsetNonceRule, _newValsetNonceItem)
	}

	logs, sub, err := _Peggy.contract.WatchLogs(opts, "ValsetUpdatedEvent", _newValsetNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggyValsetUpdatedEvent)
				if err := _Peggy.contract.UnpackLog(event, "ValsetUpdatedEvent", log); err != nil {
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

// ParseValsetUpdatedEvent is a log parse operation binding the contract event 0xc6d025c076bafcdd040f00632d5e280b3a5188963f110f8c70c4f810184b30f3.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, address[] _validators, uint256[] _powers)
func (_Peggy *PeggyFilterer) ParseValsetUpdatedEvent(log types.Log) (*PeggyValsetUpdatedEvent, error) {
	event := new(PeggyValsetUpdatedEvent)
	if err := _Peggy.contract.UnpackLog(event, "ValsetUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
