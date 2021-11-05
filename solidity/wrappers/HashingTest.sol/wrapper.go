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

// HashingTestABI is the input ABI used to generate the binding from.
const HashingTestABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_peggyId\",\"type\":\"bytes32\"}],\"name\":\"ConcatHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_peggyId\",\"type\":\"bytes32\"}],\"name\":\"ConcatHash2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_peggyId\",\"type\":\"bytes32\"}],\"name\":\"IterativeHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"}],\"name\":\"JustSaveEverything\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"}],\"name\":\"JustSaveEverythingAgain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_powers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HashingTestFuncSigs maps the 4-byte function signature to its string representation.
var HashingTestFuncSigs = map[string]string{
	"6071cbd9": "ConcatHash(address[],uint256[],uint256,bytes32)",
	"0caff28b": "ConcatHash2(address[],uint256[],uint256,bytes32)",
	"74df6ae4": "IterativeHash(address[],uint256[],uint256,bytes32)",
	"884403e2": "JustSaveEverything(address[],uint256[],uint256)",
	"715dff7e": "JustSaveEverythingAgain(address[],uint256[],uint256)",
	"d32e81a5": "lastCheckpoint()",
	"ccf0e74c": "state_nonce()",
	"2b939281": "state_powers(uint256)",
	"2afbb62e": "state_validators(uint256)",
}

// HashingTestBin is the compiled bytecode used for deploying new contracts.
var HashingTestBin = "0x608060405234801561001057600080fd5b50610915806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063715dff7e11610066578063715dff7e1461011157806374df6ae414610124578063884403e214610111578063ccf0e74c14610137578063d32e81a51461014057610093565b80630caff28b146100985780632afbb62e146100ad5780632b939281146100dd5780636071cbd9146100fe575b600080fd5b6100ab6100a63660046106c4565b610149565b005b6100c06100bb366004610735565b610197565b6040516001600160a01b0390911681526020015b60405180910390f35b6100f06100eb366004610735565b6101c1565b6040519081526020016100d4565b6100ab61010c3660046106c4565b6101e2565b6100ab61011f36600461065a565b6102b4565b6100ab6101323660046106c4565b6102e3565b6100f060035481565b6100f060005481565b6040516918da1958dadc1bda5b9d60b21b90600090610174908490849087908a908a906020016107ec565b60408051601f198184030181529190528051602090910120600055505050505050565b600181815481106101a757600080fd5b6000918252602090912001546001600160a01b0316905081565b600281815481106101d157600080fd5b600091825260209091200154905081565b60408051602081018390526918da1958dadc1bda5b9d60b21b9181018290526060810184905260009060800160405160208183030381529060405280519060200120905060008660405160200161023991906107bf565b60405160208183030381529060405280519060200120905060008660405160200161026491906107d9565b60408051808303601f190181528282528051602091820120818401969096528282019490945260608083019590955280518083039095018552608090910190525081519101206000555050505050565b82516102c79060019060208601906104cb565b5081516102db906002906020850190610530565b506003555050565b60408051602081018390526918da1958dadc1bda5b9d60b21b9181018290526060810184905260009060800160405160208183030381529060405280519060200120905060005b86518110156104c0578015610417578561034560018361089c565b8151811061036357634e487b7160e01b600052603260045260246000fd5b602002602001015186828151811061038b57634e487b7160e01b600052603260045260246000fd5b602002602001015111156104175760405162461bcd60e51b815260206004820152604360248201527f56616c696461746f7220706f776572206d757374206e6f74206265206869676860448201527f6572207468616e2070726576696f75732076616c696461746f7220696e2062616064820152620e8c6d60eb1b608482015260a40160405180910390fd5b8187828151811061043857634e487b7160e01b600052603260045260246000fd5b602002602001015187838151811061046057634e487b7160e01b600052603260045260246000fd5b6020026020010151604051602001610494939291909283526001600160a01b03919091166020830152604082015260600190565b6040516020818303038152906040528051906020012091508060016104b99190610884565b905061032a565b506000555050505050565b828054828255906000526020600020908101928215610520579160200282015b8281111561052057825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906104eb565b5061052c92915061056b565b5090565b828054828255906000526020600020908101928215610520579160200282015b82811115610520578251825591602001919060010190610550565b5b8082111561052c576000815560010161056c565b600082601f830112610590578081fd5b813560206105a56105a083610860565b61082f565b82815281810190858301838502870184018810156105c1578586fd5b855b858110156105f35781356001600160a01b03811681146105e1578788fd5b845292840192908401906001016105c3565b5090979650505050505050565b600082601f830112610610578081fd5b813560206106206105a083610860565b828152818101908583018385028701840188101561063c578586fd5b855b858110156105f35781358452928401929084019060010161063e565b60008060006060848603121561066e578283fd5b833567ffffffffffffffff80821115610685578485fd5b61069187838801610580565b945060208601359150808211156106a6578384fd5b506106b386828701610600565b925050604084013590509250925092565b600080600080608085870312156106d9578081fd5b843567ffffffffffffffff808211156106f0578283fd5b6106fc88838901610580565b95506020870135915080821115610711578283fd5b5061071e87828801610600565b949794965050505060408301359260600135919050565b600060208284031215610746578081fd5b5035919050565b6000815180845260208085019450808401835b838110156107855781516001600160a01b031687529582019590820190600101610760565b509495945050505050565b6000815180845260208085019450808401835b83811015610785578151875295820195908201906001016107a3565b6000602082526107d2602083018461074d565b9392505050565b6000602082526107d26020830184610790565b600086825285602083015284604083015260a0606083015261081160a083018561074d565b82810360808401526108238185610790565b98975050505050505050565b604051601f8201601f1916810167ffffffffffffffff81118282101715610858576108586108c9565b604052919050565b600067ffffffffffffffff82111561087a5761087a6108c9565b5060209081020190565b60008219821115610897576108976108b3565b500190565b6000828210156108ae576108ae6108b3565b500390565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212203165536843868778aa06608c6c05a755a3ba61e4a5b4475cd4a5245785e9828d64736f6c63430008020033"

// DeployHashingTest deploys a new Ethereum contract, binding an instance of HashingTest to it.
func DeployHashingTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashingTest, error) {
	parsed, err := abi.JSON(strings.NewReader(HashingTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HashingTestBin), backend)
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

// ConcatHash is a paid mutator transaction binding the contract method 0x6071cbd9.
//
// Solidity: function ConcatHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestTransactor) ConcatHash(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.contract.Transact(opts, "ConcatHash", _validators, _powers, _valsetNonce, _peggyId)
}

// ConcatHash is a paid mutator transaction binding the contract method 0x6071cbd9.
//
// Solidity: function ConcatHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestSession) ConcatHash(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.ConcatHash(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _peggyId)
}

// ConcatHash is a paid mutator transaction binding the contract method 0x6071cbd9.
//
// Solidity: function ConcatHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestTransactorSession) ConcatHash(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.ConcatHash(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _peggyId)
}

// ConcatHash2 is a paid mutator transaction binding the contract method 0x0caff28b.
//
// Solidity: function ConcatHash2(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestTransactor) ConcatHash2(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.contract.Transact(opts, "ConcatHash2", _validators, _powers, _valsetNonce, _peggyId)
}

// ConcatHash2 is a paid mutator transaction binding the contract method 0x0caff28b.
//
// Solidity: function ConcatHash2(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestSession) ConcatHash2(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.ConcatHash2(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _peggyId)
}

// ConcatHash2 is a paid mutator transaction binding the contract method 0x0caff28b.
//
// Solidity: function ConcatHash2(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestTransactorSession) ConcatHash2(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.ConcatHash2(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _peggyId)
}

// IterativeHash is a paid mutator transaction binding the contract method 0x74df6ae4.
//
// Solidity: function IterativeHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestTransactor) IterativeHash(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.contract.Transact(opts, "IterativeHash", _validators, _powers, _valsetNonce, _peggyId)
}

// IterativeHash is a paid mutator transaction binding the contract method 0x74df6ae4.
//
// Solidity: function IterativeHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestSession) IterativeHash(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.IterativeHash(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _peggyId)
}

// IterativeHash is a paid mutator transaction binding the contract method 0x74df6ae4.
//
// Solidity: function IterativeHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _peggyId) returns()
func (_HashingTest *HashingTestTransactorSession) IterativeHash(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _peggyId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.IterativeHash(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _peggyId)
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
