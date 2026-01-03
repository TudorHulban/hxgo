package dsl

// Elements are never attributes
func El(name string, nodes ...Node) Node {
	open := []byte("<" + name + ">")
	close := []byte("</" + name + ">")

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
				attrs = append(attrs, out.HTMLParts...)
			} else {
				children = append(children, out.HTMLParts...)
			}
			styles = append(styles, out.Styles...)
		}

		parts := make([][]byte, 0, 2+len(attrs)+len(children))
		parts = append(parts, open)
		parts = append(parts, attrs...)
		parts = append(parts, []byte(">"))
		parts = append(parts, children...)
		parts = append(parts, close)

		return NodeOutput{
			IsAttr:    false,
			HTMLParts: parts,
			Styles:    styles,
		}
	}
}

func ElWId(name, cssID string, nodes ...Node) Node {
	return func() NodeOutput {
		htmlParts, styles := renderNodesWithCSSId(name, cssID, nodes...)

		return NodeOutput{
			IsAttr:    false,
			HTMLParts: htmlParts,
			Styles:    styles,
		}
	}
}
