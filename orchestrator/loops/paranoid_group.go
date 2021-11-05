package loops

import (
	"sync"
	"time"

	multierror "github.com/hashicorp/go-multierror"
)

// ParanoidGroup is a special primitive to run groups of goroutines, e.g. loops.
// If one of them fails (exits with an error), the wait group will unblock the flow,
// so the user has option to cancel the rest of routines or restart that single routine.
type ParanoidGroup struct {
	initOnce  sync.Once
	closeOnce sync.Once

	wg    *sync.WaitGroup
	errCh chan error
}

func (p *ParanoidGroup) Go(fn func() error) {
	var firstRun bool

	p.initOnce.Do(func() {
		firstRun = true
		p.wg = new(sync.WaitGroup)
		p.wg.Add(1)
		p.errCh = make(chan error, 10)
	})

	if !firstRun {
		p.wg.Add(1)
	}

	go func() {
		defer p.wg.Done()

		if err := fn(); err != nil {
			p.errCh <- err
		}
	}()
}

func (p *ParanoidGroup) Initialized() bool {
	return p.wg != nil
}

// Wait returns an error if one or more tasks failed, otherwise all tasks exited
// with no errors.
func (p *ParanoidGroup) Wait() error {
	p.closeOnce.Do(func() {
		go func() {
			p.wg.Wait()
			close(p.errCh)
		}()
	})

	var errs *multierror.Error

	for {
		select {
		case err, ok := <-p.errCh:
			if !ok {
				// wait group is done!
				return nil
			}

			if errs == nil {
				errs = new(multierror.Error)
			}
			// got new error, don't just exit now, there might be more
			errs = multierror.Append(errs, err)
		default:
			// no new errors in channel, but collected some already
			if errs != nil {
				return errs
			}

			// untight the loop
			time.Sleep(time.Millisecond)
		}
	}
}
