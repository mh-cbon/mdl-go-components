package handlers

import (
	"net/http"

	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
)

func Button(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Buttons",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var t1 *components.Button
	var line *components.Slice

	// -
	line = components.NewSlice()

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetRipple(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetRaised(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetRipple(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetDisabled(true)
	t1.SetRaised(true)
	t1.SetRipple(true)
	line.AddComponent(t1)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetColored(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetColored(true)
	t1.SetRipple(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetColored(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetColored(true)
	t1.SetRipple(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetColored(true)
	t1.SetRipple(true)
	t1.SetDisabled(true)
	line.AddComponent(t1)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetAccentColored(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetAccentColored(true)
	t1.SetRipple(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetAccentColored(true)
	t1.SetRaised(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetAccentColored(true)
	t1.SetRipple(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetLabel("button")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetAccentColored(true)
	t1.SetRipple(true)
	t1.SetDisabled(true)
	line.AddComponent(t1)

	data.Components = append(data.Components, line)

	//-
	line = components.NewSlice()

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRipple(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRaised(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetRipple(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetRipple(true)
	t1.SetDisabled(true)
	line.AddComponent(t1)

	data.Components = append(data.Components, line)

	//-
	line = components.NewSlice()

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetColored(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRipple(true)
	t1.SetColored(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetColored(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetRipple(true)
	t1.SetColored(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetRipple(true)
	t1.SetDisabled(true)
	t1.SetColored(true)
	line.AddComponent(t1)

	data.Components = append(data.Components, line)

	//-
	line = components.NewSlice()

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetAccentColored(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRipple(true)
	t1.SetAccentColored(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetAccentColored(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetRipple(true)
	t1.SetAccentColored(true)
	line.AddComponent(t1)

	t1 = components.NewButton()
	t1.SetIcon("add")
	t1.SetName("button1")
	t1.SetRaised(true)
	t1.SetRipple(true)
	t1.SetDisabled(true)
	t1.SetAccentColored(true)
	line.AddComponent(t1)

	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
