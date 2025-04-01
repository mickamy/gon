package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/cmd/destroy"
	"github.com/mickamy/gon/cmd/generate"
	initCmd "github.com/mickamy/gon/cmd/init"
	"github.com/mickamy/gon/cmd/install"
	"github.com/mickamy/gon/cmd/version"
)

var (
	versionFlag bool
	formatOpt   string
)

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
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if err := runFormatting(formatOpt); err != nil {
			fmt.Printf("⚠️ Failed to format code: %v\n", err)
		}
	},
}

func init() {
	Cmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "Print the version and exit")
	Cmd.PersistentFlags().StringVar(&formatOpt, "format", "none", "Format generated code (none, gofmt, goimports, all)")

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

func runFormatting(opt string) error {
	switch opt {
	case "gofmt":
		return runCommand("go", "fmt", "./...")
	case "goimports":
		return runCommand("goimports", "-w", ".")
	case "all":
		if err := runCommand("go", "fmt", "./..."); err != nil {
			return err
		}
		return runCommand("goimports", "-w", ".")
	case "none":
		return nil
	default:
		return fmt.Errorf("unknown format option: %s", opt)
	}
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
