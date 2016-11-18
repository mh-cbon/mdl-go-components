package handlers

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"net/http"
)

func Slider(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Slider",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var t1 *components.Slider
	var line *components.Slice

	// -
	line = components.NewSlice()
	t1 = components.NewSlider()
	t1.SetName("the_file")
	t1.SetMin("0")
	t1.SetMax("100")
	t1.SetStep("5")
	t1.SetValue("25")
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewSlider()
	t1.SetName("the_file")
	t1.SetLabel("Choose a value")
	t1.SetMin("0")
	t1.SetMax("100")
	t1.SetStep("5")
	t1.SetValue("25")
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewSlider()
	t1.SetName("the_file")
	t1.SetPlaceHolderOnly(true)
	t1.SetError("it gone south :x")
	t1.SetMin("0")
	t1.SetMax("100")
	t1.SetStep("5")
	t1.SetValue("25")
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewSlider()
	t1.SetName("the_file")
	t1.SetLabel("Choose a value")
	t1.SetError("it gone south :x")
	t1.SetMin("0")
	t1.SetMax("100")
	t1.SetStep("5")
	t1.SetValue("25")
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
