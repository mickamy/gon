package templates

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func Copy(src string, dest string) error {
	f, err := FS.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open embedded file %q: %w", src, err)
	}
	defer func(f fs.File) {
		_ = f.Close()
	}(f)

	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	out, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create output file %q: %w", dest, err)
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	if _, err := io.Copy(out, f); err != nil {
		return fmt.Errorf("failed to write file %q: %w", dest, err)
	}

	return nil
}
