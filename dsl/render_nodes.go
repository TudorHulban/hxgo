package dsl

// Empty nodes still emit tags because the renderer treats
// the element itself as meaningful, not just its children.
func renderNodes(name string, nodes ...Node) ([]byte, []Style) {
	var attrs [][]byte
	var children [][]byte
	var styles []Style

	// Collect fragments
	for _, n := range nodes {
		if n == nil {
			continue
		}
		out := n()
		if out.IsAttr {
			attrs = append(attrs, out.HTMLParts...)
		} else {
			children = append(children, out.HTMLParts...)
		}
		styles = append(styles, out.Styles...)
	}

	// Static fragments
	openTagStart := []byte("<" + name)
	closeTag := []byte("</" + name + ">")
	gt := []byte(">")

	// Count total size
	total := len(openTagStart) + len(gt) + len(closeTag)
	for _, a := range attrs {
		total += len(a)
	}
	for _, c := range children {
		total += len(c)
	}

	// Single allocation
	html := make([]byte, 0, total)

	// Build final HTML
	html = append(html, openTagStart...)
	for _, a := range attrs {
		html = append(html, a...)
	}
	html = append(html, gt...)
	for _, c := range children {
		html = append(html, c...)
	}
	html = append(html, closeTag...)

	return html, styles
}

func renderNodesWithCSSId(name, cssID string, nodes ...Node) ([][]byte, []Style) {
	var attrs [][]byte
	var children [][]byte
	var styles []Style

	// Inject id attribute
	if cssID != "" {
		idAttr := []byte(` id="` + cssID + `"`)
		attrs = append(attrs, idAttr)
	}

	// Collect fragments
	for _, n := range nodes {
		if n == nil {
			continue
		}
		out := n()
		if out.IsAttr {
			attrs = append(attrs, out.HTMLParts...)
		} else {
			children = append(children, out.HTMLParts...)
		}
		styles = append(styles, out.Styles...)
	}

	// Build fragment list (NOT concatenated HTML)
	openTagStart := []byte("<" + name)
	closeTag := []byte("</" + name + ">")
	gt := []byte(">")

	parts := make([][]byte, 0, 2+len(attrs)+len(children))
	parts = append(parts, openTagStart)
	parts = append(parts, attrs...)
	parts = append(parts, gt)
	parts = append(parts, children...)
	parts = append(parts, closeTag)

	return parts, styles
}
