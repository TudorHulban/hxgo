package card

import (
	"testing"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/widgets/base"
	"github.com/stretchr/testify/require"
)

func TestHorizontalCard(t *testing.T) {
	fragment := WidgetCardHorizontal(
		&ParamsWidgetCardHorizontal{
			WidgetCardHorizontalInfo: WidgetCardHorizontalInfo{
				Title: "Washing Head",
				Price: "40",

				ImageSquareSize: "220",
				ImageSource:     "https://images.pexels.com/photos/668353/pexels-photo-668353.jpeg",

				Text: "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam",
				Highlights: []*base.Highlight{
					{
						Label: "Time",
						Text:  "30 min",
					},
					{
						Label: "Equipment",
						Text:  "Shampoo",
					},
				},
			},

			CurrencySimbol: "RON",
			PriceCaption:   "Pret",

			CSSFlexGap: "20px",
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
