package dsl

import "unsafe"

func A(children ...Node) Node {
	return El(
		"a",
		children...,
	)
}

func Aside(children ...Node) Node {
	return El(
		"aside",
		children...,
	)
}

func Class(name string) Node {
	return Node{
		fn:   renderClass,
		data: unsafe.Pointer(&name),
	}
}

func renderClass(a *Accumulator, p unsafe.Pointer) {
	name := *(*string)(p)
	a.Write1(` class="`)
	a.Write1(name)
	a.Write1(`"`)
}

func Div(children ...Node) Node {
	return Node{
		fn:       renderDiv,
		data:     nil,
		children: children,
	}
}

func renderDiv(a *Accumulator, _ unsafe.Pointer) {
	a.Write1("<div")
	// attributes would be rendered by children before closing '>'
	a.Write1(">")
	// children rendered by walk()
	a.Write1("</div>")
}

func Label(children ...Node) Node {
	return El(
		"label",
		children...,
	)
}

func Form(children ...Node) Node {
	return El(
		"form",
		children...,
	)
}

func FormWithID(cssID string, children ...Node) Node {
	return ElWId(
		"form",
		cssID,
		children...,
	)
}

func H1(children ...Node) Node {
	return El(
		"h1",
		children...,
	)
}

func H2(children ...Node) Node {
	return El(
		"h2",
		children...,
	)
}

func H3(children ...Node) Node {
	return El(
		"h3",
		children...,
	)
}

func H4(children ...Node) Node {
	return El(
		"h4",
		children...,
	)
}

func H5(children ...Node) Node {
	return El(
		"h5",
		children...,
	)
}

func H6(children ...Node) Node {
	return El(
		"h6",
		children...,
	)
}

func Img(children ...Node) Node {
	return El(
		"img",
		children...,
	)
}

func Input(children ...Node) Node {
	return El(
		"input",
		children...,
	)
}

func Li(children ...Node) Node {
	return El(
		"li",
		children...,
	)
}

func Link(children ...Node) Node {
	return El(
		"link",
		children...,
	)
}

func TextArea(children ...Node) Node {
	return El(
		"textarea",
		children...,
	)
}

func Table(children ...Node) Node {
	return El(
		"table",
		children...,
	)
}

func THead(children ...Node) Node {
	return El(
		"thead",
		children...,
	)
}

func TBody(children ...Node) Node {
	return El(
		"tbody",
		children...,
	)
}

func Th(children ...Node) Node {
	return El(
		"th",
		children...,
	)
}

func Tr(children ...Node) Node {
	return El(
		"tr",
		children...,
	)
}

func Nav(children ...Node) Node {
	return El(
		"nav",
		children...,
	)
}

func P(children ...Node) Node {
	return El(
		"p",
		children...,
	)
}

func Ul(children ...Node) Node {
	return El(
		"ul",
		children...,
	)
}

func Ol(children ...Node) Node {
	return El(
		"ol",
		children...,
	)
}

func Span(children ...Node) Node {
	return El(
		"span",
		children...,
	)
}
