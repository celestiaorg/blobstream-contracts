package client

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
	umeeapp "github.com/umee-network/umee/app"
)

// NewClientContext creates a new Cosmos Client context, where chainID
// corresponds to Cosmos chain ID, fromSpec is either name of the key, or bech32-address
// of the Cosmos account. Keyring is required to contain the specified key.
func NewClientContext(chainId, fromSpec string, kb keyring.Keyring) (client.Context, error) {
	var keyInfo keyring.Info

	if kb != nil {
		addr, err := cosmostypes.AccAddressFromBech32(fromSpec)
		if err == nil {
			keyInfo, err = kb.KeyByAddress(addr)
			if err != nil {
				err = errors.Wrapf(err, "failed to load key info by address %s", addr.String())
				return client.Context{}, err
			}
		} else {
			// failed to parse Bech32, is it a name?
			keyInfo, err = kb.Key(fromSpec)
			if err != nil {
				err = errors.Wrapf(err, "no key in keyring for name: %s", fromSpec)
				return client.Context{}, err
			}
		}
	}

	encodingConfig := umeeapp.MakeEncodingConfig()
	clientCtx := client.Context{
		ChainID:           chainId,
		JSONCodec:         encodingConfig.Marshaler,
		InterfaceRegistry: encodingConfig.InterfaceRegistry,
		Output:            os.Stderr,
		OutputFormat:      "json",
		BroadcastMode:     "block",
		UseLedger:         false,
		Simulate:          false,
		GenerateOnly:      false,
		Offline:           false,
		SkipConfirm:       true,
		TxConfig:          encodingConfig.TxConfig,
		AccountRetriever:  authtypes.AccountRetriever{},
	}

	if keyInfo != nil {
		clientCtx = clientCtx.WithKeyring(kb)
		clientCtx = clientCtx.WithFromAddress(keyInfo.GetAddress())
		clientCtx = clientCtx.WithFromName(keyInfo.GetName())
		clientCtx = clientCtx.WithFrom(keyInfo.GetName())
	}

	return clientCtx, nil
}
