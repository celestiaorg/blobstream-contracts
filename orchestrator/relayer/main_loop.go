package relayer

import (
	"context"
	"time"

	retry "github.com/avast/retry-go"
	log "github.com/xlab/suplog"

	"github.com/InjectiveLabs/peggo/orchestrator/loops"
)

const defaultLoopDur = 1 * time.Minute

func (s *peggyRelayer) Start(ctx context.Context) error {
	logger := log.WithField("loop", "RelayerMainLoop")

	return loops.RunLoop(ctx, defaultLoopDur, func() error {
		var pg loops.ParanoidGroup
		if s.valsetRelayEnabled {
			logger.Info("Valset Relay Enabled. Starting to relay valsets to Ethereum")
			pg.Go(func() error {
				return retry.Do(func() error {
					return s.RelayValsets(ctx)
				}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
					logger.WithError(err).Warningf("failed to relay Valsets, will retry (%d)", n)
				}))
			})
		}

		if s.batchRelayEnabled {
			logger.Info("Batch Relay Enabled. Starting to relay batches to Ethereum")
			pg.Go(func() error {
				return retry.Do(func() error {
					return s.RelayBatches(ctx)
				}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
					logger.WithError(err).Warningf("failed to relay TxBatches, will retry (%d)", n)
				}))
			})
		}

		if pg.Initialized() {
			if err := pg.Wait(); err != nil {
				logger.WithError(err).Errorln("got error, loop exits")
				return err
			}
		}
		return nil
	})
}
