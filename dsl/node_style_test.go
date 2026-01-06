package dsl

import "testing"

func BenchmarkStyleOnly(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_ = Style{
			Selector: "div",
			Props:    [][2]string{{"color", "red"}},
		}
	}
}

func BenchmarkStyleCall(b *testing.B) {
	b.ReportAllocs()

	el := GetStyledNode(
		Text("hello"),
		Style{
			Selector: "div",
			Props:    [][2]string{{"color", "red"}},
		},
	)

	b.ResetTimer()

	for b.Loop() {
		_ = Render(el)
	}
}
