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
		`<input\s+name="name"\s+id="rnd-1"\s+class="mdl-textfield__input\s+"\s+type="text"\s+value=""\s+>`,
		`<input\s+id="rnd-2"\s+class="mdl-textfield__input\s+"\s+type="text"\s+value=""\s+readonly="readonly"\s+/>`,
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
