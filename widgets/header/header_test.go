package header

import (
	"testing"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/stretchr/testify/require"
)

func TestHeader(t *testing.T) {
	fragment := WidgetHeader(
		&ParamsWidgetHeader{
			Title:             "Washing Head",
			BreadcrumbCaption: "Barbershop/Washing Head",

			ParamsImage: components.ParamsImage{
				ImageSquareSize: "160",
				ImageAlt:        "Washing Head",
				ImageSource:     "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/hero-bg-1.jpg",
			},
		},
	)

	writer, errWriter := helpers.GetFileWriter(t.Name() + ".html")
	require.NoError(t, errWriter)

	defer writer.Close()

	page := components.Page{
		Title: t.Name(),

		Head: []dsl.Node{
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("css_base.css"),
			),
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("css_site.css"),
			),
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("header.css"),
			),
		},

		Body: []dsl.Node{
			fragment,
		},
	}

	writer.Write(
		dsl.Render(
			page.Build(),
		),
	)
}
