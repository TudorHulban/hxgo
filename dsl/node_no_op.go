package dsl

import "io"

// no‑op node that is NOT a Node.
var noop Node = func(w io.Writer) (bool, []Style, error) {
	return false, nil, nil
}
