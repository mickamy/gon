package generate

import (
	"github.com/spf13/cobra"

	"github.com/mickamy/gon/cmd/generate/model"
	"github.com/mickamy/gon/cmd/generate/repository"
	"github.com/mickamy/gon/cmd/generate/scaffold"
)

var Cmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate code for the application",
	Long:    `Generate code for the application, including models, DTOs, and other components.`,
}

func init() {
	Cmd.AddCommand(repository.Cmd)
	Cmd.AddCommand(model.Cmd)
	Cmd.AddCommand(scaffold.Cmd)
}
