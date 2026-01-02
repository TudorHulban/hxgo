package dsl

import (
	"bytes"
	"database/sql"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

type benchNode struct{}

func (n benchNode) isAttribute() bool { return false }

func (n benchNode) Render(w io.Writer) ([]Style, error) {
	return nil,
		errors.New("some error")
}

// BenchmarkRenderNodes/0._error_input-16         	 4707579	       253.1 ns/op	     129 B/op	       4 allocs/op
// BenchmarkRenderNodes/1._empty_input-16         	 8272828	       143.7 ns/op	      48 B/op	       1 allocs/op
// BenchmarkRenderNodes/2._valid_HTML_input-16    	 2955564	       406.2 ns/op	     130 B/op	       5 allocs/op
// BenchmarkRenderNodes/3._valid_CSS_input-16     	 3859956	       310.8 ns/op	     193 B/op	       4 allocs/op
// BenchmarkRenderNodes/4._valid_HTML-CSS_input-16         	 2542068	       470.8 ns/op	     210 B/op	       6 allocs/op

func BenchmarkRenderNodes(b *testing.B) {
	tests := []struct {
		description string
		elName      sql.NullString
		nodes       []Node
		wantErr     bool
	}{
		{
			description: "0. error input",
			elName:      sql.NullString{String: "some el", Valid: true},
			nodes:       []Node{benchNode{}},
			wantErr:     true,
		},
		{
			description: "1. empty input",
			elName:      sql.NullString{Valid: false},
			nodes:       nil,
		},
		{
			description: "2. valid HTML input",
			elName:      sql.NullString{String: "div", Valid: true},
			nodes: []Node{
				Div(AttrClass("card")),
			},
		},
		{
			description: "3. valid CSS input",
			elName:      sql.NullString{String: "div", Valid: true},
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
			elName:      sql.NullString{String: "div", Valid: true},
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
