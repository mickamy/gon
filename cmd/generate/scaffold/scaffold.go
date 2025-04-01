package scaffold

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/cmd/generate/model"
	"github.com/mickamy/gon/cmd/generate/repository"
	"github.com/mickamy/gon/cmd/generate/usecase"
	"github.com/mickamy/gon/internal/gon"
)

var domain string

var Cmd = &cobra.Command{
	Use:   "scaffold [name] [fields]",
	Short: "Generate model, repository, usecase, and handler for a domain entity",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fields := args[1:]

		if domain == "" {
			fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
			domain = name
		}

		capitalizedName := gon.Capitalize(name)
		subcommands := []struct {
			cmd  *cobra.Command
			args []string
		}{
			{model.Cmd, append([]string{name}, fields...)},
			{repository.Cmd, []string{name}},
			{usecase.Cmd, []string{"Get" + capitalizedName}},
			{usecase.Cmd, []string{"Create" + capitalizedName}},
			{usecase.Cmd, []string{"Update" + capitalizedName}},
			{usecase.Cmd, []string{"Delete" + capitalizedName}},
		}

		for _, sub := range subcommands {
			_ = sub.cmd.Flags().Set("domain", domain)
			if sub.cmd.RunE == nil {
				continue
			}
			if err := sub.cmd.RunE(sub.cmd, sub.args); err != nil {
				return fmt.Errorf("failed to execute %s: %w", sub.cmd.Use, err)
			}
		}

		fmt.Println("🎉 Scaffold complete.")
		return nil
	},
}

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}
