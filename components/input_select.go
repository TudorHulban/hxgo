package components

import (
	"strings"

	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/hx"
)

type InputSelect struct {
	CSSDivID    string
	CSSDivClass string

	CSSInputName string
	CSSInputID   string

	ElementName      string
	LabelElementName string // translated value

	SelectedValue string
	SelectValues  []string

	Action hx.Action

	WithEmptyOption bool
	IsDisabled      bool
}

const optionEmpty = `<option label=" "></option>`

func (el InputSelect) optionSelect(value string) string {
	return helpers.Sprintf(
		`<option value="%s"%s>%s</option>`,

		strings.ToLower(value),
		helpers.Ternary(
			value == el.SelectedValue,

			" selected",
			"",
		),
		value,
	)
}

func (el InputSelect) generateOptions() string {
	options := helpers.ForEachValue(
		el.SelectValues,
		el.optionSelect,
	)

	if el.WithEmptyOption {
		options = append(
			[]string{
				optionEmpty,
			},
			options...,
		)
	}

	return strings.Join(options, "\n")
}

// attributes should be surrounded by space / " ".
func (el InputSelect) selectAttributes() string {
	var sb strings.Builder

	sb.WriteString(
		helpers.Ternary(
			len(el.CSSInputID) > 0,

			`id="`+el.CSSInputID+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		helpers.Ternary(
			len(el.CSSInputName) > 0,

			`name="`+strings.ToLower(el.CSSInputName)+`"`+" ",
			`name="`+strings.ToLower(helpers.Coalesce(el.ElementName, el.LabelElementName))+`"`+" ",
		),
	)

	sb.WriteString(
		helpers.Ternary(
			len(el.Action.Method) > 0,

			el.Action.Method+`="`+el.Action.Endpoint+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		helpers.Ternary(
			len(el.Action.Swaps) > 0,

			hx.HXSwap+`="`+helpers.SanitizeCSSIds(el.Action.Swaps)+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		helpers.Ternary(
			len(el.Action.Sends) > 0,

			hx.HXSend+`="`+helpers.SanitizeCSSIds(el.Action.Sends)+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		helpers.Ternary(
			len(el.Action.Require) > 0,

			hx.HXRequire+`="`+helpers.SanitizeCSSIds(el.Action.Require)+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		helpers.Ternary(
			len(el.Action.OnChangeEnable) > 0,

			hx.HXOnChangeEnable+`="`+helpers.SanitizeCSSIds(el.Action.OnChangeEnable)+`"`+" ",
			"",
		),
	)

	sb.WriteString(
		helpers.Ternary(
			el.IsDisabled,

			"disabled"+" ",
			"",
		),
	)

	return sb.String()
}

func (el InputSelect) tagLabel() string {
	return helpers.Ternary(
		len(el.CSSInputID) > 0,

		helpers.Sprintf(
			`<label for="%s">%s:</label>`,

			el.CSSInputID,
			el.LabelElementName,
		),
		helpers.Sprintf(
			`<label>%s:</label>`,

			el.LabelElementName,
		),
	)
}

func (el InputSelect) RawSelect() dsl.Node {
	return dsl.Raw(
		helpers.Sprintf(
			`<select %s>%s</select>`,

			el.selectAttributes(),
			el.generateOptions(),
		),
	)
}

func (el InputSelect) Raw() dsl.Node {
	return dsl.Raw(
		helpers.Sprintf(
			`<div%s%s>%s</div>`,

			helpers.Ternary(
				len(el.CSSDivID) > 0,

				` id="`+el.CSSDivID+`"`,
				"",
			),

			helpers.Ternary(
				len(el.CSSDivClass) > 0,

				` class="`+el.CSSDivClass+`"`,
				"",
			),

			strings.Join(
				[]string{
					el.tagLabel(),
					helpers.Sprintf(
						`<select %s>%s</select>`,

						el.selectAttributes(),
						el.generateOptions(),
					),
				},
				"\n",
			),
		),
	)
}
