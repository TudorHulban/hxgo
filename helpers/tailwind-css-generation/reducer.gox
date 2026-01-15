package tailwindcssgeneration

import (
	"regexp"
	"strings"
)

// CSSReducer extracts only the CSS rules that correspond to the used Tailwind
// classes. The reducer operates purely on text and does not require a full CSS
// parser. It performs deterministic substring and pattern matching to locate
// class selectors and their associated rule blocks.
//
// The reducer preserves:
//   - standard utility selectors (e.g. ".flex")
//   - escaped selectors (e.g. ".sm\\:text-white")
//   - responsive blocks (e.g. "@media (min-width: 640px) { ... }")
//   - pseudo-class variants (e.g. ".hover\\:bg-red-500:hover")
//
// The reducer does not attempt to reorder, compress, or otherwise transform
// the CSS. It only extracts the minimal set of rules required for the used
// classes.
func CSSReducer(fullCSS string, usedClasses []string) string {
	var builder strings.Builder

	// Precompile a regex that captures a CSS rule block:
	//   .selector { ... }
	//   @media (...) { ... }
	//
	// This is a simplified pattern suitable for Tailwind's generated CSS.
	rulePattern := regexp.MustCompile(`(?s)([^{}]+)\{[^{}]*\}`)

	matches := rulePattern.FindAllString(fullCSS, -1)
	if len(matches) == 0 {
		return ""
	}

	// Build a set for quick membership checks.
	classSet := make(map[string]bool, len(usedClasses))
	for _, c := range usedClasses {
		classSet[c] = true
	}

	for _, rule := range matches {
		selector := extractSelector(rule)

		// Check if the selector contains any used class.
		if selectorMatches(selector, classSet) {
			builder.WriteString(rule)
			builder.WriteString("\n")
		}
	}

	return builder.String()
}

// extractSelector returns the selector portion of a CSS rule.
// Example input: ".flex { display: flex; }"
// Output: ".flex"
func extractSelector(rule string) string {
	idx := strings.Index(rule, "{")
	if idx == -1 {
		return ""
	}
	return strings.TrimSpace(rule[:idx])
}

// selectorMatches determines whether the selector references any of the
// used Tailwind classes. Tailwind classes may appear in escaped form,
// such as ".sm\\:text-white".
func selectorMatches(selector string, classSet map[string]bool) bool {
	// Remove leading dots and whitespace.
	s := strings.TrimSpace(selector)
	s = strings.TrimPrefix(s, ".")

	// Tailwind selectors may contain multiple classes or variants.
	// Split on commas to handle grouped selectors.
	parts := strings.Split(s, ",")

	for _, p := range parts {
		p = strings.TrimSpace(p)

		// Remove pseudo-classes (e.g. ":hover").
		base := p
		if idx := strings.Index(base, ":"); idx != -1 {
			base = base[:idx]
		}

		// Unescape Tailwind's escaped colon.
		base = strings.ReplaceAll(base, `\:`, ":")

		if classSet[base] {
			return true
		}
	}

	return false
}
