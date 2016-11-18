package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestInputSlice(t *testing.T) {

	input := components.NewInputTextSlice()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetValues([]string{"thevalue", "thevalue2"})
	input.SetError("This is an error!")

	expectations := []string{
		`<div\s+class="mdl-textfield\s+mdl-js-textfield\s+is-invalid\s+mdl-textfield--floating-label\s+">`,
		`<input\s+id="id-0"\s+name="name"\s+class="mdl-textfield__input\s+"\s+type="text"\s+value="thevalue"\s+>`,
		`<label class="mdl-textfield__label" for="id-0">label</label>`,
		`<span class="mdl-textfield__error">This is an error!</span>`,
		`<input\s+id="id-1"\s+name="name"\s+class="mdl-textfield__input\s+"\s+type="text"\s+value="thevalue2"\s+>`,
		`<label class="mdl-textfield__label" for="id-1">label</label>`,
	}

	validateComponent(t, input, expectations)
}
