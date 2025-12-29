package components

import (
	"github.com/TudorHulban/hxgo/dsl"
)

type ParamsH struct {
	CSSID string
	Text  string
}

func H1Raw(params *ParamsH) string {
	return `<h1 id="` + params.CSSID + `">` + params.Text + `</h1>`
}

func H2Raw(params *ParamsH) string {
	return `<h2 id="` + params.CSSID + `">` + params.Text + `</h2>`
}

func H3Raw(params *ParamsH) string {
	return `<h3 id="` + params.CSSID + `">` + params.Text + `</h3>`
}

func H4Raw(params *ParamsH) string {
	return `<h4 id="` + params.CSSID + `">` + params.Text + `</h4>`
}

func H5Raw(params *ParamsH) string {
	return `<h5 id="` + params.CSSID + `">` + params.Text + `</h5>`
}

func H6Raw(params *ParamsH) string {
	return `<h6 id="` + params.CSSID + `">` + params.Text + `</h6>`
}

func H3(params *ParamsH) dsl.Node {
	return dsl.H3(
		dsl.AttrIDLength(
			params.CSSID,
		),

		dsl.Text(
			params.Text,
		),
	)
}
