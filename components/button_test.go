package components

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/hx"
	"github.com/stretchr/testify/require"
)

func TestButton(t *testing.T) {
	fmt.Println("manual test … skipping")
	t.Skip(
		"manual testing",
	)

	el := ButtonSubmitWCSS(
		&ParamsButtonSubmit{
			Label:            "Press me!",
			HXActionType:     hx.HXPOST,
			HXActionEndpoint: "/",
			CSSID:            t.Name(),
		},
		[]dsl.Node{
			dsl.Styled(
				dsl.Div(), // child node (empty div or whatever wrapper you expect)
				dsl.Style{
					Selector: "#" + t.Name(),
					Props: [][2]string{
						{"display", "inline-block"},
						{"padding", "10px 18px"},
						{"background", "#4a6cf7"},
						{"color", "white"},
						{"border", "none"},
						{"border-radius", "6px"},
						{"font-size", "45px"},
						{"cursor", "pointer"},
						{"transition", "background 0.2s ease, box-shadow 0.2s ease"},
					},
				},
				dsl.Style{
					Selector: "#" + t.Name() + ":hover",
					Props: [][2]string{
						{"background", "#3d5be0"},
						{"box-shadow", "0 4px 14px rgba(0,0,0,0.15)"},
					},
				},
				dsl.Style{
					Selector: "#" + t.Name() + ":active",
					Props: [][2]string{
						{"background", "#344fcc"},
						{"box-shadow", "0 2px 8px rgba(0,0,0,0.2)"},
					},
				},
			),
		},
	)

	html, css := dsl.RenderHTMLandCSS(el)
	require.NotZero(t, html, "valid HTML")
	require.NotZero(t, css, "valid CSS")

	fmt.Println(
		string(html),
	)
	fmt.Println(css)

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(
				[]byte(
					dsl.HTMLwithDataCSS(html, css),
				),
			)
		},
	)

	fmt.Println(
		"Open http://localhost:8080 and press Ctrl-C when done",
	)

	http.ListenAndServe(":8080", nil)
}

func BenchmarkButtonSubmit(b *testing.B) {
	b.ReportAllocs()

	params := &ParamsButtonSubmit{
		CSSID:             "btn-submit",
		CSSClass:          "btn primary",
		Label:             "Submit",
		HXActionType:      "hx-post",
		HXActionEndpoint:  "/submit",
		HXRedirectTo:      "/done",
		HXSwapElements:    []string{"#content"},
		HXRequireElements: []string{"#form"},
		HXSendElements:    []string{"#input"},
		HXEnableElements:  []string{"#ok"},
		HXDisableElements: []string{"#no"},
		IsHXUpload:        false,
		IsDisabled:        false,
		IsNewTab:          false,
	}

	b.ResetTimer()

	el := ButtonSubmit(params)

	for b.Loop() {
		_ = dsl.RenderHTMLWithCapacity(128, el)
	}
}
