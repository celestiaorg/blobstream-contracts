package relayer

import (
	"context"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/ethereum/util"
	wrappers "github.com/celestiaorg/quantum-gravity-bridge/ethereum/solidity/wrappers/QuantumGravityBridge.sol"
	"github.com/umee-network/umee/x/peggy/types"
)

const defaultBlocksToSearch = 2000

// FindLatestValset finds the latest valset on the Peggy contract by looking back through the event
// history and finding the most recent ValsetUpdatedEvent. Most of the time this will be very fast
// as the latest update will be in recent blockchain history and the search moves from the present
// backwards in time. In the case that the validator set has not been updated for a very long time
// this will take longer.
func (s *peggyRelayer) FindLatestValset(ctx context.Context) (*types.Valset, error) {
	latestHeader, err := s.ethProvider.HeaderByNumber(ctx, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to get latest header")
		return nil, err
	}
	currentBlock := latestHeader.Number.Uint64()

	peggyFilterer, err := wrappers.NewQuantumGravityBridgeFilterer(s.peggyContract.Address(), s.ethProvider)
	if err != nil {
		err = errors.Wrap(err, "failed to init Peggy events filterer")
		return nil, err
	}

	latestEthereumValsetNonce, err := s.peggyContract.GetValsetNonce(ctx, s.peggyContract.FromAddress())
	if err != nil {
		err = errors.Wrap(err, "failed to get latest Valset nonce")
		return nil, err
	}

	cosmosValset, err := s.cosmosQueryClient.ValsetAt(ctx, latestEthereumValsetNonce.Uint64())
	if err != nil {
		err = errors.Wrap(err, "failed to get cosmos Valset")
		return nil, err
	}

	for currentBlock > 0 {
		s.logger.Debug().
			Uint64("block", currentBlock).
			Msg("about to submit a Valset or Batch looking back into the history to find the last Valset Update")

		var endSearchBlock uint64
		if currentBlock <= defaultBlocksToSearch {
			endSearchBlock = 0
		} else {
			endSearchBlock = currentBlock - defaultBlocksToSearch
		}

		var valsetUpdatedEvents []*wrappers.QuantumGravityBridgeValidatorSetUpdatedEvent
		iter, err := peggyFilterer.FilterValidatorSetUpdatedEvent(&bind.FilterOpts{
			Start: endSearchBlock,
			End:   &currentBlock,
		}, nil)

		if err != nil {
			err = errors.Wrap(err, "failed to filter past ValsetUpdated events from Ethereum")
			return nil, err
		}

		for iter.Next() {
			valsetUpdatedEvents = append(valsetUpdatedEvents, iter.Event)
		}

		iter.Close()

		// by default the lowest found valset goes first, we want the highest
		//
		// TODO(xlab): this follows the original impl, but sort might be skipped there:
		// we could access just the latest element later.
		sort.Sort(sort.Reverse(PeggyValsetUpdatedEvents(valsetUpdatedEvents)))

		s.logger.Debug().
			Interface("valset_updated_events", valsetUpdatedEvents).
			Msg("found ValsetUpdated events")

		// we take only the first event if we find any at all.
		if len(valsetUpdatedEvents) > 0 {
			event := valsetUpdatedEvents[0]
			//valsetHash := event.ValidatorSetHash
			//
			//for idx, p := range event.ValidatorSetHash {
			//	valset.Members = append(valset.Members, &types.BridgeValidator{
			//		Power:           p.Uint64(),
			//		EthereumAddress: event.Validators[idx].Hex(),
			//	})
			//}

			// FIXME: check according hash and nonce only
			s.checkIfValsetsDiffer(cosmosValset, event.ValidatorSetHash, event.Nonce)
			// FIXME: actually return valset that matches the event.ValidatorSetHash
			return &types.Valset{}, nil
		}

		currentBlock = endSearchBlock
	}

	return nil, ErrNotFound
}

var ErrNotFound = errors.New("not found")

type PeggyValsetUpdatedEvents []*wrappers.QuantumGravityBridgeValidatorSetUpdatedEvent

func (a PeggyValsetUpdatedEvents) Len() int { return len(a) }
func (a PeggyValsetUpdatedEvents) Less(i, j int) bool {
	return a[i].Nonce.Cmp(a[j].Nonce) < 0
}
func (a PeggyValsetUpdatedEvents) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// This function exists to provide a warning if Cosmos and Ethereum have different validator sets
// for a given nonce. In the mundane version of this warning the validator sets disagree on sorting order
// which can happen if some relayer uses an unstable sort, or in a case of a mild griefing attack.
// The Peggy contract validates signatures in order of highest to lowest power. That way it can exit
// the loop early once a vote has enough power, if a relayer where to submit things in the reverse order
// they could grief users of the contract into paying more in gas.
// The other (and far worse) way a disagreement here could occur is if validators are colluding to steal
// funds from the Peggy contract and have submitted a hijacking update. If slashing for off Cosmos chain
// Ethereum signatures is implemented you would put that handler here.
func (s *peggyRelayer) checkIfValsetsDiffer(cosmosValset *types.Valset, ethValsetHash [32]byte, ethNonce *big.Int) {
	if cosmosValset == nil && (ethNonce == nil || len(ethNonce.Bits())==0) {
		// bootstrapping case
		return
	} else if cosmosValset == nil {
		s.logger.Error().
			Uint64("eth_valset_nonce", ethNonce.Uint64()).
			Msg("cosmos does not have a valset for nonce from Ethereum chain. Possible bridge hijacking!")
		return
	}

	if cosmosValset.Nonce != ethNonce.Uint64() {

		s.logger.Error().
			Uint64("eth_valset_nonce", ethNonce.Uint64()).
			Uint64("cosmos_valset_nonce", cosmosValset.Nonce).
			Msg("cosmos does have a wrong valset nonce, differs from Ethereum chain. Possible bridge hijacking!")
		return
	}

	BridgeValidators(cosmosValset.Members).Sort()
	// FIXME: hash the validators according to same logic as on ethereum and compare to
	// ethValsetHash
	//for idx, member := range cosmosValset.Members {
	//	if ethereumValset.Members[idx].EthereumAddress != member.EthereumAddress {
	//		s.logger.Error().Msg("valsets are different, a sorting error?")
	//	}
	//	if ethereumValset.Members[idx].Power != member.Power {
	//		s.logger.Error().Msg("valsets are different, a sorting error?")
	//	}
	//}
}

type BridgeValidators []*types.BridgeValidator

// Sort sorts the validators by power
func (b BridgeValidators) Sort() {
	sort.Slice(b, func(i, j int) bool {
		if b[i].Power == b[j].Power {
			// Secondary sort on eth address in case powers are equal
			return util.EthAddrLessThan(b[i].EthereumAddress, b[j].EthereumAddress)
		}
		return b[i].Power > b[j].Power
	})
}

// HasDuplicates returns true if there are duplicates in the set
func (b BridgeValidators) HasDuplicates() bool {
	m := make(map[string]struct{}, len(b))
	for i := range b {
		m[b[i].EthereumAddress] = struct{}{}
	}
	return len(m) != len(b)
}

// GetPowers returns only the power values for all members
func (b BridgeValidators) GetPowers() []uint64 {
	r := make([]uint64, len(b))
	for i := range b {
		r[i] = b[i].Power
	}
	return r
}
