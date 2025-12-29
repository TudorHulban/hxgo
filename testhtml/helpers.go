package testhtml

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func parseHTML(s string) (*html.Node, error) {
	// The parser expects a full HTML document, so we wrap the fragment
	r := strings.NewReader(s)
	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	// html.Parse always returns a full document:
	// <html><head></head><body> ... </body></html>
	// We want the <body> children (the actual fragment)
	return findBody(node), nil
}

func findBody(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "body" {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if b := findBody(c); b != nil {
			return b
		}
	}
	return nil
}

func compareAttributes(got, want []html.Attribute, path string) string {
	normalize := func(attrs []html.Attribute) map[string]string {
		m := make(map[string]string)
		for _, a := range attrs {
			m[a.Key] = strings.TrimSpace(a.Val)
		}

		return m
	}

	gm := normalize(got)
	wm := normalize(want)

	if len(gm) != len(wm) {
		return fmt.Sprintf(
			"%s: attribute count mismatch: got %v, want %v", path, gm, wm,
		)
	}

	for k, gv := range gm {
		wv, ok := wm[k]
		if !ok {
			return fmt.Sprintf(
				"%s: missing attribute %q", path, k,
			)
		}
		if gv != wv {
			return fmt.Sprintf(
				"%s: attribute %q mismatch: got %q, want %q", path, k, gv, wv,
			)
		}
	}

	return ""
}

func compareNodes(got, want *html.Node) string {
	return compareNodeRecursive(got, want, "")
}

func filterMeaningful(n *html.Node) []*html.Node {
	var out []*html.Node

	for c := n; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode && strings.TrimSpace(c.Data) == "" {
			continue
		}
		out = append(out, c)
	}

	return out
}

func compareNodeRecursive(got, want *html.Node, path string) string {
	// Skip over document nodes and go directly to children
	if got.Type == html.DocumentNode {
		got = got.FirstChild
	}
	if want.Type == html.DocumentNode {
		want = want.FirstChild
	}

	// Compare node types
	if got.Type != want.Type {
		return fmt.Sprintf("%s: node type mismatch: got %v, want %v", path, got.Type, want.Type)
	}

	// Compare element names
	if got.Type == html.ElementNode && got.Data != want.Data {
		return fmt.Sprintf("%s: tag mismatch: got <%s>, want <%s>", path, got.Data, want.Data)
	}

	// Compare attributes (normalized)
	if diff := compareAttributes(got.Attr, want.Attr, path); diff != "" {
		return diff
	}

	// Compare text nodes
	if got.Type == html.TextNode {
		gt := strings.TrimSpace(got.Data)
		wt := strings.TrimSpace(want.Data)
		if gt != wt {
			return fmt.Sprintf("%s: text mismatch: got %q, want %q", path, gt, wt)
		}
		return ""
	}

	// Compare children count
	gotChildren := filterMeaningful(got.FirstChild)
	wantChildren := filterMeaningful(want.FirstChild)

	if len(gotChildren) != len(wantChildren) {
		return fmt.Sprintf(
			"%s: children count mismatch: got %d, want %d",
			path, len(gotChildren), len(wantChildren),
		)
	}

	// Compare children recursively
	for i := range gotChildren {
		childPath := fmt.Sprintf("%s/%s[%d]", path, got.Data, i)
		if diff := compareNodeRecursive(gotChildren[i], wantChildren[i], childPath); diff != "" {
			return diff
		}
	}

	return ""
}
