package dsl

import (
	"testing"
)

// interfaces
// BenchmarkRenderNodes/0._error_input-16         	 4914921	       241.5 ns/op	     128 B/op	       3 allocs/op
// BenchmarkRenderNodes/1._empty_input-16         	 5310886	       225.3 ns/op	     112 B/op	       2 allocs/op
// BenchmarkRenderNodes/2._valid_HTML_input-16    	 3290210	       364.5 ns/op	     128 B/op	       3 allocs/op
// BenchmarkRenderNodes/3._valid_CSS_input-16     	 4227207	       281.9 ns/op	     192 B/op	       3 allocs/op
// BenchmarkRenderNodes/4._valid_HTML-CSS_input-16         	 2795299	       425.7 ns/op	     208 B/op	       4 allocs/op

// functional
// BenchmarkRenderHTMLandCSS/1._empty_input-16         	26171468	        45.10 ns/op	       0 B/op	       0 allocs/op
// BenchmarkRenderHTMLandCSS/2._valid_HTML_input-16    	  550248	      1965 ns/op	     552 B/op	       6 allocs/op
// BenchmarkRenderHTMLandCSS/3._valid_CSS_input-16     	  157761	      6663 ns/op	    1344 B/op	      15 allocs/op
// BenchmarkRenderHTMLandCSS/4._valid_HTML-CSS_input-16         	  132006	      9041 ns/op	    1896 B/op	      21 allocs/op

func BenchmarkRenderHTMLandCSS(b *testing.B) {
	tests := []struct {
		description string
		nodes       []Node
	}{
		{
			description: "1. empty input",
			nodes:       nil,
		},
		{
			description: "2. valid HTML input",
			nodes: []Node{
				Div(Class("card")),
			},
		},
		{
			description: "2a. valid HTML input",
			nodes: []Node{
				Div(AttrClass("card")),
			},
		},
		{
			description: "2b. valid HTML input",
			nodes: []Node{
				Div(AttrClass("card")),
				Div(AttrClass("card")),
			},
		},
		{
			description: "3. valid CSS input",
			nodes: []Node{
				Styled(
					Noop,
					Style{
						Selector: ".card",
						Props: map[string]string{
							"padding":       "20px",
							"border-radius": "8px",
							"box-shadow":    "0 4px 12px rgba(0,0,0,0.1)",
						},
						Media: "min-width: 768px",
					},
				),
			},
		},
		{
			description: "4. valid HTML-CSS input",
			nodes: []Node{
				Div(Class("card")),
				Styled(
					Noop,
					Style{
						Selector: ".card",
						Props: map[string]string{
							"padding":       "20px",
							"border-radius": "8px",
							"box-shadow":    "0 4px 12px rgba(0,0,0,0.1)",
						},
						Media: "min-width: 768px",
					},
				),
			},
		},
	}

	for _, tt := range tests {
		b.Run(
			tt.description,
			func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					_, _ = RenderHTMLandCSS(tt.nodes...)
				}
			},
		)
	}
}
