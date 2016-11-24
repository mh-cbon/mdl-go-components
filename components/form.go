package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Form struct {
	mgc.ViewComponent
	base.Node
	*ComponentWithAComponentsSlice
	Error string
}

func NewForm() *Form {
	ret := &Form{}
	ret.ComponentWithAComponentsSlice = NewComponentWithAComponentsSlice()
	ret.SetBlock("mgc/form")
	return ret
}

func (view *Form) Render(args ...interface{}) (string, error) {
	view.GetRenderContext().SetDefaultTo(view.Components)
	return view.GetRenderContext().RenderComponent(view, args)
}
func (view *Form) Translate(t base.Translator) {
	view.Components.Translate(t)
}
func (view *Form) SetErrors(p base.ErrorProvider) {
	view.Components.SetErrors(p)
}

func (view *Form) SetMethod(b string) *Form {
	view.Attr.Set("method", b)
	return view
}
func (view *Form) GetMethod() string {
	return view.Attr.GetValue("method")
}
func (view *Form) SetAction(b string) *Form {
	view.Attr.Set("action", b)
	return view
}
func (view *Form) GetAction() string {
	return view.Attr.GetValue("action")
}
func (view *Form) SetEncType(b string) *Form {
	view.Attr.Set("enctype", b)
	return view
}
func (view *Form) GetEncType() string {
	return view.Attr.GetValue("enctype")
}
func (view *Form) SetTarget(b string) *Form {
	view.Attr.Set("target", b)
	return view
}
func (view *Form) GetTarget() string {
	return view.Attr.GetValue("target")
}
