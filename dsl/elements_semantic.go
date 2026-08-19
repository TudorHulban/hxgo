package dsl

func Article(children ...Node) Node {
	return el(
		"article",
		children...,
	)
}

// The <header> HTML element represents introductory content,
// typically a group of introductory or navigational aids.
// It may contain some heading elements but also a logo,
// a search form, an author name, and other elements.
func Header(children ...Node) Node {
	return el(
		"header",
		children...,
	)
}

// The <nav> HTML element represents a section of a page
// whose purpose is to provide navigation links,
// either within the current document or to other documents.
// Common examples of navigation sections are menus,
// tables of contents, and indexes.
func Nav(children ...Node) Node {
	return el(
		"nav",
		children...,
	)
}

// The <section> HTML element represents a generic standalone section of a document,
// which doesn't have a more specific semantic element to represent it.
// Sections should always have a heading, with very few exceptions.
func Section(children ...Node) Node {
	return el(
		"section",
		children...,
	)
}
