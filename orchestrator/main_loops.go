package orchestrator

import (
	"context"
	"errors"
	"math"
	"math/big"
	"time"

	"github.com/avast/retry-go"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/umee-network/peggo/orchestrator/cosmos"
	"github.com/umee-network/peggo/orchestrator/loops"
	"github.com/umee-network/umee/x/peggy/types"
	log "github.com/xlab/suplog"
)

const defaultLoopDur = 60 * time.Second

// Start combines the all major roles required to make
// up the Orchestrator, all of these are async loops.
func (p *peggyOrchestrator) Start(ctx context.Context) error {
	var pg loops.ParanoidGroup

	pg.Go(func() error {
		return p.EthOracleMainLoop(ctx)
	})
	pg.Go(func() error {
		return p.BatchRequesterLoop(ctx)
	})
	pg.Go(func() error {
		return p.EthSignerMainLoop(ctx)
	})
	pg.Go(func() error {
		return p.RelayerMainLoop(ctx)
	})

	return pg.Wait()
}

// EthOracleMainLoop is responsible for making sure that Ethereum events are retrieved from the Ethereum blockchain
// and ferried over to Cosmos where they will be used to issue tokens or process batches.
//
// TODO this loop requires a method to bootstrap back to the correct event nonce when restarted
func (p *peggyOrchestrator) EthOracleMainLoop(ctx context.Context) (err error) {
	logger := log.WithField("loop", "EthOracleMainLoop")
	lastResync := time.Now()
	var lastCheckedBlock uint64

	if err := retry.Do(func() (err error) {
		lastCheckedBlock, err = p.GetLastCheckedBlock(ctx)
		if lastCheckedBlock == 0 {
			peggyParams, err := p.cosmosQueryClient.PeggyParams(ctx)
			if err != nil {
				log.WithError(err).Fatalln("failed to query peggy params, is injectived running?")
			}
			lastCheckedBlock = peggyParams.BridgeContractStartHeight
		}
		return
	}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
		logger.WithError(err).Warningf("failed to get last checked block, will retry (%d)", n)
	})); err != nil {
		logger.WithError(err).Errorln("got error, loop exits")
		return err
	}

	logger.WithField("lastCheckedBlock", lastCheckedBlock).Infoln("Start scanning for events")

	return loops.RunLoop(ctx, defaultLoopDur, func() error {
		// Relays events from Ethereum -> Cosmos
		var currentBlock uint64
		if err := retry.Do(func() (err error) {
			currentBlock, err = p.CheckForEvents(ctx, lastCheckedBlock)
			return
		}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
			logger.WithError(err).Warningf("error during Eth event checking, will retry (%d)", n)
		})); err != nil {
			logger.WithError(err).Errorln("got error, loop exits")
			return err
		}

		lastCheckedBlock = currentBlock

		/*
			Auto re-sync to catch up the nonce. Reasons why event nonce fall behind.
				1. It takes some time for events to be indexed on Ethereum. So if peggo queried events immediately as block produced, there is a chance the event is missed.
				   we need to re-scan this block to ensure events are not missed due to indexing delay.
				2. if validator was in UnBonding state, the claims broadcasted in last iteration are failed.
				3. if infura call failed while filtering events, the peggo missed to broadcast claim events occured in last iteration.
		**/
		if time.Since(lastResync) >= 48*time.Hour {
			if err := retry.Do(func() (err error) {
				lastCheckedBlock, err = p.GetLastCheckedBlock(ctx)
				return
			}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
				logger.WithError(err).Warningf("failed to get last checked block, will retry (%d)", n)
			})); err != nil {
				logger.WithError(err).Errorln("got error, loop exits")
				return err
			}
			lastResync = time.Now()
			logger.WithFields(log.Fields{"lastResync": lastResync, "lastCheckedBlock": lastCheckedBlock}).Infoln("Auto resync")
		}

		return nil
	})
}

// EthSignerMainLoop simply signs off on any batches or validator sets provided by the validator
// since these are provided directly by a trusted Cosmsos node they can simply be assumed to be
// valid and signed off on.
func (p *peggyOrchestrator) EthSignerMainLoop(ctx context.Context) (err error) {
	logger := log.WithField("loop", "EthSignerMainLoop")

	var peggyID common.Hash
	if err := retry.Do(func() (err error) {
		peggyID, err = p.peggyContract.GetPeggyID(ctx, p.peggyContract.FromAddress())
		return
	}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
		logger.WithError(err).Warningf("failed to get PeggyID from Ethereum contract, will retry (%d)", n)
	})); err != nil {
		logger.WithError(err).Errorln("got error, loop exits")
		return err
	}
	logger.Debugf("received peggyID %s", peggyID.Hex())

	return loops.RunLoop(ctx, defaultLoopDur, func() error {
		var oldestUnsignedValsets []*types.Valset
		if err := retry.Do(func() error {
			oldestValsets, err := p.cosmosQueryClient.OldestUnsignedValsets(ctx, p.peggyBroadcastClient.AccFromAddress())
			if err != nil {
				if err == cosmos.ErrNotFound || oldestValsets == nil {
					logger.Debugln("no Valset waiting to be signed")
					return nil
				}

				return err
			}
			oldestUnsignedValsets = oldestValsets
			return nil
		}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
			logger.WithError(err).Warningf("failed to get unsigned Valset for signing, will retry (%d)", n)
		})); err != nil {
			logger.WithError(err).Errorln("got error, loop exits")
			return err
		}

		for _, oldestValset := range oldestUnsignedValsets {
			logger.Infoln("Sending Valset confirm for %d", oldestValset.Nonce)
			if err := retry.Do(func() error {
				return p.peggyBroadcastClient.SendValsetConfirm(ctx, p.ethFrom, peggyID, oldestValset)
			}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
				logger.WithError(err).Warningf("failed to sign and send Valset confirmation to Cosmos, will retry (%d)", n)
			})); err != nil {
				logger.WithError(err).Errorln("got error, loop exits")
				return err
			}
		}

		var oldestUnsignedTransactionBatch *types.OutgoingTxBatch
		if err := retry.Do(func() error {
			// sign the last unsigned batch, TODO check if we already have signed this
			txBatch, err := p.cosmosQueryClient.OldestUnsignedTransactionBatch(ctx, p.peggyBroadcastClient.AccFromAddress())
			if err != nil {
				if err == cosmos.ErrNotFound || txBatch == nil {
					logger.Debugln("no TransactionBatch waiting to be signed")
					return nil
				}
				return err
			}
			oldestUnsignedTransactionBatch = txBatch
			return nil
		}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
			logger.WithError(err).Warningf("failed to get unsigned TransactionBatch for signing, will retry (%d)", n)
		})); err != nil {
			logger.WithError(err).Errorln("got error, loop exits")
			return err
		}

		if oldestUnsignedTransactionBatch != nil {
			logger.Infoln("Sending TransactionBatch confirm for BatchNonce %d", oldestUnsignedTransactionBatch.BatchNonce)
			if err := retry.Do(func() error {
				return p.peggyBroadcastClient.SendBatchConfirm(ctx, p.ethFrom, peggyID, oldestUnsignedTransactionBatch)
			}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
				logger.WithError(err).Warningf("failed to sign and send TransactionBatch confirmation to Cosmos, will retry (%d)", n)
			})); err != nil {
				logger.WithError(err).Errorln("got error, loop exits")
				return err
			}
		}
		return nil
	})
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

/*
Not required any more. The valset request are generated in endblocker of peggy module automatically. Also MsgSendValsetRequest is removed on peggy module.

func (p *peggyOrchestrator) ValsetRequesterLoop(ctx context.Context) (err error) {
	logger := log.WithField("loop", "ValsetRequesterLoop")

	return loops.RunLoop(ctx, defaultLoopDur, func() error {
		var latestValsets []*types.Valset
		var currentValset *types.Valset

		var pg loops.ParanoidGroup

		pg.Go(func() error {
			return retry.Do(func() (err error) {
				latestValsets, err = p.cosmosQueryClient.LatestValsets(ctx)
				return
			}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
				logger.WithError(err).Warningf("failed to get latest valsets, will retry (%d)", n)
			}))
		})

		pg.Go(func() error {
			return retry.Do(func() (err error) {
				currentValset, err = p.cosmosQueryClient.CurrentValset(ctx)
				return
			}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
				logger.WithError(err).Warningf("failed to get current valset, will retry (%d)", n)
			}))
		})

		if err := pg.Wait(); err != nil {
			logger.WithError(err).Errorln("got error, loop exits")
			return err
		}

		if len(latestValsets) == 0 {
			retry.Do(func() error {
				return p.peggyBroadcastClient.SendValsetRequest(ctx)
			}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
				logger.WithError(err).Warningf("failed to request Valset to be formed, will retry (%d)", n)
			}))
		} else {
			// if the power difference is more than 1% different than the last valset
			if valPowerDiff(latestValsets[0], currentValset) > 0.01 {
				log.Debugln("power difference is more than 1%% different than the last valset. Sending valset request")

				retry.Do(func() error {
					return p.peggyBroadcastClient.SendValsetRequest(ctx)
				}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
					logger.WithError(err).Warningf("failed to request Valset to be formed, will retry (%d)", n)
				}))
			}
		}

		return nil
	})
}
**/

func (p *peggyOrchestrator) BatchRequesterLoop(ctx context.Context) (err error) {
	logger := p.logger.With().Str("loop", "BatchRequesterLoop").Logger()

	return loops.RunLoop(ctx, defaultLoopDur, func() error {
		// Each loop performs the following:
		//
		// - get All the denominations
		// - check if threshold is met
		// - broadcast Request batch
		var pg loops.ParanoidGroup

		pg.Go(func() error {
			var unbatchedTokensWithFees []*types.BatchFees

			if err := retry.Do(func() (err error) {
				unbatchedTokensWithFees, err = p.cosmosQueryClient.UnbatchedTokensWithFees(ctx)
				return
			}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
				logger.Err(err).Uint("retry", n).Msg("failed to get UnbatchedTokensWithFees; retrying...")
			})); err != nil {
				// non-fatal, just alert
				logger.Warn().Msg("unable to get UnbatchedTokensWithFees for the token")
				return nil
			}

			if len(unbatchedTokensWithFees) > 0 {
				logger.Debug().Msg("checking if token fees meets set threshold amount and send batch request")
				for _, unbatchedToken := range unbatchedTokensWithFees {
					return retry.Do(func() (err error) {
						// Check if the token is present in cosmos denom. If so, send batch
						// request with cosmosDenom.
						tokenAddr := ethcmn.HexToAddress(unbatchedToken.Token)

						var denom string
						resp, err := p.cosmosQueryClient.ERC20ToDenom(ctx, tokenAddr)
						if err != nil {
							logger.Err(err).Str("token_contract", tokenAddr.String()).Msg("failed to get denom, won't request for a batch")
							// do not return error, just continue with the next unbatched tx.
							return nil
						}

						denom = resp.GetDenom()

						// send batch request only if fee threshold is met
						if p.CheckFeeThreshold(tokenAddr, unbatchedToken.TotalFees, p.minBatchFeeUSD) {
							logger.Info().Str("token_contract", tokenAddr.String()).Str("denom", denom).Msg("sending batch request")
							_ = p.peggyBroadcastClient.SendRequestBatch(ctx, denom)
						}

						return nil
					}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
						logger.Err(err).Uint("retry", n).Msg("failed to get LatestUnbatchOutgoingTx; retrying...")
					}))
				}
			} else {
				logger.Debug().Msg("no outgoing withdraw tx or unbatched token fee less than threshold")
			}

			return nil
		})

		return pg.Wait()
	})
}

func (p *peggyOrchestrator) CheckFeeThreshold(erc20Contract common.Address, totalFee cosmtypes.Int, minFeeInUSD float64) bool {
	if minFeeInUSD == 0 || p.priceFeeder == nil {
		return true
	}

	tokenPriceInUSD, err := p.priceFeeder.QueryUSDPrice(erc20Contract)
	if err != nil {
		return false
	}

	tokenPriceInUSDDec := decimal.NewFromFloat(tokenPriceInUSD)
	totalFeeInUSDDec := decimal.NewFromBigInt(totalFee.BigInt(), -18).Mul(tokenPriceInUSDDec)
	minFeeInUSDDec := decimal.NewFromFloat(minFeeInUSD)

	if totalFeeInUSDDec.GreaterThan(minFeeInUSDDec) {
		return true
	}
	return false
}

func (p *peggyOrchestrator) RelayerMainLoop(ctx context.Context) (err error) {
	if p.relayer != nil {
		return p.relayer.Start(ctx)
	} else {
		return errors.New("relayer is nil")
	}
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
