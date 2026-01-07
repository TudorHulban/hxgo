package dsl

import "unsafe"

type renderer func(*accumulator, unsafe.Pointer)

// Node: function pointer plus data pointer.
//
// Data-oriented design.
// Manual control over execution.
// No interface dispatch.
// No virtual method tables.
type Node struct {
	fn   renderer
	data unsafe.Pointer

	children    []Node
	isAttribute bool
	isCSS       bool
}

func (n *Node) Add(children ...Node) {
	n.children = append(n.children, children...)
}
