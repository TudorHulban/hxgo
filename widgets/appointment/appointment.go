package appointment

import (
	"github.com/TudorHulban/hxgo/components/buttons"
	"github.com/TudorHulban/hxgo/components/inputs"
	"github.com/TudorHulban/hxgo/dsl"
	inputdate "github.com/TudorHulban/hxgo/widgets/input-date"
	inputslots "github.com/TudorHulban/hxgo/widgets/input-slots"
)

type ParamsWidgetAppointment struct {
	inputslots.ParamsWidgetSlots
	inputdate.ParamsWidgetInputDate

	SelectLabel  string
	SelectValues []string

	buttons.ParamsButtonSubmit
}

type ResponseWidgetAppointment struct {
	CSS []dsl.CSS

	HTML           dsl.Node
	LinkJavascript dsl.Node
}

func WidgetAppointment(params *ParamsWidgetAppointment) *ResponseWidgetAppointment {
	nodesInputDate := inputdate.WidgetInputDate(
		&params.ParamsWidgetInputDate,
	)

	inputSimple := inputs.InputSelect{
		CSSDivID: "resource-selection",

		LabelElementName: params.SelectLabel,
		SelectValues:     params.SelectValues,

		WithEmptyOption: true,
	}

	return &ResponseWidgetAppointment{
		LinkJavascript: nodesInputDate.LinkJavascript,
		CSS:            []dsl.CSS{CSSAppointment},

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

					inputslots.WidgetSlots(
						&params.ParamsWidgetSlots,
					),
				),
			),

			buttons.ButtonSubmit(
				&params.ParamsButtonSubmit,
			),
		),
	}
}

func CSSAppointment() string {
	return `
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
		`
}
