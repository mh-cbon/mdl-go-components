package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type CheckIcon struct {
	mgc.ViewComponent
	base.Node
	base.NodeWithOption
	InputClasses base.ClassList
	InputAttr    base.AttrList
	Icon         string
}

func NewCheckIcon() *CheckIcon {
	ret := &CheckIcon{}
	ret.SetBlock("mgc/form_checkicon")
	return ret
}

func (i *CheckIcon) SetIcon(s string) {
	i.Icon = s
}
func (i *CheckIcon) GetIcon() string {
	return i.Icon
}

func (view *CheckIcon) SetName(b string) {
	view.InputAttr.Set("name", b)
}
func (view *CheckIcon) GetName() string {
	return view.InputAttr.GetValue("name")
}

func (view *CheckIcon) SetId(b string) {
	view.InputAttr.Set("id", b)
}
func (view *CheckIcon) GetId() string {
	return view.InputAttr.GetValue("id")
}

func (view *CheckIcon) SetRipple(b bool) {
	if b {
		view.Classes.Add("mdl-js-ripple-effect")
	} else {
		view.Classes.Remove("mdl-js-ripple-effect")
	}
}
func (view *CheckIcon) SetDisabled(b bool) {
	if b {
		view.InputAttr.Set("disabled", "disabled")
	} else {
		view.InputAttr.Remove("disabled")
	}
}

func (view *CheckIcon) SetLabel(b interface{}) {
	view.Option.SetLabel(b)
}
func (view *CheckIcon) GetLabel() interface{} {
	return view.Option.GetLabel()
}

func (view *CheckIcon) SetValue(b string) {
	view.Option.SetValue(b)
}
func (view *CheckIcon) GetValue() string {
	return view.Option.GetValue()
}

func (view *CheckIcon) SetChecked(b bool) {
	view.Option.SetChecked(b)
}
func (view *CheckIcon) IsChecked() bool {
	return view.Option.IsChecked()
}

func (view *CheckIcon) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
