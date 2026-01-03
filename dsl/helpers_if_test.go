package dsl

import "testing"

func BenchmarkIfCall(b *testing.B) {
	b.ReportAllocs()

	el := If(false, Text("hi!"))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Render(el)
	}
}
