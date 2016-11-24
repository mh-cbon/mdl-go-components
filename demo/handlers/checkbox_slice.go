package handlers

import (
	"net/http"

	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
)

func Checkbox_slice(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Checbkox slices",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var t1 *components.CheckboxSlice
	var line *components.Slice

	// -
	line = components.NewSlice()
	t1 = components.NewCheckboxSlice()
	t1.SetName("cb_slice")
	t1.AddOption("value", "label")
	t1.AddOption("value2", "label2")
	t1.SetValues([]string{"value"})
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewCheckboxSlice()
	t1.SetName("cb_slice1")
	t1.SetDisabled(true)
	t1.AddOption("value", "label")
	t1.AddOption("value2", "label2")
	t1.SetValues([]string{"value"})
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewCheckboxSlice()
	t1.SetName("cb_slice2")
	t1.SetRipple(true)
	t1.AddOption("value", "label")
	t1.AddOption("value2", "label2")
	t1.SetValues([]string{"value"})
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewRadioSlice()
	t1.SetName("radio_slice")
	t1.AddOption("value", "label")
	t1.AddOption("value2", "label2")
	t1.SetValues([]string{"value"})
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewRadioSlice()
	t1.SetName("radio_slice1")
	t1.SetDisabled(true)
	t1.AddOption("value", "label")
	t1.AddOption("value2", "label2")
	t1.SetValues([]string{"value"})
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewRadioSlice()
	t1.SetName("radio_slice2")
	t1.SetRipple(true)
	t1.AddOption("value", "label")
	t1.AddOption("value2", "label2")
	t1.SetValues([]string{"value"})
	line.AddComponent(t1)
	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
