package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestInputFile(t *testing.T) {

	input := components.NewInputFile()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetError("This is an error!")

	expectations := []string{
		`custom-js-inputfile\s+is-invalid\s+mdl-textfield--floating-label`,
		`<input[^>]+type="text"[^>]+id="id"[^>]+/>`,
		`<input[^>]+readonly="readonly"[^>]+id="id"[^>]+/>`,
		`<input[^>]+id="id"[^>]+class="mdl-textfield__input\s+"[^>]+/>`,
		`<input[^>]+id="id"[^>]+value=""[^>]+/>`,
		`<input\s+name="name"\s+type="file"\s+/>`,
		`<label class="mdl-textfield__label" for="id">label</label>`,
		`<span class="mdl-textfield__error">This is an error!</span>`,
	}

	validateComponent(t, input, expectations)
}
