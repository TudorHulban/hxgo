package dsl

import "unsafe"

// Elements are never attributes.
// Children slice should be immutable.
func El(tag string, children ...Node) Node {
	data := &struct { // intentionally move to heap.
		tag      string
		children []Node
	}{tag, children}

	return Node{
		fn:   renderEl,
		data: unsafe.Pointer(data),
	}
}

func renderEl(a *Acc, p unsafe.Pointer) {
	d := (*struct {
		tag      string
		children []Node
	})(p)

	a.Write("<")
	a.Write(d.tag)

	// 1. render attributes
	for i := range d.children {
		if d.children[i].isAttribute {
			d.children[i].fn(a, d.children[i].data)
		}
	}

	a.Write(">")

	// 2. render normal children
	for i := range d.children {
		if !d.children[i].isAttribute {
			walk(a, d.children[i])
		}
	}

	a.Write("</")
	a.Write(d.tag)
	a.Write(">")
}

// Children slice should be immutable.
func ElWId(tag, cssID string, children ...Node) Node {
	// pack both strings into a tiny struct so we pass ONE pointer
	data := &struct { // intentionally move to heap.
		tag   string
		cssID string
	}{tag, cssID}

	return Node{
		fn:       renderElWId,
		data:     unsafe.Pointer(data),
		children: children,
	}
}

func renderElWId(a *Acc, p unsafe.Pointer) {
	d := (*struct {
		tag   string
		cssID string
	})(p)

	a.Write("<")
	a.Write(d.tag)

	if d.cssID != "" {
		a.Write(` id="`)
		a.Write(d.cssID)
		a.Write(`"`)
	}

	a.Write(">")

	// children rendered by walk()

	a.Write("</")
	a.Write(d.tag)
	a.Write(">")
}
