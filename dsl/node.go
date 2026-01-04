package dsl

import "unsafe"

// Single builder used by the renderer.
// Accumulator accumulates HTML bytes and CSS styles during rendering.
// For better performance,
// use NewAcc with an estimated size based on the typical output.
type Accumulator struct {
	html []byte
	css  []Style
}

// NewAccumulator creates an Accumulator with pre-allocated capacity.
// estimatedHTMLSize should be the approximate size of the rendered HTML in bytes.
// Pre-allocating reduces memory allocations during rendering.
// Example: for a typical page of ~10KB, use NewAccumulator(10240, 16)
func NewAccumulator(estimatedHTMLSize, estimatedCSSRules int) *Accumulator {
	return &Accumulator{
		html: make([]byte, 0, estimatedHTMLSize),
		css:  make([]Style, 0, estimatedCSSRules),
	}
}

func (a *Accumulator) Write1(value string) {
	a.html = append(a.html, value...)
}

// Fused writes.
func (a *Accumulator) Write3(value1, value2, value3 string) {
	a.html = append(a.html, value1...)
	a.html = append(a.html, value2...)
	a.html = append(a.html, value3...)
}

// Fused writes.
func (a *Accumulator) Write5(value1, value2, value3, value4, value5 string) {
	a.html = append(a.html, value1...)
	a.html = append(a.html, value2...)
	a.html = append(a.html, value3...)
	a.html = append(a.html, value4...)
	a.html = append(a.html, value5...)
}

type Renderer func(*Accumulator, unsafe.Pointer)

// Node: function pointer plus data pointer.
type Node struct {
	fn          Renderer
	data        unsafe.Pointer
	children    []Node
	isAttribute bool
}
