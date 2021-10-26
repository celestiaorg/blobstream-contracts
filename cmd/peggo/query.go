package peggo

import (
	"github.com/spf13/cobra"
)

func getQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Query commands that can get state info from Peggy",
		Long: `Query commands that can get state info from Peggy.

Inputs in the CLI commands can be provided via flags or environment variables. If
using the later, prefix the environment variable with PEGGO_ and the named of the
flag (e.g. PEGGO_COSMOS_PK).`,
	}

	// TODO: Add commands. Injective's Peggo doesn't have any at the moment.
	cmd.AddCommand()

	return cmd
}
