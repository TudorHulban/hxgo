package dsl

import "testing"

func BenchmarkClassCall(b *testing.B) {
	b.ReportAllocs()

	node := Class("hi!")

	b.ResetTimer()

	for b.Loop() {
		_ = node()
	}
}
