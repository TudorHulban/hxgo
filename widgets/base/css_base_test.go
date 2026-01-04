package base

import (
	"testing"

	pagecss "github.com/TudorHulban/hx-core/page-css"
	"github.com/TudorHulban/hxgo/helpers"
	"github.com/stretchr/testify/require"
)

func TestBase(t *testing.T) {
	cssPage := pagecss.NewCSSPage(
		CSSBase,
	)

	writerCSS, errWriterCSS := helpers.GetFileWriter("generated_base.css")
	require.NoError(t, errWriterCSS)

	defer writerCSS.Close()

	cssPage.GetCSSAccurateBeautifiedTo(
		writerCSS,
		&pagecss.ParamsSpaces{
			NumberSpaces:              5,
			IncrementWithNumberSpaces: 2,
		},
	)
}
