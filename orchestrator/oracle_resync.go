package orchestrator

import (
	"context"
)

// GetLastCheckedBlock retrieves the last claim event this oracle has relayed to Cosmos.
func (s *peggyOrchestrator) GetLastCheckedBlock(ctx context.Context) (uint64, error) {
	lastClaimEvent, err := s.cosmosQueryClient.LastClaimEventByAddr(ctx, s.peggyBroadcastClient.AccFromAddress())
	if err != nil {
		return uint64(0), err
	}

	return lastClaimEvent.EthereumEventHeight, nil
}
