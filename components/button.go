package components

import (
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
)

type ParamsButtonSubmit struct {
	Label    string
	CSSClass string
	CSSID    string

	HXActionType     string
	HXActionEndpoint string

	HXRedirectTo      string
	HXSwapElements    []string // includes CSS ID sanitization
	HXRequireElements []string // includes CSS ID sanitization
	HXSendElements    []string // includes CSS ID sanitization

	HXEnableElements  []string // includes CSS ID sanitization
	HXDisableElements []string // includes CSS ID sanitization

	IsHXUpload bool
	IsDisabled bool
	IsNewTab   bool
}

func ButtonSubmit(params *ParamsButtonSubmit) dsl.Node {
	return dsl.Button(
		dsl.AttrType("submit"),
		dsl.AttrIDLength(params.CSSID),
		dsl.AttrClassLength(params.CSSClass),

		dsl.Text(
			params.Label,
		),

		dsl.If(
			len(params.HXActionType) == 0 || params.IsDisabled,

			dsl.Attr(
				"disabled",
			),
		),

		dsl.If(
			len(params.HXActionType) > 0,

			dsl.AttrWithValue(
				params.HXActionType,
				params.HXActionEndpoint,
			),
		),

		dsl.AttrHXRedirectLength(params.HXRedirectTo),
		dsl.AttrHXSwap(params.HXSwapElements...),
		dsl.AttrHXRequire(params.HXRequireElements...),
		dsl.AttrHXSend(params.HXSendElements...),
		dsl.AttrHXEnable(params.HXEnableElements...),
		dsl.AttrHXDisable(params.HXDisableElements...),
		dsl.AttrHXUpload(params.IsHXUpload),

		dsl.If(
			params.IsNewTab,

			dsl.OnClick(
				helpers.Sprintf(
					`window.open('%s','_blank')`,
					params.HXActionEndpoint,
				),
			),
		),
	)
}

func ButtonSubmitWCSS(params *ParamsButtonSubmit, css []dsl.Node) dsl.Node {
	return dsl.Button(
		append(
			css,

			dsl.AttrType("submit"),
			dsl.AttrIDLength(params.CSSID),
			dsl.AttrClassLength(params.CSSClass),

			dsl.Text(
				params.Label,
			),

			dsl.If(
				len(params.HXActionType) == 0 || params.IsDisabled,

				dsl.Attr(
					"disabled",
				),
			),

			dsl.If(
				len(params.HXActionType) > 0,

				dsl.AttrWithValue(
					params.HXActionType,
					params.HXActionEndpoint,
				),
			),

			dsl.AttrHXRedirectLength(params.HXRedirectTo),
			dsl.AttrHXSwap(params.HXSwapElements...),
			dsl.AttrHXRequire(params.HXRequireElements...),
			dsl.AttrHXSend(params.HXSendElements...),
			dsl.AttrHXEnable(params.HXEnableElements...),
			dsl.AttrHXDisable(params.HXDisableElements...),
			dsl.AttrHXUpload(params.IsHXUpload),

			dsl.If(
				params.IsNewTab,

				dsl.OnClick(
					helpers.Sprintf(
						`window.open('%s','_blank')`,
						params.HXActionEndpoint,
					),
				),
			),
		)...,
	)
}
