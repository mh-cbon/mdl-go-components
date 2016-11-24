package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
)

func ChipAutocomplete(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Chip autocomplete",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var line *components.Slice
	var t1 *components.ChipAutocomplete

	// -
	line = components.NewSlice()

	t1 = components.NewChipAutocomplete()
	t1.SetLabel("Type text to complete")
	t1.SetError("Wronggg!!!")
	t1.SetUrlPlaceholder("itsinput")
	t1.SetUrlCompleter("/chip_autocomplete_list?input=itsinput")
	t1.SetUrlCreator("/chip_autocomplete_create")

	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewChipAutocomplete()
	t1.SetLabel("Type text to complete")
	t1.SetUrlCompleter("/chip_autocomplete_list?input=itsinput")
	t1.SetUrlPlaceholder("itsinput")

	t1.AddOption("value", "text")
	t1.AddOption("value2", "text 2")

	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}

var chipResults = []string{
	"ada",
	"golang",
	"java",
	"php",
	"scala",
	"javascript",
	"c",
	"c++",
	"c#",
	"pascal",
	"asm",
	"lisp",
	"haskell",
}

type ChipResult struct {
	Value string
	Text  string
}

func ChipAutocompleteList(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	limits := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(limits)
	results := []ChipResult{}
	for i, v := range chipResults {
		if len(v) >= len(input) && v[0:len(input)] == input {
			results = append(results, ChipResult{strconv.Itoa(i), v})
		}
	}
	if limit > 0 {
		results = results[0:limit]
	}
	blob, _ := json.Marshal(results)
	w.Write(blob)
}

type ChipCreateResponse struct {
	Data       ChipResult
	Valid      bool
	HasFailure bool
	Failure    string
}

func ChipAutocompleteCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	value := r.Form.Get("Value")

	res := ChipCreateResponse{}
	if value == "" {
		res.Valid = false
		res.HasFailure = true
		res.Failure = "Need a value"
	} else if value == "error" {
		res.Valid = false
		res.HasFailure = true
		res.Failure = "beep boop, demo error"
	} else {
		i := strconv.Itoa(len(chipResults))
		res.Data.Value = i
		res.Data.Text = value
		chipResults = append(chipResults, value)
		res.Valid = true
	}
	blob, _ := json.Marshal(res)
	w.Write(blob)
}
