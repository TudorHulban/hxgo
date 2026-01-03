package dsl

import "unsafe"

// ------------------------------------------------------------
// Accumulator: single builder used by the renderer
// ------------------------------------------------------------

type Acc struct {
	Buf    []byte
	Styles []Style
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
