package dsl

func Article(children ...Node) Writer {
	return El(
		"article",
		children...,
	)
}
