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

func (co *CSSContribution) BackgroundColor(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"background-color", v},
	)
	return co
}

func (co *CSSContribution) BackgroundImage(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"background-image", v},
	)
	return co
}

func (co *CSSContribution) BackgroundSize(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"background-size", v},
	)
	return co
}

func (co *CSSContribution) BackgroundRepeat(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"background-repeat", v},
	)
	return co
}

func (co *CSSContribution) BackgroundPosition(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"background-position", v},
	)
	return co
}

func (co *CSSContribution) BorderTop(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"border-top", v},
	)
	return co
}

func (co *CSSContribution) BorderRight(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"border-right", v},
	)
	return co
}

func (co *CSSContribution) BorderBottom(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"border-bottom", v},
	)
	return co
}

func (co *CSSContribution) BorderLeft(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"border-left", v},
	)
	return co
}

func (co *CSSContribution) BorderColor(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"border-color", v},
	)
	return co
}

func (co *CSSContribution) BorderStyle(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"border-style", v},
	)
	return co
}

func (co *CSSContribution) BorderWidth(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"border-width", v},
	)
	return co
}

func (co *CSSContribution) FontSize(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"font-size", v},
	)

	return co
}

func (co *CSSContribution) Cursor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"cursor",
			value,
		},
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

func (co *CSSContribution) OutlineColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"outline-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OutlineStyle(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"outline-style",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OutlineWidth(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"outline-width",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Filter(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"filter",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BackdropFilter(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"backdrop-filter",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Transform(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"transform",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TransformOrigin(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"transform-origin",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Animation(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationDuration(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-duration",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationTimingFunction(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-timing-function",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationDelay(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-delay",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationIterationCount(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-iteration-count",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationDirection(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-direction",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationFillMode(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-fill-mode",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationPlayState(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-play-state",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AspectRatio(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"aspect-ratio",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BackfaceVisibility(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"backface-visibility",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderCollapse(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-collapse",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderRadius(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-radius",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BoxShadow(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"box-shadow",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BoxSizing(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"box-sizing",
			value,
		},
	)

	return co
}

func (co *CSSContribution) CaptionSide(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"caption-side",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ClipPath(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"clip-path",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Content(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"content",
			value,
		},
	)

	return co
}

func (co *CSSContribution) CounterIncrement(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"counter-increment",
			value,
		},
	)

	return co
}

func (co *CSSContribution) CounterReset(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"counter-reset",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Direction(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"direction",
			value,
		},
	)

	return co
}

func (co *CSSContribution) EmptyCells(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"empty-cells",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontKerning(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-kerning",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontStyle(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-style",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontVariant(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-variant",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ImageRendering(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"image-rendering",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Isolation(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"isolation",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ObjectViewBox(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"object-view-box",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ColumnCount(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"column-count",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ColumnGap(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"column-gap",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ColumnRule(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"column-rule",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ColumnWidth(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"column-width",
			value,
		},
	)

	return co
}

func (co *CSSContribution) GridTemplateColumns(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"grid-template-columns",
			value,
		},
	)

	return co
}

func (co *CSSContribution) GridTemplateRows(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"grid-template-rows",
			value,
		},
	)

	return co
}

func (co *CSSContribution) GridTemplateAreas(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"grid-template-areas",
			value,
		},
	)

	return co
}

func (co *CSSContribution) GridColumn(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"grid-column",
			value,
		},
	)

	return co
}

func (co *CSSContribution) GridRow(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"grid-row",
			value,
		},
	)

	return co
}

func (co *CSSContribution) GridAutoFlow(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"grid-auto-flow",
			value,
		},
	)

	return co
}

func (co *CSSContribution) GridAutoColumns(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"grid-auto-columns",
			value,
		},
	)

	return co
}

func (co *CSSContribution) GridAutoRows(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"grid-auto-rows",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Flex(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"flex",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FlexGrow(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"flex-grow",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FlexShrink(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"flex-shrink",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FlexBasis(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"flex-basis",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AlignContent(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"align-content",
			value,
		},
	)

	return co
}

func (co *CSSContribution) JustifyItems(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"justify-items",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PlaceItems(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"place-items",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PlaceContent(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"place-content",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PlaceSelf(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"place-self",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Order(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"order",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Orphans(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"orphans",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OutlineOffset(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"outline-offset",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PageBreakAfter(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"page-break-after",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PageBreakBefore(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"page-break-before",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PageBreakInside(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"page-break-inside",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Perspective(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"perspective",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PerspectiveOrigin(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"perspective-origin",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Quotes(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"quotes",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Resize(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"resize",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollBehavior(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-behavior",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollMargin(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-margin",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollPadding(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-padding",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ShapeOutside(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"shape-outside",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TabSize(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"tab-size",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TableLayout(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"table-layout",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextIndent(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-indent",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextOverflow(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-overflow",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TouchAction(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"touch-action",
			value,
		},
	)

	return co
}

func (co *CSSContribution) UnicodeBidi(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"unicode-bidi",
			value,
		},
	)

	return co
}

func (co *CSSContribution) UserZoom(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"user-zoom",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Float(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"float",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Clear(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"clear",
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

func (co *CSSContribution) PaddingTop(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"padding-top",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PaddingRight(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"padding-right",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PaddingBottom(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"padding-bottom",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PaddingLeft(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"padding-left",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontWeight(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-weight",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontFamily(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-family",
			value,
		},
	)

	return co
}

func (co *CSSContribution) LineHeight(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"line-height",
			value,
		},
	)

	return co
}

func (co *CSSContribution) LetterSpacing(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"letter-spacing",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextAlign(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-align",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextTransform(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-transform",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextDecoration(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-decoration",
			value,
		},
	)

	return co
}

func (co *CSSContribution) WhiteSpace(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"white-space",
			value,
		},
	)

	return co
}

func (co *CSSContribution) WordBreak(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"word-break",
			value,
		},
	)

	return co
}

func (co *CSSContribution) WillChange(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"will-change",
			value,
		},
	)

	return co
}

func (co *CSSContribution) WritingMode(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"writing-mode",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Zoom(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"zoom",
			value,
		},
	)

	return co
}

func (co *CSSContribution) RubyAlign(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"ruby-align",
			value,
		},
	)

	return co
}

func (co *CSSContribution) RubyPosition(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"ruby-position",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Mask(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskImage(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-image",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskMode(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-mode",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskRepeat(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-repeat",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskPosition(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-position",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskSize(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-size",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskComposite(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-composite",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskOrigin(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-origin",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskClip(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-clip",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskBorder(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-border",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskBorderMode(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-border-mode",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskBorderOutset(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-border-outset",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskBorderRepeat(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-border-repeat",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskBorderSlice(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-border-slice",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskBorderSource(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-border-source",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaskBorderWidth(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mask-border-width",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Hyphens(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"hyphens",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ImageOrientation(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"image-orientation",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ImageResolution(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"image-resolution",
			value,
		},
	)

	return co
}

func (co *CSSContribution) InitialLetter(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"initial-letter",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ListStyle(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"list-style",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ObjectFit(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"object-fit",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ObjectPosition(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"object-position",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Visibility(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"visibility",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PointerEvents(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"pointer-events",
			value,
		},
	)

	return co
}

func (co *CSSContribution) UserSelect(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"user-select",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Outline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"outline",
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

func (co *CSSContribution) Margin(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"margin", v},
	)
	return co
}

func (co *CSSContribution) MarginTop(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"margin-top", v},
	)
	return co
}

func (co *CSSContribution) MarginRight(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"margin-right", v},
	)
	return co
}

func (co *CSSContribution) MarginBottom(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"margin-bottom", v},
	)
	return co
}

func (co *CSSContribution) MarginLeft(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"margin-left", v},
	)
	return co
}

func (co *CSSContribution) Width(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"width", v},
	)
	return co
}

func (co *CSSContribution) Height(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"height", v},
	)
	return co
}

func (co *CSSContribution) MaxWidth(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"max-width", v},
	)
	return co
}

func (co *CSSContribution) MinWidth(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"min-width", v},
	)
	return co
}

func (co *CSSContribution) MaxHeight(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"max-height", v},
	)
	return co
}

func (co *CSSContribution) MinHeight(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"min-height", v},
	)
	return co
}

func (co *CSSContribution) Gap(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"gap", v},
	)
	return co
}

func (co *CSSContribution) RowGap(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"row-gap", v},
	)
	return co
}

func (co *CSSContribution) AlignItems(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"align-items", v},
	)
	return co
}

func (co *CSSContribution) JustifyContent(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"justify-content", v},
	)
	return co
}

func (co *CSSContribution) FlexDirection(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"flex-direction", v},
	)
	return co
}

func (co *CSSContribution) FlexWrap(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"flex-wrap", v},
	)
	return co
}

func (co *CSSContribution) Position(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"position", v},
	)
	return co
}

func (co *CSSContribution) ZIndex(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"z-index", v},
	)
	return co
}

func (co *CSSContribution) Top(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"top", v},
	)
	return co
}

func (co *CSSContribution) Right(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"right", v},
	)
	return co
}

func (co *CSSContribution) Bottom(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"bottom", v},
	)
	return co
}

func (co *CSSContribution) Left(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"left", v},
	)
	return co
}

func (co *CSSContribution) Overflow(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"overflow", v},
	)
	return co
}

func (co *CSSContribution) OverflowX(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"overflow-x", v},
	)
	return co
}

func (co *CSSContribution) OverflowY(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"overflow-y", v},
	)
	return co
}

func (co *CSSContribution) Opacity(v string) *CSSContribution {
	co.DeclarativeStyle = append(co.DeclarativeStyle,
		[2]string{"opacity", v},
	)
	return co
}

func (co *CSSContribution) LineBreak(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"line-break",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverflowWrap(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overflow-wrap",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BreakBefore(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"break-before",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BreakAfter(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"break-after",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BreakInside(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"break-inside",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextCombineUpright(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-combine-upright",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextOrientation(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-orientation",
			value,
		},
	)

	return co
}

func (co *CSSContribution) HangingPunctuation(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"hanging-punctuation",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontFeatureSettings(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-feature-settings",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontVariantCaps(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-variant-caps",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontVariantNumeric(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-variant-numeric",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontVariantLigatures(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-variant-ligatures",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontVariantEastAsian(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-variant-east-asian",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontSynthesis(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-synthesis",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontOpticalSizing(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-optical-sizing",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontPalette(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-palette",
			value,
		},
	)

	return co
}

func (co *CSSContribution) FontSizeAdjust(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"font-size-adjust",
			value,
		},
	)

	return co
}

func (co *CSSContribution) HyphenateCharacter(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"hyphenate-character",
			value,
		},
	)

	return co
}

func (co *CSSContribution) HyphenateLimitChars(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"hyphenate-limit-chars",
			value,
		},
	)

	return co
}

func (co *CSSContribution) HyphenateLimitLines(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"hyphenate-limit-lines",
			value,
		},
	)

	return co
}

func (co *CSSContribution) HyphenateLimitZone(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"hyphenate-limit-zone",
			value,
		},
	)

	return co
}

func (co *CSSContribution) LineClamp(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"line-clamp",
			value,
		},
	)

	return co
}

func (co *CSSContribution) NavUp(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"nav-up",
			value,
		},
	)

	return co
}

func (co *CSSContribution) NavDown(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"nav-down",
			value,
		},
	)

	return co
}

func (co *CSSContribution) NavLeft(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"nav-left",
			value,
		},
	)

	return co
}

func (co *CSSContribution) NavRight(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"nav-right",
			value,
		},
	)

	return co
}

func (co *CSSContribution) NavIndex(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"nav-index",
			value,
		},
	)

	return co
}

func (co *CSSContribution) NavAuto(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"nav-auto",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollSnapType(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-snap-type",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollSnapAlign(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-snap-align",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollSnapStop(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-snap-stop",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollMarginTop(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-margin-top",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollMarginRight(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-margin-right",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollMarginBottom(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-margin-bottom",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollMarginLeft(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-margin-left",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollPaddingTop(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-padding-top",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollPaddingRight(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-padding-right",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollPaddingBottom(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-padding-bottom",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollPaddingLeft(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-padding-left",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ShapeMargin(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"shape-margin",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ShapeImageThreshold(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"shape-image-threshold",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ShapeInside(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"shape-inside",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextEmphasis(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-emphasis",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextEmphasisColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-emphasis-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextEmphasisPosition(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-emphasis-position",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextEmphasisStyle(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-emphasis-style",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextRendering(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-rendering",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextShadow(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-shadow",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TouchActionDelay(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"touch-action-delay",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TransformStyle(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"transform-style",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Translate(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"translate",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TranslateX(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"translate-x",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TranslateY(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"translate-y",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TranslateZ(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"translate-z",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Scale(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scale",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScaleX(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scale-x",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScaleY(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scale-y",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScaleZ(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scale-z",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Rotate(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"rotate",
			value,
		},
	)

	return co
}

func (co *CSSContribution) RotateX(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"rotate-x",
			value,
		},
	)

	return co
}

func (co *CSSContribution) RotateY(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"rotate-y",
			value,
		},
	)

	return co
}

func (co *CSSContribution) RotateZ(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"rotate-z",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BackdropOpacity(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"backdrop-opacity",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BackdropBrightness(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"backdrop-brightness",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BackdropContrast(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"backdrop-contrast",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BackdropGrayscale(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"backdrop-grayscale",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BackdropHueRotate(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"backdrop-hue-rotate",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BackdropInvert(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"backdrop-invert",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BackdropSaturate(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"backdrop-saturate",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BackdropSepia(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"backdrop-sepia",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MixBlendMode(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"mix-blend-mode",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BackgroundBlendMode(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"background-blend-mode",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderSpacing(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-spacing",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderBlock(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-block",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderInline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-inline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderBlockColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-block-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderBlockStyle(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-block-style",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderBlockWidth(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-block-width",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderBlockStart(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-block-start",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderBlockEnd(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-block-end",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderInlineColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-inline-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderInlineStyle(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-inline-style",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderInlineWidth(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-inline-width",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderInlineStart(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-inline-start",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BorderInlineEnd(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"border-inline-end",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Inset(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"inset",
			value,
		},
	)

	return co
}

func (co *CSSContribution) InsetBlock(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"inset-block",
			value,
		},
	)

	return co
}

func (co *CSSContribution) InsetBlockStart(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"inset-block-start",
			value,
		},
	)

	return co
}

func (co *CSSContribution) InsetBlockEnd(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"inset-block-end",
			value,
		},
	)

	return co
}

func (co *CSSContribution) InsetInline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"inset-inline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) InsetInlineStart(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"inset-inline-start",
			value,
		},
	)

	return co
}

func (co *CSSContribution) InsetInlineEnd(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"inset-inline-end",
			value,
		},
	)

	return co
}

func (co *CSSContribution) BlockSize(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"block-size",
			value,
		},
	)

	return co
}

func (co *CSSContribution) InlineSize(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"inline-size",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MinBlockSize(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"min-block-size",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaxBlockSize(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"max-block-size",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MinInlineSize(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"min-inline-size",
			value,
		},
	)

	return co
}

func (co *CSSContribution) MaxInlineSize(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"max-inline-size",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverflowBlock(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overflow-block",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverflowInline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overflow-inline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehavior(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorX(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-x",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorY(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-y",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorBlock(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-block",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorInline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-inline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarWidth(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-width",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarGutter(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-gutter",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AccentColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"accent-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Appearance(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"appearance",
			value,
		},
	)

	return co
}

func (co *CSSContribution) CaretColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"caret-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) CaretShape(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"caret-shape",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ColorScheme(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"color-scheme",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ForcedColorAdjust(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"forced-color-adjust",
			value,
		},
	)

	return co
}

func (co *CSSContribution) PrintColorAdjust(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"print-color-adjust",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextDecorationColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-decoration-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextDecorationLine(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-decoration-line",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextDecorationThickness(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-decoration-thickness",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextUnderlineOffset(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-underline-offset",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextUnderlinePosition(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-underline-position",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextJustify(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-justify",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextSpacingTrim(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-spacing-trim",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextWrap(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-wrap",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextWrapStyle(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-wrap-style",
			value,
		},
	)

	return co
}

func (co *CSSContribution) TextWrapMode(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"text-wrap-mode",
			value,
		},
	)

	return co
}

func (co *CSSContribution) HangingPunctuationLast(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"hanging-punctuation-last",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorYBlock(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-y-block",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorYInline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-y-inline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorXBlock(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-x-block",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorXInline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-x-inline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollTimelineName(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-timeline-name",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollTimelineAxis(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-timeline-axis",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollTimelineScope(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-timeline-scope",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ViewTimelineName(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"view-timeline-name",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ViewTimelineAxis(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"view-timeline-axis",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ViewTimelineInset(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"view-timeline-inset",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationTimeline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-timeline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationRange(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-range",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationRangeStart(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-range-start",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationRangeEnd(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-range-end",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ContainerName(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"container-name",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ContainerType(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"container-type",
			value,
		},
	)

	return co
}

func (co *CSSContribution) Container(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"container",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ViewTransitionName(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"view-transition-name",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ViewTransitionClass(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"view-transition-class",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ViewTransitionGroup(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"view-transition-group",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ViewTransitionRoot(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"view-transition-root",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarColorTrack(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-color-track",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarColorThumb(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-color-thumb",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarGutterStable(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-gutter-stable",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarGutterBothEdges(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-gutter-both-edges",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorInlineStart(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-inline-start",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorInlineEnd(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-inline-end",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorBlockStart(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-block-start",
			value,
		},
	)

	return co
}

func (co *CSSContribution) OverscrollBehaviorBlockEnd(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"overscroll-behavior-block-end",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollTimeline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scroll-timeline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ViewTimeline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"view-timeline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ViewTimelineInsetBlock(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"view-timeline-inset-block",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ViewTimelineInsetInline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"view-timeline-inset-inline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationComposition(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-composition",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationDelayStart(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-delay-start",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationDelayEnd(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-delay-end",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationDurationStart(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-duration-start",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationDurationEnd(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-duration-end",
			value,
		},
	)

	return co
}

func (co *CSSContribution) AnimationTimelineScope(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"animation-timeline-scope",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarShadow(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-shadow",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackPieceColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-piece-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarButtonColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-button-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarCornerColor(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-corner-color",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarWidthThin(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-width-thin",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarWidthAuto(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-width-auto",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarWidthNone(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-width-none",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarGutterForce(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-gutter-force",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarGutterStableBoth(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-gutter-stable-both",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackBorder(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-border",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbBorder(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-border",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackRadius(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-radius",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbRadius(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-radius",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackShadow(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-shadow",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbShadow(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-shadow",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackOpacity(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-opacity",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbOpacity(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-opacity",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackInset(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-inset",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbInset(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-inset",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackBlendMode(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-blend-mode",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbBlendMode(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-blend-mode",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackFilter(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-filter",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbFilter(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-filter",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackBackdropFilter(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-backdrop-filter",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbBackdropFilter(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-backdrop-filter",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackAnimation(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-animation",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbAnimation(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-animation",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackTransition(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-transition",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbTransition(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-transition",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackOutline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-outline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbOutline(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-outline",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackOutlineOffset(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-outline-offset",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarThumbOutlineOffset(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-thumb-outline-offset",
			value,
		},
	)

	return co
}

func (co *CSSContribution) ScrollbarTrackOpacityTransition(value string) *CSSContribution {
	co.DeclarativeStyle = append(
		co.DeclarativeStyle,
		[2]string{
			"scrollbar-track-opacity-transition",
			value,
		},
	)

	return co
}
