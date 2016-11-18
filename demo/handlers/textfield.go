package handlers

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"net/http"
)

func Textfield(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Textfields",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var line *components.Slice
	var t1 *components.Input

	// -
	line = components.NewSlice()

	t1 = components.NewInputText()
	t1.SetLabel("textfield empty")
	t1.SetName("textfield1")
	t1.SetPlaceHolderOnly(true)
	line.Add(t1)

	t1 = components.NewInputText()
	t1.SetLabel("textfield with value")
	t1.SetName("textfield1")
	t1.SetValue("some\">")
	t1.SetPlaceHolderOnly(true)
	line.Add(t1)

	t1 = components.NewInputText()
	t1.SetLabel("textfield with value and error")
	t1.SetName("textfield1")
	t1.SetValue("some")
	t1.SetError("an error")
	t1.SetPlaceHolderOnly(true)
	line.Add(t1)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewInputText()
	t1.SetLabel("textfield empty")
	t1.SetName("textfield1")
	line.Add(t1)

	t1 = components.NewInputText()
	t1.SetLabel("textfield with value")
	t1.SetName("textfield1")
	t1.SetValue("some")
	line.Add(t1)

	t1 = components.NewInputText()
	t1.SetLabel("textfield with value and error")
	t1.SetName("textfield1")
	t1.SetValue("some")
	t1.SetError("an error")
	line.Add(t1)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewInputPassword()
	t1.SetLabel("password empty")
	t1.SetName("textfield1")
	line.Add(t1)

	t1 = components.NewInputPassword()
	t1.SetLabel("password with value")
	t1.SetName("textfield1")
	t1.SetValue("some")
	line.Add(t1)

	t1 = components.NewInputPassword()
	t1.SetLabel("password with value and error")
	t1.SetName("textfield1")
	t1.SetValue("some")
	t1.SetError("an error")
	line.Add(t1)

	data.Components = append(data.Components, line)

	renderComponents(w, data)
}
