package dsl

import (
	"io"
	"strings"
)

func collectTailwindMethods(children []Node) []string {
	var out []string

	for _, c := range children {
		if !c.isAttribute {
			continue
		}

		if !c.isCSS {
			continue
		}

		name, _ := c.GetAttributeNameValue()

		if name == "TW" {
			continue
		}

		if !isTailwindKey(name) {
			continue
		}

		out = append(out, name)
	}

	return out
}

func isTailwindKey(name string) bool {
	return strings.HasPrefix(name, "Text")
}

func printElementTW(p *DSLPrinter, n Node, indent int) {
	p.writeIndentN(indent)

	_, _ = p.w.Write([]byte("dsl."))
	_, _ = p.w.Write([]byte(capitalize(n.GetTagName())))
	_, _ = p.w.Write([]byte("(\n"))

	// collect Tailwind CSS-like attributes
	methods := collectTailwindMethods(n.children)

	if len(methods) > 0 {
		p.writeIndentN(indent + 1)
		_, _ = p.w.Write([]byte("dsl.TW()"))

		for _, m := range methods {
			_, _ = p.w.Write([]byte("."))
			_, _ = p.w.Write([]byte(m))
			_, _ = p.w.Write([]byte("()"))
		}

		_, _ = p.w.Write([]byte(",\n"))
	}

	// print remaining children (non-CSS)
	for _, child := range n.children {
		if child.isCSS {
			continue // Skip CSS nodes - already handled above
		}

		printNodeTW(p, child, indent+1)
	}

	p.writeIndentN(indent)
	_, _ = p.w.Write([]byte(")\n"))
}

func printNodeTW(p *DSLPrinter, n Node, indent int) {
	switch {
	case n.isText:
		p.writeIndentN(indent)
		p.printText(n)
		_, _ = p.w.Write([]byte(",\n"))

	case n.isCSS:
		// Tailwind CSS attributes handled in printElementTW - skip here
		return

	case n.isAttribute:
		p.writeIndentN(indent)
		p.printAttribute(n)
		_, _ = p.w.Write([]byte(",\n"))

	default:
		printElementTW(p, n, indent)
	}
}

func PrintDSLWithTailwind(w io.Writer, n Node) {
	p := &DSLPrinter{w: w}

	printElementTW(p, n, 0)
}
