package dsl

import (
	"fmt"
	"strings"
	"unsafe"
)

type renderer func(*accumulator, unsafe.Pointer)

// Node: function pointer plus data pointer.
//
// Data-oriented design.
// Manual control over execution.
// No interface dispatch.
// No virtual method tables.
type Node struct {
	fn   renderer
	data unsafe.Pointer

	children []Node

	isAttribute bool
	isCSS       bool
	isText      bool
}

func (n Node) IsZero() bool {
	return n.fn == nil &&
		n.data == nil &&
		len(n.children) == 0 &&
		!n.isAttribute &&
		!n.isCSS &&
		!n.isText
}

func (n *Node) Add(children ...Node) {
	n.children = append(n.children, children...)
}

func (n Node) GetTagName() string {
	if n.isAttribute || n.isText || n.isCSS {
		return ""
	}

	name := (*struct{ name string })(n.data).name

	return strings.TrimPrefix(name, "<")
}

func (n Node) GetAttributeNameValue() (string, string) {
	if !n.isAttribute {
		return "", ""
	}

	d := (*struct {
		name  string
		value string
	})(n.data)

	return d.name, d.value
}

func (n *Node) Canonical() string {
	switch {
	case n.isText:
		s := *(*string)(n.data)

		return fmt.Sprintf(`text(%q)`, s)

	case n.isAttribute:
		v := *(*string)(n.data)

		return fmt.Sprintf(`class="%s"`, v)

	case n.isCSS:
		d := (*struct {
			name  string
			value string
		})(n.data)

		return fmt.Sprintf("%s:%s", d.name, d.value)

	default:
		d := (*struct {
			tag      string
			attrs    []Node
			children []Node
		})(n.data)

		var attrs []string
		for i := range d.attrs {
			attrs = append(
				attrs,
				d.attrs[i].Canonical(),
			)
		}

		var kids []string
		for i := range d.children {
			kids = append(
				kids,
				d.children[i].Canonical(),
			)
		}

		if len(attrs) == 0 && len(kids) == 0 {
			return d.tag
		}

		if len(attrs) == 0 {
			return fmt.Sprintf(
				"%s(%s)",

				d.tag,
				strings.Join(kids, ", "),
			)
		}

		if len(kids) == 0 {
			return fmt.Sprintf(
				"%s[%s]",

				d.tag,
				strings.Join(attrs, ", "),
			)
		}

		return fmt.Sprintf(
			"%s[%s](%s)",

			d.tag,
			strings.Join(attrs, ", "),
			strings.Join(kids, ", "),
		)
	}
}
