package dsl

import (
	"unsafe"
)

func Class(name string) Node {
	namePtr := new(string)
	*namePtr = name

	return Node{
		fn:   renderClass,
		data: unsafe.Pointer(namePtr),

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
