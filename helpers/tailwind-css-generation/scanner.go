package tailwindcssgeneration

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ScanConfig defines the folders that will be scanned for DSL method usage.
type ScanConfig struct {
	Folders []string
}

// ScanUsedMethods walks the configured folders, inspects all .go files,
// and detects occurrences of DSL method calls based on the known method names.
// The scanner is fully agnostic to the DSL implementation and CSS source.
// It only performs static text matching of ".MethodName(" patterns.
func ScanUsedMethods(cfg ScanConfig, allMethods []MethodInfo) (map[string]bool, error) {
	used := make(map[string]bool)

	// Build a list of method names for efficient scanning.
	methodNames := make([]string, 0, len(allMethods))
	for _, m := range allMethods {
		methodNames = append(methodNames, m.Name)
	}

	for _, folder := range cfg.Folders {
		err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if !strings.HasSuffix(path, ".go") {
				return nil
			}

			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			src := string(data)

			// Detect method usage by simple substring matching.
			for _, name := range methodNames {
				pattern := regexp.MustCompile(`\.` + regexp.QuoteMeta(name) + `\s*\(`)
				if pattern.MatchString(src) {
					used[name] = true
				}
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return used, nil
}
