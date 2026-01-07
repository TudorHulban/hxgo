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
	)

	el.Add(
		dsl.NewCSSFor("#"+t.Name()).
			WithBreakpoint("768px").
			Background("#4a6cf7").
			Border("none").
			Color("white").
			Cursor("pointer").
			Display("inline-block").
			FontSize("45px").
			Padding("10px 18px").
			Radius("6px").
			Transition("background 0.2s ease, box-shadow 0.2s ease").
			ShadowBox("0 4px 14px rgba(0,0,0,0.15)").
			AsNode(),

		dsl.NewCSSFor("#"+t.Name()+":hover").
			WithBreakpoint("768px").
			Background("#3d5be0").
			ShadowBox("0 4px 14px rgba(0,0,0,0.15)").
			AsNode(),

		dsl.NewCSSFor("#"+t.Name()+":active").
			WithBreakpoint("768px").
			Background("#344fcc").
			ShadowBox("0 2px 8px rgba(0,0,0,0.2)").
			AsNode(),
	)

	html, css := dsl.RenderHTMLandStyles(el)
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

	// http.ListenAndServe(":8080", nil)
}

// BenchmarkButtonSubmit-12    	 2860005	       420.4 ns/op	     432 B/op	       3 allocs/op
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
