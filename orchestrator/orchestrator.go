package orchestrator

import (
	"context"
	"time"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"

	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/coingecko"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/cosmos/tmclient"

	sidechain "github.com/celestiaorg/quantum-gravity-bridge/orchestrator/cosmos"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/keystore"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/peggy"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/provider"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/relayer"
)

type PeggyOrchestrator interface {
	Start(ctx context.Context) error
	// CheckForEvents(ctx context.Context, startingBlock uint64) (currentBlock uint64, err error)
	GetLastCheckedBlock(ctx context.Context) (uint64, error)
	EthSignerMainLoop(ctx context.Context) error
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
	loopsDuration        time.Duration
	cosmosBlockTime      time.Duration

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
	loopDuration time.Duration,
	cosmosBlockTime time.Duration,
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
		loopsDuration:        loopDuration,
		cosmosBlockTime:      cosmosBlockTime,
	}

	for _, option := range options {
		option(orch)
	}

	return orch
}
