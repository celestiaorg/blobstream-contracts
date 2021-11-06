package solidity

import (
	"bufio"
	"bytes"
	"crypto/ecdsa"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var cosmosCfg *cosmtypes.Config

func init() {
	if _, ok := os.LookupEnv("QGB_TEST_EVM_RPC"); !ok {
		os.Setenv("QGB_TEST_EVM_RPC", "http://localhost:8545")
	}
	if _, ok := os.LookupEnv("QGB_TEST_COVERAGE"); !ok {
		os.Setenv("QGB_TEST_COVERAGE", "0")
	}
	if _, ok := os.LookupEnv("QGB_TEST_COVERAGE_MODE"); !ok {
		os.Setenv("QGB_TEST_COVERAGE_MODE", "set")
	}
	if _, ok := os.LookupEnv("QGB_TEST_BECH32_PREFIX"); !ok {
		os.Setenv("QGB_TEST_BECH32_PREFIX", "inj")
	}
	if _, ok := os.LookupEnv("QGB_TEST_BIP44_COIN"); !ok {
		os.Setenv("QGB_TEST_BIP44_COIN", "60")
	}
	if _, ok := os.LookupEnv("QGB_TEST_BIP44_HDPATH"); !ok {
		os.Setenv("QGB_TEST_BIP44_HDPATH", "m/44'/60'/0'/0")
	}
	if _, ok := os.LookupEnv("QGB_TEST_SIGNING_ALGO"); !ok {
		os.Setenv("QGB_TEST_SIGNING_ALGO", "secp256k1")
	}

	readEnv()

	var (
		// Bech32Prefix defines the Bech32 prefix used for EthAccounts on the Celestia Chain
		Bech32Prefix = os.Getenv("QGB_TEST_BECH32_PREFIX")

		// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
		Bech32PrefixAccAddr = Bech32Prefix
		// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
		Bech32PrefixAccPub = Bech32Prefix + cosmtypes.PrefixPublic
		// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
		Bech32PrefixValAddr = Bech32Prefix + cosmtypes.PrefixValidator + cosmtypes.PrefixOperator
		// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
		Bech32PrefixValPub = Bech32Prefix + cosmtypes.PrefixValidator + cosmtypes.PrefixOperator + cosmtypes.PrefixPublic
		// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
		Bech32PrefixConsAddr = Bech32Prefix + cosmtypes.PrefixValidator + cosmtypes.PrefixConsensus
		// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
		Bech32PrefixConsPub = Bech32Prefix + cosmtypes.PrefixValidator + cosmtypes.PrefixConsensus + cosmtypes.PrefixPublic

		// Bip44CoinType satisfies EIP84. See https://github.com/ethereum/EIPs/issues/84 for more info.
		Bip44CoinType, _ = strconv.Atoi(os.Getenv("QGB_TEST_BIP44_COIN"))
		// BIP44HDPath is the BIP44 HD path used on Ethereum.
		BIP44HDPath = os.Getenv("QGB_TEST_BIP44_HDPATH")
	)

	cosmosCfg = cosmtypes.GetConfig()
	cosmosCfg.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	cosmosCfg.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	cosmosCfg.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	cosmosCfg.SetCoinType(uint32(Bip44CoinType))
	cosmosCfg.SetFullFundraiserPath(BIP44HDPath)

	for idx, a := range EthAccounts {
		a.Parse()
		EthAccounts[idx] = a
	}

	for idx, a := range CosmosAccounts {
		a.Parse()
		CosmosAccounts[idx] = a
	}
}

// readEnv is a special utility that reads `.env` file into actual environment variables
// of the current app, similar to `dotenv` Node package.
func readEnv() {
	if envdata, _ := ioutil.ReadFile(".env"); len(envdata) > 0 {
		s := bufio.NewScanner(bytes.NewReader(envdata))
		for s.Scan() {
			parts := strings.Split(s.Text(), "=")
			if len(parts) != 2 {
				continue
			}
			strValue := strings.Trim(parts[1], `"`)
			_ = os.Setenv(parts[0], strValue)
		}
	}
}

// Ganache snapshot
// Mnemonic:      concert load couple harbor equip island argue ramp clarify fence smart topic
// Base HD Path:  m/44'/60'/0'/0/{account_index}
var EthAccounts = []Account{
	{Address: "0x5409ED021D9299bf6814279A6A1411A7e866A631", Key: "0xf2f48ee19680706196e2e339e5da3491186e0c4c5030670656b0e0164837257d"},
	{Address: "0x6Ecbe1DB9EF729CBe972C83Fb886247691Fb6beb", Key: "0x5d862464fe9303452126c8bc94274b8c5f9874cbd219789b3eb2128075a76f72"},
	{Address: "0xE36Ea790bc9d7AB70C55260C66D52b1eca985f84", Key: "0xdf02719c4df8b9b8ac7f551fcb5d9ef48fa27eef7a66453879f4d8fdc6e78fb1"},
	{Address: "0xE834EC434DABA538cd1b9Fe1582052B880BD7e63", Key: "0xff12e391b79415e941a94de3bf3a9aee577aed0731e297d5cfa0b8a1e02fa1d0"},
	{Address: "0x78dc5D2D739606d31509C31d654056A45185ECb6", Key: "0x752dd9cf65e68cfaba7d60225cbdbc1f4729dd5e5507def72815ed0d8abc6249"},
	{Address: "0xA8dDa8d7F5310E4A9E24F8eBA77E091Ac264f872", Key: "0xefb595a0178eb79a8df953f87c5148402a224cdf725e88c0146727c6aceadccd"},
	{Address: "0x06cEf8E666768cC40Cc78CF93d9611019dDcB628", Key: "0x83c6d2cc5ddcf9711a6d59b417dc20eb48afd58d45290099e5987e3d768f328f"},
	{Address: "0x4404ac8bd8F9618D27Ad2f1485AA1B2cFD82482D", Key: "0xbb2d3f7c9583780a7d3904a2f55d792707c345f21de1bacb2d389934d82796b2"},
	{Address: "0x7457d5E02197480Db681D3fdF256c7acA21bDc12", Key: "0xb2fd4d29c1390b71b8795ae81196bfd60293adf99f9d32a0aff06288fcdac55f"},
	{Address: "0x91c987bf62D25945dB517BDAa840A6c661374402", Key: "0x23cb7121166b9a2f93ae0b7c05bde02eae50d64449b2cbb42bc84e9d38d6cc89"},
}

var CosmosAccounts = []Account{
	// validator 1
	{Mnemonic: "copper push brief egg scan entry inform record adjust fossil boss egg comic alien upon aspect dry avoid interest fury window hint race symptom"},
	// validator 2
	{Mnemonic: "maximum display century economy unlock van census kite error heart snow filter midnight usage egg venture cash kick motor survey drastic edge muffin visual"},
	// validator 3
	{Mnemonic: "keep liar demand upon shed essence tip undo eagle run people strong sense another salute double peasant egg royal hair report winner student diamond"},
	// user
	{Mnemonic: "pony glide frown crisp unfold lawn cup loan trial govern usual matrix theory wash fresh address pioneer between meadow visa buffalo keep gallery swear"},
}

func getEthAddresses(accounts ...Account) []common.Address {
	addresses := make([]common.Address, 0, len(accounts))
	for _, a := range accounts {
		addresses = append(addresses, a.EthAddress)
	}

	return addresses
}

func getSigningKeys(accounts ...Account) []*ecdsa.PrivateKey {
	privkeys := make([]*ecdsa.PrivateKey, 0, len(accounts))
	for _, a := range accounts {
		privkeys = append(privkeys, a.EthPrivKey)
	}

	return privkeys
}

func getSigningKeysForAddresses(addresses []common.Address, accounts ...Account) []*ecdsa.PrivateKey {
	privkeys := make([]*ecdsa.PrivateKey, len(addresses))

	for i, a := range addresses {
		for _, acc := range accounts {
			if acc.EthAddress == a {
				privkeys[i] = acc.EthPrivKey
				break
			}
		}
	}

	return privkeys
}

type Account struct {
	Address  string
	Key      string
	Mnemonic string

	EthAddress common.Address
	EthPrivKey *ecdsa.PrivateKey

	CosmosAccAddress cosmtypes.AccAddress
	CosmosValAddress cosmtypes.ValAddress
}

func (a *Account) Parse() {
	if len(a.Mnemonic) > 0 {
		// derive address and privkey from the provided mnemonic

		algo, err := keyring.NewSigningAlgoFromString(os.Getenv("QGB_TEST_SIGNING_ALGO"), keyring.SigningAlgoList{
			hd.Secp256k1,
		})
		orPanic(err)

		pkBytes, err := algo.Derive()(a.Mnemonic, "", cosmosCfg.GetFullFundraiserPath())

		a.EthPrivKey, err = crypto.ToECDSA(pkBytes)
		orPanic(err)

		a.Address = crypto.PubkeyToAddress(a.EthPrivKey.PublicKey).Hex()
	} else if len(a.Key) > 0 {
		// privkey should be parsed as hex
		pk, err := crypto.HexToECDSA(a.Key[2:])
		orPanic(err)
		a.EthPrivKey = pk
	}

	if ethAddress := common.HexToAddress(a.Address); common.IsHexAddress(a.Address) && ethAddress != zeroAddress {
		// provided an Eth address

		if a.EthPrivKey != nil {
			if !bytes.Equal(crypto.PubkeyToAddress(a.EthPrivKey.PublicKey).Bytes(), ethAddress.Bytes()) {
				panic(errors.Errorf("privkey doesn't match address: %s", ethAddress.Hex()))
			}
		}

		a.EthAddress = ethAddress
		a.CosmosAccAddress = cosmtypes.AccAddress(a.EthAddress.Bytes())
		a.CosmosValAddress = cosmtypes.ValAddress(a.EthAddress.Bytes())
	} else if accAddress, err := cosmtypes.AccAddressFromBech32(a.Address); err == nil {
		// provided a Bech32 address

		a.EthAddress = common.BytesToAddress(accAddress.Bytes())
		if a.EthAddress == zeroAddress {
			panic(errors.Errorf("unsupported address: %s (check your Bech32 prefix)", a.Address))
		}

		if a.EthPrivKey != nil {
			if !bytes.Equal(crypto.PubkeyToAddress(a.EthPrivKey.PublicKey).Bytes(), accAddress.Bytes()) {
				panic(errors.Errorf("privkey doesn't match address: %s", accAddress.String()))
			}
		}

		a.CosmosAccAddress = accAddress
		a.CosmosValAddress = cosmtypes.ValAddress(accAddress.Bytes())
	} else if err != nil {
		panic(errors.Wrapf(err, "failed to parse address: %s", a.Address))
	} else {
		panic(errors.Errorf("unsupported address: %s", a.Address))
	}
}
