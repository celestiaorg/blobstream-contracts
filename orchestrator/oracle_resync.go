package orchestrator

import (
	"context"

	"github.com/pkg/errors"
)

const defaultBlocksToSearch = 2000

// GetLastCheckedBlock retrieves the last claim event this oracle has relayed to Cosmos.
func (s *peggyOrchestrator) GetLastCheckedBlock(ctx context.Context) (uint64, error) {
	latestHeader, err := s.ethProvider.HeaderByNumber(ctx, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to get latest header")
		return 0, err
	}
	currentBlock := latestHeader.Number.Uint64()

	lastClaimEvent, err := s.cosmosQueryClient.LastClaimEventByAddr(ctx, s.peggyBroadcastClient.AccFromAddress())
	if err != nil {
		return currentBlock, err
	} else if lastClaimEvent.EthereumEventNonce == 0 {
		// TODO: Instead of currentBlock,return the block height at which peggy contract is deployed.
		return currentBlock, nil
	}

	return lastClaimEvent.EthereumEventHeight, nil
}

var ErrNotFound = errors.New("not found")
