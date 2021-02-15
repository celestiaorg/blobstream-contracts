package peggy

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

// Gets the latest transaction batch nonce
func (s *peggyContract) GetTxBatchNonce(
	ctx context.Context,
	erc20ContractAddress common.Address,
	callerAddress common.Address,
) (*big.Int, error) {

	nonce, err := s.ethPeggy.LastBatchNonce(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	}, erc20ContractAddress)

	if err != nil {
		err = errors.Wrap(err, "LastBatchNonce call failed")
		return nil, err
	}

	return nonce, nil
}

// Gets the latest validator set nonce
func (s *peggyContract) GetValsetNonce(
	ctx context.Context,
	callerAddress common.Address,
) (*big.Int, error) {

	nonce, err := s.ethPeggy.StateLastValsetNonce(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	})

	if err != nil {
		err = errors.Wrap(err, "StateLastValsetNonce call failed")
		return nil, err
	}

	return nonce, nil
}

// Gets the peggyID
func (s *peggyContract) GetPeggyID(
	ctx context.Context,
	callerAddress common.Address,
) (common.Hash, error) {

	peggyID, err := s.ethPeggy.StatePeggyId(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	})

	if err != nil {
		err = errors.Wrap(err, "StatePeggyId call failed")
		return common.Hash{}, err
	}

	return peggyID, nil
}

func (s *peggyContract) GetERC20Symbol(
	ctx context.Context,
	erc20ContractAddress common.Address,
	callerAddress common.Address,
) (symbol string, err error) {

	erc20Wrapper := bind.NewBoundContract(erc20ContractAddress, erc20ABI, s.ethProvider, nil, nil)

	callOpts := &bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	}
	var out []interface{}
	err = erc20Wrapper.Call(callOpts, &out, "symbol")
	if err != nil {
		err = errors.Wrap(err, "ERC20 [symbol] call failed")
		return "", err
	}

	symbol = *abi.ConvertType(out[0], new(string)).(*string)

	return symbol, nil
}
