package dsl

type Tailwind string

func TW() Tailwind { return "" }

func (t Tailwind) AsNode() Node {
	return AttrWithValue(
		"class",
		string(t),
	)
}

// Layout (12)

// display: block
func (t Tailwind) Block() Tailwind { return t + " block" }

// display: block (sm)
func (t Tailwind) BlockSm() Tailwind { return t + " sm:block" }

// display: block (md)
func (t Tailwind) BlockMd() Tailwind { return t + " md:block" }

// display: block (lg)
func (t Tailwind) BlockLg() Tailwind { return t + " lg:block" }

// display: block (xl)
func (t Tailwind) BlockXl() Tailwind { return t + " xl:block" }

// display: block (2xl)
func (t Tailwind) BlockX2l() Tailwind { return t + " 2xl:block" }

// display: inline-block
func (t Tailwind) InlineBlock() Tailwind    { return t + " inline-block" }
func (t Tailwind) InlineBlockSm() Tailwind  { return t + " sm:inline-block" }
func (t Tailwind) InlineBlockMd() Tailwind  { return t + " md:inline-block" }
func (t Tailwind) InlineBlockLg() Tailwind  { return t + " lg:inline-block" }
func (t Tailwind) InlineBlockXl() Tailwind  { return t + " xl:inline-block" }
func (t Tailwind) InlineBlockX2l() Tailwind { return t + " 2xl:inline-block" }

// display: inline
func (t Tailwind) Inline() Tailwind    { return t + " inline" }
func (t Tailwind) InlineSm() Tailwind  { return t + " sm:inline" }
func (t Tailwind) InlineMd() Tailwind  { return t + " md:inline" }
func (t Tailwind) InlineLg() Tailwind  { return t + " lg:inline" }
func (t Tailwind) InlineXl() Tailwind  { return t + " xl:inline" }
func (t Tailwind) InlineX2l() Tailwind { return t + " 2xl:inline" }

// display: flex
func (t Tailwind) Flex() Tailwind    { return t + " flex" }
func (t Tailwind) FlexSm() Tailwind  { return t + " sm:flex" }
func (t Tailwind) FlexMd() Tailwind  { return t + " md:flex" }
func (t Tailwind) FlexLg() Tailwind  { return t + " lg:flex" }
func (t Tailwind) FlexXl() Tailwind  { return t + " xl:flex" }
func (t Tailwind) FlexX2l() Tailwind { return t + " 2xl:flex" }

// display: grid
func (t Tailwind) Grid() Tailwind    { return t + " grid" }
func (t Tailwind) GridSm() Tailwind  { return t + " sm:grid" }
func (t Tailwind) GridMd() Tailwind  { return t + " md:grid" }
func (t Tailwind) GridLg() Tailwind  { return t + " lg:grid" }
func (t Tailwind) GridXl() Tailwind  { return t + " xl:grid" }
func (t Tailwind) GridX2l() Tailwind { return t + " 2xl:grid" }

// width: 100%
func (t Tailwind) WFull() Tailwind    { return t + " w-full" }
func (t Tailwind) WFullSm() Tailwind  { return t + " sm:w-full" }
func (t Tailwind) WFullMd() Tailwind  { return t + " md:w-full" }
func (t Tailwind) WFullLg() Tailwind  { return t + " lg:w-full" }
func (t Tailwind) WFullXl() Tailwind  { return t + " xl:w-full" }
func (t Tailwind) WFullX2l() Tailwind { return t + " 2xl:w-full" }

// width: auto
func (t Tailwind) WAuto() Tailwind    { return t + " w-auto" }
func (t Tailwind) WAutoSm() Tailwind  { return t + " sm:w-auto" }
func (t Tailwind) WAutoMd() Tailwind  { return t + " md:w-auto" }
func (t Tailwind) WAutoLg() Tailwind  { return t + " lg:w-auto" }
func (t Tailwind) WAutoXl() Tailwind  { return t + " xl:w-auto" }
func (t Tailwind) WAutoX2l() Tailwind { return t + " 2xl:w-auto" }

// max-width: full
func (t Tailwind) MaxWFull() Tailwind    { return t + " max-w-full" }
func (t Tailwind) MaxWFullSm() Tailwind  { return t + " sm:max-w-full" }
func (t Tailwind) MaxWFullMd() Tailwind  { return t + " md:max-w-full" }
func (t Tailwind) MaxWFullLg() Tailwind  { return t + " lg:max-w-full" }
func (t Tailwind) MaxWFullXl() Tailwind  { return t + " xl:max-w-full" }
func (t Tailwind) MaxWFullX2l() Tailwind { return t + " 2xl:max-w-full" }

// height: full
func (t Tailwind) HFull() Tailwind    { return t + " h-full" }
func (t Tailwind) HFullSm() Tailwind  { return t + " sm:h-full" }
func (t Tailwind) HFullMd() Tailwind  { return t + " md:h-full" }
func (t Tailwind) HFullLg() Tailwind  { return t + " lg:h-full" }
func (t Tailwind) HFullXl() Tailwind  { return t + " xl:h-full" }
func (t Tailwind) HFullX2l() Tailwind { return t + " 2xl:h-full" }

// position: relative
func (t Tailwind) Relative() Tailwind    { return t + " relative" }
func (t Tailwind) RelativeSm() Tailwind  { return t + " sm:relative" }
func (t Tailwind) RelativeMd() Tailwind  { return t + " md:relative" }
func (t Tailwind) RelativeLg() Tailwind  { return t + " lg:relative" }
func (t Tailwind) RelativeXl() Tailwind  { return t + " xl:relative" }
func (t Tailwind) RelativeX2l() Tailwind { return t + " 2xl:relative" }

// position: absolute
func (t Tailwind) Absolute() Tailwind    { return t + " absolute" }
func (t Tailwind) AbsoluteSm() Tailwind  { return t + " sm:absolute" }
func (t Tailwind) AbsoluteMd() Tailwind  { return t + " md:absolute" }
func (t Tailwind) AbsoluteLg() Tailwind  { return t + " lg:absolute" }
func (t Tailwind) AbsoluteXl() Tailwind  { return t + " xl:absolute" }
func (t Tailwind) AbsoluteX2l() Tailwind { return t + " 2xl:absolute" }

// overflow: hidden
func (t Tailwind) OverflowHidden() Tailwind    { return t + " overflow-hidden" }
func (t Tailwind) OverflowHiddenSm() Tailwind  { return t + " sm:overflow-hidden" }
func (t Tailwind) OverflowHiddenMd() Tailwind  { return t + " md:overflow-hidden" }
func (t Tailwind) OverflowHiddenLg() Tailwind  { return t + " lg:overflow-hidden" }
func (t Tailwind) OverflowHiddenXl() Tailwind  { return t + " xl:overflow-hidden" }
func (t Tailwind) OverflowHiddenX2l() Tailwind { return t + " 2xl:overflow-hidden" }

// Flexbox (10)

// flex-direction: row
func (t Tailwind) FlexRow() Tailwind    { return t + " flex-row" }
func (t Tailwind) FlexRowSm() Tailwind  { return t + " sm:flex-row" }
func (t Tailwind) FlexRowMd() Tailwind  { return t + " md:flex-row" }
func (t Tailwind) FlexRowLg() Tailwind  { return t + " lg:flex-row" }
func (t Tailwind) FlexRowXl() Tailwind  { return t + " xl:flex-row" }
func (t Tailwind) FlexRowX2l() Tailwind { return t + " 2xl:flex-row" }

// flex-direction: column
func (t Tailwind) FlexCol() Tailwind    { return t + " flex-col" }
func (t Tailwind) FlexColSm() Tailwind  { return t + " sm:flex-col" }
func (t Tailwind) FlexColMd() Tailwind  { return t + " md:flex-col" }
func (t Tailwind) FlexColLg() Tailwind  { return t + " lg:flex-col" }
func (t Tailwind) FlexColXl() Tailwind  { return t + " xl:flex-col" }
func (t Tailwind) FlexColX2l() Tailwind { return t + " 2xl:flex-col" }

// align-items: center
func (t Tailwind) ItemsCenter() Tailwind    { return t + " items-center" }
func (t Tailwind) ItemsCenterSm() Tailwind  { return t + " sm:items-center" }
func (t Tailwind) ItemsCenterMd() Tailwind  { return t + " md:items-center" }
func (t Tailwind) ItemsCenterLg() Tailwind  { return t + " lg:items-center" }
func (t Tailwind) ItemsCenterXl() Tailwind  { return t + " xl:items-center" }
func (t Tailwind) ItemsCenterX2l() Tailwind { return t + " 2xl:items-center" }

// align-items: flex-start
func (t Tailwind) ItemsStart() Tailwind    { return t + " items-start" }
func (t Tailwind) ItemsStartSm() Tailwind  { return t + " sm:items-start" }
func (t Tailwind) ItemsStartMd() Tailwind  { return t + " md:items-start" }
func (t Tailwind) ItemsStartLg() Tailwind  { return t + " lg:items-start" }
func (t Tailwind) ItemsStartXl() Tailwind  { return t + " xl:items-start" }
func (t Tailwind) ItemsStartX2l() Tailwind { return t + " 2xl:items-start" }

// align-items: flex-end
func (t Tailwind) ItemsEnd() Tailwind    { return t + " items-end" }
func (t Tailwind) ItemsEndSm() Tailwind  { return t + " sm:items-end" }
func (t Tailwind) ItemsEndMd() Tailwind  { return t + " md:items-end" }
func (t Tailwind) ItemsEndLg() Tailwind  { return t + " lg:items-end" }
func (t Tailwind) ItemsEndXl() Tailwind  { return t + " xl:items-end" }
func (t Tailwind) ItemsEndX2l() Tailwind { return t + " 2xl:items-end" }

// justify-content: center
func (t Tailwind) JustifyCenter() Tailwind    { return t + " justify-center" }
func (t Tailwind) JustifyCenterSm() Tailwind  { return t + " sm:justify-center" }
func (t Tailwind) JustifyCenterMd() Tailwind  { return t + " md:justify-center" }
func (t Tailwind) JustifyCenterLg() Tailwind  { return t + " lg:justify-center" }
func (t Tailwind) JustifyCenterXl() Tailwind  { return t + " xl:justify-center" }
func (t Tailwind) JustifyCenterX2l() Tailwind { return t + " 2xl:justify-center" }

// justify-content: space-between
func (t Tailwind) JustifyBetween() Tailwind    { return t + " justify-between" }
func (t Tailwind) JustifyBetweenSm() Tailwind  { return t + " sm:justify-between" }
func (t Tailwind) JustifyBetweenMd() Tailwind  { return t + " md:justify-between" }
func (t Tailwind) JustifyBetweenLg() Tailwind  { return t + " lg:justify-between" }
func (t Tailwind) JustifyBetweenXl() Tailwind  { return t + " xl:justify-between" }
func (t Tailwind) JustifyBetweenX2l() Tailwind { return t + " 2xl:justify-between" }

// justify-content: space-around
func (t Tailwind) JustifyAround() Tailwind    { return t + " justify-around" }
func (t Tailwind) JustifyAroundSm() Tailwind  { return t + " sm:justify-around" }
func (t Tailwind) JustifyAroundMd() Tailwind  { return t + " md:justify-around" }
func (t Tailwind) JustifyAroundLg() Tailwind  { return t + " lg:justify-around" }
func (t Tailwind) JustifyAroundXl() Tailwind  { return t + " xl:justify-around" }
func (t Tailwind) JustifyAroundX2l() Tailwind { return t + " 2xl:justify-around" }

// gap: 1rem
func (t Tailwind) Gap4() Tailwind    { return t + " gap-4" }
func (t Tailwind) Gap4Sm() Tailwind  { return t + " sm:gap-4" }
func (t Tailwind) Gap4Md() Tailwind  { return t + " md:gap-4" }
func (t Tailwind) Gap4Lg() Tailwind  { return t + " lg:gap-4" }
func (t Tailwind) Gap4Xl() Tailwind  { return t + " xl:gap-4" }
func (t Tailwind) Gap4X2l() Tailwind { return t + " 2xl:gap-4" }

// gap: 0.5rem
func (t Tailwind) Gap2() Tailwind    { return t + " gap-2" }
func (t Tailwind) Gap2Sm() Tailwind  { return t + " sm:gap-2" }
func (t Tailwind) Gap2Md() Tailwind  { return t + " md:gap-2" }
func (t Tailwind) Gap2Lg() Tailwind  { return t + " lg:gap-2" }
func (t Tailwind) Gap2Xl() Tailwind  { return t + " xl:gap-2" }
func (t Tailwind) Gap2X2l() Tailwind { return t + " 2xl:gap-2" }

// Typography (12)

// font-size: small
func (t Tailwind) TextSm() Tailwind    { return t + " text-sm" }
func (t Tailwind) TextSmSm() Tailwind  { return t + " sm:text-sm" }
func (t Tailwind) TextSmMd() Tailwind  { return t + " md:text-sm" }
func (t Tailwind) TextSmLg() Tailwind  { return t + " lg:text-sm" }
func (t Tailwind) TextSmXl() Tailwind  { return t + " xl:text-sm" }
func (t Tailwind) TextSmX2l() Tailwind { return t + " 2xl:text-sm" }

// font-size: base
func (t Tailwind) TextBase() Tailwind    { return t + " text-base" }
func (t Tailwind) TextBaseSm() Tailwind  { return t + " sm:text-base" }
func (t Tailwind) TextBaseMd() Tailwind  { return t + " md:text-base" }
func (t Tailwind) TextBaseLg() Tailwind  { return t + " lg:text-base" }
func (t Tailwind) TextBaseXl() Tailwind  { return t + " xl:text-base" }
func (t Tailwind) TextBaseX2l() Tailwind { return t + " 2xl:text-base" }

// font-size: large
func (t Tailwind) TextLg() Tailwind    { return t + " text-lg" }
func (t Tailwind) TextLgSm() Tailwind  { return t + " sm:text-lg" }
func (t Tailwind) TextLgMd() Tailwind  { return t + " md:text-lg" }
func (t Tailwind) TextLgLg() Tailwind  { return t + " lg:text-lg" }
func (t Tailwind) TextLgXl() Tailwind  { return t + " xl:text-lg" }
func (t Tailwind) TextLgX2l() Tailwind { return t + " 2xl:text-lg" }

// font-size: xl
func (t Tailwind) TextXl() Tailwind    { return t + " text-xl" }
func (t Tailwind) TextXlSm() Tailwind  { return t + " sm:text-xl" }
func (t Tailwind) TextXlMd() Tailwind  { return t + " md:text-xl" }
func (t Tailwind) TextXlLg() Tailwind  { return t + " lg:text-xl" }
func (t Tailwind) TextXlXl() Tailwind  { return t + " xl:text-xl" }
func (t Tailwind) TextXlX2l() Tailwind { return t + " 2xl:text-xl" }

// font-weight: normal
func (t Tailwind) FontNormal() Tailwind    { return t + " font-normal" }
func (t Tailwind) FontNormalSm() Tailwind  { return t + " sm:font-normal" }
func (t Tailwind) FontNormalMd() Tailwind  { return t + " md:font-normal" }
func (t Tailwind) FontNormalLg() Tailwind  { return t + " lg:font-normal" }
func (t Tailwind) FontNormalXl() Tailwind  { return t + " xl:font-normal" }
func (t Tailwind) FontNormalX2l() Tailwind { return t + " 2xl:font-normal" }

// font-weight: medium
func (t Tailwind) FontMedium() Tailwind    { return t + " font-medium" }
func (t Tailwind) FontMediumSm() Tailwind  { return t + " sm:font-medium" }
func (t Tailwind) FontMediumMd() Tailwind  { return t + " md:font-medium" }
func (t Tailwind) FontMediumLg() Tailwind  { return t + " lg:font-medium" }
func (t Tailwind) FontMediumXl() Tailwind  { return t + " xl:font-medium" }
func (t Tailwind) FontMediumX2l() Tailwind { return t + " 2xl:font-medium" }

// font-weight: semibold
func (t Tailwind) FontSemibold() Tailwind    { return t + " font-semibold" }
func (t Tailwind) FontSemiboldSm() Tailwind  { return t + " sm:font-semibold" }
func (t Tailwind) FontSemiboldMd() Tailwind  { return t + " md:font-semibold" }
func (t Tailwind) FontSemiboldLg() Tailwind  { return t + " lg:font-semibold" }
func (t Tailwind) FontSemiboldXl() Tailwind  { return t + " xl:font-semibold" }
func (t Tailwind) FontSemiboldX2l() Tailwind { return t + " 2xl:font-semibold" }

// font-weight: bold
func (t Tailwind) FontBold() Tailwind    { return t + " font-bold" }
func (t Tailwind) FontBoldSm() Tailwind  { return t + " sm:font-bold" }
func (t Tailwind) FontBoldMd() Tailwind  { return t + " md:font-bold" }
func (t Tailwind) FontBoldLg() Tailwind  { return t + " lg:font-bold" }
func (t Tailwind) FontBoldXl() Tailwind  { return t + " xl:font-bold" }
func (t Tailwind) FontBoldX2l() Tailwind { return t + " 2xl:font-bold" }

// line-height: normal
func (t Tailwind) LeadingNormal() Tailwind    { return t + " leading-normal" }
func (t Tailwind) LeadingNormalSm() Tailwind  { return t + " sm:leading-normal" }
func (t Tailwind) LeadingNormalMd() Tailwind  { return t + " md:leading-normal" }
func (t Tailwind) LeadingNormalLg() Tailwind  { return t + " lg:leading-normal" }
func (t Tailwind) LeadingNormalXl() Tailwind  { return t + " xl:leading-normal" }
func (t Tailwind) LeadingNormalX2l() Tailwind { return t + " 2xl:leading-normal" }

// line-height: snug
func (t Tailwind) LeadingSnug() Tailwind    { return t + " leading-snug" }
func (t Tailwind) LeadingSnugSm() Tailwind  { return t + " sm:leading-snug" }
func (t Tailwind) LeadingSnugMd() Tailwind  { return t + " md:leading-snug" }
func (t Tailwind) LeadingSnugLg() Tailwind  { return t + " lg:leading-snug" }
func (t Tailwind) LeadingSnugXl() Tailwind  { return t + " xl:leading-snug" }
func (t Tailwind) LeadingSnugX2l() Tailwind { return t + " 2xl:leading-snug" }

// text-align: center
func (t Tailwind) TextCenter() Tailwind    { return t + " text-center" }
func (t Tailwind) TextCenterSm() Tailwind  { return t + " sm:text-center" }
func (t Tailwind) TextCenterMd() Tailwind  { return t + " md:text-center" }
func (t Tailwind) TextCenterLg() Tailwind  { return t + " lg:text-center" }
func (t Tailwind) TextCenterXl() Tailwind  { return t + " xl:text-center" }
func (t Tailwind) TextCenterX2l() Tailwind { return t + " 2xl:text-center" }

// text-align: left
func (t Tailwind) TextLeft() Tailwind    { return t + " text-left" }
func (t Tailwind) TextLeftSm() Tailwind  { return t + " sm:text-left" }
func (t Tailwind) TextLeftMd() Tailwind  { return t + " md:text-left" }
func (t Tailwind) TextLeftLg() Tailwind  { return t + " lg:text-left" }
func (t Tailwind) TextLeftXl() Tailwind  { return t + " xl:text-left" }
func (t Tailwind) TextLeftX2l() Tailwind { return t + " 2xl:text-left" }

// Spacing (20)

// padding: 0.5rem
func (t Tailwind) P2() Tailwind    { return t + " p-2" }
func (t Tailwind) P2Sm() Tailwind  { return t + " sm:p-2" }
func (t Tailwind) P2Md() Tailwind  { return t + " md:p-2" }
func (t Tailwind) P2Lg() Tailwind  { return t + " lg:p-2" }
func (t Tailwind) P2Xl() Tailwind  { return t + " xl:p-2" }
func (t Tailwind) P2X2l() Tailwind { return t + " 2xl:p-2" }

// padding: 1rem
func (t Tailwind) P4() Tailwind    { return t + " p-4" }
func (t Tailwind) P4Sm() Tailwind  { return t + " sm:p-4" }
func (t Tailwind) P4Md() Tailwind  { return t + " md:p-4" }
func (t Tailwind) P4Lg() Tailwind  { return t + " lg:p-4" }
func (t Tailwind) P4Xl() Tailwind  { return t + " xl:p-4" }
func (t Tailwind) P4X2l() Tailwind { return t + " 2xl:p-4" }

// padding-x: 1rem
func (t Tailwind) Px4() Tailwind    { return t + " px-4" }
func (t Tailwind) Px4Sm() Tailwind  { return t + " sm:px-4" }
func (t Tailwind) Px4Md() Tailwind  { return t + " md:px-4" }
func (t Tailwind) Px4Lg() Tailwind  { return t + " lg:px-4" }
func (t Tailwind) Px4Xl() Tailwind  { return t + " xl:px-4" }
func (t Tailwind) Px4X2l() Tailwind { return t + " 2xl:px-4" }

// padding-y: 0.5rem
func (t Tailwind) Py2() Tailwind    { return t + " py-2" }
func (t Tailwind) Py2Sm() Tailwind  { return t + " sm:py-2" }
func (t Tailwind) Py2Md() Tailwind  { return t + " md:py-2" }
func (t Tailwind) Py2Lg() Tailwind  { return t + " lg:py-2" }
func (t Tailwind) Py2Xl() Tailwind  { return t + " xl:py-2" }
func (t Tailwind) Py2X2l() Tailwind { return t + " 2xl:py-2" }

// padding-left: 1rem
func (t Tailwind) Pl4() Tailwind    { return t + " pl-4" }
func (t Tailwind) Pl4Sm() Tailwind  { return t + " sm:pl-4" }
func (t Tailwind) Pl4Md() Tailwind  { return t + " md:pl-4" }
func (t Tailwind) Pl4Lg() Tailwind  { return t + " lg:pl-4" }
func (t Tailwind) Pl4Xl() Tailwind  { return t + " xl:pl-4" }
func (t Tailwind) Pl4X2l() Tailwind { return t + " 2xl:pl-4" }

// padding-right: 1rem
func (t Tailwind) Pr4() Tailwind    { return t + " pr-4" }
func (t Tailwind) Pr4Sm() Tailwind  { return t + " sm:pr-4" }
func (t Tailwind) Pr4Md() Tailwind  { return t + " md:pr-4" }
func (t Tailwind) Pr4Lg() Tailwind  { return t + " lg:pr-4" }
func (t Tailwind) Pr4Xl() Tailwind  { return t + " xl:pr-4" }
func (t Tailwind) Pr4X2l() Tailwind { return t + " 2xl:pr-4" }

// margin: 0.5rem
func (t Tailwind) M2() Tailwind    { return t + " m-2" }
func (t Tailwind) M2Sm() Tailwind  { return t + " sm:m-2" }
func (t Tailwind) M2Md() Tailwind  { return t + " md:m-2" }
func (t Tailwind) M2Lg() Tailwind  { return t + " lg:m-2" }
func (t Tailwind) M2Xl() Tailwind  { return t + " xl:m-2" }
func (t Tailwind) M2X2l() Tailwind { return t + " 2xl:m-2" }

// margin: 1rem
func (t Tailwind) M4() Tailwind    { return t + " m-4" }
func (t Tailwind) M4Sm() Tailwind  { return t + " sm:m-4" }
func (t Tailwind) M4Md() Tailwind  { return t + " md:m-4" }
func (t Tailwind) M4Lg() Tailwind  { return t + " lg:m-4" }
func (t Tailwind) M4Xl() Tailwind  { return t + " xl:m-4" }
func (t Tailwind) M4X2l() Tailwind { return t + " 2xl:m-4" }

// margin-top: 1rem
func (t Tailwind) Mt4() Tailwind    { return t + " mt-4" }
func (t Tailwind) Mt4Sm() Tailwind  { return t + " sm:mt-4" }
func (t Tailwind) Mt4Md() Tailwind  { return t + " md:mt-4" }
func (t Tailwind) Mt4Lg() Tailwind  { return t + " lg:mt-4" }
func (t Tailwind) Mt4Xl() Tailwind  { return t + " xl:mt-4" }
func (t Tailwind) Mt4X2l() Tailwind { return t + " 2xl:mt-4" }

// margin-bottom: 1rem
func (t Tailwind) Mb4() Tailwind    { return t + " mb-4" }
func (t Tailwind) Mb4Sm() Tailwind  { return t + " sm:mb-4" }
func (t Tailwind) Mb4Md() Tailwind  { return t + " md:mb-4" }
func (t Tailwind) Mb4Lg() Tailwind  { return t + " lg:mb-4" }
func (t Tailwind) Mb4Xl() Tailwind  { return t + " xl:mb-4" }
func (t Tailwind) Mb4X2l() Tailwind { return t + " 2xl:mb-4" }

// margin-left: 1rem
func (t Tailwind) Ml4() Tailwind    { return t + " ml-4" }
func (t Tailwind) Ml4Sm() Tailwind  { return t + " sm:ml-4" }
func (t Tailwind) Ml4Md() Tailwind  { return t + " md:ml-4" }
func (t Tailwind) Ml4Lg() Tailwind  { return t + " lg:ml-4" }
func (t Tailwind) Ml4Xl() Tailwind  { return t + " xl:ml-4" }
func (t Tailwind) Ml4X2l() Tailwind { return t + " 2xl:ml-4" }

// margin-right: 1rem
func (t Tailwind) Mr4() Tailwind    { return t + " mr-4" }
func (t Tailwind) Mr4Sm() Tailwind  { return t + " sm:mr-4" }
func (t Tailwind) Mr4Md() Tailwind  { return t + " md:mr-4" }
func (t Tailwind) Mr4Lg() Tailwind  { return t + " lg:mr-4" }
func (t Tailwind) Mr4Xl() Tailwind  { return t + " xl:mr-4" }
func (t Tailwind) Mr4X2l() Tailwind { return t + " 2xl:mr-4" }

// space-y: 0.5rem
func (t Tailwind) SpaceY2() Tailwind    { return t + " space-y-2" }
func (t Tailwind) SpaceY2Sm() Tailwind  { return t + " sm:space-y-2" }
func (t Tailwind) SpaceY2Md() Tailwind  { return t + " md:space-y-2" }
func (t Tailwind) SpaceY2Lg() Tailwind  { return t + " lg:space-y-2" }
func (t Tailwind) SpaceY2Xl() Tailwind  { return t + " xl:space-y-2" }
func (t Tailwind) SpaceY2X2l() Tailwind { return t + " 2xl:space-y-2" }

// space-y: 1rem
func (t Tailwind) SpaceY4() Tailwind    { return t + " space-y-4" }
func (t Tailwind) SpaceY4Sm() Tailwind  { return t + " sm:space-y-4" }
func (t Tailwind) SpaceY4Md() Tailwind  { return t + " md:space-y-4" }
func (t Tailwind) SpaceY4Lg() Tailwind  { return t + " lg:space-y-4" }
func (t Tailwind) SpaceY4Xl() Tailwind  { return t + " xl:space-y-4" }
func (t Tailwind) SpaceY4X2l() Tailwind { return t + " 2xl:space-y-4" }

// space-x: 0.5rem
func (t Tailwind) SpaceX2() Tailwind    { return t + " space-x-2" }
func (t Tailwind) SpaceX2Sm() Tailwind  { return t + " sm:space-x-2" }
func (t Tailwind) SpaceX2Md() Tailwind  { return t + " md:space-x-2" }
func (t Tailwind) SpaceX2Lg() Tailwind  { return t + " lg:space-x-2" }
func (t Tailwind) SpaceX2Xl() Tailwind  { return t + " xl:space-x-2" }
func (t Tailwind) SpaceX2X2l() Tailwind { return t + " 2xl:space-x-2" }

// space-x: 1rem
func (t Tailwind) SpaceX4() Tailwind    { return t + " space-x-4" }
func (t Tailwind) SpaceX4Sm() Tailwind  { return t + " sm:space-x-4" }
func (t Tailwind) SpaceX4Md() Tailwind  { return t + " md:space-x-4" }
func (t Tailwind) SpaceX4Lg() Tailwind  { return t + " lg:space-x-4" }
func (t Tailwind) SpaceX4Xl() Tailwind  { return t + " xl:space-x-4" }
func (t Tailwind) SpaceX4X2l() Tailwind { return t + " 2xl:space-x-4" }

// gap: 0.25rem
func (t Tailwind) Gap1() Tailwind    { return t + " gap-1" }
func (t Tailwind) Gap1Sm() Tailwind  { return t + " sm:gap-1" }
func (t Tailwind) Gap1Md() Tailwind  { return t + " md:gap-1" }
func (t Tailwind) Gap1Lg() Tailwind  { return t + " lg:gap-1" }
func (t Tailwind) Gap1Xl() Tailwind  { return t + " xl:gap-1" }
func (t Tailwind) Gap1X2l() Tailwind { return t + " 2xl:gap-1" }

// gap: 1rem (alt)
func (t Tailwind) Gap4Alt() Tailwind    { return t + " gap-4" }
func (t Tailwind) Gap4AltSm() Tailwind  { return t + " sm:gap-4" }
func (t Tailwind) Gap4AltMd() Tailwind  { return t + " md:gap-4" }
func (t Tailwind) Gap4AltLg() Tailwind  { return t + " lg:gap-4" }
func (t Tailwind) Gap4AltXl() Tailwind  { return t + " xl:gap-4" }
func (t Tailwind) Gap4AltX2l() Tailwind { return t + " 2xl:gap-4" }

// inset: 0
func (t Tailwind) Inset0() Tailwind    { return t + " inset-0" }
func (t Tailwind) Inset0Sm() Tailwind  { return t + " sm:inset-0" }
func (t Tailwind) Inset0Md() Tailwind  { return t + " md:inset-0" }
func (t Tailwind) Inset0Lg() Tailwind  { return t + " lg:inset-0" }
func (t Tailwind) Inset0Xl() Tailwind  { return t + " xl:inset-0" }
func (t Tailwind) Inset0X2l() Tailwind { return t + " 2xl:inset-0" }

// inset-x: 0
func (t Tailwind) InsetX0() Tailwind    { return t + " inset-x-0" }
func (t Tailwind) InsetX0Sm() Tailwind  { return t + " sm:inset-x-0" }
func (t Tailwind) InsetX0Md() Tailwind  { return t + " md:inset-x-0" }
func (t Tailwind) InsetX0Lg() Tailwind  { return t + " lg:inset-x-0" }
func (t Tailwind) InsetX0Xl() Tailwind  { return t + " xl:inset-x-0" }
func (t Tailwind) InsetX0X2l() Tailwind { return t + " 2xl:inset-x-0" }

// Colors (12)

// text color: white
func (t Tailwind) TextWhite() Tailwind    { return t + " text-white" }
func (t Tailwind) TextWhiteSm() Tailwind  { return t + " sm:text-white" }
func (t Tailwind) TextWhiteMd() Tailwind  { return t + " md:text-white" }
func (t Tailwind) TextWhiteLg() Tailwind  { return t + " lg:text-white" }
func (t Tailwind) TextWhiteXl() Tailwind  { return t + " xl:text-white" }
func (t Tailwind) TextWhiteX2l() Tailwind { return t + " 2xl:text-white" }

// text color: black
func (t Tailwind) TextBlack() Tailwind    { return t + " text-black" }
func (t Tailwind) TextBlackSm() Tailwind  { return t + " sm:text-black" }
func (t Tailwind) TextBlackMd() Tailwind  { return t + " md:text-black" }
func (t Tailwind) TextBlackLg() Tailwind  { return t + " lg:text-black" }
func (t Tailwind) TextBlackXl() Tailwind  { return t + " xl:text-black" }
func (t Tailwind) TextBlackX2l() Tailwind { return t + " 2xl:text-black" }

// text color: gray-700
func (t Tailwind) TextGray700() Tailwind    { return t + " text-gray-700" }
func (t Tailwind) TextGray700Sm() Tailwind  { return t + " sm:text-gray-700" }
func (t Tailwind) TextGray700Md() Tailwind  { return t + " md:text-gray-700" }
func (t Tailwind) TextGray700Lg() Tailwind  { return t + " lg:text-gray-700" }
func (t Tailwind) TextGray700Xl() Tailwind  { return t + " xl:text-gray-700" }
func (t Tailwind) TextGray700X2l() Tailwind { return t + " 2xl:text-gray-700" }

// background: white
func (t Tailwind) BgWhite() Tailwind    { return t + " bg-white" }
func (t Tailwind) BgWhiteSm() Tailwind  { return t + " sm:bg-white" }
func (t Tailwind) BgWhiteMd() Tailwind  { return t + " md:bg-white" }
func (t Tailwind) BgWhiteLg() Tailwind  { return t + " lg:bg-white" }
func (t Tailwind) BgWhiteXl() Tailwind  { return t + " xl:bg-white" }
func (t Tailwind) BgWhiteX2l() Tailwind { return t + " 2xl:bg-white" }

// background: black
func (t Tailwind) BgBlack() Tailwind    { return t + " bg-black" }
func (t Tailwind) BgBlackSm() Tailwind  { return t + " sm:bg-black" }
func (t Tailwind) BgBlackMd() Tailwind  { return t + " md:bg-black" }
func (t Tailwind) BgBlackLg() Tailwind  { return t + " lg:bg-black" }
func (t Tailwind) BgBlackXl() Tailwind  { return t + " xl:bg-black" }
func (t Tailwind) BgBlackX2l() Tailwind { return t + " 2xl:bg-black" }

// background: gray-100
func (t Tailwind) BgGray100() Tailwind    { return t + " bg-gray-100" }
func (t Tailwind) BgGray100Sm() Tailwind  { return t + " sm:bg-gray-100" }
func (t Tailwind) BgGray100Md() Tailwind  { return t + " md:bg-gray-100" }
func (t Tailwind) BgGray100Lg() Tailwind  { return t + " lg:bg-gray-100" }
func (t Tailwind) BgGray100Xl() Tailwind  { return t + " xl:bg-gray-100" }
func (t Tailwind) BgGray100X2l() Tailwind { return t + " 2xl:bg-gray-100" }

// background: blue-500
func (t Tailwind) BgBlue500() Tailwind    { return t + " bg-blue-500" }
func (t Tailwind) BgBlue500Sm() Tailwind  { return t + " sm:bg-blue-500" }
func (t Tailwind) BgBlue500Md() Tailwind  { return t + " md:bg-blue-500" }
func (t Tailwind) BgBlue500Lg() Tailwind  { return t + " lg:bg-blue-500" }
func (t Tailwind) BgBlue500Xl() Tailwind  { return t + " xl:bg-blue-500" }
func (t Tailwind) BgBlue500X2l() Tailwind { return t + " 2xl:bg-blue-500" }

// background: red-500
func (t Tailwind) BgRed500() Tailwind    { return t + " bg-red-500" }
func (t Tailwind) BgRed500Sm() Tailwind  { return t + " sm:bg-red-500" }
func (t Tailwind) BgRed500Md() Tailwind  { return t + " md:bg-red-500" }
func (t Tailwind) BgRed500Lg() Tailwind  { return t + " lg:bg-red-500" }
func (t Tailwind) BgRed500Xl() Tailwind  { return t + " xl:bg-red-500" }
func (t Tailwind) BgRed500X2l() Tailwind { return t + " 2xl:bg-red-500" }

// border color: gray-300
func (t Tailwind) BorderGray300() Tailwind    { return t + " border-gray-300" }
func (t Tailwind) BorderGray300Sm() Tailwind  { return t + " sm:border-gray-300" }
func (t Tailwind) BorderGray300Md() Tailwind  { return t + " md:border-gray-300" }
func (t Tailwind) BorderGray300Lg() Tailwind  { return t + " lg:border-gray-300" }
func (t Tailwind) BorderGray300Xl() Tailwind  { return t + " xl:border-gray-300" }
func (t Tailwind) BorderGray300X2l() Tailwind { return t + " 2xl:border-gray-300" }

// border color: blue-500
func (t Tailwind) BorderBlue500() Tailwind    { return t + " border-blue-500" }
func (t Tailwind) BorderBlue500Sm() Tailwind  { return t + " sm:border-blue-500" }
func (t Tailwind) BorderBlue500Md() Tailwind  { return t + " md:border-blue-500" }
func (t Tailwind) BorderBlue500Lg() Tailwind  { return t + " lg:border-blue-500" }
func (t Tailwind) BorderBlue500Xl() Tailwind  { return t + " xl:border-blue-500" }
func (t Tailwind) BorderBlue500X2l() Tailwind { return t + " 2xl:border-blue-500" }

// text opacity: 75%
func (t Tailwind) TextOpacity75() Tailwind    { return t + " text-opacity-75" }
func (t Tailwind) TextOpacity75Sm() Tailwind  { return t + " sm:text-opacity-75" }
func (t Tailwind) TextOpacity75Md() Tailwind  { return t + " md:text-opacity-75" }
func (t Tailwind) TextOpacity75Lg() Tailwind  { return t + " lg:text-opacity-75" }
func (t Tailwind) TextOpacity75Xl() Tailwind  { return t + " xl:text-opacity-75" }
func (t Tailwind) TextOpacity75X2l() Tailwind { return t + " 2xl:text-opacity-75" }

// bg opacity: 75%
func (t Tailwind) BgOpacity75() Tailwind    { return t + " bg-opacity-75" }
func (t Tailwind) BgOpacity75Sm() Tailwind  { return t + " sm:bg-opacity-75" }
func (t Tailwind) BgOpacity75Md() Tailwind  { return t + " md:bg-opacity-75" }
func (t Tailwind) BgOpacity75Lg() Tailwind  { return t + " lg:bg-opacity-75" }
func (t Tailwind) BgOpacity75Xl() Tailwind  { return t + " xl:bg-opacity-75" }
func (t Tailwind) BgOpacity75X2l() Tailwind { return t + " 2xl:bg-opacity-75" }

// Borders & Radius (8)

// border: 1px solid
func (t Tailwind) Border() Tailwind    { return t + " border" }
func (t Tailwind) BorderSm() Tailwind  { return t + " sm:border" }
func (t Tailwind) BorderMd() Tailwind  { return t + " md:border" }
func (t Tailwind) BorderLg() Tailwind  { return t + " lg:border" }
func (t Tailwind) BorderXl() Tailwind  { return t + " xl:border" }
func (t Tailwind) BorderX2l() Tailwind { return t + " 2xl:border" }

// border width: 2px
func (t Tailwind) Border2() Tailwind    { return t + " border-2" }
func (t Tailwind) Border2Sm() Tailwind  { return t + " sm:border-2" }
func (t Tailwind) Border2Md() Tailwind  { return t + " md:border-2" }
func (t Tailwind) Border2Lg() Tailwind  { return t + " lg:border-2" }
func (t Tailwind) Border2Xl() Tailwind  { return t + " xl:border-2" }
func (t Tailwind) Border2X2l() Tailwind { return t + " 2xl:border-2" }

// border radius: small
func (t Tailwind) Rounded() Tailwind    { return t + " rounded" }
func (t Tailwind) RoundedSm() Tailwind  { return t + " sm:rounded" }
func (t Tailwind) RoundedMd() Tailwind  { return t + " md:rounded" }
func (t Tailwind) RoundedLg() Tailwind  { return t + " lg:rounded" }
func (t Tailwind) RoundedXl() Tailwind  { return t + " xl:rounded" }
func (t Tailwind) RoundedX2l() Tailwind { return t + " 2xl:rounded" }

// border radius: medium
func (t Tailwind) RoundedMdSize() Tailwind    { return t + " rounded-md" }
func (t Tailwind) RoundedMdSizeSm() Tailwind  { return t + " sm:rounded-md" }
func (t Tailwind) RoundedMdSizeMd() Tailwind  { return t + " md:rounded-md" }
func (t Tailwind) RoundedMdSizeLg() Tailwind  { return t + " lg:rounded-md" }
func (t Tailwind) RoundedMdSizeXl() Tailwind  { return t + " xl:rounded-md" }
func (t Tailwind) RoundedMdSizeX2l() Tailwind { return t + " 2xl:rounded-md" }

// border radius: large
func (t Tailwind) RoundedLgSize() Tailwind    { return t + " rounded-lg" }
func (t Tailwind) RoundedLgSizeSm() Tailwind  { return t + " sm:rounded-lg" }
func (t Tailwind) RoundedLgSizeMd() Tailwind  { return t + " md:rounded-lg" }
func (t Tailwind) RoundedLgSizeLg() Tailwind  { return t + " lg:rounded-lg" }
func (t Tailwind) RoundedLgSizeXl() Tailwind  { return t + " xl:rounded-lg" }
func (t Tailwind) RoundedLgSizeX2l() Tailwind { return t + " 2xl:rounded-lg" }

// border radius: full
func (t Tailwind) RoundedFull() Tailwind    { return t + " rounded-full" }
func (t Tailwind) RoundedFullSm() Tailwind  { return t + " sm:rounded-full" }
func (t Tailwind) RoundedFullMd() Tailwind  { return t + " md:rounded-full" }
func (t Tailwind) RoundedFullLg() Tailwind  { return t + " lg:rounded-full" }
func (t Tailwind) RoundedFullXl() Tailwind  { return t + " xl:rounded-full" }
func (t Tailwind) RoundedFullX2l() Tailwind { return t + " 2xl:rounded-full" }

// border style: dashed
func (t Tailwind) BorderDashed() Tailwind    { return t + " border-dashed" }
func (t Tailwind) BorderDashedSm() Tailwind  { return t + " sm:border-dashed" }
func (t Tailwind) BorderDashedMd() Tailwind  { return t + " md:border-dashed" }
func (t Tailwind) BorderDashedLg() Tailwind  { return t + " lg:border-dashed" }
func (t Tailwind) BorderDashedXl() Tailwind  { return t + " xl:border-dashed" }
func (t Tailwind) BorderDashedX2l() Tailwind { return t + " 2xl:border-dashed" }

// border style: dotted
func (t Tailwind) BorderDotted() Tailwind    { return t + " border-dotted" }
func (t Tailwind) BorderDottedSm() Tailwind  { return t + " sm:border-dotted" }
func (t Tailwind) BorderDottedMd() Tailwind  { return t + " md:border-dotted" }
func (t Tailwind) BorderDottedLg() Tailwind  { return t + " lg:border-dotted" }
func (t Tailwind) BorderDottedXl() Tailwind  { return t + " xl:border-dotted" }
func (t Tailwind) BorderDottedX2l() Tailwind { return t + " 2xl:border-dotted" }

// Effects (6)

// shadow: small
func (t Tailwind) Shadow() Tailwind    { return t + " shadow" }
func (t Tailwind) ShadowSm() Tailwind  { return t + " sm:shadow" }
func (t Tailwind) ShadowMd() Tailwind  { return t + " md:shadow" }
func (t Tailwind) ShadowLg() Tailwind  { return t + " lg:shadow" }
func (t Tailwind) ShadowXl() Tailwind  { return t + " xl:shadow" }
func (t Tailwind) ShadowX2l() Tailwind { return t + " 2xl:shadow" }

// shadow: medium
func (t Tailwind) ShadowMdSize() Tailwind    { return t + " shadow-md" }
func (t Tailwind) ShadowMdSizeSm() Tailwind  { return t + " sm:shadow-md" }
func (t Tailwind) ShadowMdSizeMd() Tailwind  { return t + " md:shadow-md" }
func (t Tailwind) ShadowMdSizeLg() Tailwind  { return t + " lg:shadow-md" }
func (t Tailwind) ShadowMdSizeXl() Tailwind  { return t + " xl:shadow-md" }
func (t Tailwind) ShadowMdSizeX2l() Tailwind { return t + " 2xl:shadow-md" }

// opacity: 50%
func (t Tailwind) Opacity50() Tailwind    { return t + " opacity-50" }
func (t Tailwind) Opacity50Sm() Tailwind  { return t + " sm:opacity-50" }
func (t Tailwind) Opacity50Md() Tailwind  { return t + " md:opacity-50" }
func (t Tailwind) Opacity50Lg() Tailwind  { return t + " lg:opacity-50" }
func (t Tailwind) Opacity50Xl() Tailwind  { return t + " xl:opacity-50" }
func (t Tailwind) Opacity50X2l() Tailwind { return t + " 2xl:opacity-50" }

// opacity: 75%
func (t Tailwind) Opacity75() Tailwind    { return t + " opacity-75" }
func (t Tailwind) Opacity75Sm() Tailwind  { return t + " sm:opacity-75" }
func (t Tailwind) Opacity75Md() Tailwind  { return t + " md:opacity-75" }
func (t Tailwind) Opacity75Lg() Tailwind  { return t + " lg:opacity-75" }
func (t Tailwind) Opacity75Xl() Tailwind  { return t + " xl:opacity-75" }
func (t Tailwind) Opacity75X2l() Tailwind { return t + " 2xl:opacity-75" }

// blur: small
func (t Tailwind) BlurSm() Tailwind    { return t + " blur-sm" }
func (t Tailwind) BlurSmSm() Tailwind  { return t + " sm:blur-sm" }
func (t Tailwind) BlurSmMd() Tailwind  { return t + " md:blur-sm" }
func (t Tailwind) BlurSmLg() Tailwind  { return t + " lg:blur-sm" }
func (t Tailwind) BlurSmXl() Tailwind  { return t + " xl:blur-sm" }
func (t Tailwind) BlurSmX2l() Tailwind { return t + " 2xl:blur-sm" }

// blur: none
func (t Tailwind) BlurNone() Tailwind    { return t + " blur-none" }
func (t Tailwind) BlurNoneSm() Tailwind  { return t + " sm:blur-none" }
func (t Tailwind) BlurNoneMd() Tailwind  { return t + " md:blur-none" }
func (t Tailwind) BlurNoneLg() Tailwind  { return t + " lg:blur-none" }
func (t Tailwind) BlurNoneXl() Tailwind  { return t + " xl:blur-none" }
func (t Tailwind) BlurNoneX2l() Tailwind { return t + " 2xl:blur-none" }

// Transitions (6)

// enable transitions
func (t Tailwind) Transition() Tailwind    { return t + " transition" }
func (t Tailwind) TransitionSm() Tailwind  { return t + " sm:transition" }
func (t Tailwind) TransitionMd() Tailwind  { return t + " md:transition" }
func (t Tailwind) TransitionLg() Tailwind  { return t + " lg:transition" }
func (t Tailwind) TransitionXl() Tailwind  { return t + " xl:transition" }
func (t Tailwind) TransitionX2l() Tailwind { return t + " 2xl:transition" }

// transition: colors
func (t Tailwind) TransitionColors() Tailwind    { return t + " transition-colors" }
func (t Tailwind) TransitionColorsSm() Tailwind  { return t + " sm:transition-colors" }
func (t Tailwind) TransitionColorsMd() Tailwind  { return t + " md:transition-colors" }
func (t Tailwind) TransitionColorsLg() Tailwind  { return t + " lg:transition-colors" }
func (t Tailwind) TransitionColorsXl() Tailwind  { return t + " xl:transition-colors" }
func (t Tailwind) TransitionColorsX2l() Tailwind { return t + " 2xl:transition-colors" }

// transition: opacity
func (t Tailwind) TransitionOpacity() Tailwind    { return t + " transition-opacity" }
func (t Tailwind) TransitionOpacitySm() Tailwind  { return t + " sm:transition-opacity" }
func (t Tailwind) TransitionOpacityMd() Tailwind  { return t + " md:transition-opacity" }
func (t Tailwind) TransitionOpacityLg() Tailwind  { return t + " lg:transition-opacity" }
func (t Tailwind) TransitionOpacityXl() Tailwind  { return t + " xl:transition-opacity" }
func (t Tailwind) TransitionOpacityX2l() Tailwind { return t + " 2xl:transition-opacity" }

// duration: 150ms
func (t Tailwind) Duration150() Tailwind    { return t + " duration-150" }
func (t Tailwind) Duration150Sm() Tailwind  { return t + " sm:duration-150" }
func (t Tailwind) Duration150Md() Tailwind  { return t + " md:duration-150" }
func (t Tailwind) Duration150Lg() Tailwind  { return t + " lg:duration-150" }
func (t Tailwind) Duration150Xl() Tailwind  { return t + " xl:duration-150" }
func (t Tailwind) Duration150X2l() Tailwind { return t + " 2xl:duration-150" }

// duration: 300ms
func (t Tailwind) Duration300() Tailwind    { return t + " duration-300" }
func (t Tailwind) Duration300Sm() Tailwind  { return t + " sm:duration-300" }
func (t Tailwind) Duration300Md() Tailwind  { return t + " md:duration-300" }
func (t Tailwind) Duration300Lg() Tailwind  { return t + " lg:duration-300" }
func (t Tailwind) Duration300Xl() Tailwind  { return t + " xl:duration-300" }
func (t Tailwind) Duration300X2l() Tailwind { return t + " 2xl:duration-300" }

// ease: in-out
func (t Tailwind) EaseInOut() Tailwind    { return t + " ease-in-out" }
func (t Tailwind) EaseInOutSm() Tailwind  { return t + " sm:ease-in-out" }
func (t Tailwind) EaseInOutMd() Tailwind  { return t + " md:ease-in-out" }
func (t Tailwind) EaseInOutLg() Tailwind  { return t + " lg:ease-in-out" }
func (t Tailwind) EaseInOutXl() Tailwind  { return t + " xl:ease-in-out" }
func (t Tailwind) EaseInOutX2l() Tailwind { return t + " 2xl:ease-in-out" }

// Misc (6)

// cursor: pointer
func (t Tailwind) CursorPointer() Tailwind    { return t + " cursor-pointer" }
func (t Tailwind) CursorPointerSm() Tailwind  { return t + " sm:cursor-pointer" }
func (t Tailwind) CursorPointerMd() Tailwind  { return t + " md:cursor-pointer" }
func (t Tailwind) CursorPointerLg() Tailwind  { return t + " lg:cursor-pointer" }
func (t Tailwind) CursorPointerXl() Tailwind  { return t + " xl:cursor-pointer" }
func (t Tailwind) CursorPointerX2l() Tailwind { return t + " 2xl:cursor-pointer" }

// cursor: default
func (t Tailwind) CursorDefault() Tailwind    { return t + " cursor-default" }
func (t Tailwind) CursorDefaultSm() Tailwind  { return t + " sm:cursor-default" }
func (t Tailwind) CursorDefaultMd() Tailwind  { return t + " md:cursor-default" }
func (t Tailwind) CursorDefaultLg() Tailwind  { return t + " lg:cursor-default" }
func (t Tailwind) CursorDefaultXl() Tailwind  { return t + " xl:cursor-default" }
func (t Tailwind) CursorDefaultX2l() Tailwind { return t + " 2xl:cursor-default" }

// select: none
func (t Tailwind) SelectNone() Tailwind    { return t + " select-none" }
func (t Tailwind) SelectNoneSm() Tailwind  { return t + " sm:select-none" }
func (t Tailwind) SelectNoneMd() Tailwind  { return t + " md:select-none" }
func (t Tailwind) SelectNoneLg() Tailwind  { return t + " lg:select-none" }
func (t Tailwind) SelectNoneXl() Tailwind  { return t + " xl:select-none" }
func (t Tailwind) SelectNoneX2l() Tailwind { return t + " 2xl:select-none" }

// select: text
func (t Tailwind) SelectText() Tailwind    { return t + " select-text" }
func (t Tailwind) SelectTextSm() Tailwind  { return t + " sm:select-text" }
func (t Tailwind) SelectTextMd() Tailwind  { return t + " md:select-text" }
func (t Tailwind) SelectTextLg() Tailwind  { return t + " lg:select-text" }
func (t Tailwind) SelectTextXl() Tailwind  { return t + " xl:select-text" }
func (t Tailwind) SelectTextX2l() Tailwind { return t + " 2xl:select-text" }

// appearance: none
func (t Tailwind) AppearanceNone() Tailwind    { return t + " appearance-none" }
func (t Tailwind) AppearanceNoneSm() Tailwind  { return t + " sm:appearance-none" }
func (t Tailwind) AppearanceNoneMd() Tailwind  { return t + " md:appearance-none" }
func (t Tailwind) AppearanceNoneLg() Tailwind  { return t + " lg:appearance-none" }
func (t Tailwind) AppearanceNoneXl() Tailwind  { return t + " xl:appearance-none" }
func (t Tailwind) AppearanceNoneX2l() Tailwind { return t + " 2xl:appearance-none" }

// outline: none
func (t Tailwind) OutlineNone() Tailwind    { return t + " outline-none" }
func (t Tailwind) OutlineNoneSm() Tailwind  { return t + " sm:outline-none" }
func (t Tailwind) OutlineNoneMd() Tailwind  { return t + " md:outline-none" }
func (t Tailwind) OutlineNoneLg() Tailwind  { return t + " lg:outline-none" }
func (t Tailwind) OutlineNoneXl() Tailwind  { return t + " xl:outline-none" }
func (t Tailwind) OutlineNoneX2l() Tailwind { return t + " 2xl:outline-none" }
