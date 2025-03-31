package generate

import (
	"github.com/spf13/cobra"

	"github.com/mickamy/gon/cmd/generate/model"
)

var Cmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code for the application",
	Long:  `Generate code for the application, including models, DTOs, and other components.`,
}

func init() {
	Cmd.AddCommand(model.Cmd)
}
