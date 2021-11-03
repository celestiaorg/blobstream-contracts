package relayer

import (
	"context"
	"time"

	retry "github.com/avast/retry-go"

	"github.com/umee-network/peggo/orchestrator/loops"
)

const defaultLoopDur = 5 * time.Minute

func (s *peggyRelayer) Start(ctx context.Context) error {
	logger := s.logger.With().Str("loop", "RelayerMainLoop").Logger()

	return loops.RunLoop(ctx, defaultLoopDur, func() error {
		var pg loops.ParanoidGroup
		if s.valsetRelayEnabled {
			logger.Info().Msg("valset Relay Enabled. Starting to relay valsets to Ethereum")
			pg.Go(func() error {
				return retry.Do(func() error {
					return s.RelayValsets(ctx)
				}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
					logger.Err(err).Uint("retry", n).Msg("failed to relay Valsets; retrying...")
				}))
			})
		}

		if s.batchRelayEnabled {
			logger.Info().Msg("batch Relay Enabled. Starting to relay batches to Ethereum")
			pg.Go(func() error {
				return retry.Do(func() error {
					return s.RelayBatches(ctx)
				}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
					logger.Err(err).Uint("retry", n).Msg("failed to relay TxBatches; retrying...")
				}))
			})
		}

		if pg.Initialized() {
			if err := pg.Wait(); err != nil {
				logger.Err(err).Msg("got error, loop exits")
				return err
			}
		}
		return nil
	})
}
