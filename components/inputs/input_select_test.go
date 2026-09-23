package inputs

import (
	"testing"

	"github.com/TudorHulban/hxgo/hx"
	"github.com/stretchr/testify/require"
)

func Test_Element_SelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass:      "class-div",
		CSSInputID:       "id-input",
		LabelElementName: "label-value",

		Action: hx.Action{
			Swaps: []string{
				"id1",
			},
		},

		SelectOptions: []Option{
			{
				Value: "a1",
				Label: "a",
			},
			{
				Value: "b2",
				Label: "b",
			},
		},
	}

	// <div  class="class-div"><label for="id-input">label-value:</label>
	// <select id="id-input" name="label" hx-swap="#id1" ><option value="a">a</option>
	// <option value="b">b</option></select>
	// </div>

	output := el.Raw()
	c := output.Canonical()

	// fmt.Println(
	// 	"elem:",
	// 	c,
	// )

	require.Contains(t,
		c,
		"label",
	)
}

func Test_NoLabel_Element_SelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass: "class-div",
		CSSInputID:  "id-input",

		Action: hx.Action{
			Swaps: []string{
				"id1",
			},
		},

		SelectOptions: []Option{
			{
				Value: "a1",
				Label: "a",
			},
			{
				Value: "b2",
				Label: "b",
			},
		},
	}

	output := el.Raw()

	require.NotContains(t,
		output.Canonical(),
		"label",
	)
}

func Test_Empty_Element_SelectInput(t *testing.T) {
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

	output := el.Raw()

	// fmt.Println(
	// 	"elem:",
	// 	output.Canonical(),
	// )

	require.NotNil(t, output)
}

func Test_OnChange_Element_SelectInput(t *testing.T) {
	el := InputSelect{
		CSSDivClass:      "class-div",
		CSSInputID:       "id-input",
		LabelElementName: "label",

		Action: hx.Action{
			OnChangeEnable: []string{
				"id1",
			},
		},

		SelectOptions: []Option{
			{
				Value: "a1",
				Label: "a",
			},
			{
				Value: "b2",
				Label: "b",
			},
		},
	}

	// <div  class="class-div"><label for="id-input">label:</label>
	// <select id="id-input" name="label" hx-swap="#id1" ><option value="a">a</option>
	// <option value="b">b</option></select>
	// </div>

	output := el.Raw()

	// fmt.Println(
	// 	"elem:",
	// 	output.Canonical(),
	// )

	require.NotNil(t, output)
}
