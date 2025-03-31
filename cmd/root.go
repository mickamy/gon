package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/cmd/generate"
	initCmd "github.com/mickamy/gon/cmd/init"
)

var Cmd = &cobra.Command{
	Use:   "gon",
	Short: "gon is a code generator for Go with Clean Architecture layout",
	Long:  `gon scaffolds models, repositories, usecases, and handlers with opinionated structure.`,
}

func init() {
	Cmd.AddCommand(initCmd.Cmd)
	Cmd.AddCommand(generate.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
