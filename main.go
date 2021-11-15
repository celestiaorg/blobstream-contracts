package main

import (
	"os"

	"github.com/umee-network/peggo/cmd/peggo"
)

func main() {
	cmd := peggo.NewRootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
