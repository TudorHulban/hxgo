package dsl

import "testing"

func BenchmarkElOnly(b *testing.B) {
	b.ReportAllocs()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = El("div") // do not call the closure
	}
}

func BenchmarkElCall(b *testing.B) {
	b.ReportAllocs()

	el := El("div", Text("hi!"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = el()
	}
}

func BenchmarkElWIdCall(b *testing.B) {
	b.ReportAllocs()
	el := ElWId("div", "myid", Text("hello"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = el()
	}
}

func BenchmarkDiv(b *testing.B) {
	b.ReportAllocs()

	div := El("div", Text("hi!"))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		out := div()
		_ = out.HTMLParts // or out.HTML if interface version
	}
}

func BenchmarkDivRender(b *testing.B) {
	b.ReportAllocs()

	div := El("div", Text("hello"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = RenderHTMLandCSS(div) // interface or B‑Prime version
	}
}
