package dsl

import "unsafe"

func Button(children ...Node) Node {
	return el(
		"button",
		children...,
	)
}

func Body(children ...Node) Node {
	return el(
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

func renderDoctype(a *accumulator, _ unsafe.Pointer) {
	a.Write1("<!doctype html>")
}

func Head(children ...Node) Node {
	return el(
		"head",
		children...,
	)
}

func HTML(children ...Node) Node {
	return el(
		"html",
		children...,
	)
}

func Meta(children ...Node) Node {
	return el(
		"meta",
		children...,
	)
}

func Title(children ...Node) Node {
	return el(
		"title",
		children...,
	)
}

func Script(children ...Node) Node {
	return el(
		"script",
		children...,
	)
}

func SVG(children ...Node) Node {
	return el(
		"svg",
		children...,
	)
}

func Path(children ...Node) Node {
	return el(
		"path",
		children...,
	)
}
