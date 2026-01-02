package dsl

import (
	"io"
	"text/template"
)

func Text(text string) Node {
	return func(w io.Writer) (IsAttribute, Render) {
		return func() bool {
				return false
			},

			func(wr io.Writer) ([]Style, error) {
				_, errWrite := w.Write(
					[]byte(
						template.HTMLEscapeString(text),
					),
				)

				return nil, errWrite
			}
	}
}

func Raw(text string) Node {
	return func(w io.Writer) (IsAttribute, Render) {
		return func() bool {
				return false
			},

			func(wr io.Writer) ([]Style, error) {
				_, errWrite := w.Write(
					[]byte(
						[]byte(text),
					),
				)

				return nil, errWrite
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
