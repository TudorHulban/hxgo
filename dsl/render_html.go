package dsl

import (
	"encoding/base64"

	"github.com/TudorHulban/hxgo/helpers"
)

func RenderHTMLandCSS(nodes ...Node) ([]byte, string) {
	if len(nodes) == 0 {
		return []byte{}, ""
	}

	var a Acc

	for i := range nodes {
		walk(&a, nodes[i])
	}

	// HTML is already fully assembled
	html := a.html

	// No styles collected → no CSS
	if len(a.styles) == 0 {
		return html, ""
	}

	// Build CSS
	css := newSmartCSSCollector()
	for _, s := range a.styles {
		css.Add(s)
	}

	return html, css.String()
}

func HTMLwithDataCSS(html []byte, css string) string {
	return helpers.Sprintf(
		`<html><head><link rel="stylesheet" href="data:text/css;base64,%s"></head><body>%s</body></html>`,

		base64.StdEncoding.EncodeToString([]byte(css)),
		string(html),
	)
}

func walk(a *Acc, n Node) {
	if n.fn == nil {
		return
	}

	n.fn(a, n.data)

	for i := range n.children {
		walk(a, n.children[i])
	}
}

// Renderer: single-pass traversal
func Render(nodes ...Node) []byte {
	var a Acc

	for i := range nodes {
		walk(&a, nodes[i])
	}

	return a.html
}
