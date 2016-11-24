package components_test

import (
	"testing"

	"github.com/mh-cbon/mdl-go-components/components"
)

func TestExpand(t *testing.T) {

	expand := components.NewExpand()
	expand.SetId("id")

	bt := components.NewCheckbox()
	bt.GetOption().SetLabel("labelopt")
	bt.GetOption().SetValue("valueopt")
	bt.SetName("name")
	bt.SetId("id")
	expand.SetBt(bt)

	child := components.NewInputConfirmText()
	child.SetError("This is an error!")

	child.InputLeft.SetName("left-name")
	child.InputLeft.SetId("left-id")
	child.InputLeft.SetLabel("Left label")
	child.InputLeft.SetValue("Left value")

	child.InputRight.SetName("right-name")
	child.InputRight.SetId("right-id")
	child.InputRight.SetLabel("Right label")
	child.InputRight.SetValue("Right value")
	expand.AddComponent(child)

	expectations := []string{
		`<div class="custom-js-expander custom-expander\s+">`,
		`<div class="custom-expander-bt">\s+<label\s+class="mdl-checkbox mdl-js-checkbox "\s+for="id">\s+<input\s+name="name"\s+id="id"\s+type="checkbox"\s+value="valueopt"\s+class="mdl-checkbox__input "\s+/>\s+<span class="mdl-checkbox__label">labelopt</span>\s+</label>\s+</div>`,
		`<div class="custom-expander-container">\s+<div\s+class="mdl-textfield[^"]+">`,
	}
	validateComponent(t, expand, expectations)
}

func TestExpand_Expanded(t *testing.T) {

	expand := components.NewExpand()
	expand.SetId("id")

	bt := components.NewText()
	bt.SetContent("click me!")
	expand.SetBt(bt)

	text := components.NewText()
	text.SetContent("hello content!")
	expand.AddComponent(text)
	expand.AddComponent(text) // yes, twice!

	expectations := []string{
		`<div class="custom-js-expander custom-expander\s+">`,
		`<div class="custom-expander-bt">\s+click me!\s+</div>`,
		`<div class="custom-expander-container">\s+hello content!\s+hello content!\s+</div>`,
	}
	validateComponent(t, expand, expectations)
}
