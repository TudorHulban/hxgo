package components

import (
	"testing"

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

	output := el.Raw()

	out := output()
	require.False(t, out.IsAttr)
	require.Zero(t, out.HTML)

	// expectedOutput := helpers.Sprintf(
	// 	`<div class="class-div"><label>%s:</label><input type="text"  name="label" hx-max=50></div>`,
	// 	el.LabelElementName,
	// )
}
