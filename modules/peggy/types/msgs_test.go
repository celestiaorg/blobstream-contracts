package types

import (
	"bytes"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestValidateMsgSetOrchestratorAddresses(t *testing.T) {
	var (
		ethAddress                   = common.HexToAddress("0xb462864E395d88d6bc7C5dd5F3F5eb4cc2599255")
		cosmosAddress sdk.AccAddress = bytes.Repeat([]byte{0x1}, sdk.AddrLen)
		valAddress    sdk.AccAddress = bytes.Repeat([]byte{0x1}, sdk.AddrLen)
	)
	specs := map[string]struct {
		srcCosmosAddr sdk.AccAddress
		srcValAddr    sdk.AccAddress
		srcETHAddr    common.Address
		expErr        bool
	}{
		"all good": {
			srcCosmosAddr: cosmosAddress,
			srcValAddr:    valAddress,
			srcETHAddr:    ethAddress,
		},
		"empty validator address": {
			srcETHAddr:    ethAddress,
			srcCosmosAddr: cosmosAddress,
			expErr:        true,
		},
		"invalid validator address": {
			srcValAddr:    []byte{0x1},
			srcCosmosAddr: cosmosAddress,
			srcETHAddr:    ethAddress,
			expErr:        true,
		},
		"empty cosmos address": {
			srcValAddr: valAddress,
			srcETHAddr: ethAddress,
			expErr:     true,
		},
		"invalid cosmos address": {
			srcCosmosAddr: []byte{0x1},
			srcValAddr:    valAddress,
			srcETHAddr:    ethAddress,
			expErr:        true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			msg := NewMsgSetOrchestratorAddress(spec.srcValAddr, spec.srcCosmosAddr, spec.srcETHAddr)
			// when
			err := msg.ValidateBasic()
			if spec.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}

}
