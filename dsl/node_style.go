package dsl

import "unsafe"

type Style struct {
	Selector string
	Props    map[string]string
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

func renderStyled(a *Acc, p unsafe.Pointer) {
	styles := *(*[]Style)(p)
	a.styles = append(a.styles, styles...)
}
