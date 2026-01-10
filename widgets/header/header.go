package header

import (
	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
)

type ParamsWidgetHeader struct {
	Title             string
	BreadcrumbCaption string

	components.ParamsImage
}

func WidgetHeader(params *ParamsWidgetHeader) dsl.Node {
	return dsl.Div(
		dsl.AttrClass("page-header-info"),

		dsl.H1(
			dsl.AttrClass("entry-title"),
			dsl.Text(params.Title),
		),

		dsl.Div(
			dsl.AttrClass("header-images"),

			dsl.Img(
				dsl.AttrWithValue("src", params.ImageSource),
				dsl.AttrWithValue("alt", params.ImageAlt),
			),
		),

		dsl.Div(
			dsl.AttrClass("header-breadcrumb"),

			dsl.Span(
				dsl.Text(
					params.BreadcrumbCaption,
				),
			),
		),
	)
}
