package relayer

import (
	"context"

	"github.com/InjectiveLabs/peggo/orchestrator/cosmos"
	"github.com/InjectiveLabs/peggo/orchestrator/cosmos/tmclient"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/peggy"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/provider"
	"github.com/InjectiveLabs/peggo/orchestrator/metrics"
	"github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

type PeggyRelayer interface {
	Start(ctx context.Context) error

	FindLatestValset(ctx context.Context) (*types.Valset, error)
	RelayBatches(ctx context.Context) error
	RelayValsets(ctx context.Context) error
}

type peggyRelayer struct {
	svcTags metrics.Tags

	tmClient             tmclient.TendermintClient
	cosmosQueryClient    cosmos.PeggyQueryClient
	peggyContract        peggy.PeggyContract
	ethProvider          provider.EVMProvider
	valsetRelayEnabled   bool
	relayValsetOffsetDur string
	batchRelayEnabled    bool
	relayBatchOffsetDur  string
}

func NewPeggyRelayer(
	cosmosQueryClient cosmos.PeggyQueryClient,
	tmClient tmclient.TendermintClient,
	peggyContract peggy.PeggyContract,
	valsetRelayEnabled bool,
	relayValsetOffsetDur string,
	batchRelayEnabled bool,
	relayBatchOffsetDur string,
) PeggyRelayer {
	return &peggyRelayer{
		tmClient:             tmClient,
		cosmosQueryClient:    cosmosQueryClient,
		peggyContract:        peggyContract,
		ethProvider:          peggyContract.Provider(),
		valsetRelayEnabled:   valsetRelayEnabled,
		relayValsetOffsetDur: relayValsetOffsetDur,
		batchRelayEnabled:    batchRelayEnabled,
		relayBatchOffsetDur:  relayBatchOffsetDur,
		svcTags: metrics.Tags{
			"svc": "peggy_relayer",
		},
	}
}
