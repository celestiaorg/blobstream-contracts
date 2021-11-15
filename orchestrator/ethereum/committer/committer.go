package committer

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/umee-network/peggo/orchestrator/ethereum/provider"
)

// EVMCommitter defines an interface for submitting transactions
// into Ethereum, Matic, and other EVM-compatible networks.
type EVMCommitter interface {
	FromAddress() common.Address
	Provider() provider.EVMProvider
	SendTx(
		ctx context.Context,
		recipient common.Address,
		txData []byte,
	) (txHash common.Hash, err error)
}

type EVMCommitterOption func(o *options) error

type options struct {
	GasPrice   decimal.Decimal
	GasLimit   uint64
	RPCTimeout time.Duration
}

func defaultOptions() *options {
	v, _ := decimal.NewFromString("20")
	return &options{
		GasPrice:   v.Shift(9), // 20 gwei
		GasLimit:   1000000,
		RPCTimeout: 10 * time.Second,
	}
}

func applyOptions(o *options, opts ...EVMCommitterOption) error {
	for _, oo := range opts {
		if err := oo(o); err != nil {
			err = errors.Wrap(err, "failed to apply option to EVMCommitter")
			return err
		}
	}

	return nil
}

func OptionGasPriceFromString(str string) EVMCommitterOption {
	return func(o *options) error {
		gasPrice, err := decimal.NewFromString(str)
		if err != nil {
			err = errors.Wrap(err, "unable to parse gas price from string to decimal")
			return err
		}

		o.GasPrice = gasPrice
		return nil
	}
}

func OptionGasPriceFromDecimal(gasPrice decimal.Decimal) EVMCommitterOption {
	return func(o *options) error {
		o.GasPrice = gasPrice
		return nil
	}
}

func OptionGasPriceFromBigInt(i *big.Int) EVMCommitterOption {
	return func(o *options) error {
		o.GasPrice = decimal.NewFromBigInt(i, 0)
		return nil
	}
}

func OptionGasLimit(limit uint64) EVMCommitterOption {
	return func(o *options) error {
		o.GasLimit = limit
		return nil
	}
}

func TxBroadcastTimeout(dur time.Duration) EVMCommitterOption {
	return func(o *options) error {
		o.RPCTimeout = dur
		return nil
	}
}
