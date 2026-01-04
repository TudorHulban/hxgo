package dsl

import "unsafe"

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

func Doctype(child Node) Node {
	return Node{
		fn:       renderDoctype,
		data:     nil,
		children: []Node{child},
	}
}

func renderDoctype(a *Accumulator, _ unsafe.Pointer) {
	a.Write1("<!doctype html>")
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
