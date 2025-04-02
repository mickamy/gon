package generate

import (
	"github.com/spf13/cobra"

	"github.com/mickamy/gon/cmd/generate/di"
	"github.com/mickamy/gon/cmd/generate/fixture"
	"github.com/mickamy/gon/cmd/generate/handler"
	"github.com/mickamy/gon/cmd/generate/model"
	"github.com/mickamy/gon/cmd/generate/repository"
	"github.com/mickamy/gon/cmd/generate/scaffold"
	"github.com/mickamy/gon/cmd/generate/usecase"
)

var Cmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate code for the application",
	Long:    `Generate code for the application, including models, DTOs, and other components.`,
}

func init() {
	Cmd.AddCommand(di.Cmd)
	Cmd.AddCommand(fixture.Cmd)
	Cmd.AddCommand(handler.Cmd)
	Cmd.AddCommand(repository.Cmd)
	Cmd.AddCommand(model.Cmd)
	Cmd.AddCommand(scaffold.Cmd)
	Cmd.AddCommand(usecase.Cmd)
}
