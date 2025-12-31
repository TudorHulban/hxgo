package dsl

import (
	"bytes"
)

func RenderHTMLandCSS(nodes ...Node) ([]byte, string, error) {
	var htmlBuf bytes.Buffer
	collector := newSmartCSSCollector()

	for _, node := range nodes {
		if node == nil {
			continue
		}

		// Single pass: render HTML and get styles
		styles, errRender := node.Render(&htmlBuf)
		if errRender != nil {
			return nil, "", errRender
		}

		for _, s := range styles {
			collector.Add(s)
		}
	}

	return htmlBuf.Bytes(), collector.String(), nil
}
