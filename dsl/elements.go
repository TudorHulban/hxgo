package dsl

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
	return func() NodeOutput {
		// Render child node first
		child := node()

		// Pre-size: "<!doctype html>" + child HTML
		size := len("<!doctype html>") + len(child.HTML)
		html := make([]byte, 0, size)

		html = append(html, "<!doctype html>"...)
		html = append(html, child.HTML...)

		return NodeOutput{
			IsAttr: false,
			HTML:   html,
			Styles: child.Styles,
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
