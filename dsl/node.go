package dsl

type NodeOutput struct {
	IsAttr bool
	HTML   []byte
	Styles []Style
}

type Node func() NodeOutput
