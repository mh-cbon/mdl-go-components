package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestInputText(t *testing.T) {

	input := components.NewInputText()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetValue("thevalue")
	input.SetError("This is an error!")

	expectations := []string{
		`mdl-js-textfield\s+is-invalid\s+mdl-textfield--floating-label`,
		`<label class="mdl-textfield__label" for="id">label</label>`,
		`<span class="mdl-textfield__error">This is an error!</span>`,
		`<input\s+name="name"\s+id="id"\s+class="mdl-textfield__input\s+"\s+type="text"\s+value="thevalue"\s+>`,
	}

	validateComponent(t, input, expectations)
}

func TestPassword(t *testing.T) {

	input := components.NewInputPassword()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetValue("thevalue")
	input.SetError("This is an error!")

	expectations := []string{
		`mdl-js-textfield\s+is-invalid\s+mdl-textfield--floating-label`,
		`<label class="mdl-textfield__label" for="id">label</label>`,
		`<span class="mdl-textfield__error">This is an error!</span>`,
		`<input\s+name="name"\s+id="id"\s+class="mdl-textfield__input\s+"\s+type="password"\s+value="thevalue"\s+>`,
	}

	validateComponent(t, input, expectations)
}

func TestHidden(t *testing.T) {

	input := components.NewInputHidden()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetValue("thevalue")
	input.SetError("This is an error!")

	expectations := []string{
		`<input\s+class=""\s+type="hidden"\s+name="name"\s+value="thevalue"\s+/>`,
	}

	validateComponent(t, input, expectations)
}
