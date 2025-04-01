package repository

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/gon"
	"github.com/mickamy/gon/templates"
)

type TemplateData struct {
	EntityName      string
	LowerEntityName string
	BasePackage     string
	DatabasePackage string
	DomainName      string
}

var Cmd = &cobra.Command{
	Use:   "repository [name] [fields]",
	Short: "Generate a repository to retrieve domain model",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := gon.Capitalize(args[0])
		if domain == "" {
			fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", name)
		}

		data := TemplateData{
			EntityName:      name,
			LowerEntityName: gon.Uncapitalize(name),
			BasePackage:     gon.Config.BasePackage,
			DatabasePackage: gon.Config.DatabasePackage,
			DomainName:      gon.Uncapitalize(domain),
		}

		fmt.Println("📄 Generating repository file...")
		if err := renderToFile(data, filepath.Join("internal", "domain", domain, "repository", fmt.Sprintf("%s_repository.go", strings.ToLower(name)))); err != nil {
			return err
		}

		fmt.Println("✅ Repository file generated successfully.")
		return nil
	},
}

var domain string

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}

func renderToFile(data TemplateData, outPath string) error {
	b, err := templates.FS.ReadFile("repository.tmpl")
	if err != nil {
		return err
	}
	tmplContent := string(b)

	tmpl, err := template.New("repository").Parse(tmplContent)
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
