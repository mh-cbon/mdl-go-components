package handlers

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"net/http"
)

func InputFile(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Input file",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var t1 *components.InputFile
	var line *components.Slice

	// -
	line = components.NewSlice()
	t1 = components.NewInputFile()
	t1.SetName("the_file")
	t1.SetLabel("Choose a file")
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewInputFile()
	t1.SetName("the_file")
	t1.SetLabel("Choose a file")
	t1.SetMultiple(true)
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewInputFile()
	t1.SetName("the_file")
	t1.SetLabel("Choose a file")
	t1.SetError("it gone south :x")
	t1.SetClearIcon("delete_forever")
	t1.SetAttachIcon("account_box")
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
