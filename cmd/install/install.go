package install

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mickamy/gon/internal/config"
	"github.com/mickamy/gon/internal/templates"
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
	if err := copyDatabaseFile(cfg); err != nil {
		fmt.Printf("💥 Failed to generate database file: %v\n", err)
		return err
	}

	if err := copyTestUtilFiles(cfg); err != nil {
		fmt.Printf("💥 Failed to generate test util files: %v\n", err)
		return err
	}

	fmt.Println("📁 Creating templates...")
	if err := copyTemplateFiles(cfg); err != nil {
		fmt.Printf("💥 Failed to create templates: %v\n", err)
		return err
	}

	fmt.Println("📦 Installing gomock...")
	if err := exec.Command("go", "get", "-tool", "github.com/golang/mock/mockgen@latest").Run(); err != nil {
		fmt.Println("⚠️ Failed to install gomock:", err)
	} else {
		fmt.Println("✅ gomock installed successfully.")
	}

	fmt.Println("📦 Installing DI framework...")
	if err := exec.Command("go", "get", "-tool", cfg.DIFramework.InstallPackage()+"@latest").Run(); err != nil {
		fmt.Println("⚠️ Failed to install DI framework:", err)
	} else {
		fmt.Println("✅ DI framework installed successfully.")
	}

	return nil
}

func copyDatabaseFile(cfg *config.Config) error {
	driver := cfg.DBDriver
	path := filepath.Join(cfg.DatabasePackagePath(), driver.String()+".go")
	fmt.Printf("🧱 Generating database file: %s...\n", path)

	if _, err := os.Stat(path); err == nil {
		fmt.Printf("📄 %s already exists. Skipping generation.\n", path)
		return nil
	}

	if err := templates.Copy("defaults/database/"+driver.String()+".go.tmpl", path); err != nil {
		return err
	}

	return nil
}

func copyTestUtilFiles(cfg *config.Config) error {
	testUtilFiles := map[string]string{
		filepath.Join(cfg.TestUtilDir, "httptestutil", "request.go"): "defaults/httptestutil/request.go.tmpl",
	}

	switch cfg.WebFramework {
	case config.WebFrameworkEcho:
		testUtilFiles[filepath.Join(cfg.TestUtilDir, "httptestutil", "echo.go")] = "defaults/httptestutil/echo.go.tmpl"
	}

	for destPath, embedPath := range testUtilFiles {
		if _, err := os.Stat(destPath); err == nil {
			fmt.Printf("📄 %s already exists. Skipping.\n", destPath)
			continue
		}

		if err := templates.Copy(embedPath, destPath); err != nil {
			return err
		}

		fmt.Printf("✅ Created template: %s\n", destPath)
	}

	return nil
}

func copyTemplateFiles(cfg *config.Config) error {
	templateFiles := map[string]func() string{
		cfg.ModelTemplate: func() string {
			return "defaults/model.tmpl"
		},
		cfg.ModelTestTemplate: func() string {
			return "defaults/model_test.tmpl"
		},
		cfg.FixtureTemplate: func() string {
			return "defaults/fixture.tmpl"
		},
		cfg.RepositoryTemplate: func() string {
			return "defaults/repository_" + cfg.DBDriver.String() + ".tmpl"
		},
		cfg.RepositoryTestTemplate: func() string {
			return "defaults/repository_test_" + cfg.DBDriver.String() + ".tmpl"
		},
		cfg.UsecaseTemplate: func() string {
			return "defaults/usecase.tmpl"
		},
		cfg.UsecaseTestTemplate: func() string {
			return "defaults/usecase_test.tmpl"
		},
		cfg.HandlerTemplate: func() string {
			return "defaults/handler_" + cfg.WebFramework.String() + ".tmpl"
		},
		cfg.HandlerTestTemplate: func() string {
			return "defaults/handler_test_" + cfg.WebFramework.String() + ".tmpl"
		},
		cfg.DiTemplate: func() string {
			return "defaults/di_" + cfg.DIFramework.String() + ".tmpl"
		},
	}

	for destPath, embedPathFn := range templateFiles {
		embedPath := embedPathFn()
		if _, err := os.Stat(destPath); err == nil {
			fmt.Printf("📄 %s already exists. Skipping.\n", destPath)
			continue
		}

		if err := templates.Copy(embedPath, destPath); err != nil {
			return fmt.Errorf("failed to copy template src=[%s] dest=[%s]: %w", embedPath, destPath, err)
		}

		fmt.Printf("✅ Created template: %s\n", destPath)
	}

	return nil
}
