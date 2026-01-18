package dsl

import (
	"io"
	"strconv"
)

type DSLPrinter struct {
	w io.Writer

	indentUnit  string
	indent      int
	isMultiLine bool
}

func printDSL(w io.Writer, n Node) {
	if n.fn == nil {
		return
	}

	p := DSLPrinter{
		w: w,
	}

	p.printNode(n)
}

func PrintDSL(w io.Writer, n Node, indentUnit ...string) {
	if n.fn == nil {
		return
	}

	unit := "    " // default: 4 spaces
	if len(indentUnit) > 0 {
		unit = indentUnit[0]
	}

	p := DSLPrinter{
		w:           w,
		indent:      0,
		indentUnit:  unit,
		isMultiLine: true,
	}

	p.printNode(n)
}

func (p *DSLPrinter) writeIndent() {
	for i := 0; i < p.indent; i++ {
		p.w.Write(
			[]byte(p.indentUnit),
		)
	}
}

func (p *DSLPrinter) printNode(n Node) {
	switch {
	case n.isText:
		p.printText(n)
	case n.isAttribute:
		p.printAttribute(n)
	default:
		p.printElement(n)
	}
}

func (p *DSLPrinter) printElement(n Node) {
	tag := extractTagName(n.data)

	p.writeIndent()
	p.w.Write([]byte(_DSL))
	p.w.Write([]byte("."))
	p.w.Write([]byte(capitalize(tag)))
	p.w.Write([]byte("("))

	if p.isMultiLine {
		p.w.Write([]byte("\n"))
	}

	p.indent++

	for i := range n.children {
		p.printNode(n.children[i])
		p.w.Write([]byte(","))

		if p.isMultiLine {
			p.w.Write([]byte("\n"))
		}
	}

	p.indent--

	p.writeIndent()
	p.w.Write([]byte(")"))
}

func (p *DSLPrinter) printAttribute(n Node) {
	d := (*struct {
		key string
		val string
	})(n.data)

	p.writeIndent()
	p.w.Write([]byte(_DSL))
	p.w.Write([]byte("."))
	p.w.Write([]byte(capitalize(d.key)))
	p.w.Write([]byte("("))
	p.w.Write([]byte(strconv.Quote(d.val)))
	p.w.Write([]byte(")"))
}

func (p *DSLPrinter) printText(n Node) {
	s := *(*string)(n.data)

	p.writeIndent()
	p.w.Write([]byte(_DSL))
	p.w.Write([]byte(".Text("))
	p.w.Write([]byte(strconv.Quote(s)))
	p.w.Write([]byte(")"))
}
