package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev"

func String() string {
	if Version == "dev" {
		return "dev (built from source; try GitHub release for accurate version)"
	}
	return Version
}

var Cmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Show gon version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gon version %s\n", String())
	},
}
