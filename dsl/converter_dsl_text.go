package dsl

import "unsafe"

func renderDSLText(a *accumulator, p unsafe.Pointer) {
	d := (*struct{ value string })(p)

	a.Write1(d.value)
}

func DSLText(s string) Node {
	data := &struct{ value string }{value: s}

	return Node{
		fn:   renderDSLText,
		data: unsafe.Pointer(data),
	}
}
