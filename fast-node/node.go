package fastnode

import (
	"unsafe"
)

// ------------------------------------------------------------
// Accumulator: single builder used by the renderer
// ------------------------------------------------------------

type Acc struct {
	Buf []byte
}

func (a *Acc) Write(s string) {
	a.Buf = append(a.Buf, s...)
}

func (a *Acc) WriteBytes(b []byte) {
	a.Buf = append(a.Buf, b...)
}

// ------------------------------------------------------------
// Node: function pointer + data pointer
// ------------------------------------------------------------

type Accumulator func(*Acc, unsafe.Pointer)

type Node struct {
	fn       Accumulator
	data     unsafe.Pointer
	children []Node
}

// ------------------------------------------------------------
// Renderer: single-pass traversal
// ------------------------------------------------------------

func Render(nodes ...Node) []byte {
	var a Acc
	for i := range nodes {
		walk(&a, nodes[i])
	}
	return a.Buf
}

func walk(a *Acc, n Node) {
	n.fn(a, n.data)
	for i := range n.children {
		walk(a, n.children[i])
	}
}

// ------------------------------------------------------------
// Helpers
// ------------------------------------------------------------

func Text(s string) Node {
	return Node{
		fn:   renderText,
		data: unsafe.Pointer(&s),
	}
}

func renderText(a *Acc, p unsafe.Pointer) {
	s := *(*string)(p)
	a.Write(s)
}

func Class(name string) Node {
	return Node{
		fn:   renderClass,
		data: unsafe.Pointer(&name),
	}
}

func renderClass(a *Acc, p unsafe.Pointer) {
	name := *(*string)(p)
	a.Write(` class="`)
	a.Write(name)
	a.Write(`"`)
}

func Div(children ...Node) Node {
	return Node{
		fn:       renderDiv,
		data:     nil,
		children: children,
	}
}

func renderDiv(a *Acc, _ unsafe.Pointer) {
	a.Write("<div")
	// attributes would be rendered by children before closing '>'
	a.Write(">")
	// children rendered by walk()
	a.Write("</div>")
}
