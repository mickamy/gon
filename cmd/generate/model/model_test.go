package model_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mickamy/gon/cmd/generate/model"
	"github.com/mickamy/gon/internal/config"
)

func TestGenerateModel(t *testing.T) {
	t.Parallel()
	tmp := t.TempDir()

	cfg := config.New(config.Config{
		BasePackage:        "example.com/test/project",
		OutputDir:          "internal/domain",
		DBDriver:           config.DBDriverGorm,
		WebFramework:       config.WebFrameworkEcho,
		DatabasePackage:    "example.com/test/project/internal/infra/storage/database",
		ModelTemplate:      "defaults/model.tmpl",
		RepositoryTemplate: "defaults/repository_gorm.tmpl",
	})

	args := []string{"User", "name:string", "age:int"}

	// change working directory to temp
	oldWd, _ := os.Getwd()
	defer func() {
		_ = os.Chdir(oldWd)
	}()
	_ = os.Chdir(tmp)

	err := model.Generate(cfg, args, "user")
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}

	expectedPath := filepath.Join(tmp, "internal", "domain", "user", "model", "user_model.go")
	data, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("expected file not found: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "type User struct") {
		t.Errorf("generated model missing type definition")
	}
	if !strings.Contains(content, "Name string") || !strings.Contains(content, "Age int") {
		t.Errorf("generated model missing fields")
	}
}
