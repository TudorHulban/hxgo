package card

import (
	"testing"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/stretchr/testify/require"
)

func TestVerticalCard(t *testing.T) {
	fragment := WidgetCardVertical(
		&ParamsWidgetCardVertical{
			WidgetCardVerticalInfo: WidgetCardVerticalInfo{
				Title: "Jumpin Jack",
				Price: "4000",

				ImageSquareSize: "160",
				ImageSource:     "https://images.pexels.com/photos/19040825/pexels-photo-19040825.jpeg",

				ActionEndpoint: "TBD",
			},

			CurrencySimbol: "RON",
			PriceCaption:   "Price",
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
				dsl.Href("card.css"),
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
