package dsl

import "testing"

func BenchmarkAttrWithValueOnly(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_ = AttrWithValue("data-x", "hi!")
	}
}

func BenchmarkAttrWithValueCall(b *testing.B) {
	b.ReportAllocs()

	el := AttrWithValue("data-x", "hi!")

	b.ResetTimer()

	for b.Loop() {
		_ = Render(el)
	}
}
