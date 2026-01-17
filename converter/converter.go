package converter

import (
	"fmt"
	"strings"

	"github.com/TudorHulban/hxgo/helpers"
	"golang.org/x/net/html"
)

var nodeProcessorMap = map[string]NodeFunc{
	"a":    processAnchor,
	"div":  processDiv,
	"h1":   processH1,
	"h2":   processH2,
	"h3":   processH3,
	"h4":   processH4,
	"h5":   processH5,
	"h6":   processH6,
	"img":  processImg,
	"ol":   processOrderedList,
	"p":    processP,
	"li":   processListItem,
	"nav":  processNav,
	"span": processSpan,
}

func traverseAST(node *html.Node) []string {
	if node == nil {
		return nil
	}

	var result []string

	switch node.Type {
	case html.ElementNode:
		if processFunc, exists := nodeProcessorMap[node.Data]; exists {
			nodeContent := processFunc(node)

			childCount := 0
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				childCount++
			}

			childrenContent := make([]string, 0, childCount)

			// Recursively process children and inject into the current node
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				childrenContent = append(
					childrenContent,
					traverseAST(child)...,
				)
			}

			// inject any children before the closing parentheses
			if len(childrenContent) > 0 {
				// Insert children content into the node string
				// Find the position of the closing parentheses for insertion
				injectionPoint := strings.LastIndex(nodeContent, "\n)")
				if injectionPoint != -1 {
					// Inject the children before the closing parentheses
					nodeContent = nodeContent[:injectionPoint] +
						strings.Join(childrenContent, "") +
						nodeContent[injectionPoint:]
				}
			}

			// Append the node with its children injected
			result = append(result, nodeContent)
		} else {
			// If node type is not handled, traverse children
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				result = append(
					result,
					traverseAST(child)...,
				)
			}
		}

	case html.TextNode:
		replacer := strings.NewReplacer("\n", "", "\t", "")

		if strings.TrimSpace(node.Data) != "" {
			result = append(
				result,
				helpers.Sprintf(
					"\n"+`%s.Text("%s"),`,

					_DSL,
					replacer.Replace(node.Data),
				),
			)
		}

	case html.DocumentNode:
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			result = append(
				result,
				traverseAST(child)...,
			)
		}

	case html.DoctypeNode:
		result = append(
			result,
			helpers.Sprintf(
				"<!DOCTYPE %s>",
				node.Data,
			),
		)

	case html.CommentNode:
		result = append(
			result,
			helpers.Sprintf(
				"<!-- %s -->",
				node.Data,
			),
		)

	default:
		// Non-element nodes can be ignored or logged
		fmt.Printf(
			"Encountered non-element node: %v (%s).\n",

			node.Type,
			node.Data,
		)
	}

	return result
}
