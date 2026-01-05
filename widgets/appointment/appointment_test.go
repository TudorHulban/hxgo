package wappointment

import (
	"testing"
	"time"

	pagecss "github.com/TudorHulban/hx-core/page-css"
	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/widgets/base"
	winputdate "github.com/TudorHulban/hxgo/widgets/input-date"
	winputslots "github.com/TudorHulban/hxgo/widgets/input-slots"
	"github.com/stretchr/testify/require"
)

func TestAppointment(t *testing.T) {
	fragment := WidgetAppointment(
		&ParamsWidgetAppointment{
			SelectLabel: "Doctor",
			SelectValues: []string{
				"John Smith",
				"Martha Doe",
			},

			ParamsWidgetSlots: winputslots.ParamsWidgetSlots{
				SubmitEndpoint: "xxx",
				NumberColumns:  1,

				SlotsInfo: []*winputslots.InfoSlot{
					{
						ResourceID: 1,
						SlotID:     1000,
						Caption:    "10 00 - dr. John Smith",
					},
					{
						ResourceID: 2,
						SlotID:     1030,
						Caption:    "10 30 - dr. Martha Doe",
					},
					{
						ResourceID: 1,
						SlotID:     1100,
						Caption:    "11 00 - dr. John Smith",
					},
					{
						ResourceID: 2,
						SlotID:     1100,
						Caption:    "11 00 - dr. Martha Doe",
					},
				},
			},
			ParamsWidgetInputDate: winputdate.ParamsWidgetInputDate{
				CSSID: "schedule",

				DateValue:   time.Now(),
				HowManyDays: 3,
			},
			ParamsButtonSubmit: components.ParamsButtonSubmit{
				Label:    "Submit",
				CSSClass: "btn-submit",
				CSSID:    winputslots.ButtonSubmitCSSID,
			},
		},
	)

	page := components.Page{
		Title: t.Name(),

		Head: []dsl.Node{
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("generated.css"),
			),
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("https://cdn.jsdelivr.net/npm/flatpickr/dist/flatpickr.min.css"),
			),
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("https://npmcdn.com/flatpickr/dist/themes/dark.css"),
			),
		},

		Body: []dsl.Node{
			fragment.LinkJavascript,
			fragment.HTML,
		},
	}

	writer, errWriterCSS := helpers.GetFileWriter("generated.css")
	require.NoError(t, errWriterCSS)

	defer writer.Close()

	cssPage := pagecss.NewCSSPage(
		append(
			[]func() *pagecss.CSSElement{
				base.CSSBase,
				base.CSSSite,
			},
			fragment.CSS...,
		)...,
	)

	cssPage.GetCSSAccurateBeautifiedTo(
		writer,
		&pagecss.ParamsSpaces{
			NumberSpaces:              5,
			IncrementWithNumberSpaces: 2,
		},
	)

	writerHTML, errWriterHTML := helpers.GetFileWriter(t.Name() + ".html")
	require.NoError(t, errWriterHTML)

	defer writerHTML.Close()

	writer.Write(
		dsl.Render(
			page.Build(),
		),
	)
}
