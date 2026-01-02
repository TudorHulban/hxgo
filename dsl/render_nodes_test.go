package dsl

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func benchNode(w io.Writer) (IsAttribute, Render) {
	return func() bool {
			return false
		},
		func(io.Writer) ([]Style, error) {
			return nil,
				errors.New("some error")
		}
}

// BenchmarkRenderNodes/0._error_input-16         	 4914921	       241.5 ns/op	     128 B/op	       3 allocs/op
// BenchmarkRenderNodes/1._empty_input-16         	 5310886	       225.3 ns/op	     112 B/op	       2 allocs/op
// BenchmarkRenderNodes/2._valid_HTML_input-16    	 3290210	       364.5 ns/op	     128 B/op	       3 allocs/op
// BenchmarkRenderNodes/3._valid_CSS_input-16     	 4227207	       281.9 ns/op	     192 B/op	       3 allocs/op
// BenchmarkRenderNodes/4._valid_HTML-CSS_input-16         	 2795299	       425.7 ns/op	     208 B/op	       4 allocs/op

func BenchmarkRenderNodes(b *testing.B) {
	tests := []struct {
		description string
		elName      string
		nodes       []Node
		wantErr     bool
	}{
		{
			description: "0. error input",
			elName:      "some el",
			nodes:       []Node{benchNode},
			wantErr:     true,
		},
		{
			description: "1. empty input",
			elName:      "",
			nodes:       nil,
		},
		{
			description: "2. valid HTML input",
			elName:      "div",
			nodes: []Node{
				Div(AttrClass("card")),
			},
		},
		{
			description: "3. valid CSS input",
			elName:      "div",
			nodes: []Node{
				Styled(
					[]Style{
						{
							Selector: ".card",
							Props: map[string]string{
								"padding":       "20px",
								"border-radius": "8px",
								"box-shadow":    "0 4px 12px rgba(0,0,0,0.1)",
							},

							Media: "min-width: 768px",
						},
						{
							Selector: ".card:hover",
							Props: map[string]string{
								"box-shadow": "0 8px 24px rgba(0,0,0,0.2)",
							},
						},
					}...,
				),
			},
		},
		{
			description: "4. valid HTML-CSS input",
			elName:      "div",
			nodes: []Node{
				Div(AttrClass("card")),
				Styled(
					[]Style{
						{
							Selector: ".card",
							Props: map[string]string{
								"padding":       "20px",
								"border-radius": "8px",
								"box-shadow":    "0 4px 12px rgba(0,0,0,0.1)",
							},

							Media: "min-width: 768px",
						},
						{
							Selector: ".card:hover",
							Props: map[string]string{
								"box-shadow": "0 8px 24px rgba(0,0,0,0.2)",
							},
						},
					}...,
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
					var buf bytes.Buffer

					_, errRender := renderNodes(&buf, tt.elName, tt.nodes...)
					if tt.wantErr {
						require.Error(b, errRender)

						continue
					}

					require.NoError(b, errRender)
				}
			},
		)
	}
}

// BenchmarkRenderNodesWithCSSId/0._error_input-16         	 3703798	       322.6 ns/op	     144 B/op	       4 allocs/op
// BenchmarkRenderNodesWithCSSId/1._empty_input-16         	 3853051	       311.9 ns/op	     128 B/op	       3 allocs/op
// BenchmarkRenderNodesWithCSSId/2._valid_HTML_input-16    	 2681136	       445.2 ns/op	     144 B/op	       4 allocs/op
// BenchmarkRenderNodesWithCSSId/3._valid_CSS_input-16     	 3231702	       372.3 ns/op	     208 B/op	       4 allocs/op
// BenchmarkRenderNodesWithCSSId/4._valid_HTML-CSS_input-16         	 2386543	       505.6 ns/op	     224 B/op	       5 allocs/op

func BenchmarkRenderNodesWithCSSId(b *testing.B) {
	tests := []struct {
		description string
		elName      string
		nodes       []Node
		wantErr     bool
	}{
		{
			description: "0. error input",
			elName:      "some el",
			nodes:       []Node{benchNode},
			wantErr:     true,
		},
		{
			description: "1. empty input",
			elName:      "",
			nodes:       nil,
		},
		{
			description: "2. valid HTML input",
			elName:      "div",
			nodes: []Node{
				Div(AttrClass("card")),
			},
		},
		{
			description: "3. valid CSS input",
			elName:      "div",
			nodes: []Node{
				Styled(
					[]Style{
						{
							Selector: ".card",
							Props: map[string]string{
								"padding":       "20px",
								"border-radius": "8px",
								"box-shadow":    "0 4px 12px rgba(0,0,0,0.1)",
							},

							Media: "min-width: 768px",
						},
						{
							Selector: ".card:hover",
							Props: map[string]string{
								"box-shadow": "0 8px 24px rgba(0,0,0,0.2)",
							},
						},
					}...,
				),
			},
		},
		{
			description: "4. valid HTML-CSS input",
			elName:      "div",
			nodes: []Node{
				Div(AttrClass("card")),
				Styled(
					[]Style{
						{
							Selector: ".card",
							Props: map[string]string{
								"padding":       "20px",
								"border-radius": "8px",
								"box-shadow":    "0 4px 12px rgba(0,0,0,0.1)",
							},

							Media: "min-width: 768px",
						},
						{
							Selector: ".card:hover",
							Props: map[string]string{
								"box-shadow": "0 8px 24px rgba(0,0,0,0.2)",
							},
						},
					}...,
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
					var buf bytes.Buffer

					_, errRender := renderNodesWithCSSId(&buf, tt.elName, "css ID", tt.nodes...)
					if tt.wantErr {
						require.Error(b, errRender)

						continue
					}

					require.NoError(b, errRender)
				}
			},
		)
	}
}
