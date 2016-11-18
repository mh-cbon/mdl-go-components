package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestCheckboxSlice(t *testing.T) {
	input := components.NewCheckboxSlice()
	input.SetName("name")
	input.SetId("id")
	input.SetValue("value2")
	input.AddOption("valueopt", "labelopt")
	input.AddOption("value2", "label2")

	expectations := []string{
		`<label[^>]+class="[^"]*mdl-checkbox mdl-js-checkbox[^"]*"[^>]+>`,

		`<label[^>]+for="id-0">`,
		`<input[^>]+id="id-0"[^>]+name="name"[^>]+type="checkbox"[^>]+>`,
		`<input[^>]+type="checkbox"[^>]+value="valueopt"[^>]+>`,
		`<input[^>]+type="checkbox"[^>]+class="[^"]*mdl-checkbox__input[^"]*"[^>]+>`,
		`<span class="mdl-checkbox__label">labelopt</span>`,

		`<label[^>]+for="id-1">`,
		`<input[^>]+id="id-1"[^>]+name="name"[^>]+type="checkbox"[^>]+>`,
		`<input[^>]+type="checkbox"[^>]+value="value2"[^>]+>`,
		`<input[^>]+type="checkbox"[^>]+class="[^"]*mdl-checkbox__input[^"]*"[^>]+>`,
		`<span class="mdl-checkbox__label">label2</span>`,

		`<span class="mdl-textfield__error"></span>`,
	}
	validateComponent(t, input, expectations)
}

func TestCheckboxSliceWithError(t *testing.T) {
	input := components.NewCheckboxSlice()
	input.SetName("name")
	input.SetId("id")
	input.SetError("This is an error!")
	input.SetValue("value2")
	input.AddOption("valueopt", "labelopt")
	input.AddOption("value2", "label2")

	expectations := []string{
		`<label[^>]+class="[^"]*mdl-checkbox mdl-js-checkbox[^"]*"[^>]+>`,

		`<label[^>]+for="id-0">`,
		`<input[^>]+id="id-0"[^>]+name="name"[^>]+type="checkbox"[^>]+>`,
		`<input[^>]+type="checkbox"[^>]+value="valueopt"[^>]+>`,
		`<input[^>]+type="checkbox"[^>]+class="[^"]*mdl-checkbox__input[^"]*"[^>]+>`,
		`<span class="mdl-checkbox__label">labelopt</span>`,

		`<label[^>]+for="id-1">`,
		`<input[^>]+id="id-1"[^>]+name="name"[^>]+type="checkbox"[^>]+>`,
		`<input[^>]+type="checkbox"[^>]+value="value2"[^>]+>`,
		`<input[^>]+type="checkbox"[^>]+class="[^"]*mdl-checkbox__input[^"]*"[^>]+>`,
		`<span class="mdl-checkbox__label">label2</span>`,

		`<span class="mdl-textfield__error">This is an error!</span>`,
	}
	validateComponent(t, input, expectations)
}

func TestCheckboxSliceWithValues(t *testing.T) {
	input := components.NewCheckboxSlice()
	input.SetName("name")
	input.SetId("id")
	input.SetError("This is an error!")
	input.SetValue("value2")
	input.AddOption("valueopt", "labelopt")
	input.AddOption("value2", "label2")
	input.CheckOptions([]string{"value2"})

	expectations := []string{
		`<input[^>]+type="checkbox"[^>]+value="valueopt"\s+class="mdl-checkbox__input\s+"\s+/>`,
		`<input[^>]+type="checkbox"[^>]+value="value2"[^>]+checked="checked"[^>]+>`,

	}
	validateComponent(t, input, expectations)
}

func TestRadioSlice(t *testing.T) {

	input := components.NewRadioSlice()
	input.SetName("name")
	input.SetId("id")
	input.SetValue("value2")
	input.AddOption("valueopt", "labelopt")
	input.AddOption("value2", "label2")

	expectations := []string{
		`<label[^>]+class="[^"]*mdl-radio mdl-js-radio[^"]*"[^>]+>`,

		`<label[^>]+for="id-0">`,
		`<input[^>]+id="id-0"[^>]+name="name"[^>]+type="radio"[^>]+>`,
		`<input[^>]+type="radio"[^>]+value="valueopt"[^>]+>`,
		`<input[^>]+type="radio"[^>]+class="[^"]*mdl-radio__button[^"]*"[^>]+>`,
		`<span class="mdl-radio__label">labelopt</span>`,

		`<label[^>]+for="id-1">`,
		`<input[^>]+id="id-1"[^>]+name="name"[^>]+type="radio"[^>]+>`,
		`<input[^>]+type="radio"[^>]+value="value2"[^>]+>`,
		`<input[^>]+type="radio"[^>]+class="[^"]*mdl-radio__button[^"]*"[^>]+>`,
		`<span class="mdl-radio__label">label2</span>`,

		`<span class="mdl-textfield__error"></span>`,
	}
	validateComponent(t, input, expectations)
}

func TestRadioSliceWithError(t *testing.T) {

	input := components.NewRadioSlice()
	input.SetName("name")
	input.SetId("id")
	input.SetError("This is an error!")
	input.SetValue("value2")
	input.AddOption("valueopt", "labelopt")
	input.AddOption("value2", "label2")

	expectations := []string{
		`<label[^>]+class="[^"]*mdl-radio mdl-js-radio[^"]*"[^>]+>`,

		`<label[^>]+for="id-0">`,
		`<input[^>]+id="id-0"[^>]+name="name"[^>]+type="radio"[^>]+>`,
		`<input[^>]+type="radio"[^>]+value="valueopt"[^>]+>`,
		`<input[^>]+type="radio"[^>]+class="[^"]*mdl-radio__button[^"]*"[^>]+>`,
		`<span class="mdl-radio__label">labelopt</span>`,

		`<label[^>]+for="id-1">`,
		`<input[^>]+id="id-1"[^>]+name="name"[^>]+type="radio"[^>]+>`,
		`<input[^>]+type="radio"[^>]+value="value2"[^>]+>`,
		`<input[^>]+type="radio"[^>]+class="[^"]*mdl-radio__button[^"]*"[^>]+>`,
		`<span class="mdl-radio__label">label2</span>`,

		`<span class="mdl-textfield__error">This is an error!</span>`,
	}
	validateComponent(t, input, expectations)
}

func TestRadioSliceWithValues(t *testing.T) {

	input := components.NewRadioSlice()
	input.SetName("name")
	input.SetId("id")
	input.SetError("This is an error!")
	input.SetValue("value2")
	input.AddOption("valueopt", "labelopt")
	input.AddOption("value2", "label2")
	input.CheckOptions([]string{"value2"})

	expectations := []string{
		`<input[^>]+type="radio"[^>]+value="valueopt"\s+class="mdl-radio__button\s+"\s+/>`,
		`<input[^>]+type="radio"[^>]+value="value2"[^>]+checked="checked"[^>]+>`,
	}
	validateComponent(t, input, expectations)
}
