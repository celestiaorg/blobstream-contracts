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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataRootTupleRootNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidatorSetNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MalformedCurrentValidatorSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SuppliedValidatorSetInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRootTupleRoot\",\"type\":\"bytes32\"}],\"name\":\"DataRootTupleRootEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"powerThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"validatorSetHash\",\"type\":\"bytes32\"}],\"name\":\"ValidatorSetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_validatorSetHash\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_dataRootTupleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_eventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValidatorSetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorSetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_dataRootTupleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"submitDataRootTupleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_oldNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newPowerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_newValidatorSetHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"_currentValidatorSet\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_sigs\",\"type\":\"tuple[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"_tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0806040523461002a57306080526117299081610030823960805181818161041801526104f60152f35b600080fdfe6080604081815260048036101561001557600080fd5b600092833560e01c90816305d85c1314610993575080631f3302a91461085f578063226fe7be146106e15780634f1ef2861461047c57806352d1902d146104035780635433218c146103e4578063715018a614610387578063817f985b146103605780638da5cb5b14610337578063ad3cb1cc146102b2578063cdade86614610293578063e23eb32614610138578063e5a2b5d2146101155763f2fde38b146100bd57600080fd5b34610111576020366003190112610111576100d6610bad565b916100df610fca565b6001600160a01b038316156100fb57836100f884610ff6565b80f35b51631e4fbdf760e01b8152908101839052602490fd5b8280fd5b5050346101345781600319360112610134576020906097549051908152f35b5080fd5b50346101115760a036600319011261011157803591604435916001600160401b0360643581811161028f576101709036908401610ac4565b909160843590811161028b576101899036908501610af9565b9160985494609754956096549060018101809111610278578a036102695784830361025a576101c46101bb8488610c71565b88602435610cf7565b0361024c575092610234927f6614d037bde4905e31ca5ff05de61964c267f28b0320ed49e59f7d99752e1c4f979592879560209851898101906f0e8e4c2dce6c2c6e8d2dedc84c2e8c6d60831b82528c898201528960608201526060815261022b81610b29565b51902093610d79565b8460985584865260998352818187205551908152a280f35b8651630bbdaec960e11b8152fd5b50865163c6617b7b60e01b8152fd5b50865163e869766d60e01b8152fd5b634e487b7160e01b8c526011835260248cfd5b8780fd5b8680fd5b5050346101345781600319360112610134576020906098549051908152f35b509134610334578060031936011261033457815190828201908282106001600160401b038311176103215750610313935082526005815260208101640352e302e360dc1b815282519384926020845251809281602086015285850190610bde565b601f01601f19168101030190f35b634e487b7160e01b815260418552602490fd5b80fd5b50503461013457816003193601126101345760645490516001600160a01b039091168152602090f35b50346101115760203660031901126101115760209282913581526099845220549051908152f35b83346103345780600319360112610334576103a0610fca565b606480546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b5050346101345781600319360112610134576020906096549051908152f35b509134610334578060031936011261033457507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316300361046f57602090517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b5163703e46dd60e11b8152fd5b50908060031936011261011157610491610bad565b9060249384356001600160401b038111610134573660238201121561013457808501356104bd81610bc3565b946104ca85519687610b75565b81865260209182870193368a83830101116106dd578186928b8693018737880101526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081163081149081156106af575b5061069f5761052f610fca565b81169585516352d1902d60e01b815283818a818b5afa869181610670575b50610569575050505050505191634c9c8ce360e01b8352820152fd5b9088888894938c7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc9182810361065b5750853b15610647575080546001600160a01b031916821790558451889392917fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b8580a282511561062957505061061b9582915190845af4913d1561061f573d61060d61060482610bc3565b92519283610b75565b81528581943d92013e611690565b5080f35b5060609250611690565b95509550505050503461063b57505080f35b63b398979f60e01b8152fd5b8651634c9c8ce360e01b8152808501849052fd5b8751632a87526960e21b815280860191909152fd5b9091508481813d8311610698575b6106888183610b75565b8101031261028f5751903861054d565b503d61067e565b855163703e46dd60e11b81528890fd5b9050817f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141538610522565b8580fd5b50346101115760603660031901126101115780359060243591604435927ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0e93845460ff81881c16159485809661084c575b1580610829575b61081c57509186917fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c938660016001600160401b031983161789556107fd575b50610785818387610cf7565b85609855609655816097556107986112c1565b6107a06112c1565b6107a933610ff6565b82519182526020820152a26107bc578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808280f35b68ffffffffffffffffff19166801000000000000000117875538610779565b875162dc149f60e41b8152fd5b50303b1580610839575b15610739565b5060016001600160401b03831614610833565b5060016001600160401b03831610610732565b509034610111576003199160803684011261098f578160231936011261098f5781516001600160401b03918184018381118382101761097c5784526024358252602095604435878401526064359584871161013457606090873603011261033457845193606085018581108282111761096957865286830135908111610134578601903660238301121561033457828201356108fa81610b96565b9261090788519485610b75565b81845260248a85019260051b8201019236841161033457509060248a9201905b83821061095a57505050506109519495604491855260248101358886015201358584015235610c25565b90519015158152f35b81358152908201908201610927565b634e487b7160e01b835260418452602483fd5b634e487b7160e01b875260418252602487fd5b8380fd5b848385346101115760c03660031901126101115781359160443590606435906001600160401b0360843581811161028b576109d19036908401610ac4565b92909160a435908111610ac0576109eb9036908301610af9565b916098546097549b6096549160018101809111610aad578b03610a9f5750838603610a9057610a26610a1d8787610c71565b8d602435610cf7565b03610a8257509187989991610a6993610a6187897fe55fb3cbbfe29b13c7f8a35ef23127e7df9ab88df16bac166ad254a20f02414c9c610cf7565b958693610d79565b609655816097558460985582519182526020820152a280f35b8751630bbdaec960e11b8152fd5b50875163c6617b7b60e01b8152fd5b6368a35ffd60e11b81529050fd5b634e487b7160e01b8d526011845260248dfd5b8880fd5b9181601f84011215610af4578235916001600160401b038311610af4576020808501948460061b010111610af457565b600080fd5b9181601f84011215610af4578235916001600160401b038311610af45760208085019460608502010111610af457565b608081019081106001600160401b03821117610b4457604052565b634e487b7160e01b600052604160045260246000fd5b606081019081106001600160401b03821117610b4457604052565b90601f801991011681019081106001600160401b03821117610b4457604052565b6001600160401b038111610b445760051b60200190565b600435906001600160a01b0382168203610af457565b6001600160401b038111610b4457601f01601f191660200190565b60005b838110610bf15750506000910152565b8181015183820152602001610be1565b6001019081600111610c0f57565b634e487b7160e01b600052601160045260246000fd5b916098548311610c6957610c669260005260996020526040600020546020604051938051828601520151604084015260408352610c6183610b5a565b610ed2565b90565b505050600090565b60409160405180926020926020830195816040850160208952526060840192946000906000915b848310610cbd575050505050610cb7925003601f198101835282610b75565b51902090565b919395509193863560018060a01b038116809103610111578582819260019452858a01358682015201970193019091879593969492610c98565b916040519160208301936918da1958dadc1bda5b9d60b21b85526040840152606083015260808201526080815260a081018181106001600160401b03821117610b445760405251902090565b9190811015610d53576060020190565b634e487b7160e01b600052603260045260246000fd5b9190811015610d535760061b0190565b9493929460009360005b838110610dac575b505050505090915010610d9a57565b60405163cabeb65560e01b8152600490fd5b610db7818987610d43565b60209081810135159081610ec5575b81610eb0575b50610ea757610ddc828686610d69565b6001600160a01b03903581811690819003610af457610dfc848c8a610d43565b91610e63610e5b6040948551878101907f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252603c8b818301528152610e4181610b5a565b519020610e4d8261103f565b888884013593013591611469565b91909161151e565b1603610e975750610e75828686610d69565b01358601809611610c0f5786861015610e92576001905b01610d83565b610d8b565b51638baa579f60e01b8152600490fd5b50600190610e8c565b60ff9150610ebd9061103f565b161538610dcc565b6040810135159150610dc6565b6040820180519293919260018111610fab5750835151610fa2575b6020840191825182511115610f98576000610f47602092604051610f3660218287810194878652610f26815180928b8686019101610bde565b8101036001810184520182610b75565b604051928392839251928391610bde565b8101039060025afa15610f8c57600051935191825115610f755790610f7193949151905190611136565b1490565b5051600114159050610f85571490565b5050600090565b6040513d6000823e3d90fd5b5050505050600090565b50505050600090565b610fbc855151916020870151611079565b14610eed5750505050600090565b6064546001600160a01b03163303610fde57565b60405163118cdaa760e01b8152336004820152602490fd5b606480546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b3560ff81168103610af45790565b90610100918203918211610c0f57565b600019810191908211610c0f57565b91908203918211610c0f57565b909160005b60018481831b10156110aa5781018091111561107e57634e487b7160e01b600052601160045260246000fd5b5092919092610100908103908111610c0f576110c59061104d565b9160016110d18461105d565b1b916110dc8361105d565b81116110e85750505090565b91925090600183036110fc57505050600190565b8261110d611119946111139361106c565b9261106c565b90611079565b610c6690610c01565b8051821015610d535760209160051b010190565b90918215611260576001831461120f578351156111ca5761115683611302565b611169611163865161105d565b8661134f565b928181106111a6578161110d61118996936111839361106c565b90611136565b6111a08261119a610c66945161105d565b90611122565b51611427565b6111b09450611136565b906111c3610c66929161119a815161105d565b5190611427565b60405162461bcd60e51b815260206004820181905260248201527f6578706563746564206174206c65617374206f6e6520696e6e657220686173686044820152606490fd5b929150505161121b5790565b60405162461bcd60e51b815260206004820152601760248201527f756e657870656374656420696e6e6572206861736865730000000000000000006044820152606490fd5b60405162461bcd60e51b815260206004820152603360248201527f63616e6e6f742063616c6c20636f6d70757465526f6f744861736820776974686044820152722030206e756d626572206f66206c656176657360681b6064820152608490fd5b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0e5460401c16156112f057565b604051631afcd79f60e31b8152600490fd5b6001808210610af4578180916000925b6113385750506000198101908111610c0f576001901b9081146113325790565b60011c90565b90916000198114610c0f57810191811c9081611312565b91908251801590811561141c575b506113bf5761136b81610b96565b906113796040519283610b75565b808252601f1961138882610b96565b0136602084013760005b8181106113a0575090925050565b806113ad60019287611122565b516113b88286611122565b5201611392565b60405162461bcd60e51b815260206004820152602f60248201527f496e76616c69642072616e67653a205f626567696e206f72205f656e6420617260448201526e65206f7574206f6620626f756e647360881b6064820152608490fd5b90508110153861135d565b6114566000916020936040519085820192600160f81b84526021830152604182015260418152610f3681610b29565b8101039060025afa15610f8c5760005190565b9291907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083116115125760ff16601b81141580611507575b6114fb579160809493916020936040519384528484015260408301526060820152600093849182805260015afa156114ee5781516001600160a01b038116156114e8579190565b50600190565b50604051903d90823e3d90fd5b50505050600090600490565b50601c8114156114a1565b50505050600090600390565b600581101561167a578061152f5750565b6001810361157c5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606490fd5b600281036115c95760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606490fd5b600381036116215760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608490fd5b60041461162a57565b60405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608490fd5b634e487b7160e01b600052602160045260246000fd5b906116b757508051156116a557805190602001fd5b604051630a12f52160e11b8152600490fd5b815115806116ea575b6116c8575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156116c056fea2646970667358221220cdf9b4c7e4615a93c9460cea22a2ec60e132e8bb1d6f9e3e79121fe494b3fbf864736f6c63430008160033",
}

// WrappersABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappersMetaData.ABI instead.
var WrappersABI = WrappersMetaData.ABI

// WrappersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WrappersMetaData.Bin instead.
var WrappersBin = WrappersMetaData.Bin

// DeployWrappers deploys a new Ethereum contract, binding an instance of Wrappers to it.
func DeployWrappers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Wrappers, error) {
	parsed, err := WrappersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrappersBin), backend)
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

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Wrappers *WrappersCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wrappers.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Wrappers *WrappersSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Wrappers.Contract.UPGRADEINTERFACEVERSION(&_Wrappers.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Wrappers *WrappersCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Wrappers.Contract.UPGRADEINTERFACEVERSION(&_Wrappers.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wrappers *WrappersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wrappers.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wrappers *WrappersSession) Owner() (common.Address, error) {
	return _Wrappers.Contract.Owner(&_Wrappers.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wrappers *WrappersCallerSession) Owner() (common.Address, error) {
	return _Wrappers.Contract.Owner(&_Wrappers.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Wrappers *WrappersCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Wrappers.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Wrappers *WrappersSession) ProxiableUUID() ([32]byte, error) {
	return _Wrappers.Contract.ProxiableUUID(&_Wrappers.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Wrappers *WrappersCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Wrappers.Contract.ProxiableUUID(&_Wrappers.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x226fe7be.
//
// Solidity: function initialize(uint256 _nonce, uint256 _powerThreshold, bytes32 _validatorSetHash) returns()
func (_Wrappers *WrappersTransactor) Initialize(opts *bind.TransactOpts, _nonce *big.Int, _powerThreshold *big.Int, _validatorSetHash [32]byte) (*types.Transaction, error) {
	return _Wrappers.contract.Transact(opts, "initialize", _nonce, _powerThreshold, _validatorSetHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x226fe7be.
//
// Solidity: function initialize(uint256 _nonce, uint256 _powerThreshold, bytes32 _validatorSetHash) returns()
func (_Wrappers *WrappersSession) Initialize(_nonce *big.Int, _powerThreshold *big.Int, _validatorSetHash [32]byte) (*types.Transaction, error) {
	return _Wrappers.Contract.Initialize(&_Wrappers.TransactOpts, _nonce, _powerThreshold, _validatorSetHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x226fe7be.
//
// Solidity: function initialize(uint256 _nonce, uint256 _powerThreshold, bytes32 _validatorSetHash) returns()
func (_Wrappers *WrappersTransactorSession) Initialize(_nonce *big.Int, _powerThreshold *big.Int, _validatorSetHash [32]byte) (*types.Transaction, error) {
	return _Wrappers.Contract.Initialize(&_Wrappers.TransactOpts, _nonce, _powerThreshold, _validatorSetHash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wrappers *WrappersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wrappers.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wrappers *WrappersSession) RenounceOwnership() (*types.Transaction, error) {
	return _Wrappers.Contract.RenounceOwnership(&_Wrappers.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wrappers *WrappersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Wrappers.Contract.RenounceOwnership(&_Wrappers.TransactOpts)
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

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wrappers *WrappersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Wrappers.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wrappers *WrappersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Wrappers.Contract.TransferOwnership(&_Wrappers.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wrappers *WrappersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Wrappers.Contract.TransferOwnership(&_Wrappers.TransactOpts, newOwner)
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

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Wrappers *WrappersTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Wrappers.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Wrappers *WrappersSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Wrappers.Contract.UpgradeToAndCall(&_Wrappers.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Wrappers *WrappersTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Wrappers.Contract.UpgradeToAndCall(&_Wrappers.TransactOpts, newImplementation, data)
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

// WrappersInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Wrappers contract.
type WrappersInitializedIterator struct {
	Event *WrappersInitialized // Event containing the contract specifics and raw log

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
func (it *WrappersInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappersInitialized)
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
		it.Event = new(WrappersInitialized)
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
func (it *WrappersInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappersInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappersInitialized represents a Initialized event raised by the Wrappers contract.
type WrappersInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Wrappers *WrappersFilterer) FilterInitialized(opts *bind.FilterOpts) (*WrappersInitializedIterator, error) {

	logs, sub, err := _Wrappers.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WrappersInitializedIterator{contract: _Wrappers.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Wrappers *WrappersFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WrappersInitialized) (event.Subscription, error) {

	logs, sub, err := _Wrappers.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappersInitialized)
				if err := _Wrappers.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Wrappers *WrappersFilterer) ParseInitialized(log types.Log) (*WrappersInitialized, error) {
	event := new(WrappersInitialized)
	if err := _Wrappers.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Wrappers contract.
type WrappersOwnershipTransferredIterator struct {
	Event *WrappersOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WrappersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappersOwnershipTransferred)
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
		it.Event = new(WrappersOwnershipTransferred)
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
func (it *WrappersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappersOwnershipTransferred represents a OwnershipTransferred event raised by the Wrappers contract.
type WrappersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Wrappers *WrappersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WrappersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Wrappers.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WrappersOwnershipTransferredIterator{contract: _Wrappers.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Wrappers *WrappersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WrappersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Wrappers.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappersOwnershipTransferred)
				if err := _Wrappers.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Wrappers *WrappersFilterer) ParseOwnershipTransferred(log types.Log) (*WrappersOwnershipTransferred, error) {
	event := new(WrappersOwnershipTransferred)
	if err := _Wrappers.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
