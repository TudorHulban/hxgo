package wfooter

import (
	"fmt"
	"testing"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	pagecss "github.com/TudorHulban/hx-core/page-css"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	"github.com/TudorHulban/hx-widgets/base"
	"github.com/TudorHulban/hx-widgets/helpers"
	"github.com/stretchr/testify/require"
)

func TestFooter(t *testing.T) {
	fragment := WidgetFooter()

	page := hxcomponents.Page{
		Title: t.Name(),

		Head: []hxprimitives.Node{
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("generated.css"),
			),
		},

		Body: []hxprimitives.Node{
			fragment,
		},
	}

	writerCSS, errWriterCSS := helpers.GetFileWriter("generated.css")
	require.NoError(t, errWriterCSS)

	defer writerCSS.Close()

	cssPage := pagecss.NewCSSPage(
		base.CSSBase,
		base.CSSSite,
		CSSWidgetFooter,
	)

	cssPage.GetCSSAccurateBeautifiedTo(
		writerCSS,
		&pagecss.ParamsSpaces{
			NumberSpaces:              5,
			IncrementWithNumberSpaces: 2,
		},
	)

	writerHTML, errWriterHTML := helpers.GetFileWriter(t.Name() + ".html")
	require.NoError(t, errWriterHTML)

	defer writerHTML.Close()

	fmt.Println(
		page.Build().Render(writerHTML),
	)
}
