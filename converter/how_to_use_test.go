package converter

import (
	"context"
	"fmt"
	"net/http"
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
	require.True(t,
		doc.FirstChild != nil,
	)

	nodes := traverseAST(doc)
	require.NotEmpty(t, nodes)

	fmt.Println(
		nodes,
	)
}

func TestWithURL(t *testing.T) {
	testURL := "https://tailwind-nextjs-starter-blog.vercel.app/"

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

	nodes := traverseAST(doc)
	require.NotEmpty(t, nodes)

	fmt.Println(
		nodes,
	)
}
