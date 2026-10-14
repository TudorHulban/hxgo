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

// RenderConvertedHTML should be used when testing the HTML converter.
func RenderConvertedHTML(nodes ...Node) []byte {
	if len(nodes) == 0 {
		return []byte{}
	}

	var a accumulator

	for i := range nodes {
		converterWalk(&a, nodes[i])
	}

	// HTML is already fully assembled
	return a.html
}
