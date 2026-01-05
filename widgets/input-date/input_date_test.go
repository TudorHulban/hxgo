package winputdate

import (
	"testing"
	"time"

	"github.com/TudorHulban/hxgo/components"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/stretchr/testify/require"
)

func TestWidgetInputDate(t *testing.T) {
	fragment := WidgetInputDate(
		&ParamsWidgetInputDate{
			CSSID: "schedule",

			DateValue:   time.Now(),
			HowManyDays: 3,
		},
	)

	writer, errWriter := helpers.GetFileWriter(t.Name() + ".html")
	require.NoError(t, errWriter)

	defer writer.Close()

	page := components.Page{
		Title: t.Name(),

		Head: []dsl.Node{
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("css_base.css"),
			),
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("css_site.css"),
			),
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("https://cdn.jsdelivr.net/npm/flatpickr/dist/flatpickr.min.css"),
			),
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("https://npmcdn.com/flatpickr/dist/themes/dark.css"),
			),
			dsl.Link(
				dsl.Rel("stylesheet"),
				dsl.Href("input_date.css"),
			),
		},

		Body: []dsl.Node{
			fragment.LinkJavascript,
			fragment.HTML,
		},
	}

	writer.Write(
		dsl.Render(
			page.Build(),
		),
	)
}
