package dsl

import (
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
