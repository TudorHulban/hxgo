package dsl

import (
	"encoding/base64"

	"github.com/TudorHulban/hxgo/helpers"
)

func walk(a *accumulator, n Node) {
	if n.fn == nil {
		return
	}

	n.fn(a, n.data)

	for i := range n.children {
		walk(a, n.children[i])
	}
}

func walkHTML(a *accumulator, n Node) {
	if n.fn == nil {
		return
	}

	n.fn(a, n.data)

	for i := range n.children {
		if n.children[i].isCSS {
			continue
		}

		walkHTML(a, n.children[i])
	}
}

// Render renders nodes to HTML bytes in single-pass traversal.
// For better performance with large outputs, use RenderHTMLWithCapacity.
func Render(nodes ...Node) []byte {
	if len(nodes) == 0 {
		return []byte{}
	}

	var a accumulator

	for i := range nodes {
		walkHTML(&a, nodes[i])
	}

	// HTML is already fully assembled
	return a.html
}

// RenderHX introduces
// the separator needed by HX client for parsing the response.
func RenderHX(nodes ...Node) []byte {
	if len(nodes) == 0 {
		return []byte{}
	}

	var a accumulator

	for i := range nodes {
		walkHTML(&a, nodes[i])

		a.Write1("|\n")
	}

	// HTML is already fully assembled
	return a.html
}

// RenderHTMLWithCapacity renders nodes to HTML with pre-allocated capacity.
// estimatedSize should be the approximate output size in bytes.
// This reduces allocations significantly for known output sizes.
func RenderHTMLWithCapacity(estimatedSize int, nodes ...Node) []byte {
	if len(nodes) == 0 {
		return []byte{}
	}

	a := newAccumulator(estimatedSize, 0)

	for i := range nodes {
		walkHTML(a, nodes[i])
	}

	return a.html
}

// RenderHXHTMLWithCapacity introduces
// the separator needed by HX client for parsing the response.
func RenderHXHTMLWithCapacity(estimatedSize int, nodes ...Node) []byte {
	if len(nodes) == 0 {
		return []byte{}
	}

	a := newAccumulator(estimatedSize, 0)

	for i := range nodes {
		walkHTML(a, nodes[i])

		a.Write1("|\n")
	}

	return a.html
}

func RenderHTMLandStyles(nodes ...Node) ([]byte, string) {
	if len(nodes) == 0 {
		return []byte{}, ""
	}

	var a accumulator

	for i := range nodes {
		walk(&a, nodes[i])
	}

	if len(a.css) == 0 {
		return a.html, "" // HTML is already fully assembled
	}

	return a.html, a.EmitStyles()
}

func RenderFull(nodes ...Node) ([]byte, string, string) {
	if len(nodes) == 0 {
		return nil, "", ""
	}

	var a accumulator

	for i := range nodes {
		walk(&a, nodes[i])
	}

	if len(a.css) == 0 {
		return a.html, "", "" // HTML is already fully assembled
	}

	styles, css := a.BuildCSS()

	return a.html, styles, css
}

// RenderHTMLandStylesWithCapacity renders nodes with pre-allocated capacity.
// estimatedHTMLSize: approximate HTML output size in bytes
// estimatedCSSRules: approximate number of CSS rules
func RenderHTMLandStylesWithCapacity(estimatedHTMLSize, estimatedCSSRules int, nodes ...Node) ([]byte, string) {
	if len(nodes) == 0 {
		return []byte{}, ""
	}

	a := newAccumulator(estimatedHTMLSize, estimatedCSSRules)

	for i := range nodes {
		walk(a, nodes[i])
	}

	if len(a.css) == 0 {
		return a.html, "" // HTML is already fully assembled
	}

	return a.html, a.EmitStyles()
}

func HTMLwithDataCSS(html []byte, css string) string {
	return helpers.Sprintf(
		`<html><head><link rel="stylesheet" href="data:text/css;base64,%s"></head><body>%s</body></html>`,

		base64.StdEncoding.EncodeToString([]byte(css)),
		string(html),
	)
}
