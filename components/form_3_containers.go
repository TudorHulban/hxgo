package components

import "github.com/TudorHulban/hxgo/dsl"

type ParamsNewFormThreeContainers struct {
	SymbolEntry   string // form logo symbol
	TextForm      string
	CSSIDTextForm string

	IDEnclosingDiv    string
	IDForm            string
	IDContainersDiv   string // link media queries to this ID.
	IDContainerLeft   string
	IDContainerMiddle string
	IDContainerRight  string

	ElementsInputLeft   []dsl.Node
	ElementsInputMiddle []dsl.Node
	ElementsInputRight  []dsl.Node

	Buttons []dsl.Node
}

func NewFormThreeContainers(params *ParamsNewFormThreeContainers) dsl.Node {
	return dsl.Div(
		dsl.AttrIDLength(params.IDEnclosingDiv),

		dsl.If(
			len(params.SymbolEntry) > 0 || len(params.TextForm) > 0,

			dsl.Div(
				dsl.AttrCSS(
					"display: flex;",
				),

				dsl.If(
					len(params.SymbolEntry) > 0,

					dsl.Span(
						dsl.AttrClass("material-symbols-outlined"),
						dsl.Text(
							params.SymbolEntry,
						),
					),
				),

				dsl.If(
					len(params.TextForm) > 0,

					H3(
						&ParamsH{
							Text:  params.TextForm,
							CSSID: params.CSSIDTextForm,
						},
					),
				),
			),
		),

		dsl.FormWithID(
			params.IDForm,

			append(
				[]dsl.Node{
					dsl.AutocompleteOff(),

					dsl.Div(
						[]dsl.Node{
							dsl.AttrID(params.IDContainersDiv),

							dsl.ContainerDiv(
								&dsl.ParamsContainerDiv{
									IDContainer: params.IDContainerLeft,
									Elements:    params.ElementsInputLeft,
								},
							),

							dsl.ContainerDiv(
								&dsl.ParamsContainerDiv{
									IDContainer: params.IDContainerMiddle,
									Elements:    params.ElementsInputMiddle,
								},
							),

							dsl.ContainerDiv(
								&dsl.ParamsContainerDiv{
									IDContainer: params.IDContainerRight,
									Elements:    params.ElementsInputRight,
								},
							),
						}...,
					),
				},

				dsl.If(
					len(params.Buttons) > 0,

					dsl.Div(
						append(
							[]dsl.Node{
								dsl.AttrCSS(
									`display:flex;width:100%;flex-wrap:nowrap;gap:10px;`,
								),
							},
							params.Buttons...,
						)...,
					),
				),
			)...,
		),
	)
}
