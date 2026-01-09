package dsl

import "unsafe"

// Elements are never attributes.
// Children slice should be immutable.
//
//go:inline
func el(tag string, children ...Node) Node {
	var attrs []Node // Separate attributes from children during construction

	var nodes []Node

	for i := range children {
		if children[i].isAttribute {
			attrs = append(attrs, children[i])
		} else {
			nodes = append(nodes, children[i])
		}
	}

	data := &struct {
		tag      string
		attrs    []Node
		children []Node
	}{tag, attrs, nodes}

	return Node{
		fn:   renderEl,
		data: unsafe.Pointer(data),
	}
}

func renderEl(a *accumulator, p unsafe.Pointer) {
	d := (*struct {
		tag      string
		attrs    []Node
		children []Node
	})(p)

	a.Write1("<")
	a.Write1(d.tag)

	// Render attributes - single iteration
	for i := range d.attrs {
		d.attrs[i].fn(a, d.attrs[i].data)
	}

	a.Write1(">")

	// Render children - single iteration
	for i := range d.children {
		walk(a, d.children[i])
	}

	a.Write3("</", d.tag, ">")
}

// Children slice should be immutable.
func elWId(tag, cssID string, children ...Node) Node {
	// Separate attributes from children during construction
	var attrs []Node

	var nodes []Node

	for i := range children {
		if children[i].isAttribute {
			attrs = append(attrs, children[i])
		} else {
			nodes = append(nodes, children[i])
		}
	}

	data := &struct {
		tag      string
		cssID    string
		attrs    []Node
		children []Node
	}{tag, cssID, attrs, nodes}

	return Node{
		fn:   renderElWId,
		data: unsafe.Pointer(data),
	}
}

func renderElWId(a *accumulator, p unsafe.Pointer) {
	d := (*struct {
		tag      string
		cssID    string
		attrs    []Node
		children []Node
	})(p)

	a.Write1("<")
	a.Write1(d.tag)

	if d.cssID != "" {
		a.Write1(` id="`)
		a.Write1(d.cssID)
		a.Write1(`"`)
	}

	// Render attributes - single iteration
	for i := range d.attrs {
		d.attrs[i].fn(a, d.attrs[i].data)
	}

	a.Write1(">")

	// Render children - single iteration
	for i := range d.children {
		walk(a, d.children[i])
	}

	a.Write1("</")
	a.Write1(d.tag)
	a.Write1(">")
}
