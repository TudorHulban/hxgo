package components

import (
	"fmt"
	"os"
	"testing"

	"github.com/TudorHulban/hxgo/hx"
	"github.com/stretchr/testify/require"
)

func TestElementSelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass:      "class-div",
		CSSInputID:       "id-input",
		LabelElementName: "label",

		Action: hx.Action{
			Swaps: []string{
				"id1",
			},
		},

		SelectValues: []string{
			"a",
			"b",
		},
	}

	// <div  class="class-div"><label for="id-input">label:</label>
	// <select id="id-input" name="label" hx-swap="#id1" ><option value="a">a</option>
	// <option value="b">b</option></select>
	// </div>

	fnIsAttribute, fnRender := el.Raw()(os.Stdout)
	require.False(t, fnIsAttribute())

	fmt.Println(
		fnRender(os.Stdout),
	)
}

func TestEmptyElementSelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass:      "class-div",
		CSSInputID:       "id-input",
		LabelElementName: "label",

		Action: hx.Action{
			Sends: []string{
				"id1",
			},
		},
	}

	// <div  class="class-div"><label for="id-input">label:</label>
	// <select id="id-input" name="label" hx-send="#id1" ></select>
	// </div>

	fnIsAttribute, fnRender := el.Raw()(os.Stdout)
	require.False(t, fnIsAttribute())

	fmt.Println(
		fnRender(os.Stdout),
	)
}

func TestOnChangeSelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass:      "class-div",
		CSSInputID:       "id-input",
		LabelElementName: "label",

		Action: hx.Action{
			OnChangeEnable: []string{
				"id1",
			},
		},

		SelectValues: []string{
			"a",
			"b",
		},
	}

	// <div  class="class-div"><label for="id-input">label:</label>
	// <select id="id-input" name="label" hx-swap="#id1" ><option value="a">a</option>
	// <option value="b">b</option></select>
	// </div>

	fnIsAttribute, fnRender := el.Raw()(os.Stdout)
	require.False(t, fnIsAttribute())

	fmt.Println(
		fnRender(os.Stdout),
	)
}
