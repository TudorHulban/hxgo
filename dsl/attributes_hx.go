package dsl

import (
	"strings"

	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/hx"
)

func AttrHXRedirect(value string) Node {
	return AttrWithValue(
		hx.HXRedirect,
		value,
	)
}

func AttrHXRedirectLength(value string) Node {
	if len(value) > 0 {
		return AttrWithValue(
			hx.HXRedirect,
			value,
		)
	}

	return Noop
}

func AttrHXRequire(values ...string) Node {
	return AttrIf(
		len(values) > 0,

		hx.HXRequire,
		strings.Join(
			helpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXSwap(values ...string) Node {
	return AttrIf(
		len(values) > 0,

		hx.HXSwap,
		strings.Join(
			helpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXSend(values ...string) Node {
	return AttrIf(
		len(values) > 0,

		hx.HXSend,
		strings.Join(
			helpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXEnable(values ...string) Node {
	return AttrIf(
		len(values) > 0,

		hx.HXDirectEnable,
		strings.Join(
			helpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXDisable(values ...string) Node {
	return AttrIf(
		len(values) > 0,

		hx.HXDirectDisable,
		strings.Join(
			helpers.SanitizeCSSIdsSlice(values...),
			",",
		),
	)
}

func AttrHXUpload(withUpload bool) Node {
	if withUpload {
		return Attr(
			hx.HXUpload,
		)
	}

	return Noop
}

func AttrHXPOST(value string) Node {
	return AttrWithValue(
		hx.HXPOST,
		value,
	)
}

func AttrHXShow(values ...string) Node {
	return AttrIf(
		len(values) > 0,

		hx.HXShow,
		strings.Join(
			values,
			",",
		),
	)
}

func AttrHXHide(values ...string) Node {
	return AttrIf(
		len(values) > 0,

		hx.HXHide,
		strings.Join(
			values,
			",",
		),
	)
}
