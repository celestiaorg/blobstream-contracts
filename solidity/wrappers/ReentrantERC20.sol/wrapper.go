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

// LogicCallArgs is an auto generated low-level Go binding around an user-defined struct.
type LogicCallArgs struct {
	TransferAmounts        []*big.Int
	TransferTokenContracts []common.Address
	FeeAmounts             []*big.Int
	FeeTokenContracts      []common.Address
	LogicContractAddress   common.Address
	Payload                []byte
	TimeOut                *big.Int
	InvalidationId         [32]byte
	InvalidationNonce      *big.Int
}

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220dc53804e208c55422efac5a2a226775bc4a6f1c7e1d19ec1db187980fe33604564736f6c63430008020033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// CosmosERC20ABI is the input ABI used to generate the binding from.
const CosmosERC20ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_peggyAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CosmosERC20FuncSigs maps the 4-byte function signature to its string representation.
var CosmosERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// CosmosERC20Bin is the compiled bytecode used for deploying new contracts.
var CosmosERC20Bin = "0x60806040526000196006553480156200001757600080fd5b5060405162000bcd38038062000bcd8339810160408190526200003a91620002e1565b8251839083906200005390600390602085019062000188565b5080516200006990600490602084019062000188565b50506005805460ff199081166012171660ff84161790555062000095846006546200009f60201b60201c565b50505050620003f8565b6001600160a01b038216620000fa5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b806002546200010a919062000380565b6002556001600160a01b0382166000908152602081905260409020546200013390829062000380565b6001600160a01b038316600081815260208181526040808320949094559251848152919290917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b8280546200019690620003a5565b90600052602060002090601f016020900481019282620001ba576000855562000205565b82601f10620001d557805160ff191683800117855562000205565b8280016001018555821562000205579182015b8281111562000205578251825591602001919060010190620001e8565b506200021392915062000217565b5090565b5b8082111562000213576000815560010162000218565b600082601f8301126200023f578081fd5b81516001600160401b03808211156200025c576200025c620003e2565b604051601f8301601f19908116603f01168101908282118183101715620002875762000287620003e2565b81604052838152602092508683858801011115620002a3578485fd5b8491505b83821015620002c65785820183015181830184015290820190620002a7565b83821115620002d757848385830101525b9695505050505050565b60008060008060808587031215620002f7578384fd5b84516001600160a01b03811681146200030e578485fd5b60208601519094506001600160401b03808211156200032b578485fd5b62000339888389016200022e565b945060408701519150808211156200034f578384fd5b506200035e878288016200022e565b925050606085015160ff8116811462000375578182fd5b939692955090935050565b60008219821115620003a057634e487b7160e01b81526011600452602481fd5b500190565b600281046001821680620003ba57607f821691505b60208210811415620003dc57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6107c580620004086000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012957806370a082311461013c57806395d89b411461014f578063a457c2d714610157578063a9059cbb1461016a578063dd62ed3e1461017d576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101b6565b6040516100c391906106bc565b60405180910390f35b6100df6100da366004610693565b610248565b60405190151581526020016100c3565b6002545b6040519081526020016100c3565b6100df61010f366004610658565b61025e565b60055460405160ff90911681526020016100c3565b6100df610137366004610693565b6102b0565b6100f361014a366004610605565b6102e7565b6100b6610306565b6100df610165366004610693565b610315565b6100df610178366004610693565b61034c565b6100f361018b366004610626565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101c59061073e565b80601f01602080910402602001604051908101604052809291908181526020018280546101f19061073e565b801561023e5780601f106102135761010080835404028352916020019161023e565b820191906000526020600020905b81548152906001019060200180831161022157829003601f168201915b5050505050905090565b6000610255338484610359565b50600192915050565b600061026b848484610483565b6001600160a01b0384166000908152600160209081526040808320338085529252909120546102a69186916102a1908690610727565b610359565b5060019392505050565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916102559185906102a190869061070f565b6001600160a01b0381166000908152602081905260409020545b919050565b6060600480546101c59061073e565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916102559185906102a1908690610727565b6000610255338484610483565b6001600160a01b0383166103c05760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084015b60405180910390fd5b6001600160a01b0382166104215760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103b7565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b0383166104e75760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103b7565b6001600160a01b0382166105495760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103b7565b6001600160a01b03831660009081526020819052604090205461056d908290610727565b6001600160a01b03808516600090815260208190526040808220939093559084168152205461059d90829061070f565b6001600160a01b038381166000818152602081815260409182902094909455518481529092918616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610476565b80356001600160a01b038116811461030157600080fd5b600060208284031215610616578081fd5b61061f826105ee565b9392505050565b60008060408385031215610638578081fd5b610641836105ee565b915061064f602084016105ee565b90509250929050565b60008060006060848603121561066c578081fd5b610675846105ee565b9250610683602085016105ee565b9150604084013590509250925092565b600080604083850312156106a5578182fd5b6106ae836105ee565b946020939093013593505050565b6000602080835283518082850152825b818110156106e8578581018301518582016040015282016106cc565b818111156106f95783604083870101525b50601f01601f1916929092016040019392505050565b6000821982111561072257610722610779565b500190565b60008282101561073957610739610779565b500390565b60028104600182168061075257607f821691505b6020821081141561077357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea26469706673582212206aed947a8645fbfdadadf80661cdd65f3a348fcc935a9c5314fb0223a2acd30264736f6c63430008020033"

// DeployCosmosERC20 deploys a new Ethereum contract, binding an instance of CosmosERC20 to it.
func DeployCosmosERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _peggyAddress common.Address, _name string, _symbol string, _decimals uint8) (common.Address, *types.Transaction, *CosmosERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(CosmosERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CosmosERC20Bin), backend, _peggyAddress, _name, _symbol, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosmosERC20{CosmosERC20Caller: CosmosERC20Caller{contract: contract}, CosmosERC20Transactor: CosmosERC20Transactor{contract: contract}, CosmosERC20Filterer: CosmosERC20Filterer{contract: contract}}, nil
}

// CosmosERC20 is an auto generated Go binding around an Ethereum contract.
type CosmosERC20 struct {
	CosmosERC20Caller     // Read-only binding to the contract
	CosmosERC20Transactor // Write-only binding to the contract
	CosmosERC20Filterer   // Log filterer for contract events
}

// CosmosERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type CosmosERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmosERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmosERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmosERC20Session struct {
	Contract     *CosmosERC20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CosmosERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmosERC20CallerSession struct {
	Contract *CosmosERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CosmosERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmosERC20TransactorSession struct {
	Contract     *CosmosERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CosmosERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type CosmosERC20Raw struct {
	Contract *CosmosERC20 // Generic contract binding to access the raw methods on
}

// CosmosERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmosERC20CallerRaw struct {
	Contract *CosmosERC20Caller // Generic read-only contract binding to access the raw methods on
}

// CosmosERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmosERC20TransactorRaw struct {
	Contract *CosmosERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmosERC20 creates a new instance of CosmosERC20, bound to a specific deployed contract.
func NewCosmosERC20(address common.Address, backend bind.ContractBackend) (*CosmosERC20, error) {
	contract, err := bindCosmosERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20{CosmosERC20Caller: CosmosERC20Caller{contract: contract}, CosmosERC20Transactor: CosmosERC20Transactor{contract: contract}, CosmosERC20Filterer: CosmosERC20Filterer{contract: contract}}, nil
}

// NewCosmosERC20Caller creates a new read-only instance of CosmosERC20, bound to a specific deployed contract.
func NewCosmosERC20Caller(address common.Address, caller bind.ContractCaller) (*CosmosERC20Caller, error) {
	contract, err := bindCosmosERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20Caller{contract: contract}, nil
}

// NewCosmosERC20Transactor creates a new write-only instance of CosmosERC20, bound to a specific deployed contract.
func NewCosmosERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*CosmosERC20Transactor, error) {
	contract, err := bindCosmosERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20Transactor{contract: contract}, nil
}

// NewCosmosERC20Filterer creates a new log filterer instance of CosmosERC20, bound to a specific deployed contract.
func NewCosmosERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*CosmosERC20Filterer, error) {
	contract, err := bindCosmosERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20Filterer{contract: contract}, nil
}

// bindCosmosERC20 binds a generic wrapper to an already deployed contract.
func bindCosmosERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CosmosERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmosERC20 *CosmosERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmosERC20.Contract.CosmosERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmosERC20 *CosmosERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmosERC20.Contract.CosmosERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmosERC20 *CosmosERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmosERC20.Contract.CosmosERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmosERC20 *CosmosERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmosERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmosERC20 *CosmosERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmosERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmosERC20 *CosmosERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmosERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CosmosERC20 *CosmosERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CosmosERC20 *CosmosERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CosmosERC20.Contract.Allowance(&_CosmosERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CosmosERC20 *CosmosERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CosmosERC20.Contract.Allowance(&_CosmosERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CosmosERC20 *CosmosERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CosmosERC20 *CosmosERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _CosmosERC20.Contract.BalanceOf(&_CosmosERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CosmosERC20 *CosmosERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CosmosERC20.Contract.BalanceOf(&_CosmosERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CosmosERC20 *CosmosERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CosmosERC20 *CosmosERC20Session) Decimals() (uint8, error) {
	return _CosmosERC20.Contract.Decimals(&_CosmosERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CosmosERC20 *CosmosERC20CallerSession) Decimals() (uint8, error) {
	return _CosmosERC20.Contract.Decimals(&_CosmosERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosmosERC20 *CosmosERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosmosERC20 *CosmosERC20Session) Name() (string, error) {
	return _CosmosERC20.Contract.Name(&_CosmosERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosmosERC20 *CosmosERC20CallerSession) Name() (string, error) {
	return _CosmosERC20.Contract.Name(&_CosmosERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosmosERC20 *CosmosERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosmosERC20 *CosmosERC20Session) Symbol() (string, error) {
	return _CosmosERC20.Contract.Symbol(&_CosmosERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosmosERC20 *CosmosERC20CallerSession) Symbol() (string, error) {
	return _CosmosERC20.Contract.Symbol(&_CosmosERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosmosERC20 *CosmosERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosmosERC20 *CosmosERC20Session) TotalSupply() (*big.Int, error) {
	return _CosmosERC20.Contract.TotalSupply(&_CosmosERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosmosERC20 *CosmosERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _CosmosERC20.Contract.TotalSupply(&_CosmosERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Approve(&_CosmosERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Approve(&_CosmosERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.DecreaseAllowance(&_CosmosERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.DecreaseAllowance(&_CosmosERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.IncreaseAllowance(&_CosmosERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.IncreaseAllowance(&_CosmosERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Transfer(&_CosmosERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Transfer(&_CosmosERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.TransferFrom(&_CosmosERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.TransferFrom(&_CosmosERC20.TransactOpts, sender, recipient, amount)
}

// CosmosERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CosmosERC20 contract.
type CosmosERC20ApprovalIterator struct {
	Event *CosmosERC20Approval // Event containing the contract specifics and raw log

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
func (it *CosmosERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmosERC20Approval)
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
		it.Event = new(CosmosERC20Approval)
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
func (it *CosmosERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmosERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmosERC20Approval represents a Approval event raised by the CosmosERC20 contract.
type CosmosERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CosmosERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CosmosERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20ApprovalIterator{contract: _CosmosERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CosmosERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CosmosERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmosERC20Approval)
				if err := _CosmosERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) ParseApproval(log types.Log) (*CosmosERC20Approval, error) {
	event := new(CosmosERC20Approval)
	if err := _CosmosERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmosERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CosmosERC20 contract.
type CosmosERC20TransferIterator struct {
	Event *CosmosERC20Transfer // Event containing the contract specifics and raw log

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
func (it *CosmosERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmosERC20Transfer)
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
		it.Event = new(CosmosERC20Transfer)
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
func (it *CosmosERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmosERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmosERC20Transfer represents a Transfer event raised by the CosmosERC20 contract.
type CosmosERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CosmosERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CosmosERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20TransferIterator{contract: _CosmosERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CosmosERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CosmosERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmosERC20Transfer)
				if err := _CosmosERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) ParseTransfer(log types.Log) (*CosmosERC20Transfer, error) {
	event := new(CosmosERC20Transfer)
	if err := _CosmosERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Bin is the compiled bytecode used for deploying new contracts.
var ERC20Bin = "0x60806040523480156200001157600080fd5b5060405162000a5d38038062000a5d8339810160408190526200003491620001ce565b81516200004990600390602085019062000075565b5080516200005f90600490602084019062000075565b50506005805460ff191660121790555062000288565b828054620000839062000235565b90600052602060002090601f016020900481019282620000a75760008555620000f2565b82601f10620000c257805160ff1916838001178555620000f2565b82800160010185558215620000f2579182015b82811115620000f2578251825591602001919060010190620000d5565b506200010092915062000104565b5090565b5b8082111562000100576000815560010162000105565b600082601f8301126200012c578081fd5b81516001600160401b038082111562000149576200014962000272565b604051601f8301601f19908116603f0116810190828211818310171562000174576200017462000272565b8160405283815260209250868385880101111562000190578485fd5b8491505b83821015620001b3578582018301518183018401529082019062000194565b83821115620001c457848385830101525b9695505050505050565b60008060408385031215620001e1578182fd5b82516001600160401b0380821115620001f8578384fd5b62000206868387016200011b565b935060208501519150808211156200021c578283fd5b506200022b858286016200011b565b9150509250929050565b6002810460018216806200024a57607f821691505b602082108114156200026c57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6107c580620002986000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012957806370a082311461013c57806395d89b411461014f578063a457c2d714610157578063a9059cbb1461016a578063dd62ed3e1461017d576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101b6565b6040516100c391906106bc565b60405180910390f35b6100df6100da366004610693565b610248565b60405190151581526020016100c3565b6002545b6040519081526020016100c3565b6100df61010f366004610658565b61025e565b60055460405160ff90911681526020016100c3565b6100df610137366004610693565b6102b0565b6100f361014a366004610605565b6102e7565b6100b6610306565b6100df610165366004610693565b610315565b6100df610178366004610693565b61034c565b6100f361018b366004610626565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101c59061073e565b80601f01602080910402602001604051908101604052809291908181526020018280546101f19061073e565b801561023e5780601f106102135761010080835404028352916020019161023e565b820191906000526020600020905b81548152906001019060200180831161022157829003601f168201915b5050505050905090565b6000610255338484610359565b50600192915050565b600061026b848484610483565b6001600160a01b0384166000908152600160209081526040808320338085529252909120546102a69186916102a1908690610727565b610359565b5060019392505050565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916102559185906102a190869061070f565b6001600160a01b0381166000908152602081905260409020545b919050565b6060600480546101c59061073e565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916102559185906102a1908690610727565b6000610255338484610483565b6001600160a01b0383166103c05760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084015b60405180910390fd5b6001600160a01b0382166104215760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103b7565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b0383166104e75760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103b7565b6001600160a01b0382166105495760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103b7565b6001600160a01b03831660009081526020819052604090205461056d908290610727565b6001600160a01b03808516600090815260208190526040808220939093559084168152205461059d90829061070f565b6001600160a01b038381166000818152602081815260409182902094909455518481529092918616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610476565b80356001600160a01b038116811461030157600080fd5b600060208284031215610616578081fd5b61061f826105ee565b9392505050565b60008060408385031215610638578081fd5b610641836105ee565b915061064f602084016105ee565b90509250929050565b60008060006060848603121561066c578081fd5b610675846105ee565b9250610683602085016105ee565b9150604084013590509250925092565b600080604083850312156106a5578182fd5b6106ae836105ee565b946020939093013593505050565b6000602080835283518082850152825b818110156106e8578581018301518582016040015282016106cc565b818111156106f95783604083870101525b50601f01601f1916929092016040019392505050565b6000821982111561072257610722610779565b500190565b60008282101561073957610739610779565b500390565b60028104600182168061075257607f821691505b6020821081141561077357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220af0f475b501aec1496483f4fb62b29139400d4e2ab27f931f799b877ac5ade0264736f6c63430008020033"

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, tokenName string, tokenSymbol string) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend, tokenName, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggyABI is the input ABI used to generate the binding from.
const PeggyABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_peggyId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_cosmosDenom\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"ERC20DeployedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_invalidationId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_invalidationNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"LogicCallEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"SendToCosmosEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"TransactionBatchExecutedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_newValsetNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"ValsetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cosmosDenom\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"deployERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Address\",\"type\":\"address\"}],\"name\":\"lastBatchNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_invalidation_id\",\"type\":\"bytes32\"}],\"name\":\"lastLogicCallNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendToCosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"state_invalidationMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"state_lastBatchNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastEventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_peggyId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stuff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_currentValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_currentPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_currentValsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_destinations\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_batchTimeout\",\"type\":\"uint256\"}],\"name\":\"submitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_currentValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_currentPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_currentValsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"transferAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"transferTokenContracts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"feeTokenContracts\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"logicContractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timeOut\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"invalidationId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"invalidationNonce\",\"type\":\"uint256\"}],\"internalType\":\"structLogicCallArgs\",\"name\":\"_args\",\"type\":\"tuple\"}],\"name\":\"submitLogicCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_currentValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_currentPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_theHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"}],\"name\":\"testCheckValidatorSignatures\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_peggyId\",\"type\":\"bytes32\"}],\"name\":\"testMakeCheckpoint\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_newPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_newValsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_currentValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_currentPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_currentValsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"updateValset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PeggyFuncSigs maps the 4-byte function signature to its string representation.
var PeggyFuncSigs = map[string]string{
	"f7955637": "deployERC20(string,string,string,uint8)",
	"011b2174": "lastBatchNonce(address)",
	"c9d194d5": "lastLogicCallNonce(bytes32)",
	"1ffbe7f9": "sendToCosmos(address,bytes32,uint256)",
	"7dfb6f86": "state_invalidationMapping(bytes32)",
	"df97174b": "state_lastBatchNonces(address)",
	"73b20547": "state_lastEventNonce()",
	"f2b53307": "state_lastValsetCheckpoint()",
	"b56561fe": "state_lastValsetNonce()",
	"69dd3908": "state_peggyId()",
	"e5a2b5d2": "state_powerThreshold()",
	"9c755f2f": "stuff()",
	"83b435db": "submitBatch(address[],uint256[],uint256,uint8[],bytes32[],bytes32[],uint256[],address[],uint256[],uint256,address,uint256)",
	"0c246c82": "submitLogicCall(address[],uint256[],uint256,uint8[],bytes32[],bytes32[],(uint256[],address[],uint256[],address[],address,bytes,uint256,bytes32,uint256))",
	"db7c4e57": "testCheckValidatorSignatures(address[],uint256[],uint8[],bytes32[],bytes32[],bytes32,uint256)",
	"c227c30b": "testMakeCheckpoint(address[],uint256[],uint256,bytes32)",
	"e3cb9f62": "updateValset(address[],uint256[],uint256,address[],uint256[],uint256,uint8[],bytes32[],bytes32[])",
}

// PeggyBin is the compiled bytecode used for deploying new contracts.
var PeggyBin = "0x6080604052600060045560006005553480156200001b57600080fd5b506040516200347c3803806200347c8339810160408190526200003e91620002af565b600160005580518251146200009a5760405162461bcd60e51b815260206004820152601f60248201527f4d616c666f726d65642063757272656e742076616c696461746f72207365740060448201526064015b60405180910390fd5b6000805b82518110156200010457828181518110620000c957634e487b7160e01b600052603260045260246000fd5b602002602001015182620000de9190620004e6565b915084821115620000ef5762000104565b80620000fb8162000501565b9150506200009e565b508381116200017c5760405162461bcd60e51b815260206004820152603c60248201527f5375626d69747465642076616c696461746f7220736574207369676e6174757260448201527f657320646f206e6f74206861766520656e6f75676820706f7765722e00000000606482015260840162000091565b60006200018c84848389620001e9565b6006879055600786905560018190556040519091506000907fc6d025c076bafcdd040f00632d5e280b3a5188963f110f8c70c4f810184b30f390620001d5908790879062000414565b60405180910390a25050505050506200054b565b6040516000906918da1958dadc1bda5b9d60b21b90829062000218908590849088908b908b9060200162000446565b60408051808303601f190181529190528051602090910120979650505050505050565b600082601f8301126200024c578081fd5b81516020620002656200025f83620004c0565b6200048d565b828152818101908583018385028701840188101562000282578586fd5b855b85811015620002a25781518452928401929084019060010162000284565b5090979650505050505050565b60008060008060808587031215620002c5578384fd5b845160208087015160408801519296509450906001600160401b0380821115620002ed578485fd5b818801915088601f83011262000301578485fd5b8151620003126200025f82620004c0565b81815284810190848601868402860187018d10156200032f578889fd5b8895505b83861015620003685780516001600160a01b03811681146200035357898afd5b83526001959095019491860191860162000333565b5060608b0151909750945050508083111562000382578384fd5b505062000392878288016200023b565b91505092959194509250565b6000815180845260208085019450808401835b83811015620003d85781516001600160a01b031687529582019590820190600101620003b1565b509495945050505050565b6000815180845260208085019450808401835b83811015620003d857815187529582019590820190600101620003f6565b6000604082526200042960408301856200039e565b82810360208401526200043d8185620003e3565b95945050505050565b600086825285602083015284604083015260a060608301526200046d60a08301856200039e565b8281036080840152620004818185620003e3565b98975050505050505050565b604051601f8201601f191681016001600160401b0381118282101715620004b857620004b862000535565b604052919050565b60006001600160401b03821115620004dc57620004dc62000535565b5060209081020190565b60008219821115620004fc57620004fc6200051f565b500190565b60006000198214156200051857620005186200051f565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b612f21806200055b6000396000f3fe60806040523480156200001157600080fd5b5060043610620001245760003560e01c8063b56561fe11620000b1578063df97174b116200007b578063df97174b1462000235578063e3cb9f621462000258578063e5a2b5d2146200026f578063f2b533071462000279578063f795563714620002835762000124565b8063b56561fe14620001da578063c227c30b14620001e4578063c9d194d514620001fb578063db7c4e57146200021e5762000124565b806373b2054711620000f357806373b20547146200018c5780637dfb6f86146200019657806383b435db14620001b95780639c755f2f14620001d05762000124565b8063011b217414620001295780630c246c8214620001525780631ffbe7f9146200016b57806369dd39081462000182575b600080fd5b620001406200013a36600462001786565b6200029a565b60405190815260200160405180910390f35b620001696200016336600462001bb9565b620002b9565b005b620001696200017c366004620017a3565b62000739565b6200014060065481565b6200014060055481565b62000140620001a736600462001d5d565b60036020526000908152604090205481565b62000169620001ca36600462001a02565b620007e9565b6200014060085481565b6200014060045481565b62000169620001f536600462001cc4565b62000bb7565b620001406200020c36600462001d5d565b60009081526003602052604090205490565b620001696200022f366004620017d8565b62000bcc565b620001406200024636600462001786565b60026020526000908152604090205481565b6200016962000269366004620018c6565b62000be6565b6200014060075481565b6200014060015481565b620001696200029436600462001d76565b62000dac565b6001600160a01b0381166000908152600260205260409020545b919050565b60026000541415620002e85760405162461bcd60e51b8152600401620002df90620021fa565b60405180910390fd5b600260005560c081015143106200032e5760405162461bcd60e51b8152602060048201526009602482015268151a5b5959081bdd5d60ba1b6044820152606401620002df565b61010081015160e082015160009081526003602052604090205410620003bd5760405162461bcd60e51b815260206004820152603d60248201527f4e657720696e76616c69646174696f6e206e6f6e6365206d757374206265206760448201527f726561746572207468616e207468652063757272656e74206e6f6e63650000006064820152608401620002df565b85518751148015620003d0575083518751145b8015620003de575082518751145b8015620003ec575081518751145b6200040b5760405162461bcd60e51b8152600401620002df90620021c3565b6001546200041e88888860065462000e58565b146200043e5760405162461bcd60e51b8152600401620002df9062002166565b60208101515181515114620004a05760405162461bcd60e51b815260206004820152602160248201527f4d616c666f726d6564206c697374206f6620746f6b656e207472616e736665726044820152607360f81b6064820152608401620002df565b80606001515181604001515114620004f45760405162461bcd60e51b81526020600482015260166024820152754d616c666f726d6564206c697374206f66206665657360501b6044820152606401620002df565b6000600654681b1bd9da58d0d85b1b60ba1b836000015184602001518560400151866060015187608001518860a001518960c001518a60e001518b6101000151604051602001620005509b9a9998979695949392919062001fa0565b6040516020818303038152906040528051906020012090506200057b88888787878660075462000eac565b61010082015160e08301516000908152600360205260408120919091555b825151811015620006355762000620836080015184600001518381518110620005d257634e487b7160e01b600052603260045260246000fd5b602002602001015185602001518481518110620005ff57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b0316620010d59092919063ffffffff16565b806200062c81620022d4565b91505062000599565b5060006200064c83608001518460a001516200113f565b905060005b836040015151811015620006cb57620006b633856040015183815181106200068957634e487b7160e01b600052603260045260246000fd5b602002602001015186606001518481518110620005ff57634e487b7160e01b600052603260045260246000fd5b80620006c281620022d4565b91505062000651565b50600554620006dc9060016200228a565b600581905560e08401516101008501516040517f7c2bb24f8e1b3725cb613d7f11ef97d9745cc97a0e40f730621c052d684077a19362000721939291869190620020c8565b60405180910390a15050600160005550505050505050565b600260005414156200075f5760405162461bcd60e51b8152600401620002df90620021fa565b60026000556200077b6001600160a01b0384163330846200118a565b6005546200078b9060016200228a565b6005819055604051839133916001600160a01b038716917fd7767894d73c589daeca9643f445f03d7be61aad2950c117e7cbff4176fca7e491620007d791878252602082015260400190565b60405180910390a45050600160005550565b600260005414156200080f5760405162461bcd60e51b8152600401620002df90620021fa565b600260008181556001600160a01b0384168152602091909152604090205483116200089c5760405162461bcd60e51b815260206004820152603660248201527f4e6577206261746368206e6f6e6365206d7573742062652067726561746572206044820152757468616e207468652063757272656e74206e6f6e636560501b6064820152608401620002df565b804310620009135760405162461bcd60e51b815260206004820152603b60248201527f42617463682074696d656f7574206d757374206265206772656174657220746860448201527f616e207468652063757272656e7420626c6f636b2068656967687400000000006064820152608401620002df565b8a518c5114801562000926575088518c51145b801562000934575087518c51145b801562000942575086518c51145b620009615760405162461bcd60e51b8152600401620002df90620021c3565b600154620009748d8d8d60065462000e58565b14620009945760405162461bcd60e51b8152600401620002df9062002166565b84518651148015620009a7575083518651145b620009f55760405162461bcd60e51b815260206004820152601f60248201527f4d616c666f726d6564206261746368206f66207472616e73616374696f6e73006044820152606401620002df565b62000a538c8c8b8b8b6006546f0e8e4c2dce6c2c6e8d2dedc84c2e8c6d60831b8d8d8d8d8d8d60405160200162000a3498979695949392919062002050565b6040516020818303038152906040528051906020012060075462000eac565b6001600160a01b0382166000908152600260205260408120849055805b875181101562000b375762000aea87828151811062000a9f57634e487b7160e01b600052603260045260246000fd5b602002602001015189838151811062000ac857634e487b7160e01b600052603260045260246000fd5b6020026020010151866001600160a01b0316620010d59092919063ffffffff16565b85818151811062000b0b57634e487b7160e01b600052603260045260246000fd5b60200260200101518262000b2091906200228a565b91508062000b2e81620022d4565b91505062000a70565b5062000b4e6001600160a01b0384163383620010d5565b5060055462000b5f9060016200228a565b60058190556040519081526001600160a01b0383169084907f02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab7089060200160405180910390a35050600160005550505050505050505050565b62000bc58484848462000e58565b5050505050565b62000bdd8787878787878762000eac565b50505050505050565b83871162000c5d5760405162461bcd60e51b815260206004820152603760248201527f4e65772076616c736574206e6f6e6365206d757374206265206772656174657260448201527f207468616e207468652063757272656e74206e6f6e63650000000000000000006064820152608401620002df565b875189511462000cb05760405162461bcd60e51b815260206004820152601b60248201527f4d616c666f726d6564206e65772076616c696461746f722073657400000000006044820152606401620002df565b8451865114801562000cc3575082518651145b801562000cd1575081518651145b801562000cdf575080518651145b62000cfe5760405162461bcd60e51b8152600401620002df90620021c3565b60015462000d1187878760065462000e58565b1462000d315760405162461bcd60e51b8152600401620002df9062002166565b600062000d438a8a8a60065462000e58565b905062000d5887878686868660075462000eac565b6001819055600488905560405188907fc6d025c076bafcdd040f00632d5e280b3a5188963f110f8c70c4f810184b30f39062000d98908d908d9062001f27565b60405180910390a250505050505050505050565b60003084848460405162000dc09062001470565b62000dcf949392919062001ed8565b604051809103906000f08015801562000dec573d6000803e3d6000fd5b509050600554600162000e0091906200228a565b60058190556040516001600160a01b038316917f82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c79162000e49918991899189918991906200210f565b60405180910390a25050505050565b6040516000906918da1958dadc1bda5b9d60b21b90829062000e87908590849088908b908b9060200162001f59565b60408051601f198184030181529190528051602090910120925050505b949350505050565b6000805b8851811015620010535786818151811062000edb57634e487b7160e01b600052603260045260246000fd5b602002602001015160ff166000146200103e5762000f9d89828151811062000f1357634e487b7160e01b600052603260045260246000fd5b60200260200101518589848151811062000f3d57634e487b7160e01b600052603260045260246000fd5b602002602001015189858151811062000f6657634e487b7160e01b600052603260045260246000fd5b602002602001015189868151811062000f8f57634e487b7160e01b600052603260045260246000fd5b6020026020010151620011ca565b62000ff75760405162461bcd60e51b815260206004820152602360248201527f56616c696461746f72207369676e617475726520646f6573206e6f74206d617460448201526231b41760e91b6064820152608401620002df565b8781815181106200101857634e487b7160e01b600052603260045260246000fd5b6020026020010151826200102d91906200228a565b9150828211156200103e5762001053565b806200104a81620022d4565b91505062000eb0565b50818111620010cb5760405162461bcd60e51b815260206004820152603c60248201527f5375626d69747465642076616c696461746f7220736574207369676e6174757260448201527f657320646f206e6f74206861766520656e6f75676820706f7765722e000000006064820152608401620002df565b5050505050505050565b6040516001600160a01b0383166024820152604481018290526200113a90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915262001295565b505050565b60606200118383836040518060400160405280601e81526020017f416464726573733a206c6f772d6c6576656c2063616c6c206661696c656400008152506200136e565b9392505050565b6040516001600160a01b0380851660248301528316604482015260648101829052620011c49085906323b872dd60e01b9060840162001102565b50505050565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018590526000908190605c0160408051601f1981840301815282825280516020918201206000845290830180835281905260ff8816918301919091526060820186905260808201859052915060019060a0016020604051602081039080840390855afa1580156200126b573d6000803e3d6000fd5b505050602060405103516001600160a01b0316876001600160a01b03161491505095945050505050565b6000620012ec826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166200136e9092919063ffffffff16565b8051909150156200113a57808060200190518101906200130d919062001d3b565b6200113a5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401620002df565b606062000ea484846000856060843b620013cb5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401620002df565b600080866001600160a01b03168587604051620013e9919062001eba565b60006040518083038185875af1925050503d806000811462001428576040519150601f19603f3d011682016040523d82523d6000602084013e6200142d565b606091505b509150915081156200144357915062000ea49050565b805115620014545780518082602001fd5b8360405162461bcd60e51b8152600401620002df9190620020fa565b610bcd806200231f83390190565b80356001600160a01b0381168114620002b457600080fd5b600082601f830112620014a7578081fd5b81356020620014c0620014ba8362002264565b62002231565b8281528181019085830183850287018401881015620014dd578586fd5b855b858110156200150657620014f3826200147e565b84529284019290840190600101620014df565b5090979650505050505050565b600082601f83011262001524578081fd5b8135602062001537620014ba8362002264565b828152818101908583018385028701840188101562001554578586fd5b855b85811015620015065781358452928401929084019060010162001556565b600082601f83011262001585578081fd5b8135602062001598620014ba8362002264565b8281528181019085830183850287018401881015620015b5578586fd5b855b858110156200150657620015cb8262001774565b84529284019290840190600101620015b7565b600082601f830112620015ef578081fd5b81356001600160401b038111156200160b576200160b62002308565b62001620601f8201601f191660200162002231565b81815284602083860101111562001635578283fd5b816020850160208301379081016020019190915292915050565b600061012080838503121562001663578182fd5b6200166e8162002231565b91505081356001600160401b03808211156200168957600080fd5b620016978583860162001513565b83526020840135915080821115620016ae57600080fd5b620016bc8583860162001496565b60208401526040840135915080821115620016d657600080fd5b620016e48583860162001513565b60408401526060840135915080821115620016fe57600080fd5b6200170c8583860162001496565b60608401526200171f608085016200147e565b608084015260a08401359150808211156200173957600080fd5b506200174884828501620015de565b60a08301525060c082013560c082015260e082013560e082015261010080830135818301525092915050565b803560ff81168114620002b457600080fd5b60006020828403121562001798578081fd5b62001183826200147e565b600080600060608486031215620017b8578182fd5b620017c3846200147e565b95602085013595506040909401359392505050565b600080600080600080600060e0888a031215620017f3578283fd5b87356001600160401b03808211156200180a578485fd5b620018188b838c0162001496565b985060208a01359150808211156200182e578485fd5b6200183c8b838c0162001513565b975060408a013591508082111562001852578485fd5b620018608b838c0162001574565b965060608a013591508082111562001876578485fd5b620018848b838c0162001513565b955060808a01359150808211156200189a578485fd5b50620018a98a828b0162001513565b93505060a0880135915060c0880135905092959891949750929550565b60008060008060008060008060006101208a8c031215620018e5578283fd5b89356001600160401b0380821115620018fc578485fd5b6200190a8d838e0162001496565b9a5060208c013591508082111562001920578485fd5b6200192e8d838e0162001513565b995060408c0135985060608c01359150808211156200194b578485fd5b620019598d838e0162001496565b975060808c01359150808211156200196f578485fd5b6200197d8d838e0162001513565b965060a08c0135955060c08c01359150808211156200199a578485fd5b620019a88d838e0162001574565b945060e08c0135915080821115620019be578384fd5b620019cc8d838e0162001513565b93506101008c0135915080821115620019e3578283fd5b50620019f28c828d0162001513565b9150509295985092959850929598565b6000806000806000806000806000806000806101808d8f03121562001a25578586fd5b6001600160401b038d35111562001a3a578586fd5b62001a498e8e358f0162001496565b9b506001600160401b0360208e0135111562001a63578586fd5b62001a758e60208f01358f0162001513565b9a5060408d013599506001600160401b0360608e0135111562001a96578586fd5b62001aa88e60608f01358f0162001574565b98506001600160401b0360808e0135111562001ac2578586fd5b62001ad48e60808f01358f0162001513565b97506001600160401b0360a08e0135111562001aee578586fd5b62001b008e60a08f01358f0162001513565b96506001600160401b0360c08e0135111562001b1a578586fd5b62001b2c8e60c08f01358f0162001513565b95506001600160401b0360e08e0135111562001b46578283fd5b62001b588e60e08f01358f0162001496565b94506001600160401b036101008e0135111562001b73578283fd5b62001b868e6101008f01358f0162001513565b93506101208d0135925062001b9f6101408e016200147e565b91506101608d013590509295989b509295989b509295989b565b600080600080600080600060e0888a03121562001bd4578081fd5b87356001600160401b038082111562001beb578283fd5b62001bf98b838c0162001496565b985060208a013591508082111562001c0f578283fd5b62001c1d8b838c0162001513565b975060408a0135965060608a013591508082111562001c3a578283fd5b62001c488b838c0162001574565b955060808a013591508082111562001c5e578283fd5b62001c6c8b838c0162001513565b945060a08a013591508082111562001c82578283fd5b62001c908b838c0162001513565b935060c08a013591508082111562001ca6578283fd5b5062001cb58a828b016200164f565b91505092959891949750929550565b6000806000806080858703121562001cda578182fd5b84356001600160401b038082111562001cf1578384fd5b62001cff8883890162001496565b9550602087013591508082111562001d15578384fd5b5062001d248782880162001513565b949794965050505060408301359260600135919050565b60006020828403121562001d4d578081fd5b8151801515811462001183578182fd5b60006020828403121562001d6f578081fd5b5035919050565b6000806000806080858703121562001d8c578182fd5b84356001600160401b038082111562001da3578384fd5b62001db188838901620015de565b9550602087013591508082111562001dc7578384fd5b62001dd588838901620015de565b9450604087013591508082111562001deb578384fd5b5062001dfa87828801620015de565b92505062001e0b6060860162001774565b905092959194509250565b6000815180845260208085019450808401835b8381101562001e505781516001600160a01b03168752958201959082019060010162001e29565b509495945050505050565b6000815180845260208085019450808401835b8381101562001e505781518752958201959082019060010162001e6e565b6000815180845262001ea6816020860160208601620022a5565b601f01601f19169290920160200192915050565b6000825162001ece818460208701620022a5565b9190910192915050565b6001600160a01b038516815260806020820181905260009062001efe9083018662001e8c565b828103604084015262001f12818662001e8c565b91505060ff8316606083015295945050505050565b60006040825262001f3c604083018562001e16565b828103602084015262001f50818562001e5b565b95945050505050565b600086825285602083015284604083015260a0606083015262001f8060a083018562001e16565b828103608084015262001f94818562001e5b565b98975050505050505050565b60006101608d83528c602084015280604084015262001fc28184018d62001e5b565b9050828103606084015262001fd8818c62001e16565b9050828103608084015262001fee818b62001e5b565b905082810360a084015262002004818a62001e16565b6001600160a01b03891660c085015283810360e0850152905062002029818862001e8c565b61010084019690965250506101208101929092526101409091015298975050505050505050565b60006101008a8352896020840152806040840152620020728184018a62001e5b565b9050828103606084015262002088818962001e16565b905082810360808401526200209e818862001e5b565b60a084019690965250506001600160a01b039290921660c083015260e09091015295945050505050565b600085825284602083015260806040830152620020e9608083018562001e8c565b905082606083015295945050505050565b60006020825262001183602083018462001e8c565b600060a082526200212460a083018862001e8c565b828103602084015262002138818862001e8c565b905082810360408401526200214e818762001e8c565b60ff9590951660608401525050608001529392505050565b6020808252603f908201527f537570706c6965642063757272656e742076616c696461746f727320616e642060408201527f706f7765727320646f206e6f74206d6174636820636865636b706f696e742e00606082015260800190565b6020808252601f908201527f4d616c666f726d65642063757272656e742076616c696461746f722073657400604082015260600190565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b604051601f8201601f191681016001600160401b03811182821017156200225c576200225c62002308565b604052919050565b60006001600160401b0382111562002280576200228062002308565b5060209081020190565b60008219821115620022a057620022a0620022f2565b500190565b60005b83811015620022c2578181015183820152602001620022a8565b83811115620011c45750506000910152565b6000600019821415620022eb57620022eb620022f2565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfe60806040526000196006553480156200001757600080fd5b5060405162000bcd38038062000bcd8339810160408190526200003a91620002e1565b8251839083906200005390600390602085019062000188565b5080516200006990600490602084019062000188565b50506005805460ff199081166012171660ff84161790555062000095846006546200009f60201b60201c565b50505050620003f8565b6001600160a01b038216620000fa5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b806002546200010a919062000380565b6002556001600160a01b0382166000908152602081905260409020546200013390829062000380565b6001600160a01b038316600081815260208181526040808320949094559251848152919290917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b8280546200019690620003a5565b90600052602060002090601f016020900481019282620001ba576000855562000205565b82601f10620001d557805160ff191683800117855562000205565b8280016001018555821562000205579182015b8281111562000205578251825591602001919060010190620001e8565b506200021392915062000217565b5090565b5b8082111562000213576000815560010162000218565b600082601f8301126200023f578081fd5b81516001600160401b03808211156200025c576200025c620003e2565b604051601f8301601f19908116603f01168101908282118183101715620002875762000287620003e2565b81604052838152602092508683858801011115620002a3578485fd5b8491505b83821015620002c65785820183015181830184015290820190620002a7565b83821115620002d757848385830101525b9695505050505050565b60008060008060808587031215620002f7578384fd5b84516001600160a01b03811681146200030e578485fd5b60208601519094506001600160401b03808211156200032b578485fd5b62000339888389016200022e565b945060408701519150808211156200034f578384fd5b506200035e878288016200022e565b925050606085015160ff8116811462000375578182fd5b939692955090935050565b60008219821115620003a057634e487b7160e01b81526011600452602481fd5b500190565b600281046001821680620003ba57607f821691505b60208210811415620003dc57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b6107c580620004086000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012957806370a082311461013c57806395d89b411461014f578063a457c2d714610157578063a9059cbb1461016a578063dd62ed3e1461017d576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101b6565b6040516100c391906106bc565b60405180910390f35b6100df6100da366004610693565b610248565b60405190151581526020016100c3565b6002545b6040519081526020016100c3565b6100df61010f366004610658565b61025e565b60055460405160ff90911681526020016100c3565b6100df610137366004610693565b6102b0565b6100f361014a366004610605565b6102e7565b6100b6610306565b6100df610165366004610693565b610315565b6100df610178366004610693565b61034c565b6100f361018b366004610626565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101c59061073e565b80601f01602080910402602001604051908101604052809291908181526020018280546101f19061073e565b801561023e5780601f106102135761010080835404028352916020019161023e565b820191906000526020600020905b81548152906001019060200180831161022157829003601f168201915b5050505050905090565b6000610255338484610359565b50600192915050565b600061026b848484610483565b6001600160a01b0384166000908152600160209081526040808320338085529252909120546102a69186916102a1908690610727565b610359565b5060019392505050565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916102559185906102a190869061070f565b6001600160a01b0381166000908152602081905260409020545b919050565b6060600480546101c59061073e565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916102559185906102a1908690610727565b6000610255338484610483565b6001600160a01b0383166103c05760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084015b60405180910390fd5b6001600160a01b0382166104215760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103b7565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b0383166104e75760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103b7565b6001600160a01b0382166105495760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103b7565b6001600160a01b03831660009081526020819052604090205461056d908290610727565b6001600160a01b03808516600090815260208190526040808220939093559084168152205461059d90829061070f565b6001600160a01b038381166000818152602081815260409182902094909455518481529092918616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610476565b80356001600160a01b038116811461030157600080fd5b600060208284031215610616578081fd5b61061f826105ee565b9392505050565b60008060408385031215610638578081fd5b610641836105ee565b915061064f602084016105ee565b90509250929050565b60008060006060848603121561066c578081fd5b610675846105ee565b9250610683602085016105ee565b9150604084013590509250925092565b600080604083850312156106a5578182fd5b6106ae836105ee565b946020939093013593505050565b6000602080835283518082850152825b818110156106e8578581018301518582016040015282016106cc565b818111156106f95783604083870101525b50601f01601f1916929092016040019392505050565b6000821982111561072257610722610779565b500190565b60008282101561073957610739610779565b500390565b60028104600182168061075257607f821691505b6020821081141561077357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea26469706673582212206aed947a8645fbfdadadf80661cdd65f3a348fcc935a9c5314fb0223a2acd30264736f6c63430008020033a2646970667358221220f94049f0dd7262ebc6d8c72bf30e43b3badfdc5c05c0a1b8ad1e089836515e3a64736f6c63430008020033"

// DeployPeggy deploys a new Ethereum contract, binding an instance of Peggy to it.
func DeployPeggy(auth *bind.TransactOpts, backend bind.ContractBackend, _peggyId [32]byte, _powerThreshold *big.Int, _validators []common.Address, _powers []*big.Int) (common.Address, *types.Transaction, *Peggy, error) {
	parsed, err := abi.JSON(strings.NewReader(PeggyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PeggyBin), backend, _peggyId, _powerThreshold, _validators, _powers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Peggy{PeggyCaller: PeggyCaller{contract: contract}, PeggyTransactor: PeggyTransactor{contract: contract}, PeggyFilterer: PeggyFilterer{contract: contract}}, nil
}

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

// LastLogicCallNonce is a free data retrieval call binding the contract method 0xc9d194d5.
//
// Solidity: function lastLogicCallNonce(bytes32 _invalidation_id) view returns(uint256)
func (_Peggy *PeggyCaller) LastLogicCallNonce(opts *bind.CallOpts, _invalidation_id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "lastLogicCallNonce", _invalidation_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLogicCallNonce is a free data retrieval call binding the contract method 0xc9d194d5.
//
// Solidity: function lastLogicCallNonce(bytes32 _invalidation_id) view returns(uint256)
func (_Peggy *PeggySession) LastLogicCallNonce(_invalidation_id [32]byte) (*big.Int, error) {
	return _Peggy.Contract.LastLogicCallNonce(&_Peggy.CallOpts, _invalidation_id)
}

// LastLogicCallNonce is a free data retrieval call binding the contract method 0xc9d194d5.
//
// Solidity: function lastLogicCallNonce(bytes32 _invalidation_id) view returns(uint256)
func (_Peggy *PeggyCallerSession) LastLogicCallNonce(_invalidation_id [32]byte) (*big.Int, error) {
	return _Peggy.Contract.LastLogicCallNonce(&_Peggy.CallOpts, _invalidation_id)
}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_Peggy *PeggyCaller) StateInvalidationMapping(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "state_invalidationMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_Peggy *PeggySession) StateInvalidationMapping(arg0 [32]byte) (*big.Int, error) {
	return _Peggy.Contract.StateInvalidationMapping(&_Peggy.CallOpts, arg0)
}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_Peggy *PeggyCallerSession) StateInvalidationMapping(arg0 [32]byte) (*big.Int, error) {
	return _Peggy.Contract.StateInvalidationMapping(&_Peggy.CallOpts, arg0)
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

// Stuff is a free data retrieval call binding the contract method 0x9c755f2f.
//
// Solidity: function stuff() view returns(uint256)
func (_Peggy *PeggyCaller) Stuff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Peggy.contract.Call(opts, &out, "stuff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stuff is a free data retrieval call binding the contract method 0x9c755f2f.
//
// Solidity: function stuff() view returns(uint256)
func (_Peggy *PeggySession) Stuff() (*big.Int, error) {
	return _Peggy.Contract.Stuff(&_Peggy.CallOpts)
}

// Stuff is a free data retrieval call binding the contract method 0x9c755f2f.
//
// Solidity: function stuff() view returns(uint256)
func (_Peggy *PeggyCallerSession) Stuff() (*big.Int, error) {
	return _Peggy.Contract.Stuff(&_Peggy.CallOpts)
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

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_Peggy *PeggyTransactor) DeployERC20(opts *bind.TransactOpts, _cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Peggy.contract.Transact(opts, "deployERC20", _cosmosDenom, _name, _symbol, _decimals)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_Peggy *PeggySession) DeployERC20(_cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Peggy.Contract.DeployERC20(&_Peggy.TransactOpts, _cosmosDenom, _name, _symbol, _decimals)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_Peggy *PeggyTransactorSession) DeployERC20(_cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Peggy.Contract.DeployERC20(&_Peggy.TransactOpts, _cosmosDenom, _name, _symbol, _decimals)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x1ffbe7f9.
//
// Solidity: function sendToCosmos(address _tokenContract, bytes32 _destination, uint256 _amount) returns()
func (_Peggy *PeggyTransactor) SendToCosmos(opts *bind.TransactOpts, _tokenContract common.Address, _destination [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Peggy.contract.Transact(opts, "sendToCosmos", _tokenContract, _destination, _amount)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x1ffbe7f9.
//
// Solidity: function sendToCosmos(address _tokenContract, bytes32 _destination, uint256 _amount) returns()
func (_Peggy *PeggySession) SendToCosmos(_tokenContract common.Address, _destination [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Peggy.Contract.SendToCosmos(&_Peggy.TransactOpts, _tokenContract, _destination, _amount)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x1ffbe7f9.
//
// Solidity: function sendToCosmos(address _tokenContract, bytes32 _destination, uint256 _amount) returns()
func (_Peggy *PeggyTransactorSession) SendToCosmos(_tokenContract common.Address, _destination [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Peggy.Contract.SendToCosmos(&_Peggy.TransactOpts, _tokenContract, _destination, _amount)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x83b435db.
//
// Solidity: function submitBatch(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_Peggy *PeggyTransactor) SubmitBatch(opts *bind.TransactOpts, _currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Peggy.contract.Transact(opts, "submitBatch", _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x83b435db.
//
// Solidity: function submitBatch(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_Peggy *PeggySession) SubmitBatch(_currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Peggy.Contract.SubmitBatch(&_Peggy.TransactOpts, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x83b435db.
//
// Solidity: function submitBatch(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_Peggy *PeggyTransactorSession) SubmitBatch(_currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Peggy.Contract.SubmitBatch(&_Peggy.TransactOpts, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// SubmitLogicCall is a paid mutator transaction binding the contract method 0x0c246c82.
//
// Solidity: function submitLogicCall(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, (uint256[],address[],uint256[],address[],address,bytes,uint256,bytes32,uint256) _args) returns()
func (_Peggy *PeggyTransactor) SubmitLogicCall(opts *bind.TransactOpts, _currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _args LogicCallArgs) (*types.Transaction, error) {
	return _Peggy.contract.Transact(opts, "submitLogicCall", _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _args)
}

// SubmitLogicCall is a paid mutator transaction binding the contract method 0x0c246c82.
//
// Solidity: function submitLogicCall(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, (uint256[],address[],uint256[],address[],address,bytes,uint256,bytes32,uint256) _args) returns()
func (_Peggy *PeggySession) SubmitLogicCall(_currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _args LogicCallArgs) (*types.Transaction, error) {
	return _Peggy.Contract.SubmitLogicCall(&_Peggy.TransactOpts, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _args)
}

// SubmitLogicCall is a paid mutator transaction binding the contract method 0x0c246c82.
//
// Solidity: function submitLogicCall(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, (uint256[],address[],uint256[],address[],address,bytes,uint256,bytes32,uint256) _args) returns()
func (_Peggy *PeggyTransactorSession) SubmitLogicCall(_currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _args LogicCallArgs) (*types.Transaction, error) {
	return _Peggy.Contract.SubmitLogicCall(&_Peggy.TransactOpts, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _args)
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

// PeggyERC20DeployedEventIterator is returned from FilterERC20DeployedEvent and is used to iterate over the raw logs and unpacked data for ERC20DeployedEvent events raised by the Peggy contract.
type PeggyERC20DeployedEventIterator struct {
	Event *PeggyERC20DeployedEvent // Event containing the contract specifics and raw log

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
func (it *PeggyERC20DeployedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggyERC20DeployedEvent)
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
		it.Event = new(PeggyERC20DeployedEvent)
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
func (it *PeggyERC20DeployedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggyERC20DeployedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggyERC20DeployedEvent represents a ERC20DeployedEvent event raised by the Peggy contract.
type PeggyERC20DeployedEvent struct {
	CosmosDenom   string
	TokenContract common.Address
	Name          string
	Symbol        string
	Decimals      uint8
	EventNonce    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterERC20DeployedEvent is a free log retrieval operation binding the contract event 0x82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7.
//
// Solidity: event ERC20DeployedEvent(string _cosmosDenom, address indexed _tokenContract, string _name, string _symbol, uint8 _decimals, uint256 _eventNonce)
func (_Peggy *PeggyFilterer) FilterERC20DeployedEvent(opts *bind.FilterOpts, _tokenContract []common.Address) (*PeggyERC20DeployedEventIterator, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}

	logs, sub, err := _Peggy.contract.FilterLogs(opts, "ERC20DeployedEvent", _tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &PeggyERC20DeployedEventIterator{contract: _Peggy.contract, event: "ERC20DeployedEvent", logs: logs, sub: sub}, nil
}

// WatchERC20DeployedEvent is a free log subscription operation binding the contract event 0x82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7.
//
// Solidity: event ERC20DeployedEvent(string _cosmosDenom, address indexed _tokenContract, string _name, string _symbol, uint8 _decimals, uint256 _eventNonce)
func (_Peggy *PeggyFilterer) WatchERC20DeployedEvent(opts *bind.WatchOpts, sink chan<- *PeggyERC20DeployedEvent, _tokenContract []common.Address) (event.Subscription, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}

	logs, sub, err := _Peggy.contract.WatchLogs(opts, "ERC20DeployedEvent", _tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggyERC20DeployedEvent)
				if err := _Peggy.contract.UnpackLog(event, "ERC20DeployedEvent", log); err != nil {
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

// ParseERC20DeployedEvent is a log parse operation binding the contract event 0x82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7.
//
// Solidity: event ERC20DeployedEvent(string _cosmosDenom, address indexed _tokenContract, string _name, string _symbol, uint8 _decimals, uint256 _eventNonce)
func (_Peggy *PeggyFilterer) ParseERC20DeployedEvent(log types.Log) (*PeggyERC20DeployedEvent, error) {
	event := new(PeggyERC20DeployedEvent)
	if err := _Peggy.contract.UnpackLog(event, "ERC20DeployedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeggyLogicCallEventIterator is returned from FilterLogicCallEvent and is used to iterate over the raw logs and unpacked data for LogicCallEvent events raised by the Peggy contract.
type PeggyLogicCallEventIterator struct {
	Event *PeggyLogicCallEvent // Event containing the contract specifics and raw log

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
func (it *PeggyLogicCallEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeggyLogicCallEvent)
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
		it.Event = new(PeggyLogicCallEvent)
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
func (it *PeggyLogicCallEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeggyLogicCallEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeggyLogicCallEvent represents a LogicCallEvent event raised by the Peggy contract.
type PeggyLogicCallEvent struct {
	InvalidationId    [32]byte
	InvalidationNonce *big.Int
	ReturnData        []byte
	EventNonce        *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogicCallEvent is a free log retrieval operation binding the contract event 0x7c2bb24f8e1b3725cb613d7f11ef97d9745cc97a0e40f730621c052d684077a1.
//
// Solidity: event LogicCallEvent(bytes32 _invalidationId, uint256 _invalidationNonce, bytes _returnData, uint256 _eventNonce)
func (_Peggy *PeggyFilterer) FilterLogicCallEvent(opts *bind.FilterOpts) (*PeggyLogicCallEventIterator, error) {

	logs, sub, err := _Peggy.contract.FilterLogs(opts, "LogicCallEvent")
	if err != nil {
		return nil, err
	}
	return &PeggyLogicCallEventIterator{contract: _Peggy.contract, event: "LogicCallEvent", logs: logs, sub: sub}, nil
}

// WatchLogicCallEvent is a free log subscription operation binding the contract event 0x7c2bb24f8e1b3725cb613d7f11ef97d9745cc97a0e40f730621c052d684077a1.
//
// Solidity: event LogicCallEvent(bytes32 _invalidationId, uint256 _invalidationNonce, bytes _returnData, uint256 _eventNonce)
func (_Peggy *PeggyFilterer) WatchLogicCallEvent(opts *bind.WatchOpts, sink chan<- *PeggyLogicCallEvent) (event.Subscription, error) {

	logs, sub, err := _Peggy.contract.WatchLogs(opts, "LogicCallEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeggyLogicCallEvent)
				if err := _Peggy.contract.UnpackLog(event, "LogicCallEvent", log); err != nil {
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

// ParseLogicCallEvent is a log parse operation binding the contract event 0x7c2bb24f8e1b3725cb613d7f11ef97d9745cc97a0e40f730621c052d684077a1.
//
// Solidity: event LogicCallEvent(bytes32 _invalidationId, uint256 _invalidationNonce, bytes _returnData, uint256 _eventNonce)
func (_Peggy *PeggyFilterer) ParseLogicCallEvent(log types.Log) (*PeggyLogicCallEvent, error) {
	event := new(PeggyLogicCallEvent)
	if err := _Peggy.contract.UnpackLog(event, "LogicCallEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Destination   [32]byte
	Amount        *big.Int
	EventNonce    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSendToCosmosEvent is a free log retrieval operation binding the contract event 0xd7767894d73c589daeca9643f445f03d7be61aad2950c117e7cbff4176fca7e4.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, bytes32 indexed _destination, uint256 _amount, uint256 _eventNonce)
func (_Peggy *PeggyFilterer) FilterSendToCosmosEvent(opts *bind.FilterOpts, _tokenContract []common.Address, _sender []common.Address, _destination [][32]byte) (*PeggySendToCosmosEventIterator, error) {

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

// WatchSendToCosmosEvent is a free log subscription operation binding the contract event 0xd7767894d73c589daeca9643f445f03d7be61aad2950c117e7cbff4176fca7e4.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, bytes32 indexed _destination, uint256 _amount, uint256 _eventNonce)
func (_Peggy *PeggyFilterer) WatchSendToCosmosEvent(opts *bind.WatchOpts, sink chan<- *PeggySendToCosmosEvent, _tokenContract []common.Address, _sender []common.Address, _destination [][32]byte) (event.Subscription, error) {

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

// ParseSendToCosmosEvent is a log parse operation binding the contract event 0xd7767894d73c589daeca9643f445f03d7be61aad2950c117e7cbff4176fca7e4.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, bytes32 indexed _destination, uint256 _amount, uint256 _eventNonce)
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

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReentrancyGuardBin is the compiled bytecode used for deploying new contracts.
var ReentrancyGuardBin = "0x6080604052348015600f57600080fd5b506001600055603f8060226000396000f3fe6080604052600080fdfea26469706673582212205687fc138b3ca5642581697badb2f0b5b3647cfdaebefce990bca7c0ce25096864736f6c63430008020033"

// DeployReentrancyGuard deploys a new Ethereum contract, binding an instance of ReentrancyGuard to it.
func DeployReentrancyGuard(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReentrancyGuard, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReentrancyGuardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// ReentrantERC20ABI is the input ABI used to generate the binding from.
const ReentrantERC20ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_peggyAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ReentrantERC20FuncSigs maps the 4-byte function signature to its string representation.
var ReentrantERC20FuncSigs = map[string]string{
	"a9059cbb": "transfer(address,uint256)",
}

// ReentrantERC20Bin is the compiled bytecode used for deploying new contracts.
var ReentrantERC20Bin = "0x608060405234801561001057600080fd5b506040516104f73803806104f783398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610082565b600060208284031215610065578081fd5b81516001600160a01b038116811461007b578182fd5b9392505050565b610466806100916000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a9059cbb14610030575b600080fd5b61004361003e366004610187565b610057565b604051901515815260200160405180910390f35b604080516000808252602080830182815283850183815260608086018581526101a087018852608080880183905260a080890184905260c0808a0185905260e0808b01869052610100808c018b9052610120808d018890526101408d018c90526101608d018c90526101808d018c90528d519081018e52888152808b018d9052808e018990529687018c90529386018a90529185018490528401889052830187905282018690528554885187815295860198899052630612364160e11b909852949692949193919287926001600160a01b0390911690630c246c8290610149908990889087908b808960248401610372565b600060405180830381600087803b15801561016357600080fd5b505af1158015610177573d6000803e3d6000fd5b5050505050505050505092915050565b60008060408385031215610199578182fd5b82356001600160a01b03811681146101af578283fd5b946020939093013593505050565b6000815180845260208085019450808401835b838110156101f55781516001600160a01b0316875295820195908201906001016101d0565b509495945050505050565b6000815180845260208085019450808401835b838110156101f557815187529582019590820190600101610213565b6000815180845260208085019450808401835b838110156101f557815160ff1687529582019590820190600101610242565b60008151808452815b818110156102865760208185018101518683018201520161026a565b818111156102975782602083870101525b50601f01601f19169290920160200192915050565b600061012082518185526102c282860182610200565b915050602083015184820360208601526102dc82826101bd565b915050604083015184820360408601526102f68282610200565b9150506060830151848203606086015261031082826101bd565b915050608083015161032d60808601826001600160a01b03169052565b5060a083015184820360a08601526103458282610261565b91505060c083015160c085015260e083015160e08501526101008084015181860152508091505092915050565b60e08082528851908201819052600090602090610100840190828c01845b828110156103b75781516001600160a01b0316845260208401935090840190600101610390565b505050838103828501526103cb818b610200565b91505087604084015282810360608401526103e6818861022f565b905082810360808401526103fa8187610200565b905082810360a084015261040e8186610200565b905082810360c084015261042281856102ac565b9a995050505050505050505056fea2646970667358221220cafdb574fa1b1954b6f16fa7feb0db00a700d2ec195a5e8a7c41915a82acffe664736f6c63430008020033"

// DeployReentrantERC20 deploys a new Ethereum contract, binding an instance of ReentrantERC20 to it.
func DeployReentrantERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _peggyAddress common.Address) (common.Address, *types.Transaction, *ReentrantERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrantERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReentrantERC20Bin), backend, _peggyAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReentrantERC20{ReentrantERC20Caller: ReentrantERC20Caller{contract: contract}, ReentrantERC20Transactor: ReentrantERC20Transactor{contract: contract}, ReentrantERC20Filterer: ReentrantERC20Filterer{contract: contract}}, nil
}

// ReentrantERC20 is an auto generated Go binding around an Ethereum contract.
type ReentrantERC20 struct {
	ReentrantERC20Caller     // Read-only binding to the contract
	ReentrantERC20Transactor // Write-only binding to the contract
	ReentrantERC20Filterer   // Log filterer for contract events
}

// ReentrantERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrantERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrantERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrantERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrantERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrantERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrantERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrantERC20Session struct {
	Contract     *ReentrantERC20   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrantERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrantERC20CallerSession struct {
	Contract *ReentrantERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ReentrantERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrantERC20TransactorSession struct {
	Contract     *ReentrantERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ReentrantERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrantERC20Raw struct {
	Contract *ReentrantERC20 // Generic contract binding to access the raw methods on
}

// ReentrantERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrantERC20CallerRaw struct {
	Contract *ReentrantERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ReentrantERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrantERC20TransactorRaw struct {
	Contract *ReentrantERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrantERC20 creates a new instance of ReentrantERC20, bound to a specific deployed contract.
func NewReentrantERC20(address common.Address, backend bind.ContractBackend) (*ReentrantERC20, error) {
	contract, err := bindReentrantERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrantERC20{ReentrantERC20Caller: ReentrantERC20Caller{contract: contract}, ReentrantERC20Transactor: ReentrantERC20Transactor{contract: contract}, ReentrantERC20Filterer: ReentrantERC20Filterer{contract: contract}}, nil
}

// NewReentrantERC20Caller creates a new read-only instance of ReentrantERC20, bound to a specific deployed contract.
func NewReentrantERC20Caller(address common.Address, caller bind.ContractCaller) (*ReentrantERC20Caller, error) {
	contract, err := bindReentrantERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrantERC20Caller{contract: contract}, nil
}

// NewReentrantERC20Transactor creates a new write-only instance of ReentrantERC20, bound to a specific deployed contract.
func NewReentrantERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ReentrantERC20Transactor, error) {
	contract, err := bindReentrantERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrantERC20Transactor{contract: contract}, nil
}

// NewReentrantERC20Filterer creates a new log filterer instance of ReentrantERC20, bound to a specific deployed contract.
func NewReentrantERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ReentrantERC20Filterer, error) {
	contract, err := bindReentrantERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrantERC20Filterer{contract: contract}, nil
}

// bindReentrantERC20 binds a generic wrapper to an already deployed contract.
func bindReentrantERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrantERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrantERC20 *ReentrantERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrantERC20.Contract.ReentrantERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrantERC20 *ReentrantERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrantERC20.Contract.ReentrantERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrantERC20 *ReentrantERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrantERC20.Contract.ReentrantERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrantERC20 *ReentrantERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrantERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrantERC20 *ReentrantERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrantERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrantERC20 *ReentrantERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrantERC20.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ReentrantERC20 *ReentrantERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReentrantERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ReentrantERC20 *ReentrantERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReentrantERC20.Contract.Transfer(&_ReentrantERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ReentrantERC20 *ReentrantERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReentrantERC20.Contract.Transfer(&_ReentrantERC20.TransactOpts, recipient, amount)
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ac354f89b7d1e539f1399f03f55aa2733ed1b517fea56584883d3bb648986a7964736f6c63430008020033"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}
