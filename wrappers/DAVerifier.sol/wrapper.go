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
	Height     *big.Int
	DataRoot   [32]byte
	SquareSize *big.Int
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidDataRootTupleToDataRootTupleRootProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"InvalidRowsToDataRootProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"InvalidSharesToRowsProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalDataLengthAndNumberOfSharesProofs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalRowsProofsAndRowsRootsNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalShareProofsAndRowsRootsNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractQuantumGravityBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beginKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endKey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"NamespaceID\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"NamespaceID\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"sideNodes\",\"type\":\"tuple[]\"}],\"internalType\":\"structNamespaceMerkleMultiproof[]\",\"name\":\"shareProofs\",\"type\":\"tuple[]\"},{\"internalType\":\"NamespaceID\",\"name\":\"namespaceID\",\"type\":\"bytes8\"},{\"components\":[{\"internalType\":\"NamespaceID\",\"name\":\"min\",\"type\":\"bytes8\"},{\"internalType\":\"NamespaceID\",\"name\":\"max\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode[]\",\"name\":\"rowsRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof[]\",\"name\":\"rowsProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"squareSize\",\"type\":\"uint256\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"attestationProof\",\"type\":\"tuple\"}],\"internalType\":\"structSharesProof\",\"name\":\"_sharesProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60803461007457601f61153538819003918201601f19168301916001600160401b038311848410176100795780849260209460405283398101031261007457516001600160a01b0381169081900361007457600080546001600160a01b0319169190911790556040516114a590816100908239f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806330f56249146100605763e78cea921461003257600080fd5b3461005b57600036600319011261005b576000546040516001600160a01b039091168152602090f35b600080fd5b3461005b57604036600319011261005b576001600160401b036004351161005b5760c06004353603600319011261005b57610140604052600435600401356001600160401b03811161005b57600435013660238201121561005b5760048101356100c981610a60565b916100d76040519384610a3f565b8183526024602084019260051b8201019036821161005b5760248101925b8284106109955784608052602460043501356001600160401b03811161005b57600435013660238201121561005b57600481013561013281610a60565b916101406040519384610a3f565b8183526024602084019260051b8201019036821161005b5760248101925b82841061091f5760a0859052610178600435604401610a77565b60c052600435606401356001600160401b03811161005b576101a1906004369181350101610a8c565b60e052600435608401356001600160401b03811161005b57366023826004350101121561005b576004818135010135906101da82610a60565b916101e86040519384610a3f565b80835260208301913660248360051b836004350101011161005b576024816004350101925b60248360051b8360043501010184106108ea5761010085815260043560a401356001600160401b03811161005b57600435018036039060a060031983011261005b5760606040519261025e84610a24565b60048301358452602319011261005b576040519061027b82610a24565b602481013582526044810135602083015260648101356040830152602083019182526084810135906001600160401b03821161005b5760046102c39294939436920101610b27565b90604081019182528060a0608001526024359260018060a01b0360005416915190519251604080519485936367a0c5ab60e11b8552600485015280516024850152602081015160448501520151606483015260a06084830152610104820190805191606060a485015282518091526020610124850193019060005b8181106108d1575050508260209492604083878495015160c4850152015160e483015203915afa90811561064e57600091610896575b50156108845781515160e05151036108725760005b82515181101561043c576104116001600160401b0360c01b806103b184606060800151610bd4565b5151169060206103c685606060800151610bd4565b5101511660406103db85606060800151610bd4565b5101519060405192602084015260288301526030908183015281526103ff81610a24565b61040a838651610bd4565b5184610c45565b156104245761041f90610bc5565b610389565b602490604051906301dd096960e21b82526004820152fd5b60a0515160e0515103610860576000805b60a051805183101561049b5761049591610489602061046f8661048f95610bd4565b51015161048186602060800151610bd4565b515190610c1d565b90610c38565b91610bc5565b9061044d565b50608051510361084e576000805b60a05180518310156108435760206104c4846104d693610bd4565b51015161048184602060800151610bd4565b906001600160401b0360c01b604060800151169160406104fb85606060800151610bd4565b5101516040519361050b85610a24565b8085526020850152604084015261052784602060800151610bd4565b51936001600160401b0360c01b60406080015116946080519261054a8186610c38565b6105548682610c1d565b9461055e86610a60565b9561056c6040519788610a3f565b80875261057b601f1991610a60565b0160005b818110610832575050865b8281106108145750505061059e8451611066565b9460005b855181101561065a57602060006106016105bc848a610bd4565b516105c5611047565b506105f58d6105e7604051938492878985015260218401526029830190610fbe565b03601f198101835282610a3f565b60405191828092610fbe565b039060025afa1561064e57610649906000516040519061062082610a24565b8b82528b60208301526040820152610638828a610bd4565b526106438189610bd4565b50610bc5565b6105a2565b6040513d6000823e3d90fd5b50929650925093600093610672604087015151611066565b936000955b87518082141580610806575b156106d057816106996106c4926106ca946110b6565b906106a88a60408d0151610bd4565b516106b38b8b610bd4565b526106be8a8a610bd4565b50610c38565b96610bc5565b95610677565b5050935090959193506106e66020860151611134565b808060011b04600214811517156107f0576000816107159360018394811b92811b106107e7575b829089611178565b505093905b604086015180518610156107505761074260409261073b8861074894610bd4565b5190611324565b95610bc5565b94905061071a565b509294509250936001600160401b0360c01b8251166001600160401b0360c01b8251161491826107c7575b826107b4575b50501561079b576107959161048f91610c38565b906104a9565b60405163cef8a4cb60e01b815260048101849052602490fd5b6040919250810151910151148480610781565b602080820151908301516001600160c01b0319908116911614925061077b565b6001915061070d565b634e487b7160e01b600052601160045260246000fd5b506040890151518810610683565b8061082261082d9284610bd4565b51610638828a610bd4565b61058a565b806060602080938b0101520161057f565b602060405160018152f35b60405163efc454a560e01b8152600490fd5b604051636031acbb60e01b8152600490fd5b604051634f17331d60e01b8152600490fd5b60405163320f037560e21b8152600490fd5b90506020813d6020116108c9575b816108b160209383610a3f565b8101031261005b5751801515810361005b5783610374565b3d91506108a4565b825185528795506020948501949092019160010161033e565b6001600160401b0384351161005b576020806024926109123685893588600435010101610b27565b815201940193905061020d565b83356001600160401b03811161005b5782016060602319823603011261005b576040519161094c83610a24565b60248201358352604482013560208401526064820135926001600160401b03841161005b57610985602094936024869536920101610a8c565b604082015281520193019261015e565b83356001600160401b03811161005b578201903660438301121561005b576024820135906001600160401b038211610a0e576040516109de601f8401601f191660200182610a3f565b8281526044933685858301011161005b5760208481969582966000940183860137830101528152019301926100f5565b634e487b7160e01b600052604160045260246000fd5b606081019081106001600160401b03821117610a0e57604052565b90601f801991011681019081106001600160401b03821117610a0e57604052565b6001600160401b038111610a0e5760051b60200190565b35906001600160c01b03198216820361005b57565b81601f8201121561005b57803590610aa382610a60565b92604090610ab382519586610a3f565b838552602091828601918360608097028601019481861161005b578401925b858410610ae3575050505050505090565b868483031261005b578487918451610afa81610a24565b610b0387610a77565b8152610b10838801610a77565b838201528587013586820152815201930192610ad2565b91909160608184031261005b5760405190610b4182610a24565b819381356001600160401b03811161005b5782019080601f8301121561005b57813590610b6d82610a60565b91610b7b6040519384610a3f565b808352602093848085019260051b82010192831161005b578401905b828210610bb65750505083528082013590830152604090810135910152565b81358152908401908401610b97565b60001981146107f05760010190565b8051821015610be85760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b906101009182039182116107f057565b6000198101919082116107f057565b919082039182116107f057565b60010190816001116107f057565b919082018092116107f057565b6040820180519193909291600190818111610ef25750825151610ee8575b6020830192835185511115610edd5760206105e793610c966040516105f581600098899586888401526021830190610fbe565b039060025afa15610ed25782519481515115610eba57829081958051925b610db5575b505160001991828201918211610da15703610d54575b8293925b610ce1575b50505050501490565b9091929394818601868111610d405783518051821015610d3857610d0f9291610d0991610bd4565b51610fe9565b94848101809111610d24579392919083610cd3565b634e487b7160e01b84526011600452602484fd5b505094610cd8565b634e487b7160e01b85526011600452602485fd5b9091939482515182870190878211610da157811015610d965790610d7c610d83928551610bd4565b5190610fe9565b94848101809111610d2457939190610ccf565b505050509250505090565b634e487b7160e01b86526011600452602486fd5b9091819793949697519387891b948515610ea657891c94858a1b95808704821490151715610e9257610de79086610c38565b60001992818401918211610e70578551821015610e84575094865151928a01928a8411610e7057831015610e6257610e20908451610c1d565b88831b1115610e4e57610d7c610e37928751610bd4565b965b868101809111610da157959392919084610cb4565b610d09610e5c928751610bd4565b96610e39565b505050505050509250505090565b634e487b7160e01b89526011600452602489fd5b949350509796949350610cb9565b634e487b7160e01b88526011600452602488fd5b634e487b7160e01b88526012600452602488fd5b519495949293505003610ecc57501490565b91505090565b6040513d84823e3d90fd5b505050505050600090565b5050505050600090565b610f03845151916020860151610f12565b14610c63575050505050600090565b909160005b60018481831b1015610f4357810180911115610f1757634e487b7160e01b600052601160045260246000fd5b50929190926101009081039081116107f057610f5e90610bfe565b916001610f6a84610c0e565b1b91610f7583610c0e565b8111610f815750505090565b9192509060018303610f9557505050600190565b82610fa6610fb294610fac93610c1d565b92610c1d565b90610f12565b610fbb90610c2a565b90565b9081519160005b838110610fd6575050016000815290565b8060208092840101518185015201610fc5565b9060405191600160f81b6020840152602183015260418201526041815260808101908082106001600160401b03831117610a0e5781600091602093604052607f19906110358382610fbe565b03019060025afa1561064e5760005190565b6040519061105482610a24565b60006040838281528260208201520152565b9061107082610a60565b61107d6040519182610a3f565b828152809261108e601f1991610a60565b019060005b82811061109f57505050565b6020906110aa611047565b82828501015201611093565b90918160005b83151580611129575b156110dd576110d390610bc5565b9260011c926110bc565b92506110e99193610c1d565b90816000925b611115575060001982019182116107f05781600192821161110f57501b90565b90501b90565b9161111f90610bc5565b9160011c806110ef565b5083600116156110c5565b600180821061005b57600082805b611165575060001981019081116107f05781901b918214611161575090565b1c90565b9061116f90610bc5565b90821c80611142565b939594929092611186611047565b5060016111938484610c1d565b146112405784518211801590611232575b611216576111d1906111be6111b98585610c1d565b611134565b976111c98986610c38565b858789611178565b5091976111ec96929591946111e69190610c38565b91611178565b91949093929091151560011461120d5761120591611324565b929190600090565b50929190600090565b93604091935061122a9596925001516112a4565b929391929091565b5060208501518310156111a4565b9394959180869492945111159081611277575b506112665750604061122a9401516112a4565b805161122a95509193919250611285565b905060208601511138611253565b9061129c9291949394611296611047565b506112ba565b919392909190565b906112b29291611296611047565b919390929190565b908093926112c6611047565b508251908115918215611319575b50811561130e575b506112fd576112ea91610bd4565b51600183018093116107f0579190600090565b5050611307611047565b9190600190565b9050811015386112dc565b8310159150386112d4565b9061132d611047565b50815181516001600160c01b03199081169391811661134c8582611436565b948183036114055782945b8360208601511691604080960151868660208501511693015193875195600160f81b60208801526021870152602986015260318501526051840152605983015260618201526061815260a08101908082106001600160401b03831117610a0e57818452602091600091609f19906113ce8382610fbe565b03019060025afa156113fa5760005192818351956113eb87610a24565b16855216602084015282015290565b50513d6000823e3d90fd5b808303611419578260208501511694611357565b611430836020860151168460208801511690611455565b94611357565b906001600160c01b03198082169083161015611450575090565b905090565b906001600160c01b0319808216908316111561145057509056fea2646970667358221220a8fb3a15ab49be121bc2f13f4bc4a830cc67d1e4ef78b7752dd841235e0bfb4b64736f6c63430008140033",
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

// Verify is a free data retrieval call binding the contract method 0x30f56249.
//
// Solidity: function verify((bytes[],(uint256,uint256,(bytes8,bytes8,bytes32)[])[],bytes8,(bytes8,bytes8,bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32,uint256),(bytes32[],uint256,uint256))) _sharesProof, bytes32 _root) view returns(bool)
func (_Wrappers *WrappersCaller) Verify(opts *bind.CallOpts, _sharesProof SharesProof, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _Wrappers.contract.Call(opts, &out, "verify", _sharesProof, _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x30f56249.
//
// Solidity: function verify((bytes[],(uint256,uint256,(bytes8,bytes8,bytes32)[])[],bytes8,(bytes8,bytes8,bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32,uint256),(bytes32[],uint256,uint256))) _sharesProof, bytes32 _root) view returns(bool)
func (_Wrappers *WrappersSession) Verify(_sharesProof SharesProof, _root [32]byte) (bool, error) {
	return _Wrappers.Contract.Verify(&_Wrappers.CallOpts, _sharesProof, _root)
}

// Verify is a free data retrieval call binding the contract method 0x30f56249.
//
// Solidity: function verify((bytes[],(uint256,uint256,(bytes8,bytes8,bytes32)[])[],bytes8,(bytes8,bytes8,bytes32)[],(bytes32[],uint256,uint256)[],(uint256,(uint256,bytes32,uint256),(bytes32[],uint256,uint256))) _sharesProof, bytes32 _root) view returns(bool)
func (_Wrappers *WrappersCallerSession) Verify(_sharesProof SharesProof, _root [32]byte) (bool, error) {
	return _Wrappers.Contract.Verify(&_Wrappers.CallOpts, _sharesProof, _root)
}
