package base

import (
	"testing"

	"github.com/TudorHulban/hxgo/dsl"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/stretchr/testify/require"
)

func TestBase(t *testing.T) {
	cssContributionBase := dsl.CSSContribution{
		ProceduralCSS: []dsl.CSS{
			CSSBase,
		},
	}

	el := cssContributionBase.AsNode()

	html, styles, css := dsl.RenderFull(el)
	require.Zero(t, html)
	require.Zero(t, styles)
	require.NotZero(t, css)

	writerCSS, errWriterCSS := helpers.GetFileWriter("generated_base.css")
	require.NoError(t, errWriterCSS)

	defer writerCSS.Close()

	writerCSS.Write(
		[]byte(css),
	)
}
