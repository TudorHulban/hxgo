package widgets

import (
	"testing"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	wcard "github.com/TudorHulban/hxgo/widgets/card"
	wheader "github.com/TudorHulban/hxgo/widgets/header"
	whero "github.com/TudorHulban/hxgo/widgets/hero"
	"github.com/stretchr/testify/require"
)

func Page() dsl.Node {
	img1Src := "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-1.jpg"
	img2Src := "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-2.jpg"
	img3Src := "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-3.jpg"
	img4Src := "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/service-4.jpg"

	action1 := "https://themes.getmotopress.com/bro-barbershop/service/washing-head/"
	action2 := "https://themes.getmotopress.com/bro-barbershop/service/beard-trim/"
	action3 := "https://themes.getmotopress.com/bro-barbershop/service/mustache-haircut/"
	action4 := "https://themes.getmotopress.com/bro-barbershop/service/mens-haircut/"

	return dsl.Div(
		wheader.WidgetHeader(
			&wheader.ParamsWidgetHeader{
				Title: "Barber Shop",

				ParamsImage: components.ParamsImage{
					ImageSquareSize: "160",
					ImageAlt:        "Barber SHop",
					ImageSource:     "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/hero-bg-1.jpg",
				},
			},
		),

		whero.WidgetHero(
			&whero.ParamsWidgetHero{
				Title:   "Expert cuts and classic styles",
				Message: "Experience the finest grooming services in town. from traditional haircuts to modern styling, we've got you covered.",

				ButtonPrimaryInfo: components.ParamsARef{
					CSSClass: "button primary",
					Caption:  "Book appointment",
				},
				ButtonSecondaryInfo: components.ParamsARef{
					CSSClass: "button secondary",
					Caption:  "View services",
				},

				ParamsImage: components.ParamsImage{
					ImageSquareSize: "400",
					ImageAlt:        "Barber Shop",
					ImageSource:     "https://themes.getmotopress.com/bro-barbershop/wp-content/uploads/sites/64/2024/01/appointment-img.jpg",
				},
			},
		),

		wcard.WidgetCards(
			&wcard.ParamsWidgetCards{
				Title: "Popular Services",

				CurrencySimbol: "RON",
				PriceCaption:   "Pret",

				CSSFlexGap: "20px",

				Cards: []*wcard.WidgetCardVerticalInfo{
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
		),
	)
}

func TestPage(t *testing.T) {
	writer, errWriter := helpers.GetFileWriter(t.Name() + ".html")
	require.NoError(t, errWriter)

	defer writer.Close()

	page := components.Page{
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
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("hero.css"),
			),
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("cards.css"),
			),
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("card.css"),
			),
		},

		Body: []dsl.Node{
			Page(),
		},
	}

	writer.Write(
		dsl.Render(
			page.Build(),
		),
	)
}
