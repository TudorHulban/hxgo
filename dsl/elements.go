package dsl

import (
	"io"
)

func Button(children ...Node) Writer {
	return El(
		"button",
		children...,
	)
}

func Body(children ...Node) Writer {
	return El(
		"body",
		children...,
	)
}

func Doctype(node Node) Writer {
	return Writer(
		func(w io.Writer) ([]Style, error) {
			if _, errWrite := w.Write([]byte("<!doctype html>")); errWrite != nil {
				return nil, errWrite
			}

			return node.Render(w)
		},
	)
}

func Head(children ...Node) Writer {
	return El(
		"head",
		children...,
	)
}

func HTML(children ...Node) Writer {
	return El(
		"html",
		children...,
	)
}

func Meta(children ...Node) Writer {
	return El(
		"meta",
		children...,
	)
}

func Title(children ...Node) Writer {
	return El(
		"title",
		children...,
	)
}

func Script(children ...Node) Writer {
	return El(
		"script",
		children...,
	)
}

func SVG(children ...Node) Writer {
	return El(
		"svg",
		children...,
	)
}

func Path(children ...Node) Writer {
	return El(
		"path",
		children...,
	)
}
