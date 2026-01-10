package dsl

import (
	"sort"
	"strings"
)

// Single builder used by the renderer.
// accumulator accumulates HTML bytes and CSS styles during rendering.
// For better performance,
// use NewAcc with an estimated size based on the typical output.
type accumulator struct {
	html []byte
	css  []CSSContribution
}

// newAccumulator creates an Accumulator with pre-allocated capacity.
// estimatedHTMLSize should be the approximate size of the rendered HTML in bytes.
// Pre-allocating reduces memory allocations during rendering.
// Example: for a typical page of ~10KB, use newAccumulator(10240, 16)
func newAccumulator(estimatedHTMLSize, estimatedCSSRules int) *accumulator {
	return &accumulator{
		html: make([]byte, 0, estimatedHTMLSize),
		css:  make([]CSSContribution, 0, estimatedCSSRules),
	}
}

func (a *accumulator) Write1(value string) {
	a.html = append(a.html, value...)
}

// Fused writes.
func (a *accumulator) Write3(value1, value2, value3 string) {
	a.html = append(a.html, value1...)
	a.html = append(a.html, value2...)
	a.html = append(a.html, value3...)
}

// Fused writes.
func (a *accumulator) Write5(value1, value2, value3, value4, value5 string) {
	a.html = append(a.html, value1...)
	a.html = append(a.html, value2...)
	a.html = append(a.html, value3...)
	a.html = append(a.html, value4...)
	a.html = append(a.html, value5...)
}

func normalizeCSS(text string) string {
	text = strings.TrimSpace(text)

	lines := strings.Split(text, "\n")

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return strings.Join(lines, "\n")
}

func stripTrailingBraces(text string) string {
	text = strings.TrimRight(text, " \t\r\n") // Remove trailing whitespace

	// Remove ALL trailing braces
	for len(text) > 0 && text[len(text)-1] == '}' {
		text = text[:len(text)-1]
		text = strings.TrimRight(text, " \t\r\n")
	}

	return text
}

func (a *accumulator) groupCSS() ([]CSSContributionKey, map[CSSContributionKey]*group) {
	groups := make(map[CSSContributionKey]*group)
	order := make([]CSSContributionKey, 0, len(a.css))

	for _, contribution := range a.css {
		key := contribution.CSSContributionKey

		g, exists := groups[key]
		if !exists {
			g = newGroup()
			groups[key] = g

			order = append(order, key)
		}

		// merge declarative
		for _, pair := range contribution.DeclarativeStyle {
			g.Declarative[pair[0]] = pair[1]
		}

		// dedupe procedural
		for _, fn := range contribution.ProceduralCSS {
			g.Procedural.Add(fn)
		}
	}

	return order, groups
}

func (*accumulator) emitStyles(order []CSSContributionKey, groups map[CSSContributionKey]*group) string {
	var sb strings.Builder

	for _, key := range order {
		g := groups[key]
		if len(g.Declarative) == 0 {
			continue
		}

		if key.Breakpoint != "" {
			sb.WriteString("@media (")

			if key.DesktopFirst {
				sb.WriteString("min-width:")
			} else {
				sb.WriteString("max-width:")
			}

			sb.WriteString(key.Breakpoint)
			sb.WriteString(") {\n")
		}

		sb.WriteString("  ")
		sb.WriteString(key.Selector)
		sb.WriteString(" {\n")

		names := make([]string, 0, len(g.Declarative))

		for k := range g.Declarative {
			names = append(names, k)
		}

		sort.Strings(names)

		for _, name := range names {
			sb.WriteString("    ")
			sb.WriteString(name)
			sb.WriteString(": ")
			sb.WriteString(g.Declarative[name])
			sb.WriteString(";\n")
		}

		sb.WriteString("  }\n")

		if key.Breakpoint != "" {
			sb.WriteString("}\n")
		}
	}

	return sb.String()
}

func (*accumulator) emitProcedural(order []CSSContributionKey, groups map[CSSContributionKey]*group) string {
	var sb strings.Builder

	for _, key := range order {
		g := groups[key]
		if len(g.Procedural.order) == 0 {
			continue
		}

		for _, fn := range g.Procedural.order {
			raw := normalizeCSS(fn())
			raw = stripTrailingBraces(raw)

			// ----------------------------------------------------
			// CASE 1: No media query
			// ----------------------------------------------------
			if key.Breakpoint == "" {
				// GLOBAL CSS (no selector)
				if key.Selector == "" {
					sb.WriteString(raw)
					sb.WriteString("\n")

					continue
				}

				// SELECTOR CSS
				sb.WriteString(key.Selector)
				sb.WriteString("{")
				sb.WriteString(raw)
				sb.WriteString("}\n")

				continue
			}

			// ----------------------------------------------------
			// CASE 2: Media query
			// ----------------------------------------------------

			// @media (min|max-width: X)
			sb.WriteString("@media (")

			if key.DesktopFirst {
				sb.WriteString("min-width:")
			} else {
				sb.WriteString("max-width:")
			}

			sb.WriteString(key.Breakpoint)
			sb.WriteString(") {\n")

			// GLOBAL CSS inside media query
			if key.Selector == "" {
				sb.WriteString(raw)
				sb.WriteString("\n")
				sb.WriteString("}\n")

				continue
			}

			// SELECTOR CSS inside media query
			sb.WriteString("  ")
			sb.WriteString(key.Selector)
			sb.WriteString(" {\n")
			sb.WriteString("    ")
			sb.WriteString(raw)
			sb.WriteString("\n  }\n")

			sb.WriteString("}\n")
		}
	}

	return sb.String()
}

func (a *accumulator) EmitStyles() string {
	return a.emitStyles(
		a.groupCSS(),
	)
}

func (a *accumulator) EmitProcedural() string {
	return a.emitProcedural(
		a.groupCSS(),
	)
}

func (a *accumulator) BuildCSS() (string, string) {
	order, groups := a.groupCSS()

	return a.emitStyles(order, groups), a.emitProcedural(order, groups)
}
