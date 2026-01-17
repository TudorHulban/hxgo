package converter

import "golang.org/x/net/html"

func extractAttributes(node *html.Node) map[string]string {
	attributes := make(map[string]string)

	for _, attr := range node.Attr {
		attributes[attr.Key] = attr.Val
	}

	return attributes
}
