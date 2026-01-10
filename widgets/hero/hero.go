package hero

import (
	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
)

type ParamsWidgetHero struct {
	Title   string
	Message string

	ButtonPrimaryInfo   components.ParamsARef
	ButtonSecondaryInfo components.ParamsARef

	components.ParamsImage
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
					components.ARefRaw(
						&params.ButtonPrimaryInfo,
					),
				),

				dsl.Raw(
					components.ARefRaw(
						&params.ButtonSecondaryInfo,
					),
				),
			),
		),

		dsl.Div(
			dsl.AttrClass(
				"hero-image",
			),

			components.Image(
				&params.ParamsImage,
			),
		),
	)
}
