package dsl

import (
	"fmt"
	"unsafe"

	"github.com/TudorHulban/hxgo/helpers"
)

type CSSContributionKey struct {
	Selector     string
	Breakpoint   string
	DesktopFirst bool
}

type CSS func() string

type CSSContribution struct {
	CSSContributionKey

	DeclarativeStyle [][2]string
	ProceduralCSS    []CSS
}

func NewGeneralCSS() *CSSContribution {
	return &CSSContribution{}
}

func NewCSSFor(selector string) *CSSContribution {
	return &CSSContribution{
		CSSContributionKey: CSSContributionKey{
			Selector: selector,
		},
	}
}

func NewCSSForClass(class string) *CSSContribution {
	return &CSSContribution{
		CSSContributionKey: CSSContributionKey{
			Selector: "." + helpers.RemoveBeforeFirstLetter(class),
		},
	}
}

func NewCSSForID(id string) *CSSContribution {
	return &CSSContribution{
		CSSContributionKey: CSSContributionKey{
			Selector: "#" + helpers.RemoveBeforeFirstLetter(id),
		},
	}
}

func (co *CSSContribution) AddStyles(styles ...[2]string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle, styles...)

	return co
}

func (co *CSSContribution) AddCSS(css ...CSS) *CSSContribution {
	co.ProceduralCSS = append(co.ProceduralCSS, css...)

	return co
}

func (co *CSSContribution) AtMin(px int) *CSSContribution {
	co.Breakpoint = fmt.Sprintf("%dpx", px)
	co.DesktopFirst = true

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

func (co *CSSContribution) WithBreakpoint(point string) *CSSContribution {
	co.Breakpoint = point

	return co
}

func (co *CSSContribution) WithStyles(styles ...[2]string) Node {
	co.DeclarativeStyle = append(co.DeclarativeStyle, styles...)

	return co.AsNode()
}

func renderCSSContribution(a *accumulator, data unsafe.Pointer) {
	co := (*CSSContribution)(data)

	a.css = append(a.css, *co)
}

func (co *CSSContribution) AsNode() Node {
	return Node{
		fn:    renderCSSContribution,
		data:  unsafe.Pointer(co),
		isCSS: true,
	}
}
