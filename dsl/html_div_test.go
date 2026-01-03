package dsl

import "testing"

func BenchmarkElOnly(b *testing.B) {
	b.ReportAllocs()

	b.ResetTimer()

	for b.Loop() {
		_ = El("div")
	}
}

func BenchmarkElCall(b *testing.B) {
	b.ReportAllocs()

	el := El("div", Text("hi!"))

	b.ResetTimer()

	for b.Loop() {
		_ = Render(el)
	}
}

func BenchmarkElWIdCall(b *testing.B) {
	b.ReportAllocs()

	el := ElWId("div", "myid", Text("hello"))

	b.ResetTimer()

	for b.Loop() {
		_ = Render(el)
	}
}

func BenchmarkDiv(b *testing.B) {
	b.ReportAllocs()

	el := El("div", Text("hi!"))

	b.ResetTimer()

	for b.Loop() {
		_ = Render(el)
	}
}

func BenchmarkDivRender(b *testing.B) {
	b.ReportAllocs()

	div := El("div", Text("hello"))

	b.ResetTimer()

	for b.Loop() {
		_, _ = RenderHTMLandCSS(div)
	}
}
