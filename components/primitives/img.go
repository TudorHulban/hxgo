package primitives

import (
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
)

type ParamsImage struct {
	ImageSquareSize string
	ImageSource     string
	ImageAlt        string
}

func Image(params *ParamsImage) dsl.Node {
	return dsl.Img(
		dsl.AttrWithValue(
			"sizes",

			helpers.Sprintf(
				"auto, (max-width: %s) 100vw, %s",

				params.ImageSquareSize,
				params.ImageSquareSize,
			),
		),
		dsl.AttrCSS("object-fit:cover;"),
		dsl.AttrClass("attachment-post-thumbnail"),
		dsl.AttrWithValue("decoding", "async"),
		dsl.AttrWithValue("height", params.ImageSquareSize),
		dsl.AttrWithValue("width", params.ImageSquareSize),

		dsl.AttrImgAlternativeText(params.ImageAlt),
		dsl.AttrImgSource(params.ImageSource),

		dsl.AttrWithValue("loading", "lazy"),
	)
}
