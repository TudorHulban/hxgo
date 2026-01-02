package dsl

import "io"

// no‑op node that is NOT a Node.
var noop Node = func(w io.Writer) (IsAttribute, Render) {
	return func() bool { return false },
		func(io.Writer) ([]Style, error) {
			return nil, nil
		}
}
