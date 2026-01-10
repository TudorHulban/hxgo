package dsl

import (
	"testing"
)

// BenchmarkRenderHTMLandCSS/1._empty_input-16         	783166084	         1.502 ns/op	       0 B/op	       0 allocs/op
// BenchmarkRenderHTMLandCSS/2a._valid_HTML_input-16   	10475170	       113.7 ns/op	     104 B/op	       4 allocs/op
// BenchmarkRenderHTMLandCSS/2b._valid_HTML_input-16   	10379050	       115.4 ns/op	     104 B/op	       4 allocs/op
// BenchmarkRenderHTMLandCSS/2c._valid_HTML_input-16   	 7093837	       168.8 ns/op	     168 B/op	       5 allocs/op
// BenchmarkRenderHTMLandCSS/3._valid_CSS_input-16     	 1556341	       770.0 ns/op	    1504 B/op	      16 allocs/op
// BenchmarkRenderHTMLandCSS/4._valid_HTML-CSS_input-16         	 1566343	       765.8 ns/op	    1504 B/op	      16 allocs/op

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
					BoxShadow("0 4px 12px rgba(0,0,0,0.1)").
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
					BoxShadow("0 4px 12px rgba(0,0,0,0.1)").
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

// BenchmarkRenderHTML/1._empty_input-16         	852422221	         1.408 ns/op	       0 B/op	       0 allocs/op
// BenchmarkRenderHTML/2a._valid_HTML_input-16   	10765213	       110.5 ns/op	     104 B/op	       4 allocs/op
// BenchmarkRenderHTML/2b._valid_HTML_input-16   	10093743	       115.2 ns/op	     104 B/op	       4 allocs/op
// BenchmarkRenderHTML/2c._valid_HTML_input-16   	 6892093	       174.1 ns/op	     168 B/op	       5 allocs/op
// BenchmarkRenderHTML/3._valid_CSS_input-16     	17756401	        66.22 ns/op	     144 B/op	       2 allocs/op
// BenchmarkRenderHTML/4._valid_HTML-CSS_input-16         	18267277	        66.51 ns/op	     144 B/op	       2 allocs/op

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
					BoxShadow("0 4px 12px rgba(0,0,0,0.1)").
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
					BoxShadow("0 4px 12px rgba(0,0,0,0.1)").
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

// BenchmarkRenderHTMLWCapacity/1._empty_input-16         	726109297	         1.633 ns/op	       0 B/op	       0 allocs/op
// BenchmarkRenderHTMLWCapacity/2a._valid_HTML_input-16   	15281486	        76.71 ns/op	     176 B/op	       2 allocs/op
// BenchmarkRenderHTMLWCapacity/2b._valid_HTML_input-16   	15013795	        79.47 ns/op	     176 B/op	       2 allocs/op
// BenchmarkRenderHTMLWCapacity/2c._valid_HTML_input-16   	 9801856	       122.6 ns/op	     304 B/op	       2 allocs/op
// BenchmarkRenderHTMLWCapacity/3._valid_CSS_input-16     	17068366	        69.20 ns/op	     144 B/op	       2 allocs/op
// BenchmarkRenderHTMLWCapacity/4._valid_HTML-CSS_input-16         	13251045	        90.34 ns/op	     272 B/op	       3 allocs/op

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
					BoxShadow("0 4px 12px rgba(0,0,0,0.1)").
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
					BoxShadow("0 4px 12px rgba(0,0,0,0.1)").
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
