package dsl

import (
	"testing"
)

// interfaces
// BenchmarkRenderNodes/1._empty_input-16         	 5310886	       225.3 ns/op	     112 B/op	       2 allocs/op
// BenchmarkRenderNodes/2._valid_HTML_input-16    	 3290210	       364.5 ns/op	     128 B/op	       3 allocs/op
// BenchmarkRenderNodes/3._valid_CSS_input-16     	 4227207	       281.9 ns/op	     192 B/op	       3 allocs/op
// BenchmarkRenderNodes/4._valid_HTML-CSS_input-16         	 2795299	       425.7 ns/op	     208 B/op	       4 allocs/op

// unsafe
// BenchmarkRenderHTMLandCSS/1._empty_input-12         	355249774	         3.381 ns/op	       0 B/op	       0 allocs/op
// BenchmarkRenderHTMLandCSS/2a._valid_HTML_input-12   	 7262638	       164.0 ns/op	     104 B/op	       4 allocs/op
// BenchmarkRenderHTMLandCSS/2b._valid_HTML_input-12   	 6672134	       177.0 ns/op	     104 B/op	       4 allocs/op
// BenchmarkRenderHTMLandCSS/2c._valid_HTML_input-12   	 4571163	       259.8 ns/op	     168 B/op	       5 allocs/op
// BenchmarkRenderHTMLandCSS/3._valid_CSS_input-12     	  764766	      1345 ns/op	    1360 B/op	      15 allocs/op
// BenchmarkRenderHTMLandCSS/4._valid_HTML-CSS_input-12         	  710091	      1508 ns/op	    1416 B/op	      18 allocs/op

// BenchmarkRenderHTMLandCSS/1._empty_input-16         	99964186	        12.75 ns/op	       0 B/op	       0 allocs/op
// BenchmarkRenderHTMLandCSS/2a._valid_HTML_input-16   	 1619826	       711.1 ns/op	     104 B/op	       4 allocs/op
// BenchmarkRenderHTMLandCSS/2b._valid_HTML_input-16   	 1610301	       733.4 ns/op	     104 B/op	       4 allocs/op
// BenchmarkRenderHTMLandCSS/2c._valid_HTML_input-16   	 1000000	      1131 ns/op	     168 B/op	       5 allocs/op
// BenchmarkRenderHTMLandCSS/3._valid_CSS_input-16     	  182564	      6128 ns/op	    1504 B/op	      16 allocs/op
// BenchmarkRenderHTMLandCSS/4._valid_HTML-CSS_input-16         	  183285	      6048 ns/op	    1504 B/op	      16 allocs/op

// go test -bench=RenderHTMLandCSS -benchmem -run=^$ -memprofile=mem.pprof
// go tool pprof -pdf mem.pprof > mem.pdf

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
			description: "2a. valid HTML input",
			nodes: []Node{
				Div(Class("card")),
			},
		},
		{
			description: "2b. valid HTML input",
			nodes: []Node{
				Div(AttrClass("card")),
			},
		},
		{
			description: "2c. valid HTML input",
			nodes: []Node{
				Div(AttrClass("card")),
				Div(AttrClass("card")),
			},
		},
		{
			description: "3. valid CSS input",
			nodes: []Node{
				NewCSSFor(".card").
					WithBreakpoint("768px").
					Padding("20px").
					ShadowBox("0 4px 12px rgba(0,0,0,0.1)").
					Radius("8px").
					AsNode(),
			},
		},
		{
			description: "4. valid HTML-CSS input",
			nodes: []Node{
				NewCSSForClass("card").
					WithBreakpoint("768px").
					Padding("20px").
					Radius("8px").
					ShadowBox("0 4px 12px rgba(0,0,0,0.1)").
					AsNode(),
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
					_, _ = RenderHTMLandStyles(tt.nodes...)
				}
			},
		)
	}
}

// BenchmarkRenderHTML/1._empty_input-16         	93749204	        12.07 ns/op	       0 B/op	       0 allocs/op
// BenchmarkRenderHTML/2a._valid_HTML_input-16   	 1688866	       728.2 ns/op	     104 B/op	       4 allocs/op
// BenchmarkRenderHTML/2b._valid_HTML_input-16   	 1576878	       763.5 ns/op	     104 B/op	       4 allocs/op
// BenchmarkRenderHTML/2c._valid_HTML_input-16   	 1033178	      1174 ns/op	     168 B/op	       5 allocs/op
// BenchmarkRenderHTML/3._valid_CSS_input-16     	 2327839	       526.4 ns/op	     144 B/op	       2 allocs/op
// BenchmarkRenderHTML/4._valid_HTML-CSS_input-16         	 2352422	       513.6 ns/op	     144 B/op	       2 allocs/op

func BenchmarkRenderHTML(b *testing.B) {
	tests := []struct {
		description string
		nodes       []Node
	}{
		{
			description: "1. empty input",
			nodes:       nil,
		},
		{
			description: "2a. valid HTML input",
			nodes: []Node{
				Div(Class("card")),
			},
		},
		{
			description: "2b. valid HTML input",
			nodes: []Node{
				Div(AttrClass("card")),
			},
		},
		{
			description: "2c. valid HTML input",
			nodes: []Node{
				Div(AttrClass("card")),
				Div(AttrClass("card")),
			},
		},
		{
			description: "3. valid CSS input",
			nodes: []Node{
				NewCSSForClass("card").
					WithBreakpoint("768px").
					Padding("20px").
					Radius("8px").
					ShadowBox("0 4px 12px rgba(0,0,0,0.1)").
					AsNode(),
			},
		},
		{
			description: "4. valid HTML-CSS input",
			nodes: []Node{
				NewCSSForClass("card").
					WithBreakpoint("768px").
					Padding("20px").
					Radius("8px").
					ShadowBox("0 4px 12px rgba(0,0,0,0.1)").
					AsNode(),
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
					_ = Render(tt.nodes...)
				}
			},
		)
	}
}

// BenchmarkRenderHTMLWCapacity/1._empty_input-16         	475330550	        12.06 ns/op	       0 B/op	       0 allocs/op
// BenchmarkRenderHTMLWCapacity/2a._valid_HTML_input-16   	 2438256	       492.3 ns/op	     176 B/op	       2 allocs/op
// BenchmarkRenderHTMLWCapacity/2b._valid_HTML_input-16   	 2293862	       516.6 ns/op	     176 B/op	       2 allocs/op
// BenchmarkRenderHTMLWCapacity/2c._valid_HTML_input-16   	 1615639	       753.2 ns/op	     304 B/op	       2 allocs/op
// BenchmarkRenderHTMLWCapacity/3._valid_CSS_input-16     	 2321086	       504.8 ns/op	     112 B/op	       2 allocs/op
// BenchmarkRenderHTMLWCapacity/4._valid_HTML-CSS_input-16         	 1457641	       823.6 ns/op	     240 B/op	       3 allocs/op

func BenchmarkRenderHTMLWCapacity(b *testing.B) {
	tests := []struct {
		description   string
		nodes         []Node
		capacityBytes int
	}{
		{
			description: "1. empty input",
			nodes:       nil,
		},
		{
			description: "2a. valid HTML input",
			nodes: []Node{
				Div(Class("card")),
			},
			capacityBytes: 128,
		},
		{
			description: "2b. valid HTML input",
			nodes: []Node{
				Div(AttrClass("card")),
			},
			capacityBytes: 128,
		},
		{
			description: "2c. valid HTML input",
			nodes: []Node{
				Div(AttrClass("card")),
				Div(AttrClass("card")),
			},
			capacityBytes: 256,
		},
		{
			description: "3. valid CSS input",
			nodes: []Node{
				NewCSSForClass("card").
					WithBreakpoint("768px").
					Padding("20px").
					Radius("8px").
					ShadowBox("0 4px 12px rgba(0,0,0,0.1)").
					AsNode(),
			},
		},
		{
			description: "4. valid HTML-CSS input",
			nodes: []Node{
				NewCSSForClass("card").
					WithBreakpoint("768px").
					Padding("20px").
					Radius("8px").
					ShadowBox("0 4px 12px rgba(0,0,0,0.1)").
					AsNode(),
			},
			capacityBytes: 128,
		},
	}

	for _, tt := range tests {
		b.Run(
			tt.description,
			func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					_ = RenderHTMLWithCapacity(
						tt.capacityBytes,
						tt.nodes...,
					)
				}
			},
		)
	}
}
