package winputdate

import (
	"time"

	pagecss "github.com/TudorHulban/hx-core/page-css"
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
)

type ParamsWidgetInputDate struct {
	CSSID string `valid:"required"`

	DateValue   time.Time
	HowManyDays uint8
}

type ResponseWidgetInputDate struct {
	HTML           dsl.Node
	LinkJavascript dsl.Node
}

func WidgetInputDate(params *ParamsWidgetInputDate) *ResponseWidgetInputDate {
	return &ResponseWidgetInputDate{
		HTML: dsl.Div(
			dsl.AttrClass("input-date"),

			dsl.Raw(
				helpers.Sprintf(
					`<input class="flatpickr-input" type="text" id="%s" value="%s" min="%s" max="%s"/>`,

					params.CSSID,
					params.DateValue.Format("2006-01-02"),
					params.DateValue.AddDate(0, 0, 1).Format("2006-01-02"),
					params.DateValue.AddDate(0, 0, int(params.HowManyDays)).Format("2006-01-02"),
				),
			),

			dsl.Raw(
				helpers.Sprintf(
					`<script>
					document.addEventListener('DOMContentLoaded', function() {
					const inputElement = document.getElementById('%s');
					if (inputElement) {
							flatpickr(inputElement, {
								dateFormat: "Y-m-d",
								inline: true,
								defaultDate: inputElement.value,
								minDate: inputElement.getAttribute('min'),
								maxDate: inputElement.getAttribute('max')
							});
						}
					});
					</script>`,

					params.CSSID,
				),
			),
		),

		LinkJavascript: dsl.Raw(
			`<script src="https://cdn.jsdelivr.net/npm/flatpickr"></script>`,
		),
	}
}

func CSSInputDate() *pagecss.CSSElement {
	return &pagecss.CSSElement{}
}
