package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestCheckbox(t *testing.T) {

	input := components.NewCheckbox()
	input.GetOption().SetLabel("labelopt")
	input.GetOption().SetValue("valueopt")
	input.SetName("name")
	input.SetId("id")
	input.SetChecked(true)

	expectations := []string{
		`<label[^>]+class="[^"]*mdl-checkbox mdl-js-checkbox[^"]*"[^>]+for="id">`,
		`<input[^>]+id="id"[^>]+type="checkbox"[^>]+>`,
		`<input[^>]+name="name"[^>]+type="checkbox"[^>]+>`,
		`<input[^>]+type="checkbox"[^>]+checked="checked"[^>]+>`,
		`<input[^>]+type="checkbox"[^>]+value="valueopt"[^>]+>`,
		`<span class="mdl-checkbox__label">labelopt</span>`,
	}
	validateComponent(t, input, expectations)
}

func TestRadio(t *testing.T) {

	input := components.NewRadio()
	input.GetOption().SetLabel("labelopt")
	input.GetOption().SetValue("valueopt")
	input.SetName("name")
	input.SetId("id")
	input.SetChecked(true)

	expectations := []string{
		`<label[^>]+class="[^"]*mdl-radio mdl-js-radio[^"]*"[^>]+for="id">`,
		`<input[^>]+id="id"[^>]+type="radio"[^>]+>`,
		`<input[^>]+name="name"[^>]+type="radio"[^>]+>`,
		`<input[^>]+type="radio"[^>]+checked="checked"[^>]+>`,
		`<input[^>]+type="radio"[^>]+value="valueopt"[^>]+>`,
		`<span class="mdl-radio__label">labelopt</span>`,
	}
	validateComponent(t, input, expectations)
}
