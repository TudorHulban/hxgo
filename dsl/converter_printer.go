package dsl

import (
	"strconv"
	"strings"
	"unsafe"
)

func extractTagName(data unsafe.Pointer) string {
	d := (*struct {
		openTag  []byte
		closeTag []byte
	})(data)

	// openTag = "<div"
	// skip '<'
	return string(d.openTag[1:])
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func printElement(n Node) string {
	tag := extractTagName(n.data)

	var b strings.Builder

	b.WriteString(_DSL)
	b.WriteString(".")
	b.WriteString(
		capitalize(tag),
	)
	b.WriteString("(")

	for i := range n.children {
		b.WriteString(PrintDSL(n.children[i]))
		b.WriteString(",")
	}

	b.WriteString(")")

	return b.String()
}

func printAttribute(n Node) string {
	d := (*struct {
		key string
		val string
	})(n.data)

	var b strings.Builder

	b.WriteString(_DSL)
	b.WriteString(".")
	b.WriteString(capitalize(d.key))
	b.WriteString("(")
	b.WriteString(strconv.Quote(d.val))
	b.WriteString(")")

	return b.String()
}

func printText(n Node) string {
	s := *(*string)(n.data)

	var b strings.Builder

	b.WriteString(_DSL)
	b.WriteString(".Text(")
	b.WriteString(strconv.Quote(s))
	b.WriteString(")")

	return b.String()
}

func PrintDSL(n Node) string {
	if n.fn == nil {
		return ""
	}

	if n.isAttribute {
		return printAttribute(n)
	}

	if n.isText {
		return printText(n)
	}

	return printElement(n)
}
