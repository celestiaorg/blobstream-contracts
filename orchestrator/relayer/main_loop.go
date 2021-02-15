package relayer

import (
	"context"
	"sync"
	"time"

	log "github.com/xlab/suplog"
)

const defaultLoopDur = 10 * time.Second

func (s *peggyRelayer) RunLoop() {
	t := time.NewTimer(0)
	for range t.C {
		ctx, cancelFn := context.WithTimeout(context.Background(), defaultLoopDur)

		wg := new(sync.WaitGroup)

		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := s.relayValsets(ctx); err != nil {
				log.Errorln(err)
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := s.relayBatches(ctx); err != nil {
				log.Errorln(err)
			}
		}()

		wg.Wait()
		cancelFn()
		t.Reset(defaultLoopDur)
	}
}
