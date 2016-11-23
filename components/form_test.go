package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestForm(t *testing.T) {

	form := components.NewForm()

	form.AddText("name")
	form.AddFile("file")

	expectations := []string{
		`<form\s+id="rnd-0"\s+class="">`,
		`<input[^>]+name="name"[^>]+id="rnd-1"[^>]+>`,
		`<input[^>]+id="rnd-1"[^>]+class="mdl-textfield__input\s+"[^>]+>`,
		`<input[^>]+id="rnd-1"[^>]+type="text"[^>]+>`,
		`<input[^>]+id="rnd-1"[^>]+value=""[^>]+>`,

		`<input[^>]+type="text"[^>]+id="rnd-2"[^>]+/>`,
		`<input[^>]+readonly="readonly"[^>]+id="rnd-2"[^>]+/>`,
		`<input[^>]+id="rnd-2"[^>]+class="mdl-textfield__input\s+"[^>]+/>`,
		`<input[^>]+id="rnd-2"[^>]+value=""[^>]+/>`,

	}
	validateComponent(t, form, expectations)

}

func TestFormAttr(t *testing.T) {

	form := components.NewForm()

	form.SetAction("http://whatever/")
	form.SetMethod("POST")
	form.SetTarget("_blank")

	expectations := []string{
		`<form[^>]+method="POST"[^>]+>`,
		`<form[^>]+action="http://whatever/"[^>]+>`,
		`<form[^>]+target="_blank"[^>]+>`,
	}
	validateComponent(t, form, expectations)

}
