package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Checker interface {
	IsChecked() bool
}

type Expand struct {
	mgc.ViewComponent
	Node
	Bt       mgc.ViewComponentRenderer
	Children []mgc.ViewComponentRenderer
	Expanded bool
}

func NewExpand() *Expand {
	ret := &Expand{}
	ret.SetBlock("mgc/component_expand")
	return ret
}

func (view *Expand) AddChild(some mgc.ViewComponentRenderer) {
	view.Children = append(view.Children, some)
}

func (view *Expand) SetBt(b mgc.ViewComponentRenderer) {
	view.Bt = b
	if v, ok := b.(Checker); ok {
		view.SetExpanded(v.IsChecked())
	}
}

func (view *Expand) SetExpanded(b bool) {
	view.Expanded = b
}
func (view *Expand) IsExpanded() bool {
	return view.Expanded
}

func (view *Expand) Render(args ...interface{}) (string, error) {
	view.GetRenderContext().SetDefaultTo(view.Bt)
	for _, c := range view.Children {
		view.GetRenderContext().SetDefaultTo(c)
	}
	return view.GetRenderContext().RenderComponent(view, args)
}
