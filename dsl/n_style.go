package dsl

import "io"

type Style struct {
	Selector string
	Props    map[string]string
	Media    string // optional media query
}

type NodeStyle struct {
	Styles []Style
}

func Styled(styles ...Style) Node {
	return &NodeStyle{Styles: styles}
}

func (n NodeStyle) isAttribute() bool { return false }

func (n NodeStyle) Render(w io.Writer) ([]Style, error) {
	return n.Styles, nil
}
