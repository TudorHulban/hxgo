package dsl

import "testing"

func BenchmarkStyleOnly(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_ = CSSContribution{
			CSSContributionKey: CSSContributionKey{
				Selector: "div",
			},
			DeclarativeStyle: [][2]string{{"color", "red"}},
		}
	}
}

func BenchmarkHTMLCall(b *testing.B) {
	b.ReportAllocs()

	el := Div(
		Text("hello"),
		CSSContribution{
			CSSContributionKey: CSSContributionKey{
				Selector: "div",
			},
			DeclarativeStyle: [][2]string{{"color", "red"}},
		}.
			AsNode(),
	)

	b.ResetTimer()

	for b.Loop() {
		_ = Render(el)
	}
}

func BenchmarkCSSCall(b *testing.B) {
	b.ReportAllocs()

	el := Div(
		Text("hello"),
		CSSContribution{
			CSSContributionKey: CSSContributionKey{
				Selector: "div",
			},
			DeclarativeStyle: [][2]string{{"color", "red"}},
		}.
			AsNode(),
	)

	b.ResetTimer()

	for b.Loop() {
		_, _, _ = RenderFull(el)
	}
}
