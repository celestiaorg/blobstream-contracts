package peggo

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "peggo",
		Short: "Peggo is a companion executable for orchestrating a Peggy validator",
	}

	cmd.PersistentFlags().String(flagLogLevel, zerolog.InfoLevel.String(), "logging level")
	cmd.PersistentFlags().String(flagLogFormat, logLevelJSON, "logging format (text|json)")
	cmd.PersistentFlags().String(flagSvcWaitTimeout, "1m", "Standard wait timeout for external services (e.g. Cosmos daemon gRPC connection)")

	cmd.AddCommand(
		getOrchestratorCmd(),
		getQueryCmd(),
		getTxCmd(),
		getVersionCmd(),
	)

	return cmd
}

func getLogger(cmd *cobra.Command) (zerolog.Logger, error) {
	logLevelStr, err := cmd.Flags().GetString(flagLogLevel)
	if err != nil {
		return zerolog.Logger{}, err
	}

	logLvl, err := zerolog.ParseLevel(logLevelStr)
	if err != nil {
		return zerolog.Logger{}, err
	}

	logFormat, err := cmd.Flags().GetString(flagLogFormat)
	if err != nil {
		return zerolog.Logger{}, err
	}

	var logWriter io.Writer
	switch logFormat {
	case logLevelJSON:
		logWriter = os.Stderr

	case logLevelText:
		logWriter = zerolog.ConsoleWriter{Out: os.Stderr}

	default:
		return zerolog.Logger{}, fmt.Errorf("invalid logging format: %s", logFormat)
	}

	return zerolog.New(logWriter).Level(logLvl).With().Timestamp().Logger(), nil
}

// parseServerConfig returns a server configuration, given a command Context,
// by parsing the following in order of precedence:
//
// - flags
// - environment variables
// - configuration file (TOML) (TODO)
func parseServerConfig(cmd *cobra.Command) (*koanf.Koanf, error) {
	konfig := koanf.New(".")

	// load from file first (if provided)
	// TODO: Support config files if/when needed.
	// if configPath := ctx.String(config.ConfigPath); len(configPath) != 0 {
	// 	if err := konfig.Load(file.Provider(configPath), toml.Parser()); err != nil {
	// 		return nil, err
	// 	}
	// }

	// load from environment variables
	if err := konfig.Load(env.Provider("PEGGO_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(strings.TrimPrefix(s, "PEGGO_")), "_", "-", -1)
	}), nil); err != nil {
		return nil, err
	}

	// finally, load from command line flags
	if err := konfig.Load(posflag.Provider(cmd.Flags(), ".", konfig), nil); err != nil {
		return nil, err
	}

	return konfig, nil
}
