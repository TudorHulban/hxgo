package tailwindcssgeneration

var Dictionary = map[string]string{
	"Absolute": `
.absolute {
  position: absolute;
}`,

	"AbsoluteLg": `
@media (min-width: 1024px) {
  .lg\:absolute {
    position: absolute;
  }
}`,

	"AbsoluteMd": `
@media (min-width: 768px) {
  .md\:absolute {
    position: absolute;
  }
}`,

	"AbsoluteSm": `
@media (min-width: 640px) {
  .sm\:absolute {
    position: absolute;
  }
}`,

	"AbsoluteX2l": `
@media (min-width: 1536px) {
  .2xl\:absolute {
    position: absolute;
  }
}`,

	"AbsoluteXl": `
@media (min-width: 1280px) {
  .xl\:absolute {
    position: absolute;
  }
}`,

	"AppearanceNone": `
.appearance-none {
  appearance: none;
}`,

	"AppearanceNoneLg": `
@media (min-width: 1024px) {
  .lg\:appearance-none {
    appearance: none;
  }
}`,

	"AppearanceNoneMd": `
@media (min-width: 768px) {
  .md\:appearance-none {
    appearance: none;
  }
}`,

	"AppearanceNoneSm": `
@media (min-width: 640px) {
  .sm\:appearance-none {
    appearance: none;
  }
}`,

	"AppearanceNoneX2l": `
@media (min-width: 1536px) {
  .2xl\:appearance-none {
    appearance: none;
  }
}`,

	"AppearanceNoneXl": `
@media (min-width: 1280px) {
  .xl\:appearance-none {
    appearance: none;
  }
}`,

	"BgBlack": `
.bg-black {
  background-color: #000;
}`,

	"BgBlackLg": `
@media (min-width: 1024px) {
  .lg\:bg-black {
    background-color: #000;
  }
}`,

	"BgBlackMd": `
@media (min-width: 768px) {
  .md\:bg-black {
    background-color: #000;
  }
}`,

	"BgBlackSm": `
@media (min-width: 640px) {
  .sm\:bg-black {
    background-color: #000;
  }
}`,

	"BgBlackX2l": `
@media (min-width: 1536px) {
  .2xl\:bg-black {
    background-color: #000;
  }
}`,

	"BgBlackXl": `
@media (min-width: 1280px) {
  .xl\:bg-black {
    background-color: #000;
  }
}`,

	"BgBlue500": `
.bg-blue-500 {
  --tw-bg-opacity: 1;
  background-color: rgb(59 130 246 / var(--tw-bg-opacity));
}`,

	"BgBlue500Lg": `
@media (min-width: 1024px) {
  .lg\:bg-blue-500 {
    --tw-bg-opacity: 1;
    background-color: rgb(59 130 246 / var(--tw-bg-opacity));
  }
}`,

	"BgBlue500Md": `
@media (min-width: 768px) {
  .md\:bg-blue-500 {
    --tw-bg-opacity: 1;
    background-color: rgb(59 130 246 / var(--tw-bg-opacity));
  }
}`,

	"BgBlue500Sm": `
@media (min-width: 640px) {
  .sm\:bg-blue-500 {
    --tw-bg-opacity: 1;
    background-color: rgb(59 130 246 / var(--tw-bg-opacity));
  }
}`,

	"BgBlue500X2l": `
@media (min-width: 1536px) {
  .2xl\:bg-blue-500 {
    --tw-bg-opacity: 1;
    background-color: rgb(59 130 246 / var(--tw-bg-opacity));
  }
}`,

	"BgBlue500Xl": `
@media (min-width: 1280px) {
  .xl\:bg-blue-500 {
    --tw-bg-opacity: 1;
    background-color: rgb(59 130 246 / var(--tw-bg-opacity));
  }
}`,

	"BgGray100": `
.bg-gray-100 {
  --tw-bg-opacity: 1;
  background-color: rgb(243 244 246 / var(--tw-bg-opacity));
}`,

	"BgGray100Lg": `
@media (min-width: 1024px) {
  .lg\:bg-gray-100 {
    --tw-bg-opacity: 1;
    background-color: rgb(243 244 246 / var(--tw-bg-opacity));
  }
}`,

	"BgGray100Md": `
@media (min-width: 768px) {
  .md\:bg-gray-100 {
    --tw-bg-opacity: 1;
    background-color: rgb(243 244 246 / var(--tw-bg-opacity));
  }
}`,

	"BgGray100Sm": `
@media (min-width: 640px) {
  .sm\:bg-gray-100 {
    --tw-bg-opacity: 1;
    background-color: rgb(243 244 246 / var(--tw-bg-opacity));
  }
}`,

	"BgGray100X2l": `
@media (min-width: 1536px) {
  .2xl\:bg-gray-100 {
    --tw-bg-opacity: 1;
    background-color: rgb(243 244 246 / var(--tw-bg-opacity));
  }
}`,

	"BgGray100Xl": `
@media (min-width: 1280px) {
  .xl\:bg-gray-100 {
    --tw-bg-opacity: 1;
    background-color: rgb(243 244 246 / var(--tw-bg-opacity));
  }
}`,

	"Block": `
.block {
  display: block;
}`,

	"BlockMd": `
@media (min-width: 768px) {
  .md\:block {
    display: block;
  }
}`,

	"BlockSm": `
@media (min-width: 640px) {
  .sm\:block {
    display: block;
  }
}`,

	"BlockX2l": `
@media (min-width: 1536px) {
  .2xl\:block {
    display: block;
  }
}`,

	"BlockLg": `
@media (min-width: 1024px) {
  .lg\:block {
    display: block;
  }
}`,

	"BlockXl": `
@media (min-width: 1280px) {
  .xl\:block {
    display: block;
  }
}`,

	"BlurNone": `
.blur-none {
  --tw-blur: none;
  filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
}`,

	"BlurNoneLg": `
@media (min-width: 1024px) {
  .lg\:blur-none {
    --tw-blur: none;
    filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
  }
}`,

	"BlurNoneMd": `
@media (min-width: 768px) {
  .md\:blur-none {
    --tw-blur: none;
    filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
  }
}`,

	"BlurNoneSm": `
@media (min-width: 640px) {
  .sm\:blur-none {
    --tw-blur: none;
    filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
  }
}`,

	"BlurNoneX2l": `
@media (min-width: 1536px) {
  .2xl\:blur-none {
    --tw-blur: none;
    filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
  }
}`,

	"BlurNoneXl": `
@media (min-width: 1280px) {
  .xl\:blur-none {
    --tw-blur: none;
    filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
  }
}`,

	"BlurSm": `
.blur-sm {
  --tw-blur: blur(4px);
  filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
}`,

	"BlurSmLg": `
@media (min-width: 1024px) {
  .lg\:blur-sm {
    --tw-blur: blur(4px);
    filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
  }
}`,

	"BlurSmMd": `
@media (min-width: 768px) {
  .md\:blur-sm {
    --tw-blur: blur(4px);
    filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
  }
}`,

	"BlurSmSm": `
@media (min-width: 640px) {
  .sm\:blur-sm {
    --tw-blur: blur(4px);
    filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
  }
}`,

	"BlurSmX2l": `
@media (min-width: 1536px) {
  .2xl\:blur-sm {
    --tw-blur: blur(4px);
    filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
  }
}`,

	"BlurSmXl": `
@media (min-width: 1280px) {
  .xl\:blur-sm {
    --tw-blur: blur(4px);
    filter: var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow);
  }
}`,

	"Border": `
.border {
  border-width: 1px;
}`,

	"Border2": `
.border-2 {
  border-width: 2px;
}`,

	"Border2Lg": `
@media (min-width: 1024px) {
  .lg\:border-2 {
    border-width: 2px;
  }
}`,

	"Border2Md": `
@media (min-width: 768px) {
  .md\:border-2 {
    border-width: 2px;
  }
}`,

	"Border2Sm": `
@media (min-width: 640px) {
  .sm\:border-2 {
    border-width: 2px;
  }
}`,

	"Border2X2l": `
@media (min-width: 1536px) {
  .2xl\:border-2 {
    border-width: 2px;
  }
}`,

	"Border2Xl": `
@media (min-width: 1280px) {
  .xl\:border-2 {
    border-width: 2px;
  }
}`,

	"BorderBlue500": `
.border-blue-500 {
  --tw-border-opacity: 1;
  border-color: rgb(59 130 246 / var(--tw-border-opacity));
}`,

	"BorderBlue500Lg": `
@media (min-width: 1024px) {
  .lg\:border-blue-500 {
    --tw-border-opacity: 1;
    border-color: rgb(59 130 246 / var(--tw-border-opacity));
  }
}`,

	"BorderBlue500Md": `
@media (min-width: 768px) {
  .md\:border-blue-500 {
    --tw-border-opacity: 1;
    border-color: rgb(59 130 246 / var(--tw-border-opacity));
  }
}`,

	"BorderBlue500Sm": `
@media (min-width: 640px) {
  .sm\:border-blue-500 {
    --tw-border-opacity: 1;
    border-color: rgb(59 130 246 / var(--tw-border-opacity));
  }
}`,

	"BorderBlue500X2l": `
@media (min-width: 1536px) {
  .2xl\:border-blue-500 {
    --tw-border-opacity: 1;
    border-color: rgb(59 130 246 / var(--tw-border-opacity));
  }
}`,

	"BorderBlue500Xl": `
@media (min-width: 1280px) {
  .xl\:border-blue-500 {
    --tw-border-opacity: 1;
    border-color: rgb(59 130 246 / var(--tw-border-opacity));
  }
}`,

	"BorderDashed": `
.border-dashed {
  border-style: dashed;
}`,

	"BorderDashedLg": `
@media (min-width: 1024px) {
  .lg\:border-dashed {
    border-style: dashed;
  }
}`,

	"BorderDashedMd": `
@media (min-width: 768px) {
  .md\:border-dashed {
    border-style: dashed;
  }
}`,

	"BorderDashedSm": `
@media (min-width: 640px) {
  .sm\:border-dashed {
    border-style: dashed;
  }
}`,

	"BorderDashedX2l": `
@media (min-width: 1536px) {
  .2xl\:border-dashed {
    border-style: dashed;
  }
}`,

	"BorderDashedXl": `
@media (min-width: 1280px) {
  .xl\:border-dashed {
    border-style: dashed;
  }
}`,

	"BorderDotted": `
.border-dotted {
  border-style: dotted;
}`,

	"BorderDottedLg": `
@media (min-width: 1024px) {
  .lg\:border-dotted {
    border-style: dotted;
  }
}`,

	"BorderDottedMd": `
@media (min-width: 768px) {
  .md\:border-dotted {
    border-style: dotted;
  }
}`,

	"BorderDottedSm": `
@media (min-width: 640px) {
  .sm\:border-dotted {
    border-style: dotted;
  }
}`,

	"BorderDottedX2l": `
@media (min-width: 1536px) {
  .2xl\:border-dotted {
    border-style: dotted;
  }
}`,

	"BorderDottedXl": `
@media (min-width: 1280px) {
  .xl\:border-dotted {
    border-style: dotted;
  }
}`,

	"BorderGray300": `
.border-gray-300 {
  --tw-border-opacity: 1;
  border-color: rgb(209 213 219 / var(--tw-border-opacity));
}`,

	"BorderGray300Lg": `
@media (min-width: 1024px) {
  .lg\:border-gray-300 {
    --tw-border-opacity: 1;
    border-color: rgb(209 213 219 / var(--tw-border-opacity));
  }
}`,

	"BorderGray300Md": `
@media (min-width: 768px) {
  .md\:border-gray-300 {
    --tw-border-opacity: 1;
    border-color: rgb(209 213 219 / var(--tw-border-opacity));
  }
}`,

	"BorderGray300Sm": `
@media (min-width: 640px) {
  .sm\:border-gray-300 {
    --tw-border-opacity: 1;
    border-color: rgb(209 213 219 / var(--tw-border-opacity));
  }
}`,

	"BorderGray300X2l": `
@media (min-width: 1536px) {
  .2xl\:border-gray-300 {
    --tw-border-opacity: 1;
    border-color: rgb(209 213 219 / var(--tw-border-opacity));
  }
}`,

	"BorderGray300Xl": `
@media (min-width: 1280px) {
  .xl\:border-gray-300 {
    --tw-border-opacity: 1;
    border-color: rgb(209 213 219 / var(--tw-border-opacity));
  }
}`,

	"BorderLg": `
@media (min-width: 1024px) {
  .lg\:border {
    border-width: 1px;
  }
}`,

	"BorderMd": `
@media (min-width: 768px) {
  .md\:border {
    border-width: 1px;
  }
}`,

	"BorderSm": `
@media (min-width: 640px) {
  .sm\:border {
    border-width: 1px;
  }
}`,

	"BorderX2l": `
@media (min-width: 1536px) {
  .2xl\:border {
    border-width: 1px;
  }
}`,

	"BorderXl": `
@media (min-width: 1280px) {
  .xl\:border {
    border-width: 1px;
  }
}`,

	"CursorDefault": `
.cursor-default {
  cursor: default;
}`,

	"CursorDefaultLg": `
@media (min-width: 1024px) {
  .lg\:cursor-default {
    cursor: default;
  }
}`,

	"CursorDefaultMd": `
@media (min-width: 768px) {
  .md\:cursor-default {
    cursor: default;
  }
}`,

	"CursorDefaultSm": `
@media (min-width: 640px) {
  .sm\:cursor-default {
    cursor: default;
  }
}`,

	"CursorDefaultX2l": `
@media (min-width: 1536px) {
  .2xl\:cursor-default {
    cursor: default;
  }
}`,

	"CursorDefaultXl": `
@media (min-width: 1280px) {
  .xl\:cursor-default {
    cursor: default;
  }
}`,

	"CursorPointer": `
.cursor-pointer {
  cursor: pointer;
}`,

	"CursorPointerLg": `
@media (min-width: 1024px) {
  .lg\:cursor-pointer {
    cursor: pointer;
  }
}`,

	"CursorPointerMd": `
@media (min-width: 768px) {
  .md\:cursor-pointer {
    cursor: pointer;
  }
}`,

	"CursorPointerSm": `
@media (min-width: 640px) {
  .sm\:cursor-pointer {
    cursor: pointer;
  }
}`,

	"CursorPointerX2l": `
@media (min-width: 1536px) {
  .2xl\:cursor-pointer {
    cursor: pointer;
  }
}`,

	"CursorPointerXl": `
@media (min-width: 1280px) {
  .xl\:cursor-pointer {
    cursor: pointer;
  }
}`,

	"Duration150": `
.duration-150 {
  transition-duration: 150ms;
}`,

	"Duration150Lg": `
@media (min-width: 1024px) {
  .lg\:duration-150 {
    transition-duration: 150ms;
  }
}`,

	"Duration150Md": `
@media (min-width: 768px) {
  .md\:duration-150 {
    transition-duration: 150ms;
  }
}`,

	"Duration150Sm": `
@media (min-width: 640px) {
  .sm\:duration-150 {
    transition-duration: 150ms;
  }
}`,

	"Duration150X2l": `
@media (min-width: 1536px) {
  .2xl\:duration-150 {
    transition-duration: 150ms;
  }
}`,

	"Duration150Xl": `
@media (min-width: 1280px) {
  .xl\:duration-150 {
    transition-duration: 150ms;
  }
}`,

	"Duration300": `
.duration-300 {
  transition-duration: 300ms;
}`,

	"Duration300Lg": `
@media (min-width: 1024px) {
  .lg\:duration-300 {
    transition-duration: 300ms;
  }
}`,

	"Duration300Md": `
@media (min-width: 768px) {
  .md\:duration-300 {
    transition-duration: 300ms;
  }
}`,

	"Duration300Sm": `
@media (min-width: 640px) {
  .sm\:duration-300 {
    transition-duration: 300ms;
  }
}`,

	"Duration300X2l": `
@media (min-width: 1536px) {
  .2xl\:duration-300 {
    transition-duration: 300ms;
  }
}`,

	"Duration300Xl": `
@media (min-width: 1280px) {
  .xl\:duration-300 {
    transition-duration: 300ms;
  }
}`,

	"EaseInOut": `
.ease-in-out {
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}`,

	"EaseInOutLg": `
@media (min-width: 1024px) {
  .lg\:ease-in-out {
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
}`,

	"EaseInOutMd": `
@media (min-width: 768px) {
  .md\:ease-in-out {
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
}`,

	"EaseInOutSm": `
@media (min-width: 640px) {
  .sm\:ease-in-out {
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
}`,

	"EaseInOutX2l": `
@media (min-width: 1536px) {
  .2xl\:ease-in-out {
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
}`,

	"EaseInOutXl": `
@media (min-width: 1280px) {
  .xl\:ease-in-out {
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
}`,

	"Flex": `
.flex {
  display: flex;
}`,

	"FlexCol": `
.flex-col {
  flex-direction: column;
}`,

	"FlexColLg": `
@media (min-width: 1024px) {
  .lg\:flex-col {
    flex-direction: column;
  }
}`,

	"FlexColMd": `
@media (min-width: 768px) {
  .md\:flex-col {
    flex-direction: column;
  }
}`,

	"FlexColSm": `
@media (min-width: 640px) {
  .sm\:flex-col {
    flex-direction: column;
  }
}`,

	"FlexColX2l": `
@media (min-width: 1536px) {
  .2xl\:flex-col {
    flex-direction: column;
  }
}`,

	"FlexColXl": `
@media (min-width: 1280px) {
  .xl\:flex-col {
    flex-direction: column;
  }
}`,

	"FlexLg": `
@media (min-width: 1024px) {
  .lg\:flex {
    display: flex;
  }
}`,

	"FlexMd": `
@media (min-width: 768px) {
  .md\:flex {
    display: flex;
  }
}`,

	"FlexRow": `
.flex-row {
  flex-direction: row;
}`,

	"FlexRowLg": `
@media (min-width: 1024px) {
  .lg\:flex-row {
    flex-direction: row;
  }
}`,

	"FlexRowMd": `
@media (min-width: 768px) {
  .md\:flex-row {
    flex-direction: row;
  }
}`,

	"FlexRowSm": `
@media (min-width: 640px) {
  .sm\:flex-row {
    flex-direction: row;
  }
}`,

	"FlexRowX2l": `
@media (min-width: 1536px) {
  .2xl\:flex-row {
    flex-direction: row;
  }
}`,

	"FlexRowXl": `
@media (min-width: 1280px) {
  .xl\:flex-row {
    flex-direction: row;
  }
}`,

	"FlexSm": `
@media (min-width: 640px) {
  .sm\:flex {
    display: flex;
  }
}`,

	"FlexX2l": `
@media (min-width: 1536px) {
  .2xl\:flex {
    display: flex;
  }
}`,

	"FlexXl": `
@media (min-width: 1280px) {
  .xl\:flex {
    display: flex;
  }
}`,

	"FontBold": `
.font-bold {
  font-weight: 700;
}`,

	"FontBoldLg": `
@media (min-width: 1024px) {
  .lg\:font-bold {
    font-weight: 700;
  }
}`,

	"FontBoldMd": `
@media (min-width: 768px) {
  .md\:font-bold {
    font-weight: 700;
  }
}`,

	"FontBoldSm": `
@media (min-width: 640px) {
  .sm\:font-bold {
    font-weight: 700;
  }
}`,

	"FontBoldX2l": `
@media (min-width: 1536px) {
  .2xl\:font-bold {
    font-weight: 700;
  }
}`,

	"FontBoldXl": `
@media (min-width: 1280px) {
  .xl\:font-bold {
    font-weight: 700;
  }
}`,

	"FontMedium": `
.font-medium {
  font-weight: 500;
}`,

	"FontMediumLg": `
@media (min-width: 1024px) {
  .lg\:font-medium {
    font-weight: 500;
  }
}`,

	"FontMediumMd": `
@media (min-width: 768px) {
  .md\:font-medium {
    font-weight: 500;
  }
}`,

	"FontMediumSm": `
@media (min-width: 640px) {
  .sm\:font-medium {
    font-weight: 500;
  }
}`,

	"FontMediumX2l": `
@media (min-width: 1536px) {
  .2xl\:font-medium {
    font-weight: 500;
  }
}`,

	"FontMediumXl": `
@media (min-width: 1280px) {
  .xl\:font-medium {
    font-weight: 500;
  }
}`,

	"FontNormal": `
.font-normal {
  font-weight: 400;
}`,

	"FontNormalLg": `
@media (min-width: 1024px) {
  .lg\:font-normal {
    font-weight: 400;
  }
}`,

	"FontNormalMd": `
@media (min-width: 768px) {
  .md\:font-normal {
    font-weight: 400;
  }
}`,

	"FontNormalSm": `
@media (min-width: 640px) {
  .sm\:font-normal {
    font-weight: 400;
  }
}`,

	"FontNormalX2l": `
@media (min-width: 1536px) {
  .2xl\:font-normal {
    font-weight: 400;
  }
}`,

	"FontNormalXl": `
@media (min-width: 1280px) {
  .xl\:font-normal {
    font-weight: 400;
  }
}`,

	"FontSemibold": `
.font-semibold {
  font-weight: 600;
}`,

	"FontSemiboldLg": `
@media (min-width: 1024px) {
  .lg\:font-semibold {
    font-weight: 600;
  }
}`,

	"FontSemiboldMd": `
@media (min-width: 768px) {
  .md\:font-semibold {
    font-weight: 600;
  }
}`,

	"FontSemiboldSm": `
@media (min-width: 640px) {
  .sm\:font-semibold {
    font-weight: 600;
  }
}`,

	"FontSemiboldX2l": `
@media (min-width: 1536px) {
  .2xl\:font-semibold {
    font-weight: 600;
  }
}`,

	"FontSemiboldXl": `
@media (min-width: 1280px) {
  .xl\:font-semibold {
    font-weight: 600;
  }
}`,

	"Gap1": `
.gap-1 {
  gap: 0.25rem;
}`,

	"Gap1Lg": `
@media (min-width: 1024px) {
  .lg\:gap-1 {
    gap: 0.25rem;
  }
}`,

	"Gap1Md": `
@media (min-width: 768px) {
  .md\:gap-1 {
    gap: 0.25rem;
  }
}`,

	"Gap1Sm": `
@media (min-width: 640px) {
  .sm\:gap-1 {
    gap: 0.25rem;
  }
}`,

	"Gap1X2l": `
@media (min-width: 1536px) {
  .2xl\:gap-1 {
    gap: 0.25rem;
  }
}`,

	"Gap1Xl": `
@media (min-width: 1280px) {
  .xl\:gap-1 {
    gap: 0.25rem;
  }
}`,

	"Gap2": `
.gap-2 {
  gap: 0.5rem;
}`,

	"Gap2Lg": `
@media (min-width: 1024px) {
  .lg\:gap-2 {
    gap: 0.5rem;
  }
}`,

	"Gap2Md": `
@media (min-width: 768px) {
  .md\:gap-2 {
    gap: 0.5rem;
  }
}`,

	"Gap2Sm": `
@media (min-width: 640px) {
  .sm\:gap-2 {
    gap: 0.5rem;
  }
}`,

	"Gap2X2l": `
@media (min-width: 1536px) {
  .2xl\:gap-2 {
    gap: 0.5rem;
  }
}`,

	"Gap2Xl": `
@media (min-width: 1280px) {
  .xl\:gap-2 {
    gap: 0.5rem;
  }
}`,

	"Gap4": `
.gap-4 {
  gap: 1rem;
}`,

	"Gap4Alt": `
.gap-4 {
  gap: 1rem;
}`,

	"Gap4AltLg": `
@media (min-width: 1024px) {
  .lg\:gap-4 {
    gap: 1rem;
  }
}`,

	"Gap4AltMd": `
@media (min-width: 768px) {
  .md\:gap-4 {
    gap: 1rem;
  }
}`,

	"Gap4AltSm": `
@media (min-width: 640px) {
  .sm\:gap-4 {
    gap: 1rem;
  }
}`,

	"Gap4AltX2l": `
@media (min-width: 1536px) {
  .2xl\:gap-4 {
    gap: 1rem;
  }
}`,

	"Gap4AltXl": `
@media (min-width: 1280px) {
  .xl\:gap-4 {
    gap: 1rem;
  }
}`,

	"Gap4Lg": `
@media (min-width: 1024px) {
  .lg\:gap-4 {
    gap: 1rem;
  }
}`,

	"Gap4Md": `
@media (min-width: 768px) {
  .md\:gap-4 {
    gap: 1rem;
  }
}`,

	"Gap4Sm": `
@media (min-width: 640px) {
  .sm\:gap-4 {
    gap: 1rem;
  }
}`,

	"Gap4X2l": `
@media (min-width: 1536px) {
  .2xl\:gap-4 {
    gap: 1rem;
  }
}`,

	"Gap4Xl": `
@media (min-width: 1280px) {
  .xl\:gap-4 {
    gap: 1rem;
  }
}`,

	"Grid": `
.grid {
  display: grid;
}`,

	"GridLg": `
@media (min-width: 1024px) {
  .lg\:grid {
    display: grid;
  }
}`,

	"GridMd": `
@media (min-width: 768px) {
  .md\:grid {
    display: grid;
  }
}`,

	"GridSm": `
@media (min-width: 640px) {
  .sm\:grid {
    display: grid;
  }
}`,

	"GridX2l": `
@media (min-width: 1536px) {
  .2xl\:grid {
    display: grid;
  }
}`,

	"GridXl": `
@media (min-width: 1280px) {
  .xl\:grid {
    display: grid;
  }
}`,

	"HFull": `
.h-full {
  height: 100%;
}`,

	"HFullLg": `
@media (min-width: 1024px) {
  .lg\:h-full {
    height: 100%;
  }
}`,

	"HFullMd": `
@media (min-width: 768px) {
  .md\:h-full {
    height: 100%;
  }
}`,

	"HFullSm": `
@media (min-width: 640px) {
  .sm\:h-full {
    height: 100%;
  }
}`,

	"HFullX2l": `
@media (min-width: 1536px) {
  .2xl\:h-full {
    height: 100%;
  }
}`,

	"HFullXl": `
@media (min-width: 1280px) {
  .xl\:h-full {
    height: 100%;
  }
}`,

	"Inline": `
.inline {
  display: inline;
}`,

	"InlineBlock": `
.inline-block {
  display: inline-block;
}`,

	"InlineBlockLg": `
@media (min-width: 1024px) {
  .lg\:inline-block {
    display: inline-block;
  }
}`,

	"InlineBlockMd": `
@media (min-width: 768px) {
  .md\:inline-block {
    display: inline-block;
  }
}`,

	"InlineBlockSm": `
@media (min-width: 640px) {
  .sm\:inline-block {
    display: inline-block;
  }
}`,

	"InlineBlockX2l": `
@media (min-width: 1536px) {
  .2xl\:inline-block {
    display: inline-block;
  }
}`,

	"InlineBlockXl": `
@media (min-width: 1280px) {
  .xl\:inline-block {
    display: inline-block;
  }
}`,

	"InlineLg": `
@media (min-width: 1024px) {
  .lg\:inline {
    display: inline;
  }
}`,

	"InlineMd": `
@media (min-width: 768px) {
  .md\:inline {
    display: inline;
  }
}`,

	"InlineSm": `
@media (min-width: 640px) {
  .sm\:inline {
    display: inline;
  }
}`,

	"InlineX2l": `
@media (min-width: 1536px) {
  .2xl\:inline {
    display: inline;
  }
}`,

	"InlineXl": `
@media (min-width: 1280px) {
  .xl\:inline {
    display: inline;
  }
}`,

	"Inset0": `
.inset-0 {
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
}`,

	"Inset0Lg": `
@media (min-width: 1024px) {
  .lg\:inset-0 {
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
  }
}`,

	"Inset0Md": `
@media (min-width: 768px) {
  .md\:inset-0 {
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
  }
}`,

	"Inset0Sm": `
@media (min-width: 640px) {
  .sm\:inset-0 {
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
  }
}`,

	"Inset0X2l": `
@media (min-width: 1536px) {
  .2xl\:inset-0 {
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
  }
}`,

	"Inset0Xl": `
@media (min-width: 1280px) {
  .xl\:inset-0 {
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
  }
}`,

	"InsetX0": `
.inset-x-0 {
  left: 0;
  right: 0;
}`,

	"InsetX0Lg": `
@media (min-width: 1024px) {
  .lg\:inset-x-0 {
    left: 0;
    right: 0;
  }
}`,

	"InsetX0Md": `
@media (min-width: 768px) {
  .md\:inset-x-0 {
    left: 0;
    right: 0;
  }
}`,

	"InsetX0Sm": `
@media (min-width: 640px) {
  .sm\:inset-x-0 {
    left: 0;
    right: 0;
  }
}`,

	"InsetX0X2l": `
@media (min-width: 1536px) {
  .2xl\:inset-x-0 {
    left: 0;
    right: 0;
  }
}`,

	"InsetX0Xl": `
@media (min-width: 1280px) {
  .xl\:inset-x-0 {
    left: 0;
    right: 0;
  }
}`,

	"ItemsCenter": `
.items-center {
  align-items: center;
}`,

	"ItemsCenterLg": `
@media (min-width: 1024px) {
  .lg\:items-center {
    align-items: center;
  }
}`,

	"ItemsCenterMd": `
@media (min-width: 768px) {
  .md\:items-center {
    align-items: center;
  }
}`,

	"ItemsCenterSm": `
@media (min-width: 640px) {
  .sm\:items-center {
    align-items: center;
  }
}`,

	"ItemsCenterX2l": `
@media (min-width: 1536px) {
  .2xl\:items-center {
    align-items: center;
  }
}`,

	"ItemsCenterXl": `
@media (min-width: 1280px) {
  .xl\:items-center {
    align-items: center;
  }
}`,

	"ItemsEnd": `
.items-end {
  align-items: flex-end;
}`,

	"ItemsEndLg": `
@media (min-width: 1024px) {
  .lg\:items-end {
    align-items: flex-end;
  }
}`,

	"ItemsEndMd": `
@media (min-width: 768px) {
  .md\:items-end {
    align-items: flex-end;
  }
}`,

	"ItemsEndSm": `
@media (min-width: 640px) {
  .sm\:items-end {
    align-items: flex-end;
  }
}`,

	"ItemsEndX2l": `
@media (min-width: 1536px) {
  .2xl\:items-end {
    align-items: flex-end;
  }
}`,

	"ItemsEndXl": `
@media (min-width: 1280px) {
  .xl\:items-end {
    align-items: flex-end;
  }
}`,

	"ItemsStart": `
.items-start {
  align-items: flex-start;
}`,

	"ItemsStartLg": `
@media (min-width: 1024px) {
  .lg\:items-start {
    align-items: flex-start;
  }
}`,

	"ItemsStartMd": `
@media (min-width: 768px) {
  .md\:items-start {
    align-items: flex-start;
  }
}`,

	"ItemsStartSm": `
@media (min-width: 640px) {
  .sm\:items-start {
    align-items: flex-start;
  }
}`,

	"ItemsStartX2l": `
@media (min-width: 1536px) {
  .2xl\:items-start {
    align-items: flex-start;
  }
}`,

	"ItemsStartXl": `
@media (min-width: 1280px) {
  .xl\:items-start {
    align-items: flex-start;
  }
}`,

	"JustifyAround": `
.justify-around {
  justify-content: space-around;
}`,

	"JustifyAroundLg": `
@media (min-width: 1024px) {
  .lg\:justify-around {
    justify-content: space-around;
  }
}`,

	"JustifyAroundMd": `
@media (min-width: 768px) {
  .md\:justify-around {
    justify-content: space-around;
  }
}`,

	"JustifyAroundSm": `
@media (min-width: 640px) {
  .sm\:justify-around {
    justify-content: space-around;
  }
}`,

	"JustifyAroundX2l": `
@media (min-width: 1536px) {
  .2xl\:justify-around {
    justify-content: space-around;
  }
}`,

	"JustifyAroundXl": `
@media (min-width: 1280px) {
  .xl\:justify-around {
    justify-content: space-around;
  }
}`,

	"JustifyBetween": `
.justify-between {
  justify-content: space-between;
}`,

	"JustifyBetweenLg": `
@media (min-width: 1024px) {
  .lg\:justify-between {
    justify-content: space-between;
  }
}`,

	"JustifyBetweenMd": `
@media (min-width: 768px) {
  .md\:justify-between {
    justify-content: space-between;
  }
}`,

	"JustifyBetweenSm": `
@media (min-width: 640px) {
  .sm\:justify-between {
    justify-content: space-between;
  }
}`,

	"JustifyBetweenX2l": `
@media (min-width: 1536px) {
  .2xl\:justify-between {
    justify-content: space-between;
  }
}`,

	"JustifyBetweenXl": `
@media (min-width: 1280px) {
  .xl\:justify-between {
    justify-content: space-between;
  }
}`,

	"JustifyCenter": `
.justify-center {
  justify-content: center;
}`,

	"JustifyCenterLg": `
@media (min-width: 1024px) {
  .lg\:justify-center {
    justify-content: center;
  }
}`,

	"JustifyCenterMd": `
@media (min-width: 768px) {
  .md\:justify-center {
    justify-content: center;
  }
}`,

	"JustifyCenterSm": `
@media (min-width: 640px) {
  .sm\:justify-center {
    justify-content: center;
  }
}`,

	"JustifyCenterX2l": `
@media (min-width: 1536px) {
  .2xl\:justify-center {
    justify-content: center;
  }
}`,

	"JustifyCenterXl": `
@media (min-width: 1280px) {
  .xl\:justify-center {
    justify-content: center;
  }
}`,

	"LeadingNormal": `
.leading-normal {
  line-height: 1.5;
}`,

	"LeadingNormalLg": `
@media (min-width: 1024px) {
  .lg\:leading-normal {
    line-height: 1.5;
  }
}`,

	"LeadingNormalMd": `
@media (min-width: 768px) {
  .md\:leading-normal {
    line-height: 1.5;
  }
}`,

	"LeadingNormalSm": `
@media (min-width: 640px) {
  .sm\:leading-normal {
    line-height: 1.5;
  }
}`,

	"LeadingNormalX2l": `
@media (min-width: 1536px) {
  .2xl\:leading-normal {
    line-height: 1.5;
  }
}`,

	"LeadingNormalXl": `
@media (min-width: 1280px) {
  .xl\:leading-normal {
    line-height: 1.5;
  }
}`,

	"LeadingSnug": `
.leading-snug {
  line-height: 1.375;
}`,

	"LeadingSnugLg": `
@media (min-width: 1024px) {
  .lg\:leading-snug {
    line-height: 1.375;
  }
}`,

	"LeadingSnugMd": `
@media (min-width: 768px) {
  .md\:leading-snug {
    line-height: 1.375;
  }
}`,

	"LeadingSnugSm": `
@media (min-width: 640px) {
  .sm\:leading-snug {
    line-height: 1.375;
  }
}`,

	"LeadingSnugX2l": `
@media (min-width: 1536px) {
  .2xl\:leading-snug {
    line-height: 1.375;
  }
}`,

	"LeadingSnugXl": `
@media (min-width: 1280px) {
  .xl\:leading-snug {
    line-height: 1.375;
  }
}`,

	"M2": `
.m-2 {
  margin: 0.5rem;
}`,

	"M2Lg": `
@media (min-width: 1024px) {
  .lg\:m-2 {
    margin: 0.5rem;
  }
}`,

	"M2Md": `
@media (min-width: 768px) {
  .md\:m-2 {
    margin: 0.5rem;
  }
}`,

	"M2Sm": `
@media (min-width: 640px) {
  .sm\:m-2 {
    margin: 0.5rem;
  }
}`,

	"M2X2l": `
@media (min-width: 1536px) {
  .2xl\:m-2 {
    margin: 0.5rem;
  }
}`,

	"M2Xl": `
@media (min-width: 1280px) {
  .xl\:m-2 {
    margin: 0.5rem;
  }
}`,

	"M4": `
.m-4 {
  margin: 1rem;
}`,

	"M4Lg": `
@media (min-width: 1024px) {
  .lg\:m-4 {
    margin: 1rem;
  }
}`,

	"M4Md": `
@media (min-width: 768px) {
  .md\:m-4 {
    margin: 1rem;
  }
}`,

	"M4Sm": `
@media (min-width: 640px) {
  .sm\:m-4 {
    margin: 1rem;
  }
}`,

	"M4X2l": `
@media (min-width: 1536px) {
  .2xl\:m-4 {
    margin: 1rem;
  }
}`,

	"M4Xl": `
@media (min-width: 1280px) {
  .xl\:m-4 {
    margin: 1rem;
  }
}`,

	"MaxWFull": `
.max-w-full {
  max-width: 100%;
}`,

	"MaxWFullLg": `
@media (min-width: 1024px) {
  .lg\:max-w-full {
    max-width: 100%;
  }
}`,

	"MaxWFullMd": `
@media (min-width: 768px) {
  .md\:max-w-full {
    max-width: 100%;
  }
}`,

	"MaxWFullSm": `
@media (min-width: 640px) {
  .sm\:max-w-full {
    max-width: 100%;
  }
}`,

	"MaxWFullX2l": `
@media (min-width: 1536px) {
  .2xl\:max-w-full {
    max-width: 100%;
  }
}`,

	"MaxWFullXl": `
@media (min-width: 1280px) {
  .xl\:max-w-full {
    max-width: 100%;
  }
}`,

	"Mb4": `
.mb-4 {
  margin-bottom: 1rem;
}`,

	"Mb4Lg": `
@media (min-width: 1024px) {
  .lg\:mb-4 {
    margin-bottom: 1rem;
  }
}`,

	"Mb4Md": `
@media (min-width: 768px) {
  .md\:mb-4 {
    margin-bottom: 1rem;
  }
}`,

	"Mb4Sm": `
@media (min-width: 640px) {
  .sm\:mb-4 {
    margin-bottom: 1rem;
  }
}`,

	"Mb4X2l": `
@media (min-width: 1536px) {
  .2xl\:mb-4 {
    margin-bottom: 1rem;
  }
}`,

	"Mb4Xl": `
@media (min-width: 1280px) {
  .xl\:mb-4 {
    margin-bottom: 1rem;
  }
}`,

	"Ml4": `
.ml-4 {
  margin-left: 1rem;
}`,

	"Ml4Lg": `
@media (min-width: 1024px) {
  .lg\:ml-4 {
    margin-left: 1rem;
  }
}`,

	"Ml4Md": `
@media (min-width: 768px) {
  .md\:ml-4 {
    margin-left: 1rem;
  }
}`,

	"Ml4Sm": `
@media (min-width: 640px) {
  .sm\:ml-4 {
    margin-left: 1rem;
  }
}`,

	"Ml4X2l": `
@media (min-width: 1536px) {
  .2xl\:ml-4 {
    margin-left: 1rem;
  }
}`,

	"Ml4Xl": `
@media (min-width: 1280px) {
  .xl\:ml-4 {
    margin-left: 1rem;
  }
}`,

	"Mr4": `
.mr-4 {
  margin-right: 1rem;
}`,

	"Mr4Lg": `
@media (min-width: 1024px) {
  .lg\:mr-4 {
    margin-right: 1rem;
  }
}`,

	"Mr4Md": `
@media (min-width: 768px) {
  .md\:mr-4 {
    margin-right: 1rem;
  }
}`,

	"Mr4Sm": `
@media (min-width: 640px) {
  .sm\:mr-4 {
    margin-right: 1rem;
  }
}`,

	"Mr4X2l": `
@media (min-width: 1536px) {
  .2xl\:mr-4 {
    margin-right: 1rem;
  }
}`,

	"Mr4Xl": `
@media (min-width: 1280px) {
  .xl\:mr-4 {
    margin-right: 1rem;
  }
}`,

	"Mt4": `
.mt-4 {
  margin-top: 1rem;
}`,

	"Mt4Lg": `
@media (min-width: 1024px) {
  .lg\:mt-4 {
    margin-top: 1rem;
  }
}`,

	"Mt4Md": `
@media (min-width: 768px) {
  .md\:mt-4 {
    margin-top: 1rem;
  }
}`,

	"Mt4Sm": `
@media (min-width: 640px) {
  .sm\:mt-4 {
    margin-top: 1rem;
  }
}`,

	"Mt4X2l": `
@media (min-width: 1536px) {
  .2xl\:mt-4 {
    margin-top: 1rem;
  }
}`,

	"Mt4Xl": `
@media (min-width: 1280px) {
  .xl\:mt-4 {
    margin-top: 1rem;
  }
}`,

	"Opacity50": `
.opacity-50 {
  opacity: 0.5;
}`,

	"Opacity50Lg": `
@media (min-width: 1024px) {
  .lg\:opacity-50 {
    opacity: 0.5;
  }
}`,

	"Opacity50Md": `
@media (min-width: 768px) {
  .md\:opacity-50 {
    opacity: 0.5;
  }
}`,

	"Opacity50Sm": `
@media (min-width: 640px) {
  .sm\:opacity-50 {
    opacity: 0.5;
  }
}`,

	"Opacity50X2l": `
@media (min-width: 1536px) {
  .2xl\:opacity-50 {
    opacity: 0.5;
  }
}`,

	"Opacity50Xl": `
@media (min-width: 1280px) {
  .xl\:opacity-50 {
    opacity: 0.5;
  }
}`,

	"Opacity75": `
.opacity-75 {
  opacity: 0.75;
}`,

	"Opacity75Lg": `
@media (min-width: 1024px) {
  .lg\:opacity-75 {
    opacity: 0.75;
  }
}`,

	"Opacity75Md": `
@media (min-width: 768px) {
  .md\:opacity-75 {
    opacity: 0.75;
  }
}`,

	"Opacity75Sm": `
@media (min-width: 640px) {
  .sm\:opacity-75 {
    opacity: 0.75;
  }
}`,

	"Opacity75X2l": `
@media (min-width: 1536px) {
  .2xl\:opacity-75 {
    opacity: 0.75;
  }
}`,

	"Opacity75Xl": `
@media (min-width: 1280px) {
  .xl\:opacity-75 {
    opacity: 0.75;
  }
}`,

	"OutlineNone": `
.outline-none {
  outline: 2px solid transparent;
  outline-offset: 2px;
}`,

	"OutlineNoneLg": `
@media (min-width: 1024px) {
  .lg\:outline-none {
    outline: 2px solid transparent;
    outline-offset: 2px;
  }
}`,

	"OutlineNoneMd": `
@media (min-width: 768px) {
  .md\:outline-none {
    outline: 2px solid transparent;
    outline-offset: 2px;
  }
}`,

	"OutlineNoneSm": `
@media (min-width: 640px) {
  .sm\:outline-none {
    outline: 2px solid transparent;
    outline-offset: 2px;
  }
}`,

	"OutlineNoneX2l": `
@media (min-width: 1536px) {
  .2xl\:outline-none {
    outline: 2px solid transparent;
    outline-offset: 2px;
  }
}`,

	"OutlineNoneXl": `
@media (min-width: 1280px) {
  .xl\:outline-none {
    outline: 2px solid transparent;
    outline-offset: 2px;
  }
}`,

	"OverflowHidden": `
.overflow-hidden {
  overflow: hidden;
}`,

	"OverflowHiddenLg": `
@media (min-width: 1024px) {
  .lg\:overflow-hidden {
    overflow: hidden;
  }
}`,

	"OverflowHiddenMd": `
@media (min-width: 768px) {
  .md\:overflow-hidden {
    overflow: hidden;
  }
}`,

	"OverflowHiddenSm": `
@media (min-width: 640px) {
  .sm\:overflow-hidden {
    overflow: hidden;
  }
}`,

	"OverflowHiddenX2l": `
@media (min-width: 1536px) {
  .2xl\:overflow-hidden {
    overflow: hidden;
  }
}`,

	"OverflowHiddenXl": `
@media (min-width: 1280px) {
  .xl\:overflow-hidden {
    overflow: hidden;
  }
}`,

	"P2": `
.p-2 {
  padding: 0.5rem;
}`,

	"P2Lg": `
@media (min-width: 1024px) {
  .lg\:p-2 {
    padding: 0.5rem;
  }
}`,

	"P2Md": `
@media (min-width: 768px) {
  .md\:p-2 {
    padding: 0.5rem;
  }
}`,

	"P2Sm": `
@media (min-width: 640px) {
  .sm\:p-2 {
    padding: 0.5rem;
  }
}`,

	"P2X2l": `
@media (min-width: 1536px) {
  .2xl\:p-2 {
    padding: 0.5rem;
  }
}`,

	"P2Xl": `
@media (min-width: 1280px) {
  .xl\:p-2 {
    padding: 0.5rem;
  }
}`,

	"P4": `
.p-4 {
  padding: 1rem;
}`,

	"P4Lg": `
@media (min-width: 1024px) {
  .lg\:p-4 {
    padding: 1rem;
  }
}`,

	"P4Md": `
@media (min-width: 768px) {
  .md\:p-4 {
    padding: 1rem;
  }
}`,

	"P4Sm": `
@media (min-width: 640px) {
  .sm\:p-4 {
    padding: 1rem;
  }
}`,

	"P4X2l": `
@media (min-width: 1536px) {
  .2xl\:p-4 {
    padding: 1rem;
  }
}`,

	"P4Xl": `
@media (min-width: 1280px) {
  .xl\:p-4 {
    padding: 1rem;
  }
}`,

	"Pl4": `
.pl-4 {
  padding-left: 1rem;
}`,

	"Pl4Lg": `
@media (min-width: 1024px) {
  .lg\:pl-4 {
    padding-left: 1rem;
  }
}`,

	"Pl4Md": `
@media (min-width: 768px) {
  .md\:pl-4 {
    padding-left: 1rem;
  }
}`,

	"Pl4Sm": `
@media (min-width: 640px) {
  .sm\:pl-4 {
    padding-left: 1rem;
  }
}`,

	"Pl4X2l": `
@media (min-width: 1536px) {
  .2xl\:pl-4 {
    padding-left: 1rem;
  }
}`,

	"Pl4Xl": `
@media (min-width: 1280px) {
  .xl\:pl-4 {
    padding-left: 1rem;
  }
}`,

	"Pr4": `
.pr-4 {
  padding-right: 1rem;
}`,

	"Pr4Lg": `
@media (min-width: 1024px) {
  .lg\:pr-4 {
    padding-right: 1rem;
  }
}`,

	"Pr4Md": `
@media (min-width: 768px) {
  .md\:pr-4 {
    padding-right: 1rem;
  }
}`,

	"Pr4Sm": `
@media (min-width: 640px) {
  .sm\:pr-4 {
    padding-right: 1rem;
  }
}`,

	"Pr4X2l": `
@media (min-width: 1536px) {
  .2xl\:pr-4 {
    padding-right: 1rem;
  }
}`,

	"Pr4Xl": `
@media (min-width: 1280px) {
  .xl\:pr-4 {
    padding-right: 1rem;
  }
}`,

	"Px4": `
.px-4 {
  padding-left: 1rem;
  padding-right: 1rem;
}`,

	"Px4Lg": `
@media (min-width: 1024px) {
  .lg\:px-4 {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}`,

	"Px4Md": `
@media (min-width: 768px) {
  .md\:px-4 {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}`,

	"Px4Sm": `
@media (min-width: 640px) {
  .sm\:px-4 {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}`,

	"Px4X2l": `
@media (min-width: 1536px) {
  .2xl\:px-4 {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}`,

	"Px4Xl": `
@media (min-width: 1280px) {
  .xl\:px-4 {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}`,

	"Py2": `
.py-2 {
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
}`,

	"Py2Lg": `
@media (min-width: 1024px) {
  .lg\:py-2 {
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
  }
}`,

	"Py2Md": `
@media (min-width: 768px) {
  .md\:py-2 {
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
  }
}`,

	"Py2Sm": `
@media (min-width: 640px) {
  .sm\:py-2 {
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
  }
}`,

	"Py2X2l": `
@media (min-width: 1536px) {
  .2xl\:py-2 {
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
  }
}`,

	"Py2Xl": `
@media (min-width: 1280px) {
  .xl\:py-2 {
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
  }
}`,

	"Relative": `
.relative {
  position: relative;
}`,

	"RelativeLg": `
@media (min-width: 1024px) {
  .lg\:relative {
    position: relative;
  }
}`,

	"RelativeMd": `
@media (min-width: 768px) {
  .md\:relative {
    position: relative;
  }
}`,

	"RelativeSm": `
@media (min-width: 640px) {
  .sm\:relative {
    position: relative;
  }
}`,

	"RelativeX2l": `
@media (min-width: 1536px) {
  .2xl\:relative {
    position: relative;
  }
}`,

	"RelativeXl": `
@media (min-width: 1280px) {
  .xl\:relative {
    position: relative;
  }
}`,

	"Rounded": `
.rounded {
  border-radius: 0.25rem;
}`,

	"RoundedFull": `
.rounded-full {
  border-radius: 9999px;
}`,

	"RoundedFullLg": `
@media (min-width: 1024px) {
  .lg\:rounded-full {
    border-radius: 9999px;
  }
}`,

	"RoundedFullMd": `
@media (min-width: 768px) {
  .md\:rounded-full {
    border-radius: 9999px;
  }
}`,

	"RoundedFullSm": `
@media (min-width: 640px) {
  .sm\:rounded-full {
    border-radius: 9999px;
  }
}`,

	"RoundedFullX2l": `
@media (min-width: 1536px) {
  .2xl\:rounded-full {
    border-radius: 9999px;
  }
}`,

	"RoundedFullXl": `
@media (min-width: 1280px) {
  .xl\:rounded-full {
    border-radius: 9999px;
  }
}`,

	"RoundedLg": `
.rounded-lg {
  border-radius: 0.5rem;
}`,

	"RoundedLgSize": `
.rounded-lg {
  border-radius: 0.5rem;
}`,

	"RoundedLgSizeLg": `
@media (min-width: 1024px) {
  .lg\:rounded-lg {
    border-radius: 0.5rem;
  }
}`,

	"RoundedLgSizeMd": `
@media (min-width: 768px) {
  .md\:rounded-lg {
    border-radius: 0.5rem;
  }
}`,

	"RoundedLgSizeSm": `
@media (min-width: 640px) {
  .sm\:rounded-lg {
    border-radius: 0.5rem;
  }
}`,

	"RoundedLgSizeX2l": `
@media (min-width: 1536px) {
  .2xl\:rounded-lg {
    border-radius: 0.5rem;
  }
}`,

	"RoundedLgSizeXl": `
@media (min-width: 1280px) {
  .xl\:rounded-lg {
    border-radius: 0.5rem;
  }
}`,

	"RoundedMd": `
.rounded-md {
  border-radius: 0.375rem;
}`,

	"RoundedMdSize": `
.rounded-md {
  border-radius: 0.375rem;
}`,

	"RoundedMdSizeLg": `
@media (min-width: 1024px) {
  .lg\:rounded-md {
    border-radius: 0.375rem;
  }
}`,

	"RoundedMdSizeMd": `
@media (min-width: 768px) {
  .md\:rounded-md {
    border-radius: 0.375rem;
  }
}`,

	"RoundedMdSizeSm": `
@media (min-width: 640px) {
  .sm\:rounded-md {
    border-radius: 0.375rem;
  }
}`,

	"RoundedMdSizeX2l": `
@media (min-width: 1536px) {
  .2xl\:rounded-md {
    border-radius: 0.375rem;
  }
}`,

	"RoundedMdSizeXl": `
@media (min-width: 1280px) {
  .xl\:rounded-md {
    border-radius: 0.375rem;
  }
}`,

	"RoundedSm": `
.rounded-sm {
  border-radius: 0.125rem;
}`,

	"RoundedX2l": `
@media (min-width: 1536px) {
  .2xl\:rounded {
    border-radius: 0.25rem;
  }
}`,

	"RoundedXl": `
.rounded-xl {
  border-radius: 0.75rem;
}`,

	"SelectNone": `
.select-none {
  user-select: none;
}`,

	"SelectNoneLg": `
@media (min-width: 1024px) {
  .lg\:select-none {
    user-select: none;
  }
}`,

	"SelectNoneMd": `
@media (min-width: 768px) {
  .md\:select-none {
    user-select: none;
  }
}`,

	"SelectNoneSm": `
@media (min-width: 640px) {
  .sm\:select-none {
    user-select: none;
  }
}`,

	"SelectNoneX2l": `
@media (min-width: 1536px) {
  .2xl\:select-none {
    user-select: none;
  }
}`,

	"SelectNoneXl": `
@media (min-width: 1280px) {
  .xl\:select-none {
    user-select: none;
  }
}`,

	"SelectText": `
.select-text {
  user-select: text;
}`,

	"SelectTextLg": `
@media (min-width: 1024px) {
  .lg\:select-text {
    user-select: text;
  }
}`,

	"SelectTextMd": `
@media (min-width: 768px) {
  .md\:select-text {
    user-select: text;
  }
}`,

	"SelectTextSm": `
@media (min-width: 640px) {
  .sm\:select-text {
    user-select: text;
  }
}`,

	"SelectTextX2l": `
@media (min-width: 1536px) {
  .2xl\:select-text {
    user-select: text;
  }
}`,

	"SelectTextXl": `
@media (min-width: 1280px) {
  .xl\:select-text {
    user-select: text;
  }
}`,

	"Shadow": `
.shadow {
  --tw-shadow: 0 1px 3px 0 rgba(0,0,0,0.1), 0 1px 2px 0 rgba(0,0,0,0.06);
  box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
}`,

	"ShadowLg": `
@media (min-width: 1024px) {
  .lg\:shadow {
    --tw-shadow: 0 10px 15px -3px rgba(0,0,0,0.1), 0 4px 6px -2px rgba(0,0,0,0.05);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
  }
}`,

	"ShadowMd": `
@media (min-width: 768px) {
  .md\:shadow {
    --tw-shadow: 0 4px 6px -1px rgba(0,0,0,0.1), 0 2px 4px -1px rgba(0,0,0,0.06);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
  }
}`,

	"ShadowMdSize": `
.shadow-md {
  --tw-shadow: 0 4px 6px -1px rgba(0,0,0,0.1), 0 2px 4px -1px rgba(0,0,0,0.06);
  box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
}`,

	"ShadowMdSizeLg": `
@media (min-width: 1024px) {
  .lg\:shadow-md {
    --tw-shadow: 0 4px 6px -1px rgba(0,0,0,0.1), 0 2px 4px -1px rgba(0,0,0,0.06);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
  }
}`,

	"ShadowMdSizeMd": `
@media (min-width: 768px) {
  .md\:shadow-md {
    --tw-shadow: 0 4px 6px -1px rgba(0,0,0,0.1), 0 2px 4px -1px rgba(0,0,0,0.06);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
  }
}`,

	"ShadowMdSizeSm": `
@media (min-width: 640px) {
  .sm\:shadow-md {
    --tw-shadow: 0 4px 6px -1px rgba(0,0,0,0.1), 0 2px 4px -1px rgba(0,0,0,0.06);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
  }
}`,

	"ShadowMdSizeX2l": `
@media (min-width: 1536px) {
  .2xl\:shadow-md {
    --tw-shadow: 0 4px 6px -1px rgba(0,0,0,0.1), 0 2px 4px -1px rgba(0,0,0,0.06);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
  }
}`,

	"ShadowMdSizeXl": `
@media (min-width: 1280px) {
  .xl\:shadow-md {
    --tw-shadow: 0 4px 6px -1px rgba(0,0,0,0.1), 0 2px 4px -1px rgba(0,0,0,0.06);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
  }
}`,

	"ShadowSm": `
.shadow-sm {
  --tw-shadow: 0 1px 2px 0 rgba(0,0,0,0.05);
  box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
}`,

	"ShadowX2l": `
@media (min-width: 1536px) {
  .2xl\:shadow {
    --tw-shadow: 0 25px 50px -12px rgba(0,0,0,0.25);
    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
  }
}`,

	"ShadowXl": `
.shadow-xl {
  --tw-shadow: 0 20px 25px -5px rgba(0,0,0,0.1), 0 10px 10px -5px rgba(0,0,0,0.04);
  box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow);
}`,

	"SpaceX2": `
.space-x-2 > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 0;
  margin-right: calc(0.5rem * var(--tw-space-x-reverse));
  margin-left: calc(0.5rem * calc(1 - var(--tw-space-x-reverse)));
}`,

	"SpaceX2Lg": `
@media (min-width: 1024px) {
  .lg\:space-x-2 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(0.5rem * var(--tw-space-x-reverse));
    margin-left: calc(0.5rem * calc(1 - var(--tw-space-x-reverse)));
  }
}`,

	"SpaceX2Md": `
@media (min-width: 768px) {
  .md\:space-x-2 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(0.5rem * var(--tw-space-x-reverse));
    margin-left: calc(0.5rem * calc(1 - var(--tw-space-x-reverse)));
  }
}`,

	"SpaceX2Sm": `
@media (min-width: 640px) {
  .sm\:space-x-2 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(0.5rem * var(--tw-space-x-reverse));
    margin-left: calc(0.5rem * calc(1 - var(--tw-space-x-reverse)));
  }
}`,

	"SpaceX2X2l": `
@media (min-width: 1536px) {
  .2xl\:space-x-2 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(0.5rem * var(--tw-space-x-reverse));
    margin-left: calc(0.5rem * calc(1 - var(--tw-space-x-reverse)));
  }
}`,

	"SpaceX2Xl": `
@media (min-width: 1280px) {
  .xl\:space-x-2 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(0.5rem * var(--tw-space-x-reverse));
    margin-left: calc(0.5rem * calc(1 - var(--tw-space-x-reverse)));
  }
}`,

	"SpaceX4": `
.space-x-4 > :not([hidden]) ~ :not([hidden]) {
  --tw-space-x-reverse: 0;
  margin-right: calc(1rem * var(--tw-space-x-reverse));
  margin-left: calc(1rem * calc(1 - var(--tw-space-x-reverse)));
}`,

	"SpaceX4Lg": `
@media (min-width: 1024px) {
  .lg\:space-x-4 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(1rem * var(--tw-space-x-reverse));
    margin-left: calc(1rem * calc(1 - var(--tw-space-x-reverse)));
  }
}`,

	"SpaceX4Md": `
@media (min-width: 768px) {
  .md\:space-x-4 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(1rem * var(--tw-space-x-reverse));
    margin-left: calc(1rem * calc(1 - var(--tw-space-x-reverse)));
  }
}`,

	"SpaceX4Sm": `
@media (min-width: 640px) {
  .sm\:space-x-4 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(1rem * var(--tw-space-x-reverse));
    margin-left: calc(1rem * calc(1 - var(--tw-space-x-reverse)));
  }
}`,

	"SpaceX4X2l": `
@media (min-width: 1536px) {
  .2xl\:space-x-4 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(1rem * var(--tw-space-x-reverse));
    margin-left: calc(1rem * calc(1 - var(--tw-space-x-reverse)));
  }
}`,

	"SpaceX4Xl": `
@media (min-width: 1280px) {
  .xl\:space-x-4 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-x-reverse: 0;
    margin-right: calc(1rem * var(--tw-space-x-reverse));
    margin-left: calc(1rem * calc(1 - var(--tw-space-x-reverse)));
  }
}`,

	"SpaceY2": `
.space-y-2 > :not([hidden]) ~ :not([hidden]) {
  --tw-space-y-reverse: 0;
  margin-top: calc(0.5rem * calc(1 - var(--tw-space-y-reverse)));
  margin-bottom: calc(0.5rem * var(--tw-space-y-reverse));
}`,

	"SpaceY2Lg": `
@media (min-width: 1024px) {
  .lg\:space-y-2 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-y-reverse: 0;
    margin-top: calc(0.5rem * calc(1 - var(--tw-space-y-reverse)));
    margin-bottom: calc(0.5rem * var(--tw-space-y-reverse));
  }
}`,

	"SpaceY2Md": `
@media (min-width: 768px) {
  .md\:space-y-2 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-y-reverse: 0;
    margin-top: calc(0.5rem * calc(1 - var(--tw-space-y-reverse)));
    margin-bottom: calc(0.5rem * var(--tw-space-y-reverse));
  }
}`,

	"SpaceY2Sm": `
@media (min-width: 640px) {
  .sm\:space-y-2 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-y-reverse: 0;
    margin-top: calc(0.5rem * calc(1 - var(--tw-space-y-reverse)));
    margin-bottom: calc(0.5rem * var(--tw-space-y-reverse));
  }
}`,

	"SpaceY2X2l": `
@media (min-width: 1536px) {
  .2xl\:space-y-2 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-y-reverse: 0;
    margin-top: calc(0.5rem * calc(1 - var(--tw-space-y-reverse)));
    margin-bottom: calc(0.5rem * var(--tw-space-y-reverse));
  }
}`,

	"SpaceY2Xl": `
@media (min-width: 1280px) {
  .xl\:space-y-2 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-y-reverse: 0;
    margin-top: calc(0.5rem * calc(1 - var(--tw-space-y-reverse)));
    margin-bottom: calc(0.5rem * var(--tw-space-y-reverse));
  }
}`,

	"SpaceY4": `
.space-y-4 > :not([hidden]) ~ :not([hidden]) {
  --tw-space-y-reverse: 0;
  margin-top: calc(1rem * calc(1 - var(--tw-space-y-reverse)));
  margin-bottom: calc(1rem * var(--tw-space-y-reverse));
}`,

	"SpaceY4Lg": `
@media (min-width: 1024px) {
  .lg\:space-y-4 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-y-reverse: 0;
    margin-top: calc(1rem * calc(1 - var(--tw-space-y-reverse)));
    margin-bottom: calc(1rem * var(--tw-space-y-reverse));
  }
}`,

	"SpaceY4Md": `
@media (min-width: 768px) {
  .md\:space-y-4 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-y-reverse: 0;
    margin-top: calc(1rem * calc(1 - var(--tw-space-y-reverse)));
    margin-bottom: calc(1rem * var(--tw-space-y-reverse));
  }
}`,

	"SpaceY4Sm": `
@media (min-width: 640px) {
  .sm\:space-y-4 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-y-reverse: 0;
    margin-top: calc(1rem * calc(1 - var(--tw-space-y-reverse)));
    margin-bottom: calc(1rem * var(--tw-space-y-reverse));
  }
}`,

	"SpaceY4X2l": `
@media (min-width: 1536px) {
  .2xl\:space-y-4 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-y-reverse: 0;
    margin-top: calc(1rem * calc(1 - var(--tw-space-y-reverse)));
    margin-bottom: calc(1rem * var(--tw-space-y-reverse));
  }
}`,

	"SpaceY4Xl": `
@media (min-width: 1280px) {
  .xl\:space-y-4 > :not([hidden]) ~ :not([hidden]) {
    --tw-space-y-reverse: 0;
    margin-top: calc(1rem * calc(1 - var(--tw-space-y-reverse)));
    margin-bottom: calc(1rem * var(--tw-space-y-reverse));
  }
}`,

	"TextBase": `
.text-base {
  font-size: 1rem;
  line-height: 1.5;
}`,

	"TextBaseLg": `
@media (min-width: 1024px) {
  .lg\:text-base {
    font-size: 1rem;
    line-height: 1.5;
  }
}`,

	"TextBaseMd": `
@media (min-width: 768px) {
  .md\:text-base {
    font-size: 1rem;
    line-height: 1.5;
  }
}`,

	"TextBaseSm": `
@media (min-width: 640px) {
  .sm\:text-base {
    font-size: 1rem;
    line-height: 1.5;
  }
}`,

	"TextBaseX2l": `
@media (min-width: 1536px) {
  .2xl\:text-base {
    font-size: 1rem;
    line-height: 1.5;
  }
}`,

	"TextBaseXl": `
@media (min-width: 1280px) {
  .xl\:text-base {
    font-size: 1rem;
    line-height: 1.5;
  }
}`,

	"TextBlack": `
.text-black {
  color: #000;
}`,

	"TextBlackLg": `
@media (min-width: 1024px) {
  .lg\:text-black {
    color: #000;
  }
}`,

	"TextBlackMd": `
@media (min-width: 768px) {
  .md\:text-black {
    color: #000;
  }
}`,

	"TextBlackSm": `
@media (min-width: 640px) {
  .sm\:text-black {
    color: #000;
  }
}`,

	"TextBlackX2l": `
@media (min-width: 1536px) {
  .2xl\:text-black {
    color: #000;
  }
}`,

	"TextBlackXl": `
@media (min-width: 1280px) {
  .xl\:text-black {
    color: #000;
  }
}`,

	"TextCenter": `
.text-center {
  text-align: center;
}`,

	"TextCenterLg": `
@media (min-width: 1024px) {
  .lg\:text-center {
    text-align: center;
  }
}`,

	"TextCenterMd": `
@media (min-width: 768px) {
  .md\:text-center {
    text-align: center;
  }
}`,

	"TextCenterSm": `
@media (min-width: 640px) {
  .sm\:text-center {
    text-align: center;
  }
}`,

	"TextCenterX2l": `
@media (min-width: 1536px) {
  .2xl\:text-center {
    text-align: center;
  }
}`,

	"TextCenterXl": `
@media (min-width: 1280px) {
  .xl\:text-center {
    text-align: center;
  }
}`,

	"TextGray700": `
.text-gray-700 {
  --tw-text-opacity: 1;
  color: rgb(55 65 81 / var(--tw-text-opacity));
}`,

	"TextGray700Lg": `
@media (min-width: 1024px) {
  .lg\:text-gray-700 {
    --tw-text-opacity: 1;
    color: rgb(55 65 81 / var(--tw-text-opacity));
  }
}`,

	"TextGray700Md": `
@media (min-width: 768px) {
  .md\:text-gray-700 {
    --tw-text-opacity: 1;
    color: rgb(55 65 81 / var(--tw-text-opacity));
  }
}`,

	"TextGray700Sm": `
@media (min-width: 640px) {
  .sm\:text-gray-700 {
    --tw-text-opacity: 1;
    color: rgb(55 65 81 / var(--tw-text-opacity));
  }
}`,

	"TextGray700X2l": `
@media (min-width: 1536px) {
  .2xl\:text-gray-700 {
    --tw-text-opacity: 1;
    color: rgb(55 65 81 / var(--tw-text-opacity));
  }
}`,

	"TextGray700Xl": `
@media (min-width: 1280px) {
  .xl\:text-gray-700 {
    --tw-text-opacity: 1;
    color: rgb(55 65 81 / var(--tw-text-opacity));
  }
}`,

	"TextLeft": `
.text-left {
  text-align: left;
}`,

	"TextLeftLg": `
@media (min-width: 1024px) {
  .lg\:text-left {
    text-align: left;
  }
}`,

	"TextLeftMd": `
@media (min-width: 768px) {
  .md\:text-left {
    text-align: left;
  }
}`,

	"TextLeftSm": `
@media (min-width: 640px) {
  .sm\:text-left {
    text-align: left;
  }
}`,

	"TextLeftX2l": `
@media (min-width: 1536px) {
  .2xl\:text-left {
    text-align: left;
  }
}`,

	"TextLeftXl": `
@media (min-width: 1280px) {
  .xl\:text-left {
    text-align: left;
  }
}`,

	"TextLg": `
.text-lg {
  font-size: 1.125rem;
  line-height: 1.75rem;
}`,

	"TextLgLg": `
@media (min-width: 1024px) {
  .lg\:text-lg {
    font-size: 1.125rem;
    line-height: 1.75rem;
  }
}`,

	"TextLgMd": `
@media (min-width: 768px) {
  .md\:text-lg {
    font-size: 1.125rem;
    line-height: 1.75rem;
  }
}`,

	"TextLgSm": `
@media (min-width: 640px) {
  .sm\:text-lg {
    font-size: 1.125rem;
    line-height: 1.75rem;
  }
}`,

	"TextLgX2l": `
@media (min-width: 1536px) {
  .2xl\:text-lg {
    font-size: 1.125rem;
    line-height: 1.75rem;
  }
}`,

	"TextLgXl": `
@media (min-width: 1280px) {
  .xl\:text-lg {
    font-size: 1.125rem;
    line-height: 1.75rem;
  }
}`,

	"BgWhite": `
.bg-white {
  --tw-bg-opacity: 1;
  background-color: rgb(255 255 255 / var(--tw-bg-opacity));
}`,

	"BgWhiteLg": `
@media (min-width: 1024px) {
  .lg\:bg-white {
    --tw-bg-opacity: 1;
    background-color: rgb(255 255 255 / var(--tw-bg-opacity));
  }
}`,

	"BgWhiteMd": `
@media (min-width: 768px) {
  .md\:bg-white {
    --tw-bg-opacity: 1;
    background-color: rgb(255 255 255 / var(--tw-bg-opacity));
  }
}`,

	"BgWhiteSm": `
@media (min-width: 640px) {
  .sm\:bg-white {
    --tw-bg-opacity: 1;
    background-color: rgb(255 255 255 / var(--tw-bg-opacity));
  }
}`,

	"BgWhiteX2l": `
@media (min-width: 1536px) {
  .2xl\:bg-white {
    --tw-bg-opacity: 1;
    background-color: rgb(255 255 255 / var(--tw-bg-opacity));
  }
}`,

	"BgWhiteXl": `
@media (min-width: 1280px) {
  .xl\:bg-white {
    --tw-bg-opacity: 1;
    background-color: rgb(255 255 255 / var(--tw-bg-opacity));
  }
}`,

	"BgRed500": `
.bg-red-500 {
  --tw-bg-opacity: 1;
  background-color: rgb(239 68 68 / var(--tw-bg-opacity));
}`,

	"BgRed500Lg": `
@media (min-width: 1024px) {
  .lg\:bg-red-500 {
    --tw-bg-opacity: 1;
    background-color: rgb(239 68 68 / var(--tw-bg-opacity));
  }
}`,

	"BgRed500Md": `
@media (min-width: 768px) {
  .md\:bg-red-500 {
    --tw-bg-opacity: 1;
    background-color: rgb(239 68 68 / var(--tw-bg-opacity));
  }
}`,

	"BgRed500Sm": `
@media (min-width: 640px) {
  .sm\:bg-red-500 {
    --tw-bg-opacity: 1;
    background-color: rgb(239 68 68 / var(--tw-bg-opacity));
  }
}`,

	"BgRed500X2l": `
@media (min-width: 1536px) {
  .2xl\:bg-red-500 {
    --tw-bg-opacity: 1;
    background-color: rgb(239 68 68 / var(--tw-bg-opacity));
  }
}`,

	"BgRed500Xl": `
@media (min-width: 1280px) {
  .xl\:bg-red-500 {
    --tw-bg-opacity: 1;
    background-color: rgb(239 68 68 / var(--tw-bg-opacity));
  }
}`,

	"TextOpacity75": `
.text-opacity-75 {
  --tw-text-opacity: 0.75;
}`,

	"TextOpacity75Lg": `
@media (min-width: 1024px) {
  .lg\:text-opacity-75 {
    --tw-text-opacity: 0.75;
  }
}`,

	"TextOpacity75Md": `
@media (min-width: 768px) {
  .md\:text-opacity-75 {
    --tw-text-opacity: 0.75;
  }
}`,

	"TextOpacity75Sm": `
@media (min-width: 640px) {
  .sm\:text-opacity-75 {
    --tw-text-opacity: 0.75;
  }
}`,

	"TextOpacity75X2l": `
@media (min-width: 1536px) {
  .2xl\:text-opacity-75 {
    --tw-text-opacity: 0.75;
  }
}`,

	"TextOpacity75Xl": `
@media (min-width: 1280px) {
  .xl\:text-opacity-75 {
    --tw-text-opacity: 0.75;
  }
}`,

	"BgOpacity75": `
.bg-opacity-75 {
  --tw-bg-opacity: 0.75;
}`,

	"BgOpacity75Lg": `
@media (min-width: 1024px) {
  .lg\:bg-opacity-75 {
    --tw-bg-opacity: 0.75;
  }
}`,

	"BgOpacity75Md": `
@media (min-width: 768px) {
  .md\:bg-opacity-75 {
    --tw-bg-opacity: 0.75;
  }
}`,

	"BgOpacity75Sm": `
@media (min-width: 640px) {
  .sm\:bg-opacity-75 {
    --tw-bg-opacity: 0.75;
  }
}`,

	"BgOpacity75X2l": `
@media (min-width: 1536px) {
  .2xl\:bg-opacity-75 {
    --tw-bg-opacity: 0.75;
  }
}`,

	"BgOpacity75Xl": `
@media (min-width: 1280px) {
  .xl\:bg-opacity-75 {
    --tw-bg-opacity: 0.75;
  }
}`,

	"TextSm": `
.text-sm {
  font-size: 0.875rem;
  line-height: 1.25rem;
}`,

	"TextSmLg": `
@media (min-width: 1024px) {
  .lg\:text-sm {
    font-size: 0.875rem;
    line-height: 1.25rem;
  }
}`,

	"TextSmMd": `
@media (min-width: 768px) {
  .md\:text-sm {
    font-size: 0.875rem;
    line-height: 1.25rem;
  }
}`,

	"TextSmSm": `
@media (min-width: 640px) {
  .sm\:text-sm {
    font-size: 0.875rem;
    line-height: 1.25rem;
  }
}`,

	"TextSmX2l": `
@media (min-width: 1536px) {
  .2xl\:text-sm {
    font-size: 0.875rem;
    line-height: 1.25rem;
  }
}`,

	"TextSmXl": `
@media (min-width: 1280px) {
  .xl\:text-sm {
    font-size: 0.875rem;
    line-height: 1.25rem;
  }
}`,

	"TextWhite": `
.text-white {
  --tw-text-opacity: 1;
  color: rgb(255 255 255 / var(--tw-text-opacity));
}`,

	"TextWhiteLg": `
@media (min-width: 1024px) {
  .lg\:text-white {
    --tw-text-opacity: 1;
    color: rgb(255 255 255 / var(--tw-text-opacity));
  }
}`,

	"TextWhiteMd": `
@media (min-width: 768px) {
  .md\:text-white {
    --tw-text-opacity: 1;
    color: rgb(255 255 255 / var(--tw-text-opacity));
  }
}`,

	"TextWhiteSm": `
@media (min-width: 640px) {
  .sm\:text-white {
    --tw-text-opacity: 1;
    color: rgb(255 255 255 / var(--tw-text-opacity));
  }
}`,

	"TextWhiteX2l": `
@media (min-width: 1536px) {
  .2xl\:text-white {
    --tw-text-opacity: 1;
    color: rgb(255 255 255 / var(--tw-text-opacity));
  }
}`,

	"TextWhiteXl": `
@media (min-width: 1280px) {
  .xl\:text-white {
    --tw-text-opacity: 1;
    color: rgb(255 255 255 / var(--tw-text-opacity));
  }
}`,

	"TextXl": `
.text-xl {
  font-size: 1.25rem;
  line-height: 1.75rem;
}`,

	"TextXlLg": `
@media (min-width: 1024px) {
  .lg\:text-xl {
    font-size: 1.25rem;
    line-height: 1.75rem;
  }
}`,

	"TextXlMd": `
@media (min-width: 768px) {
  .md\:text-xl {
    font-size: 1.25rem;
    line-height: 1.75rem;
  }
}`,

	"TextXlSm": `
@media (min-width: 640px) {
  .sm\:text-xl {
    font-size: 1.25rem;
    line-height: 1.75rem;
  }
}`,

	"TextXlX2l": `
@media (min-width: 1536px) {
  .2xl\:text-xl {
    font-size: 1.25rem;
    line-height: 1.75rem;
  }
}`,

	"TextXlXl": `
@media (min-width: 1280px) {
  .xl\:text-xl {
    font-size: 1.25rem;
    line-height: 1.75rem;
  }
}`,

	"Transition": `
.transition {
  transition-property: background-color, border-color, color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
  transition-duration: 150ms;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}`,

	"TransitionColors": `
.transition-colors {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
}`,

	"TransitionColorsLg": `
@media (min-width: 1024px) {
  .lg\:transition-colors {
    transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
  }
}`,

	"TransitionColorsMd": `
@media (min-width: 768px) {
  .md\:transition-colors {
    transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
  }
}`,

	"TransitionColorsSm": `
@media (min-width: 640px) {
  .sm\:transition-colors {
    transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
  }
}`,

	"TransitionColorsX2l": `
@media (min-width: 1536px) {
  .2xl\:transition-colors {
    transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
  }
}`,

	"TransitionColorsXl": `
@media (min-width: 1280px) {
  .xl\:transition-colors {
    transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
  }
}`,

	"TransitionLg": `
@media (min-width: 1024px) {
  .lg\:transition {
    transition-property: background-color, border-color, color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
    transition-duration: 150ms;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
}`,

	"TransitionMd": `
@media (min-width: 768px) {
  .md\:transition {
    transition-property: background-color, border-color, color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
    transition-duration: 150ms;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
}`,

	"TransitionOpacity": `
.transition-opacity {
  transition-property: opacity;
}`,

	"TransitionOpacityLg": `
@media (min-width: 1024px) {
  .lg\:transition-opacity {
    transition-property: opacity;
  }
}`,

	"TransitionOpacityMd": `
@media (min-width: 768px) {
  .md\:transition-opacity {
    transition-property: opacity;
  }
}`,

	"TransitionOpacitySm": `
@media (min-width: 640px) {
  .sm\:transition-opacity {
    transition-property: opacity;
  }
}`,

	"TransitionOpacityX2l": `
@media (min-width: 1536px) {
  .2xl\:transition-opacity {
    transition-property: opacity;
  }
}`,

	"TransitionOpacityXl": `
@media (min-width: 1280px) {
  .xl\:transition-opacity {
    transition-property: opacity;
  }
}`,

	"TransitionSm": `
.transition-sm {
  transition-duration: 75ms;
}`,

	"TransitionX2l": `
@media (min-width: 1536px) {
  .2xl\:transition {
    transition-property: background-color, border-color, color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
    transition-duration: 150ms;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
}`,

	"TransitionXl": `
@media (min-width: 1280px) {
  .xl\:transition {
    transition-property: background-color, border-color, color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter;
    transition-duration: 150ms;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  }
}`,

	"WAuto": `
.w-auto {
  width: auto;
}`,

	"WAutoLg": `
@media (min-width: 1024px) {
  .lg\:w-auto {
    width: auto;
  }
}`,

	"WAutoMd": `
@media (min-width: 768px) {
  .md\:w-auto {
    width: auto;
  }
}`,

	"WAutoSm": `
@media (min-width: 640px) {
  .sm\:w-auto {
    width: auto;
  }
}`,

	"WAutoX2l": `
@media (min-width: 1536px) {
  .2xl\:w-auto {
    width: auto;
  }
}`,

	"WAutoXl": `
@media (min-width: 1280px) {
  .xl\:w-auto {
    width: auto;
  }
}`,

	"WFull": `
.w-full {
  width: 100%;
}`,

	"WFullLg": `
@media (min-width: 1024px) {
  .lg\:w-full {
    width: 100%;
  }
}`,

	"WFullMd": `
@media (min-width: 768px) {
  .md\:w-full {
    width: 100%;
  }
}`,

	"WFullSm": `
@media (min-width: 640px) {
  .sm\:w-full {
    width: 100%;
  }
}`,

	"WFullX2l": `
@media (min-width: 1536px) {
  .2xl\:w-full {
    width: 100%;
  }
}`,

	"WFullXl": `
@media (min-width: 1280px) {
  .xl\:w-full {
    width: 100%;
  }
}`,
}
