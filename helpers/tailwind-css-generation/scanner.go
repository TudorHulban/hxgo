package tailwindcssgeneration

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type ScanConfig struct {
	Folders []string
	Files   []string

	BuilderFunc      string // start token
	BuilderQualifier string // token builder package
	EndMethod        string // end token
}

func scanSourceForChains(src string, cfg ScanConfig, methodNames map[string]struct{}, used map[string]struct{}) {
	startPatterns := []string{cfg.BuilderFunc + "()"}
	if cfg.BuilderQualifier != "" {
		startPatterns = append(startPatterns, cfg.BuilderQualifier+"."+cfg.BuilderFunc+"()")
	}

	// Use regex to find .AsNode( with possible whitespace
	endPattern := regexp.MustCompile(`\.\s*` + regexp.QuoteMeta(cfg.EndMethod) + `\s*\(`)

	for _, start := range startPatterns {
		searchIdx := 0
		for {
			i := strings.Index(src[searchIdx:], start)
			if i == -1 {
				break
			}

			chainStart := searchIdx + i + len(start)

			// Find end pattern from chainStart onwards
			loc := endPattern.FindStringIndex(src[chainStart:])
			if loc == nil {
				searchIdx = chainStart
				continue
			}

			chainEnd := chainStart + loc[0] // loc[0] is the start of the match

			region := src[chainStart:chainEnd]

			// Split by '.' to get method calls
			parts := strings.Split(region, ".")
			for _, part := range parts {
				part = strings.TrimSpace(part)
				if part == "" {
					continue
				}

				if idx := strings.Index(part, "("); idx != -1 {
					methodName := strings.TrimSpace(part[:idx])
					if methodName != "" {
						if _, ok := methodNames[methodName]; ok {
							used[methodName] = struct{}{}
						}
					}
				}
			}

			searchIdx = chainStart + loc[1] // loc[1] is the end of the match, move past it
		}
	}
}

func ScanUsedMethods(cfg ScanConfig, methods []MethodInfo) ([]string, error) {
	if cfg.BuilderFunc == "" {
		cfg.BuilderFunc = "TW"
	}
	if cfg.EndMethod == "" {
		cfg.EndMethod = "AsNode"
	}
	if cfg.BuilderQualifier == "" {
		cfg.BuilderQualifier = "dsl"
	}

	methodNames := make(map[string]struct{}, len(methods))
	for _, m := range methods {
		methodNames[m.Name] = struct{}{}
	}

	used := make(map[string]struct{})

	// 1. Scan folders
	for _, folder := range cfg.Folders {
		err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if !strings.HasSuffix(path, ".go") {
				return nil
			}

			src, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			scanSourceForChains(string(src), cfg, methodNames, used)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	// 2. Scan explicit files
	for _, file := range cfg.Files {
		if !strings.HasSuffix(file, ".go") {
			continue
		}
		src, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		scanSourceForChains(string(src), cfg, methodNames, used)
	}

	out := make([]string, 0, len(used))
	for name := range used {
		out = append(out, name)
	}
	sort.Strings(out)
	return out, nil
}
