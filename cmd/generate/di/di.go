package di

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/caseconv"
	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/scan"
	"github.com/mickamy/gon/internal/templates"
)

type RepositoryData struct {
	Name string
}

type UseCaseData struct {
	Name string
}

type HandlerData struct {
	Name string
}
type TemplateData struct {
	Repositories  []RepositoryData
	UseCases      []UseCaseData
	Handlers      []HandlerData
	DomainPackage string
}

var Cmd = &cobra.Command{
	Use:   "di [domain]",
	Short: "Generate a DI file for dependency injection",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		domain := args[0]
		if err := Generate(cfg, domain); err != nil {
			return fmt.Errorf("⚠️ Failed to generate di file: %w", err)
		}
		fmt.Println("✅ DI file generated successfully.")
		return nil
	},
}

func Generate(cfg *config.Config, domain string) error {
	fs, err := findFiles(domain)
	if err != nil {
		return err
	}

	handlers, err := parseHandlers(fs.handlers)
	if err != nil {
		return err
	}

	data := TemplateData{
		Repositories:  parseRepositories(fs.repositories),
		UseCases:      parseUseCases(fs.usecases),
		Handlers:      handlers,
		DomainPackage: cfg.DomainPackage(domain),
	}

	outPath := filepath.Join(cfg.OutputDir, domain, "di", domain+"_di.go")
	if err := templates.Render(cfg.DiTemplate, data, outPath); err != nil {
		return err
	}

	return nil
}

type diTarget struct {
	repositories []string
	usecases     []string
	handlers     []string
}

func findFiles(domain string) (diTarget, error) {
	var target diTarget
	if err := filepath.Walk("./internal/domain/"+domain, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		if strings.Contains(path, "repository") {
			target.repositories = append(target.repositories, path)
		} else if strings.Contains(path, "usecase") {
			target.usecases = append(target.usecases, path)
		} else if strings.Contains(path, "handler") {
			target.handlers = append(target.handlers, path)
		}
		return nil
	}); err != nil {
		return target, fmt.Errorf("failed to walk through domain directory: %w", err)
	}

	return target, nil
}

func parseRepositories(repositories []string) []RepositoryData {
	var data []RepositoryData
	for _, repo := range repositories {
		base := filepath.Base(repo)
		name := strings.TrimSuffix(base, "_repository.go")
		data = append(data, RepositoryData{Name: caseconv.PascalCase(name)})
	}
	return data
}

func parseUseCases(usecases []string) []UseCaseData {
	var data []UseCaseData
	for _, uc := range usecases {
		base := filepath.Base(uc)
		name := strings.TrimSuffix(base, "_use_case.go")
		data = append(data, UseCaseData{Name: caseconv.PascalCase(name)})
	}
	return data
}

func parseHandlers(handlers []string) ([]HandlerData, error) {
	var data []HandlerData
	for _, handler := range handlers {
		funcs, err := scan.Handlers(handler)
		if err != nil {
			return nil, fmt.Errorf("failed to scan handlers: %w", err)
		}
		for _, f := range funcs {
			data = append(data, HandlerData{Name: f})
		}
	}
	return data, nil
}
