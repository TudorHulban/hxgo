package dsl

func A(children ...Node) Node {
	return el(
		"a",
		children...,
	)
}

func Aside(children ...Node) Node {
	return el(
		"aside",
		children...,
	)
}

func Div(children ...Node) Node {
	return el(
		"div",
		children...,
	)
}

func Label(children ...Node) Node {
	return el(
		"label",
		children...,
	)
}

func Form(children ...Node) Node {
	return el(
		"form",
		children...,
	)
}

func FormWithID(cssID string, children ...Node) Node {
	return elWId(
		"form",
		cssID,
		children...,
	)
}

func H1(children ...Node) Node {
	return el(
		"h1",
		children...,
	)
}

func H2(children ...Node) Node {
	return el(
		"h2",
		children...,
	)
}

func H3(children ...Node) Node {
	return el(
		"h3",
		children...,
	)
}

func H4(children ...Node) Node {
	return el(
		"h4",
		children...,
	)
}

func H5(children ...Node) Node {
	return el(
		"h5",
		children...,
	)
}

func H6(children ...Node) Node {
	return el(
		"h6",
		children...,
	)
}

func Input(children ...Node) Node {
	return el(
		"input",
		children...,
	)
}

func Li(children ...Node) Node {
	return el(
		"li",
		children...,
	)
}

func Link(children ...Node) Node {
	return el(
		"link",
		children...,
	)
}

func TextArea(children ...Node) Node {
	return el(
		"textarea",
		children...,
	)
}

func Table(children ...Node) Node {
	return el(
		"table",
		children...,
	)
}

func THead(children ...Node) Node {
	return el(
		"thead",
		children...,
	)
}

func TBody(children ...Node) Node {
	return el(
		"tbody",
		children...,
	)
}

func Th(children ...Node) Node {
	return el(
		"th",
		children...,
	)
}

func Tr(children ...Node) Node {
	return el(
		"tr",
		children...,
	)
}

func Nav(children ...Node) Node {
	return el(
		"nav",
		children...,
	)
}

func P(children ...Node) Node {
	return el(
		"p",
		children...,
	)
}

func Ul(children ...Node) Node {
	return el(
		"ul",
		children...,
	)
}

func Ol(children ...Node) Node {
	return el(
		"ol",
		children...,
	)
}

func Span(children ...Node) Node {
	return el(
		"span",
		children...,
	)
}

// inline‑element set

func Small(children ...Node) Node {
	return el(
		"small",
		children...,
	)
}

func Em(children ...Node) Node {
	return el(
		"em",
		children...,
	)
}

func Strong(children ...Node) Node {
	return el(
		"strong",
		children...,
	)
}

func Mark(children ...Node) Node {
	return el(
		"mark",
		children...,
	)
}

func Sub(children ...Node) Node {
	return el(
		"sub",
		children...,
	)
}

func Sup(children ...Node) Node {
	return el(
		"sup",
		children...,
	)
}

func Code(children ...Node) Node {
	return el(
		"code",
		children...,
	)
}

func I(children ...Node) Node {
	return el(
		"i",
		children...,
	)
}

func B(children ...Node) Node {
	return el(
		"b",
		children...,
	)
}

func U(children ...Node) Node {
	return el(
		"u",
		children...,
	)
}

func S(children ...Node) Node {
	return el(
		"s",
		children...,
	)
}

func Q(children ...Node) Node {
	return el(
		"q",
		children...,
	)
}
