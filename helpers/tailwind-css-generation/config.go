package tailwindcssgeneration

import (
	"fmt"
	"os"
	"os/exec"
)

type TailwindConfig struct {
	CLIPath    string   // e.g. "npx"
	CLIArgs    []string // e.g. {"tailwindcss", "-c", "tailwind.config.js", "-o", "public/tailwind.css"}
	WorkingDir string   // optional: "", or project root
	Env        []string // optional: environment variables
}

// RunTailwind executes the Tailwind CLI to generate the final CSS file.
func RunTailwind(cfg TailwindConfig) error {
	cmd := exec.Command(cfg.CLIPath, cfg.CLIArgs...)
	if cfg.WorkingDir != "" {
		cmd.Dir = cfg.WorkingDir
	}
	if len(cfg.Env) > 0 {
		cmd.Env = append(os.Environ(), cfg.Env...)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// OrchestratorConfig defines the parameters required by the external test
// to run the full Tailwind CSS reduction pipeline. This struct exists only
// to group configuration values in a single, explicit contract.
type OrchestratorConfig struct {
	// ZeroValueDSL is the zero-value instance of the DSL type
	// (for example: dsl.TW()).
	ZeroValueDSL any

	// ScanFolders lists the folders that will be inspected for DSL method usage.
	ScanFolders []string

	// CSSURL is the remote location of the full Tailwind CSS file.
	// The orchestrator performs a simple HTTP GET request to retrieve it.
	CSSURL string

	// OutputPath is the filesystem location where the reduced CSS will be written.
	OutputPath string
}

// Run executes the full pipeline without involving the testing framework.
// This variant is intended for non-test usage or for embedding into custom
// test harnesses. It performs the same steps as GenerateReducedTailwindCSS
// but returns errors instead of using assertions.
func Run(cfg OrchestratorConfig) error {
	// Step 1: Extract DSL methods.
	allMethods, err := ExtractAllMethods(cfg.ZeroValueDSL)
	if err != nil {
		return err
	}
	if len(allMethods) == 0 {
		return fmt.Errorf("no DSL methods extracted")
	}

	// Step 2: Scan source folders.
	usedNames, err := ScanUsedMethods(
		ScanConfig{Folders: cfg.ScanFolders},
		allMethods,
	)
	if err != nil {
		return err
	}

	// Step 3: Resolve used classes.
	usedMethods := ResolveUsedMethods(allMethods, usedNames)
	usedClasses := ClassesFromMethods(usedMethods)
	if len(usedClasses) == 0 {
		return fmt.Errorf("no used classes detected")
	}

	// Step 4: Fetch full CSS.
	fullCSS, err := GETFetcher(cfg.CSSURL)
	if err != nil {
		return err
	}
	if fullCSS == "" {
		return fmt.Errorf("fetched CSS is empty")
	}

	// Step 5: Reduce CSS.
	reduced := CSSReducer(fullCSS, usedClasses)
	if reduced == "" {
		return fmt.Errorf("reduced CSS is empty")
	}

	// Step 6: Write output.
	err = CSSWriter(cfg.OutputPath, reduced)
	if err != nil {
		return err
	}

	return nil
}
