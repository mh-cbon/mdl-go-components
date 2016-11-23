package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Checker interface {
	IsChecked() bool
	SetChecked(b bool)
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

func (view *Expand) Translate(t Translator) {
  if v, ok := view.Bt.(NodeTranslator); ok {
    v.Translate(t)
  }
	for _, c := range view.Children {
		if v, ok := c.(NodeTranslator); ok {
			v.Translate(t)
		}
	}
}
func (view *Expand) SetErrors(p ErrorProvider) {
  if v, ok := view.Bt.(NodeErrorsSetter); ok {
    v.SetErrors(p)
  }
	for _, c := range view.Children {
		if v, ok := c.(NodeErrorsSetter); ok {
			v.SetErrors(p)
		}
	}
}

func (view *Expand) Add(some mgc.ViewComponentRenderer) {
	view.Children = append(view.Children, some)
}
func (view *Expand) AddChild(some mgc.ViewComponentRenderer) {
	view.Add(some)
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
	for _, c := range view.Children {
		view.GetRenderContext().SetDefaultTo(c)
	}
	return view.GetRenderContext().RenderComponent(view, args)
}
