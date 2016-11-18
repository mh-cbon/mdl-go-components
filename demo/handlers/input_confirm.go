package handlers

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"net/http"
)

func InputConfirm(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Input confirm",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var t1 *components.InputConfirm
	var line *components.Slice

	// -
	line = components.NewSlice()
	t1 = components.NewInputConfirmText()
	t1.InputLeft.SetName("left")
	t1.InputLeft.SetLabel("Type the text")
	t1.InputRight.SetLabel("Type again")
	t1.InputRight.SetName("right")
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewInputConfirmPassword()
	t1.InputLeft.SetLabel("Type the pasword")
	t1.InputLeft.SetName("left")
	t1.InputRight.SetLabel("Type again")
	t1.InputRight.SetName("right")
	t1.SetPlaceHolderOnly(true)
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewInputConfirmPassword()
	t1.InputLeft.SetLabel("Type the pasword")
	t1.InputLeft.SetName("left")
	t1.InputLeft.SetValue("aaa")
	t1.InputRight.SetLabel("Type again")
	t1.InputRight.SetName("right")
	t1.SetError("Oh no")
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
