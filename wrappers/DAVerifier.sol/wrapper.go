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

// AttestationProof is an auto generated low-level Go binding around an user-defined struct.
type AttestationProof struct {
	TupleRootNonce *big.Int
	Tuple          DataRootTuple
	Proof          BinaryMerkleProof
}

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

// NamespaceMerkleMultiproof is an auto generated low-level Go binding around an user-defined struct.
type NamespaceMerkleMultiproof struct {
	BeginKey  *big.Int
	EndKey    *big.Int
	SideNodes []NamespaceNode
}

// NamespaceNode is an auto generated low-level Go binding around an user-defined struct.
type NamespaceNode struct {
	Min    [8]byte
	Max    [8]byte
	Digest [32]byte
}

// SharesProof is an auto generated low-level Go binding around an user-defined struct.
type SharesProof struct {
	Data             [][]byte
	ShareProofs      []NamespaceMerkleMultiproof
	NamespaceID      [8]byte
	RowsRoots        []NamespaceNode
	RowsProofs       []BinaryMerkleProof
	AttestationProof AttestationProof
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207eadb80a7580ae404794766e2bdef99e27e061ef3413cdabf850b8262fd6415964736f6c63430008090033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208f868974afcbbf9b36869e2284f3eb5ebe16ddff65eb1593f2f10039702dce5664736f6c63430008090033",
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

// DAVerifierMetaData contains all meta data concerning the DAVerifier contract.
var DAVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidDataRootTupleToDataRootTupleRootProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"InvalidRowsToDataRootProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"InvalidSharesToRowsProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalDataLengthAndNumberOfSharesProofs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalRowsProofsAndRowsRootsNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalShareProofsAndRowsRootsNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractQuantumGravityBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes8\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes8\",\"name\":\"namespaceID\",\"type\":\"bytes8\"},{\"components\":[{\"internalType\":\"bytes8\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowsRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowsProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e78cea92": "bridge()",
		"1f7b9136": "verify((bytes[],(uint256,uint256,(bytes8,bytes8,bytes32)[])[],bytes8,(bytes8,bytes8,bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))),bytes32)",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051611c3a380380611c3a83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b611ba7806100936000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631f7b91361461003b578063e78cea9214610063575b600080fd5b61004e610049366004611847565b61008e565b60405190151581526020015b60405180910390f35b600054610076906001600160a01b031681565b6040516001600160a01b03909116815260200161005a565b6000805460a0840151805160208201516040928301519251631f3302a960e01b81526001600160a01b0390941693631f3302a9936100d093929160040161195c565b60206040518083038186803b1580156100e857600080fd5b505afa1580156100fc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061012091906119df565b61013d5760405163320f037560e21b815260040160405180910390fd5b8260600151518360800151511461016757604051634f17331d60e01b815260040160405180910390fd5b60005b8360800151518110156102825760008460600151828151811061018f5761018f611a01565b602002602001015160000151856060015183815181106101b1576101b1611a01565b602002602001015160200151866060015184815181106101d3576101d3611a01565b602002602001015160400151604051602001610211939291906001600160c01b03199384168152919092166008820152601081019190915260300190565b604051602081830303815290604052905061024a848660800151848151811061023c5761023c611a01565b6020026020010151836104a1565b61026f576040516301dd096960e21b8152600481018390526024015b60405180910390fd5b508061027a81611a2d565b91505061016a565b50826060015151836020015151146102ad57604051636031acbb60e01b815260040160405180910390fd5b6000805b84602001515181101561032a57846020015181815181106102d4576102d4611a01565b602002602001015160000151856020015182815181106102f6576102f6611a01565b60200260200101516020015161030c9190611a48565b6103169083611a5f565b91508061032281611a2d565b9150506102b1565b50835151811461034d5760405163efc454a560e01b815260040160405180910390fd5b6000805b8560200151518110156104935760008660200151828151811061037657610376611a01565b6020026020010151600001518760200151838151811061039857610398611a01565b6020026020010151602001516103ae9190611a48565b90506000604051806060016040528089604001516001600160c01b031916815260200189604001516001600160c01b0319168152602001896060015185815181106103fb576103fb611a01565b6020026020010151604001518152509050610452818960200151858151811061042657610426611a01565b60200260200101518a6040015161044d8c6000015189888b6104489190611a5f565b610711565b6107c9565b6104725760405163cef8a4cb60e01b815260048101849052602401610266565b61047c8285611a5f565b93505050808061048b90611a2d565b915050610351565b506001925050505b92915050565b600060018360400151116104c457825151156104bf5750600061070a565b6104e6565b6104d6836020015184604001516108a6565b835151146104e65750600061070a565b82604001518360200151106104fd5750600061070a565b60006105088361092d565b84515190915061053257836040015160011415610528578414905061070a565b600091505061070a565b60208401516001905b60208601516000906001841b90610553908290611a77565b61055d9190611a99565b90506000600161056f81861b84611a5f565b6105799190611a48565b90508760400151811061058d575050610653565b91508161059b600185611a48565b885151116105b15760009550505050505061070a565b6105bc600185611a48565b6001901b8289602001516105d09190611a48565b101561060d5787516106069086906105e9600188611a48565b815181106105f9576105f9611a01565b60200260200101516109a2565b945061063f565b875161063c9061061e600187611a48565b8151811061062e5761062e611a01565b6020026020010151866109a2565b94505b61064a600185611a5f565b9350505061053b565b600186604001516106649190611a48565b81146106ac57610675600183611a48565b86515111610689576000935050505061070a565b855161069c9084906105e9600186611a48565b92506106a9600183611a5f565b91505b8551516106ba600184611a48565b10156107035785516106ef906106d1600185611a48565b815181106106e1576106e1611a01565b6020026020010151846109a2565b92506106fc600183611a5f565b91506106ac565b5050841490505b9392505050565b6060600061071f8484611a48565b6001600160401b038111156107365761073661134b565b60405190808252806020026020018201604052801561076957816020015b60608152602001906001900390816107545790505b509050835b838110156107c05785818151811061078857610788611a01565b60200260200101518282815181106107a2576107a2611a01565b602002602001018190525080806107b890611a2d565b91505061076e565b50949350505050565b60008082516001600160401b038111156107e5576107e561134b565b60405190808252806020026020018201604052801561083057816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816108035790505b50905060005b8351811015610890576108628585838151811061085557610855611a01565b6020026020010151610a20565b82828151811061087457610874611a01565b60200260200101819052508061088990611a2d565b9050610836565b5061089c868683610ade565b9695505050505050565b60006108b182610ca7565b6108bd90610100611a48565b905060006108cc600183611a48565b6001901b90506001816108df9190611a48565b84116108eb575061049b565b80600114156108fe57600191505061049b565b61091a61090b8286611a48565b6109158386611a48565b6108a6565b610925906001611a5f565b91505061049b565b60006002600060f81b83604051602001610948929190611af3565b60408051601f198184030181529082905261096291611b17565b602060405180830381855afa15801561097f573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061049b9190611b23565b604051600160f81b6020820152602181018390526041810182905260009060029060610160408051601f19818403018152908290526109e091611b17565b602060405180830381855afa1580156109fd573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061070a9190611b23565b604080516060810182526000808252602082018190529181019190915260006002600060f81b8585604051602001610a5a93929190611b3c565b60408051601f1981840301815290829052610a7491611b17565b602060405180830381855afa158015610a91573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190610ab49190611b23565b604080516060810182526001600160c01b0319969096168087526020870152850152509192915050565b6000808360400151516001600160401b03811115610afe57610afe61134b565b604051908082528060200260200182016040528015610b4957816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610b1c5790505b5090506000805b85518214801590610b65575085604001515181105b15610bd5576000610b7a838860000151610cd4565b905086604001518281518110610b9257610b92611a01565b6020026020010151848381518110610bac57610bac611a01565b6020908102919091010152610bc18184611a5f565b92505080610bce90611a2d565b9050610b50565b506000610be58660200151610d25565b610bf0906002611a99565b90506001811015610bff575060015b6000610c396040518060c0016040528089815260200188815260200160008152602001848152602001600081526020016000815250610d67565b80516020820151919250905b886040015151811015610c8f57610c7d83600001518a604001518381518110610c7057610c70611a01565b6020026020010151610f88565b9150610c8881611a2d565b9050610c45565b50610c9a818a611102565b9998505050505050505050565b60005b81816001901b1015610cc857610cc1600182611a5f565b9050610caa565b61049b81610100611a48565b600080610ce084611154565b905060006001610cf8610cf38787611a48565b611187565b610d029190611a48565b905080821115610d19576001901b915061049b9050565b506001901b9392505050565b60006001821015610d3557600080fd5b6000610d4083611187565b90506000610d4f600183611a48565b6001901b90508381141561070a5760011c9392505050565b610d6f611310565b81604001518260600151610d839190611a48565b60011415610deb57604082015182515111801590610da957508151602001516040830151105b15610dcb5761049b82602001518360a0015184602001515185608001516111aa565b61049b826000015160400151836080015184606001518560a001516111ef565b8151516060830151111580610e095750815160200151604083015110155b15610e2e5761049b826000015160400151836080015184606001518560a001516111ef565b6000610e4c83604001518460600151610e479190611a48565b610d25565b90506000610ea86040518060c00160405280866000015181526020018660200151815260200186604001518152602001848760400151610e8c9190611a5f565b8152602001866080015181526020018660a00151815250610d67565b90506000610f046040518060c001604052808760000151815260200187602001518152602001858860400151610ede9190611a5f565b815260200187606001518152602001846020015181526020018460400151815250610d67565b905080606001511515600115151415610f4657506040805160808101825282518152602080840151908201529181015190820152600060608201529392505050565b81518151600091610f5691610f88565b604080516080810182529182526020848101519083015292830151928101929092525060006060820152949350505050565b604080516060810182526000808252602082018190529181019190915282518251600091610fb591611236565b84519091506000906001600160c01b03199081161415610fde57506001600160c01b0319611012565b83516001600160c01b03199081161415610ffd57506020840151611012565b61100f85602001518560200151611255565b90505b84516020808701516040808901518851898501518a8401519351600160f81b968101969096526001600160c01b03199687166021870152938616602986015260318501919091528416605184015292166059820152606181019190915260009060029060810160408051601f198184030181529082905261109291611b17565b602060405180830381855afa1580156110af573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906110d29190611b23565b604080516060810182526001600160c01b0319958616815293909416602084015292820192909252949350505050565b805182516000916001600160c01b0319918216911614801561113f575081602001516001600160c01b03191683602001516001600160c01b031916145b801561070a5750506040908101519101511490565b60006001815b83158015906111695750818416155b1561070a578061117881611a2d565b915050600184901c935061115a565b6000805b821561049b578061119b81611a2d565b915050600183901c925061118b565b6111b2611310565b60008060006111c288888861126d565b60408051608081018252938452602084019890985296820152941515606086015250929695505050505050565b6111f7611310565b600080600061120788888861126d565b604080516080810182529384526020840192909252908201879052151560608201529350505050949350505050565b600060c082811c9084901c101561124e57508161049b565b508061049b565b600060c082811c9084901c111561124e57508161049b565b604080516060810182526000808252602082018190529181019190915260008085516000148061129e575085518510155b806112a95750838510155b156112d757505060408051606081018252600080825260208201819052918101919091529050826001611307565b8585815181106112e9576112e9611a01565b60200260200101518560016112fe9190611a5f565b60009250925092505b93509350939050565b6040805160e08101825260006080820181815260a0830182905260c08301829052825260208201819052918101829052606081019190915290565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156113835761138361134b565b60405290565b60405160c081016001600160401b03811182821017156113835761138361134b565b604051601f8201601f191681016001600160401b03811182821017156113d3576113d361134b565b604052919050565b60006001600160401b038211156113f4576113f461134b565b5060051b60200190565b6000601f838184011261141057600080fd5b82356020611425611420836113db565b6113ab565b82815260059290921b8501810191818101908784111561144457600080fd5b8287015b848110156114da5780356001600160401b03808211156114685760008081fd5b818a0191508a603f83011261147d5760008081fd5b858201356040828211156114935761149361134b565b6114a4828b01601f191689016113ab565b92508183528c818386010111156114bb5760008081fd5b8181850189850137506000908201870152845250918301918301611448565b50979650505050505050565b80356001600160c01b0319811681146114fe57600080fd5b919050565b600082601f83011261151457600080fd5b81356020611524611420836113db565b8281526060928302850182019282820191908785111561154357600080fd5b8387015b858110156115985781818a03121561155f5760008081fd5b611567611361565b611570826114e6565b815261157d8683016114e6565b81870152604082810135908201528452928401928101611547565b5090979650505050505050565b600082601f8301126115b657600080fd5b813560206115c6611420836113db565b82815260059290921b840181019181810190868411156115e557600080fd5b8286015b848110156116705780356001600160401b03808211156116095760008081fd5b908801906060828b03601f19018113156116235760008081fd5b61162b611361565b838801358152604080850135828a015291840135918383111561164e5760008081fd5b61165c8d8a85880101611503565b9082015286525050509183019183016115e9565b509695505050505050565b60006060828403121561168d57600080fd5b611695611361565b905081356001600160401b038111156116ad57600080fd5b8201601f810184136116be57600080fd5b803560206116ce611420836113db565b82815260059290921b830181019181810190878411156116ed57600080fd5b938201935b8385101561170b578435825293820193908201906116f2565b808652505080850135818501525050506040820135604082015292915050565b600082601f83011261173c57600080fd5b8135602061174c611420836113db565b82815260059290921b8401810191818101908684111561176b57600080fd5b8286015b848110156116705780356001600160401b0381111561178e5760008081fd5b61179c8986838b010161167b565b84525091830191830161176f565b600081830360808112156117bd57600080fd5b6117c5611361565b915082358252604080601f19830112156117de57600080fd5b805191508082016001600160401b0383821081831117156118015761180161134b565b81835260208601358452828601356020850152836020860152606086013593508084111561182e57600080fd5b505061183c8583860161167b565b908301525092915050565b6000806040838503121561185a57600080fd5b82356001600160401b038082111561187157600080fd5b9084019060c0828703121561188557600080fd5b61188d611389565b82358281111561189c57600080fd5b6118a8888286016113fe565b8252506020830135828111156118bd57600080fd5b6118c9888286016115a5565b6020830152506118db604084016114e6565b60408201526060830135828111156118f257600080fd5b6118fe88828601611503565b60608301525060808301358281111561191657600080fd5b6119228882860161172b565b60808301525060a08301358281111561193a57600080fd5b611946888286016117aa565b60a0830152509660209590950135955050505050565b838152600060208451818401528085015160408401526080606084015260e08301845160606080860152818151808452610100870191508483019350600092505b808310156119bd578351825292840192600192909201919084019061199d565b509286015160a0860152505060409093015160c0909201919091525092915050565b6000602082840312156119f157600080fd5b8151801515811461070a57600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611a4157611a41611a17565b5060010190565b600082821015611a5a57611a5a611a17565b500390565b60008219821115611a7257611a72611a17565b500190565b600082611a9457634e487b7160e01b600052601260045260246000fd5b500490565b6000816000190483118215151615611ab357611ab3611a17565b500290565b6000815160005b81811015611ad95760208185018101518683015201611abf565b81811115611ae8576000828601525b509290920192915050565b6001600160f81b0319831681526000611b0f6001830184611ab8565b949350505050565b600061070a8284611ab8565b600060208284031215611b3557600080fd5b5051919050565b6001600160f81b0319841681526001600160c01b0319831660018201526000611b686009830184611ab8565b9594505050505056fea2646970667358221220cd273ffe145ce7cb5ca91014ded5b93e85a0f0542307e64ff6e5a0dc7cef57e864736f6c63430008090033",
}

// DAVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use DAVerifierMetaData.ABI instead.
var DAVerifierABI = DAVerifierMetaData.ABI

// Deprecated: Use DAVerifierMetaData.Sigs instead.
// DAVerifierFuncSigs maps the 4-byte function signature to its string representation.
var DAVerifierFuncSigs = DAVerifierMetaData.Sigs

// DAVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DAVerifierMetaData.Bin instead.
var DAVerifierBin = DAVerifierMetaData.Bin

// DeployDAVerifier deploys a new Ethereum contract, binding an instance of DAVerifier to it.
func DeployDAVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *DAVerifier, error) {
	parsed, err := DAVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DAVerifierBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DAVerifier{DAVerifierCaller: DAVerifierCaller{contract: contract}, DAVerifierTransactor: DAVerifierTransactor{contract: contract}, DAVerifierFilterer: DAVerifierFilterer{contract: contract}}, nil
}

// DAVerifier is an auto generated Go binding around an Ethereum contract.
type DAVerifier struct {
	DAVerifierCaller     // Read-only binding to the contract
	DAVerifierTransactor // Write-only binding to the contract
	DAVerifierFilterer   // Log filterer for contract events
}

// DAVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAVerifierSession struct {
	Contract     *DAVerifier       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAVerifierCallerSession struct {
	Contract *DAVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DAVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAVerifierTransactorSession struct {
	Contract     *DAVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DAVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type DAVerifierRaw struct {
	Contract *DAVerifier // Generic contract binding to access the raw methods on
}

// DAVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAVerifierCallerRaw struct {
	Contract *DAVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// DAVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAVerifierTransactorRaw struct {
	Contract *DAVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAVerifier creates a new instance of DAVerifier, bound to a specific deployed contract.
func NewDAVerifier(address common.Address, backend bind.ContractBackend) (*DAVerifier, error) {
	contract, err := bindDAVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAVerifier{DAVerifierCaller: DAVerifierCaller{contract: contract}, DAVerifierTransactor: DAVerifierTransactor{contract: contract}, DAVerifierFilterer: DAVerifierFilterer{contract: contract}}, nil
}

// NewDAVerifierCaller creates a new read-only instance of DAVerifier, bound to a specific deployed contract.
func NewDAVerifierCaller(address common.Address, caller bind.ContractCaller) (*DAVerifierCaller, error) {
	contract, err := bindDAVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAVerifierCaller{contract: contract}, nil
}

// NewDAVerifierTransactor creates a new write-only instance of DAVerifier, bound to a specific deployed contract.
func NewDAVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*DAVerifierTransactor, error) {
	contract, err := bindDAVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAVerifierTransactor{contract: contract}, nil
}

// NewDAVerifierFilterer creates a new log filterer instance of DAVerifier, bound to a specific deployed contract.
func NewDAVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*DAVerifierFilterer, error) {
	contract, err := bindDAVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAVerifierFilterer{contract: contract}, nil
}

// bindDAVerifier binds a generic wrapper to an already deployed contract.
func bindDAVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DAVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAVerifier *DAVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAVerifier.Contract.DAVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAVerifier *DAVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAVerifier.Contract.DAVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAVerifier *DAVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAVerifier.Contract.DAVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAVerifier *DAVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAVerifier *DAVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAVerifier *DAVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAVerifier.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_DAVerifier *DAVerifierCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAVerifier.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_DAVerifier *DAVerifierSession) Bridge() (common.Address, error) {
	return _DAVerifier.Contract.Bridge(&_DAVerifier.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_DAVerifier *DAVerifierCallerSession) Bridge() (common.Address, error) {
	return _DAVerifier.Contract.Bridge(&_DAVerifier.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x1f7b9136.
//
// Solidity: function verify((bytes[],(uint256,uint256,(bytes8,bytes8,bytes32)[])[],bytes8,(bytes8,bytes8,bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof, bytes32 _root) view returns(bool)
func (_DAVerifier *DAVerifierCaller) Verify(opts *bind.CallOpts, _sharesProof SharesProof, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _DAVerifier.contract.Call(opts, &out, "verify", _sharesProof, _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x1f7b9136.
//
// Solidity: function verify((bytes[],(uint256,uint256,(bytes8,bytes8,bytes32)[])[],bytes8,(bytes8,bytes8,bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof, bytes32 _root) view returns(bool)
func (_DAVerifier *DAVerifierSession) Verify(_sharesProof SharesProof, _root [32]byte) (bool, error) {
	return _DAVerifier.Contract.Verify(&_DAVerifier.CallOpts, _sharesProof, _root)
}

// Verify is a free data retrieval call binding the contract method 0x1f7b9136.
//
// Solidity: function verify((bytes[],(uint256,uint256,(bytes8,bytes8,bytes32)[])[],bytes8,(bytes8,bytes8,bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof, bytes32 _root) view returns(bool)
func (_DAVerifier *DAVerifierCallerSession) Verify(_sharesProof SharesProof, _root [32]byte) (bool, error) {
	return _DAVerifier.Contract.Verify(&_DAVerifier.CallOpts, _sharesProof, _root)
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220333ac9e98839ed5aa8c16d658fffd68e1cdbc5fb35ef9efab71b442b1f20217b64736f6c63430008090033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"_tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_IDAOracle *IDAOracleCaller) VerifyAttestation(opts *bind.CallOpts, _tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	var out []interface{}
	err := _IDAOracle.contract.Call(opts, &out, "verifyAttestation", _tupleRootNonce, _tuple, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_IDAOracle *IDAOracleSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _IDAOracle.Contract.VerifyAttestation(&_IDAOracle.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_IDAOracle *IDAOracleCallerSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _IDAOracle.Contract.VerifyAttestation(&_IDAOracle.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// NamespaceMerkleTreeMetaData contains all meta data concerning the NamespaceMerkleTree contract.
var NamespaceMerkleTreeMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122004c65e75b2c0b05bf08acd3738ddbb9ef874634d0b4adfadfd6f697dd36f19c764736f6c63430008090033",
}

// NamespaceMerkleTreeABI is the input ABI used to generate the binding from.
// Deprecated: Use NamespaceMerkleTreeMetaData.ABI instead.
var NamespaceMerkleTreeABI = NamespaceMerkleTreeMetaData.ABI

// NamespaceMerkleTreeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NamespaceMerkleTreeMetaData.Bin instead.
var NamespaceMerkleTreeBin = NamespaceMerkleTreeMetaData.Bin

// DeployNamespaceMerkleTree deploys a new Ethereum contract, binding an instance of NamespaceMerkleTree to it.
func DeployNamespaceMerkleTree(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NamespaceMerkleTree, error) {
	parsed, err := NamespaceMerkleTreeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NamespaceMerkleTreeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NamespaceMerkleTree{NamespaceMerkleTreeCaller: NamespaceMerkleTreeCaller{contract: contract}, NamespaceMerkleTreeTransactor: NamespaceMerkleTreeTransactor{contract: contract}, NamespaceMerkleTreeFilterer: NamespaceMerkleTreeFilterer{contract: contract}}, nil
}

// NamespaceMerkleTree is an auto generated Go binding around an Ethereum contract.
type NamespaceMerkleTree struct {
	NamespaceMerkleTreeCaller     // Read-only binding to the contract
	NamespaceMerkleTreeTransactor // Write-only binding to the contract
	NamespaceMerkleTreeFilterer   // Log filterer for contract events
}

// NamespaceMerkleTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NamespaceMerkleTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NamespaceMerkleTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NamespaceMerkleTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NamespaceMerkleTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NamespaceMerkleTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NamespaceMerkleTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NamespaceMerkleTreeSession struct {
	Contract     *NamespaceMerkleTree // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NamespaceMerkleTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NamespaceMerkleTreeCallerSession struct {
	Contract *NamespaceMerkleTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// NamespaceMerkleTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NamespaceMerkleTreeTransactorSession struct {
	Contract     *NamespaceMerkleTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// NamespaceMerkleTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NamespaceMerkleTreeRaw struct {
	Contract *NamespaceMerkleTree // Generic contract binding to access the raw methods on
}

// NamespaceMerkleTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NamespaceMerkleTreeCallerRaw struct {
	Contract *NamespaceMerkleTreeCaller // Generic read-only contract binding to access the raw methods on
}

// NamespaceMerkleTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NamespaceMerkleTreeTransactorRaw struct {
	Contract *NamespaceMerkleTreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNamespaceMerkleTree creates a new instance of NamespaceMerkleTree, bound to a specific deployed contract.
func NewNamespaceMerkleTree(address common.Address, backend bind.ContractBackend) (*NamespaceMerkleTree, error) {
	contract, err := bindNamespaceMerkleTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NamespaceMerkleTree{NamespaceMerkleTreeCaller: NamespaceMerkleTreeCaller{contract: contract}, NamespaceMerkleTreeTransactor: NamespaceMerkleTreeTransactor{contract: contract}, NamespaceMerkleTreeFilterer: NamespaceMerkleTreeFilterer{contract: contract}}, nil
}

// NewNamespaceMerkleTreeCaller creates a new read-only instance of NamespaceMerkleTree, bound to a specific deployed contract.
func NewNamespaceMerkleTreeCaller(address common.Address, caller bind.ContractCaller) (*NamespaceMerkleTreeCaller, error) {
	contract, err := bindNamespaceMerkleTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NamespaceMerkleTreeCaller{contract: contract}, nil
}

// NewNamespaceMerkleTreeTransactor creates a new write-only instance of NamespaceMerkleTree, bound to a specific deployed contract.
func NewNamespaceMerkleTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*NamespaceMerkleTreeTransactor, error) {
	contract, err := bindNamespaceMerkleTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NamespaceMerkleTreeTransactor{contract: contract}, nil
}

// NewNamespaceMerkleTreeFilterer creates a new log filterer instance of NamespaceMerkleTree, bound to a specific deployed contract.
func NewNamespaceMerkleTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*NamespaceMerkleTreeFilterer, error) {
	contract, err := bindNamespaceMerkleTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NamespaceMerkleTreeFilterer{contract: contract}, nil
}

// bindNamespaceMerkleTree binds a generic wrapper to an already deployed contract.
func bindNamespaceMerkleTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NamespaceMerkleTreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NamespaceMerkleTree *NamespaceMerkleTreeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NamespaceMerkleTree.Contract.NamespaceMerkleTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NamespaceMerkleTree *NamespaceMerkleTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NamespaceMerkleTree.Contract.NamespaceMerkleTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NamespaceMerkleTree *NamespaceMerkleTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NamespaceMerkleTree.Contract.NamespaceMerkleTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NamespaceMerkleTree *NamespaceMerkleTreeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NamespaceMerkleTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NamespaceMerkleTree *NamespaceMerkleTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NamespaceMerkleTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NamespaceMerkleTree *NamespaceMerkleTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NamespaceMerkleTree.Contract.contract.Transact(opts, method, params...)
}

// QuantumGravityBridgeMetaData contains all meta data concerning the QuantumGravityBridge contract.
var QuantumGravityBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_validatorSetHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataRootTupleRootNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidatorSetNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MalformedCurrentValidatorSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SuppliedValidatorSetInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRootTupleRoot\",\"type\":\"bytes32\"}],\"name\":\"DataRootTupleRootEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"powerThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"validatorSetHash\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_dataRootTupleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_eventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValidatorSetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorSetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_dataRootTupleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"submitDataRootTupleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_oldNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newPowerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_newValidatorSetHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"_tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"817f985b": "state_dataRootTupleRoots(uint256)",
		"cdade866": "state_eventNonce()",
		"5433218c": "state_lastValidatorSetCheckpoint()",
		"e5a2b5d2": "state_powerThreshold()",
		"e23eb326": "submitDataRootTupleRoot(uint256,uint256,bytes32,(address,uint256)[],(uint8,bytes32,bytes32)[])",
		"05d85c13": "updateValidatorSet(uint256,uint256,uint256,bytes32,(address,uint256)[],(uint8,bytes32,bytes32)[])",
		"1f3302a9": "verifyAttestation(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516113c73803806113c783398101604081905261002f916100c9565b604080516918da1958dadc1bda5b9d60b21b6020808301919091528183018690526060820185905260808083018590528351808403909101815260a08301808552815191909201206002879055600081905560018690559085905260c08201849052915185917fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c919081900360e00190a2505050506100f7565b6000806000606084860312156100de57600080fd5b8351925060208401519150604084015190509250925092565b6112c1806101066000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063817f985b1161005b578063817f985b146100d6578063cdade866146100f6578063e23eb326146100ff578063e5a2b5d21461011257600080fd5b806305d85c13146100825780631f3302a9146100975780635433218c146100bf575b600080fd5b610095610090366004610d98565b61011b565b005b6100aa6100a5366004610ec1565b61021c565b60405190151581526020015b60405180910390f35b6100c860005481565b6040519081526020016100b6565b6100c86100e4366004610fe3565b60036020526000908152604090205481565b6100c860025481565b61009561010d366004610ffc565b61027e565b6100c860015481565b600254600180549061012e90839061109f565b8a1461014d576040516368a35ffd60e11b815260040160405180910390fd5b84831461016d5760405163c6617b7b60e01b815260040160405180910390fd5b600061017987876103bd565b90506000546101898b84846103f1565b146101a757604051630bbdaec960e11b815260040160405180910390fd5b60006101b48c8b8b6103f1565b90506101c488888888858861043e565b600081905560018a905560028c9055604080518b8152602081018b90528d917fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c910160405180910390a2505050505050505050505050565b600060025484111561023057506000610277565b6000848152600360209081526040808320548151875181850152928701518383015281518084038301815260609093019091529190610272908390869061054c565b925050505b9392505050565b600254600180549061029190839061109f565b89146102b05760405163e869766d60e01b815260040160405180910390fd5b8483146102d05760405163c6617b7b60e01b815260040160405180910390fd5b60006102dc87876103bd565b90506000546102ec8a84846103f1565b1461030a57604051630bbdaec960e11b815260040160405180910390fd5b604080516f0e8e4c2dce6c2c6e8d2dedc84c2e8c6d60831b6020808301919091528183018d905260608083018c9052835180840390910181526080909201909252805191012061035e88888888858861043e565b60028b905560008b815260036020526040908190208a9055518b907f6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f906103a8908c815260200190565b60405180910390a25050505050505050505050565b600082826040516020016103d29291906110d3565b6040516020818303038152906040528051906020012090505b92915050565b604080516918da1958dadc1bda5b9d60b21b6020808301919091528183019590955260608101939093526080808401929092528051808403909201825260a0909201909152805191012090565b6000805b868110156105215761046a86868381811061045f5761045f61112a565b9050606002016107ba565b156104745761050f565b6104bd8888838181106104895761048961112a565b61049f9260206040909202019081019150611140565b858888858181106104b2576104b261112a565b9050606002016107ee565b6104da57604051638baa579f60e01b815260040160405180910390fd5b8787828181106104ec576104ec61112a565b9050604002016020013582610501919061109f565b915082821061050f57610521565b806105198161115b565b915050610442565b50818110156105435760405163cabeb65560e01b815260040160405180910390fd5b50505050505050565b6000600183604001511161056f578251511561056a57506000610277565b610591565b61058183602001518460400151610888565b8351511461059157506000610277565b82604001518360200151106105a857506000610277565b60006105b38361090f565b8451519091506105dd578360400151600114156105d35784149050610277565b6000915050610277565b60208401516001905b60208601516000906001841b906105fe908290611176565b6106089190611198565b90506000600161061a81861b8461109f565b61062491906111b7565b9050876040015181106106385750506106fe565b9150816106466001856111b7565b8851511161065c57600095505050505050610277565b6106676001856111b7565b6001901b82896020015161067b91906111b7565b10156106b85787516106b19086906106946001886111b7565b815181106106a4576106a461112a565b6020026020010151610984565b94506106ea565b87516106e7906106c96001876111b7565b815181106106d9576106d961112a565b602002602001015186610984565b94505b6106f560018561109f565b935050506105e6565b6001866040015161070f91906111b7565b8114610757576107206001836111b7565b865151116107345760009350505050610277565b85516107479084906106946001866111b7565b925061075460018361109f565b91505b8551516107656001846111b7565b10156107ae57855161079a9061077c6001856111b7565b8151811061078c5761078c61112a565b602002602001015184610984565b92506107a760018361109f565b9150610757565b50509093149392505050565b600060208201351580156107d057506040820135155b80156103eb57506107e460208301836111ce565b60ff161592915050565b600080610848846040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b905061086a8161085b60208601866111ce565b85602001358660400135610a02565b6001600160a01b0316856001600160a01b0316149150509392505050565b600061089382610a2a565b61089f906101006111b7565b905060006108ae6001836111b7565b6001901b90506001816108c191906111b7565b84116108cd57506103eb565b80600114156108e05760019150506103eb565b6108fc6108ed82866111b7565b6108f783866111b7565b610888565b61090790600161109f565b9150506103eb565b60006002600060f81b8360405160200161092a92919061122c565b60408051601f198184030181529082905261094491611250565b602060405180830381855afa158015610961573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906103eb919061125c565b604051600160f81b6020820152602181018390526041810182905260009060029060610160408051601f19818403018152908290526109c291611250565b602060405180830381855afa1580156109df573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190610277919061125c565b6000806000610a1387878787610a57565b91509150610a2081610b44565b5095945050505050565b60005b81816001901b1015610a4b57610a4460018261109f565b9050610a2d565b6103eb816101006111b7565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610a8e5750600090506003610b3b565b8460ff16601b14158015610aa657508460ff16601c14155b15610ab75750600090506004610b3b565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610b0b573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610b3457600060019250925050610b3b565b9150600090505b94509492505050565b6000816004811115610b5857610b58611275565b1415610b615750565b6001816004811115610b7557610b75611275565b1415610bc85760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064015b60405180910390fd5b6002816004811115610bdc57610bdc611275565b1415610c2a5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610bbf565b6003816004811115610c3e57610c3e611275565b1415610c975760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610bbf565b6004816004811115610cab57610cab611275565b1415610d045760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610bbf565b50565b60008083601f840112610d1957600080fd5b50813567ffffffffffffffff811115610d3157600080fd5b6020830191508360208260061b8501011115610d4c57600080fd5b9250929050565b60008083601f840112610d6557600080fd5b50813567ffffffffffffffff811115610d7d57600080fd5b602083019150836020606083028501011115610d4c57600080fd5b60008060008060008060008060c0898b031215610db457600080fd5b88359750602089013596506040890135955060608901359450608089013567ffffffffffffffff80821115610de857600080fd5b610df48c838d01610d07565b909650945060a08b0135915080821115610e0d57600080fd5b50610e1a8b828c01610d53565b999c989b5096995094979396929594505050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610e6757610e67610e2e565b60405290565b6040516060810167ffffffffffffffff81118282101715610e6757610e67610e2e565b604051601f8201601f1916810167ffffffffffffffff81118282101715610eb957610eb9610e2e565b604052919050565b60008060008385036080811215610ed757600080fd5b8435935060206040601f1983011215610eef57600080fd5b610ef7610e44565b86820135815260408701358282015293506060860135915067ffffffffffffffff80831115610f2557600080fd5b9186019160608389031215610f3957600080fd5b610f41610e6d565b833582811115610f5057600080fd5b8401601f81018a13610f6157600080fd5b803583811115610f7357610f73610e2e565b8060051b9350610f84858501610e90565b818152938201850193858101908c861115610f9e57600080fd5b928601925b85841015610fbc57833582529286019290860190610fa3565b80855250505050828401358382015260408401356040820152809450505050509250925092565b600060208284031215610ff557600080fd5b5035919050565b600080600080600080600060a0888a03121561101757600080fd5b873596506020880135955060408801359450606088013567ffffffffffffffff8082111561104457600080fd5b6110508b838c01610d07565b909650945060808a013591508082111561106957600080fd5b506110768a828b01610d53565b989b979a50959850939692959293505050565b634e487b7160e01b600052601160045260246000fd5b600082198211156110b2576110b2611089565b500190565b80356001600160a01b03811681146110ce57600080fd5b919050565b6020808252818101839052600090604080840186845b8781101561111d576001600160a01b03611102836110b7565b168352818501358584015291830191908301906001016110e9565b5090979650505050505050565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561115257600080fd5b610277826110b7565b600060001982141561116f5761116f611089565b5060010190565b60008261119357634e487b7160e01b600052601260045260246000fd5b500490565b60008160001904831182151516156111b2576111b2611089565b500290565b6000828210156111c9576111c9611089565b500390565b6000602082840312156111e057600080fd5b813560ff8116811461027757600080fd5b6000815160005b8181101561121257602081850181015186830152016111f8565b81811115611221576000828601525b509290920192915050565b6001600160f81b031983168152600061124860018301846111f1565b949350505050565b600061027782846111f1565b60006020828403121561126e57600080fd5b5051919050565b634e487b7160e01b600052602160045260246000fdfea264697066735822122056447b3dc649b76e9fd0ffea72fda78eaab4ee470f6a3364c14b3d24c5e0499964736f6c63430008090033",
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
func DeployQuantumGravityBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _nonce *big.Int, _powerThreshold *big.Int, _validatorSetHash [32]byte) (common.Address, *types.Transaction, *QuantumGravityBridge, error) {
	parsed, err := QuantumGravityBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QuantumGravityBridgeBin), backend, _nonce, _powerThreshold, _validatorSetHash)
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

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) StateEventNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "state_eventNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) StateEventNonce() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StateEventNonce(&_QuantumGravityBridge.CallOpts)
}

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) StateEventNonce() (*big.Int, error) {
	return _QuantumGravityBridge.Contract.StateEventNonce(&_QuantumGravityBridge.CallOpts)
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
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeCaller) VerifyAttestation(opts *bind.CallOpts, _tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	var out []interface{}
	err := _QuantumGravityBridge.contract.Call(opts, &out, "verifyAttestation", _tupleRootNonce, _tuple, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _QuantumGravityBridge.Contract.VerifyAttestation(&_QuantumGravityBridge.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_QuantumGravityBridge *QuantumGravityBridgeCallerSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _QuantumGravityBridge.Contract.VerifyAttestation(&_QuantumGravityBridge.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactor) SubmitDataRootTupleRoot(opts *bind.TransactOpts, _newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.contract.Transact(opts, "submitDataRootTupleRoot", _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeSession) SubmitDataRootTupleRoot(_newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.SubmitDataRootTupleRoot(&_QuantumGravityBridge.TransactOpts, _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorSession) SubmitDataRootTupleRoot(_newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.SubmitDataRootTupleRoot(&_QuantumGravityBridge.TransactOpts, _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactor) UpdateValidatorSet(opts *bind.TransactOpts, _newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.contract.Transact(opts, "updateValidatorSet", _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeSession) UpdateValidatorSet(_newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.UpdateValidatorSet(&_QuantumGravityBridge.TransactOpts, _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_QuantumGravityBridge *QuantumGravityBridgeTransactorSession) UpdateValidatorSet(_newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _QuantumGravityBridge.Contract.UpdateValidatorSet(&_QuantumGravityBridge.TransactOpts, _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
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
