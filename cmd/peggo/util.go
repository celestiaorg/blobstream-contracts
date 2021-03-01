package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	ethcmn "github.com/ethereum/go-ethereum/common"
	log "github.com/xlab/suplog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// readEnv is a special utility that reads `.env` file into actual environment variables
// of the current app, similar to `dotenv` Node package.
func readEnv() {
	if envdata, _ := ioutil.ReadFile(".env"); len(envdata) > 0 {
		s := bufio.NewScanner(bytes.NewReader(envdata))
		for s.Scan() {
			parts := strings.Split(s.Text(), "=")
			if len(parts) != 2 {
				continue
			}
			strValue := strings.Trim(parts[1], `"`)
			if err := os.Setenv(parts[0], strValue); err != nil {
				log.WithField("name", parts[0]).WithError(err).Warningln("failed to override ENV variable")
			}
		}
	}
}

// stdinConfirm checks the user's confirmation, if not forced to Yes
func stdinConfirm(msg string) bool {
	var response string

	fmt.Print(msg)

	if _, err := fmt.Scanln(&response); err != nil {
		log.WithError(err).Errorln("failed to confirm the action")
		return false
	}

	switch strings.ToLower(strings.TrimSpace(response)) {
	case "y", "yes":
		return true
	default:
		return false
	}
}

// parseERC20ContractMapping converts list of denom:address pairs to a proper typed map.
func parseERC20ContractMapping(items []string) map[string]ethcmn.Address {
	res := make(map[string]ethcmn.Address)

	for _, item := range items {
		// item is a pair denom:address
		parts := strings.Split(item, ":")
		addr := ethcmn.HexToAddress(parts[1])

		if len(parts) != 2 || len(parts[0]) == 0 || addr == (ethcmn.Address{}) {
			log.Fatalln("failed to parse ERC20 mapping: check that all inputs contain valid denom:address pairs")
		}

		res[parts[0]] = addr
	}

	return res
}

// logLevel converts vague log level name into typed level.
func logLevel(s string) log.Level {
	switch s {
	case "1", "error":
		return log.ErrorLevel
	case "2", "warn":
		return log.WarnLevel
	case "3", "info":
		return log.InfoLevel
	case "4", "debug":
		return log.DebugLevel
	default:
		return log.FatalLevel
	}
}

// toBool is used to parse vague bool definition into typed bool.
func toBool(s string) bool {
	switch strings.ToLower(s) {
	case "true", "1", "t", "yes":
		return true
	default:
		return false
	}
}

// duration parses duration from string with a provided default fallback.
func duration(s string, defaults time.Duration) time.Duration {
	dur, err := time.ParseDuration(s)
	if err != nil {
		dur = defaults
	}
	return dur
}

// checkStatsdPrefix ensures that the statsd prefix really
// have "." at end.
func checkStatsdPrefix(s string) string {
	if !strings.HasSuffix(s, ".") {
		return s + "."
	}
	return s
}

// waitForService awaits an active ClientConn to a GRPC service.
func waitForService(ctx context.Context, clientconn *grpc.ClientConn) {
	for {
		select {
		case <-ctx.Done():
			log.Fatalln("GRPC service wait timed out")
		default:
			state := clientconn.GetState()

			if state != connectivity.Ready {
				log.WithField("state", state.String()).Warningln("state of GRPC connection not ready")
				time.Sleep(5 * time.Second)
				continue
			}

			return
		}
	}
}

// orShutdown fatals the app if there was an error.
func orShutdown(err error) {
	if err != nil && err != grpc.ErrServerStopped {
		log.WithError(err).Fatalln("unable to start peggo orchestrator")
	}
}
