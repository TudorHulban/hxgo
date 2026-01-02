package dsl

import (
	"io"
)

func El(name string, nodes ...Node) Writer {
	return func(w io.Writer) ([]Style, error) {
		return renderNodes(
			w,

			name,
			nodes...,
		)
	}
}

func ElWId(name, cssID string, nodes ...Node) Writer {
	return func(w io.Writer) ([]Style, error) {
		return renderNodesWithCSSId(
			w,

			name,
			cssID,
			nodes...,
		)
	}
}
