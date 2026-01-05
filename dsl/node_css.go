package dsl

import "unsafe"

type DesktopFirst bool

type ActualCSS func() string

type CSSKey struct {
	Selector         string
	InflexionPointPX uint16
}

type CSS struct {
	Selector         string
	InflexionPointPX uint16
	DesktopFirst     DesktopFirst
	ActualCSS        ActualCSS
}

func CSSNode(rules ...CSS) Node {
	return Node{
		fn:   renderCSS,
		data: unsafe.Pointer(&rules),
	}
}

func renderCSS(a *accumulator, p unsafe.Pointer) {
	a.initialize() // ← one-time cost only if needed

	rules := *(*[]CSS)(p)

	for i := range rules {
		rule := &rules[i]

		key := CSSKey{
			Selector:         rule.Selector,
			InflexionPointPX: rule.InflexionPointPX,
		}

		// ensure map for this key exists
		bucket, exists := a.registryCSS[key]
		if !exists {
			bucket = make(map[*ActualCSS]DesktopFirst)
			a.registryCSS[key] = bucket
		}

		// dedupe by function pointer identity
		if _, exists := bucket[&rule.ActualCSS]; exists {
			continue
		}

		bucket[&rule.ActualCSS] = rule.DesktopFirst
	}
}
