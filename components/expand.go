package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Checker interface {
	IsChecked() bool
	SetChecked(b bool)
}

type Expand struct {
	mgc.ViewComponent
	base.Node
	*ComponentWithAComponentsSlice
	Bt       mgc.ViewComponentRenderer
	Expanded bool
}

func NewExpand() *Expand {
	ret := &Expand{}
	ret.SetBlock("mgc/expand")
	ret.ComponentWithAComponentsSlice = NewComponentWithAComponentsSlice()
	return ret
}

func (view *Expand) Translate(t base.Translator) {
	if v, ok := view.Bt.(base.NodeTranslator); ok {
		v.Translate(t)
	}
	view.Components.Translate(t)
}
func (view *Expand) SetErrors(p base.ErrorProvider) {
	if v, ok := view.Bt.(base.NodeErrorsSetter); ok {
		v.SetErrors(p)
	}
	view.Components.SetErrors(p)
}

func (view *Expand) SetBt(b mgc.ViewComponentRenderer) {
	view.Bt = b
	if v, ok := b.(Checker); ok {
		view.SetExpanded(v.IsChecked())
	}
}

func (view *Expand) SetExpanded(b bool) {
	view.Expanded = b
	if view.Bt != nil {
		if v, ok := view.Bt.(Checker); ok {
			v.SetChecked(b)
		}
	}
}
func (view *Expand) IsExpanded() bool {
	return view.Expanded
}

func (view *Expand) Render(args ...interface{}) (string, error) {
	view.GetRenderContext().SetDefaultTo(view.Bt)
	view.GetRenderContext().SetDefaultTo(view.Components)
	return view.GetRenderContext().RenderComponent(view, args)
}
