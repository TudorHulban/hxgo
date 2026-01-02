package dsl

import "io"

type Style struct {
	Selector string
	Props    map[string]string
	Media    string // optional media query
}

// Style nodes are not attributes.
// no HTML output, only styles.
func Styled(styles ...Style) Node {
	return func(w io.Writer) (bool, []Style, error) {
		return false, styles, nil
	}
}
