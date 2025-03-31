package model_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mickamy/gon/cmd"
)

func TestGenerateModel(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	oldWd, _ := os.Getwd()
	defer func(dir string) {
		err := os.Chdir(dir)
		if err != nil {
			t.Fatalf("failed to change directory: %v", err)
		}
	}(oldWd)
	_ = os.Chdir(tmp)

	root := cmd.Cmd
	root.SetArgs([]string{"generate", "model", "User", "name:string", "age:int"})

	err := root.Execute()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}

	expectedPath := filepath.Join(tmp, "internal", "domain", "model", "user_model.go")
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
