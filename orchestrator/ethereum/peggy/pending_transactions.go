package peggy

import (
	"context"

	"time"

	log "github.com/xlab/suplog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type PendingTxInputList []string

func (p PendingTxInputList) AddPendingTxInput(txInput string) {
	// Enqueue pending tx input
	p = append(p, txInput)

	// Persisting top 100 pending txs of peggy contract only.
	if len(p) > 100 {
		// Dequeue pending tx input
		p[0] = "" // to avoid memory leak
		p = p[1:]
	}
}

func (p PendingTxInputList) IsPendingTxInput(txInput string) bool {
	for _, pendingTxInput := range p {
		if pendingTxInput == txInput {
			return true
		}
	}
	return false
}

func (s *peggyContract) SubscribeToPendingTxs(alchemyWebsocketURL string) {
	args := map[string]interface{}{
		"address": s.peggyAddress.Hex(),
	}

	wsClient, err := rpc.Dial(alchemyWebsocketURL)
	if err != nil {
		log.WithField("Websocket endpoint", alchemyWebsocketURL).WithError(err).Fatalln("Failed to connect to Ethereum Alchemy websocket")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Subscribe to Transactions
	ch := make(chan *RPCTransaction, 16)
	_, err = wsClient.EthSubscribe(ctx, ch, "alchemy_filteredNewFullPendingTransactions", args)
	log.WithField("Subscription error", alchemyWebsocketURL).WithError(err).Fatalln("Failed to subscribe to pending transactions")

	for {
		// Check that the transaction was send over the channel
		pendingTransaction := <-ch
		s.pendingTxInputList.AddPendingTxInput(string(pendingTransaction.Input))
	}
}

// RPCTransaction represents a transaction that will serialize to the RPC representation of a transaction
type RPCTransaction struct {
	BlockHash        *common.Hash         `json:"blockHash"`
	BlockNumber      *hexutil.Big         `json:"blockNumber"`
	From             common.Address       `json:"from"`
	Gas              hexutil.Uint64       `json:"gas"`
	GasPrice         *hexutil.Big         `json:"gasPrice"`
	GasFeeCap        *hexutil.Big         `json:"maxFeePerGas,omitempty"`
	GasTipCap        *hexutil.Big         `json:"maxPriorityFeePerGas,omitempty"`
	Hash             common.Hash          `json:"hash"`
	Input            hexutil.Bytes        `json:"input"`
	Nonce            hexutil.Uint64       `json:"nonce"`
	To               *common.Address      `json:"to"`
	TransactionIndex *hexutil.Uint64      `json:"transactionIndex"`
	Value            *hexutil.Big         `json:"value"`
	Type             hexutil.Uint64       `json:"type"`
	Accesses         *ethTypes.AccessList `json:"accessList,omitempty"`
	ChainID          *hexutil.Big         `json:"chainId,omitempty"`
	V                *hexutil.Big         `json:"v"`
	R                *hexutil.Big         `json:"r"`
	S                *hexutil.Big         `json:"s"`
}
