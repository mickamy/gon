package destroy

import (
	"github.com/spf13/cobra"

	"github.com/mickamy/gon/cmd/destroy/handler"
	"github.com/mickamy/gon/cmd/destroy/model"
	"github.com/mickamy/gon/cmd/destroy/repository"
	"github.com/mickamy/gon/cmd/destroy/scaffold"
	"github.com/mickamy/gon/cmd/destroy/usecase"
)

var Cmd = &cobra.Command{
	Use:     "destroy",
	Aliases: []string{"d"},
	Short:   "Destroy code for the application",
	Long:    `Destroy code for the application, including models, DTOs, and other components.`,
}

func init() {
	Cmd.AddCommand(handler.Cmd)
	Cmd.AddCommand(model.Cmd)
	Cmd.AddCommand(repository.Cmd)
	Cmd.AddCommand(scaffold.Cmd)
	Cmd.AddCommand(usecase.Cmd)
}
