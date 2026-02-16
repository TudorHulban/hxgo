package dsl

import "testing"

func BenchmarkClassCall(b *testing.B) {
	b.ReportAllocs()

	el := Class("hi!")

	b.ResetTimer()

	for b.Loop() {
		_ = RenderConvertedHTML(el)
	}
}
