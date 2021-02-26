package tmclient

import (
	"context"
	"strings"

	log "github.com/xlab/suplog"

	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type TendermintClient interface {
	GetBlock(ctx context.Context, height int64) (*tmctypes.ResultBlock, error)
	GetLatestBlockHeight(ctx context.Context) (int64, error)
	GetTxs(ctx context.Context, block *tmctypes.ResultBlock) ([]*ctypes.ResultTx, error)
	GetValidatorSet(ctx context.Context, height int64) (*tmctypes.ResultValidators, error)
}

type tmClient struct {
	rpcClient rpcclient.Client
}

func NewRPCClient(rpcNodeAddr string) TendermintClient {
	rpcClient, err := rpchttp.NewWithTimeout(rpcNodeAddr, "/websocket", 10)
	if err != nil {
		log.WithError(err).Fatalln("failed to init rpcClient")
	}

	return &tmClient{
		rpcClient: rpcClient,
	}
}

// GetBlock queries for a block by height. An error is returned if the query fails.
func (c *tmClient) GetBlock(ctx context.Context, height int64) (*tmctypes.ResultBlock, error) {
	return c.rpcClient.Block(ctx, &height)
}

// GetLatestBlockHeight returns the latest block height on the active chain.
func (c *tmClient) GetLatestBlockHeight(ctx context.Context) (int64, error) {
	status, err := c.rpcClient.Status(ctx)
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight

	return height, nil
}

// GetTxs queries for all the transactions in a block height.
// It uses `Tx` RPC method to query for the transaction.
func (c *tmClient) GetTxs(ctx context.Context, block *tmctypes.ResultBlock) ([]*ctypes.ResultTx, error) {
	txs := make([]*ctypes.ResultTx, 0, len(block.Block.Txs))

	for _, tmTx := range block.Block.Txs {
		tx, err := c.rpcClient.Tx(ctx, tmTx.Hash(), true)
		if err != nil {
			if strings.HasSuffix(err.Error(), "not found") {
				log.WithError(err).Errorln("failed to get Tx by hash")
				continue
			}

			return nil, err
		}

		txs = append(txs, tx)
	}

	return txs, nil
}

// GetValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (c *tmClient) GetValidatorSet(ctx context.Context, height int64) (*tmctypes.ResultValidators, error) {
	return c.rpcClient.Validators(ctx, &height, nil, nil)
}
