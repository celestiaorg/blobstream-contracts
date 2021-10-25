package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	cosmcrypto "github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/usbwallet"
	"github.com/ethereum/go-ethereum/common"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/umee-network/peggo/orchestrator/ethereum/keystore"
)

const defaultKeyringKeyName = "validator"

var emptyCosmosAddress = cosmtypes.AccAddress{}

func initCosmosKeyring(
	cosmosKeyringDir *string,
	cosmosKeyringAppName *string,
	cosmosKeyringBackend *string,
	cosmosKeyFrom *string,
	cosmosKeyPassphrase *string,
	cosmosPrivKey *string,
	cosmosUseLedger *bool,
) (cosmtypes.AccAddress, keyring.Keyring, error) {
	switch {
	case len(*cosmosPrivKey) > 0:
		if *cosmosUseLedger {
			err := errors.New("cannot combine ledger and privkey options")
			return emptyCosmosAddress, nil, err
		}

		pkBytes, err := hexToBytes(*cosmosPrivKey)
		if err != nil {
			err = errors.Wrap(err, "failed to hex-decode cosmos account privkey")
			return emptyCosmosAddress, nil, err
		}

		cosmosAccPk := &secp256k1.PrivKey{
			Key: pkBytes,
		}

		addressFromPk := cosmtypes.AccAddress(cosmosAccPk.PubKey().Address().Bytes())

		var keyName string

		// check that if cosmos 'From' specified separately, it must match the provided privkey,
		if len(*cosmosKeyFrom) > 0 {
			addressFrom, err := cosmtypes.AccAddressFromBech32(*cosmosKeyFrom)
			if err == nil {
				if !bytes.Equal(addressFrom.Bytes(), addressFromPk.Bytes()) {
					err = errors.Errorf("expected account address %s but got %s from the private key", addressFrom.String(), addressFromPk.String())
					return emptyCosmosAddress, nil, err
				}
			} else {
				// use it as a name then
				keyName = *cosmosKeyFrom
			}
		}

		if len(keyName) == 0 {
			keyName = defaultKeyringKeyName
		}

		// wrap a PK into a Keyring
		kb, err := KeyringForPrivKey(keyName, cosmosAccPk)
		return addressFromPk, kb, err

	case len(*cosmosKeyFrom) > 0:
		var fromIsAddress bool
		addressFrom, err := cosmtypes.AccAddressFromBech32(*cosmosKeyFrom)
		if err == nil {
			fromIsAddress = true
		}

		var passReader io.Reader = os.Stdin
		if len(*cosmosKeyPassphrase) > 0 {
			passReader = newPassReader(*cosmosKeyPassphrase)
		}

		var absoluteKeyringDir string
		if filepath.IsAbs(*cosmosKeyringDir) {
			absoluteKeyringDir = *cosmosKeyringDir
		} else {
			absoluteKeyringDir, _ = filepath.Abs(*cosmosKeyringDir)
		}

		kb, err := keyring.New(
			*cosmosKeyringAppName,
			*cosmosKeyringBackend,
			absoluteKeyringDir,
			passReader,
			hd.Secp256k1Option(),
		)
		if err != nil {
			err = errors.Wrap(err, "failed to init keyring")
			return emptyCosmosAddress, nil, err
		}

		var keyInfo keyring.Info
		if fromIsAddress {
			if keyInfo, err = kb.KeyByAddress(addressFrom); err != nil {
				err = errors.Wrapf(err, "couldn't find an entry for the key %s in keybase", addressFrom.String())
				return emptyCosmosAddress, nil, err
			}
		} else {
			if keyInfo, err = kb.Key(*cosmosKeyFrom); err != nil {
				err = errors.Wrapf(err, "could not find an entry for the key '%s' in keybase", *cosmosKeyFrom)
				return emptyCosmosAddress, nil, err
			}
		}

		switch keyType := keyInfo.GetType(); keyType {
		case keyring.TypeLocal:
			// kb has a key and it's totally usable
			return keyInfo.GetAddress(), kb, nil
		case keyring.TypeLedger:
			// the kb stores references to ledger keys, so we must explicitly
			// check that. kb doesn't know how to scan HD keys - they must be added manually before
			if *cosmosUseLedger {
				return keyInfo.GetAddress(), kb, nil
			}
			err := errors.Errorf("'%s' key is a ledger reference, enable ledger option", keyInfo.GetName())
			return emptyCosmosAddress, nil, err
		case keyring.TypeOffline:
			err := errors.Errorf("'%s' key is an offline key, not supported yet", keyInfo.GetName())
			return emptyCosmosAddress, nil, err
		case keyring.TypeMulti:
			err := errors.Errorf("'%s' key is an multisig key, not supported yet", keyInfo.GetName())
			return emptyCosmosAddress, nil, err
		default:
			err := errors.Errorf("'%s' key  has unsupported type: %s", keyInfo.GetName(), keyType)
			return emptyCosmosAddress, nil, err
		}

	default:
		err := errors.New("insufficient cosmos key details provided")
		return emptyCosmosAddress, nil, err
	}
}

var emptyEthAddress = ethcmn.Address{}

func initEthereumAccountsManager(
	ethChainID uint64,
	ethKeystoreDir *string,
	ethKeyFrom *string,
	ethPassphrase *string,
	ethPrivKey *string,
	ethUseLedger *bool,
) (
	ethKeyFromAddress ethcmn.Address,
	signerFn bind.SignerFn,
	personalSignFn keystore.PersonalSignFn,
	err error,
) {
	switch {
	case *ethUseLedger:
		if ethKeyFrom == nil {
			err := errors.New("cannot use Ledger without from address specified")
			return emptyEthAddress, nil, nil, err
		}

		ethKeyFromAddress = ethcmn.HexToAddress(*ethKeyFrom)
		if ethKeyFromAddress == (ethcmn.Address{}) {
			err = errors.Wrap(err, "failed to parse Ethereum from address")
			return emptyEthAddress, nil, nil, err
		}

		ledgerBackend, err := usbwallet.NewLedgerHub()
		if err != nil {
			err = errors.Wrap(err, "failed to connect with Ethereum app on Ledger device")
			return emptyEthAddress, nil, nil, err
		}

		signerFn = func(from common.Address, tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
			acc := accounts.Account{
				Address: from,
			}

			wallets := ledgerBackend.Wallets()
			for _, w := range wallets {
				if err := w.Open(""); err != nil {
					err = errors.Wrap(err, "failed to connect to wallet on Ledger device")
					return nil, err
				}

				if !w.Contains(acc) {
					if err := w.Close(); err != nil {
						err = errors.Wrap(err, "failed to disconnect the wallet on Ledger device")
						return nil, err
					}

					continue
				}

				tx, err = w.SignTx(acc, tx, new(big.Int).SetUint64(ethChainID))
				_ = w.Close()
				return tx, err
			}

			return nil, errors.Errorf("account %s not found on Ledger", from.String())
		}

		personalSignFn = func(from common.Address, data []byte) (sig []byte, err error) {
			acc := accounts.Account{
				Address: from,
			}

			wallets := ledgerBackend.Wallets()
			for _, w := range wallets {
				if err := w.Open(""); err != nil {
					err = errors.Wrap(err, "failed to connect to wallet on Ledger device")
					return nil, err
				}

				if !w.Contains(acc) {
					if err := w.Close(); err != nil {
						err = errors.Wrap(err, "failed to disconnect the wallet on Ledger device")
						return nil, err
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

	case len(*ethPrivKey) > 0:
		ethPk, err := crypto.HexToECDSA(*ethPrivKey)
		if err != nil {
			err = errors.Wrap(err, "failed to hex-decode Ethereum ECDSA Private Key")
			return emptyEthAddress, nil, nil, err
		}

		ethAddressFromPk := ethcrypto.PubkeyToAddress(ethPk.PublicKey)

		if len(*ethKeyFrom) > 0 {
			addr := ethcmn.HexToAddress(*ethKeyFrom)
			if addr == (ethcmn.Address{}) {
				err = errors.Wrap(err, "failed to parse Ethereum from address")
				return emptyEthAddress, nil, nil, err
			} else if addr != ethAddressFromPk {
				err = errors.Wrap(err, "Ethereum from address does not match address from ECDSA Private Key")
				return emptyEthAddress, nil, nil, err
			}
		}

		txOpts, err := bind.NewKeyedTransactorWithChainID(ethPk, new(big.Int).SetUint64(ethChainID))
		if err != nil {
			err = errors.New("failed to init NewKeyedTransactorWithChainID")
			return emptyEthAddress, nil, nil, err
		}

		personalSignFn, err := keystore.PrivateKeyPersonalSignFn(ethPk)
		if err != nil {
			err = errors.New("failed to init PrivateKeyPersonalSignFn")
			return emptyEthAddress, nil, nil, err
		}

		return txOpts.From, txOpts.Signer, personalSignFn, nil

	case len(*ethKeystoreDir) > 0:
		if ethKeyFrom == nil {
			err := errors.New("cannot use Ethereum keystore without from address specified")
			return emptyEthAddress, nil, nil, err
		}

		ethKeyFromAddress = ethcmn.HexToAddress(*ethKeyFrom)
		if ethKeyFromAddress == (ethcmn.Address{}) {
			err = errors.Wrap(err, "failed to parse Ethereum from address")
			return emptyEthAddress, nil, nil, err
		}

		if info, err := os.Stat(*ethKeystoreDir); err != nil || !info.IsDir() {
			err = errors.New("failed to locate keystore dir")
			return emptyEthAddress, nil, nil, err
		}

		ks, err := keystore.New(*ethKeystoreDir)
		if err != nil {
			err = errors.Wrap(err, "failed to load keystore")
			return emptyEthAddress, nil, nil, err
		}

		var pass string
		if len(*ethPassphrase) > 0 {
			pass = *ethPassphrase
		} else {
			pass, err = ethPassFromStdin()
			if err != nil {
				return emptyEthAddress, nil, nil, err
			}
		}

		signerFn, err := ks.SignerFn(ethChainID, ethKeyFromAddress, pass)
		if err != nil {
			err = errors.Wrapf(err, "failed to load key for %s", ethKeyFromAddress)
			return emptyEthAddress, nil, nil, err
		}

		personalSignFn, err := ks.PersonalSignFn(ethKeyFromAddress, pass)
		if err != nil {
			err = errors.Wrapf(err, "failed to load key for %s", ethKeyFromAddress)
			return emptyEthAddress, nil, nil, err
		}

		return ethKeyFromAddress, signerFn, personalSignFn, nil

	default:
		err := errors.New("insufficient ethereum key details provided")
		return emptyEthAddress, nil, nil, err
	}
}

func ethPassFromStdin() (string, error) {
	fmt.Print("Passphrase for Ethereum account: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		err := errors.Wrap(err, "failed to read password from stdin")
		return "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(password), nil
}

func newPassReader(pass string) io.Reader {
	return &passReader{
		pass: pass,
		buf:  new(bytes.Buffer),
	}
}

type passReader struct {
	pass string
	buf  *bytes.Buffer
}

var _ io.Reader = &passReader{}

func (r *passReader) Read(p []byte) (n int, err error) {
	n, err = r.buf.Read(p)
	if err == io.EOF || n == 0 {
		r.buf.WriteString(r.pass + "\n")

		n, err = r.buf.Read(p)
	}

	return
}

// KeyringForPrivKey creates a temporary in-mem keyring for a PrivKey.
// Allows to init Context when the key has been provided in plaintext and parsed.
func KeyringForPrivKey(name string, privKey cryptotypes.PrivKey) (keyring.Keyring, error) {
	kb := keyring.NewInMemory(hd.EthSecp256k1Option())
	tmpPhrase := randPhrase(64)
	armored := cosmcrypto.EncryptArmorPrivKey(privKey, tmpPhrase, privKey.Type())
	err := kb.ImportPrivKey(name, armored, tmpPhrase)
	if err != nil {
		err = errors.Wrap(err, "failed to import privkey")
		return nil, err
	}

	return kb, nil
}

func randPhrase(size int) string {
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	orPanic(err)

	return string(buf)
}

func orPanic(err error) {
	if err != nil {
		log.Panicln()
	}
}
