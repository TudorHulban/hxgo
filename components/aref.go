package components

import (
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/hx"
)

type ParamsARef struct {
	CSSClass string

	Route   string
	ItemID  string
	Caption string

	TargetsMultiswap string
	TargetsSend      string
}

func ARefRaw(params *ParamsARef) string {
	if len(params.TargetsMultiswap) == 0 {
		return helpers.Sprintf(
			`<a href="%s">%s</a>`,

			params.Route,
			params.Caption,
		)
	}

	return helpers.Sprintf(
		`<a%s href="#" %s="%s" %s="%s"%s>%s</a>`,

		helpers.Ternary(
			len(params.CSSClass) > 0,

			helpers.Sprintf(
				` class="%s" `,
				params.CSSClass,
			),
			"",
		),

		hx.HXPOST,

		helpers.Ternary(
			len(params.ItemID) > 0,

			params.Route+"/"+params.ItemID,
			params.Route,
		),

		hx.HXSwap,
		params.TargetsMultiswap,

		helpers.Ternary(
			len(params.TargetsSend) > 0,

			helpers.Sprintf(
				` %s="%s"`,
				hx.HXSend,
				params.TargetsSend,
			),
			"",
		),

		params.Caption,
	)
}
