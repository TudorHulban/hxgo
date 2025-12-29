package dsl

func A(children ...Node) Node {
	return El(
		"a",
		children...,
	)
}

func Aside(children ...Node) Writer {
	return El(
		"aside",
		children...,
	)
}

func Div(children ...Node) Writer {
	return El(
		"div",
		children...,
	)
}

func Label(children ...Node) Writer {
	return El(
		"label",
		children...,
	)
}

func Form(children ...Node) Writer {
	return El(
		"form",
		children...,
	)
}

func FormWithID(cssID string, children ...Node) Writer {
	return ElWId(
		"form",
		cssID,
		children...,
	)
}

func H1(children ...Node) Writer {
	return El(
		"h1",
		children...,
	)
}

func H2(children ...Node) Writer {
	return El(
		"h2",
		children...,
	)
}

func H3(children ...Node) Writer {
	return El(
		"h3",
		children...,
	)
}

func H4(children ...Node) Writer {
	return El(
		"h4",
		children...,
	)
}

func H5(children ...Node) Writer {
	return El(
		"h5",
		children...,
	)
}

func H6(children ...Node) Writer {
	return El(
		"h6",
		children...,
	)
}

func Img(children ...Node) Writer {
	return El(
		"img",
		children...,
	)
}

func Input(children ...Node) Writer {
	return El(
		"input",
		children...,
	)
}

func Li(children ...Node) Writer {
	return El(
		"li",
		children...,
	)
}

func Link(children ...Node) Writer {
	return El(
		"link",
		children...,
	)
}

func TextArea(children ...Node) Writer {
	return El(
		"textarea",
		children...,
	)
}

func Table(children ...Node) Writer {
	return El(
		"table",
		children...,
	)
}

func THead(children ...Node) Writer {
	return El(
		"thead",
		children...,
	)
}

func TBody(children ...Node) Writer {
	return El(
		"tbody",
		children...,
	)
}

func Th(children ...Node) Writer {
	return El(
		"th",
		children...,
	)
}

func Tr(children ...Node) Writer {
	return El(
		"tr",
		children...,
	)
}

func Nav(children ...Node) Writer {
	return El(
		"nav",
		children...,
	)
}

func P(children ...Node) Writer {
	return El(
		"p",
		children...,
	)
}

func Ul(children ...Node) Writer {
	return El(
		"ul",
		children...,
	)
}

func Ol(children ...Node) Writer {
	return El(
		"ol",
		children...,
	)
}

func Span(children ...Node) Writer {
	return El(
		"span",
		children...,
	)
}
