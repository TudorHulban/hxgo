package dsl

import (
	"database/sql"
	"io"
)

func renderNodes(w io.Writer, elName sql.NullString, nodes ...Node) ([]Style, error) {
	var collected []Style

	if elName.Valid {
		_, _ = w.Write([]byte("<" + elName.String))

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

		_, _ = w.Write([]byte(">"))

		if isTagless(elName.String) {
			return collected, nil
		}
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

	if elName.Valid && !isTagless(elName.String) {
		_, _ = w.Write([]byte("</" + elName.String + ">"))
	}

	return collected, nil
}

func renderNodesWithCSSId(w io.Writer, elName, cssID string, nodes ...Node) ([]Style, error) {
	var collected []Style

	_, _ = w.Write([]byte("<" + elName))

	_, _ = AttrIDLength(cssID).Render(w)

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

	_, _ = w.Write([]byte(">"))

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

	if !isTagless(elName) {
		_, _ = w.Write([]byte("</" + elName + ">"))
	}

	return collected, nil
}
