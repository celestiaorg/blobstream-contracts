package orchestrator

import (
	"context"

	ethcmn "github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/peggo/orchestrator/coingecko"
	"github.com/InjectiveLabs/peggo/orchestrator/cosmos/tmclient"

	sidechain "github.com/InjectiveLabs/peggo/orchestrator/cosmos"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/keystore"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/peggy"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/provider"
	"github.com/InjectiveLabs/peggo/orchestrator/metrics"
	"github.com/InjectiveLabs/peggo/orchestrator/relayer"
)

type PeggyOrchestrator interface {
	Start(ctx context.Context) error

	CheckForEvents(ctx context.Context, startingBlock uint64) (currentBlock uint64, err error)
	GetLastCheckedBlock(ctx context.Context) (uint64, error)

	EthOracleMainLoop(ctx context.Context) error
	EthSignerMainLoop(ctx context.Context) error
	BatchRequesterLoop(ctx context.Context) error
	RelayerMainLoop(ctx context.Context) error
}

type peggyOrchestrator struct {
	svcTags metrics.Tags

	tmClient             tmclient.TendermintClient
	cosmosQueryClient    sidechain.PeggyQueryClient
	peggyBroadcastClient sidechain.PeggyBroadcastClient
	peggyContract        peggy.PeggyContract
	ethProvider          provider.EVMProvider
	ethFrom              ethcmn.Address
	ethSignerFn          keystore.SignerFn
	ethPersonalSignFn    keystore.PersonalSignFn
	erc20ContractMapping map[ethcmn.Address]string
	relayer              relayer.PeggyRelayer
	minBatchFeeUSD       float64
	priceFeeder          *coingecko.CoingeckoPriceFeed
}

func NewPeggyOrchestrator(
	cosmosQueryClient sidechain.PeggyQueryClient,
	peggyBroadcastClient sidechain.PeggyBroadcastClient,
	tmClient tmclient.TendermintClient,
	peggyContract peggy.PeggyContract,
	ethFrom ethcmn.Address,
	ethSignerFn keystore.SignerFn,
	ethPersonalSignFn keystore.PersonalSignFn,
	erc20ContractMapping map[ethcmn.Address]string,
	relayer relayer.PeggyRelayer,
	minBatchFeeUSD float64,
	priceFeeder *coingecko.CoingeckoPriceFeed,

) PeggyOrchestrator {
	return &peggyOrchestrator{
		tmClient:             tmClient,
		cosmosQueryClient:    cosmosQueryClient,
		peggyBroadcastClient: peggyBroadcastClient,
		peggyContract:        peggyContract,
		ethProvider:          peggyContract.Provider(),
		ethFrom:              ethFrom,
		ethSignerFn:          ethSignerFn,
		ethPersonalSignFn:    ethPersonalSignFn,
		erc20ContractMapping: erc20ContractMapping,
		relayer:              relayer,
		minBatchFeeUSD:       minBatchFeeUSD,
		priceFeeder:          priceFeeder,
		svcTags: metrics.Tags{
			"svc": "peggy_orchestrator",
		},
	}
}
