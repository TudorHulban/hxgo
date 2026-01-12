package tailwindcssgeneration

import (
	"fmt"
	"os"
	"path/filepath"
)

// CSSWriter persists the reduced CSS content to the configured output path.
// The writer creates parent directories when necessary and overwrites any
// existing file. No additional transformations are applied to the content.
func CSSWriter(outputPath string, css string) error {
	if outputPath == "" {
		return fmt.Errorf("output path is empty")
	}

	dir := filepath.Dir(outputPath)
	if dir != "." {
		err := os.MkdirAll(dir, 0o755)
		if err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	err := os.WriteFile(outputPath, []byte(css), 0o644)
	if err != nil {
		return fmt.Errorf("failed to write css file: %w", err)
	}

	return nil
}
