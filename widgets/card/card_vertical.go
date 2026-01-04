package wcard

import (
	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
)

type WidgetCardVerticalInfo struct {
	Title string
	Price string

	ImageSquareSize string
	ImageSource     string

	ActionEndpoint string
}

type ParamsWidgetCardVertical struct {
	WidgetCardVerticalInfo

	CurrencySimbol string
	PriceCaption   string
}

func WidgetCardVertical(params *ParamsWidgetCardVertical) dsl.Node {
	return dsl.Div(
		dsl.AttrClass("has-post-thumbnail bw-to-color"),

		dsl.Div(
			dsl.AttrClass("post-wrapper"),

			dsl.P(
				dsl.AttrClass("post-thumbnail"),

				dsl.A(
					dsl.Href(params.ActionEndpoint),

					components.Image(
						&components.ParamsImage{
							ImageSquareSize: params.ImageSquareSize,
							ImageSource:     params.ImageSource,
							ImageAlt:        params.Title,
						},
					),
				),
			),

			dsl.Div(
				dsl.AttrClass("post-title-wrapper"),

				dsl.H2(
					dsl.AttrClass("entry-title post-title"),

					dsl.A(
						dsl.Href(params.ActionEndpoint),
						dsl.Text(params.Title),
					),
				),

				dsl.P(
					dsl.AttrClass("service-price"),

					dsl.Span(
						dsl.AttrClass("price-title"),
						dsl.Text(params.PriceCaption+":"),
					),

					dsl.Span(
						dsl.AttrClass("price"),

						dsl.Span(
							dsl.AttrClass("currency"),
							dsl.Text(params.CurrencySimbol),
						),

						dsl.Text(params.Price),
					),
				),
			),
		),
	)
}
