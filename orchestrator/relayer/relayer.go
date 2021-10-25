package relayer

import (
	"context"

	"github.com/umee-network/peggo/orchestrator/cosmos"
	"github.com/umee-network/peggo/orchestrator/ethereum/peggy"
	"github.com/umee-network/peggo/orchestrator/ethereum/provider"
	"github.com/umee-network/umee/x/peggy/types"
)

type PeggyRelayer interface {
	Start(ctx context.Context) error

	FindLatestValset(ctx context.Context) (*types.Valset, error)
	RelayBatches(ctx context.Context) error
	RelayValsets(ctx context.Context) error
}

type peggyRelayer struct {
	cosmosQueryClient  cosmos.PeggyQueryClient
	peggyContract      peggy.PeggyContract
	ethProvider        provider.EVMProvider
	valsetRelayEnabled bool
	batchRelayEnabled  bool
}

func NewPeggyRelayer(
	cosmosQueryClient cosmos.PeggyQueryClient,
	peggyContract peggy.PeggyContract,
	valsetRelayEnabled bool,
	batchRelayEnabled bool,
) PeggyRelayer {
	return &peggyRelayer{
		cosmosQueryClient:  cosmosQueryClient,
		peggyContract:      peggyContract,
		ethProvider:        peggyContract.Provider(),
		valsetRelayEnabled: valsetRelayEnabled,
		batchRelayEnabled:  batchRelayEnabled,
	}
}
