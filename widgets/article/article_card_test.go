package warticle

import (
	"testing"
	"time"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/hx"
	"github.com/stretchr/testify/require"
)

func TestWidgetArticleCard(t *testing.T) {
	widget := WidgetArticleCard(
		&ParamsWidgetArticleCard{
			Category:       "category-g1",
			Title:          "Article-title",
			ArticleExcerpt: "lorem ipsum dolor sit amet, consectetur adipiscing elit. sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",

			ArticleMeta: ArticleMeta{
				LastUpdate:     time.Now(),
				Author:         "John Smith",
				NumberComments: 1,
			},

			ParamsImage: components.ParamsImage{
				ImageSquareSize: "260",
				ImageSource:     "https://images.pexels.com/photos/539694/pexels-photo-539694.jpeg",
				ImageAlt:        "1",
			},

			ParamsButtonSubmit: components.ParamsButtonSubmit{
				Label: "Continue reading",

				HXActionType: hx.HXPOST,
			},
		},
	)

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
				dsl.Href("article_card.css"),
			),
		},

		Body: []dsl.Node{
			widget,
		},
	}

	writer.Write(
		dsl.Render(
			page.Build(),
		),
	)
}
