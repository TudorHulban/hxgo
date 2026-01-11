package footer

import (
	"testing"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/widgets/base"
	"github.com/stretchr/testify/require"
)

func TestFooter(t *testing.T) {
	page := components.Page{
		Title: t.Name(),

		Head: []dsl.Node{
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("generated.css"),
			),
		},

		Body: []dsl.Node{
			WidgetFooter(),
		},
	}

	writerCSS, errWriterCSS := helpers.GetFileWriter("generated.css")
	require.NoError(t, errWriterCSS)

	defer writerCSS.Close()

	cssContribution := dsl.CSSContribution{
		ProceduralCSS: []dsl.CSS{
			base.CSSBase,
			base.CSSSite,
		},
	}

	el := page.Build()
	el.Add(
		cssContribution.AsNode(),
	)

	writerHTML, errWriterHTML := helpers.GetFileWriter(t.Name() + ".html")
	require.NoError(t, errWriterHTML)

	defer writerHTML.Close()

	html, styles, css := dsl.RenderFull(el)
	require.Zero(t, styles)
	require.NotZero(t, html)
	require.NotZero(t, css)

	writerHTML.Write(
		html,
	)

	writerCSS.Write(
		[]byte(css),
	)
}
