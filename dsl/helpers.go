package dsl

import (
	"io"
	"text/template"
)

func Text(text string) Node {
	return func(w io.Writer) (bool, []Style, error) {
		_, err := w.Write([]byte(
			template.HTMLEscapeString(text),
		))
		return false, nil, err
	}
}

func Raw(text string) Node {
	return func(w io.Writer) (bool, []Style, error) {
		_, err := w.Write([]byte(text))
		return false, nil, err
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
