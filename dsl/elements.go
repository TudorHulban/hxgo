package dsl

import (
	"io"
)

func Button(children ...Node) Node {
	return El(
		"button",
		children...,
	)
}

func Body(children ...Node) Node {
	return El(
		"body",
		children...,
	)
}

func Doctype(node Node) Node {
	return func(w io.Writer) (IsAttribute, Render) {
		return func() bool {
				return false // doctype is never an attribute
			},
			func(wr io.Writer) ([]Style, error) {
				if _, err := wr.Write([]byte("<!doctype html>")); err != nil {
					return nil, err
				}

				// Now render the wrapped node
				_, renderFn := node(wr)

				return renderFn(wr)
			}
	}
}

func Head(children ...Node) Node {
	return El(
		"head",
		children...,
	)
}

func HTML(children ...Node) Node {
	return El(
		"html",
		children...,
	)
}

func Meta(children ...Node) Node {
	return El(
		"meta",
		children...,
	)
}

func Title(children ...Node) Node {
	return El(
		"title",
		children...,
	)
}

func Script(children ...Node) Node {
	return El(
		"script",
		children...,
	)
}

func SVG(children ...Node) Node {
	return El(
		"svg",
		children...,
	)
}

func Path(children ...Node) Node {
	return El(
		"path",
		children...,
	)
}
