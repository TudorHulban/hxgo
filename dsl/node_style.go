package dsl

import "unsafe"

type Style struct {
	Selector string
	Props    [][2]string
	Media    string // optional media query
}

// Style nodes are not attributes.
// no HTML output, only styles.
//
// Create the node first and then apply styles to it.
func GetStyledNode(child Node, styles ...Style) Node {
	data := &styles // // intentionally move to heap.

	return Node{
		fn:       renderStyles,
		data:     unsafe.Pointer(data),
		children: []Node{child},
	}
}

func renderStyles(a *accumulator, p unsafe.Pointer) {
	styles := *(*[]Style)(p)

	a.cssComponents = append(a.cssComponents, styles...)
}
