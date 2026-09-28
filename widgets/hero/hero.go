package hero

import (
	"github.com/TudorHulban/hxgo/components/primitives"
	"github.com/TudorHulban/hxgo/dsl"
)

type ParamsWidgetHero struct {
	Title   string
	Message string

	ButtonPrimaryInfo   primitives.ParamsARef
	ButtonSecondaryInfo primitives.ParamsARef

	primitives.ParamsImage
}

func WidgetHero(params *ParamsWidgetHero) dsl.Node {
	return dsl.Div(
		dsl.AttrClass(
			"hero",
		),

		dsl.Div(
			dsl.AttrClass(
				"hero-content",
			),

			dsl.H1(
				dsl.Text(
					params.Title,
				),
			),

			dsl.P(
				dsl.Text(
					params.Message,
				),
			),

			dsl.Div(
				dsl.AttrClass(
					"hero-buttons",
				),

				dsl.Raw(
					primitives.ARefRaw(
						&params.ButtonPrimaryInfo,
					),
				),

				dsl.Raw(
					primitives.ARefRaw(
						&params.ButtonSecondaryInfo,
					),
				),
			),
		),

		dsl.Div(
			dsl.AttrClass(
				"hero-image",
			),

			primitives.Image(
				&params.ParamsImage,
			),
		),
	)
}
