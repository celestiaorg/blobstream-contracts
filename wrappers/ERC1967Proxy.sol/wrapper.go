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
	_ = abi.ConvertType
)

// WrappersMetaData contains all meta data concerning the Wrappers contract.
var WrappersMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_logic\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405261001561000f6101b5565b906102f9565b61001d61002b565b61010c610853823961010c90f35b60405190565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061005990610031565b810190811060018060401b0382111761007157604052565b61003b565b9061008961008261002b565b928361004f565b565b5f80fd5b5f80fd5b60018060a01b031690565b6100a790610093565b90565b6100b38161009e565b036100ba57565b5f80fd5b905051906100cb826100aa565b565b5f80fd5b5f80fd5b60018060401b0381116100f1576100ed602091610031565b0190565b61003b565b5f5b838110610108575050905f910152565b8060209183015181850152016100f8565b9092919261012e610129826100d5565b610076565b9381855260208501908284011161014a57610148926100f6565b565b6100d1565b9080601f8301121561016d5781602061016a93519101610119565b90565b6100cd565b9190916040818403126101b05761018b835f83016100be565b92602082015160018060401b0381116101ab576101a8920161014f565b90565b61008f565b61008b565b6101d361095f803803806101c881610076565b928339810190610172565b9091565b90565b90565b5f1b90565b6101f66101f16101fb926101d7565b6101dd565b6101da565b90565b6102277f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6101e2565b90565b5f1c90565b90565b90565b61024961024461024e9261022f565b610232565b61022f565b90565b61025d6102629161022a565b610235565b90565b90565b61027c61027761028192610265565b610232565b61022f565b90565b634e487b7160e01b5f52601160045260245ffd5b6102a76102ad9193929361022f565b9261022f565b82039182116102b857565b610284565b6102d16102cc6102d69261022f565b6101dd565b6101da565b90565b634e487b7160e01b5f52600160045260245ffd5b156102f457565b6102d9565b9061036c916103646103096101fe565b61035e61035861035361034e61033e7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd610251565b6103486001610268565b90610298565b6102bd565b6101da565b916101da565b146102ed565b905f91610391565b565b5190565b90565b61038961038461038e92610372565b610232565b61022f565b90565b9161039b8361041b565b6103a48261036e565b6103b66103b05f610375565b9161022f565b119081156103da575b506103c9575b5050565b6103d291610514565b505f806103c5565b90505f6103bf565b6103f66103f16103fb92610093565b610232565b610093565b90565b610407906103e2565b90565b610413906103fe565b90565b5f0190565b61042481610623565b61044e7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b9161040a565b9061045761002b565b8061046181610416565b0390a2565b606090565b60018060401b03811161048757610483602091610031565b0190565b61003b565b9061049e6104998361046b565b610076565b918252565b60207f206661696c656400000000000000000000000000000000000000000000000000917f416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c5f8201520152565b6104fa602761048c565b90610507602083016104a3565b565b6105116104f0565b90565b9061053191610521610466565b509061052b610509565b91610739565b90565b60209181520190565b60207f6f74206120636f6e747261637400000000000000000000000000000000000000917f455243313936373a206e657720696d706c656d656e746174696f6e206973206e5f8201520152565b610597602d604092610534565b6105a08161053d565b0190565b6105b99060208101905f81830391015261058a565b90565b156105c357565b6105cb61002b565b62461bcd60e51b8152806105e1600482016105a4565b0390fd5b906105f660018060a01b03916101dd565b9181191691161790565b90565b9061061861061361061f9261040a565b610600565b82546105e5565b9055565b61064f906106386106338261077c565b6105bc565b5f6106496106446101fe565b61079c565b01610603565b565b60207f6e74726163740000000000000000000000000000000000000000000000000000917f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f5f8201520152565b6106ab6026604092610534565b6106b481610651565b0190565b6106cd9060208101905f81830391015261069e565b90565b156106d757565b6106df61002b565b62461bcd60e51b8152806106f5600482016106b8565b0390fd5b9061070b610706836100d5565b610076565b918252565b3d5f1461072b576107203d6106f9565b903d5f602084013e5b565b610733610466565b90610729565b5f61077593928192610749610466565b5061075b6107568261077c565b6106d0565b90602081019051915af49061076e610710565b90916107ec565b90565b5f90565b610784610778565b503b6107986107925f610375565b9161022f565b1190565b90565b5190565b6107c26107cb6020936107d0936107b98161079f565b93848093610534565b958691016100f6565b610031565b0190565b6107e99160208201915f8184039101526107a3565b90565b9190916107f7610466565b505f14610802575090565b61080b8261036e565b61081d6108175f610375565b9161022f565b115f1461082d5750805190602001fd5b61084e9061083961002b565b91829162461bcd60e51b8352600483016107d4565b0390fdfe6080604052600a6012565b6022565b5f90565b6018600e565b50601f60b5565b90565b5f8091368280378136915af43d5f803e5f14603b573d5ff35b3d5ffd5b90565b90565b5f1b90565b60596055605d92603f565b6045565b6042565b90565b60877f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc604a565b90565b5f1c90565b60018060a01b031690565b60a360a791608a565b608f565b90565b60b29054609a565b90565b60bb600e565b5060d05f60cb60c76060565b60d3565b0160aa565b90565b9056fea2646970667358221220d931939e53377533841dcd2138196534e48e81e93470d07d5a0cbd05246a364164736f6c63430008160033",
}

// WrappersABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappersMetaData.ABI instead.
var WrappersABI = WrappersMetaData.ABI

// WrappersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WrappersMetaData.Bin instead.
var WrappersBin = WrappersMetaData.Bin

// DeployWrappers deploys a new Ethereum contract, binding an instance of Wrappers to it.
func DeployWrappers(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _data []byte) (common.Address, *types.Transaction, *Wrappers, error) {
	parsed, err := WrappersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrappersBin), backend, _logic, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wrappers{WrappersCaller: WrappersCaller{contract: contract}, WrappersTransactor: WrappersTransactor{contract: contract}, WrappersFilterer: WrappersFilterer{contract: contract}}, nil
}

// Wrappers is an auto generated Go binding around an Ethereum contract.
type Wrappers struct {
	WrappersCaller     // Read-only binding to the contract
	WrappersTransactor // Write-only binding to the contract
	WrappersFilterer   // Log filterer for contract events
}

// WrappersCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappersFilterer is an auto generated log filtering Go binding around Ethereum contract events.
type WrappersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappersSession struct {
	Contract     *Wrappers         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrappersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappersCallerSession struct {
	Contract *WrappersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WrappersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappersTransactorSession struct {
	Contract     *WrappersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WrappersRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappersRaw struct {
	Contract *Wrappers // Generic contract binding to access the raw methods on
}

// WrappersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappersCallerRaw struct {
	Contract *WrappersCaller // Generic read-only contract binding to access the raw methods on
}

// WrappersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappersTransactorRaw struct {
	Contract *WrappersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappers creates a new instance of Wrappers, bound to a specific deployed contract.
func NewWrappers(address common.Address, backend bind.ContractBackend) (*Wrappers, error) {
	contract, err := bindWrappers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wrappers{WrappersCaller: WrappersCaller{contract: contract}, WrappersTransactor: WrappersTransactor{contract: contract}, WrappersFilterer: WrappersFilterer{contract: contract}}, nil
}

// NewWrappersCaller creates a new read-only instance of Wrappers, bound to a specific deployed contract.
func NewWrappersCaller(address common.Address, caller bind.ContractCaller) (*WrappersCaller, error) {
	contract, err := bindWrappers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappersCaller{contract: contract}, nil
}

// NewWrappersTransactor creates a new write-only instance of Wrappers, bound to a specific deployed contract.
func NewWrappersTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappersTransactor, error) {
	contract, err := bindWrappers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappersTransactor{contract: contract}, nil
}

// NewWrappersFilterer creates a new log filterer instance of Wrappers, bound to a specific deployed contract.
func NewWrappersFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappersFilterer, error) {
	contract, err := bindWrappers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappersFilterer{contract: contract}, nil
}

// bindWrappers binds a generic wrapper to an already deployed contract.
func bindWrappers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WrappersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wrappers *WrappersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wrappers.Contract.WrappersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wrappers *WrappersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wrappers.Contract.WrappersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wrappers *WrappersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wrappers.Contract.WrappersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wrappers *WrappersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wrappers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wrappers *WrappersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wrappers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wrappers *WrappersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wrappers.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Wrappers *WrappersTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Wrappers.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Wrappers *WrappersSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Wrappers.Contract.Fallback(&_Wrappers.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Wrappers *WrappersTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Wrappers.Contract.Fallback(&_Wrappers.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wrappers *WrappersTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wrappers.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wrappers *WrappersSession) Receive() (*types.Transaction, error) {
	return _Wrappers.Contract.Receive(&_Wrappers.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wrappers *WrappersTransactorSession) Receive() (*types.Transaction, error) {
	return _Wrappers.Contract.Receive(&_Wrappers.TransactOpts)
}

// WrappersAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Wrappers contract.
type WrappersAdminChangedIterator struct {
	Event *WrappersAdminChanged // Event containing the contract specifics and raw log

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
func (it *WrappersAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappersAdminChanged)
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
		it.Event = new(WrappersAdminChanged)
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
func (it *WrappersAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappersAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappersAdminChanged represents a AdminChanged event raised by the Wrappers contract.
type WrappersAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Wrappers *WrappersFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*WrappersAdminChangedIterator, error) {

	logs, sub, err := _Wrappers.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &WrappersAdminChangedIterator{contract: _Wrappers.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Wrappers *WrappersFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *WrappersAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Wrappers.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappersAdminChanged)
				if err := _Wrappers.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Wrappers *WrappersFilterer) ParseAdminChanged(log types.Log) (*WrappersAdminChanged, error) {
	event := new(WrappersAdminChanged)
	if err := _Wrappers.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappersBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Wrappers contract.
type WrappersBeaconUpgradedIterator struct {
	Event *WrappersBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *WrappersBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappersBeaconUpgraded)
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
		it.Event = new(WrappersBeaconUpgraded)
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
func (it *WrappersBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappersBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappersBeaconUpgraded represents a BeaconUpgraded event raised by the Wrappers contract.
type WrappersBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Wrappers *WrappersFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*WrappersBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Wrappers.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &WrappersBeaconUpgradedIterator{contract: _Wrappers.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Wrappers *WrappersFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *WrappersBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Wrappers.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappersBeaconUpgraded)
				if err := _Wrappers.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Wrappers *WrappersFilterer) ParseBeaconUpgraded(log types.Log) (*WrappersBeaconUpgraded, error) {
	event := new(WrappersBeaconUpgraded)
	if err := _Wrappers.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappersUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Wrappers contract.
type WrappersUpgradedIterator struct {
	Event *WrappersUpgraded // Event containing the contract specifics and raw log

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
func (it *WrappersUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappersUpgraded)
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
		it.Event = new(WrappersUpgraded)
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
func (it *WrappersUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappersUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappersUpgraded represents a Upgraded event raised by the Wrappers contract.
type WrappersUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Wrappers *WrappersFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WrappersUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Wrappers.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WrappersUpgradedIterator{contract: _Wrappers.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Wrappers *WrappersFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WrappersUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Wrappers.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappersUpgraded)
				if err := _Wrappers.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Wrappers *WrappersFilterer) ParseUpgraded(log types.Log) (*WrappersUpgraded, error) {
	event := new(WrappersUpgraded)
	if err := _Wrappers.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
