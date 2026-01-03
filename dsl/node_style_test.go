package dsl

import "testing"

func BenchmarkStyleOnly(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_ = Style{
			Selector: "div",
			Props:    map[string]string{"color": "red"},
		}
	}
}

func BenchmarkStyleCall(b *testing.B) {
	b.ReportAllocs()

	s := Styled(
		Text("hello"),
		Style{Selector: "div", Props: map[string]string{"color": "red"}},
	)

	b.ResetTimer()

	for b.Loop() {
		_ = s()
	}
}
