package wappointment

import (
	pagecss "github.com/TudorHulban/hx-core/page-css"
	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	winputdate "github.com/TudorHulban/hxgo/widgets/input-date"
	winputslots "github.com/TudorHulban/hxgo/widgets/input-slots"
)

type ParamsWidgetAppointment struct {
	winputslots.ParamsWidgetSlots
	winputdate.ParamsWidgetInputDate

	SelectLabel  string
	SelectValues []string

	components.ParamsButtonSubmit
}

type ResponseWidgetAppointment struct {
	HTML           dsl.Node
	LinkJavascript dsl.Node

	CSS []func() *pagecss.CSSElement
}

func WidgetAppointment(params *ParamsWidgetAppointment) *ResponseWidgetAppointment {
	nodesInputDate := winputdate.WidgetInputDate(
		&params.ParamsWidgetInputDate,
	)

	inputSimple := components.InputSelect{
		CSSDivID: "resource-selection",

		LabelElementName: params.SelectLabel,
		SelectValues:     params.SelectValues,

		WithEmptyOption: true,
	}

	return &ResponseWidgetAppointment{
		LinkJavascript: nodesInputDate.LinkJavascript,
		CSS: []func() *pagecss.CSSElement{
			winputdate.CSSInputDate,
			winputslots.CSSWidgetSlots,
			CSSAppointment,
		},

		HTML: dsl.Div(
			dsl.AttrClass(
				"appointment-container",
			),

			dsl.Div(
				dsl.AttrCSS(
					`display: flex; flex-wrap: nowrap; gap: 0.2em;`,
				),

				nodesInputDate.HTML,

				dsl.Div(
					dsl.AttrCSS(
						`display: flex; flex-direction: column; gap: 0.1em;`,
					),

					inputSimple.Raw(),

					winputslots.WidgetSlots(
						&params.ParamsWidgetSlots,
					),
				),
			),

			components.ButtonSubmit(
				&params.ParamsButtonSubmit,
			),
		),
	}
}

func CSSAppointment() *pagecss.CSSElement {
	return &pagecss.CSSElement{
		CSSAllMedias: `
		.appointment-container {
			padding: 0.3em;
			text-align: right;
			width: fit-content;
			background-color:rgb(134, 146, 138);

			#resource-selection {
				padding-top: 2.1em;
			}

			.hours-grid {
				width: 100%;
			}
		}
		`,
	}
}
