package dsl

import (
	"strings"
	"unsafe"
)

type Transformer func(Node) Node

func DeleteTextNodes() Transformer {
	return func(n Node) Node {
		if n.isText {
			return Node{} // delete
		}

		out := n
		out.children = nil

		for _, child := range n.children {
			transformed := DeleteTextNodes()(child)
			if !transformed.IsZero() {
				out.children = append(out.children, transformed)
			}
		}

		return out
	}
}

func IsClassAttribute(n Node) bool {
	if !n.isAttribute {
		return false
	}

	name, _ := n.GetAttributeNameValue()

	return name == "class"
}

func DeleteNodesWhere(pred func(Node) bool) Transformer {
	return func(n Node) Node {
		if pred(n) {
			return Node{} // delete
		}

		out := n
		out.children = nil

		for _, child := range n.children {
			transformed := DeleteNodesWhere(pred)(child)
			if !transformed.IsZero() {
				out.children = append(out.children, transformed)
			}
		}

		return out
	}
}

func DeleteAttributesNamed(name string) Transformer {
	return func(n Node) Node {
		if n.isAttribute {
			attrName, _ := n.GetAttributeNameValue()
			if attrName == name {
				return Node{} // or a sentinel meaning “delete me”
			}
		}

		// Recurse into children
		out := n
		out.children = nil

		for _, child := range n.children {
			transformed := DeleteAttributesNamed(name)(child)
			if !transformed.IsZero() {
				out.children = append(out.children, transformed)
			}
		}

		return out
	}
}

func buildTailwindDSLText(methods []string) string {
	var b strings.Builder

	b.WriteString("TW()")

	for _, m := range methods {
		b.WriteString(".")
		b.WriteString(m)
		b.WriteString("()")
	}

	return b.String()
}

func TailwindTransformer(n Node) Node {
	// Do not transform attributes, text, or CSS nodes
	if n.isAttribute || n.isText || n.isCSS {
		return n
	}

	// Extract class attribute
	var classValue string
	for _, child := range n.children {
		if child.isAttribute {
			name, value := child.GetAttributeNameValue()
			if name == "class" {
				classValue = value
				break
			}
		}
	}

	if classValue == "" {
		return n
	}

	// Parse Tailwind classes
	classes := parseTailwindClasses(classValue)
	mapping := TW().mapping()

	methods := make([]string, 0, len(classes))
	for _, className := range classes {
		if methodName, ok := mapping[className]; ok {
			methods = append(methods, methodName)
		}
	}

	if len(methods) == 0 {
		return n
	}

	dslValue := buildTailwindDSLText(methods)

	// Build new children:
	// 1. DSL text node
	// 2. All original non-attribute children
	out := n
	// out.children = []Node{dslText}
	out.children = nil

	for _, child := range n.children {
		if child.isAttribute {
			name, _ := child.GetAttributeNameValue()
			if name == "class" {
				// Replace class attribute with Tailwind DSL
				out.children = append(out.children, AttrWithValue("class", dslValue))

				continue
			}
		}

		out.children = append(out.children, child)
	}

	return out
}

func ExtractAttributeValue(n Node, fn func(string) string) Node {
	// Only operate on attribute nodes
	if !n.isAttribute {
		return n
	}

	// Read the attribute value
	_, value := n.GetAttributeNameValue()

	// Transform the value
	newValue := fn(value)

	// Emit a new attribute-with-value node (name = "")
	data := &struct {
		name  string
		value string
	}{
		name:  "",
		value: newValue,
	}

	return Node{
		fn:          renderAttribute,
		data:        unsafe.Pointer(data),
		isAttribute: true,
	}
}
