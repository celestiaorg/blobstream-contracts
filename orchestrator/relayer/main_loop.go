package relayer

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

const defaultLoopDur = 10 * time.Second

func (s *peggyRelayer) Start(ctx context.Context) error {
	t := time.NewTimer(0)
	for range t.C {
		ctx, cancelFn := context.WithTimeout(context.Background(), defaultLoopDur)

		eg, ctx := errgroup.WithContext(ctx)

		eg.Go(func() error {
			return s.RelayValsets(ctx)
		})
		eg.Go(func() error {
			return s.RelayBatches(ctx)
		})

		if err := eg.Wait(); err != nil {
			return err
		}

		cancelFn()
		t.Reset(defaultLoopDur)
	}

	return nil
}
