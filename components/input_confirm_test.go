package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestInputConfirm(t *testing.T) {

	input := components.NewInputConfirmText()
	input.SetError("This is an error!")

	input.InputLeft.SetName("left-name")
	input.InputLeft.SetId("left-id")
	input.InputLeft.SetLabel("Left label")
	input.InputLeft.SetValue("Left value")

	input.InputRight.SetName("right-name")
	input.InputRight.SetId("right-id")
	input.InputRight.SetLabel("Right label")
	input.InputRight.SetValue("Right value")

	expectations := []string{
		`<div\s+class="mdl-textfield\s+custom-input-confirm\s+custom-js-input-confirm\s+is-invalid\s+">`,
		`<input\s+name="left-name"\s+id="left-id"\s+class="mdl-textfield__input "\s+type="text"\s+value="Left value"\s+>`,
		`<label\s+class="mdl-textfield__label"\s+for="left-id"\s+>Left label</label>`,
		`<input\s+name="right-name"\s+id="right-id"\s+class="mdl-textfield__input "\s+type="text"\s+value="Right value"\s+>`,
		`<label\s+class="mdl-textfield__label"\s+for="right-id"\s+>Right label</label>`,
		`<span class="mdl-textfield__error">This is an error!</span>`,
	}
	validateComponent(t, input, expectations)
}
