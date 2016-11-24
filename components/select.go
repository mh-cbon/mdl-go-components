package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Select struct {
	mgc.ViewComponent
	base.Node
	base.NodeLabel
	base.NodePlaceholder
	base.NodeSingleError
	base.NodeWithOptions
	InputClasses base.ClassList
	InputAttr    base.AttrList
	Multiple     bool
	Icon         string
}

func NewSelect() *Select {
	ret := &Select{}
	ret.SetBlock("mgc/form_select")
	return ret
}

func (view *Select) SetErrors(p base.ErrorProvider) {
	err := p.GetError(view.GetName())
	if err != nil {
		view.SetError(err)
	}
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

func (view *Select) UpdateWindowUrl(b bool) {
	if b {
		view.Classes.Add("custom-js-select-change-url")
	} else {
		view.Classes.Remove("custom-js-select-change-url")
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

func (view *Select) Translate(t base.Translator) {
	view.NodeLabel.Translate(t)
	view.NodeWithOptions.Translate(t)
}

func (view *Select) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
