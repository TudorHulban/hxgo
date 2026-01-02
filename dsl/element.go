package dsl

import (
	"io"
)

// Elements are never attributes
func El(name string, nodes ...Node) Node {
	return func(w io.Writer) (bool, []Style, error) {
		styles, err := renderNodes(
			w,
			name,
			nodes...,
		)
		return false, styles, err
	}
}

func ElWId(name, cssID string, nodes ...Node) Node {
	return func(w io.Writer) (bool, []Style, error) {
		styles, err := renderNodesWithCSSId(
			w,
			name,
			cssID,
			nodes...,
		)
		return false, styles, err
	}
}
