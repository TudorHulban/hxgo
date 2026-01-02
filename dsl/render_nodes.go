package dsl

import (
	"bytes"
	"io"
)

// Empty nodes still emit tags because the renderer treats
// the element itself as meaningful, not just its children.
// func renderNodes(w io.Writer, elName string, nodes ...Node) ([]Style, error) {
// 	var styles []Style

// 	// Buffers
// 	var attrsBuf bytes.Buffer
// 	var childrenBuf bytes.Buffer

// 	// First pass: call each node ONCE
// 	for _, node := range nodes {
// 		if node == nil {
// 			continue
// 		}

// 		// Each node writes into its own temp buffer
// 		var tmp bytes.Buffer
// 		isAttr, st, err := node(&tmp)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if isAttr {
// 			// attribute → write into attrs buffer
// 			attrsBuf.Write(tmp.Bytes())
// 		} else {
// 			// child → write into children buffer
// 			childrenBuf.Write(tmp.Bytes())
// 		}

// 		// collect styles
// 		styles = append(styles, st...)
// 	}

// 	// Now write final HTML
// 	io.WriteString(w, "<")
// 	io.WriteString(w, elName)
// 	w.Write(attrsBuf.Bytes())
// 	io.WriteString(w, ">")
// 	w.Write(childrenBuf.Bytes())
// 	io.WriteString(w, "</")
// 	io.WriteString(w, elName)
// 	io.WriteString(w, ">")

// 	return styles, nil
// }

func renderNodes(w io.Writer, elName string, nodes ...Node) ([]Style, error) {
	var styles []Style
	var nodeBuf bytes.Buffer

	// Collect output with type info in one pass
	type nodeOutput struct {
		data   []byte
		isAttr bool
	}
	outputs := make([]nodeOutput, 0, len(nodes))

	for _, node := range nodes {
		if node == nil {
			continue
		}

		nodeBuf.Reset()
		isAttr, st, err := node(&nodeBuf)
		if err != nil {
			return nil, err
		}

		// Copy bytes once
		data := make([]byte, nodeBuf.Len())
		copy(data, nodeBuf.Bytes())
		outputs = append(outputs, nodeOutput{data, isAttr})

		styles = append(styles, st...)
	}

	// Write final HTML
	io.WriteString(w, "<")
	io.WriteString(w, elName)

	for _, out := range outputs {
		if out.isAttr {
			w.Write(out.data)
		}
	}

	io.WriteString(w, ">")

	for _, out := range outputs {
		if !out.isAttr {
			w.Write(out.data)
		}
	}

	io.WriteString(w, "</")
	io.WriteString(w, elName)
	io.WriteString(w, ">")

	return styles, nil
}

func renderNodesWithCSSId(w io.Writer, elName, cssID string, nodes ...Node) ([]Style, error) {
	var styles []Style

	// Buffers
	var attrsBuf bytes.Buffer
	var childrenBuf bytes.Buffer

	// First pass: call each node ONCE
	for _, node := range nodes {
		if node == nil {
			continue
		}

		// Each node writes into its own temp buffer
		var tmp bytes.Buffer
		isAttr, st, err := node(&tmp)
		if err != nil {
			return nil, err
		}

		if isAttr {
			// attribute → write into attrs buffer
			attrsBuf.Write(tmp.Bytes())
		} else {
			// child → write into children buffer
			childrenBuf.Write(tmp.Bytes())
		}

		// collect styles
		styles = append(styles, st...)
	}

	// Now write final HTML
	AttrIDLength(cssID)(w)

	io.WriteString(w, "<")
	io.WriteString(w, elName)
	w.Write(attrsBuf.Bytes())
	io.WriteString(w, ">")
	w.Write(childrenBuf.Bytes())
	io.WriteString(w, "</")
	io.WriteString(w, elName)
	io.WriteString(w, ">")

	return styles, nil
}
