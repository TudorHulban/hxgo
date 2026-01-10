package article

import (
	"fmt"
	"time"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
)

type ArticleMeta struct {
	LastUpdate time.Time
	Author     string

	NumberComments uint16
}

func (meta ArticleMeta) String() string {
	return fmt.Sprintf(
		`%s | by %s | %d Comments`,

		meta.LastUpdate.Format("January 2, 2006"),
		meta.Author,
		meta.NumberComments,
	)
}

type ParamsWidgetArticleCard struct {
	Category       string
	Title          string
	ArticleExcerpt string

	components.ParamsImage
	ArticleMeta

	components.ParamsButtonSubmit
}

func WidgetArticleCard(params *ParamsWidgetArticleCard) dsl.Node {
	return dsl.Article(
		dsl.AttrClass("article-card"),

		dsl.Div(
			dsl.AttrClass("article-card-image-container"),

			components.Image(
				&params.ParamsImage,
			),
		),

		dsl.Div(
			dsl.AttrClass("article-card-content"),

			dsl.P(
				dsl.AttrClass("article-card-category"),
				dsl.Text(
					params.Category,
				),
			),

			dsl.H2(
				dsl.AttrClass("article-card-title"),

				dsl.Text(
					params.Title,
				),
			),

			dsl.P(
				dsl.AttrClass("article-card-meta"),

				dsl.Text(
					params.ArticleMeta.String(),
				),
			),

			dsl.P(
				dsl.AttrClass("article-card-excerpt"),

				dsl.Text(
					params.ArticleExcerpt,
				),
			),

			dsl.Div(
				dsl.AttrClass("article-card-action"),

				components.ButtonSubmit(
					&params.ParamsButtonSubmit,
				),
			),
		),
	)
}

func CSSArticleCard() {}
