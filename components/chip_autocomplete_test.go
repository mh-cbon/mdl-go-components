package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestChipAutocomplete(t *testing.T) {

	input := components.NewChipAutocomplete()
	input.SetLabel("Type text to complete")
	input.SetError("Wronggg!!!")
	input.SetUrlPlaceholder("itsinput")
	input.SetUrlCompleter("/chip_autocomplete_list?input=itsinput")
	input.SetUrlCreator("/chip_autocomplete_create")

	expectations := []string{
		`<div[^>]+url-placeholder="itsinput"[^>]+class="custom-chipautocomplete custom-js-chipautocomplete\s+"[^>]+>`,
		`<div[^>]+url-completer="/chip_autocomplete_list\?input=itsinput"[^>]+class="custom-chipautocomplete[^"]+"[^>]+>`,
		`<div[^>]+url-creator="/chip_autocomplete_create"[^>]+class="custom-chipautocomplete custom-js-chipautocomplete\s+"[^>]+>`,
		`<div[^>]+id="rnd-0"[^>]+class="custom-chipautocomplete custom-js-chipautocomplete\s+"[^>]+>`,
		`<label class="mdl-textfield__label" for="rnd-1">Type text to complete</label>`,
		`<span class="mdl-textfield__error">Wronggg!!!</span>`,
		`<div\s+class="custom-chipautocomplete-selected\s+"\s+>\s+</div>`,
		`<div\s+class="custom-chipautocomplete-results mdl-shadow--2dp\s+"\s+>\s+<ul class="mdl-list"></ul>\s+</div>`,
		`<input\s+id="rnd-1"\s+class="mdl-textfield__input\s+"\s+type="text"\s+value=""\s+>`,
	}
	validateComponent(t, input, expectations)
}
