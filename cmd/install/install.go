package install

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/gon"
)

var Cmd = &cobra.Command{
	Use:   "install",
	Short: "Generate database file and install dependencies",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := writeDatabaseFile(); err != nil {
			fmt.Printf("💥 Failed to generate database file: %v\n", err)
			return err
		}

		fmt.Println("📦 Installing gomock...")
		if err := exec.Command("go", "install", "github.com/golang/mock/mockgen@latest").Run(); err != nil {
			fmt.Println("⚠️ Failed to install gomock:", err)
		} else {
			fmt.Println("✅ gomock installed.")
		}

		return nil
	},
}

func writeDatabaseFile() error {
	driver := gon.Config.DefaultDriver
	path := filepath.Join(gon.Config.DatabasePackagePath(), driver.String()+".go")
	fmt.Printf("🧱 Generating %s...\n", path)

	if _, err := os.Stat(path); err == nil {
		fmt.Printf("%s already exists. Skipping.", path)
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
		return fmt.Errorf("❌ cannot generate database file: unsupported driver %q", driver)
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return err
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
