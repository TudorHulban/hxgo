package dsl

import "testing"

func BenchmarkStyleOnly(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_ = NewCSSFor("div").Color("red")
	}
}

func BenchmarkHTMLCall(b *testing.B) {
	b.ReportAllocs()

	el := Div(
		Text("hello"),
		NewCSSFor("div").Color("red").AsNode(),
	)

	b.ResetTimer()

	for b.Loop() {
		_ = Render(el)
	}
}

func BenchmarkCSSCall(b *testing.B) {
	b.ReportAllocs()

	el := Div(
		Text("hello"),
		NewCSSFor("div").Color("red").AsNode(),
	)

	b.ResetTimer()

	for b.Loop() {
		_, _, _ = RenderFull(el)
	}
}
