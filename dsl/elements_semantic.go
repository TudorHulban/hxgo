package dsl

// TODO: add <header>, <nav>, <section>

func Article(children ...Node) Writer {
	return El(
		"article",
		children...,
	)
}
