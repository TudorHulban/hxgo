package dsl

import (
	"context"
	"net/http"
	"os"
	"strings"
	"testing"
	"unsafe"

	"github.com/TudorHulban/hxgo/helpers"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestHTMLToDSL(t *testing.T) {
	testCases := []struct {
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

	for _, tc := range testCases {
		t.Run(
			tc.description,
			func(t *testing.T) {
				doc, errParse := html.Parse(strings.NewReader(tc.htmlInput))
				require.NoError(t, errParse)

				root := ConvertHTML(doc)

				var out strings.Builder

				printDSL(&out, root)

				require.Equal(t,
					tc.expectedDSL,
					out.String(),
				)
			},
		)
	}
}

func TestPrintDSL(t *testing.T) {
	tests := []struct {
		description string
		expectedDSL string

		node Node
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
				var out strings.Builder

				printDSL(&out, tt.node)

				require.Equal(t,
					tt.expectedDSL,
					out.String(),
				)
			},
		)
	}
}

func TestFullPageWithURL(t *testing.T) {
	const testURL = "https://templates.iqonic.design/product/lite/logik/html/dist/dashboard/"

	req, errRequest := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		testURL,
		nil,
	)
	require.NoError(t, errRequest)

	resp, errCall := http.DefaultClient.Do(req)
	require.NoError(t, errCall)

	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	doc, errParse := html.Parse(resp.Body)
	require.NoError(t, errParse)
	require.NotNil(t, doc)
	require.NotEmpty(t, doc)

	root := ConvertHTML(doc)

	PrintDSL(os.Stdout, root)
}

func TestElementsWithURL(t *testing.T) {
	const testURL = "https://templates.iqonic.design/product/lite/logik/html/dist/dashboard/"

	req, errRequest := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		testURL,
		nil,
	)
	require.NoError(t, errRequest)

	resp, errCall := http.DefaultClient.Do(req)
	require.NoError(t, errCall)

	defer resp.Body.Close()

	require.Equal(t,
		http.StatusOK,
		resp.StatusCode,
	)

	doc, errParse := html.Parse(resp.Body)
	require.NoError(t, errParse)
	require.NotNil(t, doc)
	require.NotEmpty(t, doc)

	elements := ConvertDOMElements(
		doc,
		DOMNode{
			CSSClass: "card-body",
		},
	)

	elements.PrintWithDescription(os.Stdout)
}
