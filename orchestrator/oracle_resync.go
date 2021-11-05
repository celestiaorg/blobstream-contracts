package orchestrator

import (
	"context"
	"github.com/InjectiveLabs/peggo/orchestrator/metrics"
)

// GetLastCheckedBlock retrieves the last claim event this oracle has relayed to Cosmos.
func (s *peggyOrchestrator) GetLastCheckedBlock(ctx context.Context) (uint64, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	lastClaimEvent, err := s.cosmosQueryClient.LastClaimEventByAddr(ctx, s.peggyBroadcastClient.AccFromAddress())
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		return uint64(0), err
	}

	return lastClaimEvent.EthereumEventHeight, nil
}
