package dsl

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type ResultDOMConversion struct {
	Description string
	Node        Node
}

func (r ResultDOMConversion) PrintWithDescription(w io.Writer) {
	_, _ = w.Write([]byte(r.Description))
	_, _ = w.Write([]byte(":\n"))

	PrintDSL(w, r.Node)
}

func (result ResultDOMConversion) PrintWithTransformers(writer io.Writer, transformers ...Transformer) {
	_, _ = writer.Write([]byte(result.Description))
	_, _ = writer.Write([]byte(":\n"))

	for _, transformer := range transformers {
		result.Node = transformer(result.Node)
	}

	PrintDSL(writer, result.Node)
}

type ResultsDOMConversion []ResultDOMConversion

func (r ResultsDOMConversion) PrintWithDescription(w io.Writer) {
	for i := range r {
		r[i].PrintWithDescription(w)

		_, _ = w.Write([]byte("\n"))
	}
}

type DOMNode struct {
	CSSID    string
	CSSClass string
}

func (d DOMNode) String() string {
	if len(d.CSSID) > 0 && len(d.CSSClass) > 0 {
		return "#" + d.CSSID + " ." + d.CSSClass
	}

	if len(d.CSSID) > 0 {
		return "#" + d.CSSID
	}

	if len(d.CSSClass) > 0 {
		return "." + d.CSSClass
	}

	return ""
}

func getAttr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}

	return ""
}

func matchesSelector(node *html.Node, selector DOMNode) bool {
	if selector.CSSID != "" {
		elementID := getAttr(node, "id")
		if elementID == selector.CSSID {
			return true
		}
	}

	if selector.CSSClass != "" {
		classAttr := getAttr(node, "class")
		if classAttr != "" {
			for classToken := range strings.FieldsSeq(classAttr) {
				if classToken == selector.CSSClass {
					return true
				}
			}
		}
	}

	return false
}

func findMatches(root *html.Node, sel DOMNode) []*html.Node {
	var out []*html.Node

	var walk func(*html.Node)

	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && matchesSelector(n, sel) {
			out = append(out, n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(root)

	return out
}

func ConvertDOMElements(doc *html.Node, selectors ...DOMNode) ResultsDOMConversion {
	var results []ResultDOMConversion

	for _, sel := range selectors {
		matches := findMatches(doc, sel)

		// take only the first match
		if len(matches) > 0 {
			results = append(
				results,
				ResultDOMConversion{
					Description: sel.String(),
					Node:        ConvertHTML(matches[0]),
				},
			)
		}
	}

	return results
}
