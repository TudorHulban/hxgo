package tailwindcssgeneration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// GenerateReducedTailwindCSS is the orchestration entry point used by the
// external test. It performs the complete pipeline:
//
//  1. Extract all DSL methods.
//  2. Scan source folders for method usage.
//  3. Resolve used methods and classes.
//  4. Fetch the full CSS via HTTP GET.
//  5. Reduce the CSS to only the used classes.
//  6. Write the reduced CSS to the configured output path.
//
// The function is intentionally linear and explicit to preserve clarity
// and deterministic behavior.
func GenerateReducedTailwindCSS(
	t *testing.T,
	tailwindZero any,
	scanFolders []string,
	cssURL string,
	outputPath string,
) {
	t.Helper()

	// Step 1: Extract DSL methods.
	allMethods, err := ExtractAllMethods(tailwindZero)
	require.NoError(t, err)
	require.NotEmpty(t, allMethods)

	// Step 2: Scan source folders.
	usedNames, err := ScanUsedMethods(
		ScanConfig{Folders: scanFolders},
		allMethods,
	)
	require.NoError(t, err)

	// Step 3: Resolve used classes.
	usedMethods := ResolveUsedMethods(allMethods, usedNames)
	usedClasses := ClassesFromMethods(usedMethods)
	require.NotEmpty(t, usedClasses)

	// Step 4: Fetch full CSS.
	fullCSS, err := GETFetcher(cssURL)
	require.NoError(t, err)
	require.NotEmpty(t, fullCSS)

	// Step 5: Reduce CSS.
	reduced := CSSReducer(fullCSS, usedClasses)
	require.NotEmpty(t, reduced)

	// Step 6: Write output.
	err = CSSWriter(outputPath, reduced)
	require.NoError(t, err)
}
