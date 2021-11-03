package peggo

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

func stdinConfirm(msg string) bool {
	var response string

	fmt.Fprint(os.Stderr, msg)

	if _, err := fmt.Scanln(&response); err != nil {
		fmt.Fprintf(os.Stderr, "failed to confirm action: %s\n", err)
		return false
	}

	switch strings.ToLower(strings.TrimSpace(response)) {
	case "y", "yes":
		return true

	default:
		return false
	}
}

func hexToBytes(str string) ([]byte, error) {
	str = strings.TrimPrefix(str, "0x")

	data, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// waitForService awaits an active connection to a gRPC service.
func waitForService(ctx context.Context, clientconn *grpc.ClientConn) {
	for {
		select {
		case <-ctx.Done():
			fmt.Fprintln(os.Stderr, "gRPC service wait timed out")
			os.Exit(1)

		default:
			state := clientconn.GetState()

			if state != connectivity.Ready {
				fmt.Fprintf(os.Stderr, "state of gRPC connection not ready: %s\n", state)
				time.Sleep(5 * time.Second)
				continue
			}

			return
		}
	}
}
