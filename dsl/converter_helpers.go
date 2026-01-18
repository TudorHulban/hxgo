package dsl

import (
	"strings"
	"unsafe"

	"golang.org/x/net/html"
)

func hasClass(n *html.Node, class string) bool {
	if n.Type != html.ElementNode {
		return false
	}

	for _, a := range n.Attr {
		if a.Key == "class" {
			parts := strings.Fields(a.Val)
			for i := range parts {
				if parts[i] == class {
					return true
				}
			}
		}
	}

	return false
}

func hasID(n *html.Node, id string) bool {
	if n.Type != html.ElementNode {
		return false
	}

	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return true
		}
	}

	return false
}

func isTriviallyEmpty(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return false
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case html.TextNode:
			if strings.TrimSpace(c.Data) != "" {
				return false
			}

		case html.ElementNode:
			return false
		}
	}

	return true
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func extractTagName(data unsafe.Pointer) string {
	d := (*struct {
		openTag  []byte
		closeTag []byte
	})(data)

	// openTag = "<div"
	// skip '<'
	return string(d.openTag[1:])
}
