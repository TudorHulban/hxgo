package dsl

type ParamsContainerDiv struct {
	IDContainer string
	Elements    []Node
}

func ContainerDiv(params *ParamsContainerDiv) Node {
	return Div(
		append(
			params.Elements,

			If(
				len(params.IDContainer) > 0,

				AttrID(params.IDContainer),
			),
		)...,
	)
}
