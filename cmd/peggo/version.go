package peggo

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const (
	flagFormat = "format"
)

var (
	// Version defines the application version (defined at compile time)
	Version = ""

	// Commit defines the application commit hash (defined at compile time)
	Commit = ""
)

type versionInfo struct {
	Version string `json:"version" yaml:"version"`
	Commit  string `json:"commit" yaml:"commit"`
	Go      string `json:"go" yaml:"go"`
}

func getVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print binary version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			verInfo := versionInfo{
				Version: Version,
				Commit:  Commit,
				Go:      fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
			}

			format, err := cmd.Flags().GetString(flagFormat)
			if err != nil {
				return err
			}

			var bz []byte

			switch format {
			case "json":
				bz, err = json.Marshal(verInfo)

			default:
				bz, err = yaml.Marshal(&verInfo)
			}

			if err != nil {
				return err
			}

			_, err = fmt.Println(string(bz))
			return err
		},
	}

	versionCmd.Flags().String(flagFormat, "text", "Print the version in the given format (text|json)")

	return versionCmd
}
