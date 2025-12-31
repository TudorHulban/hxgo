package dsl

import "io"

type Node interface {
	isAttribute() bool
	Render(w io.Writer) ([]Style, error)
}

var _ Node = attribute{}
var _ Node = Writer(
	func(io.Writer) ([]Style, error) {
		return nil, nil
	},
)
var _ Node = &NodeStyle{}
