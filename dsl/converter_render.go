package dsl

func converterWalk(a *accumulator, n Node) {
	if n.fn == nil {
		return
	}

	if n.IsAttribute {
		n.fn(a, n.Data)

		return
	}

	// Text nodes - render and done
	if n.IsText {
		n.fn(a, n.Data)

		return
	}

	// Open tag
	n.fn(a, n.Data) // writes "<tag"

	// Attributes
	for i := range n.Children {
		if n.Children[i].IsAttribute {
			converterWalk(a, n.Children[i])
		}
	}

	// Check if void element
	type tagData struct {
		openTag  []byte
		closeTag []byte
	}

	d := (*tagData)(n.Data)

	if d.closeTag == nil {
		a.html = append(a.html, '/', '>')

		return
	}

	a.html = append(a.html, '>')

	// Content children
	for i := range n.Children {
		if n.Children[i].IsAttribute || n.Children[i].IsCSS {
			continue
		}

		converterWalk(a, n.Children[i])
	}

	// Close tag
	renderElementClose(a, n.Data)
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
