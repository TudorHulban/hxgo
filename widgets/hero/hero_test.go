package whero

import (
	"testing"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/stretchr/testify/require"
)

func TestHero(t *testing.T) {
	fragment := WidgetHero(
		&ParamsWidgetHero{
			Title:   "Expert cuts and classic styles",
			Message: "experience the finest grooming services in town. from traditional haircuts to modern styling, we've got you covered.",

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
				dsl.Href("hero.css"),
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
