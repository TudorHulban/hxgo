package dsl

import (
	"encoding/base64"

	"github.com/TudorHulban/hxgo/helpers"
)

func RenderHTMLandCSS(nodes ...Node) ([]byte, string) {
	// Fast path: no nodes
	if len(nodes) == 0 {
		return nil, ""
	}

	// Fast path: single node, no CSS
	if len(nodes) == 1 {
		n := nodes[0]
		if n != nil {
			out := n()
			if len(out.Styles) == 0 {
				// compute total size from this node only
				total := 0
				for _, p := range out.HTMLParts {
					total += len(p)
				}

				html := make([]byte, 0, total)
				for _, p := range out.HTMLParts {
					html = append(html, p...)
				}

				return html, ""
			}
		}
	}

	// General path: multiple nodes and/or CSS
	var parts [][]byte
	var styles []Style

	for _, n := range nodes {
		if n == nil {
			continue
		}
		out := n()
		parts = append(parts, out.HTMLParts...)
		styles = append(styles, out.Styles...)
	}

	// compute total size
	total := 0
	for _, p := range parts {
		total += len(p)
	}

	html := make([]byte, 0, total)
	for _, p := range parts {
		html = append(html, p...)
	}

	// no styles collected → no CSS work
	if len(styles) == 0 {
		return html, ""
	}

	css := newSmartCSSCollector()
	for _, s := range styles {
		css.Add(s)
	}

	return html, css.String()
}

// func RenderHTMLandCSS(nodes ...Node) ([]byte, string) {
// 	var parts [][]byte
// 	var styles []Style

// 	for _, n := range nodes {
// 		if n == nil {
// 			continue
// 		}
// 		out := n()
// 		parts = append(parts, out.HTMLParts...)
// 		styles = append(styles, out.Styles...)
// 	}

// 	// compute total size
// 	total := 0
// 	for _, p := range parts {
// 		total += len(p)
// 	}

// 	html := make([]byte, 0, total)
// 	for _, p := range parts {
// 		html = append(html, p...)
// 	}

// 	// skip CSS collector entirely when no styles
// 	if len(styles) == 0 {
// 		// return a compile‑time constant
// 		return html, ""
// 	}

// 	css := newSmartCSSCollector()
// 	for _, s := range styles {
// 		css.Add(s)
// 	}

// 	return html, css.String()
// }

func HTMLwithDataCSS(html []byte, css string) string {
	return helpers.Sprintf(
		`<html><head><link rel="stylesheet" href="data:text/css;base64,%s"></head><body>%s</body></html>`,

		base64.StdEncoding.EncodeToString([]byte(css)),
		string(html),
	)
}
