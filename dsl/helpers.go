package dsl

import (
	"text/template"
)

func Text(text string) Node {
	return func() NodeOutput {
		escaped := template.HTMLEscapeString(text)
		html := []byte(escaped) // can be optimized to manual escape later

		return NodeOutput{
			IsAttr: false,
			HTML:   html,
			Styles: nil,
		}
	}
}

func Raw(text string) Node {
	return func() NodeOutput {
		html := []byte(text)

		return NodeOutput{
			IsAttr: false,
			HTML:   html,
			Styles: nil,
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
