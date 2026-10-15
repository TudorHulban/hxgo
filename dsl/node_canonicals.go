package dsl

import (
	"fmt"
	"strconv"
	"strings"
)

// Canonical is to be used for debugging.
// For responses use a render.
func (n Node) Canonical() string {
	switch {
	case n.isText:
		s := *(*string)(n.data)

		return fmt.Sprintf(`text(%q)`, s)

	case n.isRaw:
		s := *(*string)(n.data)

		return fmt.Sprintf(`raw(%q)`, s)

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

		for ix := range d.children {
			kids = append(
				kids,
				d.children[ix].Canonical(),
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

// CanonicalInto appends the canonical form of n to dst and returns the
// extended slice. Output is byte-for-byte identical to Canonical().
//
// Canonical is to be used for debugging.
// For responses use a render.
func (n Node) CanonicalInto(destination []byte) []byte {
	switch {
	case n.isText:
		s := *(*string)(n.data)

		destination = append(destination, "text("...)
		destination = strconv.AppendQuote(destination, s) // same escaping as fmt's %q
		destination = append(destination, ')')

	case n.isRaw:
		s := *(*string)(n.data)

		destination = append(destination, "raw("...)
		destination = strconv.AppendQuote(destination, s) // same escaping as fmt's %q
		destination = append(destination, ')')

	case n.isAttribute:
		v := *(*string)(n.data)

		destination = append(destination, `class="`...)
		destination = append(destination, v...)
		destination = append(destination, '"')

	case n.isCSS:
		d := (*struct {
			name  string
			value string
		})(n.data)

		destination = append(destination, d.name...)
		destination = append(destination, ':')
		destination = append(destination, d.value...)

	default:
		d := (*struct {
			tag      string
			attrs    []Node
			children []Node
		})(n.data)

		destination = append(destination, d.tag...)

		if len(d.attrs) > 0 {
			destination = append(destination, '[')

			for i := range d.attrs {
				if i > 0 {
					destination = append(destination, ", "...)
				}

				destination = d.attrs[i].CanonicalInto(destination)
			}

			destination = append(destination, ']')
		}

		if len(d.children) > 0 {
			destination = append(destination, '(')

			for i := range d.children {
				if i > 0 {
					destination = append(destination, ", "...)
				}

				destination = d.children[i].CanonicalInto(destination)
			}

			destination = append(destination, ')')
		}
	}

	return destination
}
