package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
	"time"
)

func TestDate(t *testing.T) {

	input := components.NewInputDate()
	input.SetLabel("label")
	input.SetName("name")
	input.SetDisplayFormat("ddd DD MMM YYYY")
	input.SetDisplayLocale("fr")
	input.SetDisplayMode("24H")
	input.SetOkText("Ok")
	input.SetCancelText("Cancel")
	input.SetId("id")
	input.SetError("This is an error!")
	if past, err := time.Parse("2006-01-02", "2008-11-11"); err != nil {
		panic(err)
	} else {
		input.SetPast(past.UTC())
	}
	if past, err := time.Parse("2006-01-02", "2024-11-11"); err != nil {
		panic(err)
	} else {
		input.SetFuture(past.UTC())
	}
	if past, err := time.Parse("2006-01-02", "2016-11-11"); err != nil {
		panic(err)
	} else {
		input.SetPresent(past.UTC())
	}
	input.SetGoFormat(time.RFC3339)

	expectations := []string{
		`<div[^>]+class="mdl-textfield\s+mdl-js-textfield\s+custom-js-datefield\s+is-invalid\s+mdl-textfield--floating-label\s*"`,
		`displayformat="ddd DD MMM YYYY"`,
		`displaylocale="fr"`,
		`past="2008-11-11T00:00:00Z"`,
		`future="2024-11-11T00:00:00Z"`,
		`mode="24H"`,
		`ok="Ok"`,
		`cancel="Cancel"`,
		`<input\s+id="id"\s+class="mdl-textfield__input\s+"\s+type="text"\s+>`,
		`<input class="custom-datefield__value" type="hidden"\s+name="name"\s+value="2016-11-11T00:00:00Z"\s+>`,
		`<label class="mdl-textfield__label" for="id">label</label>`,
		`<span class="mdl-textfield__error">This is an error!</span>`,
	}
	validateComponent(t, input, expectations)
}
