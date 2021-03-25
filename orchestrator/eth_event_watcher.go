package orchestrator

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	wrappers "github.com/InjectiveLabs/peggo/solidity/wrappers/Peggy.sol"
)

// CheckForEvents checks for events such as a deposit to the Peggy Ethereum contract or a validator set update
// or a transaction batch update. It then responds to these events by performing actions on the Cosmos chain if required
func (s *peggyOrchestrator) CheckForEvents(
	ctx context.Context,
	startingBlock uint64,
) (currentBlock uint64, err error) {
	latestHeader, err := s.ethProvider.HeaderByNumber(ctx, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to get latest header")
		return 0, err
	}

	currentBlock = latestHeader.Number.Uint64()

	if (currentBlock - startingBlock) > defaultBlocksToSearch {
		currentBlock = startingBlock + defaultBlocksToSearch
	}

	peggyFilterer, err := wrappers.NewPeggyFilterer(s.peggyContract.Address(), s.ethProvider)
	if err != nil {
		err = errors.Wrap(err, "failed to init Peggy events filterer")
		return 0, err
	}

	var sendToCosmosEvents []*wrappers.PeggySendToCosmosEvent
	{
		iter, err := peggyFilterer.FilterSendToCosmosEvent(&bind.FilterOpts{
			Start: startingBlock,
			End:   &currentBlock,
		}, nil, nil, nil)
		if err != nil {
			log.WithFields(log.Fields{
				"start": startingBlock,
				"end":   currentBlock,
			}).Debugln("failed to filter past SendToCosmos events from Ethereum")

			if !isUnknownBlockErr(err) {
				err = errors.Wrap(err, "failed to filter past SendToCosmos events from Ethereum")
				return 0, err
			} else if iter == nil {
				return 0, errors.New("no iterator returned")
			}
		}

		for iter.Next() {
			sendToCosmosEvents = append(sendToCosmosEvents, iter.Event)
		}

		iter.Close()
	}
	log.Debugln("Deposits:", sendToCosmosEvents)

	var transactionBatchExecutedEvents []*wrappers.PeggyTransactionBatchExecutedEvent
	{
		iter, err := peggyFilterer.FilterTransactionBatchExecutedEvent(&bind.FilterOpts{
			Start: startingBlock,
			End:   &currentBlock,
		}, nil, nil)
		if err != nil {
			log.WithFields(log.Fields{
				"start": startingBlock,
				"end":   currentBlock,
			}).Debugln("failed to filter past TransactionBatchExecuted events from Ethereum")

			if !isUnknownBlockErr(err) {
				err = errors.Wrap(err, "failed to filter past TransactionBatchExecuted events from Ethereum")
				return 0, err
			} else if iter == nil {
				return 0, errors.New("no iterator returned")
			}
		}

		for iter.Next() {
			transactionBatchExecutedEvents = append(transactionBatchExecutedEvents, iter.Event)
		}

		iter.Close()
	}
	log.Debugln("Withdraws:", transactionBatchExecutedEvents)

	// note that starting block overlaps with our last checked block, because we have to deal with
	// the possibility that the relayer was killed after relaying only one of multiple events in a single
	// block, so we also need this routine so make sure we don't send in the first event in this hypothetical
	// multi event block again. In theory we only send all events for every block and that will pass of fail
	// atomicly but lets not take that risk.
	lastEventNonce, err := s.cosmosQueryClient.LastEventNonce(ctx, s.peggyBroadcastClient.AccFromAddress())
	if err != nil {
		err = errors.New("failed to query last event nonce from backend")
		return 0, err
	}

	deposits := filterSendToCosmosEventsByNonce(sendToCosmosEvents, lastEventNonce)
	withdraws := filterTransactionBatchExecutedEventsByNonce(transactionBatchExecutedEvents, lastEventNonce)

	if len(deposits) > 0 {
		log.WithFields(log.Fields{
			"sender":      deposits[0].Sender.Hex(),
			"destination": sdk.AccAddress(deposits[0].Destination[12:32]).String(),
			"amount":      deposits[0].Amount.String(),
			"event_nonce": deposits[0].EventNonce.String(),
		}).Infoln("Oracle observed a deposit")
	}

	if len(withdraws) > 0 {
		log.WithFields(log.Fields{
			"nonce":          withdraws[0].BatchNonce.String(),
			"token_contract": withdraws[0].Token.Hex(),
			"event_nonce":    withdraws[0].EventNonce.String(),
		}).Infoln("Oracle observed a withdraw batch")
	}

	if len(deposits) > 0 || len(withdraws) > 0 {
		// todo get eth chain id from the chain
		if err := s.peggyBroadcastClient.SendEthereumClaims(ctx, deposits, withdraws); err != nil {
			err = errors.Wrap(err, "failed to send ethereum claims to Cosmos chain")
			return 0, err
		}
	}

	return currentBlock, nil
}

func filterSendToCosmosEventsByNonce(
	events []*wrappers.PeggySendToCosmosEvent,
	nonce uint64,
) []*wrappers.PeggySendToCosmosEvent {
	res := make([]*wrappers.PeggySendToCosmosEvent, 0, len(events))

	for _, ev := range events {
		if ev.EventNonce.Uint64() > nonce {
			res = append(res, ev)
		}
	}

	return res
}

func filterTransactionBatchExecutedEventsByNonce(
	events []*wrappers.PeggyTransactionBatchExecutedEvent,
	nonce uint64,
) []*wrappers.PeggyTransactionBatchExecutedEvent {
	res := make([]*wrappers.PeggyTransactionBatchExecutedEvent, 0, len(events))

	for _, ev := range events {
		if ev.EventNonce.Uint64() > nonce {
			res = append(res, ev)
		}
	}

	return res
}

func isUnknownBlockErr(err error) bool {
	// Geth error
	if strings.Contains(err.Error(), "unknown block") {
		return true
	}

	// Parity error
	if strings.Contains(err.Error(), "One of the blocks specified in filter") {
		return true
	}

	return false
}
