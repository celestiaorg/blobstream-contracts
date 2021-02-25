package client

import (
	"encoding/json"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/pkg/errors"
	abci "github.com/tendermint/tendermint/abci/types"
	log "github.com/xlab/suplog"
	"google.golang.org/grpc"

	ctypes "github.com/InjectiveLabs/sdk-go/chain/types"
)

func init() {
	// set the address prefixes
	config := sdk.GetConfig()

	// This is specific to Injective chain
	ctypes.SetBech32Prefixes(config)
	ctypes.SetBip44CoinType(config)
}

type CosmosClient interface {
	CanSignTransactions() bool
	FromAddress() sdk.AccAddress
	QueryClient() *grpc.ClientConn
	SyncBroadcastMsg(msgs ...sdk.Msg) (*sdk.TxResponse, error)
	QueueBroadcastMsg(msgs ...sdk.Msg) error
	ClientContext() client.Context
	Close()
}

// NewCosmosClient creates a new gRPC client that communicates with gRPC server at protoAddr.
// protoAddr must be in form "tcp://127.0.0.1:8080" or "unix:///tmp/test.sock", protocol is required.
func NewCosmosClient(ctx client.Context, protoAddr string) (CosmosClient, error) {
	conn, err := grpc.Dial(protoAddr, grpc.WithInsecure(), grpc.WithContextDialer(dialerFunc))
	if err != nil {
		err := errors.Wrapf(err, "failed to connect to the gRPC: %s", protoAddr)
		return nil, err
	}

	cc := &cosmosClient{
		ctx: ctx,
		logger: log.WithFields(log.Fields{
			"module": "peggo",
			"svc":    "cosmosClient",
		}),
		conn:      conn,
		txFactory: NewTxFactory(ctx),
		canSign:   ctx.Keyring != nil,
		syncMux:   new(sync.Mutex),
		msgC:      make(chan sdk.Msg, msgCommitBatchSizeLimit),
	}

	if cc.canSign {
		var err error

		cc.accNum, cc.accSeq, err = cc.txFactory.AccountRetriever().GetAccountNumberSequence(ctx, ctx.GetFromAddress())
		if err != nil {
			err = errors.Wrap(err, "failed to get initial account num and seq")
			return nil, err
		}

		go cc.runBatchBroadcast()
	}

	return cc, nil
}

func (c *cosmosClient) syncNonce() {
	num, seq, err := c.txFactory.AccountRetriever().GetAccountNumberSequence(c.ctx, c.ctx.GetFromAddress())
	if err != nil {
		c.logger.WithError(err).Errorln("failed to get account seq")
		return
	} else if num != c.accNum {
		c.logger.WithFields(log.Fields{
			"expected": c.accNum,
			"actual":   num,
		}).Panic("account number changed during nonce sync")
	}

	c.accSeq = seq
}

type cosmosClient struct {
	ctx       client.Context
	logger    log.Logger
	conn      *grpc.ClientConn
	txFactory tx.Factory

	fromAddress sdk.AccAddress
	doneC       chan bool
	msgC        chan sdk.Msg
	syncMux     *sync.Mutex

	accNum uint64
	accSeq uint64

	closed  int64
	canSign bool
}

func (c *cosmosClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

func (c *cosmosClient) ClientContext() client.Context {
	return c.ctx
}

func (c *cosmosClient) CanSignTransactions() bool {
	return c.canSign
}

func (c *cosmosClient) FromAddress() sdk.AccAddress {
	if !c.canSign {
		return sdk.AccAddress{}
	}

	return c.ctx.FromAddress
}

var (
	ErrQueueClosed    = errors.New("queue is closed")
	ErrEnqueueTimeout = errors.New("enqueue timeout")
	ErrReadOnly       = errors.New("client is in read-only mode")
)

func (c *cosmosClient) SyncBroadcastMsg(msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	res, err := c.broadcastTx(c.ctx, c.txFactory, true, msgs...)
	if err != nil {
		c.logger.WithField("size", 1).WithError(err).Errorln("failed to commit msg batch")
		return nil, err
	}

	c.accSeq++

	return res, nil
}

func (c *cosmosClient) broadcastTx(
	clientCtx client.Context,
	txf tx.Factory,
	await bool,
	msgs ...sdk.Msg,
) (*sdk.TxResponse, error) {

	txf, err := tx.PrepareFactory(clientCtx, txf)
	if err != nil {
		return nil, err
	}

	/*
		CONTEXT

		submitTx flow:
		txf.SimulateAndExecute() || clientCtx.Simulate
			CalculateGas
			txf.WithGas
			...

		the CalculateGas function call into the core to simulate and calculate the gas there
		but it's broken in 0.40.1, look like work in progress

		We attempt to simulate ourselves using abci.RequestQuery
		unlike simulation inside the core, this somehow requires txf must have sufficient gas at beginning

		So we can just input the big number of gas for the txf and simulate to get the correct gas value
		this will be feasible enough until the core simulation works again
	*/

	//fund the tx with abundant amount of gas
	txf = txf.WithGas(10000000000)

	builder, err := tx.BuildUnsignedTx(txf, msgs...)
	if err != nil {
		return nil, err
	}

	err = tx.Sign(txf, clientCtx.GetFromName(), builder, false)
	if err != nil {
		return nil, err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(builder.GetTx())
	if err != nil {
		return nil, err
	}

	if txf.SimulateAndExecute() || clientCtx.Simulate {

		// simulate by calling ABCI Query
		query := abci.RequestQuery{
			Path: "/app/simulate",
			Data: txBytes,
		}

		queryResult, err := clientCtx.QueryABCI(query)
		if err != nil {
			return nil, err
		}

		var simResponse sdk.SimulationResponse
		err = jsonpb.Unmarshal(strings.NewReader(string(queryResult.Value)), &simResponse)
		if err != nil {
			return nil, err
		}

		//around 750000 - 1500000
		adjusted := simResponse.GetGasUsed()
		txf = txf.WithGas(adjusted * 10000)
	}

	await = true
	if await {
		// BroadcastTxCommit - full synced commit with await
		res, err := clientCtx.BroadcastTxCommit(txBytes)
		return res, err
	}

	// BroadcastTxSync - only CheckTx, don't wait confirmation
	return clientCtx.BroadcastTxSync(txBytes)
}

func (c *cosmosClient) QueueBroadcastMsg(msgs ...sdk.Msg) error {
	if !c.canSign {
		return ErrReadOnly
	} else if atomic.LoadInt64(&c.closed) == 1 {
		return ErrQueueClosed
	}

	t := time.NewTimer(10 * time.Second)
	for _, msg := range msgs {
		select {
		case <-t.C:
			return ErrEnqueueTimeout
		case c.msgC <- msg:
		}
	}
	t.Stop()

	return nil
}

func (c *cosmosClient) Close() {
	if !c.canSign {
		return
	}

	if atomic.CompareAndSwapInt64(&c.closed, 0, 1) {
		close(c.msgC)
	}

	<-c.doneC
}

const (
	msgCommitBatchSizeLimit = 512
	msgCommitBatchTimeLimit = 500 * time.Millisecond
)

func (c *cosmosClient) runBatchBroadcast() {
	expirationTimer := time.NewTimer(msgCommitBatchTimeLimit)
	msgBatch := make([]sdk.Msg, 0, msgCommitBatchSizeLimit)

	resetBatch := func() {
		msgBatch = msgBatch[:0]

		expirationTimer.Reset(msgCommitBatchTimeLimit)
	}

	submitBatch := func() {
		c.syncMux.Lock()
		defer c.syncMux.Unlock()

		c.txFactory = c.txFactory.WithSequence(c.accSeq)
		log.Debugln("broadcastTx with nonce", c.accSeq)
		res, err := c.broadcastTx(c.ctx, c.txFactory, true, msgBatch...)
		if err != nil {
			if strings.HasPrefix(err.Error(), "account sequence mismatch") {
				c.syncNonce()
				c.txFactory = c.txFactory.WithSequence(c.accSeq)
				log.Debugln("retrying broadcastTx with nonce", c.accSeq)
				res, err = c.broadcastTx(c.ctx, c.txFactory, true, msgBatch...)
			}
			if err != nil {
				resJSON, _ := json.MarshalIndent(res, "", "\t")
				c.logger.WithField("size", len(msgBatch)).WithError(err).Errorln("failed to commit msg batch:", string(resJSON))
				return
			}
		}

		c.accSeq++
		log.Debugln("nonce incremented to ", c.accSeq)
	}

	for {
		select {
		case msg, ok := <-c.msgC:
			if !ok {
				// exit required
				if len(msgBatch) > 0 {
					submitBatch()
				}

				close(c.doneC)
				return
			}

			msgBatch = append(msgBatch, msg)

			if len(msgBatch) >= msgCommitBatchSizeLimit {
				submitBatch()
				resetBatch()
			}
		case <-expirationTimer.C:
			if len(msgBatch) > 0 {
				submitBatch()
			}

			resetBatch()
		}
	}
}
