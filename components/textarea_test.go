package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestTextarea(t *testing.T) {

	input := components.NewTextarea()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetValue("thevalue")
	input.SetError("This is an error!")

	expectations := []string{
		`<div class="mdl-textfield mdl-js-textfield\s+is-invalid\s+mdl-textfield--floating-label">`,
		`<textarea class="mdl-textfield__input"\s+id="id"\s+name="name"\s+rows= "3">thevalue</textarea>`,
		`<label class="mdl-textfield__label" for="id">label</label>`,
		`<span class="mdl-textfield__error">This is an error!</span>`,
	}

	validateComponent(t, input, expectations)
}

func TestWysiwyg(t *testing.T) {

	input := components.NewWysiwyg()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetValue("thevalue")
	input.SetError("This is an error!")

	expectations := []string{
		`<div class="mdl-textfield mdl-js-textfield custom-tinymce custom-js-tinymce\s+is-invalid\s+mdl-textfield--floating-label">`,
		`<textarea class="mdl-textfield__input"\s+id="id"\s+name="name"\s+rows="3"\s+>thevalue</textarea>`,
		`<label class="mdl-textfield__label" for="id">label</label>`,
		`<span class="mdl-textfield__error">This is an error!</span>`,
	}

	validateComponent(t, input, expectations)
}
