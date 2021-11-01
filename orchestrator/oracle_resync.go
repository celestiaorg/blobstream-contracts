package orchestrator

import (
	"context"
)

// GetLastCheckedBlock retrieves the last claim event this oracle has relayed to Cosmos.
func (p *peggyOrchestrator) GetLastCheckedBlock(ctx context.Context) (uint64, error) {

	lastClaimEvent, err := p.cosmosQueryClient.LastClaimEventByAddr(ctx, p.peggyBroadcastClient.AccFromAddress())
	if err != nil {
		return uint64(0), err
	}

	return lastClaimEvent.EthereumEventHeight, nil
}
