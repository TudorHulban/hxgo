package fastnode

import "testing"

func BenchmarkFastNode(b *testing.B) {
	n := Div(
		Class("card"),
		Text("hello"),
	)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Render(n)
	}
}
