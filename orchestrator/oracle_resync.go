package orchestrator

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	wrappers "github.com/InjectiveLabs/peggo/solidity/wrappers/Peggy.sol"
)

const defaultBlocksToSearch = 2000

// GetLastCheckedBlock retrieves the last event nonce this oracle has relayed to Cosmos
// it then uses the Ethereum indexes to determine what block the last entry
//
// TODO this should simply be stored in the deposit or withdraw claim and we
// ask the Cosmos chain, this searching is a total waste of work
func (s *peggyOrchestrator) GetLastCheckedBlock(ctx context.Context) (uint64, error) {
	latestHeader, err := s.ethProvider.HeaderByNumber(ctx, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to get latest header")
		return 0, err
	}
	currentBlock := latestHeader.Number.Uint64()

	lastEventNonce, err := s.cosmosQueryClient.LastEventNonce(ctx, s.peggyBroadcastClient.ValFromAddress())
	if err != nil {
		return 0, err
	} else if lastEventNonce == 0 {
		return currentBlock, nil
	}

	peggyFilterer, err := wrappers.NewPeggyFilterer(s.peggyContract.Address(), s.ethProvider)
	if err != nil {
		err = errors.Wrap(err, "failed to init Peggy events filterer")
		return 0, err
	}

	for currentBlock > 0 {
		log.WithField("current_block", currentBlock).
			Infoln("Oracle is resyncing, looking back into the history to find our last event nonce")

		var endSearchBlock uint64
		if currentBlock <= defaultBlocksToSearch {
			endSearchBlock = 0
		} else {
			endSearchBlock = currentBlock - defaultBlocksToSearch
		}

		var transactionBatchExecutedEvents []*wrappers.PeggyTransactionBatchExecutedEvent
		iter, err := peggyFilterer.FilterTransactionBatchExecutedEvent(&bind.FilterOpts{
			Start: endSearchBlock,
			End:   &currentBlock,
		}, nil, nil)
		if err != nil {
			err = errors.Wrap(err, "failed to filter past TransactionBatchExecuted events from Ethereum")
			return 0, err
		} else {
			for iter.Next() {
				transactionBatchExecutedEvents = append(transactionBatchExecutedEvents, iter.Event)
			}

			iter.Close()
		}

		var sendToCosmosEvents []*wrappers.PeggySendToCosmosEvent
		iter2, err := peggyFilterer.FilterSendToCosmosEvent(&bind.FilterOpts{
			Start: endSearchBlock,
			End:   &currentBlock,
		}, nil, nil, nil)
		if err != nil {
			err = errors.Wrap(err, "failed to filter past SendToCosmos events from Ethereum")
			return 0, err
		} else {
			for iter2.Next() {
				sendToCosmosEvents = append(sendToCosmosEvents, iter2.Event)
			}

			iter2.Close()
		}

		log.Debugln("Found TransactionBatchExecuted events", transactionBatchExecutedEvents)
		log.Debugln("Found SendToCosmos events", sendToCosmosEvents)

		for _, ev := range transactionBatchExecutedEvents {
			if ev.EventNonce.Uint64() == lastEventNonce {
				return ev.Raw.BlockNumber, nil
			}
		}

		for _, ev := range sendToCosmosEvents {
			if ev.EventNonce.Uint64() == lastEventNonce {
				return ev.Raw.BlockNumber, nil
			}
		}

		currentBlock = endSearchBlock
	}

	log.Warningf("Could not find the last event relayed by %s, Last Event nonce is %d but no event matching that could be found",
		s.peggyBroadcastClient.ValFromAddress(), lastEventNonce)

	return 0, ErrNotFound
}

var ErrNotFound = errors.New("not found")

const defaultRetryDur = 5 * time.Second

// /// gets the current block number, no matter how long it takes
// async fn get_block_number_with_retry(web3: &Web3) -> Uint256 {
//     let mut res = web3.eth_block_number().await;
//     while res.is_err() {
//         error!("Failed to get latest block! Is your Eth node working?");
//         delay_for(RETRY_TIME).await;
//         res = web3.eth_block_number().await;
//     }
//     res.unwrap()
// }

// /// gets the last event nonce, no matter how long it takes.
// async fn get_last_event_nonce_with_retry(
//     client: &mut PeggyQueryClient<Channel>,
//     our_cosmos_address: CosmosAddress,
// ) -> u64 {
//     let mut res = get_last_event_nonce(client, our_cosmos_address).await;
//     while res.is_err() {
//         error!("Failed to get last event nonce, is the Cosmos GRPC working?");
//         delay_for(RETRY_TIME).await;
//         res = get_last_event_nonce(client, our_cosmos_address).await;
//     }
//     res.unwrap()
// }
