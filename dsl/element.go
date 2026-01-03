package dsl

// Elements are never attributes
func El(name string, nodes ...Node) Node {
	openStart := []byte("<" + name)
	closeTag := []byte("</" + name + ">")
	gt := []byte(">")

	// CHANGED: removed `outputs := make([]NodeOutput, ...)`
	// Instead, we count directly in a single pass.
	return func() NodeOutput {
		attrCount := 0
		childCount := 0
		styleCount := 0

		// First pass: count fragments
		for _, n := range nodes {
			if n == nil {
				continue
			}
			out := n()
			if out.IsAttr {
				attrCount += len(out.HTMLParts)
			} else {
				childCount += len(out.HTMLParts)
			}
			styleCount += len(out.Styles)
		}

		// Pre-allocate parts
		partsLen := 1 + attrCount + 1 + childCount + 1
		parts := make([][]byte, partsLen)
		idx := 0

		parts[idx] = openStart
		idx++

		// CHANGED: second pass directly collects attrs/children without storing outputs
		for _, n := range nodes {
			if n == nil {
				continue
			}
			out := n()
			if out.IsAttr {
				for _, p := range out.HTMLParts {
					parts[idx] = p
					idx++
				}
			}
		}

		parts[idx] = gt
		idx++

		for _, n := range nodes {
			if n == nil {
				continue
			}
			out := n()
			if !out.IsAttr {
				for _, p := range out.HTMLParts {
					parts[idx] = p
					idx++
				}
			}
		}

		parts[idx] = closeTag

		// Styles
		var styles []Style
		if styleCount > 0 {
			styles = make([]Style, 0, styleCount)
			for _, n := range nodes {
				if n == nil {
					continue
				}
				out := n()
				styles = append(styles, out.Styles...)
			}
		}

		return NodeOutput{
			IsAttr:    false,
			HTMLParts: parts,
			Styles:    styles,
		}
	}
}

func ElWId(name, cssID string, nodes ...Node) Node {
	// Prebuild static fragments (same as El)
	openStart := make([]byte, 0, 1+len(name))
	openStart = append(openStart, '<')
	openStart = append(openStart, name...)

	closeTag := make([]byte, 0, 2+len(name)+1)
	closeTag = append(closeTag, '<', '/')
	closeTag = append(closeTag, name...)
	closeTag = append(closeTag, '>')

	gt := []byte(">")

	// CHANGED: prebuild id attribute once
	var idAttr []byte
	if cssID != "" {
		idAttr = []byte(` id="` + cssID + `"`)
	}

	return func() NodeOutput {
		// First pass: count fragments
		attrCount := 0
		childCount := 0
		styleCount := 0

		// CHANGED: if idAttr exists, count it as 1 attribute fragment
		if idAttr != nil {
			attrCount++
		}

		for _, n := range nodes {
			if n == nil {
				continue
			}
			out := n()
			if out.IsAttr {
				attrCount += len(out.HTMLParts)
			} else {
				childCount += len(out.HTMLParts)
			}
			styleCount += len(out.Styles)
		}

		// Preallocate parts: <tag + attrs + > + children + </tag>
		partsLen := 1 + attrCount + 1 + childCount + 1
		parts := make([][]byte, partsLen)
		idx := 0

		parts[idx] = openStart
		idx++

		// CHANGED: inject id attribute first
		if idAttr != nil {
			parts[idx] = idAttr
			idx++
		}

		// Collect attrs
		for _, n := range nodes {
			if n == nil {
				continue
			}
			out := n()
			if out.IsAttr {
				for _, p := range out.HTMLParts {
					parts[idx] = p
					idx++
				}
			}
		}

		parts[idx] = gt
		idx++

		// Collect children
		for _, n := range nodes {
			if n == nil {
				continue
			}
			out := n()
			if !out.IsAttr {
				for _, p := range out.HTMLParts {
					parts[idx] = p
					idx++
				}
			}
		}

		parts[idx] = closeTag

		// Styles
		var styles []Style
		if styleCount > 0 {
			styles = make([]Style, 0, styleCount)
			for _, n := range nodes {
				if n == nil {
					continue
				}
				out := n()
				styles = append(styles, out.Styles...)
			}
		}

		return NodeOutput{
			IsAttr:    false,
			HTMLParts: parts,
			Styles:    styles,
		}
	}
}
