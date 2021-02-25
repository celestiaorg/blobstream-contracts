package provider

import (
	"context"
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

type EVMProvider interface {
	bind.ContractCaller
	bind.ContractFilterer

	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

type EVMProviderWithRet interface {
	EVMProvider

	SendTransactionWithRet(ctx context.Context, tx *types.Transaction) (txHash common.Hash, err error)
}

type evmProviderWithRet struct {
	*ethclient.Client

	rc *rpc.Client
}

func NewEVMProvider(rc *rpc.Client) EVMProviderWithRet {
	return &evmProviderWithRet{
		Client: ethclient.NewClient(rc),
		rc:     rc,
	}
}

func (p *evmProviderWithRet) SendTransactionWithRet(ctx context.Context, tx *types.Transaction) (txHash common.Hash, err error) {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return common.Hash{}, err
	}

	if err := p.rc.CallContext(ctx, &txHash, "eth_sendRawTransaction", hexutil.Encode(data)); err != nil {
		return common.Hash{}, err
	}

	return txHash, nil
}

type TransactFunc func(opts *bind.TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error)

func TransactFn(p EVMProviderWithRet, contractAddress common.Address, txHashOut *common.Hash) TransactFunc {
	return func(opts *bind.TransactOpts, contract *common.Address, input []byte) (*types.Transaction, error) {
		var err error

		// Ensure a valid value field and resolve the account nonce
		value := opts.Value
		if value == nil {
			value = new(big.Int)
		}

		var nonce uint64
		if opts.Nonce == nil {
			nonce, err = p.PendingNonceAt(opts.Context, opts.From)
			if err != nil {
				return nil, errors.Errorf("failed to retrieve account nonce: %v", err)
			}
		} else {
			nonce = opts.Nonce.Uint64()
		}

		// Figure out the gas allowance and gas price values
		gasPrice := opts.GasPrice
		if gasPrice == nil {
			gasPrice, err = p.SuggestGasPrice(opts.Context)
			if err != nil {
				return nil, errors.Errorf("failed to suggest gas price: %v", err)
			}
		}

		gasLimit := opts.GasLimit
		if gasLimit == 0 {
			// Gas estimation cannot succeed without code for method invocations
			if contract != nil {
				if code, err := p.PendingCodeAt(opts.Context, contractAddress); err != nil {
					return nil, err
				} else if len(code) == 0 {
					return nil, bind.ErrNoCode
				}
			}
			// If the contract surely has code (or code is not needed), estimate the transaction
			msg := ethereum.CallMsg{From: opts.From, To: contract, GasPrice: gasPrice, Value: value, Data: input}
			gasLimit, err = p.EstimateGas(opts.Context, msg)
			if err != nil {
				return nil, errors.Errorf("failed to estimate gas needed: %v", err)
			}
		}
		// Create the transaction, sign it and schedule it for execution
		var rawTx *types.Transaction
		if contract == nil {
			rawTx = types.NewContractCreation(nonce, value, gasLimit, gasPrice, input)
		} else {
			rawTx = types.NewTransaction(nonce, contractAddress, value, gasLimit, gasPrice, input)
		}
		if opts.Signer == nil {
			return nil, errors.New("no signer to authorize the transaction with")
		}

		signedTx, err := opts.Signer(opts.From, rawTx)
		if err != nil {
			return nil, err
		}

		txHash, err := p.SendTransactionWithRet(opts.Context, signedTx)
		if err != nil {
			return nil, err
		}

		*txHashOut = txHash
		return signedTx, nil
	}
}
