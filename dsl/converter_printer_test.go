package dsl

import (
	"strings"
	"testing"
	"unsafe"

	"github.com/TudorHulban/hxgo/helpers"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestHTMLToDSL(t *testing.T) {
	tests := []struct {
		description string
		htmlInput   string
		expectedDSL string
	}{
		{
			description: "1. div with text",
			htmlInput:   `<div>Hello</div>`,
			expectedDSL: helpers.Sprintf(
				`%s.Div(%s.Text("Hello"),)`,

				_DSL,
				_DSL,
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.description,
			func(t *testing.T) {
				doc, err := html.Parse(strings.NewReader(tt.htmlInput))
				require.NoError(t, err)

				root := ConvertHTML(doc)
				out := PrintDSL(root)

				require.Equal(t, tt.expectedDSL, out)
			},
		)
	}
}

func TestPrintDSL(t *testing.T) {
	tests := []struct {
		description string
		node        Node
		expectedDSL string
	}{
		{
			description: "1. empty node returns empty string",
			node:        Node{},
			expectedDSL: "",
		},
		{
			description: "2. text node",
			node: Node{
				fn:     renderText,
				data:   unsafe.Pointer(&[]string{"Home"}[0]),
				isText: true,
			},
			expectedDSL: helpers.Sprintf(
				`%s.Text("Home")`,
				_DSL,
			),
		},
		{
			description: "3. attribute node",
			node: Node{
				fn: renderAttribute,
				data: unsafe.Pointer(&struct {
					key string
					val string
				}{
					key: "class",
					val: "breadcrumb-item",
				}),
				isAttribute: true,
			},
			expectedDSL: helpers.Sprintf(
				`%s.Class("breadcrumb-item")`,
				_DSL,
			),
		},
		{
			description: "4. element with attribute and text child",
			node: Node{
				fn:   renderElement,
				data: dataForTag("a"),
				children: []Node{
					{
						fn: renderAttribute,
						data: unsafe.Pointer(&struct {
							key string
							val string
						}{
							key: "href",
							val: "#",
						}),
						isAttribute: true,
					},
					{
						fn:     renderText,
						data:   unsafe.Pointer(&[]string{"Home"}[0]),
						isText: true,
					},
				},
			},
			expectedDSL: helpers.Sprintf(
				`%s.A(%s.Href("#"),%s.Text("Home"),)`,
				_DSL,
				_DSL,
				_DSL,
			),
		},
		{
			description: "5. nested element",
			node: Node{
				fn:   renderElement,
				data: dataForTag("li"),
				children: []Node{
					{
						fn: renderAttribute,
						data: unsafe.Pointer(&struct {
							key string
							val string
						}{
							key: "class",
							val: "breadcrumb-item",
						}),
						isAttribute: true,
					},
					{
						fn:   renderElement,
						data: dataForTag("span"),
						children: []Node{
							{
								fn:     renderText,
								data:   unsafe.Pointer(&[]string{"Text"}[0]),
								isText: true,
							},
						},
					},
				},
			},
			expectedDSL: helpers.Sprintf(
				`%s.Li(%s.Class("breadcrumb-item"),%s.Span(%s.Text("Text"),),)`,
				_DSL,
				_DSL,
				_DSL,
				_DSL,
			),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.description,
			func(t *testing.T) {
				out := PrintDSL(tt.node)
				require.Equal(t, tt.expectedDSL, out)
			},
		)
	}
}
