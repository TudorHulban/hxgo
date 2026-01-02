package dsl

import (
	"io"
)

// Elements are never attributes
func El(name string, nodes ...Node) Node {
	return func(w io.Writer) (IsAttribute, Render) {
		return func() bool {
				return false
			},

			func(wr io.Writer) ([]Style, error) {
				return renderNodes(
					wr,
					name,
					nodes...,
				)
			}
	}
}

func ElWId(name, cssID string, nodes ...Node) Node {
	return func(w io.Writer) (IsAttribute, Render) {
		return func() bool {
				return false
			},

			func(wr io.Writer) ([]Style, error) {
				return renderNodesWithCSSId(
					wr,
					name,
					cssID,
					nodes...,
				)
			}
	}
}
