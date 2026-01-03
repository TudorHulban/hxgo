package components

import (
	"database/sql"
	"strings"

	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/request"
	"github.com/TudorHulban/hxgo/timeutil"
)

type InputDateTime struct {
	CSSDivClass string
	CSSInputID  string

	ElementName      string
	LabelElementName string

	AfterTime sql.NullTime

	IsDisabled bool
}

func (el InputDateTime) Raw() dsl.Node {
	result := [4]string{
		helpers.Ternary(
			len(el.CSSDivClass) == 0,

			`<div>`,
			helpers.Sprintf(
				`<div class="%s">`,
				el.CSSDivClass,
			),
		),

		helpers.Sprintf(
			`<label for="%s">%s:</label>`,

			el.ElementName,
			el.LabelElementName,
		),

		helpers.Sprintf(
			`<input type="datetime-local" id="%s" name="%s" %s`,

			el.CSSInputID,
			helpers.Coalesce(
				el.ElementName,
				el.LabelElementName,
			),

			helpers.Ternary(
				el.AfterTime.Valid,

				helpers.Sprintf(
					`value="%s"`,
					request.TimeForInputDateTime(
						el.AfterTime.Time,
						timeutil.LayoutDateTime,
					),
				),
				"",
			),
		),

		helpers.Ternary(
			el.IsDisabled,

			"disabled></div>",
			"></div>",
		),
	}

	return dsl.Raw(
		strings.Join(result[:], "\n"),
	)
}
