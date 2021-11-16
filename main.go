package main

import (
	"os"

	"github.com/celestiaorg/quantum-gravity-bridge/cmd/peggo"
)

func main() {
	cmd := peggo.NewRootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
