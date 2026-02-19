package dsl

func walk(a *accumulator, n Node) {
	if n.fn == nil {
		return
	}

	n.fn(a, n.data)

	for i := range n.children {
		walk(a, n.children[i])
	}
}

// func walkHTML(a *accumulator, n Node) {
// 	if n.fn == nil {
// 		return
// 	}

// 	n.fn(a, n.data)

// 	for i := range n.children {
// 		if n.children[i].isCSS {
// 			continue
// 		}

// 		walkHTML(a, n.children[i])
// 	}
// }
