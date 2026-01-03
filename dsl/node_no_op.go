package dsl

// no‑op node that is NOT a Node.
var Noop Node = func() NodeOutput {
	return NodeOutput{}
}
