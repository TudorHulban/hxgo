package inputs

import (
	"fmt"
	"testing"
	"time"

	"github.com/TudorHulban/hxgo/dsl"
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

	output := el.Build()
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

	output := el.Build()

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

	output := el.Build()

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

	output := el.Build()

	// fmt.Println(
	// 	"elem:",
	// 	output.Canonical(),
	// )

	require.NotNil(t, output)
}

type ComboSelectMonth struct {
	SelectedValue string
	OptionsMonth  []Option
}

func NewComboSelectMonth(t time.Time, monthsBehind, monthsAhead int) ComboSelectMonth {
	selectedValue := fmt.Sprintf("%d-%d", t.Year(), t.Month()) // Format as ("YYYY-M").

	var options []Option

	// Loop from negative (behind) to positive (ahead)
	for ix := -monthsBehind; ix <= monthsAhead; ix++ {
		targetTime := t.AddDate(0, ix, 0)

		options = append(
			options,
			Option{
				Value: fmt.Sprintf(
					"%d-%d",
					targetTime.Year(), targetTime.Month(),
				),
				Label: targetTime.Format("Jan 2006"), // e.g., "Apr 2026", "May 2026"
			},
		)
	}

	return ComboSelectMonth{
		SelectedValue: selectedValue,
		OptionsMonth:  options,
	}
}

func (elem ComboSelectMonth) Build() *dsl.Node {
	if len(elem.OptionsMonth) == 0 {
		return nil
	}

	node := dsl.Div(
		dsl.AttrClass("selector-group"),
		dsl.AttrID("some-ID"),

		InputSelect{
			CSSInputID: "other-ID",

			SelectedValue: elem.SelectedValue,
			SelectOptions: elem.OptionsMonth,
		}.
			Build(),
	)

	return &node
}

func Test_Combo_With_SelectInput(t *testing.T) {
	combo := NewComboSelectMonth(
		time.Now(),
		3,
		3,
	)
	require.NotZero(t, combo)

	node := combo.Build()
	require.NotZero(t, node)

	// attrs, kids := node.DebugParts()
	// for i, a := range attrs {
	// 	fmt.Println("attr", i, a.IsText, a.IsAttribute, a.IsCSS, a.Data != nil)
	// }

	// for i, k := range kids {
	// 	fmt.Println("kid", i, k.IsText, k.IsAttribute, k.IsCSS, k.Data != nil)
	// }

	// fmt.Println("root:", fnName(*node))

	// for i, k := range kids {
	// 	fmt.Println("kid", i, fnName(k))
	// }

	require.NotZero(t, combo.Build().Canonical())
}
