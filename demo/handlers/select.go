package handlers

import (
	"net/http"

	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
)

func SelectField(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Select",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var t1 *components.Select
	var line *components.Slice

	// -
	line = components.NewSlice()
	t1 = components.NewSelect()
	t1.SetName("select")
	t1.SetPlaceHolderOnly(true)
	t1.SetLabel("select field")
	t1.AddOption("", "")
	t1.AddOption("value", "label")
	t1.AddOption("value2", "label2")
	t1.CheckOptions([]string{"value"})
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewSelect()
	t1.SetName("select")
	t1.SetPlaceHolderOnly(true)
	t1.SetError("this is an error")
	t1.SetLabel("select field")
	t1.AddOption("", "")
	t1.AddOption("value", "label")
	t1.AddOption("value2", "label2")
	t1.CheckOptions([]string{"value"})
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewSelect()
	t1.SetName("select")
	t1.SetLabel("select field")
	t1.AddOption("", "")
	t1.AddOption("value", "label")
	t1.AddOption("value2", "label2")
	t1.CheckOptions([]string{"value"})
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewSelect()
	t1.SetName("select")
	t1.SetLabel("select field")
	t1.SetError("this is an error")
	t1.AddOption("", "")
	t1.AddOption("value", "label")
	t1.AddOption("value2", "label2")
	t1.CheckOptions([]string{"value"})
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewSelect()
	t1.SetIcon("arrow_drop_down")
	t1.SetLabel("select field")
	t1.SetError("this is an error")
	t1.AddOption("", "")
	t1.AddOption("value", "label")
	t1.AddOption("value2", "label2")
	t1.CheckOptions([]string{"value"})
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
