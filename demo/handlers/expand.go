package handlers

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"net/http"
)

func Expand(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Expand",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var t1 *components.Expand
	var line *components.Slice

	// -
	child := components.NewInputConfirmText()
	child.SetError("This is an error!")

	child.InputLeft.SetName("left-name")
	child.InputLeft.SetId("left-id")
	child.InputLeft.SetLabel("Left label")
	child.InputLeft.SetValue("Left value")

	child.InputRight.SetName("right-name")
	child.InputRight.SetId("right-id")
	child.InputRight.SetLabel("Right label")
	child.InputRight.SetValue("Right value")

	// -
	line = components.NewSlice()
	t1 = components.NewExpand()

	bt := components.NewCheckbox()
	bt.GetOption().SetLabel("labelopt")
	bt.GetOption().SetValue("valueopt")
	bt.SetName("name")
	t1.SetBt(bt)
	t1.AddChild(child)
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewExpand()

	bt2 := components.NewCheckbox()
	bt2.GetOption().SetLabel("labelopt")
	bt2.GetOption().SetValue("valueopt")
	bt2.SetName("name")
	bt2.SetChecked(true)
	t1.SetBt(bt2)
	t1.AddChild(child)
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewExpand()

	bt3 := components.NewText()
	bt3.SetContent("Click me !")
	t1.SetBt(bt3)
	t1.AddChild(child)
	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
