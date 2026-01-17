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
	BuilderFunc      string // start token
	BuilderQualifier string // token builder package
	EndMethod        string // end token

	Folders []string
	Files   []string
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
			ixStart := strings.Index(src[searchIdx:], start)
			if ixStart == -1 {
				break
			}

			chainStart := searchIdx + ixStart + len(start)

			// Find end pattern from chainStart onwards
			loc := endPattern.FindStringIndex(src[chainStart:])
			if loc == nil {
				searchIdx = chainStart

				continue
			}

			chainEnd := chainStart + loc[0] // loc[0] is the start of the match

			region := src[chainStart:chainEnd]

			iterationMethodCalls := strings.SplitSeq(region, ".")

			iterationMethodCalls(
				func(part string) bool {
					part = strings.TrimSpace(part)
					if len(part) == 0 {
						return true
					}

					methodName, _, found := strings.Cut(part, "(")
					if !found {
						return true
					}

					methodName = strings.TrimSpace(methodName)
					if len(methodName) == 0 {
						return true
					}

					if _, exists := methodNames[methodName]; exists {
						used[methodName] = struct{}{}
					}

					return true
				},
			)

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
		errWalk := filepath.WalkDir(
			folder,
			func(path string, d fs.DirEntry, errInput error) error {
				if errInput != nil {
					return errInput
				}

				if d.IsDir() {
					return nil
				}

				if !strings.HasSuffix(path, ".go") {
					return nil
				}

				src, errRead := os.ReadFile(path)
				if errRead != nil {
					return errRead
				}

				scanSourceForChains(string(src), cfg, methodNames, used)

				return nil
			},
		)
		if errWalk != nil {
			return nil, errWalk
		}
	}

	// 2. Scan explicit files
	for _, file := range cfg.Files {
		if !strings.HasSuffix(file, ".go") {
			continue
		}

		src, errRead := os.ReadFile(file)
		if errRead != nil {
			return nil, errRead
		}

		scanSourceForChains(
			string(src),
			cfg,
			methodNames,
			used,
		)
	}

	out := make([]string, 0, len(used))
	for name := range used {
		out = append(out, name)
	}

	sort.Strings(out)

	return out, nil
}
