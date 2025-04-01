package install

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/gon"
	"github.com/mickamy/gon/templates"
)

var Cmd = &cobra.Command{
	Use:   "install",
	Short: "Generate database file, install dependencies, and prepare templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := writeDatabaseFile(); err != nil {
			fmt.Printf("💥 Failed to generate database file: %v\n", err)
			return err
		}

		fmt.Println("📁 Creating driver-specific templates...")
		if err := writeTemplateFiles(gon.Config.DefaultDriver); err != nil {
			fmt.Printf("⚠️ Failed to create templates: %v\n", err)
		}

		fmt.Println("📦 Installing gomock...")
		if err := exec.Command("go", "install", "github.com/golang/mock/mockgen@latest").Run(); err != nil {
			fmt.Println("⚠️ Failed to install gomock:", err)
		} else {
			fmt.Println("✅ gomock installed successfully.")
		}

		return nil
	},
}

func writeDatabaseFile() error {
	driver := gon.Config.DefaultDriver
	path := filepath.Join(gon.Config.DatabasePackagePath(), driver.String()+".go")
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
	case config.DriverGorm:
		content = gormFileContent
	default:
		return fmt.Errorf("❌ Failed to generate database file: unsupported driver %q", driver)
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}

	return nil
}

func writeTemplateFiles(driver config.Driver) error {
	templateFiles := map[string]func() string{
		gon.Config.ModelTemplate: func() string {
			switch driver {
			case config.DriverGorm:
				return "defaults/model.tmpl"
			default:
				return "defaults/model.tmpl"
			}
		},
		gon.Config.RepositoryTemplate: func() string {
			switch driver {
			case config.DriverGorm:
				return "defaults/repository_gorm.tmpl"
			default:
				return "defaults/repository_gorm.tmpl"
			}
		},
	}

	for destPath, embedPath := range templateFiles {
		if _, err := os.Stat(destPath); err == nil {
			fmt.Printf("📄 %s already exists. Skipping.\n", destPath)
			continue
		}

		bytes, err := templates.DefaultFS.ReadFile(embedPath())
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
