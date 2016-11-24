package handlers

import (
	"net/http"

	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
)

func Dup(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Dup",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	dup := components.NewDup()
	dup.SetId("id")
	dup.SetBtAddText("more")
	dup.SetBtRemoveText("less")

	bt := components.NewButton()
	bt.SetLabel("hello bt!")
	dup.AddComponent(bt)
	bt = components.NewButton()
	bt.SetLabel("hello bt2!")
	dup.AddComponent(bt)
	bt = components.NewButton()
	bt.SetLabel("hello bt3!")
	dup.AddComponent(bt)

	sl := components.NewSlice()
	txt := components.NewInputText()
	txt.SetLabel("hello duped bt!")
	sl.AddComponent(txt)
	txt = components.NewInputText()
	txt.SetLabel("hello duped bt!2")
	sl.AddComponent(txt)
	dup.SetDup(sl)

	data.Components = append(data.Components, dup)

	// -
	renderComponents(w, data)
}
