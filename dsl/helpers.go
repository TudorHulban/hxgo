package dsl

import (
	"text/template"
)

func Text(text string) Node {
	return func() NodeOutput {
		return NodeOutput{
			IsAttr: false,
			HTMLParts: [][]byte{
				[]byte(
					template.HTMLEscapeString(text),
				),
			},
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

	return nil
}

func Ternary(condition bool, node1, node2 Node) Node {
	if condition {
		return node1
	}

	return node2
}
