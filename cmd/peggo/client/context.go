package client

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/pkg/errors"

	keyscodec "github.com/InjectiveLabs/sdk-go/chain/crypto/codec"

	auction "github.com/InjectiveLabs/sdk-go/chain/auction/types"
	evm "github.com/InjectiveLabs/sdk-go/chain/evm/types"
	exchange "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	insurance "github.com/InjectiveLabs/sdk-go/chain/insurance/types"
	oracle "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	peggy "github.com/InjectiveLabs/sdk-go/chain/peggy/types"
	chaintypes "github.com/InjectiveLabs/sdk-go/chain/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramproposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcapplicationtypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	ibccoretypes "github.com/cosmos/ibc-go/modules/core/types"
)

// NewTxConfig initializes new Cosmos TxConfig with certain signModes enabled.
func NewTxConfig(signModes []signingtypes.SignMode) client.TxConfig {
	interfaceRegistry := types.NewInterfaceRegistry()
	keyscodec.RegisterInterfaces(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)
	exchange.RegisterInterfaces(interfaceRegistry)
	oracle.RegisterInterfaces(interfaceRegistry)
	insurance.RegisterInterfaces(interfaceRegistry)
	evm.RegisterInterfaces(interfaceRegistry)
	peggy.RegisterInterfaces(interfaceRegistry)
	chaintypes.RegisterInterfaces(interfaceRegistry)

	// more cosmos types
	authtypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	govtypes.RegisterInterfaces(interfaceRegistry)
	paramproposaltypes.RegisterInterfaces(interfaceRegistry)
	ibcapplicationtypes.RegisterInterfaces(interfaceRegistry)
	ibccoretypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)

	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return tx.NewTxConfig(marshaler, signModes)
}

// NewClientContext creates a new Cosmos Client context, where chainID
// corresponds to Cosmos chain ID, fromSpec is either name of the key, or bech32-address
// of the Cosmos account. Keyring is required to contain the specified key.
func NewClientContext(
	chainId, fromSpec string, kb keyring.Keyring,
) (client.Context, error) {
	clientCtx := client.Context{}

	interfaceRegistry := types.NewInterfaceRegistry()
	keyscodec.RegisterInterfaces(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)
	exchange.RegisterInterfaces(interfaceRegistry)
	insurance.RegisterInterfaces(interfaceRegistry)
	auction.RegisterInterfaces(interfaceRegistry)
	oracle.RegisterInterfaces(interfaceRegistry)
	evm.RegisterInterfaces(interfaceRegistry)
	peggy.RegisterInterfaces(interfaceRegistry)
	chaintypes.RegisterInterfaces(interfaceRegistry)

	// more cosmos types
	authtypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	govtypes.RegisterInterfaces(interfaceRegistry)
	paramproposaltypes.RegisterInterfaces(interfaceRegistry)
	ibcapplicationtypes.RegisterInterfaces(interfaceRegistry)
	ibccoretypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)

	marshaler := codec.NewProtoCodec(interfaceRegistry)
	encodingConfig := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig: NewTxConfig([]signingtypes.SignMode{
			signingtypes.SignMode_SIGN_MODE_DIRECT,
		}),
	}

	var keyInfo keyring.Info

	if kb != nil {
		addr, err := cosmostypes.AccAddressFromBech32(fromSpec)
		if err == nil {
			keyInfo, err = kb.KeyByAddress(addr)
			if err != nil {
				err = errors.Wrapf(err, "failed to load key info by address %s", addr.String())
				return clientCtx, err
			}
		} else {
			// failed to parse Bech32, is it a name?
			keyInfo, err = kb.Key(fromSpec)
			if err != nil {
				err = errors.Wrapf(err, "no key in keyring for name: %s", fromSpec)
				return clientCtx, err
			}
		}
	}

	clientCtx = newContext(
		chainId,
		encodingConfig,
		kb,
		keyInfo,
	)

	return clientCtx, nil
}

type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
}

func newContext(
	chainId string,
	encodingConfig EncodingConfig,
	kb keyring.Keyring,
	keyInfo keyring.Info,
) client.Context {
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

	return clientCtx
}
