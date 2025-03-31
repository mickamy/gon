package init

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

const defaultConfig = `basePackage: github.com/:name/:project
outputDir: internal/domain
defaultDriver: gorm
defaultWeb: echo
databasePackage: github.com/:name/:project/internal/storage/database
`

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a gon.yaml config file and install dependencies",
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

		fmt.Println("📦 Installing gomock...")
		if err := exec.Command("go", "install", "github.com/golang/mock/mockgen@latest").Run(); err != nil {
			fmt.Println("⚠️ Failed to install gomock:", err)
		} else {
			fmt.Println("✅ gomock installed.")
		}

		return nil
	},
}
