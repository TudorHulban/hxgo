package dsl

import "unsafe"

func Class(name string) Node {
	return Node{
		fn:   renderClass,
		data: unsafe.Pointer(&name),

		isAttribute: true,
	}
}

// TODO: move to write3?
func renderClass(a *accumulator, p unsafe.Pointer) {
	name := *(*string)(p)

	a.Write1(` class="`)
	a.Write1(name)
	a.Write1(`"`)
}
