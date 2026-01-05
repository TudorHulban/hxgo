package dsl

import "unsafe"

type Style struct {
	Selector string
	Props    [][2]string
	Media    string // optional media query
}

// Style nodes are not attributes.
// no HTML output, only styles.
func Styled(child Node, styles ...Style) Node {
	data := &styles // // intentionally move to heap.

	return Node{
		fn:       renderStyled,
		data:     unsafe.Pointer(data),
		children: []Node{child},
	}
}

func renderStyled(a *accumulator, p unsafe.Pointer) {
	styles := *(*[]Style)(p)
	a.css = append(a.css, styles...)
}
