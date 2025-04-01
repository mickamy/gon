package install

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/templates"
)

var Cmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i"},
	Short:   "Generate database file, install dependencies, and prepare templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("⚠️ Failed to load gon.yaml config: %w", err)
		}
		return RunInstall(cfg)
	},
}

func RunInstall(cfg *config.Config) error {
	if err := writeDatabaseFile(cfg); err != nil {
		fmt.Printf("💥 Failed to generate database file: %v\n", err)
		return err
	}

	fmt.Println("📁 Creating driver-specific templates...")
	if err := writeTemplateFiles(cfg); err != nil {
		fmt.Printf("⚠️ Failed to create templates: %v\n", err)
	}

	fmt.Println("📦 Installing gomock...")
	if err := exec.Command("go", "install", "github.com/golang/mock/mockgen@latest").Run(); err != nil {
		fmt.Println("⚠️ Failed to install gomock:", err)
	} else {
		fmt.Println("✅ gomock installed successfully.")
	}

	return nil
}

func writeDatabaseFile(cfg *config.Config) error {
	driver := cfg.DBDriver
	path := filepath.Join(cfg.DatabasePackagePath(), driver.String()+".go")
	fmt.Printf("🧱 Generating database file: %s...\n", path)

	if _, err := os.Stat(path); err == nil {
		fmt.Printf("📄 %s already exists. Skipping generation.\n", path)
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	var content string
	switch driver {
	case config.DBDriverGorm:
		content = gormFileContent
	default:
		return fmt.Errorf("❌ Failed to generate database file: unsupported driver %q", driver)
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}

	return nil
}

func writeTemplateFiles(cfg *config.Config) error {
	templateFiles := map[string]func() string{
		cfg.ModelTemplate: func() string {
			switch cfg.DBDriver {
			case config.DBDriverGorm:
				return "defaults/model.tmpl"
			default:
				return "defaults/model.tmpl"
			}
		},
		cfg.RepositoryTemplate: func() string {
			switch cfg.DBDriver {
			case config.DBDriverGorm:
				return "defaults/repository_gorm.tmpl"
			default:
				return "defaults/repository_gorm.tmpl"
			}
		},
		cfg.UsecaseTemplate: func() string {
			return "defaults/usecase.tmpl"
		},
	}

	for destPath, embedPathFn := range templateFiles {
		embedPath := embedPathFn()
		if _, err := os.Stat(destPath); err == nil {
			fmt.Printf("📄 %s already exists. Skipping.\n", destPath)
			continue
		}

		bytes, err := templates.DefaultFS.ReadFile(embedPath)
		if err != nil {
			fmt.Printf("⚠️ Failed to load embedded template %q: %v\n", embedPath, err)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return err
		}

		if err := os.WriteFile(destPath, bytes, 0644); err != nil {
			return err
		}

		fmt.Printf("✅ Created driver-specific template: %s\n", destPath)
	}

	return nil
}

const gormFileContent = `package database

import (
	"errors"

	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("record not found")

// DB is a wrapper of *gorm.DB
type DB struct{ *gorm.DB }
`
