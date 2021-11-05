package cosmos

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/InjectiveLabs/peggo/orchestrator/metrics"
	"github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

type PeggyQueryClient interface {
	ValsetAt(ctx context.Context, nonce uint64) (*types.Valset, error)
	CurrentValset(ctx context.Context) (*types.Valset, error)
	OldestUnsignedValsets(ctx context.Context, valAccountAddress sdk.AccAddress) ([]*types.Valset, error)
	LatestValsets(ctx context.Context) ([]*types.Valset, error)
	AllValsetConfirms(ctx context.Context, nonce uint64) ([]*types.MsgValsetConfirm, error)
	OldestUnsignedTransactionBatch(ctx context.Context, valAccountAddress sdk.AccAddress) (*types.OutgoingTxBatch, error)
	LatestTransactionBatches(ctx context.Context) ([]*types.OutgoingTxBatch, error)
	UnbatchedTokensWithFees(ctx context.Context) ([]*types.BatchFees, error)

	TransactionBatchSignatures(ctx context.Context, nonce uint64, tokenContract ethcmn.Address) ([]*types.MsgConfirmBatch, error)
	LastClaimEventByAddr(ctx context.Context, validatorAccountAddress sdk.AccAddress) (*types.LastClaimEvent, error)

	PeggyParams(ctx context.Context) (*types.Params, error)
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
		metrics.ReportFuncError(s.svcTags)
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
		metrics.ReportFuncError(s.svcTags)
		return nil, ErrNotFound
	}

	return daemonResp.Valset, nil
}

func (s *peggyQueryClient) OldestUnsignedValsets(ctx context.Context, valAccountAddress sdk.AccAddress) ([]*types.Valset, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.LastPendingValsetRequestByAddr(ctx, &types.QueryLastPendingValsetRequestByAddrRequest{
		Address: valAccountAddress.String(),
	})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query LastPendingValsetRequestByAddr from daemon")
		return nil, err
	} else if daemonResp == nil {
		metrics.ReportFuncError(s.svcTags)
		return nil, ErrNotFound
	}

	return daemonResp.Valsets, nil
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
		metrics.ReportFuncError(s.svcTags)
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
		metrics.ReportFuncError(s.svcTags)
		return nil, ErrNotFound
	}

	return daemonResp.Confirms, nil
}

func (s *peggyQueryClient) OldestUnsignedTransactionBatch(ctx context.Context, valAccountAddress sdk.AccAddress) (*types.OutgoingTxBatch, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.LastPendingBatchRequestByAddr(ctx, &types.QueryLastPendingBatchRequestByAddrRequest{
		Address: valAccountAddress.String(),
	})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query LastPendingBatchRequestByAddr from daemon")
		return nil, err
	} else if daemonResp == nil {
		metrics.ReportFuncError(s.svcTags)
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
		metrics.ReportFuncError(s.svcTags)
		return nil, ErrNotFound
	}

	return daemonResp.Batches, nil
}

func (s *peggyQueryClient) UnbatchedTokensWithFees(ctx context.Context) ([]*types.BatchFees, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.BatchFees(ctx, &types.QueryBatchFeeRequest{})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query BatchFees from daemon")
		return nil, err
	} else if daemonResp == nil {
		metrics.ReportFuncError(s.svcTags)
		return nil, ErrNotFound
	}

	return daemonResp.BatchFees, nil
}

func (s *peggyQueryClient) TransactionBatchSignatures(ctx context.Context, nonce uint64, tokenContract ethcmn.Address) ([]*types.MsgConfirmBatch, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.BatchConfirms(ctx, &types.QueryBatchConfirmsRequest{
		Nonce:           nonce,
		ContractAddress: tokenContract.String(),
	})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query BatchConfirms from daemon")
		return nil, err
	} else if daemonResp == nil {
		metrics.ReportFuncError(s.svcTags)
		return nil, ErrNotFound
	}

	return daemonResp.Confirms, nil
}

func (s *peggyQueryClient) LastClaimEventByAddr(ctx context.Context, validatorAccountAddress sdk.AccAddress) (*types.LastClaimEvent, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.LastEventByAddr(ctx, &types.QueryLastEventByAddrRequest{
		Address: validatorAccountAddress.String(),
	})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query LastEventByAddr from daemon")
		return nil, err
	} else if daemonResp == nil {
		metrics.ReportFuncError(s.svcTags)
		return nil, ErrNotFound
	}

	return daemonResp.LastClaimEvent, nil
}

func (s *peggyQueryClient) PeggyParams(ctx context.Context) (*types.Params, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	daemonResp, err := s.daemonQueryClient.Params(ctx, &types.QueryParamsRequest{})
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "failed to query PeggyParams from daemon")
		return nil, err
	} else if daemonResp == nil {
		metrics.ReportFuncError(s.svcTags)
		return nil, ErrNotFound
	}

	return &daemonResp.Params, nil
}
