package dsl

import (
	"strings"
	"unsafe"
)

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

	children []Node

	isAttribute bool
	isCSS       bool
	isText      bool
}

func (n Node) IsZero() bool {
	return n.fn == nil &&
		n.data == nil &&
		len(n.children) == 0 &&
		!n.isAttribute &&
		!n.isCSS &&
		!n.isText
}

func (n *Node) Add(children ...Node) {
	n.children = append(n.children, children...)
}

func (n Node) GetTagName() string {
	if n.isAttribute || n.isText || n.isCSS {
		return ""
	}

	name := (*struct{ name string })(n.data).name

	return strings.TrimPrefix(name, "<")
}

func (n Node) GetAttributeNameValue() (string, string) {
	if !n.isAttribute {
		return "", ""
	}

	d := (*struct {
		name  string
		value string
	})(n.data)

	return d.name, d.value
}
