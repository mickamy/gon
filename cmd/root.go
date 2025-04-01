package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/cmd/destroy"
	"github.com/mickamy/gon/cmd/generate"
	initCmd "github.com/mickamy/gon/cmd/init"
	"github.com/mickamy/gon/cmd/install"
	"github.com/mickamy/gon/cmd/version"
)

var versionFlag bool

var Cmd = &cobra.Command{
	Use:   "gon",
	Short: "gon is a code generator for Go with Clean Architecture layout",
	Long:  `gon scaffolds models, repositories, usecases, and handlers with opinionated structure.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Printf("gon version %s\n", version.Version)
			os.Exit(0)
		}
	},
}

func init() {
	Cmd.AddCommand(destroy.Cmd)
	Cmd.AddCommand(generate.Cmd)
	Cmd.AddCommand(initCmd.Cmd)
	Cmd.AddCommand(install.Cmd)
	Cmd.AddCommand(version.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
