package dsl

import "unsafe"

// Accumulator: single builder used by the renderer
type Acc struct {
	html []byte
	css  []Style
}

func (a *Acc) Write(s string) {
	a.html = append(a.html, s...)
}

func (a *Acc) WriteBytes(b []byte) {
	a.html = append(a.html, b...)
}

type Accumulator func(*Acc, unsafe.Pointer)

// Node: function pointer + data pointer
type Node struct {
	fn          Accumulator
	data        unsafe.Pointer
	children    []Node
	isAttribute bool
}
