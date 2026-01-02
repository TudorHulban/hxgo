package dsl

import (
	"encoding/base64"

	"github.com/TudorHulban/hxgo/helpers"
)

func RenderHTMLandCSS(nodes ...Node) ([]byte, string, error) {
	var htmlParts [][]byte
	var styles []Style

	// Collect structured output
	for _, n := range nodes {
		if n == nil {
			continue
		}

		out := n()

		// top-level nodes are always children, never attributes
		htmlParts = append(htmlParts, out.HTML)
		styles = append(styles, out.Styles...)
	}

	// Estimate total size
	total := 0
	for _, p := range htmlParts {
		total += len(p)
	}

	// Concatenate HTML once
	html := make([]byte, 0, total)
	for _, p := range htmlParts {
		html = append(html, p...)
	}

	// CSS collector
	collector := newSmartCSSCollector()
	for _, s := range styles {
		collector.Add(s)
	}

	return html, collector.String(), nil
}

func HTMLwithDataCSS(html []byte, css string) string {
	return helpers.Sprintf(
		`<html><head><link rel="stylesheet" href="data:text/css;base64,%s"></head><body>%s</body></html>`,

		base64.StdEncoding.EncodeToString([]byte(css)),
		string(html),
	)
}
