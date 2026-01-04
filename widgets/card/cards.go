package wcard

import (
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
)

type ParamsWidgetCards struct {
	Title string

	CurrencySimbol string
	PriceCaption   string

	CSSFlexGap string

	Cards []*WidgetCardVerticalInfo
}

func WidgetCards(params *ParamsWidgetCards) dsl.Node {
	return dsl.Div(
		dsl.Div(
			dsl.AttrClass(
				"services-list-title",
			),

			dsl.H2(
				dsl.AttrClass(
					"item-title",
				),

				dsl.Text(
					params.Title,
				),
			),
		),

		dsl.Div(
			dsl.AttrClass(
				"centered-services-list",
			),

			dsl.Div(
				append(
					[]dsl.Node{
						dsl.AttrCSS(
							helpers.Sprintf(
								"display:flex;flex-direction:row;gap:%s;",

								params.CSSFlexGap,
							),
						),

						dsl.AttrClass(
							"services-list",
						),
					},

					helpers.ForEachValue(
						params.Cards,

						func(item *WidgetCardVerticalInfo) dsl.Node {
							return WidgetCardVertical(
								&ParamsWidgetCardVertical{
									WidgetCardVerticalInfo: *item,

									CurrencySimbol: params.CurrencySimbol,
									PriceCaption:   params.PriceCaption,
								},
							)
						},
					)...,
				)...,
			),
		),
	)
}
