package relayer

import (
	"context"

	retry "github.com/avast/retry-go"

	"github.com/celestiaorg/quantum-gravity-bridge/orchestrator/loops"
)

func (s *peggyRelayer) Start(ctx context.Context) error {
	logger := s.logger.With().Str("loop", "RelayerMainLoop").Logger()

	return loops.RunLoop(ctx, s.loopDuration, func() error {
		var pg loops.ParanoidGroup
		if s.valsetRelayEnabled {
			logger.Info().Msg("valset relay enabled; starting to relay valsets to Ethereum")
			pg.Go(func() error {
				return retry.Do(func() error {
					return s.RelayValsets(ctx)
				}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
					logger.Err(err).Uint("retry", n).Msg("failed to relay valsets; retrying...")
				}))
			})
		}

		//if s.batchRelayEnabled {
		//	logger.Info().Msg("batch relay enabled; starting to relay batches to Ethereum")
		//	pg.Go(func() error {
		//		return retry.Do(func() error {
		//			return s.RelayBatches(ctx)
		//		}, retry.Context(ctx), retry.OnRetry(func(n uint, err error) {
		//			logger.Err(err).Uint("retry", n).Msg("failed to relay tx batches; retrying...")
		//		}))
		//	})
		//}

		if pg.Initialized() {
			if err := pg.Wait(); err != nil {
				logger.Err(err).Msg("main relay loop failed; exiting...")
				return err
			}
		}
		return nil
	})
}
