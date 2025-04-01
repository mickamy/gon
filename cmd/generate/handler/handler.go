package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/gon"
	"github.com/mickamy/gon/internal/templates"
)

type Action struct {
	Name   string
	Entity string
}

type TemplateData struct {
	Actions        []Action
	UsecasePackage string
}

var Cmd = &cobra.Command{
	Use:   "handler [entity] [actions]",
	Short: "Generate a handler to handle HTTP requests",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		return Generate(cfg, args)
	},
}

func Generate(cfg *config.Config, args []string) error {
	entity := gon.Capitalize(args[0])
	actions := args[1:]

	fmt.Println("📄 Generating handler file...")
	if domain == "" {
		fmt.Printf("📂 Domain not specified. Using %s as fallback.\n", entity)
		domain = gon.Uncapitalize(entity)
	}

	data := TemplateData{
		Actions:        parseAction(entity, actions),
		UsecasePackage: cfg.DomainPackage(domain) + "/usecase",
	}

	outPath := filepath.Join(cfg.OutputDir, domain, "handler", fmt.Sprintf("%s_handler.go", strings.ToLower(entity)))
	if err := renderToFile(cfg, data, outPath); err != nil {
		return err
	}

	fmt.Println("✅ Handler file generated successfully.")
	return nil
}

var domain string

func init() {
	Cmd.Flags().StringVar(&domain, "domain", "", "Domain subdirectory (e.g. 'user')")
}

func renderToFile(cfg *config.Config, data TemplateData, outPath string) error {
	tmplFile := cfg.HandlerTemplate

	var bytes []byte
	var err error

	if _, statErr := os.Stat(tmplFile); statErr == nil {
		bytes, err = os.ReadFile(tmplFile)
		if err != nil {
			return fmt.Errorf("⚠️ Failed to read local template %q: %w", tmplFile, err)
		}
	} else {
		bytes, err = templates.DefaultFS.ReadFile(tmplFile)
		if err != nil {
			bytes, err = templates.DefaultFS.ReadFile(tmplFile)
			if err != nil {
				return fmt.Errorf(
					"⚠️ Failed to read embedded template %q: %w\n💡 You might need to run `gon install` to generate default templates.",
					tmplFile, err,
				)
			}
		}
	}
	tmplContent := string(bytes)

	tmpl, err := template.New("handler").Parse(tmplContent)
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
		if err := file.Close(); err != nil {
			fmt.Printf("⚠️ Failed to close file: %v\n", err)
		}
	}(file)

	return tmpl.Execute(file, data)
}

func parseAction(entity string, raw []string) []Action {
	var actions []Action
	for _, item := range raw {
		actions = append(actions, Action{
			Name:   gon.Capitalize(item),
			Entity: entity,
		})
	}
	return actions
}
