package scan

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
)

func Handlers(dir string) ([]string, error) {
	var handlers []string

	if err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return fmt.Errorf("%s is a directory", path)
		}
		if !strings.HasSuffix(path, ".go") {
			return fmt.Errorf("%s is not a .go file", path)
		}

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return fmt.Errorf("failed to parse file %s: %w", path, err)
		}

		for _, decl := range node.Decls {
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok || funcDecl.Name == nil {
				continue
			}

			if funcDecl.Type.Results == nil || len(funcDecl.Type.Results.List) != 1 {
				continue
			}

			if ident, ok := funcDecl.Type.Results.List[0].Type.(*ast.Ident); ok && strings.HasSuffix(ident.Name, "Handler") {
				handlers = append(handlers, funcDecl.Name.Name)
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return handlers, nil
}
