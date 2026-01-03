package dsl

// TODO: add <header>, <nav>, <section>

func Article(children ...Node) Node {
	return El(
		"article",
		children...,
	)
}
