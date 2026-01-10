package dsl

// pseudo classes

func (co *CSSContribution) Hover() *CSSContribution {
	co.Selector = co.Selector + ":hover"

	return co
}

func (co *CSSContribution) Active() *CSSContribution {
	co.Selector = co.Selector + ":active"

	return co
}

func (co *CSSContribution) Display(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"display", v},
	)

	return co
}

func (co *CSSContribution) Background(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"background", v},
	)

	return co
}

func (co *CSSContribution) Border(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"border", v},
	)

	return co
}

func (co *CSSContribution) FontSize(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"font-size", v},
	)

	return co
}

func (co *CSSContribution) Cursor(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"cursor", v},
	)

	return co
}

func (co *CSSContribution) Transition(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"transition", v},
	)

	return co
}

func (co *CSSContribution) Color(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ShadowBox(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"box-shadow",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Padding(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"padding",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Radius(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-radius",
			value,
		},
	)

	return co
}
