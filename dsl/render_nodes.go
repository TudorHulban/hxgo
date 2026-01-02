package dsl

import (
	"io"
)

// Empty nodes still emit tags because the renderer treats
// the element itself as meaningful, not just its children.
func renderNodes(w io.Writer, elName string, nodes ...Node) ([]Style, error) {
	var collected []Style // no allocation if no styke nodes. grow is a trade-off.

	_, _ = io.WriteString(w, "<")
	_, _ = io.WriteString(w, elName)

	// Render attributes
	for _, node := range nodes {
		if node == nil || !node.isAttribute() {
			continue
		}
		// Attributes don't return styles, just ignore them
		if _, err := node.Render(w); err != nil {
			return nil, err
		}
	}

	_, _ = io.WriteString(w, ">")

	if isTagless(elName) {
		return collected, nil
	}

	// Render children and collect their styles in one pass
	for _, node := range nodes {
		if node == nil || node.isAttribute() {
			continue
		}

		// Render child HTML and collect its styles (from entire subtree)
		styles, err := node.Render(w)
		if err != nil {
			return nil, err
		}

		collected = append(collected, styles...)
	}

	_, _ = io.WriteString(w, "</")
	_, _ = io.WriteString(w, elName)
	_, _ = io.WriteString(w, ">")

	return collected, nil
}

func renderNodesWithCSSId(w io.Writer, elName, cssID string, nodes ...Node) ([]Style, error) {
	var collected []Style

	_, _ = io.WriteString(w, "<")
	_, _ = io.WriteString(w, elName)

	_, _ = AttrIDLength(cssID).Render(w) // introduces one allocation more than wo css ID

	// Render attributes
	for _, node := range nodes {
		if node == nil || !node.isAttribute() {
			continue
		}
		// Attributes don't return styles, just ignore them
		if _, err := node.Render(w); err != nil {
			return nil, err
		}
	}

	_, _ = io.WriteString(w, ">")

	if isTagless(elName) {
		return collected, nil
	}

	// Render children and collect their styles in one pass
	for _, node := range nodes {
		if node == nil || node.isAttribute() {
			continue
		}

		// Render child HTML and collect its styles (from entire subtree)
		styles, err := node.Render(w)
		if err != nil {
			return nil, err
		}

		collected = append(collected, styles...)
	}

	_, _ = io.WriteString(w, "</")
	_, _ = io.WriteString(w, elName)
	_, _ = io.WriteString(w, ">")

	return collected, nil
}
