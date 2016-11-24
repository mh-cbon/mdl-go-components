package handlers

import (
	"net/http"

	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
)

func Textarea(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Textarea",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var t1 *components.Textarea
	var line *components.Slice

	// -
	line = components.NewSlice()

	t1 = components.NewTextarea()
	t1.SetName("textarea")
	t1.SetLabel("textarea field")
	t1.SetValue("text here")
	line.AddComponent(t1)

	t1 = components.NewTextarea()
	t1.SetName("textarea")
	t1.SetLabel("textarea field")
	t1.SetValue("text here")
	t1.SetPlaceHolderOnly(true)
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewWysiwyg()
	t1.SetName("textarea")
	t1.SetLabel("textarea field")
	t1.SetValue("text here")
	line.AddComponent(t1)

	t1 = components.NewWysiwyg()
	t1.SetName("textarea")
	t1.SetLabel("textarea field")
	t1.SetValue("text here")
	t1.SetPlaceHolderOnly(true)
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
