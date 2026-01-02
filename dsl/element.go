package dsl

// Elements are never attributes
func El(name string, nodes ...Node) Node {
	return func() NodeOutput {
		var attrs [][]byte
		var children [][]byte
		var styles []Style

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

		// estimate capacity
		size := 1 + len(name) + 1 // "<" + name + ">"
		for _, a := range attrs {
			size += len(a)
		}
		for _, c := range children {
			size += len(c)
		}
		size += 3 + len(name) + 1 // "</" + name + ">"

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

		return NodeOutput{
			IsAttr: false,
			HTML:   html,
			Styles: styles,
		}
	}
}

func ElWId(name, cssID string, nodes ...Node) Node {
	return func() NodeOutput {
		html, styles := renderNodesWithCSSId(name, cssID, nodes...)
		return NodeOutput{
			IsAttr: false,
			HTML:   html,
			Styles: styles,
		}
	}
}
