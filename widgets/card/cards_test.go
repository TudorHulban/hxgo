package wcard

import (
	"testing"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/stretchr/testify/require"
)

func TestCards(t *testing.T) {
	img1Src := "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-1.jpg"
	img2Src := "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-2.jpg"
	img3Src := "https://images.pexels.com/photos/897262/pexels-photo-897262.jpeg"
	img4Src := "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-4.jpg"

	action1 := "https://themes.getmotopress.com/bro-barbershop/service/washing-head/"
	action2 := "https://themes.getmotopress.com/bro-barbershop/service/beard-trim/"
	action3 := "https://themes.getmotopress.com/bro-barbershop/service/mustache-haircut/"
	action4 := "https://themes.getmotopress.com/bro-barbershop/service/mens-haircut/"

	cards := WidgetCards(
		&ParamsWidgetCards{
			CurrencySimbol: "RON",
			PriceCaption:   "Pret",

			CSSFlexGap: "20px",

			Cards: []*WidgetCardVerticalInfo{
				{
					Title: "Washing Head",
					Price: "49",

					ImageSquareSize: "220",
					ImageSource:     img1Src,

					ActionEndpoint: action1,
				},
				{
					Title: "Beard Trim",
					Price: "39",

					ImageSquareSize: "220",
					ImageSource:     img2Src,

					ActionEndpoint: action2,
				},
				{
					Title: "Mustache Haircut",
					Price: "29",

					ImageSquareSize: "220",
					ImageSource:     img3Src,

					ActionEndpoint: action3,
				},
				{
					Title: "Men's Haircut",
					Price: "59",

					ImageSquareSize: "220",
					ImageSource:     img4Src,

					ActionEndpoint: action4,
				},
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
				dsl.Href("card.css"),
			),
		},

		Body: []dsl.Node{
			cards,
		},
	}

	writer.Write(
		dsl.Render(
			page.Build(),
		),
	)
}
