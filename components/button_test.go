package components

import (
	"fmt"
	"testing"

	"github.com/TudorHulban/hxgo/dsl"
	"github.com/stretchr/testify/require"
)

func TestButton(t *testing.T) {
	el := ButtonSubmitWCSS(
		&ParamsButtonSubmit{},
		[]dsl.Node{
			dsl.NodeStyle{
				Styles: []dsl.Style{
					{
						Selector: ".card:hover",
						Props: map[string]string{
							"box-shadow": "0 8px 24px rgba(0,0,0,0.2)",
						},
					},
				},
			},
		},
	)

	html, css, errRender := dsl.RenderHTMLandCSS(el)
	require.NoError(t, errRender)
	require.NotZero(t, html, "valid HTML")
	require.NotZero(t, css, "valid CSS")

	fmt.Println(
		string(html),
	)
	fmt.Println(css)
}
