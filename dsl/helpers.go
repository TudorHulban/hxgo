package dsl

import "unsafe"

func Text(text string) Node {
	return Node{
		fn: func(a *accumulator, p unsafe.Pointer) {
			s := *(*string)(p)
			a.Write1(s)
		},

		data: unsafe.Pointer(&text),
	}
}

func Raw(text string) Node {
	return Node{
		fn: func(a *accumulator, p unsafe.Pointer) {
			s := *(*string)(p)
			a.Write1(s)
		},

		data: unsafe.Pointer(&text),
	}
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
