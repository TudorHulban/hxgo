package dsl

import (
	"strings"

	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/hx"
)

func AttrHXRedirect(value string) attribute {
	return AttrWithValue(
		hx.HXRedirect,
		value,
	)
}

func AttrHXRedirectLength(value string) attribute {
	if len(value) > 0 {
		return AttrWithValue(
			hx.HXRedirect,
			value,
		)
	}

	return attribute{}
}

func AttrHXRequire(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hx.HXRequire,
		strings.Join(
			helpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXSwap(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hx.HXSwap,
		strings.Join(
			helpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXSend(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hx.HXSend,
		strings.Join(
			helpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXEnable(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hx.HXDirectEnable,
		strings.Join(
			helpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXDisable(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hx.HXDirectDisable,
		strings.Join(
			helpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXUpload(withUpload bool) attribute {
	if withUpload {
		return Attr(
			hx.HXUpload,
		)
	}

	return attribute{}
}

func AttrHXPOST(value string) attribute {
	return AttrWithValue(
		hx.HXPOST,
		value,
	)
}

func AttrHXShow(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hx.HXShow,
		strings.Join(
			values,
			",",
		),
	)
}

func AttrHXHide(values ...string) attribute {
	return AttrWithValueIf(
		len(values) > 0,

		hx.HXHide,
		strings.Join(
			values,
			",",
		),
	)
}
