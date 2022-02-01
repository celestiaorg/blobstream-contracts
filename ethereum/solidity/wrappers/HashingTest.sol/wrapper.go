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

// HashingTestMetaData contains all meta data concerning the HashingTest contract.
var HashingTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_peggyId\",\"type\":\"bytes32\"}],\"name\":\"CheckpointHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"}],\"name\":\"JustSaveEverything\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"}],\"name\":\"JustSaveEverythingAgain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_powers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cd16f185": "CheckpointHash(address[],uint256[],uint256,uint256,address,bytes32)",
		"884403e2": "JustSaveEverything(address[],uint256[],uint256)",
		"715dff7e": "JustSaveEverythingAgain(address[],uint256[],uint256)",
		"d32e81a5": "lastCheckpoint()",
		"ccf0e74c": "state_nonce()",
		"2b939281": "state_powers(uint256)",
		"2afbb62e": "state_validators(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610602806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063884403e21161005b578063884403e2146100d3578063ccf0e74c146100e8578063cd16f185146100f1578063d32e81a51461010457600080fd5b80632afbb62e146100825780632b939281146100b2578063715dff7e146100d3575b600080fd5b610095610090366004610290565b61010d565b6040516001600160a01b0390911681526020015b60405180910390f35b6100c56100c0366004610290565b610137565b6040519081526020016100a9565b6100e66100e13660046103fd565b610158565b005b6100c560035481565b6100e66100ff36600461046a565b610187565b6100c560005481565b6001818154811061011d57600080fd5b6000918252602090912001546001600160a01b0316905081565b6002818154811061014757600080fd5b600091825260209091200154905081565b825161016b9060019060208601906101db565b50815161017f906002906020850190610240565b506003555050565b6040516918da1958dadc1bda5b9d60b21b906000906101b6908490849089908c908c908b908b90602001610535565b60408051601f1981840301815291905280516020909101206000555050505050505050565b828054828255906000526020600020908101928215610230579160200282015b8281111561023057825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906101fb565b5061023c92915061027b565b5090565b828054828255906000526020600020908101928215610230579160200282015b82811115610230578251825591602001919060010190610260565b5b8082111561023c576000815560010161027c565b6000602082840312156102a257600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156102e8576102e86102a9565b604052919050565b600067ffffffffffffffff82111561030a5761030a6102a9565b5060051b60200190565b80356001600160a01b038116811461032b57600080fd5b919050565b600082601f83011261034157600080fd5b81356020610356610351836102f0565b6102bf565b82815260059290921b8401810191818101908684111561037557600080fd5b8286015b848110156103975761038a81610314565b8352918301918301610379565b509695505050505050565b600082601f8301126103b357600080fd5b813560206103c3610351836102f0565b82815260059290921b840181019181810190868411156103e257600080fd5b8286015b8481101561039757803583529183019183016103e6565b60008060006060848603121561041257600080fd5b833567ffffffffffffffff8082111561042a57600080fd5b61043687838801610330565b9450602086013591508082111561044c57600080fd5b50610459868287016103a2565b925050604084013590509250925092565b60008060008060008060c0878903121561048357600080fd5b863567ffffffffffffffff8082111561049b57600080fd5b6104a78a838b01610330565b975060208901359150808211156104bd57600080fd5b506104ca89828a016103a2565b95505060408701359350606087013592506104e760808801610314565b915060a087013590509295509295509295565b600081518084526020808501945080840160005b8381101561052a5781518752958201959082019060010161050e565b509495945050505050565b600060e082018983526020898185015288604085015260e0606085015281885180845261010086019150828a01935060005b8181101561058c5784516001600160a01b031683529383019391830191600101610567565b505084810360808601526105a081896104fa565b93505050508360a08301526105c060c08301846001600160a01b03169052565b9897505050505050505056fea26469706673582212201b245be44a52063d460dfb281e910a9a7b5ad835700c39c88eead815fb49472064736f6c634300080b0033",
}

// HashingTestABI is the input ABI used to generate the binding from.
// Deprecated: Use HashingTestMetaData.ABI instead.
var HashingTestABI = HashingTestMetaData.ABI

// Deprecated: Use HashingTestMetaData.Sigs instead.
// HashingTestFuncSigs maps the 4-byte function signature to its string representation.
var HashingTestFuncSigs = HashingTestMetaData.Sigs

// HashingTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HashingTestMetaData.Bin instead.
var HashingTestBin = HashingTestMetaData.Bin

// DeployHashingTest deploys a new Ethereum contract, binding an instance of HashingTest to it.
func DeployHashingTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashingTest, error) {
	parsed, err := HashingTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HashingTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashingTest{HashingTestCaller: HashingTestCaller{contract: contract}, HashingTestTransactor: HashingTestTransactor{contract: contract}, HashingTestFilterer: HashingTestFilterer{contract: contract}}, nil
}

// HashingTest is an auto generated Go binding around an Ethereum contract.
type HashingTest struct {
	HashingTestCaller     // Read-only binding to the contract
	HashingTestTransactor // Write-only binding to the contract
	HashingTestFilterer   // Log filterer for contract events
}

// HashingTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashingTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashingTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashingTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashingTestSession struct {
	Contract     *HashingTest      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashingTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashingTestCallerSession struct {
	Contract *HashingTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HashingTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashingTestTransactorSession struct {
	Contract     *HashingTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HashingTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashingTestRaw struct {
	Contract *HashingTest // Generic contract binding to access the raw methods on
}

// HashingTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashingTestCallerRaw struct {
	Contract *HashingTestCaller // Generic read-only contract binding to access the raw methods on
}

// HashingTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashingTestTransactorRaw struct {
	Contract *HashingTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashingTest creates a new instance of HashingTest, bound to a specific deployed contract.
func NewHashingTest(address common.Address, backend bind.ContractBackend) (*HashingTest, error) {
	contract, err := bindHashingTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashingTest{HashingTestCaller: HashingTestCaller{contract: contract}, HashingTestTransactor: HashingTestTransactor{contract: contract}, HashingTestFilterer: HashingTestFilterer{contract: contract}}, nil
}

// NewHashingTestCaller creates a new read-only instance of HashingTest, bound to a specific deployed contract.
func NewHashingTestCaller(address common.Address, caller bind.ContractCaller) (*HashingTestCaller, error) {
	contract, err := bindHashingTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashingTestCaller{contract: contract}, nil
}

// NewHashingTestTransactor creates a new write-only instance of HashingTest, bound to a specific deployed contract.
func NewHashingTestTransactor(address common.Address, transactor bind.ContractTransactor) (*HashingTestTransactor, error) {
	contract, err := bindHashingTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashingTestTransactor{contract: contract}, nil
}

// NewHashingTestFilterer creates a new log filterer instance of HashingTest, bound to a specific deployed contract.
func NewHashingTestFilterer(address common.Address, filterer bind.ContractFilterer) (*HashingTestFilterer, error) {
	contract, err := bindHashingTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashingTestFilterer{contract: contract}, nil
}

// bindHashingTest binds a generic wrapper to an already deployed contract.
func bindHashingTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashingTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashingTest *HashingTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashingTest.Contract.HashingTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashingTest *HashingTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashingTest.Contract.HashingTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashingTest *HashingTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashingTest.Contract.HashingTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashingTest *HashingTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashingTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashingTest *HashingTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashingTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashingTest *HashingTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashingTest.Contract.contract.Transact(opts, method, params...)
}

// LastCheckpoint is a free data retrieval call binding the contract method 0xd32e81a5.
//
// Solidity: function lastCheckpoint() view returns(bytes32)
func (_HashingTest *HashingTestCaller) LastCheckpoint(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HashingTest.contract.Call(opts, &out, "lastCheckpoint")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastCheckpoint is a free data retrieval call binding the contract method 0xd32e81a5.
//
// Solidity: function lastCheckpoint() view returns(bytes32)
func (_HashingTest *HashingTestSession) LastCheckpoint() ([32]byte, error) {
	return _HashingTest.Contract.LastCheckpoint(&_HashingTest.CallOpts)
}

// LastCheckpoint is a free data retrieval call binding the contract method 0xd32e81a5.
//
// Solidity: function lastCheckpoint() view returns(bytes32)
func (_HashingTest *HashingTestCallerSession) LastCheckpoint() ([32]byte, error) {
	return _HashingTest.Contract.LastCheckpoint(&_HashingTest.CallOpts)
}

// StateNonce is a free data retrieval call binding the contract method 0xccf0e74c.
//
// Solidity: function state_nonce() view returns(uint256)
func (_HashingTest *HashingTestCaller) StateNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HashingTest.contract.Call(opts, &out, "state_nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateNonce is a free data retrieval call binding the contract method 0xccf0e74c.
//
// Solidity: function state_nonce() view returns(uint256)
func (_HashingTest *HashingTestSession) StateNonce() (*big.Int, error) {
	return _HashingTest.Contract.StateNonce(&_HashingTest.CallOpts)
}

// StateNonce is a free data retrieval call binding the contract method 0xccf0e74c.
//
// Solidity: function state_nonce() view returns(uint256)
func (_HashingTest *HashingTestCallerSession) StateNonce() (*big.Int, error) {
	return _HashingTest.Contract.StateNonce(&_HashingTest.CallOpts)
}

// StatePowers is a free data retrieval call binding the contract method 0x2b939281.
//
// Solidity: function state_powers(uint256 ) view returns(uint256)
func (_HashingTest *HashingTestCaller) StatePowers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HashingTest.contract.Call(opts, &out, "state_powers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatePowers is a free data retrieval call binding the contract method 0x2b939281.
//
// Solidity: function state_powers(uint256 ) view returns(uint256)
func (_HashingTest *HashingTestSession) StatePowers(arg0 *big.Int) (*big.Int, error) {
	return _HashingTest.Contract.StatePowers(&_HashingTest.CallOpts, arg0)
}

// StatePowers is a free data retrieval call binding the contract method 0x2b939281.
//
// Solidity: function state_powers(uint256 ) view returns(uint256)
func (_HashingTest *HashingTestCallerSession) StatePowers(arg0 *big.Int) (*big.Int, error) {
	return _HashingTest.Contract.StatePowers(&_HashingTest.CallOpts, arg0)
}

// StateValidators is a free data retrieval call binding the contract method 0x2afbb62e.
//
// Solidity: function state_validators(uint256 ) view returns(address)
func (_HashingTest *HashingTestCaller) StateValidators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _HashingTest.contract.Call(opts, &out, "state_validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateValidators is a free data retrieval call binding the contract method 0x2afbb62e.
//
// Solidity: function state_validators(uint256 ) view returns(address)
func (_HashingTest *HashingTestSession) StateValidators(arg0 *big.Int) (common.Address, error) {
	return _HashingTest.Contract.StateValidators(&_HashingTest.CallOpts, arg0)
}

// StateValidators is a free data retrieval call binding the contract method 0x2afbb62e.
//
// Solidity: function state_validators(uint256 ) view returns(address)
func (_HashingTest *HashingTestCallerSession) StateValidators(arg0 *big.Int) (common.Address, error) {
	return _HashingTest.Contract.StateValidators(&_HashingTest.CallOpts, arg0)
}

// CheckpointHash is a paid mutator transaction binding the contract method 0xcd16f185.
//
// Solidity: function CheckpointHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, uint256 _rewardAmount, address _rewardToken, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestTransactor) CheckpointHash(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _rewardAmount *big.Int, _rewardToken common.Address, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.contract.Transact(opts, "CheckpointHash", _validators, _powers, _valsetNonce, _rewardAmount, _rewardToken, _peggyId)
}

// CheckpointHash is a paid mutator transaction binding the contract method 0xcd16f185.
//
// Solidity: function CheckpointHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, uint256 _rewardAmount, address _rewardToken, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestSession) CheckpointHash(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _rewardAmount *big.Int, _rewardToken common.Address, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.CheckpointHash(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _rewardAmount, _rewardToken, _peggyId)
}

// CheckpointHash is a paid mutator transaction binding the contract method 0xcd16f185.
//
// Solidity: function CheckpointHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, uint256 _rewardAmount, address _rewardToken, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestTransactorSession) CheckpointHash(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _rewardAmount *big.Int, _rewardToken common.Address, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.CheckpointHash(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _rewardAmount, _rewardToken, _peggyId)
}

// JustSaveEverything is a paid mutator transaction binding the contract method 0x884403e2.
//
// Solidity: function JustSaveEverything(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestTransactor) JustSaveEverything(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.contract.Transact(opts, "JustSaveEverything", _validators, _powers, _valsetNonce)
}

// JustSaveEverything is a paid mutator transaction binding the contract method 0x884403e2.
//
// Solidity: function JustSaveEverything(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestSession) JustSaveEverything(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.Contract.JustSaveEverything(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce)
}

// JustSaveEverything is a paid mutator transaction binding the contract method 0x884403e2.
//
// Solidity: function JustSaveEverything(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestTransactorSession) JustSaveEverything(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.Contract.JustSaveEverything(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce)
}

// JustSaveEverythingAgain is a paid mutator transaction binding the contract method 0x715dff7e.
//
// Solidity: function JustSaveEverythingAgain(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestTransactor) JustSaveEverythingAgain(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.contract.Transact(opts, "JustSaveEverythingAgain", _validators, _powers, _valsetNonce)
}

// JustSaveEverythingAgain is a paid mutator transaction binding the contract method 0x715dff7e.
//
// Solidity: function JustSaveEverythingAgain(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestSession) JustSaveEverythingAgain(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.Contract.JustSaveEverythingAgain(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce)
}

// JustSaveEverythingAgain is a paid mutator transaction binding the contract method 0x715dff7e.
//
// Solidity: function JustSaveEverythingAgain(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestTransactorSession) JustSaveEverythingAgain(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.Contract.JustSaveEverythingAgain(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce)
}
