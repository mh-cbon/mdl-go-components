package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Checkbox struct {
	mgc.ViewComponent
	Node
	NodeType
	NodeWithOption
	InputClasses ClassList
	InputAttr    AttrList
}

func NewCheckbox() *Checkbox {
	ret := &Checkbox{}
	ret.SetBlock("mgc/form_checkbox")
	ret.SetType("checkbox")
	return ret
}

func NewRadio() *Checkbox {
	ret := NewCheckbox()
	ret.SetType("radio")
	ret.SetBlock("mgc/form_radio")
	return ret
}

func NewSwitch() *Checkbox {
	ret := NewCheckbox()
	ret.SetType("checkbox")
	ret.SetBlock("mgc/form_switch")
	return ret
}

func (view *Checkbox) SetRipple(b bool) {
	if b {
		view.Classes.Add("mdl-js-ripple-effect")
	} else {
		view.Classes.Remove("mdl-js-ripple-effect")
	}
}

func (view *Checkbox) SetDisabled(b bool) {
	if b {
		view.InputAttr.Set("disabled", "disabled")
	} else {
		view.InputAttr.Remove("disabled")
	}
}

func (view *Checkbox) SetName(b string) {
	view.InputAttr.Set("name", b)
}
func (view *Checkbox) GetName() string {
	return view.InputAttr.GetValue("name")
}

func (view *Checkbox) SetId(b string) {
	view.InputAttr.Set("id", b)
}
func (view *Checkbox) GetId() string {
	return view.InputAttr.GetValue("id")
}

func (view *Checkbox) SetLabel(b interface{}) {
	view.Option.SetLabel(b)
}
func (view *Checkbox) GetLabel() interface{} {
	return view.Option.GetLabel()
}

func (view *Checkbox) SetValue(b string) {
	view.Option.SetValue(b)
}
func (view *Checkbox) GetValue() string {
	return view.Option.GetValue()
}

func (view *Checkbox) SetChecked(b bool) {
	view.Option.SetChecked(b)
}
func (view *Checkbox) IsChecked() bool {
	return view.Option.IsChecked()
}

func (view *Checkbox) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
