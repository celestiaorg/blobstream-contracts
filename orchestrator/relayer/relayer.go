package relayer

import (
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/peggy"
	"github.com/InjectiveLabs/peggo/orchestrator/ethereum/provider"
	"github.com/InjectiveLabs/peggo/orchestrator/metrics"
	"github.com/InjectiveLabs/peggo/orchestrator/sidechain"
)

type PeggyRelayer interface {
	RunLoop()
}

type peggyRelayer struct {
	svcTags metrics.Tags

	cosmosQueryClient sidechain.PeggyQueryClient
	peggyContract     peggy.PeggyContract
	ethProvider       provider.EVMProvider
}

func NewPeggyRelayer(
	cosmosQueryClient sidechain.PeggyQueryClient,
	peggyContract peggy.PeggyContract,
) PeggyRelayer {
	return &peggyRelayer{
		cosmosQueryClient: cosmosQueryClient,
		peggyContract:     peggyContract,
		ethProvider:       peggyContract.Provider(),

		svcTags: metrics.Tags{
			"svc": "peggy_relayer",
		},
	}
}
