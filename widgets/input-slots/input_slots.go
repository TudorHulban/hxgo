package winputslots

import (
	"fmt"

	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxhtml "github.com/TudorHulban/hx-core/html"
	pagecss "github.com/TudorHulban/hx-core/page-css"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

const (
	ButtonSubmitCSSClass = "time-slot"
	ButtonSubmitCSSID    = "submit"
	_SlotPrefix          = "slot"
)

type InfoSlot struct {
	Caption string

	ResourceID int64
	SlotID     int64
}

func (slot *InfoSlot) URL() string {
	return fmt.Sprintf(
		"xxx/%d/%d",

		slot.ResourceID,
		slot.SlotID,
	)
}

func (slot *InfoSlot) CSSID() string {
	return fmt.Sprintf(
		"%s%d-%d",

		_SlotPrefix,
		slot.ResourceID,
		slot.SlotID,
	)
}

type ParamsWidgetSlots struct {
	SubmitEndpoint string
	SlotsInfo      []*InfoSlot
	NumberColumns  uint8
}

func WidgetSlots(params *ParamsWidgetSlots) hxprimitives.Node {
	element := func(slot *InfoSlot) hxprimitives.Node {
		return hxprimitives.Raw(
			hxhelpers.Sprintf(
				`<button class=%s id="%s" type="button" onclick="handletimeclick('%s')">%s</button>`,

				ButtonSubmitCSSClass,
				slot.CSSID(),
				slot.URL(),
				slot.Caption,
			),
		)
	}

	rows := []hxprimitives.Node{
		hxprimitives.Raw(
			fmt.Sprintf(
				`<script>
			function handletimeclick(data){
			const elems = document.querySelectorAll('.%s');
			elems.forEach(element => {
			if (element && element.classList) {element.classList.remove('selected');
    		}
			});

			const parts = data.split('/');
  			const resourceID = parseInt(parts[1], 10);
  			const slotID = parseInt(parts[2], 10);
			const slotCSSID = "%s"+resourceID+"-"+slotID

			const slotElement = document.getElementById(slotCSSID);
			if (slotElement) {slotElement.classList.add('selected');

			const submitElement = document.getElementById('%s')
			if (submitElement) {
			submitElement.disabled = false;
			submitElement.setAttribute('hx-post', '%s/'+resourceID+'/'+slotID);
			};

			console.log('time slot clicked:', data);
  			} else {
			console.log(slotCSSID, resourceID, slotID);
			}
			};
			</script>`,

				ButtonSubmitCSSClass,
				_SlotPrefix,
				ButtonSubmitCSSID,
				params.SubmitEndpoint,
			),
		),
	}

	currentRow := make([]hxprimitives.Node, 0)

	for ix, slot := range params.SlotsInfo {
		currentRow = append(
			currentRow,
			element(slot),
		)

		if (ix+1)%int(params.NumberColumns) == 0 || ix == len(params.SlotsInfo)-1 {
			rows = append(
				rows,
				hxhtml.Div(
					append(
						[]hxprimitives.Node{
							hxprimitives.AttrClass("hours-row"),
						},

						currentRow...,
					)...,
				),
			)

			currentRow = currentRow[:0]
		}
	}

	return hxhtml.Div(
		append(
			[]hxprimitives.Node{
				hxprimitives.AttrClass("hours-grid"),
			},
			rows...,
		)...,
	)
}

func CSSWidgetSlots() *pagecss.CSSElement {
	return &pagecss.CSSElement{
		CSSAllMedias: `
		.hours-grid {
    		width: 40%;
		}

		.hours-row {
    		display: flex;
		}

		.time-slot {
    		flex-grow: 1;
		}

		.selected {
  			background-color: lightblue;
  			border: 1px solid blue;
  			font-weight: bold;
		}
		`,
	}
}
