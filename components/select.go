package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Select struct {
	mgc.ViewComponent
	Node
	NodeLabel
	NodePlaceholder
	NodeSingleError
	NodeWithOptions
	InputClasses ClassList
	InputAttr    AttrList
	Multiple     bool
	Icon         string
}

func NewSelect() *Select {
	ret := &Select{}
	ret.SetBlock("form_select")
	return ret
}

func (i *Select) SetIcon(s string) {
	i.Icon = s
}
func (i *Select) GetIcon() string {
	return i.Icon
}

func (i *Select) SetMultiple(s bool) {
	i.Multiple = s
}
func (i *Select) IsMultiple() bool {
	return i.Multiple
}

func (view *Select) SetRipple(b bool) {
	if b {
		view.Classes.Add("mdl-js-ripple-effect")
	} else {
		view.Classes.Remove("mdl-js-ripple-effect")
	}
}

func (view *Select) SetDisabled(b bool) {
	if b {
		view.InputAttr.Set("disabled", "disabled")
	} else {
		view.InputAttr.Remove("disabled")
	}
}

func (view *Select) SetName(b string) {
	view.InputAttr.Set("name", b)
}
func (view *Select) GetName() string {
	return view.InputAttr.GetValue("name")
}

func (view *Select) SetId(b string) {
	view.InputAttr.Set("id", b)
}
func (view *Select) GetId() string {
	return view.InputAttr.GetValue("id")
}

func (view *Select) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
