package dsl

import "strconv"

// Single builder used by the renderer.
// accumulator accumulates HTML bytes and CSS styles during rendering.
// For better performance,
// use NewAcc with an estimated size based on the typical output.
type accumulator struct {
	html   []byte
	styles []Style

	registryCSS map[CSSKey]map[*ActualCSS]DesktopFirst
}

// newAccumulator creates an Accumulator with pre-allocated capacity.
// estimatedHTMLSize should be the approximate size of the rendered HTML in bytes.
// Pre-allocating reduces memory allocations during rendering.
// Example: for a typical page of ~10KB, use newAccumulator(10240, 16)
func newAccumulator(estimatedHTMLSize, estimatedCSSRules int) *accumulator {
	return &accumulator{
		html:   make([]byte, 0, estimatedHTMLSize),
		styles: make([]Style, 0, estimatedCSSRules),

		// registryCSS: make(map[CSSKey]map[*ActualCSS]DesktopFirst),
	}
}

func (a *accumulator) initialize() {
	if a.registryCSS == nil {
		a.registryCSS = make(map[CSSKey]map[*ActualCSS]DesktopFirst)
	}
}

func (a *accumulator) Write1(value string) {
	a.html = append(a.html, value...)
}

// Fused writes.
func (a *accumulator) Write3(value1, value2, value3 string) {
	a.html = append(a.html, value1...)
	a.html = append(a.html, value2...)
	a.html = append(a.html, value3...)
}

// Fused writes.
func (a *accumulator) Write5(value1, value2, value3, value4, value5 string) {
	a.html = append(a.html, value1...)
	a.html = append(a.html, value2...)
	a.html = append(a.html, value3...)
	a.html = append(a.html, value4...)
	a.html = append(a.html, value5...)
}

func (a *accumulator) BuildCSS(estimatedCSS ...int) []byte {
	if len(a.registryCSS) == 0 {
		return nil
	}

	cssLength := 512
	if len(estimatedCSS) == 1 {
		cssLength = estimatedCSS[0]
	}

	result := make([]byte, 0, cssLength)

	for key, bucket := range a.registryCSS {
		for fnPtr, desktopFirst := range bucket {
			css := (*fnPtr)()

			// no media query
			if key.InflexionPointPX == 0 {
				result = append(result, key.Selector...)
				result = append(result, '{')
				result = append(result, css...)
				result = append(result, '}')

				continue
			}

			// media query
			if desktopFirst {
				result = append(result, "@media (min-width:"...)
			} else {
				result = append(result, "@media (max-width:"...)
			}

			result = strconv.AppendUint(result, uint64(key.InflexionPointPX), 10)
			result = append(result, "px) {"...)
			result = append(result, key.Selector...)
			result = append(result, '{')
			result = append(result, css...)
			result = append(result, "}}"...)
		}
	}

	return result
}
