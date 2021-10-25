package main

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
	log "github.com/xlab/suplog"

	"github.com/umee-network/peggo/orchestrator/version"
)

var app = cli.App("peggo", "Peggo is a companion executable for orchestrating a Peggy validator.")

var (
	envName        *string
	appLogLevel    *string
	svcWaitTimeout *string
)

func main() {
	readEnv()
	initGlobalOptions(
		&envName,
		&appLogLevel,
		&svcWaitTimeout,
	)

	app.Before = func() {
		log.DefaultLogger.SetLevel(logLevel(*appLogLevel))
	}

	app.Command("orchestrator", "Starts the orchestrator main loop.", orchestratorCmd)
	app.Command("q query", "Query commands that can get state info from Peggy.", queryCmdSubset)
	app.Command("tx", "Transactions for Peggy governance and maintenance.", txCmdSubset)
	app.Command("version", "Print the version information and exit.", versionCmd)

	_ = app.Run(os.Args)
}

func versionCmd(c *cli.Cmd) {
	c.Action = func() {
		fmt.Println(version.Version())
	}
}
