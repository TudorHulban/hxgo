package dsl

func converterWalk(a *accumulator, n Node) {
	if n.fn == nil {
		return
	}

	if n.isAttribute {
		n.fn(a, n.data)

		return
	}

	// Text nodes - render and done
	if n.isText {
		n.fn(a, n.data)

		return
	}

	// Open tag
	n.fn(a, n.data) // writes "<tag"

	// Attributes
	for i := range n.children {
		if n.children[i].isAttribute {
			converterWalk(a, n.children[i])
		}
	}

	if n.data == nil {
		return
	}

	// Check if void element
	type tagData struct {
		openTag  []byte
		closeTag []byte
	}

	d := (*tagData)(n.data)

	if d.closeTag == nil {
		a.html = append(a.html, '/', '>')

		return
	}

	a.html = append(a.html, '>')

	// Content children
	for i := range n.children {
		if n.children[i].isAttribute || n.children[i].isCSS {
			continue
		}

		converterWalk(a, n.children[i])
	}

	// Close tag
	renderElementClose(a, n.data)
}

func walk(a *accumulator, n Node) {
	if n.fn == nil {
		return
	}

	n.fn(a, n.data)

	for i := range n.children {
		walk(a, n.children[i])
	}
}

func walkHTML(a *accumulator, n Node) {
	if n.fn == nil {
		return
	}

	n.fn(a, n.data)

	for i := range n.children {
		if n.children[i].isCSS {
			continue
		}

		walkHTML(a, n.children[i])
	}
}
