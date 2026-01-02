package testhtml

import (
	"testing"

	"github.com/TudorHulban/hxgo/dsl"
)

func AssertRaw(t *testing.T, raw string, expectedHTML string) {
	t.Helper()

	gotNode, err := parseHTML(raw)
	if err != nil {
		t.Fatalf("failed to parse raw HTML: %v\nraw:\n%s", err, raw)
	}

	wantNode, err := parseHTML(expectedHTML)
	if err != nil {
		t.Fatalf("failed to parse expected HTML: %v\nexpected:\n%s", err, expectedHTML)
	}

	if diff := compareNodes(gotNode, wantNode); diff != "" {
		t.Fatalf("HTML mismatch:\n%s", diff)
	}
}

func AssertNode(t *testing.T, node dsl.Node, expectedHTML string) {
	t.Helper()

	out := node()
	got := string(out.HTML)

	if got != expectedHTML {
		t.Fatalf("HTML mismatch:\nexpected: %q\ngot:      %q", expectedHTML, got)
	}
}
