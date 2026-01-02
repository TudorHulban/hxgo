package dsl

import (
	"io"
)

// Empty nodes still emit tags because the renderer treats
// the element itself as meaningful, not just its children.
func renderNodes(w io.Writer, elName string, nodes ...Node) ([]Style, error) {
	var collected []Style

	_, _ = io.WriteString(w, "<")
	_, _ = io.WriteString(w, elName)

	// First pass: attributes + buffering children
	type child struct {
		render Render
	}
	var children []child

	for _, node := range nodes {
		if node == nil {
			continue
		}

		isAttrFn, renderFn := node(w)

		if isAttrFn() {
			// Attribute: render immediately
			if _, err := renderFn(w); err != nil {
				return nil, err
			}

			continue
		}

		// Child: store for later
		children = append(children, child{render: renderFn})
	}

	_, _ = io.WriteString(w, ">")

	if isTagless(elName) {
		return collected, nil
	}

	// Second phase: render children (but NOT calling node again)
	for _, c := range children {
		styles, err := c.render(w)
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

	isAttrFn, renderFn := AttrIDLength(cssID)(w)
	if isAttrFn() {
		if _, err := renderFn(w); err != nil {
			return nil, err
		}
	}

	// First pass: attributes + buffering children
	type child struct {
		render Render
	}
	var children []child

	for _, node := range nodes {
		if node == nil {
			continue
		}

		isAttrFn, renderFn := node(w)

		if isAttrFn() {
			// Attribute: render immediately
			if _, err := renderFn(w); err != nil {
				return nil, err
			}

			continue
		}

		// Child: store for later
		children = append(children, child{render: renderFn})
	}

	_, _ = io.WriteString(w, ">")

	if isTagless(elName) {
		return collected, nil
	}

	// Second phase: render children (but NOT calling node again)
	for _, c := range children {
		styles, err := c.render(w)
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
