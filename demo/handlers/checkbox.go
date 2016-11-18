package handlers

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"net/http"
)

func Checkbox(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Checbkoxes",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var t1 *components.Checkbox
	var line *components.Slice

	// -
	line = components.NewSlice()

	t1 = components.NewCheckbox()
	t1.SetLabel("checkbox regular")
	t1.SetValue("some")
	t1.SetChecked(false)
	t1.SetName("checkbox1")
	line.Add(t1)

	t1 = components.NewCheckbox()
	t1.SetLabel("checkbox disabled")
	t1.SetValue("some")
	t1.SetChecked(true)
	t1.SetDisabled(true)
	t1.SetName("checkbox1")
	line.Add(t1)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewCheckbox()
	t1.SetLabel("checkbox with ripple effect")
	t1.SetValue("some")
	t1.SetChecked(false)
	t1.SetRipple(true)
	t1.SetName("checkbox1")
	line.Add(t1)

	t1 = components.NewCheckbox()
	t1.SetLabel("checkbox checked disabled with ripple effect")
	t1.SetValue("some")
	t1.SetChecked(true)
	t1.SetRipple(true)
	t1.SetDisabled(true)
	t1.SetName("checkbox1")
	line.Add(t1)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewRadio()
	t1.SetLabel("radio regular")
	t1.SetValue("some")
	t1.SetChecked(false)
	t1.SetName("radio1")
	line.Add(t1)

	t1 = components.NewRadio()
	t1.SetLabel("radio checked with ripple effect")
	t1.SetValue("some1")
	t1.SetChecked(true)
	t1.SetRipple(true)
	t1.SetName("radio1")
	line.Add(t1)

	t1 = components.NewRadio()
	t1.SetLabel("radio disabled")
	t1.SetValue("some2")
	t1.SetChecked(false)
	t1.SetName("radio1")
	t1.SetDisabled(true)
	line.Add(t1)

	t1 = components.NewRadio()
	t1.SetLabel("radio disabled with ripple effect")
	t1.SetValue("some2")
	t1.SetChecked(false)
	t1.SetName("radio1")
	t1.SetRipple(true)
	t1.SetDisabled(true)
	line.Add(t1)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	var t2 *components.CheckIcon
	t2 = components.NewCheckIcon()
	t2.SetIcon("format_bold")
	t2.SetValue("some")
	t2.SetChecked(false)
	t2.SetName("cb2")
	line.Add(t2)

	t2 = components.NewCheckIcon()
	t2.SetIcon("format_bold")
	t2.SetValue("some1")
	t2.SetChecked(false)
	t2.SetName("cb2")
	t2.SetRipple(true)
	line.Add(t2)

	t2 = components.NewCheckIcon()
	t2.SetIcon("format_bold")
	t2.SetValue("some2")
	t2.SetChecked(true)
	t2.SetName("cb2")
	t2.SetDisabled(true)
	t2.SetRipple(true)
	line.Add(t2)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewSwitch()
	t1.SetValue("some")
	t1.SetName("switch1")
	line.Add(t1)

	t1 = components.NewSwitch()
	t1.SetValue("some")
	t1.SetName("switch2")
	t1.SetRipple(true)
	line.Add(t1)

	t1 = components.NewSwitch()
	t1.SetValue("some")
	t1.SetName("switch3")
	t1.SetDisabled(true)
	line.Add(t1)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewSwitch()
	t1.SetValue("some")
	t1.SetName("switch11")
	t1.SetChecked(true)
	line.Add(t1)

	t1 = components.NewSwitch()
	t1.SetValue("some")
	t1.SetName("switch21")
	t1.SetRipple(true)
	t1.SetChecked(true)
	line.Add(t1)

	t1 = components.NewSwitch()
	t1.SetValue("some")
	t1.SetName("switch31")
	t1.SetDisabled(true)
	t1.SetChecked(true)
	line.Add(t1)

	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewSwitch()
	t1.SetValue("some")
	t1.SetName("switch12")
	t1.SetLabel("label")
	line.Add(t1)

	t1 = components.NewSwitch()
	t1.SetValue("some")
	t1.SetName("switch22")
	t1.SetLabel("label")
	t1.SetRipple(true)
	t1.SetChecked(true)
	line.Add(t1)

	t1 = components.NewSwitch()
	t1.SetValue("some")
	t1.SetName("switch32")
	t1.SetDisabled(true)
	t1.SetLabel("label")
	line.Add(t1)

	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
