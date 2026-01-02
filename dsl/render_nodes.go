package dsl

// Empty nodes still emit tags because the renderer treats
// the element itself as meaningful, not just its children.
func renderNodes(name string, nodes ...Node) ([]byte, []Style) {
	var attrs [][]byte
	var children [][]byte
	var styles []Style

	// Collect structured output
	for _, n := range nodes {
		if n == nil {
			continue
		}
		out := n()
		if out.IsAttr {
			attrs = append(attrs, out.HTML)
		} else {
			children = append(children, out.HTML)
		}
		styles = append(styles, out.Styles...)
	}

	// Estimate final size
	size := 1 + len(name) // "<name"
	for _, a := range attrs {
		size += len(a)
	}
	size += 1 // ">"
	for _, c := range children {
		size += len(c)
	}
	size += 3 + len(name) + 1 // "</name>"

	// Build final HTML
	html := make([]byte, 0, size)
	html = append(html, '<')
	html = append(html, name...)
	for _, a := range attrs {
		html = append(html, a...)
	}
	html = append(html, '>')
	for _, c := range children {
		html = append(html, c...)
	}
	html = append(html, '<', '/')
	html = append(html, name...)
	html = append(html, '>')

	return html, styles
}

func renderNodesWithCSSId(name, cssID string, nodes ...Node) ([]byte, []Style) {
	var attrs [][]byte
	var children [][]byte
	var styles []Style

	// First: the injected id attribute
	if cssID != "" {
		idAttr := []byte(` id="` + cssID + `"`)
		attrs = append(attrs, idAttr)
	}

	// Collect structured output
	for _, n := range nodes {
		if n == nil {
			continue
		}
		out := n()
		if out.IsAttr {
			attrs = append(attrs, out.HTML)
		} else {
			children = append(children, out.HTML)
		}
		styles = append(styles, out.Styles...)
	}

	// Estimate final size
	size := 1 + len(name) // "<name"
	for _, a := range attrs {
		size += len(a)
	}
	size += 1 // ">"
	for _, c := range children {
		size += len(c)
	}
	size += 3 + len(name) + 1 // "</name>"

	// Build final HTML
	html := make([]byte, 0, size)
	html = append(html, '<')
	html = append(html, name...)
	for _, a := range attrs {
		html = append(html, a...)
	}
	html = append(html, '>')
	for _, c := range children {
		html = append(html, c...)
	}
	html = append(html, '<', '/')
	html = append(html, name...)
	html = append(html, '>')

	return html, styles
}
