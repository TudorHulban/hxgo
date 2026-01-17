package appointment

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/widgets/base"
	inputdate "github.com/TudorHulban/hxgo/widgets/input-date"
	inputslots "github.com/TudorHulban/hxgo/widgets/input-slots"
	"github.com/stretchr/testify/require"
)

func TestAppointment(t *testing.T) {
	w := WidgetAppointment(
		&ParamsWidgetAppointment{
			SelectLabel: "Doctor",
			SelectValues: []string{
				"John Smith",
				"Martha Doe",
			},

			ParamsWidgetSlots: inputslots.ParamsWidgetSlots{
				SubmitEndpoint: "xxx",
				NumberColumns:  1,

				SlotsInfo: []*inputslots.InfoSlot{
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
			ParamsWidgetInputDate: inputdate.ParamsWidgetInputDate{
				CSSID: "schedule",

				DateValue:   time.Now(),
				HowManyDays: 3,
			},
			ParamsButtonSubmit: components.ParamsButtonSubmit{
				Label:    "Submit",
				CSSClass: "btn-submit",
				CSSID:    inputslots.ButtonSubmitCSSID,
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
			w.LinkJavascript,
			w.HTML,
		},
	}

	cssContribution := dsl.CSSContribution{
		ProceduralCSS: []dsl.CSS{
			base.CSSBase,
			base.CSSSite,
		},
	}

	el := page.Build()
	el.Add(
		cssContribution.AsNode(),
	)

	writerCSS, errWriterCSS := helpers.GetFileWriter("generated.css")
	require.NoError(t, errWriterCSS)

	defer writerCSS.Close()

	writerHTML, errWriterHTML := helpers.GetFileWriter(t.Name() + ".html")
	require.NoError(t, errWriterHTML)

	defer writerHTML.Close()

	html, styles, css := dsl.RenderFull(el)
	require.Zero(t, styles)
	require.NotZero(t, html)
	require.NotZero(t, css)

	writerHTML.Write(
		html,
	)

	writerCSS.WriteString(css)

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(
				[]byte(
					dsl.HTMLwithDataCSS(
						html,
						"",
					),
				),
			)
		},
	)

	fmt.Println(
		"Open http://localhost:8080 and press Ctrl-C when done",
	)

	http.ListenAndServe(":8080", nil)
}
