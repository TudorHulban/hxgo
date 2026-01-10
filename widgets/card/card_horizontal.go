package card

import (
	"fmt"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/widgets/base"
)

type WidgetCardHorizontalInfo struct {
	Title string
	Price string

	ImageSquareSize string
	ImageSource     string

	Text       string
	Highlights []*base.Highlight
}

type ParamsWidgetCardHorizontal struct {
	CurrencySimbol string
	PriceCaption   string

	CSSFlexGap string

	WidgetCardHorizontalInfo
}

func WidgetCardHorizontal(params *ParamsWidgetCardHorizontal) dsl.Node {
	return dsl.Div(
		dsl.AttrCSS(
			helpers.Sprintf(
				"display:flex;flex-direction:row;gap:%s;",
				params.CSSFlexGap,
			),
		),

		dsl.AttrClass("has-post-thumbnail"),

		components.Image(
			&components.ParamsImage{
				ImageSquareSize: params.ImageSquareSize,
				ImageSource:     params.ImageSource,
				ImageAlt:        params.Title,
			},
		),

		dsl.Div(
			append(
				[]dsl.Node{
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

					dsl.Div(
						dsl.AttrClass("content"),

						dsl.Span(
							dsl.Text(params.Text),
						),
					),
				},

				helpers.ForEachValueWAddition(
					&helpers.ParamsForEachValueWAddition[*base.Highlight, dsl.Node]{
						Values: params.Highlights,

						Addition: func() dsl.Node {
							return dsl.Div(
								dsl.AttrClass(
									"horizontal-line",
								),
							)
						},

						Process: func(item *base.Highlight) dsl.Node {
							return dsl.Div(
								dsl.Text(
									fmt.Sprintf(
										"%s : %s",
										item.Label,
										item.Text,
									),
								),

								dsl.Div(
									dsl.AttrClass(
										"horizontal-line",
									),
								),
							)
						},
					},
				)...,
			)...,
		),
	)
}
