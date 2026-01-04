package dsl

import "testing"

func BenchmarkIfCall(b *testing.B) {
	b.ReportAllocs()

	el := If(false, Text("hi!"))

	b.ResetTimer()

	for b.Loop() {
		_ = Render(el)
	}
}
