package components

import (
	"github.com/TudorHulban/hxgo/dsl"
)

type ParamsRenderFormContainer struct {
	IDContainer string
	Elements    []dsl.Node
}

func RenderFormContainer(params *ParamsRenderFormContainer) dsl.Node {
	return dsl.Div(
		append(
			params.Elements,

			dsl.If(
				len(params.IDContainer) > 0,

				dsl.AttrID(params.IDContainer),
			),
		)...,
	)
}
