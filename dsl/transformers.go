package dsl

import (
	"unsafe"
)

const _Class = "class"

type Transformer func(Node) Node

func DeleteTextNodes() Transformer {
	return func(n Node) Node {
		if n.IsText {
			return Node{} // delete
		}

		out := n
		out.Children = nil

		for _, child := range n.Children {
			transformed := DeleteTextNodes()(child)
			if !transformed.IsZero() {
				out.Children = append(out.Children, transformed)
			}
		}

		return out
	}
}

func IsClassAttribute(n Node) bool {
	if !n.IsAttribute {
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
		out.Children = nil

		for _, child := range n.Children {
			transformed := DeleteNodesWhere(pred)(child)
			if !transformed.IsZero() {
				out.Children = append(out.Children, transformed)
			}
		}

		return out
	}
}

func DeleteAttributesNamed(name string) Transformer {
	return func(n Node) Node {
		if n.IsAttribute {
			attrName, _ := n.GetAttributeNameValue()
			if attrName == name {
				return Node{} // or a sentinel meaning “delete me”
			}
		}

		// Recurse into children
		out := n
		out.Children = nil

		for _, child := range n.Children {
			transformed := DeleteAttributesNamed(name)(child)
			if !transformed.IsZero() {
				out.Children = append(out.Children, transformed)
			}
		}

		return out
	}
}

func TailwindTransformer(n Node) Node {
	if n.IsAttribute || n.IsText || n.IsCSS {
		return n
	}

	var classValue string

	for _, child := range n.Children {
		if child.IsAttribute {
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
	out.Children = nil

	// Emit each Tailwind method as a CSS node
	for _, m := range methods {
		cssNode := AttrWithValue(m, "")
		cssNode.IsCSS = true // Mark as CSS so it's collected properly

		out.Children = append(out.Children, cssNode)
	}

	// Keep all non-class children
	for _, child := range n.Children {
		if child.IsAttribute {
			name, _ := child.GetAttributeNameValue()
			if name == _Class {
				continue
			}
		}

		out.Children = append(out.Children, child)
	}

	return out
}

func ExtractAttributeValue(n Node, fn func(string) string) Node {
	// Only operate on attribute nodes
	if !n.IsAttribute {
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
		Data:        unsafe.Pointer(data),
		IsAttribute: true,
	}
}
