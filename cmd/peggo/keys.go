package peggo

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	sdkcrypto "github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdkcryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/usbwallet"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/knadh/koanf"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/umee-network/peggo/orchestrator/ethereum/keystore"
	"golang.org/x/term"
)

const defaultKeyringKeyName = "validator"

var (
	emptyCosmosAddress = sdk.AccAddress{}
	emptyEthAddress    = ethcmn.Address{}
)

func initCosmosKeyring(konfig *koanf.Koanf) (sdk.AccAddress, keyring.Keyring, error) {
	cosmosFrom := konfig.String(flagCosmosFrom)
	cosmosPK := konfig.String(flagCosmosPK)
	cosmosPassphrase := konfig.String(flagCosmosFromPassphrase)
	cosmosKeyringDir := konfig.String(flagCosmosKeyringDir)
	cosmosUseLedger := konfig.Bool(flagCosmosUseLedger)

	switch {
	case len(cosmosPK) > 0:
		if cosmosUseLedger {
			return emptyCosmosAddress, nil, errors.New("cannot use ledger with raw private key")
		}

		pkBz, err := hexToBytes(cosmosPK)
		if err != nil {
			return emptyCosmosAddress, nil, fmt.Errorf("failed to hex decode cosmos private key: %w", err)
		}

		cosmosAccPk := &secp256k1.PrivKey{
			Key: pkBz,
		}

		addressFromPk := sdk.AccAddress(cosmosAccPk.PubKey().Address().Bytes())

		// Check that if cosmos 'From' specified separately, it must match the
		// provided privkey.
		var keyName string

		if len(cosmosFrom) > 0 {
			addressFrom, err := sdk.AccAddressFromBech32(cosmosFrom)
			if err == nil {
				if !bytes.Equal(addressFrom.Bytes(), addressFromPk.Bytes()) {
					return emptyCosmosAddress, nil, fmt.Errorf(
						"expected account address %s but got %s from the private key",
						addressFrom.String(), addressFromPk.String(),
					)
				}
			} else {
				// use it as a name then
				keyName = cosmosFrom
			}
		}

		if len(keyName) == 0 {
			keyName = defaultKeyringKeyName
		}

		// wrap a PK into a Keyring
		kb, err := keyringForPrivKey(keyName, cosmosAccPk)
		return addressFromPk, kb, err

	case len(cosmosFrom) > 0:
		var fromIsAddress bool
		addressFrom, err := sdk.AccAddressFromBech32(cosmosFrom)
		if err == nil {
			fromIsAddress = true
		}

		var passReader io.Reader
		if len(cosmosPassphrase) > 0 {
			passReader = newPassReader(cosmosPassphrase)
		} else {
			passReader = os.Stdin
		}

		var absoluteKeyringDir string
		if filepath.IsAbs(cosmosKeyringDir) {
			absoluteKeyringDir = cosmosKeyringDir
		} else {
			absoluteKeyringDir, err = filepath.Abs(cosmosKeyringDir)
			if err != nil {
				return emptyCosmosAddress, nil, err
			}
		}

		kb, err := keyring.New(
			konfig.String(flagCosmosKeyringApp),
			konfig.String(flagCosmosKeyring),
			absoluteKeyringDir,
			passReader,
		)
		if err != nil {
			return emptyCosmosAddress, nil, fmt.Errorf("failed to create keyring: %w", err)
		}

		var keyInfo keyring.Info
		if fromIsAddress {
			if keyInfo, err = kb.KeyByAddress(addressFrom); err != nil {
				return emptyCosmosAddress, nil, fmt.Errorf(
					"failed to find an entry for the key %s in the keyring: %w",
					addressFrom.String(), err,
				)
			}
		} else {
			if keyInfo, err = kb.Key(cosmosFrom); err != nil {
				return emptyCosmosAddress, nil, fmt.Errorf(
					"failed to find an entry for the key %s in the keyring: %w",
					cosmosFrom, err,
				)
			}
		}

		switch keyType := keyInfo.GetType(); keyType {
		case keyring.TypeLocal:
			// kb has a key and it's totally usable
			return keyInfo.GetAddress(), kb, nil

		case keyring.TypeLedger:
			// The keyring stores references to ledger keys, so we must explicitly
			// check that. The keyring doesn't know how to scan HD keys - they must be
			// added manually before.
			if cosmosUseLedger {
				return keyInfo.GetAddress(), kb, nil
			}

			return emptyCosmosAddress, nil, fmt.Errorf("'%s' key is a ledger reference, enable ledger option", keyInfo.GetName())

		case keyring.TypeOffline:
			return emptyCosmosAddress, nil, fmt.Errorf("'%s' key is an offline key, not supported yet", keyInfo.GetName())

		case keyring.TypeMulti:
			return emptyCosmosAddress, nil, fmt.Errorf("'%s' key is an multisig key, not supported yet", keyInfo.GetName())

		default:
			return emptyCosmosAddress, nil, fmt.Errorf("'%s' key  has unsupported type: %s", keyInfo.GetName(), keyType)
		}

	default:
		return emptyCosmosAddress, nil, errors.New("insufficient cosmos key details provided")
	}
}

func initEthereumAccountsManager(
	logger zerolog.Logger,
	ethChainID uint64,
	konfig *koanf.Koanf,
) (
	ethcmn.Address,
	bind.SignerFn,
	keystore.PersonalSignFn,
	error,
) {
	var (
		signerFn          bind.SignerFn
		ethKeyFromAddress ethcmn.Address
		personalSignFn    keystore.PersonalSignFn
	)

	ethUseLedger := konfig.Bool(flagEthUseLedger)
	ethKeyFrom := konfig.String(flagEthFrom)
	ethPrivKey := konfig.String(flagEthPK)
	ethKeystoreDir := konfig.String(flagEthKeystoreDir)
	ethPassphrase := konfig.String(flagEthPassphrase)

	switch {
	case ethUseLedger:
		if len(ethKeyFrom) == 0 {
			return emptyEthAddress, nil, nil, errors.New("cannot use Ledger without from address specified")
		}

		ethKeyFromAddress = ethcmn.HexToAddress(ethKeyFrom)
		if ethKeyFromAddress == (ethcmn.Address{}) {
			return emptyEthAddress, nil, nil, fmt.Errorf("failed to parse Ethereum from address %s", ethKeyFrom)
		}

		ledgerBackend, err := usbwallet.NewLedgerHub()
		if err != nil {
			return emptyEthAddress, nil, nil, fmt.Errorf("failed to connect with Ethereum app on Ledger device")
		}

		signerFn = func(from ethcmn.Address, tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
			acc := accounts.Account{
				Address: from,
			}

			wallets := ledgerBackend.Wallets()
			for _, w := range wallets {
				if err := w.Open(""); err != nil {
					return nil, fmt.Errorf("failed to connect to wallet on Ledger device: %w", err)
				}

				if !w.Contains(acc) {
					if err := w.Close(); err != nil {
						return nil, fmt.Errorf("failedt to disconnect the wallet on Ledger device: %w", err)
					}

					continue
				}

				tx, err = w.SignTx(acc, tx, new(big.Int).SetUint64(ethChainID))
				_ = w.Close()
				return tx, err
			}

			return nil, errors.Errorf("account %s not found on Ledger", from.String())
		}

		personalSignFn = func(from ethcmn.Address, data []byte) (sig []byte, err error) {
			acc := accounts.Account{
				Address: from,
			}

			wallets := ledgerBackend.Wallets()
			for _, w := range wallets {
				if err := w.Open(""); err != nil {
					return nil, fmt.Errorf("failed to connect to wallet on Ledger device: %w", err)
				}

				if !w.Contains(acc) {
					if err := w.Close(); err != nil {
						return nil, fmt.Errorf("failedt to disconnect the wallet on Ledger device: %w", err)
					}

					continue
				}

				sig, err = w.SignText(acc, data)
				_ = w.Close()
				return sig, err
			}

			return nil, errors.Errorf("account %s not found on Ledger", from.String())
		}

		return ethKeyFromAddress, signerFn, personalSignFn, nil

	case len(ethPrivKey) > 0:
		ethPk, err := ethcrypto.HexToECDSA(ethPrivKey)
		if err != nil {
			return emptyEthAddress, nil, nil, fmt.Errorf("failed to hex-decode Ethereum ECDSA Private Key: %w", err)
		}

		ethAddressFromPk := ethcrypto.PubkeyToAddress(ethPk.PublicKey)

		if len(ethKeyFrom) > 0 {
			addr := ethcmn.HexToAddress(ethKeyFrom)
			if addr == (ethcmn.Address{}) {
				return emptyEthAddress, nil, nil, fmt.Errorf("failed to parse Ethereum from address: %s", ethKeyFrom)
			} else if addr != ethAddressFromPk {
				return emptyEthAddress, nil, nil, errors.New("from address does not match address from Ethereum ECDSA private key")
			}
		}

		txOpts, err := bind.NewKeyedTransactorWithChainID(ethPk, new(big.Int).SetUint64(ethChainID))
		if err != nil {
			return emptyEthAddress, nil, nil, fmt.Errorf("failed to init NewKeyedTransactorWithChainID: %w", err)
		}

		personalSignFn, err := keystore.PrivateKeyPersonalSignFn(ethPk)
		if err != nil {
			return emptyEthAddress, nil, nil, fmt.Errorf("failed to init PrivateKeyPersonalSignFn: %w", err)
		}

		return txOpts.From, txOpts.Signer, personalSignFn, nil

	case len(ethKeystoreDir) > 0:
		if len(ethKeyFrom) == 0 {
			return emptyEthAddress, nil, nil, errors.New("cannot use Ethereum keystore without from address specified")
		}

		ethKeyFromAddress = ethcmn.HexToAddress(ethKeyFrom)
		if ethKeyFromAddress == (ethcmn.Address{}) {
			return emptyEthAddress, nil, nil, fmt.Errorf("failed to parse Ethereum from address: %s", ethKeyFrom)
		}

		if info, err := os.Stat(ethKeystoreDir); err != nil || !info.IsDir() {
			return emptyEthAddress, nil, nil, fmt.Errorf("failed to locate Ethereum keystore dir: %w", err)
		}

		ks, err := keystore.New(logger, ethKeystoreDir)
		if err != nil {
			return emptyEthAddress, nil, nil, fmt.Errorf("failed to load Ethereum keystore: %w", err)
		}

		var pass string
		if len(ethPassphrase) > 0 {
			pass = ethPassphrase
		} else {
			pass, err = ethPassFromStdin()
			if err != nil {
				return emptyEthAddress, nil, nil, err
			}
		}

		signerFn, err := ks.SignerFn(ethChainID, ethKeyFromAddress, pass)
		if err != nil {
			return emptyEthAddress, nil, nil, fmt.Errorf("failed to load key for %s: %w", ethKeyFromAddress, err)
		}

		personalSignFn, err := ks.PersonalSignFn(ethKeyFromAddress, pass)
		if err != nil {
			return emptyEthAddress, nil, nil, fmt.Errorf("failed to load key for %s: %w", ethKeyFromAddress, err)
		}

		return ethKeyFromAddress, signerFn, personalSignFn, nil

	default:
		return emptyEthAddress, nil, nil, errors.New("insufficient ethereum key details provided")
	}
}

func ethPassFromStdin() (string, error) {
	fmt.Fprintln(os.Stderr, "Passphrase for Ethereum account: ")
	bytePassword, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", fmt.Errorf("failed to read password from STDIN: %w", err)
	}

	password := string(bytePassword)
	return strings.TrimSpace(password), nil
}

var _ io.Reader = (*passReader)(nil)

type passReader struct {
	pass string
	buf  *bytes.Buffer
}

func newPassReader(pass string) io.Reader {
	return &passReader{
		pass: pass,
		buf:  new(bytes.Buffer),
	}
}

func (r *passReader) Read(p []byte) (n int, err error) {
	n, err = r.buf.Read(p)
	if err == io.EOF || n == 0 {
		r.buf.WriteString(r.pass + "\n")

		n, err = r.buf.Read(p)
	}

	return
}

// keyringForPrivKey creates a temporary in-mem keyring for a PrivKey.
// Allows to init Context when the key has been provided in plaintext and parsed.
func keyringForPrivKey(name string, privKey sdkcryptotypes.PrivKey) (keyring.Keyring, error) {
	tmpPhrase, err := randPhrase(64)
	if err != nil {
		return nil, err
	}

	armored := sdkcrypto.EncryptArmorPrivKey(privKey, tmpPhrase, privKey.Type())

	kb := keyring.NewInMemory()
	if err := kb.ImportPrivKey(name, armored, tmpPhrase); err != nil {
		err = errors.Wrap(err, "failed to import privkey")
		return nil, err
	}

	return kb, nil
}

func randPhrase(size int) (string, error) {
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
