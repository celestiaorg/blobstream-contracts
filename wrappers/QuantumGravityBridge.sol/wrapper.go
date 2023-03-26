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

// QgbMetaData contains all meta data concerning the Qgb contract.
var QgbMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_validatorSetHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataRootTupleRootNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidatorSetNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MalformedCurrentValidatorSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SuppliedValidatorSetInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRootTupleRoot\",\"type\":\"bytes32\"}],\"name\":\"DataRootTupleRootEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"powerThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"validatorSetHash\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_dataRootTupleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_eventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValidatorSetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorSetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_dataRootTupleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"submitDataRootTupleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_oldNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newPowerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_newValidatorSetHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"_tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604090808252346100b457606081611230803803809161002182856100ee565b8339810103126100b4577fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c82825192816020820151910151825160208101906918da1958dadc1bda5b9d60b21b825286858201528360608201528260808201526080815261008e816100bb565b519020856002556000558160015582519182526020820152a25161111e90816101128239f35b5050600080fd5b60a081019081106001600160401b038211176100d657604052565b5050634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b038211908210176100d65760405256fe6040608081526004361015610015575b50600080fd5b600090813560e01c806305d85c131461016b5780631f3302a91461013d5780635433218c1461011a578063817f985b146100e3578063cdade866146100bd578063e23eb3261461009c5763e5a2b5d21461006f575061000f565b346100985761009491506100823661042a565b60015490519081529081906020820190565b0390f35b5080fd5b5034610098576100ba6100ae3661044a565b959490949391936108e9565b51f35b50346100985761009491506100d13661042a565b60025490519081529081906020820190565b503461009857610094915061010b6100fa36610438565b600052600360205260406000205490565b90519081529081906020820190565b5034610098576100949161012d3661042a565b5490519081529081906020820190565b503461009857610094915061015a61015436610316565b91610a12565b905190151581529081906020820190565b5034610098576100ba61017d366101ff565b969590959491949392936104fa565b9181601f840112156101bd5782359167ffffffffffffffff83116101c5576020808501948460061b0101116101bd57565b505050600080fd5b50505050600080fd5b9181601f840112156101bd5782359167ffffffffffffffff83116101c557602080850194606085020101116101bd57565b60c0600319820112610275576004359160243591604435916064359167ffffffffffffffff91608435838111610268578261023c9160040161018c565b9390939260a43591821161025a57610256916004016101ce565b9091565b505050505050505050600080fd5b5050505050505050600080fd5b5050600080fd5b50634e487b7160e01b600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176102af57604052565b6102b761027c565b604052565b6060810190811067ffffffffffffffff8211176102af57604052565b6080810190811067ffffffffffffffff8211176102af57604052565b90601f8019910116810190811067ffffffffffffffff8211176102af57604052565b90600319906080828401126101bd576004359260406023198201126101c5576040519261034284610293565b602435845260209360443585820152936064359167ffffffffffffffff9081841161026857606090848603011261041e5760405193610380856102bc565b836004013582811161025a5784018160238201121561025a576004810135928311610411575b8260051b90604051936103bb868401866102f4565b845260248585019282010192831161040257602401905b8282106103f357505050835260248201359083015260440135604082015290565b813581529084019084016103d2565b50505050505050505050600080fd5b61041961027c565b6103a6565b50505050505050600080fd5b600090600319011261000f57565b602090600319011261000f5760043590565b9060a06003198301126102755760043591602435916044359167ffffffffffffffff9160643583811161041e57826104849160040161018c565b9390939260843591821161026857610256916004016101ce565b50634e487b7160e01b600052601160045260246000fd5b60019060011981116104c5570190565b6104cd61049e565b0190565b80196001116104e1575b60010190565b6104e961049e565b6104db565b811981116104c5570190565b9695919394929460025491610511600154936104b5565b8914156105db57878514156105c057610535908361052f8785610608565b91610682565b60005414156105a6577fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c966105779361056f88888c610682565b95869361074c565b60005561058381600155565b61058c84600255565b60408051918252602082019290925290819081015b0390a2565b5050505050505050506004604051630bbdaec960e11b8152fd5b50505050505050505050600460405163c6617b7b60e01b8152fd5b5050505050505050505060046040516368a35ffd60e11b8152fd5b6001600160a01b038116141561000f57565b604091825180926020928383019581818501868952526060840192946000905b83821061064c5750505050610646925003601f1981018352826102f4565b51902090565b916001919395508080958835610661816105f6565b848060a01b0316815284890135858201520196019201869492959391610628565b916040519160208301936918da1958dadc1bda5b9d60b21b85526040840152606083015260808201526080815260a0810181811067ffffffffffffffff8211176106d0575b60405251902090565b6106d861027c565b6106c7565b60019060001981146104c5570190565b50634e487b7160e01b600052603260045260246000fd5b9160609181101561071457020190565b61071c6106ed565b020190565b9190811015610732575b60061b0190565b61073a6106ed565b61072b565b35610749816105f6565b90565b600094859493905b83861061077c575b5050505050501061076957565b5060405163cabeb65560e01b8152600490fd5b909192939495610795610790888886610704565b610839565b610820576107c86107c46107b26107ad8a8989610721565b61073f565b846107be8b8b89610704565b9161086e565b1590565b610806576107e59060206107dd898888610721565b0135906104ee565b9587871015610801576107f7906106dd565b9493929190610754565b61075c565b5050505050505050506004604051638baa579f60e01b8152fd5b956107f7906106dd565b3560ff81168114156102755790565b6020810135159081610861575b8161084f575090565b60ff915061085c9061082a565b161590565b6040810135159150610846565b916108d0906108d89260405160208101917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008352603c820152603c81526108b4816102bc565b5190206108c08261082a565b6020604084013593013591610c14565b919091610a7c565b6001600160a01b0391821691161490565b9594939192600254916108fe600154936104b5565b8814156109c157868414156109a75761091c908361052f8685610608565b600054141561098e577f6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f956105a19561095f94610959878b6109db565b9361074c565b61096884600255565b8061097d856000526003602052604060002090565b556040519081529081906020820190565b50505050505050506004604051630bbdaec960e11b8152fd5b505050505050505050600460405163c6617b7b60e01b8152fd5b505050505050505050600460405163e869766d60e01b8152fd5b906040519060208201926f0e8e4c2dce6c2c6e8d2dedc84c2e8c6d60831b84526040830152606082015260608152610646816102d8565b916002548311610a53576107499260005260036020526040600020546020604051938051828601520151604084015260408352610a4e836102bc565b610d70565b505050600090565b60051115610a6557565b50634e487b7160e01b600052602160045260246000fd5b610a8581610a5b565b80610a8d5750565b610a9681610a5b565b6001811415610ae657505060405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606490fd5b610aef81610a5b565b6002811415610b3f57505060405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606490fd5b610b4881610a5b565b6003811415610ba357505060405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608490fd5b80610baf600492610a5b565b14610bb657565b5060405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608490fd5b506040513d6000823e3d90fd5b9291907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311610cbe5760ff16601b81141580610cb3575b610ca7579160809493916020936040519384528484015260408301526060820152600093849182805260015afa15610c9a575b81516001600160a01b03811615610c94579190565b50600190565b610ca2610c07565b610c7f565b50505050600090600490565b50601c811415610c4c565b50505050600090600390565b8115610cd4570490565b505050634e487b7160e01b600052601260045260246000fd5b8060001904821181151516610d00570290565b610d0861049e565b0290565b60018110610d1c575b6000190190565b610d2461049e565b610d15565b610100818110610d37570390565b610d3f61049e565b0390565b818110610d37570390565b6020918151811015610d63575b60051b010190565b610d6b6106ed565b610d5b565b6040820180519193909291600190818111610f485750825151610e61575b6020830191825185511115610f3d57610da6906110a9565b9383515115610f2457819081938051925b610e6b575b50610dc79051610d0c565b1415610e2b575b809291925b610ddf575b5050501490565b909192610deb84610d0c565b83519081511115610e2457610e1c91610e10610e1692610e0a88610d0c565b90610d4e565b51611051565b936104b5565b919080610dd3565b5092610dd8565b9092825151610e3985610d0c565b1015610e6157610e16610e5b91610e548551610e0a88610d0c565b5190611051565b90610dce565b5050505050600090565b610e8a92610e9b610e96610e8f849a97989a51898c1b97888092610cca565b610ced565b95866104ee565b610d0c565b908351821015610f18575092865151610eb389610d0c565b1015610f0a578594610ec9610ef2928451610d43565b86610ed38b610d0c565b1b1115610ef857610eec90610e548951610e0a8c610d0c565b976104b5565b94610db7565b610eec90610e108951610e0a8c610d0c565b505050505050505050600090565b93505095939295610dbc565b5193949314159150610f369050571490565b5050600090565b505050505050600090565b610f59845151916020860151610f64565b1415610e6157610d8e565b9160005b6001908382821b1015610f92576001198111610f85575b01610f68565b610f8d61049e565b610f7f565b610fab9150939193610100818110611009575b03610d29565b916001610fb784610d0c565b1b91610fc283610d0c565b8111610fce5750505090565b919250906001831415610fe357505050600190565b82610ff461100094610ffa93610d43565b92610d43565b90610f64565b610749906104d1565b61101161049e565b610fa5565b90815180926000905b82821061103a575011611030570190565b6000828201520190565b91508060208092840101518185015201839161101f565b61108a60009160209360405191600160f81b86840152602183015260418201526041815261107e816102d8565b60405191828092611016565b039060025afa1561109c575b60005190565b6110a4610c07565b611096565b600061108a6110cc61107e60209460405192839186888401526021830190611016565b03601f1981018352826102f456fea36469706673582212200d875517f47970d8a4b208740e70157ae6426146a1418f828df084fc40cd63206c6578706572696d656e74616cf564736f6c634300080b0041",
}

// QgbABI is the input ABI used to generate the binding from.
// Deprecated: Use QgbMetaData.ABI instead.
var QgbABI = QgbMetaData.ABI

// QgbBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QgbMetaData.Bin instead.
var QgbBin = QgbMetaData.Bin

// DeployQgb deploys a new Ethereum contract, binding an instance of Qgb to it.
func DeployQgb(auth *bind.TransactOpts, backend bind.ContractBackend, _nonce *big.Int, _powerThreshold *big.Int, _validatorSetHash [32]byte) (common.Address, *types.Transaction, *Qgb, error) {
	parsed, err := QgbMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QgbBin), backend, _nonce, _powerThreshold, _validatorSetHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Qgb{QgbCaller: QgbCaller{contract: contract}, QgbTransactor: QgbTransactor{contract: contract}, QgbFilterer: QgbFilterer{contract: contract}}, nil
}

// Qgb is an auto generated Go binding around an Ethereum contract.
type Qgb struct {
	QgbCaller     // Read-only binding to the contract
	QgbTransactor // Write-only binding to the contract
	QgbFilterer   // Log filterer for contract events
}

// QgbCaller is an auto generated read-only Go binding around an Ethereum contract.
type QgbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QgbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QgbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QgbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QgbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QgbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QgbSession struct {
	Contract     *Qgb              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QgbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QgbCallerSession struct {
	Contract *QgbCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QgbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QgbTransactorSession struct {
	Contract     *QgbTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QgbRaw is an auto generated low-level Go binding around an Ethereum contract.
type QgbRaw struct {
	Contract *Qgb // Generic contract binding to access the raw methods on
}

// QgbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QgbCallerRaw struct {
	Contract *QgbCaller // Generic read-only contract binding to access the raw methods on
}

// QgbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QgbTransactorRaw struct {
	Contract *QgbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQgb creates a new instance of Qgb, bound to a specific deployed contract.
func NewQgb(address common.Address, backend bind.ContractBackend) (*Qgb, error) {
	contract, err := bindQgb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Qgb{QgbCaller: QgbCaller{contract: contract}, QgbTransactor: QgbTransactor{contract: contract}, QgbFilterer: QgbFilterer{contract: contract}}, nil
}

// NewQgbCaller creates a new read-only instance of Qgb, bound to a specific deployed contract.
func NewQgbCaller(address common.Address, caller bind.ContractCaller) (*QgbCaller, error) {
	contract, err := bindQgb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QgbCaller{contract: contract}, nil
}

// NewQgbTransactor creates a new write-only instance of Qgb, bound to a specific deployed contract.
func NewQgbTransactor(address common.Address, transactor bind.ContractTransactor) (*QgbTransactor, error) {
	contract, err := bindQgb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QgbTransactor{contract: contract}, nil
}

// NewQgbFilterer creates a new log filterer instance of Qgb, bound to a specific deployed contract.
func NewQgbFilterer(address common.Address, filterer bind.ContractFilterer) (*QgbFilterer, error) {
	contract, err := bindQgb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QgbFilterer{contract: contract}, nil
}

// bindQgb binds a generic wrapper to an already deployed contract.
func bindQgb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QgbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Qgb *QgbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Qgb.Contract.QgbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Qgb *QgbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Qgb.Contract.QgbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Qgb *QgbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Qgb.Contract.QgbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Qgb *QgbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Qgb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Qgb *QgbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Qgb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Qgb *QgbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Qgb.Contract.contract.Transact(opts, method, params...)
}

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_Qgb *QgbCaller) StateDataRootTupleRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Qgb.contract.Call(opts, &out, "state_dataRootTupleRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_Qgb *QgbSession) StateDataRootTupleRoots(arg0 *big.Int) ([32]byte, error) {
	return _Qgb.Contract.StateDataRootTupleRoots(&_Qgb.CallOpts, arg0)
}

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_Qgb *QgbCallerSession) StateDataRootTupleRoots(arg0 *big.Int) ([32]byte, error) {
	return _Qgb.Contract.StateDataRootTupleRoots(&_Qgb.CallOpts, arg0)
}

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_Qgb *QgbCaller) StateEventNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Qgb.contract.Call(opts, &out, "state_eventNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_Qgb *QgbSession) StateEventNonce() (*big.Int, error) {
	return _Qgb.Contract.StateEventNonce(&_Qgb.CallOpts)
}

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_Qgb *QgbCallerSession) StateEventNonce() (*big.Int, error) {
	return _Qgb.Contract.StateEventNonce(&_Qgb.CallOpts)
}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_Qgb *QgbCaller) StateLastValidatorSetCheckpoint(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Qgb.contract.Call(opts, &out, "state_lastValidatorSetCheckpoint")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_Qgb *QgbSession) StateLastValidatorSetCheckpoint() ([32]byte, error) {
	return _Qgb.Contract.StateLastValidatorSetCheckpoint(&_Qgb.CallOpts)
}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_Qgb *QgbCallerSession) StateLastValidatorSetCheckpoint() ([32]byte, error) {
	return _Qgb.Contract.StateLastValidatorSetCheckpoint(&_Qgb.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Qgb *QgbCaller) StatePowerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Qgb.contract.Call(opts, &out, "state_powerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Qgb *QgbSession) StatePowerThreshold() (*big.Int, error) {
	return _Qgb.Contract.StatePowerThreshold(&_Qgb.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Qgb *QgbCallerSession) StatePowerThreshold() (*big.Int, error) {
	return _Qgb.Contract.StatePowerThreshold(&_Qgb.CallOpts)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_Qgb *QgbCaller) VerifyAttestation(opts *bind.CallOpts, _tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	var out []interface{}
	err := _Qgb.contract.Call(opts, &out, "verifyAttestation", _tupleRootNonce, _tuple, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_Qgb *QgbSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _Qgb.Contract.VerifyAttestation(&_Qgb.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_Qgb *QgbCallerSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _Qgb.Contract.VerifyAttestation(&_Qgb.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Qgb *QgbTransactor) SubmitDataRootTupleRoot(opts *bind.TransactOpts, _newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Qgb.contract.Transact(opts, "submitDataRootTupleRoot", _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Qgb *QgbSession) SubmitDataRootTupleRoot(_newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Qgb.Contract.SubmitDataRootTupleRoot(&_Qgb.TransactOpts, _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Qgb *QgbTransactorSession) SubmitDataRootTupleRoot(_newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Qgb.Contract.SubmitDataRootTupleRoot(&_Qgb.TransactOpts, _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Qgb *QgbTransactor) UpdateValidatorSet(opts *bind.TransactOpts, _newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Qgb.contract.Transact(opts, "updateValidatorSet", _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Qgb *QgbSession) UpdateValidatorSet(_newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Qgb.Contract.UpdateValidatorSet(&_Qgb.TransactOpts, _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Qgb *QgbTransactorSession) UpdateValidatorSet(_newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Qgb.Contract.UpdateValidatorSet(&_Qgb.TransactOpts, _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// QgbDataRootTupleRootEventIterator is returned from FilterDataRootTupleRootEvent and is used to iterate over the raw logs and unpacked data for DataRootTupleRootEvent events raised by the Qgb contract.
type QgbDataRootTupleRootEventIterator struct {
	Event *QgbDataRootTupleRootEvent // Event containing the contract specifics and raw log

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
func (it *QgbDataRootTupleRootEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QgbDataRootTupleRootEvent)
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
		it.Event = new(QgbDataRootTupleRootEvent)
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
func (it *QgbDataRootTupleRootEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QgbDataRootTupleRootEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QgbDataRootTupleRootEvent represents a DataRootTupleRootEvent event raised by the Qgb contract.
type QgbDataRootTupleRootEvent struct {
	Nonce             *big.Int
	DataRootTupleRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDataRootTupleRootEvent is a free log retrieval operation binding the contract event 0x6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f.
//
// Solidity: event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot)
func (_Qgb *QgbFilterer) FilterDataRootTupleRootEvent(opts *bind.FilterOpts, nonce []*big.Int) (*QgbDataRootTupleRootEventIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Qgb.contract.FilterLogs(opts, "DataRootTupleRootEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return &QgbDataRootTupleRootEventIterator{contract: _Qgb.contract, event: "DataRootTupleRootEvent", logs: logs, sub: sub}, nil
}

// WatchDataRootTupleRootEvent is a free log subscription operation binding the contract event 0x6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f.
//
// Solidity: event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot)
func (_Qgb *QgbFilterer) WatchDataRootTupleRootEvent(opts *bind.WatchOpts, sink chan<- *QgbDataRootTupleRootEvent, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Qgb.contract.WatchLogs(opts, "DataRootTupleRootEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QgbDataRootTupleRootEvent)
				if err := _Qgb.contract.UnpackLog(event, "DataRootTupleRootEvent", log); err != nil {
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
func (_Qgb *QgbFilterer) ParseDataRootTupleRootEvent(log types.Log) (*QgbDataRootTupleRootEvent, error) {
	event := new(QgbDataRootTupleRootEvent)
	if err := _Qgb.contract.UnpackLog(event, "DataRootTupleRootEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QgbValidatorSetUpdatedEventIterator is returned from FilterValidatorSetUpdatedEvent and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdatedEvent events raised by the Qgb contract.
type QgbValidatorSetUpdatedEventIterator struct {
	Event *QgbValidatorSetUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *QgbValidatorSetUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QgbValidatorSetUpdatedEvent)
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
		it.Event = new(QgbValidatorSetUpdatedEvent)
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
func (it *QgbValidatorSetUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QgbValidatorSetUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QgbValidatorSetUpdatedEvent represents a ValidatorSetUpdatedEvent event raised by the Qgb contract.
type QgbValidatorSetUpdatedEvent struct {
	Nonce            *big.Int
	PowerThreshold   *big.Int
	ValidatorSetHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdatedEvent is a free log retrieval operation binding the contract event 0xe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c.
//
// Solidity: event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash)
func (_Qgb *QgbFilterer) FilterValidatorSetUpdatedEvent(opts *bind.FilterOpts, nonce []*big.Int) (*QgbValidatorSetUpdatedEventIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Qgb.contract.FilterLogs(opts, "ValidatorSetUpdatedEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return &QgbValidatorSetUpdatedEventIterator{contract: _Qgb.contract, event: "ValidatorSetUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdatedEvent is a free log subscription operation binding the contract event 0xe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c.
//
// Solidity: event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash)
func (_Qgb *QgbFilterer) WatchValidatorSetUpdatedEvent(opts *bind.WatchOpts, sink chan<- *QgbValidatorSetUpdatedEvent, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Qgb.contract.WatchLogs(opts, "ValidatorSetUpdatedEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QgbValidatorSetUpdatedEvent)
				if err := _Qgb.contract.UnpackLog(event, "ValidatorSetUpdatedEvent", log); err != nil {
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
func (_Qgb *QgbFilterer) ParseValidatorSetUpdatedEvent(log types.Log) (*QgbValidatorSetUpdatedEvent, error) {
	event := new(QgbValidatorSetUpdatedEvent)
	if err := _Qgb.contract.UnpackLog(event, "ValidatorSetUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
