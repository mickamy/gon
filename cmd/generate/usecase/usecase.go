package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/gon"
)

type TemplateData struct {
	Name              string
	UncapitalizedName string
}

var Cmd = &cobra.Command{
	Use:   "usecase [name]",
	Short: "Generate a usecase to execute business logic",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		return Generate(cfg, args)
	},
}

var domain string

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}

func Generate(cfg *config.Config, args []string) error {
	name := gon.Capitalize(args[0])
	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
		domain = gon.ToSnakeCase(name)
	}

	data := TemplateData{
		Name:              name,
		UncapitalizedName: gon.Uncapitalize(name),
	}

	fmt.Println("📄 Generating usecase file...")
	outPath := filepath.Join(cfg.OutputDir, domain, "usecase", fmt.Sprintf("%s_use_case.go", gon.ToSnakeCase(name)))
	if err := renderToFile(cfg, data, outPath); err != nil {
		return err
	}

	fmt.Println("✅ Usecase file generated successfully.")
	return nil
}

func renderToFile(cfg *config.Config, data TemplateData, outPath string) error {
	b, err := os.ReadFile(cfg.UsecaseTemplate)
	if err != nil {
		return err
	}
	tmplContent := string(b)

	tmpl, err := template.New("usecase").Parse(tmplContent)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}

	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("⚠️ Failed to close file: %v\n", err)
		}
	}(file)

	return tmpl.Execute(file, data)
}
