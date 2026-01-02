package components

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/TudorHulban/hxgo/helpers"
	"github.com/TudorHulban/hxgo/hx"
	"github.com/stretchr/testify/require"
)

func TestNoIDElementSimpleInput(t *testing.T) {
	el := InputSimple{
		CSSDivClass:      "class-div",
		LabelElementName: "label",

		TypeInput: "text",
		// TextInput: "some text to show",

		Action: hx.Action{
			LengthMax: 50,
		},
	}

	// <div  class="class-div">`<label>`+el.LabelElementName+`:</label>`
	// <input type="text"  name="label" hx-max=50></div>

	fnIsAttribute, fnRender := el.Raw()(os.Stdout)
	require.False(t, fnIsAttribute())

	fmt.Println(
		fnRender(os.Stdout),
	)

	expectedOutput := helpers.Sprintf(
		`<div class="class-div"><label>%s:</label><input type="text"  name="label" hx-max=50></div>`,
		el.LabelElementName,
	)

	var buf bytes.Buffer
	styles, err := fnRender(&buf)
	require.NoError(t, err, "render error")
	require.Zero(t, styles)

	actualOutput := buf.String()
	require.Equal(t, expectedOutput, actualOutput, "unexpected output")
}
