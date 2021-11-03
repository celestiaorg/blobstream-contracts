package orchestrator

import (
	"context"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"

	"github.com/umee-network/peggo/orchestrator/coingecko"
	"github.com/umee-network/peggo/orchestrator/cosmos/tmclient"

	sidechain "github.com/umee-network/peggo/orchestrator/cosmos"
	"github.com/umee-network/peggo/orchestrator/ethereum/keystore"
	"github.com/umee-network/peggo/orchestrator/ethereum/peggy"
	"github.com/umee-network/peggo/orchestrator/ethereum/provider"
	"github.com/umee-network/peggo/orchestrator/relayer"
)

type PeggyOrchestrator interface {
	Start(ctx context.Context) error
	CheckForEvents(ctx context.Context, startingBlock uint64) (currentBlock uint64, err error)
	GetLastCheckedBlock(ctx context.Context) (uint64, error)
	EthOracleMainLoop(ctx context.Context) error
	EthSignerMainLoop(ctx context.Context) error
	BatchRequesterLoop(ctx context.Context) error
	RelayerMainLoop(ctx context.Context) error

	SetMinBatchFee(float64)
	SetPriceFeeder(*coingecko.PriceFeed)
}

type peggyOrchestrator struct {
	logger               zerolog.Logger
	tmClient             tmclient.TendermintClient
	cosmosQueryClient    sidechain.PeggyQueryClient
	peggyBroadcastClient sidechain.PeggyBroadcastClient
	peggyContract        peggy.Contract
	ethProvider          provider.EVMProvider
	ethFrom              ethcmn.Address
	ethSignerFn          keystore.SignerFn
	ethPersonalSignFn    keystore.PersonalSignFn
	relayer              relayer.PeggyRelayer

	// optional inputs with defaults
	minBatchFeeUSD float64
	priceFeeder    *coingecko.PriceFeed
}

func NewPeggyOrchestrator(
	logger zerolog.Logger,
	cosmosQueryClient sidechain.PeggyQueryClient,
	peggyBroadcastClient sidechain.PeggyBroadcastClient,
	tmClient tmclient.TendermintClient,
	peggyContract peggy.Contract,
	ethFrom ethcmn.Address,
	ethSignerFn keystore.SignerFn,
	ethPersonalSignFn keystore.PersonalSignFn,
	relayer relayer.PeggyRelayer,
	options ...func(PeggyOrchestrator),
) PeggyOrchestrator {

	orch := &peggyOrchestrator{
		logger:               logger.With().Str("module", "orchestrator").Logger(),
		tmClient:             tmClient,
		cosmosQueryClient:    cosmosQueryClient,
		peggyBroadcastClient: peggyBroadcastClient,
		peggyContract:        peggyContract,
		ethProvider:          peggyContract.Provider(),
		ethFrom:              ethFrom,
		ethSignerFn:          ethSignerFn,
		ethPersonalSignFn:    ethPersonalSignFn,
		relayer:              relayer,
	}

	for _, option := range options {
		option(orch)
	}

	return orch
}
