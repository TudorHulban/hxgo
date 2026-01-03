package dsl

import "testing"

func BenchmarkTextCall(b *testing.B) {
	b.ReportAllocs()

	node := Text("hi!")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = node()
	}
}
