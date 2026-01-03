package dsl

import "strings"

type NodeOutput struct {
	IsAttr    bool
	HTMLParts [][]byte // slice of references to static or prebuilt byte slices.
	Styles    []Style
}

func (n NodeOutput) String() string {
	var b strings.Builder

	for _, part := range n.HTMLParts {
		b.Write(part)
	}

	return b.String()
}

type Node func() NodeOutput
