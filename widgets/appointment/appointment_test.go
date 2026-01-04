package wappointment

import (
	"fmt"
	"testing"
	"time"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	pagecss "github.com/TudorHulban/hx-core/page-css"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	"github.com/TudorHulban/hx-widgets/base"
	"github.com/TudorHulban/hx-widgets/helpers"
	winputdate "github.com/TudorHulban/hx-widgets/input-date"
	winputslots "github.com/TudorHulban/hx-widgets/input-slots"
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
			ParamsButtonSubmit: hxcomponents.ParamsButtonSubmit{
				Label:    "Submit",
				CSSClass: "btn-submit",
				CSSID:    winputslots.ButtonSubmitCSSID,
			},
		},
	)

	page := hxcomponents.Page{
		Title: t.Name(),

		Head: []hxprimitives.Node{
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("generated.css"),
			),
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("https://cdn.jsdelivr.net/npm/flatpickr/dist/flatpickr.min.css"),
			),
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("https://npmcdn.com/flatpickr/dist/themes/dark.css"),
			),
		},

		Body: []hxprimitives.Node{
			fragment.LinkJavascript,
			fragment.HTML,
		},
	}

	writerCSS, errWriterCSS := helpers.GetFileWriter("generated.css")
	require.NoError(t, errWriterCSS)

	defer writerCSS.Close()

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
		writerCSS,
		&pagecss.ParamsSpaces{
			NumberSpaces:              5,
			IncrementWithNumberSpaces: 2,
		},
	)

	writerHTML, errWriterHTML := helpers.GetFileWriter(t.Name() + ".html")
	require.NoError(t, errWriterHTML)

	defer writerHTML.Close()

	fmt.Println(
		page.Build().Render(writerHTML),
	)
}
