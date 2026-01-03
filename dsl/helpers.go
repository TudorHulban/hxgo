package dsl

import (
	"text/template"
)

func Text(text string) Node {
	escaped := template.HTMLEscapeString(text)
	b := []byte(escaped)

	// CHANGED: prebuild the slice once, instead of using a literal in the closure
	parts := [][]byte{b}

	return func() NodeOutput {
		return NodeOutput{
			IsAttr:    false,
			HTMLParts: parts, // no per-call slice allocation
			Styles:    nil,
		}
	}
}

func Raw(text string) Node {
	return func() NodeOutput {
		return NodeOutput{
			IsAttr: false,
			HTMLParts: [][]byte{
				[]byte(
					text,
				),
			},
		}
	}
}

func If(condition bool, node Node) Node {
	if condition {
		return node
	}

	return Noop
}

func Ternary(condition bool, node1, node2 Node) Node {
	if condition {
		return node1
	}

	return node2
}
