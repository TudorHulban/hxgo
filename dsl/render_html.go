package dsl

import (
	"bytes"
	"encoding/base64"

	"github.com/TudorHulban/hxgo/helpers"
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

func HTMLwithDataCSS(html []byte, css string) string {
	return helpers.Sprintf(
		`<html><head><link rel="stylesheet" href="data:text/css;base64,%s"></head><body>%s</body></html>`,

		base64.StdEncoding.EncodeToString([]byte(css)),
		string(html),
	)
}
