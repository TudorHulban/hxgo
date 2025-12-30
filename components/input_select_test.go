package components

import (
	"fmt"
	"os"
	"testing"

	"github.com/TudorHulban/hxgo/hx"
)

func TestElementSelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass:      "class-div",
		CSSInputID:       "id-input",
		LabelElementName: "label",

		HX: hx.HXAction{
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

	fmt.Println(
		el.
			Raw().
			Render(os.Stdout),
	)
}

func TestEmptyElementSelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass:      "class-div",
		CSSInputID:       "id-input",
		LabelElementName: "label",

		HX: hx.HXAction{
			Sends: []string{
				"id1",
			},
		},
	}

	// <div  class="class-div"><label for="id-input">label:</label>
	// <select id="id-input" name="label" hx-send="#id1" ></select>
	// </div>

	fmt.Println(
		el.
			Raw().
			Render(os.Stdout),
	)
}

func TestOnChangeSelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass:      "class-div",
		CSSInputID:       "id-input",
		LabelElementName: "label",

		HX: hx.HXAction{
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

	fmt.Println(
		el.
			Raw().
			Render(os.Stdout),
	)
}
