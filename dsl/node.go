package dsl

import "unsafe"

type renderer func(*accumulator, unsafe.Pointer)

// Node: function pointer plus data pointer.
type Node struct {
	fn          renderer
	data        unsafe.Pointer
	children    []Node
	isAttribute bool
}
