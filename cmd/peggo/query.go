package peggo

import (
	"github.com/spf13/cobra"
)

func getQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Query commands that can get state info from Peggy",
	}

	// TODO: Add commands. Injective's Peggo doesn't have any at the moment.
	cmd.AddCommand()

	return cmd
}
