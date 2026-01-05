package wfooter

import (
	"testing"

	pagecss "github.com/TudorHulban/hx-core/page-css"
	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/widgets/base"
	"github.com/stretchr/testify/require"
)

func TestFooter(t *testing.T) {
	el := WidgetFooter()

	page := components.Page{
		Title: t.Name(),

		Head: []dsl.Node{
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("generated.css"),
			),
		},

		Body: []dsl.Node{
			el,
		},
	}

	writer, errWriterCSS := helpers.GetFileWriter("generated.css")
	require.NoError(t, errWriterCSS)

	defer writer.Close()

	cssPage := pagecss.NewCSSPage(
		base.CSSBase,
		base.CSSSite,
		CSSWidgetFooter,
	)

	cssPage.GetCSSAccurateBeautifiedTo(
		writer,
		&pagecss.ParamsSpaces{
			NumberSpaces:              5,
			IncrementWithNumberSpaces: 2,
		},
	)

	writerHTML, errWriterHTML := helpers.GetFileWriter(t.Name() + ".html")
	require.NoError(t, errWriterHTML)

	defer writerHTML.Close()

	writer.Write(
		dsl.Render(
			page.Build(),
		),
	)
}
