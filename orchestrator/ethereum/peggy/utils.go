package peggy

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	wrappers "github.com/umee-network/peggo/solidity/wrappers/Peggy.sol"
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

	erc20Wrapper, err := wrappers.NewERC20(erc20ContractAddress, s.EVMCommitter.Provider())
	if err != nil {
		err = errors.Wrap(err, "failed to get ERC20 wrapper")
		return "", err
	}

	callOpts := &bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	}

	symbol, err = erc20Wrapper.Symbol(callOpts)
	if err != nil {
		err = errors.Wrap(err, "ERC20 [symbol] call failed")
		return "", err
	}

	return symbol, nil
}

func (s *peggyContract) GetERC20Decimals(
	ctx context.Context,
	erc20ContractAddress common.Address,
	callerAddress common.Address,
) (decimals uint8, err error) {
	erc20Wrapper, err := wrappers.NewERC20(erc20ContractAddress, s.EVMCommitter.Provider())
	if err != nil {
		err = errors.Wrap(err, "failed to get ERC20 wrapper")
		return 0, err
	}

	callOpts := &bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	}

	decimals, err = erc20Wrapper.Decimals(callOpts)
	if err != nil {
		err = errors.Wrap(err, "ERC20 'decimals' call failed")
		return 0, err
	}

	return decimals, nil
}
