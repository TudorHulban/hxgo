package dsl

import "unsafe"

func Text(text string) Node {
	return Node{
		fn: func(a *Acc, p unsafe.Pointer) {
			s := *(*string)(p)
			a.Write(s)
		},

		data: unsafe.Pointer(&text),
	}
}

func Raw(s string) Node {
	return Node{
		fn:   renderRaw,
		data: unsafe.Pointer(&s),
	}
}

func renderRaw(a *Acc, p unsafe.Pointer) {
	s := *(*string)(p)
	a.Write(s) // no escaping, direct write
}

func If(condition bool, node Node) Node {
	if condition {
		return node
	}

	return Noop
}

func Ternary(condition bool, node1, node2 Node) Node {
	if condition {
		return node1
	}

	return node2
}
