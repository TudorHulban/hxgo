package dsl

import (
	"unsafe"
)

const _Class = "class"

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

	return name == _Class
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

func TailwindTransformer(n Node) Node {
	if n.isAttribute || n.isText || n.isCSS {
		return n
	}

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

	out := n
	out.children = nil

	// Emit each Tailwind method as a CSS node
	for _, m := range methods {
		cssNode := AttrWithValue(m, "")
		cssNode.isCSS = true // Mark as CSS so it's collected properly

		out.children = append(out.children, cssNode)
	}

	// Keep all non-class children
	for _, child := range n.children {
		if child.isAttribute {
			name, _ := child.GetAttributeNameValue()
			if name == _Class {
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
