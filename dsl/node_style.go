package dsl

type Style struct {
	Selector string
	Props    map[string]string
	Media    string // optional media query
}

// Style nodes are not attributes.
// no HTML output, only styles.
func Styled(child Node, style ...Style) Node {
	return func() NodeOutput {
		out := child()
		out.Styles = append(out.Styles, style...)

		return out
	}
}
