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

// WrappersMetaData contains all meta data concerning the Wrappers contract.
var WrappersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidDataRootTupleToDataRootTupleRootProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"InvalidRowsToDataRootProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"InvalidSharesToRowsProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalDataLengthAndNumberOfSharesProofs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalRowsProofsAndRowsRootsNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalShareProofsAndRowsRootsNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractQuantumGravityBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"NamespaceID\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"NamespaceID\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"internalType\":\"NamespaceID\",\"name\":\"namespaceID\",\"type\":\"bytes8\"},{\"components\":[{\"internalType\":\"NamespaceID\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"NamespaceID\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowsRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowsProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60803461007457601f61153038819003918201601f19168301916001600160401b038311848410176100795780849260209460405283398101031261007457516001600160a01b0381169081900361007457600080546001600160a01b0319169190911790556040516114a090816100908239f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c80631f7b9136146100605763e78cea921461003257600080fd5b3461005b57600036600319011261005b576000546040516001600160a01b039091168152602090f35b600080fd5b3461005b57604036600319011261005b576001600160401b036004351161005b5760c06004353603600319011261005b57610140604052600435600401356001600160401b03811161005b57600435013660238201121561005b5760048101356100c981610a5b565b916100d76040519384610a3a565b8183526024602084019260051b8201019036821161005b5760248101925b8284106109a65784608052602460043501356001600160401b03811161005b57600435013660238201121561005b57600481013561013281610a5b565b916101406040519384610a3a565b8183526024602084019260051b8201019036821161005b5760248101925b8284106109305760a0859052610178600435604401610a72565b60c052600435606401356001600160401b03811161005b576101a1906004369181350101610a87565b60e052600435608401356001600160401b03811161005b57366023826004350101121561005b576004818135010135906101da82610a5b565b916101e86040519384610a3a565b80835260208301913660248360051b836004350101011161005b576024816004350101925b60248360051b8360043501010184106108fb5761010085815260043560a401356001600160401b03811161005b576004350180360390608060031983011261005b57604080519261025d84610a1f565b60048301358452602319011261005b5760405190604082018281106001600160401b038211176108e5576040526024810135825260448101356020830152602083019182526064810135906001600160401b03821161005b5760046102c89294939436920101610b22565b90604081019182528060a0608001526024359260018060a01b03600054169151905192516020604051948593631f3302a960e01b8552600485015280516024850152015160448301526080606483015260e48201908051916060608485015282518091526020610104850193019060005b8181106108cc575050508260209492604083878495015160a4850152015160c483015203915afa90811561064957600091610891575b501561087f5781515160e051510361086d5760005b8251518110156104375761040c6001600160401b0360c01b806103ac84606060800151610bcf565b5151169060206103c185606060800151610bcf565b5101511660406103d685606060800151610bcf565b5101519060405192602084015260288301526030908183015281526103fa81610a1f565b610405838651610bcf565b5184610c40565b1561041f5761041a90610bc0565b610384565b602490604051906301dd096960e21b82526004820152fd5b60a0515160e051510361085b576000805b60a05180518310156104965761049091610484602061046a8661048a95610bcf565b51015161047c86602060800151610bcf565b515190610c18565b90610c33565b91610bc0565b90610448565b506080515103610849576000805b60a051805183101561083e5760206104bf846104d193610bcf565b51015161047c84602060800151610bcf565b906001600160401b0360c01b604060800151169160406104f685606060800151610bcf565b5101516040519361050685610a1f565b8085526020850152604084015261052284602060800151610bcf565b51936001600160401b0360c01b6040608001511694608051926105458186610c33565b61054f8682610c18565b9461055986610a5b565b956105676040519788610a3a565b808752610576601f1991610a5b565b0160005b81811061082d575050865b82811061080f575050506105998451611061565b9460005b855181101561065557602060006105fc6105b7848a610bcf565b516105c0611042565b506105f08d6105e2604051938492878985015260218401526029830190610fb9565b03601f198101835282610a3a565b60405191828092610fb9565b039060025afa1561064957610644906000516040519061061b82610a1f565b8b82528b60208301526040820152610633828a610bcf565b5261063e8189610bcf565b50610bc0565b61059d565b6040513d6000823e3d90fd5b5092965092509360009361066d604087015151611061565b936000955b87518082141580610801575b156106cb57816106946106bf926106c5946110b1565b906106a38a60408d0151610bcf565b516106ae8b8b610bcf565b526106b98a8a610bcf565b50610c33565b96610bc0565b95610672565b5050935090959193506106e1602086015161112f565b808060011b04600214811517156107eb576000816107109360018394811b92811b106107e2575b829089611173565b505093905b6040860151805186101561074b5761073d6040926107368861074394610bcf565b519061131f565b95610bc0565b949050610715565b509294509250936001600160401b0360c01b8251166001600160401b0360c01b8251161491826107c2575b826107af575b505015610796576107909161048a91610c33565b906104a4565b60405163cef8a4cb60e01b815260048101849052602490fd5b604091925081015191015114848061077c565b602080820151908301516001600160c01b03199081169116149250610776565b60019150610708565b634e487b7160e01b600052601160045260246000fd5b50604089015151881061067e565b8061081d6108289284610bcf565b51610633828a610bcf565b610585565b806060602080938b0101520161057a565b602060405160018152f35b60405163efc454a560e01b8152600490fd5b604051636031acbb60e01b8152600490fd5b604051634f17331d60e01b8152600490fd5b60405163320f037560e21b8152600490fd5b90506020813d6020116108c4575b816108ac60209383610a3a565b8101031261005b5751801515810361005b578361036f565b3d915061089f565b8251855287955060209485019490920191600101610339565b634e487b7160e01b600052604160045260246000fd5b6001600160401b0384351161005b576020806024926109233685893588600435010101610b22565b815201940193905061020d565b83356001600160401b03811161005b5782016060602319823603011261005b576040519161095d83610a1f565b60248201358352604482013560208401526064820135926001600160401b03841161005b57610996602094936024869536920101610a87565b604082015281520193019261015e565b83356001600160401b03811161005b578201903660438301121561005b576024820135906001600160401b0382116108e5576040516109ef601f8401601f191660200182610a3a565b8281526044933685858301011161005b5760208481969582966000940183860137830101528152019301926100f5565b606081019081106001600160401b038211176108e557604052565b90601f801991011681019081106001600160401b038211176108e557604052565b6001600160401b0381116108e55760051b60200190565b35906001600160c01b03198216820361005b57565b81601f8201121561005b57803590610a9e82610a5b565b92604090610aae82519586610a3a565b838552602091828601918360608097028601019481861161005b578401925b858410610ade575050505050505090565b868483031261005b578487918451610af581610a1f565b610afe87610a72565b8152610b0b838801610a72565b838201528587013586820152815201930192610acd565b91909160608184031261005b5760405190610b3c82610a1f565b819381356001600160401b03811161005b5782019080601f8301121561005b57813590610b6882610a5b565b91610b766040519384610a3a565b808352602093848085019260051b82010192831161005b578401905b828210610bb15750505083528082013590830152604090810135910152565b81358152908401908401610b92565b60001981146107eb5760010190565b8051821015610be35760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b906101009182039182116107eb57565b6000198101919082116107eb57565b919082039182116107eb57565b60010190816001116107eb57565b919082018092116107eb57565b6040820180519193909291600190818111610eed5750825151610ee3575b6020830192835185511115610ed85760206105e293610c916040516105f081600098899586888401526021830190610fb9565b039060025afa15610ecd5782519481515115610eb557829081958051925b610db0575b505160001991828201918211610d9c5703610d4f575b8293925b610cdc575b50505050501490565b9091929394818601868111610d3b5783518051821015610d3357610d0a9291610d0491610bcf565b51610fe4565b94848101809111610d1f579392919083610cce565b634e487b7160e01b84526011600452602484fd5b505094610cd3565b634e487b7160e01b85526011600452602485fd5b9091939482515182870190878211610d9c57811015610d915790610d77610d7e928551610bcf565b5190610fe4565b94848101809111610d1f57939190610cca565b505050509250505090565b634e487b7160e01b86526011600452602486fd5b9091819793949697519387891b948515610ea157891c94858a1b95808704821490151715610e8d57610de29086610c33565b60001992818401918211610e6b578551821015610e7f575094865151928a01928a8411610e6b57831015610e5d57610e1b908451610c18565b88831b1115610e4957610d77610e32928751610bcf565b965b868101809111610d9c57959392919084610caf565b610d04610e57928751610bcf565b96610e34565b505050505050509250505090565b634e487b7160e01b89526011600452602489fd5b949350509796949350610cb4565b634e487b7160e01b88526011600452602488fd5b634e487b7160e01b88526012600452602488fd5b519495949293505003610ec757501490565b91505090565b6040513d84823e3d90fd5b505050505050600090565b5050505050600090565b610efe845151916020860151610f0d565b14610c5e575050505050600090565b909160005b60018481831b1015610f3e57810180911115610f1257634e487b7160e01b600052601160045260246000fd5b50929190926101009081039081116107eb57610f5990610bf9565b916001610f6584610c09565b1b91610f7083610c09565b8111610f7c5750505090565b9192509060018303610f9057505050600190565b82610fa1610fad94610fa793610c18565b92610c18565b90610f0d565b610fb690610c25565b90565b9081519160005b838110610fd1575050016000815290565b8060208092840101518185015201610fc0565b9060405191600160f81b6020840152602183015260418201526041815260808101908082106001600160401b038311176108e55781600091602093604052607f19906110308382610fb9565b03019060025afa156106495760005190565b6040519061104f82610a1f565b60006040838281528260208201520152565b9061106b82610a5b565b6110786040519182610a3a565b8281528092611089601f1991610a5b565b019060005b82811061109a57505050565b6020906110a5611042565b8282850101520161108e565b90918160005b83151580611124575b156110d8576110ce90610bc0565b9260011c926110b7565b92506110e49193610c18565b90816000925b611110575060001982019182116107eb5781600192821161110a57501b90565b90501b90565b9161111a90610bc0565b9160011c806110ea565b5083600116156110c0565b600180821061005b57600082805b611160575060001981019081116107eb5781901b91821461115c575090565b1c90565b9061116a90610bc0565b90821c8061113d565b939594929092611181611042565b50600161118e8484610c18565b1461123b578451821180159061122d575b611211576111cc906111b96111b48585610c18565b61112f565b976111c48986610c33565b858789611173565b5091976111e796929591946111e19190610c33565b91611173565b919490939290911515600114611208576112009161131f565b929190600090565b50929190600090565b93604091935061122595969250015161129f565b929391929091565b50602085015183101561119f565b9394959180869492945111159081611272575b506112615750604061122594015161129f565b805161122595509193919250611280565b90506020860151113861124e565b906112979291949394611291611042565b506112b5565b919392909190565b906112ad9291611291611042565b919390929190565b908093926112c1611042565b508251908115918215611314575b508115611309575b506112f8576112e591610bcf565b51600183018093116107eb579190600090565b5050611302611042565b9190600190565b9050811015386112d7565b8310159150386112cf565b90611328611042565b50815181516001600160c01b0319908116939181166113478582611431565b948183036114005782945b8360208601511691604080960151868660208501511693015193875195600160f81b60208801526021870152602986015260318501526051840152605983015260618201526061815260a08101908082106001600160401b038311176108e557818452602091600091609f19906113c98382610fb9565b03019060025afa156113f55760005192818351956113e687610a1f565b16855216602084015282015290565b50513d6000823e3d90fd5b808303611414578260208501511694611352565b61142b836020860151168460208801511690611450565b94611352565b906001600160c01b0319808216908316101561144b575090565b905090565b906001600160c01b0319808216908316111561144b57509056fea264697066735822122014c2f443c36de732562dc615736a08bcf1894adcec1c76fcebefc718fa64a6ad64736f6c63430008140033",
}

// WrappersABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappersMetaData.ABI instead.
var WrappersABI = WrappersMetaData.ABI

// WrappersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WrappersMetaData.Bin instead.
var WrappersBin = WrappersMetaData.Bin

// DeployWrappers deploys a new Ethereum contract, binding an instance of Wrappers to it.
func DeployWrappers(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *Wrappers, error) {
	parsed, err := WrappersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrappersBin), backend, _bridge)
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

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Wrappers *WrappersCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wrappers.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Wrappers *WrappersSession) Bridge() (common.Address, error) {
	return _Wrappers.Contract.Bridge(&_Wrappers.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Wrappers *WrappersCallerSession) Bridge() (common.Address, error) {
	return _Wrappers.Contract.Bridge(&_Wrappers.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x1f7b9136.
//
// Solidity: function verify((bytes[],(uint256,uint256,(bytes8,bytes8,bytes32)[])[],bytes8,(bytes8,bytes8,bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof, bytes32 _root) view returns(bool)
func (_Wrappers *WrappersCaller) Verify(opts *bind.CallOpts, _sharesProof SharesProof, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _Wrappers.contract.Call(opts, &out, "verify", _sharesProof, _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x1f7b9136.
//
// Solidity: function verify((bytes[],(uint256,uint256,(bytes8,bytes8,bytes32)[])[],bytes8,(bytes8,bytes8,bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof, bytes32 _root) view returns(bool)
func (_Wrappers *WrappersSession) Verify(_sharesProof SharesProof, _root [32]byte) (bool, error) {
	return _Wrappers.Contract.Verify(&_Wrappers.CallOpts, _sharesProof, _root)
}

// Verify is a free data retrieval call binding the contract method 0x1f7b9136.
//
// Solidity: function verify((bytes[],(uint256,uint256,(bytes8,bytes8,bytes32)[])[],bytes8,(bytes8,bytes8,bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32),(bytes32[],uint256,uint256))) _sharesProof, bytes32 _root) view returns(bool)
func (_Wrappers *WrappersCallerSession) Verify(_sharesProof SharesProof, _root [32]byte) (bool, error) {
	return _Wrappers.Contract.Verify(&_Wrappers.CallOpts, _sharesProof, _root)
}
