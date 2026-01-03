package components

import (
	"strconv"
	"strings"

	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/hx"
)

type InputSimple struct {
	CSSDivID    string
	CSSDivClass string

	CSSInputName string
	CSSInputID   string

	ElementName      string
	LabelElementName string

	TypeInput string // see https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#input_types
	TextInput string

	Action hx.Action

	IsDisabled bool
}

func (el InputSimple) tagInput(toLowerElementName string) string {
	attributes := []string{
		`type="` + el.TypeInput + `"`,

		helpers.Ternary(
			len(el.CSSInputID) > 0,

			`id="`+el.CSSInputID+`"`,
			"",
		),

		`name="` + toLowerElementName + `"`,
	}

	if len(el.TextInput) > 0 && el.TextInput != "0" {
		attributes = append(
			attributes,

			`value="`+el.TextInput+`"`,
		)
	} else {
		if el.Action.LengthMax > 0 {
			buf := make([]byte, 0, 5)
			buf = strconv.AppendUint(buf, uint64(el.Action.LengthMax), 10)

			attributes = append(
				attributes,
				helpers.Sprintf(
					`%s=%s`,

					hx.HXMax,
					string(buf),
				),
			)
		}

		if el.Action.LengthMin > 0 {
			buf := make([]byte, 0, 5)
			buf = strconv.AppendUint(buf, uint64(el.Action.LengthMin), 10)

			attributes = append(
				attributes,

				helpers.Sprintf(
					`%s=%s`,

					hx.HXMin,
					string(buf),
				),
			)
		}

		if el.Action.HavePasswordsValues() {
			attributes = append(
				attributes,

				hx.HXValidationPasswords+`="`+
					helpers.SanitizeCSSId(el.Action.CSSIDValidationPasswords[0])+
					`,`+
					helpers.SanitizeCSSId(el.Action.CSSIDValidationPasswords[1])+
					`"`,
			)
		}

		if len(el.Action.CSSIDValidationDisable) > 0 {
			attributes = append(
				attributes,

				hx.HXValidationDisable+`="`+el.Action.CSSIDValidationDisable+`"`,
			)
		}
	}

	if el.IsDisabled {
		attributes = append(
			attributes,

			"disabled",
		)
	}

	return helpers.Sprintf(
		`<input %s>`,

		strings.Join(attributes, " "),
	)
}

func (el InputSimple) wrapDiv(content string) string {
	return helpers.Sprintf(
		`<div %s%s>%s</div>`,

		helpers.Ternary(
			len(el.CSSDivID) > 0,

			`id="`+el.CSSDivID+`"`,
			"",
		),
		helpers.Ternary(
			len(el.CSSDivClass) > 0,

			`class="`+el.CSSDivClass+`"`,
			"",
		),
		content,
	)
}

func (el InputSimple) RawSelect() dsl.Node {
	return dsl.Raw(
		el.tagInput(
			strings.ToLower(
				helpers.Coalesce(
					el.ElementName,
					el.LabelElementName,
				),
			),
		),
	)
}

func (el InputSimple) Raw() dsl.Node {
	return dsl.Raw(
		el.wrapDiv(
			strings.Join(
				[]string{
					helpers.Ternary(
						len(el.CSSInputID) > 0,

						`<label for="`+el.CSSInputID+`">`+el.LabelElementName+`:</label>`,
						`<label>`+el.LabelElementName+`:</label>`,
					),

					el.tagInput(
						strings.ToLower(
							helpers.Coalesce(
								el.ElementName,
								el.LabelElementName,
							),
						),
					),
				},
				"",
			),
		),
	)
}
