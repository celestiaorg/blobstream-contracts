package sidechain

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/InjectiveLabs/peggo/orchestrator/metrics"
	"github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

type PeggyQueryClient interface {
	ValsetAt(ctx context.Context, nonce uint64) (*types.Valset, error)
	CurrentValset(ctx context.Context) (*types.Valset, error)
	OldestUnsignedValset(ctx context.Context, address sdk.AccAddress) (*types.Valset, error)
	LatestValsets(ctx context.Context) ([]*types.Valset, error)
	AllValsetConfirms(ctx context.Context, nonce uint64) ([]*types.MsgValsetConfirm, error)
	OldestUnsignedTransactionBatch(ctx context.Context, address sdk.AccAddress) (*types.OutgoingTxBatch, error)
	LatestTransactionBatches(ctx context.Context) ([]*types.OutgoingTxBatch, error)
	LatestUnbatchOutgoingTx(ctx context.Context, contractAddr string) ([]*types.OutgoingTx, error)
	TransactionBatchSignatures(ctx context.Context, nonce uint64, tokenContract common.Address) ([]*types.MsgConfirmBatch, error)
	LastEventNonce(ctx context.Context, address sdk.AccAddress) (uint64, error)
}

func NewPeggyQueryClient(client types.QueryClient) PeggyQueryClient {
	return &peggyQueryClient{
		daemonQueryClient: client,

		svcTags: metrics.Tags{
			"svc": "peggy_query",
		},
	}
}

type peggyQueryClient struct {
	daemonQueryClient types.QueryClient

	svcTags metrics.Tags
}

var ErrNotFound = errors.New("not found")

func (s *peggyQueryClient) ValsetAt(ctx context.Context, nonce uint64) (*types.Valset, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.ValsetRequest(ctx, &types.QueryValsetRequestRequest{
		Nonce: nonce,
	})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query ValsetRequest from daemon")
		return nil, err
	} else if daemonResp == nil {
		return nil, ErrNotFound
	}

	return daemonResp.Valset, nil
}

func (s *peggyQueryClient) CurrentValset(ctx context.Context) (*types.Valset, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.CurrentValset(ctx, &types.QueryCurrentValsetRequest{})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query CurrentValset from daemon")
		return nil, err
	} else if daemonResp == nil {
		return nil, ErrNotFound
	}

	return daemonResp.Valset, nil
}

func (s *peggyQueryClient) OldestUnsignedValset(ctx context.Context, address sdk.AccAddress) (*types.Valset, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.LastPendingValsetRequestByAddr(ctx, &types.QueryLastPendingValsetRequestByAddrRequest{
		Address: address.String(),
	})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query LastPendingValsetRequestByAddr from daemon")
		return nil, err
	} else if daemonResp == nil {
		return nil, ErrNotFound
	}

	return daemonResp.Valset, nil
}

func (s *peggyQueryClient) LatestValsets(ctx context.Context) ([]*types.Valset, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.LastValsetRequests(ctx, &types.QueryLastValsetRequestsRequest{})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query LastValsetRequests from daemon")
		return nil, err
	} else if daemonResp == nil {
		return nil, ErrNotFound
	}

	return daemonResp.Valsets, nil
}

func (s *peggyQueryClient) AllValsetConfirms(ctx context.Context, nonce uint64) ([]*types.MsgValsetConfirm, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.ValsetConfirmsByNonce(ctx, &types.QueryValsetConfirmsByNonceRequest{
		Nonce: nonce,
	})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query ValsetConfirmsByNonce from daemon")
		return nil, err
	} else if daemonResp == nil {
		return nil, ErrNotFound
	}

	return daemonResp.Confirms, nil
}

func (s *peggyQueryClient) OldestUnsignedTransactionBatch(ctx context.Context, address sdk.AccAddress) (*types.OutgoingTxBatch, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.LastPendingBatchRequestByAddr(ctx, &types.QueryLastPendingBatchRequestByAddrRequest{
		Address: address.String(),
	})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query LastPendingBatchRequestByAddr from daemon")
		return nil, err
	} else if daemonResp == nil {
		return nil, ErrNotFound
	}

	return daemonResp.Batch, nil
}

func (s *peggyQueryClient) LatestTransactionBatches(ctx context.Context) ([]*types.OutgoingTxBatch, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.OutgoingTxBatches(ctx, &types.QueryOutgoingTxBatchesRequest{})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query OutgoingTxBatches from daemon")
		return nil, err
	} else if daemonResp == nil {
		return nil, ErrNotFound
	}

	return daemonResp.Batches, nil
}

func (s *peggyQueryClient) LatestUnbatchOutgoingTx(ctx context.Context, contractAddr string) ([]*types.OutgoingTx, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.UnbatchedTxPool(ctx, &types.QueryUnbatchedTxPoolByAddrRequest{Address: contractAddr})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query UnbatchedTxPool from daemon")
		return nil, err
	} else if daemonResp == nil {
		return nil, ErrNotFound
	}

	return daemonResp.OutgoingTxs, nil
}

func (s *peggyQueryClient) TransactionBatchSignatures(ctx context.Context, nonce uint64, tokenContract common.Address) ([]*types.MsgConfirmBatch, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.BatchConfirms(ctx, &types.QueryBatchConfirmsRequest{Nonce: nonce, ContractAddress: tokenContract.Hex()})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query BatchConfirms from daemon")
		return nil, err
	} else if daemonResp == nil {
		return nil, ErrNotFound
	}

	return daemonResp.Confirms, nil
}

func (s *peggyQueryClient) LastEventNonce(ctx context.Context, address sdk.AccAddress) (uint64, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.LastEventNonceByAddr(ctx, &types.QueryLastEventNonceByAddrRequest{
		Address: address.String(),
	})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query LastEventNonceByAddr from daemon")
		return 0, err
	} else if daemonResp == nil {
		return 0, ErrNotFound
	}

	return daemonResp.EventNonce, nil
}
