package solidity

import (
	"bufio"
	"bytes"
	"crypto/ecdsa"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func init() {
	os.Setenv("SOLIDITY_TEST_EVM_RPC", "http://localhost:8545")

	readEnv()

	for idx, a := range EthAccounts {
		a.Parse()
		EthAccounts[idx] = a
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

type Account struct {
	Address string
	Key     string

	EthAddress common.Address
	EthPrivKey *ecdsa.PrivateKey
}

func (a *Account) Parse() {
	pk, err := crypto.HexToECDSA(a.Key[2:])
	orFail(err)

	a.EthAddress = common.HexToAddress(a.Address)
	a.EthPrivKey = pk
}
