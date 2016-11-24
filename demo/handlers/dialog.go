package handlers

import (
	"net/http"

	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
)

func Dialog(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Dialog",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var t1 *components.Dialog
	var text *components.Text
	var bt *components.Button
	var link *components.Text
	var line *components.Slice

	text = components.NewText()
	text.SetContent("hello content!")

	// -
	line = components.NewSlice()

	t1 = components.NewDialog()
	t1.SetTitle("The title")
	t1.SetId("id")

	t1.Ok.SetType("submit")
	t1.Ok.Attr.Set("form", "form1")

	t1.Content = text
	line.AddComponent(t1)

	bt = components.NewButton()
	bt.SetLabel("button")
	bt.SetName("button1")
	bt.Classes.Add("custom-js-confirm-button")
	bt.Attr.Set("confirm", "#"+t1.GetId())
	line.AddComponent(bt)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewDialog()
	t1.SetTitle("The title")
	t1.SetId("id3")

	t1.Ok.SetLink("https://google.com/")

	t1.Content = text
	line.AddComponent(t1)

	bt = components.NewButton()
	bt.SetLabel("button")
	bt.SetName("button1")
	bt.Classes.Add("custom-js-confirm-button")
	bt.Attr.Set("confirm", "#"+t1.GetId())
	line.AddComponent(bt)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewDialog()
	t1.SetTitle("The title")
	t1.SetId("id2")

	t1.Ok.SetType("submit")
	t1.Ok.Attr.Set("form", "form1")

	t1.Content = text
	line.AddComponent(t1)

	link = components.NewText()
	link.SetHTMLContent(`
    <a href="#" class="custom-js-confirm-button" confirm="#` + t1.GetId() + `">click it</a>`)
	line.AddComponent(link)

	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
