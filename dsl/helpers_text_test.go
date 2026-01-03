package dsl

import "testing"

func BenchmarkTextCall(b *testing.B) {
	b.ReportAllocs()

	el := Text("hi!")

	b.ResetTimer()

	for b.Loop() {
		_ = Render(el)
	}
}
