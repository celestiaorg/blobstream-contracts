package loops

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/pkg/errors"
	log "github.com/xlab/suplog"
)

// ErrGracefulStop is a special error, if returned from within loop function, will stop that loop without
// returning any error
var ErrGracefulStop = errors.New("stop")

// Loop runs a function in the loop with a consistent interval. If execution takes longer,
// the waiting time between iteration decreases. A single iteration has a deadline and cannot run longer
// than interval itself. There is a protection from panic which could crash adjacent loops.
func RunLoop(ctx context.Context, interval time.Duration, fn func() error) (err error) {
	defer panicRecover(&err)

	delayTimer := time.NewTimer(0)
	for {
		select {
		case <-delayTimer.C:
			var start = time.Now()
			if fnErr := fn(); fnErr != nil {
				if fnErr == ErrGracefulStop {
					return nil
				}

				return fnErr
			}

			if elapsed := time.Since(start); elapsed >= interval {
				// in case of an overlap, use just interval
				delayTimer.Reset(interval)
			} else {
				delayTimer.Reset(interval - elapsed)
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func panicRecover(err *error) {
	if r := recover(); r != nil {
		if e, ok := r.(error); ok {
			*err = e

			log.WithError(e).Errorln("loop panicked with an error")
			log.Debugln(string(debug.Stack()))
			return
		}

		*err = errors.Errorf("loop panic: %v", r)
		log.Errorln(*err)
	}
}
