package components

import (
	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
)

type ParamsImage struct {
	ImageSquareSize string
	ImageSource     string
	ImageAlt        string
}

func Img(params *ParamsImage) dsl.Node {
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
		dsl.AttrWithValue("decoding", "async"),
		dsl.AttrWithValue("height", params.ImageSquareSize),
		dsl.AttrWithValue("width", params.ImageSquareSize),
		dsl.AttrWithValue("alt", params.ImageAlt),
		dsl.AttrWithValue("loading", "lazy"),
		dsl.AttrWithValue("src", params.ImageSource),
		dsl.AttrClass("attachment-post-thumbnail"),
	)
}
