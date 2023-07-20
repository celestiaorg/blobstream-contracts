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

// WrappersMetaData contains all meta data concerning the Wrappers contract.
var WrappersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_validatorSetHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataRootTupleRootNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidatorSetNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MalformedCurrentValidatorSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SuppliedValidatorSetInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRootTupleRoot\",\"type\":\"bytes32\"}],\"name\":\"DataRootTupleRootEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"powerThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"validatorSetHash\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_dataRootTupleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_eventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValidatorSetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorSetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_dataRootTupleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"submitDataRootTupleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_oldNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newPowerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_newValidatorSetHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"_tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080346100ed57601f61103b38819003918201601f1916830192916001600160401b0391828511848610176100d75781606092859260409788528339810103126100ed5781519183602082015191015191845192602084016918da1958dadc1bda5b9d60b21b815285878601528360608601528160808601526080855260a0850192858410908411176100d7577fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c94879460c09285875282519020886002556000558060015584520152a251610f4890816100f38239f35b634e487b7160e01b600052604160045260246000fd5b600080fdfe608060408181526004918236101561001657600080fd5b600092833560e01c91826305d85c131461039f575081631f3302a9146102545781635433218c14610237578163817f985b1461020f578163cdade866146101f0578163e23eb32614610093575063e5a2b5d21461007257600080fd5b3461008f578160031936011261008f576020906001549051908152f35b5080fd5b9050346101ec5760a03660031901126101ec578035916044359167ffffffffffffffff6064358181116101e8576100cd90369084016104cc565b90916084359081116101e4576100e69036908501610502565b916002549460015495600181018091116101d15789036101c3578382036101b55761011d61011483876105d4565b87602435610654565b8a54036101a757509261018f927f6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f979592879560209851898101906f0e8e4c2dce6c2c6e8d2dedc84c2e8c6d60831b82528c898201528960608201526060815261018681610533565b519020936106d7565b8460025584865260038352818187205551908152a280f35b8651630bbdaec960e11b8152fd5b865163c6617b7b60e01b8152fd5b865163e869766d60e01b8152fd5b634e487b7160e01b8b526011825260248bfd5b8780fd5b8680fd5b8280fd5b50503461008f578160031936011261008f576020906002549051908152f35b9050346101ec5760203660031901126101ec5760209282913581526003845220549051908152f35b50503461008f578160031936011261008f57602091549051908152f35b83833461008f576003199160803684011261035d578160231936011261035d57815167ffffffffffffffff948184018681118382101761038c57845260248035835260209660443588850152606435968188116103885760609088360301126103845785519460608601918683108184111761037257828852888501358181116101ec578901366023820112156101ec5785810135918211610360578160051b916103018c840186610581565b84528460808901928201019236841161035d575090848b9201905b83821061034e575050505095604491610345969786528101358886015201358584015235610a81565b90519015158152f35b8135815290820190820161031c565b80fd5b634e487b7160e01b8352604186528483fd5b634e487b7160e01b8252604185528382fd5b8480fd5b8580fd5b634e487b7160e01b845260418252602484fd5b848285346101ec5760c03660031901126101ec57813591604435906064359067ffffffffffffffff6084358181116101e4576103de90369084016104cc565b92909160a4359081116104c8576103f89036908301610502565b916002549a6001549b600181018091116104b5578a036104a9575082850361049b5761043061042786866105d4565b8c602435610654565b8a540361048d575091879899916104759361046d87897fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c9c610654565b9586936106d7565b8655816001558460025582519182526020820152a280f35b8751630bbdaec960e11b8152fd5b875163c6617b7b60e01b8152fd5b6368a35ffd60e11b8152fd5b634e487b7160e01b8c526011835260248cfd5b8880fd5b9181601f840112156104fd5782359167ffffffffffffffff83116104fd576020808501948460061b0101116104fd57565b600080fd5b9181601f840112156104fd5782359167ffffffffffffffff83116104fd57602080850194606085020101116104fd57565b6080810190811067ffffffffffffffff82111761054f57604052565b634e487b7160e01b600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761054f57604052565b90601f8019910116810190811067ffffffffffffffff82111761054f57604052565b60010190816001116105b157565b634e487b7160e01b600052601160045260246000fd5b919082018092116105b157565b60409182518092602092838301958181850186895252606084019294600090815b84831061061a575050505050610614925003601f198101835282610581565b51902090565b919395509193863560018060a01b0381168091036101ec578582819260019452858a013586820152019701930190918795939694926105f5565b916040519160208301936918da1958dadc1bda5b9d60b21b85526040840152606083015260808201526080815260a0810181811067ffffffffffffffff82111761054f5760405251902090565b91908110156106b1576060020190565b634e487b7160e01b600052603260045260246000fd5b91908110156106b15760061b0190565b93919060009485935b828510610707575b505050505050106106f557565b60405163cabeb65560e01b8152600490fd5b909192939497969561071a868a876106a1565b6020908181013515908161083f575b8161082a575b506108235761073f8786866106c7565b6001600160a01b039035818116908190036104fd5761075f898d8a6106a1565b916107c66107be6040948551878101907f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252603c8b8183015281526107a481610565565b5190206107b08261084c565b8888840135930135916109cc565b91909161085a565b16036108135750906107e5916107dd8887876106c7565b0135906105c7565b9486861015610809575b60001981146105b1576001019392919097949596976106e0565b85969798506106e8565b51638baa579f60e01b8152600490fd5b50946107ef565b60ff91506108379061084c565b16153861072f565b6040810135159150610729565b3560ff811681036104fd5790565b60058110156109b6578061086b5750565b600181036108b85760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606490fd5b600281036109055760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606490fd5b6003810361095d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608490fd5b60041461096657565b60405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608490fd5b634e487b7160e01b600052602160045260246000fd5b9291907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311610a755760ff16601b81141580610a6a575b610a5e579160809493916020936040519384528484015260408301526060820152600093849182805260015afa15610a515781516001600160a01b03811615610a4b579190565b50600190565b50604051903d90823e3d90fd5b50505050600090600490565b50601c811415610a04565b50505050600090600390565b916002548311610ac557610ac29260005260036020526040600020546020604051938051828601520151604084015260408352610abd83610565565b610b0d565b90565b505050600090565b906101009182039182116105b157565b6000198101919082116105b157565b919082039182116105b157565b80518210156106b15760209160051b010190565b6040820180519193909291600190818111610dd45750825151610dca575b6020830192835185511115610dbf576020610b5e93610b78604051610b6c81600098899586888401526021830190610e9d565b03601f198101835282610581565b60405191828092610e9d565b039060025afa15610db45782519481515115610d9c57829081958051925b610c97575b505160001991828201918211610c835703610c36575b8293925b610bc3575b50505050501490565b9091929394818601868111610c225783518051821015610c1a57610bf19291610beb91610af9565b51610ec8565b94848101809111610c06579392919083610bb5565b634e487b7160e01b84526011600452602484fd5b505094610bba565b634e487b7160e01b85526011600452602485fd5b9091939482515182870190878211610c8357811015610c785790610c5e610c65928551610af9565b5190610ec8565b94848101809111610c0657939190610bb1565b505050509250505090565b634e487b7160e01b86526011600452602486fd5b9091819793949697519387891b948515610d8857891c94858a1b95808704821490151715610d7457610cc990866105c7565b60001992818401918211610d52578551821015610d66575094865151928a01928a8411610d5257831015610d4457610d02908451610aec565b88831b1115610d3057610c5e610d19928751610af9565b965b868101809111610c8357959392919084610b96565b610beb610d3e928751610af9565b96610d1b565b505050505050509250505090565b634e487b7160e01b89526011600452602489fd5b949350509796949350610b9b565b634e487b7160e01b88526011600452602488fd5b634e487b7160e01b88526012600452602488fd5b519495949293505003610dae57501490565b91505090565b6040513d84823e3d90fd5b505050505050600090565b5050505050600090565b610de5845151916020860151610df4565b14610b2b575050505050600090565b909160005b60018481831b1015610e2557810180911115610df957634e487b7160e01b600052601160045260246000fd5b50929190926101009081039081116105b157610e4090610acd565b916001610e4c84610add565b1b91610e5783610add565b8111610e635750505090565b9192509060018303610e7757505050600190565b82610e88610e9494610e8e93610aec565b92610aec565b90610df4565b610ac2906105a3565b9081519160005b838110610eb5575050016000815290565b8060208092840101518185015201610ea4565b610ef560009160209360405191600160f81b868401526021830152604182015260418152610b6c81610533565b039060025afa15610f065760005190565b6040513d6000823e3d90fdfea26469706673582212201ad19945ea2c445e8a7415dfda1c5c81aa1425c5c61e8724c545b0e58e244e1164736f6c63430008140033",
}

// WrappersABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappersMetaData.ABI instead.
var WrappersABI = WrappersMetaData.ABI

// WrappersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WrappersMetaData.Bin instead.
var WrappersBin = WrappersMetaData.Bin

// DeployWrappers deploys a new Ethereum contract, binding an instance of Wrappers to it.
func DeployWrappers(auth *bind.TransactOpts, backend bind.ContractBackend, _nonce *big.Int, _powerThreshold *big.Int, _validatorSetHash [32]byte) (common.Address, *types.Transaction, *Wrappers, error) {
	parsed, err := WrappersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrappersBin), backend, _nonce, _powerThreshold, _validatorSetHash)
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

// WrappersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
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
	parsed, err := abi.JSON(strings.NewReader(WrappersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_Wrappers *WrappersCaller) StateDataRootTupleRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Wrappers.contract.Call(opts, &out, "state_dataRootTupleRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_Wrappers *WrappersSession) StateDataRootTupleRoots(arg0 *big.Int) ([32]byte, error) {
	return _Wrappers.Contract.StateDataRootTupleRoots(&_Wrappers.CallOpts, arg0)
}

// StateDataRootTupleRoots is a free data retrieval call binding the contract method 0x817f985b.
//
// Solidity: function state_dataRootTupleRoots(uint256 ) view returns(bytes32)
func (_Wrappers *WrappersCallerSession) StateDataRootTupleRoots(arg0 *big.Int) ([32]byte, error) {
	return _Wrappers.Contract.StateDataRootTupleRoots(&_Wrappers.CallOpts, arg0)
}

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_Wrappers *WrappersCaller) StateEventNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wrappers.contract.Call(opts, &out, "state_eventNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_Wrappers *WrappersSession) StateEventNonce() (*big.Int, error) {
	return _Wrappers.Contract.StateEventNonce(&_Wrappers.CallOpts)
}

// StateEventNonce is a free data retrieval call binding the contract method 0xcdade866.
//
// Solidity: function state_eventNonce() view returns(uint256)
func (_Wrappers *WrappersCallerSession) StateEventNonce() (*big.Int, error) {
	return _Wrappers.Contract.StateEventNonce(&_Wrappers.CallOpts)
}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_Wrappers *WrappersCaller) StateLastValidatorSetCheckpoint(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Wrappers.contract.Call(opts, &out, "state_lastValidatorSetCheckpoint")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_Wrappers *WrappersSession) StateLastValidatorSetCheckpoint() ([32]byte, error) {
	return _Wrappers.Contract.StateLastValidatorSetCheckpoint(&_Wrappers.CallOpts)
}

// StateLastValidatorSetCheckpoint is a free data retrieval call binding the contract method 0x5433218c.
//
// Solidity: function state_lastValidatorSetCheckpoint() view returns(bytes32)
func (_Wrappers *WrappersCallerSession) StateLastValidatorSetCheckpoint() ([32]byte, error) {
	return _Wrappers.Contract.StateLastValidatorSetCheckpoint(&_Wrappers.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Wrappers *WrappersCaller) StatePowerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wrappers.contract.Call(opts, &out, "state_powerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Wrappers *WrappersSession) StatePowerThreshold() (*big.Int, error) {
	return _Wrappers.Contract.StatePowerThreshold(&_Wrappers.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Wrappers *WrappersCallerSession) StatePowerThreshold() (*big.Int, error) {
	return _Wrappers.Contract.StatePowerThreshold(&_Wrappers.CallOpts)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_Wrappers *WrappersCaller) VerifyAttestation(opts *bind.CallOpts, _tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	var out []interface{}
	err := _Wrappers.contract.Call(opts, &out, "verifyAttestation", _tupleRootNonce, _tuple, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_Wrappers *WrappersSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _Wrappers.Contract.VerifyAttestation(&_Wrappers.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_Wrappers *WrappersCallerSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _Wrappers.Contract.VerifyAttestation(&_Wrappers.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Wrappers *WrappersTransactor) SubmitDataRootTupleRoot(opts *bind.TransactOpts, _newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Wrappers.contract.Transact(opts, "submitDataRootTupleRoot", _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Wrappers *WrappersSession) SubmitDataRootTupleRoot(_newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Wrappers.Contract.SubmitDataRootTupleRoot(&_Wrappers.TransactOpts, _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// SubmitDataRootTupleRoot is a paid mutator transaction binding the contract method 0xe23eb326.
//
// Solidity: function submitDataRootTupleRoot(uint256 _newNonce, uint256 _validatorSetNonce, bytes32 _dataRootTupleRoot, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Wrappers *WrappersTransactorSession) SubmitDataRootTupleRoot(_newNonce *big.Int, _validatorSetNonce *big.Int, _dataRootTupleRoot [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Wrappers.Contract.SubmitDataRootTupleRoot(&_Wrappers.TransactOpts, _newNonce, _validatorSetNonce, _dataRootTupleRoot, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Wrappers *WrappersTransactor) UpdateValidatorSet(opts *bind.TransactOpts, _newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Wrappers.contract.Transact(opts, "updateValidatorSet", _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Wrappers *WrappersSession) UpdateValidatorSet(_newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Wrappers.Contract.UpdateValidatorSet(&_Wrappers.TransactOpts, _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x05d85c13.
//
// Solidity: function updateValidatorSet(uint256 _newNonce, uint256 _oldNonce, uint256 _newPowerThreshold, bytes32 _newValidatorSetHash, (address,uint256)[] _currentValidatorSet, (uint8,bytes32,bytes32)[] _sigs) returns()
func (_Wrappers *WrappersTransactorSession) UpdateValidatorSet(_newNonce *big.Int, _oldNonce *big.Int, _newPowerThreshold *big.Int, _newValidatorSetHash [32]byte, _currentValidatorSet []Validator, _sigs []Signature) (*types.Transaction, error) {
	return _Wrappers.Contract.UpdateValidatorSet(&_Wrappers.TransactOpts, _newNonce, _oldNonce, _newPowerThreshold, _newValidatorSetHash, _currentValidatorSet, _sigs)
}

// WrappersDataRootTupleRootEventIterator is returned from FilterDataRootTupleRootEvent and is used to iterate over the raw logs and unpacked data for DataRootTupleRootEvent events raised by the Wrappers contract.
type WrappersDataRootTupleRootEventIterator struct {
	Event *WrappersDataRootTupleRootEvent // Event containing the contract specifics and raw log

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
func (it *WrappersDataRootTupleRootEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappersDataRootTupleRootEvent)
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
		it.Event = new(WrappersDataRootTupleRootEvent)
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
func (it *WrappersDataRootTupleRootEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappersDataRootTupleRootEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappersDataRootTupleRootEvent represents a DataRootTupleRootEvent event raised by the Wrappers contract.
type WrappersDataRootTupleRootEvent struct {
	Nonce             *big.Int
	DataRootTupleRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDataRootTupleRootEvent is a free log retrieval operation binding the contract event 0x6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f.
//
// Solidity: event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot)
func (_Wrappers *WrappersFilterer) FilterDataRootTupleRootEvent(opts *bind.FilterOpts, nonce []*big.Int) (*WrappersDataRootTupleRootEventIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Wrappers.contract.FilterLogs(opts, "DataRootTupleRootEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return &WrappersDataRootTupleRootEventIterator{contract: _Wrappers.contract, event: "DataRootTupleRootEvent", logs: logs, sub: sub}, nil
}

// WatchDataRootTupleRootEvent is a free log subscription operation binding the contract event 0x6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f.
//
// Solidity: event DataRootTupleRootEvent(uint256 indexed nonce, bytes32 dataRootTupleRoot)
func (_Wrappers *WrappersFilterer) WatchDataRootTupleRootEvent(opts *bind.WatchOpts, sink chan<- *WrappersDataRootTupleRootEvent, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Wrappers.contract.WatchLogs(opts, "DataRootTupleRootEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappersDataRootTupleRootEvent)
				if err := _Wrappers.contract.UnpackLog(event, "DataRootTupleRootEvent", log); err != nil {
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
func (_Wrappers *WrappersFilterer) ParseDataRootTupleRootEvent(log types.Log) (*WrappersDataRootTupleRootEvent, error) {
	event := new(WrappersDataRootTupleRootEvent)
	if err := _Wrappers.contract.UnpackLog(event, "DataRootTupleRootEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappersValidatorSetUpdatedEventIterator is returned from FilterValidatorSetUpdatedEvent and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdatedEvent events raised by the Wrappers contract.
type WrappersValidatorSetUpdatedEventIterator struct {
	Event *WrappersValidatorSetUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *WrappersValidatorSetUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappersValidatorSetUpdatedEvent)
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
		it.Event = new(WrappersValidatorSetUpdatedEvent)
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
func (it *WrappersValidatorSetUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappersValidatorSetUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappersValidatorSetUpdatedEvent represents a ValidatorSetUpdatedEvent event raised by the Wrappers contract.
type WrappersValidatorSetUpdatedEvent struct {
	Nonce            *big.Int
	PowerThreshold   *big.Int
	ValidatorSetHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdatedEvent is a free log retrieval operation binding the contract event 0xe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c.
//
// Solidity: event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash)
func (_Wrappers *WrappersFilterer) FilterValidatorSetUpdatedEvent(opts *bind.FilterOpts, nonce []*big.Int) (*WrappersValidatorSetUpdatedEventIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Wrappers.contract.FilterLogs(opts, "ValidatorSetUpdatedEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return &WrappersValidatorSetUpdatedEventIterator{contract: _Wrappers.contract, event: "ValidatorSetUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdatedEvent is a free log subscription operation binding the contract event 0xe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c.
//
// Solidity: event ValidatorSetUpdatedEvent(uint256 indexed nonce, uint256 powerThreshold, bytes32 validatorSetHash)
func (_Wrappers *WrappersFilterer) WatchValidatorSetUpdatedEvent(opts *bind.WatchOpts, sink chan<- *WrappersValidatorSetUpdatedEvent, nonce []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _Wrappers.contract.WatchLogs(opts, "ValidatorSetUpdatedEvent", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappersValidatorSetUpdatedEvent)
				if err := _Wrappers.contract.UnpackLog(event, "ValidatorSetUpdatedEvent", log); err != nil {
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
func (_Wrappers *WrappersFilterer) ParseValidatorSetUpdatedEvent(log types.Log) (*WrappersValidatorSetUpdatedEvent, error) {
	event := new(WrappersValidatorSetUpdatedEvent)
	if err := _Wrappers.contract.UnpackLog(event, "ValidatorSetUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
