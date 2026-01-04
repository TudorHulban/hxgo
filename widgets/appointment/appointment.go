package wappointment

import (
	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	pagecss "github.com/TudorHulban/hx-core/page-css"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	winputdate "github.com/TudorHulban/hx-widgets/input-date"
	winputslots "github.com/TudorHulban/hx-widgets/input-slots"
)

type ParamsWidgetAppointment struct {
	winputslots.ParamsWidgetSlots
	winputdate.ParamsWidgetInputDate

	SelectLabel  string
	SelectValues []string

	hxcomponents.ParamsButtonSubmit
}

type ResponseWidgetAppointment struct {
	HTML           hxprimitives.Node
	LinkJavascript hxprimitives.Node

	CSS []func() *pagecss.CSSElement
}

func WidgetAppointment(params *ParamsWidgetAppointment) *ResponseWidgetAppointment {
	nodesInputDate := winputdate.WidgetInputDate(
		&params.ParamsWidgetInputDate,
	)

	inputSimple := hxcomponents.InputSelect{
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

		HTML: hxhtml.Div(
			hxprimitives.AttrClass(
				"appointment-container",
			),

			hxhtml.Div(
				hxprimitives.AttrCSS(
					`display: flex; flex-wrap: nowrap; gap: 0.2em;`,
				),

				nodesInputDate.HTML,

				hxhtml.Div(
					hxprimitives.AttrCSS(
						`display: flex; flex-direction: column; gap: 0.1em;`,
					),

					inputSimple.Raw(),

					winputslots.WidgetSlots(
						&params.ParamsWidgetSlots,
					),
				),
			),

			hxcomponents.ButtonSubmit(
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
