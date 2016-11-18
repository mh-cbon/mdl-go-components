package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestSelect(t *testing.T) {

	input := components.NewSelect()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetError("This is an error!")
	input.AddOption("value1", "label1")
	input.AddOption("value2", "label2")
	input.CheckOptions([]string{"value2"})

	expectations := []string{
		`<div\s+class="[^"]*mdl-selectfield mdl-js-selectfield[^"]*">`,
		`<div\s+class="[^"]*is-invalid[^"]*">`,
		`<div\s+class="[^"]*mdl-selectfield--floating-label[^"]*">`,
		`<select[^>]+class="[^"]*mdl-selectfield__select[^"]*"[^>]+>`,
		`<select[^>]+id="id"[^>]+>`,
		`<select[^>]+name="name"[^>]+>`,

		`<option value="value1"\s+>label1</option>`,
		`<option value="value2"\s+selected="selected"\s+>label2</option>`,
		`<label class="mdl-selectfield__label" for="id">label</label>`,
		`<span class="mdl-selectfield__error">This is an error!</span>`,
	}
	validateComponent(t, input, expectations)
}
