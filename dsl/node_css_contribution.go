package dsl

import (
	"reflect"
	"unsafe"
)

type CSSContributionKey struct {
	Selector       string
	InflexionPoint string
	DesktopFirst   bool
}

type CSS func() string

type CSSContribution struct {
	CSSContributionKey

	DeclarativeStyle [][2]string
	ProceduralCSS    []CSS
}

func NewCSSContribution(params CSSContributionKey) *CSSContribution {
	return &CSSContribution{
		CSSContributionKey: params,
	}
}

func (co *CSSContribution) AddStyles(styles ...[2]string) {
	co.DeclarativeStyle = append(co.DeclarativeStyle, styles...)
}

func (co *CSSContribution) AddCSS(css ...CSS) {
	co.ProceduralCSS = append(co.ProceduralCSS, css...)
}

func renderCSSContribution(a *accumulator, data unsafe.Pointer) {
	co := (*CSSContribution)(data)

	a.css = append(a.css, *co)
}

// TODO: review
func (co CSSContribution) AsNode() Node {
	data := &co

	return Node{
		fn:    renderCSSContribution,
		data:  unsafe.Pointer(data),
		isCSS: true,
	}
}

// internal grouping structures
type procBucket struct {
	seen  map[uintptr]struct{}
	order []func() string
}

func newProcBucket() *procBucket {
	return &procBucket{
		seen:  make(map[uintptr]struct{}),
		order: make([]func() string, 0),
	}
}

func (b *procBucket) Add(fn func() string) {
	ptr := reflect.ValueOf(fn).Pointer() // TODO: review

	if _, exists := b.seen[ptr]; exists {
		return
	}

	b.seen[ptr] = struct{}{}
	b.order = append(b.order, fn)
}

type group struct {
	Declarative map[string]string // prop -> value
	Procedural  *procBucket
}

func newGroup() *group {
	return &group{
		Declarative: make(map[string]string),
		Procedural:  newProcBucket(),
	}
}
