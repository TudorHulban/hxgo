package dsl

import "unsafe"

// Elements are never attributes
func El(tag string, children ...Node) Node {
	return Node{
		fn:       renderEl,
		data:     unsafe.Pointer(&tag),
		children: children,
	}
}

func renderEl(a *Acc, p unsafe.Pointer) {
	tag := *(*string)(p)

	a.Write("<")
	a.Write(tag)
	a.Write(">")

	// children are rendered by walk()

	a.Write("</")
	a.Write(tag)
	a.Write(">")
}

func ElWId(tag, cssID string, children ...Node) Node {
	// pack both strings into a tiny struct so we pass ONE pointer
	data := struct {
		tag   string
		cssID string
	}{tag, cssID}

	return Node{
		fn:       renderElWId,
		data:     unsafe.Pointer(&data),
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
