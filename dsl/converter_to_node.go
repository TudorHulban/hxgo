package dsl

import (
	"strings"
	"unsafe"

	"golang.org/x/net/html"
)

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

	return Node{
		fn:   renderText,
		data: unsafe.Pointer(&trimmed),
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

func rendererForTag(tag string) renderer {
	switch tag {
	case "div":
		return renderDiv
	case "span":
		return renderSpan
	case "a":
		return renderA
	case "p":
		return renderP
	case "img":
		return renderImg
	case "ul":
		return renderUl
	case "ol":
		return renderOl
	case "li":
		return renderLi
	case "nav":
		return renderNav
	case "h1":
		return renderH1
	case "h2":
		return renderH2
	case "h3":
		return renderH3
	case "h4":
		return renderH4
	case "h5":
		return renderH5
	case "h6":
		return renderH6
	}

	// fallback for unknown tags
	return renderGenericElement
}

func buildDataForTag(tag string) unsafe.Pointer {
	s := tag // allocate a copy so the pointer remains stable

	return unsafe.Pointer(&s)
}

func convertElement(n *html.Node) Node {
	tag := n.Data

	node := Node{
		fn:   rendererForTag(tag),
		data: buildDataForTag(tag),
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
