package dsl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestHowToUse(t *testing.T) {
	htmlData := `
    <li class="breadcrumb-item">
        <a rel="noopener noreferrer" href="#" title="Back to homepage" class="breadcrumb-link">Home</a>
    </li>
    <li class="breadcrumb-item">
        <span class="breadcrumb-separator">Text</span>
        <a rel="noopener noreferrer" href="#" class="breadcrumb-link">Parent</a>
    </li>
    `

	doc, errParse := html.Parse(strings.NewReader(htmlData))
	require.NoError(t, errParse)
	require.NotNil(t, doc)
	require.NotNil(t, doc.FirstChild)

	// NEW: convert HTML DOM → hxgo.Node tree
	root := ConvertHTML(doc)
	require.NotNil(t, root.fn)

	// NEW: render hxgo.Node → HTML
	out := Render(root)
	require.NotEmpty(t, out)

	fmt.Println(string(out))
}
