package peggy

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// Gets the latest validator set nonce
func (s *peggyContract) GetValsetNonce(
	ctx context.Context,
	callerAddress common.Address,
) (*big.Int, error) {

	nonce, err := s.ethPeggy.StateLastValidatorSetNonce(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	})

	if err != nil {
		err = errors.Wrap(err, "StateLastValsetNonce call failed")
		return nil, err
	}

	return nonce, nil
}

// Gets the bridgeID
func (s *peggyContract) GetBridgeID(
	ctx context.Context,
	callerAddress common.Address,
) (common.Hash, error) {

	peggyID, err := s.ethPeggy.BRIDGEID(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	})

	if err != nil {
		err = errors.Wrap(err, "BRIDGEID call failed")
		return common.Hash{}, err
	}

	return peggyID, nil
}
