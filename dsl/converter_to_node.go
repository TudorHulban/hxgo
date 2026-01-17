package dsl

import (
	"strings"
	"unsafe"

	"golang.org/x/net/html"
)

// ConvertHTML calls convertElement for element nodes
// convertElement uses dataForTag to build the data structure with openTag/closeTag
// convertElement uses rendererForTag which returns renderElement for all tags
// walkHTML calls renderElement to write opening tag, then renderElementClose to write closing tag

func renderElementClose(a *accumulator, data unsafe.Pointer) {
	type tagData struct {
		openTag  []byte
		closeTag []byte
	}

	d := (*tagData)(data)

	a.html = append(a.html, d.closeTag...)
}

func renderText(a *accumulator, data unsafe.Pointer) {
	s := *(*string)(data)
	if len(s) == 0 {
		return
	}

	a.html = append(a.html, s...)
}

func convertText(n *html.Node) Node {
	trimmed := strings.TrimSpace(n.Data)
	if len(trimmed) == 0 {
		return Node{}
	}

	text := &trimmed

	return Node{
		fn:     renderText,
		data:   unsafe.Pointer(text),
		isText: true, // NEW
	}
}

func renderAttribute(a *accumulator, data unsafe.Pointer) {
	d := (*struct {
		key string
		val string
	})(data)

	a.html = append(a.html, ' ')
	a.html = append(a.html, d.key...)
	a.html = append(a.html, '=')
	a.html = append(a.html, '"')
	a.html = append(a.html, d.val...)
	a.html = append(a.html, '"')
}

func convertAttribute(a html.Attribute) Node {
	d := &struct {
		key string
		val string
	}{
		key: a.Key,
		val: a.Val,
	}

	return Node{
		fn:          renderAttribute,
		data:        unsafe.Pointer(d),
		isAttribute: true,
	}
}

func buildDataForTag(tag string) unsafe.Pointer {
	d := &struct {
		openTag  []byte
		closeTag []byte
	}{
		openTag:  append([]byte{'<'}, tag...),
		closeTag: append(append([]byte{'<', '/'}, tag...), '>'),
	}

	return unsafe.Pointer(d)
}

func convertElement(n *html.Node) Node {
	node := Node{
		fn:   renderElement,
		data: dataForTag(n.Data),
	}

	// attributes first
	for _, attr := range n.Attr {
		node.Add(convertAttribute(attr))
	}

	// children next
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		child := ConvertHTML(c)
		if child.fn != nil {
			node.Add(child)
		}
	}

	return node
}

func renderElement(a *accumulator, data unsafe.Pointer) {
	type tagData struct {
		openTag  []byte
		closeTag []byte
	}

	d := (*tagData)(data)
	a.html = append(a.html, d.openTag...)
}

func ConvertHTML(n *html.Node) Node {
	switch n.Type {
	case html.ElementNode:
		return convertElement(n)

	case html.TextNode:
		return convertText(n)

	case html.DocumentNode:
		// descend to first non-nil child with a renderer
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			child := ConvertHTML(c)
			if child.fn != nil {
				return child
			}
		}

		return Node{} // nothing useful found

	default:
		return Node{}
	}
}
