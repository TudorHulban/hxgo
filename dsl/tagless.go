package dsl

var tagless = map[string]struct{}{
	"area":   {},
	"br":     {},
	"hr":     {},
	"img":    {},
	"input":  {},
	"link":   {},
	"meta":   {},
	"source": {},
}

// TODO: should be used.
func isTagless(elName string) bool {
	_, exists := tagless[elName]

	return exists
}
