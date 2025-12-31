package dsl

import (
	"io"
)

type Writer func(io.Writer) ([]Style, error)

func (w Writer) isAttribute() bool {
	return false
}

func (w Writer) Render(wr io.Writer) ([]Style, error) {
	return w(wr)
}
