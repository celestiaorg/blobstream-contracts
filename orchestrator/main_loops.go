package orchestrator

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/xlab/suplog"
	"golang.org/x/sync/errgroup"

	"github.com/InjectiveLabs/peggo/modules/peggy/types"
	"github.com/InjectiveLabs/peggo/orchestrator/cosmos"
	"github.com/InjectiveLabs/peggo/orchestrator/relayer"
)

const defaultLoopDur = 10 * time.Second

// Start combines the all major roles required to make
// up the Orchestrator, all of these are async loops.
func (s *peggyOrchestrator) Start(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.EthOracleMainLoop(ctx)
	})
	eg.Go(func() error {
		return s.BatchRequesterLoop(ctx)
	})
	eg.Go(func() error {
		return s.EthSignerMainLoop(ctx)
	})
	eg.Go(func() error {
		return s.RelayerMainLoop(ctx)
	})
	eg.Go(func() error {
		return s.ValsetRequesterLoop(ctx)
	})

	return eg.Wait()
}

// EthOracleMainLoop is responsible for making sure that Ethereum events are retrieved from the Ethereum blockchain
// and ferried over to Cosmos where they will be used to issue tokens or process batches.
//
// TODO this loop requires a method to bootstrap back to the correct event nonce when restarted
func (s *peggyOrchestrator) EthOracleMainLoop(ctx context.Context) (err error) {
	var lastCheckedBlock uint64
	for {
		lastCheckedBlock, err = s.GetLastCheckedBlock(ctx)
		if err != nil {
			log.WithError(err).Errorln("failed to get last checked block, retry in", defaultRetryDur)
			time.Sleep(defaultRetryDur)
			continue
		}

		log.Infoln("Oracle resync complete, Oracle now operational")
		break
	}

	t := time.NewTimer(0)
	for range t.C {
		latestHeader, err := s.ethProvider.HeaderByNumber(ctx, nil)
		if err != nil {
			log.WithError(err).Errorln("failed to get latest header, retry in", defaultRetryDur)
			t.Reset(defaultRetryDur)
			continue
		}

		latestEthBlock := latestHeader.Number.Uint64()
		latestCosmosBlock, err := s.tmClient.GetLatestBlockHeight()
		if err != nil {
			log.WithError(err).Errorln("failed to get latest cosmos block height, retry in", defaultRetryDur)
			t.Reset(defaultRetryDur)
			continue
		}

		log.Debugf("[ethOracleMainLoop] Latest Eth block %d, latest Cosmos block %d",
			latestEthBlock, latestCosmosBlock,
		)

		// Relays events from Ethereum -> Cosmos
		currentBlock, err := s.CheckForEvents(ctx, lastCheckedBlock)
		if err != nil {
			log.WithError(err).Errorln("error during eht event checking, retry in", defaultRetryDur)
			t.Reset(defaultRetryDur)
			continue
		} else {
			lastCheckedBlock = currentBlock
		}

		t.Reset(defaultLoopDur)
	}

	return nil
}

// EthSignerMainLoop simply signs off on any batches or validator sets provided by the validator
// since these are provided directly by a trusted Cosmsos node they can simply be assumed to be
// valid and signed off on.
func (s *peggyOrchestrator) EthSignerMainLoop(ctx context.Context) (err error) {
	var peggyID common.Hash
	for {
		peggyID, err = s.peggyContract.GetPeggyID(ctx, s.peggyContract.FromAddress())
		if err != nil {

			log.WithError(err).Errorln("failed to get PeggyID from Ethereum contract, retry in", defaultRetryDur)
			time.Sleep(defaultRetryDur)
			continue
		}
		log.Debugf("[ethSignerMainLoop] peggyID %s", peggyID.Hex())
		break
	}

	t := time.NewTimer(0)
	for range t.C {
		latestHeader, err := s.ethProvider.HeaderByNumber(ctx, nil)
		if err != nil {
			log.WithError(err).Errorln("failed to get latest header, retry in", defaultRetryDur)
			t.Reset(defaultRetryDur)
			continue
		}

		latestEthBlock := latestHeader.Number.Uint64()
		latestCosmosBlock, err := s.tmClient.GetLatestBlockHeight()
		if err != nil {
			log.WithError(err).Errorln("failed to get latest cosmos block height, retry in", defaultRetryDur)
			t.Reset(defaultRetryDur)
			continue
		}

		log.Debugf("[ethSignerMainLoop] Latest Eth block %d, latest Cosmos block %d",
			latestEthBlock, latestCosmosBlock,
		)

		valset, err := s.cosmosQueryClient.OldestUnsignedValset(ctx, s.peggyBroadcastClient.ValFromAddress())
		if err == cosmos.ErrNotFound {
			log.Debugln("no Valset waiting to be signed")
		} else if err != nil {
			log.WithError(err).Errorln("failed to get unsigned Valset for signing, retry in", defaultRetryDur)
			t.Reset(defaultRetryDur)
			continue
		} else if valset == nil {
			log.Debugf("no valset")
			t.Reset(defaultRetryDur)
		} else {
			log.Infoln("sending Valset confirm for %d", valset.Nonce)

			if err := s.peggyBroadcastClient.SendValsetConfirm(ctx, s.ethFrom, peggyID, valset); err != nil {
				log.WithError(err).Errorln("failed to sign and send Valset confirmation to Cosmos, retry in", defaultRetryDur)
				t.Reset(defaultRetryDur)
				continue
			}
		}

		// sign the last unsigned batch, TODO check if we already have signed this
		txBatch, err := s.cosmosQueryClient.OldestUnsignedTransactionBatch(ctx, s.peggyBroadcastClient.ValFromAddress())
		if err == cosmos.ErrNotFound {
			log.Debugln("no TransactionBatch waiting to be signed")
		} else if err != nil {
			log.WithError(err).Errorln("failed to get unsigned TransactionBatch for signing, retry in", defaultRetryDur)
			t.Reset(defaultRetryDur)
			continue
		} else if txBatch == nil {
			log.Debugln("no TransactionBatch")
			t.Reset(defaultRetryDur)
		} else {
			log.Infoln("sending TransactionBatch confirm for %d", txBatch.BatchNonce)

			if err := s.peggyBroadcastClient.SendBatchConfirm(ctx, s.ethFrom, peggyID, txBatch); err != nil {
				log.WithError(err).Errorln("failed to sign and send TransactionBatch confirmation to Cosmos, retry in", defaultRetryDur)
				t.Reset(defaultRetryDur)
				continue
			}
		}

		t.Reset(defaultLoopDur)
	}

	return nil
}

// This loop doesn't have a formal role per say, anyone can request a valset
// but there does need to be some strategy to ensure requests are made. Having it
// be a function of the orchestrator makes a lot of sense as they are already online
// and have all the required funds, keys, and rpc servers setup
//
// Exactly how to balance optimizing this versus testing is an interesting discussion
// in testing we want to make sure requests are made without any powers changing on the chain
// just to simplify the test environment. But in production that's somewhat wasteful. What this
// routine does it check the current valset versus the last requested valset, if power has changed
// significantly we send in a request.
func (s *peggyOrchestrator) ValsetRequesterLoop(ctx context.Context) (err error) {
	t := time.NewTimer(0)
	for range t.C {

		latestValsets, err := s.cosmosQueryClient.LatestValsets(ctx)
		if err != nil {
			log.WithError(err).Errorln("unable to get latest valsets from Cosmos chain, retry in", defaultRetryDur)
			t.Reset(defaultRetryDur)
			continue
		}

		// This request gives normalized powers. so calculation wont work
		currentValset, err := s.cosmosQueryClient.CurrentValset(ctx)
		if err != nil {
			log.WithError(err).Errorln("unable to get current valset from Cosmos chain, retry in", defaultRetryDur)
			t.Reset(defaultRetryDur)
			continue
		}

		if len(latestValsets) == 0 {
			if err := s.peggyBroadcastClient.SendValsetRequest(ctx); err != nil {
				log.WithError(err).Warningln("valset request failed")
			}
		} else {
			// if the power difference is more than 1% different than the last valset
			if valPowerDiff(latestValsets[0], currentValset) > 0.01 {
				log.Debugln("power difference is more than 1% different than the last valset. Sending valset request")
				if err := s.peggyBroadcastClient.SendValsetRequest(ctx); err != nil {
					log.WithError(err).Warningln("valset request failed")
				}
			}
		}

		t.Reset(defaultLoopDur)
	}

	return nil
}

func (s *peggyOrchestrator) BatchRequesterLoop(ctx context.Context) (err error) {
	t := time.NewTimer(0)
	for range t.C {
		// get All the denominations
		// check if threshold is met
		// broadcast Request batch

		for coinDenom, tokenAddr := range s.erc20ContractMapping {
			unbatchTxs, err := s.cosmosQueryClient.LatestUnbatchOutgoingTx(ctx, tokenAddr)
			if err == nil && unbatchTxs != nil && len(unbatchTxs) != 0 {
				log.WithFields(log.Fields{
					"coinDenom": coinDenom,
				}).Debugln("Sending unbatchTxs:", unbatchTxs)

				if err := s.peggyBroadcastClient.SendRequestBatch(ctx, coinDenom); err != nil {
					log.WithError(err).Warningln("valset request failed")
				}
			} else if err != nil {
				log.WithError(err).Errorln("failed to get LatestUnbatchOutgoingTx")
			}
		}

		t.Reset(defaultLoopDur)
	}

	return nil
}

func (s *peggyOrchestrator) RelayerMainLoop(ctx context.Context) (err error) {
	r := relayer.NewPeggyRelayer(s.cosmosQueryClient, s.peggyContract)
	return r.Start(ctx)
}

// valPowerDiff returns the difference in power between two bridge validator sets
// TODO: this needs to be potentially refactored
func valPowerDiff(old *types.Valset, new *types.Valset) float64 {
	powers := map[string]int64{}
	var totalB int64
	// loop over b and initialize the map with their powers
	for _, bv := range old.GetMembers() {
		powers[bv.EthereumAddress] = int64(bv.Power)
		totalB += int64(bv.Power)
	}

	// subtract c powers from powers in the map, initializing
	// uninitialized keys with negative numbers
	for _, bv := range new.GetMembers() {
		if val, ok := powers[bv.EthereumAddress]; ok {
			powers[bv.EthereumAddress] = val - int64(bv.Power)
		} else {
			powers[bv.EthereumAddress] = -int64(bv.Power)
		}
	}

	var delta float64
	for _, v := range powers {
		// NOTE: we care about the absolute value of the changes
		delta += math.Abs(float64(v))
	}

	return math.Abs(delta / float64(totalB))
}

func calculateTotalValsetPower(valset *types.Valset) *big.Int {
	totalValsetPower := new(big.Int)
	for _, m := range valset.Members {
		mPower := big.NewInt(0).SetUint64(m.Power)
		totalValsetPower.Add(totalValsetPower, mPower)
	}
	return totalValsetPower
}
