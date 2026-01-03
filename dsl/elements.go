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
	// Static fragment, allocated once
	doctype := []byte("<!doctype html>")

	return func() NodeOutput {
		child := node()

		// Prepend doctype as a fragment
		parts := make([][]byte, 0, 1+len(child.HTMLParts))
		parts = append(parts, doctype)
		parts = append(parts, child.HTMLParts...)

		return NodeOutput{
			IsAttr:    false,
			HTMLParts: parts,
			Styles:    child.Styles,
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
