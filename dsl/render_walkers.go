package dsl

func walkFull(a *accumulator, n Node) {
	if n.fn == nil {
		return
	}

	n.fn(a, n.Data)

	for i := range n.Children {
		walkFull(a, n.Children[i])
	}
}

func walkHTML(a *accumulator, n Node) {
	if n.fn == nil {
		return
	}

	n.fn(a, n.Data)

	for i := range n.Children {
		if n.Children[i].IsCSS {
			continue
		}

		walkHTML(a, n.Children[i])
	}
}
