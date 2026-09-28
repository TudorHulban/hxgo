package dsl

func Img(children ...Node) Node {
	return el(
		"img",
		children...,
	)
}

func AttrImgSource(source string) Node {
	return AttrWithValue(
		"src",
		source,
	)
}

func AttrImgAlternativeText(text string) Node {
	return AttrWithValue(
		"alt",
		text,
	)
}
