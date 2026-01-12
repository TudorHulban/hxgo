package tailwindcssgeneration_test

import (
	"testing"

	"github.com/TudorHulban/hxgo/dsl"
	tailwindcssgeneration "github.com/TudorHulban/hxgo/helpers/tailwind-css-generation"
	"github.com/stretchr/testify/require"
)

func Test_GenerateReducedTailwindCSS(t *testing.T) {
	// Configuration for the test-driven Tailwind CSS reduction pipeline.
	zero := dsl.TW() // zero-value DSL instance

	scanFolders := []string{
		"./cmd",
		"./internal",
		"./ui",
	}

	// Versioned CDN URL for Tailwind CSS.
	cssURL := "https://cdn.jsdelivr.net/npm/tailwindcss@3.4.1/dist/tailwind.min.css"

	// Output location for the reduced CSS file.
	outputPath := "./tailwind.reduced.css"

	// Execute the full pipeline.
	tailwindcssgeneration.GenerateReducedTailwindCSS(
		t,
		zero,
		scanFolders,
		cssURL,
		outputPath,
	)

	// Sanity check: the file must exist after generation.
	require.FileExists(t, outputPath)
}
