package init

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

const defaultConfig = `basePackage: github.com/:name/:project
outputDir: internal/domain
defaultDriver: gorm
defaultWeb: echo
`

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a gon.yaml config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := filepath.Join(".", "gon.yaml")

		if _, err := os.Stat(path); err == nil {
			fmt.Println("gon.yaml already exists. Skipping.")
			return nil
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(f)

		tmpl, err := template.New("config").Parse(defaultConfig)
		if err != nil {
			return err
		}

		err = tmpl.Execute(f, nil)
		if err != nil {
			return err
		}

		fmt.Println("✅ gon.yaml created.")
		return nil
	},
}
