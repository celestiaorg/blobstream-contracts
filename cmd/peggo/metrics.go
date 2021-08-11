package main

import (
	"os"
	"time"

	"github.com/InjectiveLabs/peggo/orchestrator/metrics"
	cli "github.com/jawher/mow.cli"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"
)

func initMetrics(c *cli.Cmd) {
	var (
		statsdPrefix   *string
		statsdAddr     *string
		statsdStuckDur *string
		statsdMocking  *string
		statsdDisabled *string
	)

	initStatsdOptions(
		c,
		&statsdPrefix,
		&statsdAddr,
		&statsdStuckDur,
		&statsdMocking,
		&statsdDisabled,
	)

	if toBool(*statsdDisabled) {
		// initializes statsd client with a mock one with no-op enabled
		metrics.Disable()
	} else {
		go func() {
			for {
				hostname, _ := os.Hostname()
				err := metrics.Init(*statsdAddr, checkStatsdPrefix(*statsdPrefix), &metrics.StatterConfig{
					EnvName:              *envName,
					HostName:             hostname,
					StuckFunctionTimeout: duration(*statsdStuckDur, 30*time.Minute),
					MockingEnabled:       toBool(*statsdMocking) || *envName == "local",
				})
				if err != nil {
					log.WithError(err).Warningln("metrics init failed, will retry in 1 min")
					time.Sleep(time.Minute)
					continue
				}
				break
			}
			closer.Bind(func() {
				metrics.Close()
			})
		}()
	}
}
