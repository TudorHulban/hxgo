package dsl

// no‑op node that is NOT a Node.
var noop Node = func() NodeOutput {
	return NodeOutput{
		IsAttr: false,
		HTML:   nil,
		Styles: nil,
	}
}
