package dsl

import "fmt"

func (co *CSSContribution) WithBreakpoint(point string) *CSSContribution {
	co.Breakpoint = point

	return co
}

func (co *CSSContribution) AtMin(px int) *CSSContribution {
	co.Breakpoint = fmt.Sprintf("%dpx", px)
	co.DesktopFirst = true

	return co
}

// formats

func (co *CSSContribution) Mobile() *CSSContribution  { return co.AtMin(0) }
func (co *CSSContribution) Tablet() *CSSContribution  { return co.AtMin(768) }
func (co *CSSContribution) Desktop() *CSSContribution { return co.AtMin(1024) }
func (co *CSSContribution) Large() *CSSContribution   { return co.AtMin(1280) }
