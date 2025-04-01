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
databasePackage: github.com/:name/:project/internal/infra/storage/database
`

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a gon.yaml config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := writerConfigFile(); err != nil {
			fmt.Printf("💥 Failed to generate config file: %v\n", err)
			return err
		}

		return nil
	},
}

func writerConfigFile() error {
	path := filepath.Join(".", "gon.yaml")

	if _, err := os.Stat(path); err == nil {
		fmt.Println("📄 gon.yaml already exists. Skipping generation.")
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

	fmt.Println("✅ Config file gon.yaml created successfully.")

	return nil
}
