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
	case html.DocumentNode:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				return ConvertHTML(c)
			}
		}

		return Node{}

	case html.ElementNode:
		// unwrap <html>
		if n.Data == "html" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "body" {
					return ConvertHTML(c)
				}
			}

			return Node{}
		}

		// unwrap <body> with semantic filtering
		if n.Data == "body" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type != html.ElementNode {
					continue
				}

				// skip known boilerplate
				if hasID(c, "loading") {
					continue
				}

				if hasClass(c, "loader") || hasClass(c, "simple-loader") {
					continue
				}

				// skip empty wrappers
				if isTriviallyEmpty(c) {
					continue
				}

				return ConvertHTML(c)
			}

			return Node{}
		}

		return convertElement(n)

	case html.TextNode:
		return convertText(n)

	default:
		return Node{}
	}
}
