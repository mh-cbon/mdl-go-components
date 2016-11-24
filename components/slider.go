package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Slider struct {
	mgc.ViewComponent
	base.NodeLabel
	base.NodePlaceholder
	base.NodeSingleValue
	base.NodeSingleError

	Attr         base.AttrList
	Classes      base.ClassList
	InputAttr    base.AttrList
	InputClasses base.ClassList
}

func NewSlider() *Slider {
	ret := &Slider{}
	ret.SetBlock("mgc/form_slider")
	ret.SetType("range")
	return ret
}

func (view *Slider) SetErrors(p base.ErrorProvider) {
	err := p.GetError(view.GetName())
	if err != nil {
		view.SetError(err)
	}
}

func (view *Slider) SetMin(b string) {
	view.InputAttr.Set("min", b)
}
func (view *Slider) GetMin() string {
	return view.InputAttr.GetValue("min")
}

func (view *Slider) SetMax(b string) {
	view.InputAttr.Set("max", b)
}
func (view *Slider) GetMax() string {
	return view.InputAttr.GetValue("max")
}

func (view *Slider) SetStep(b string) {
	view.InputAttr.Set("step", b)
}
func (view *Slider) GetStep() string {
	return view.InputAttr.GetValue("step")
}

func (view *Slider) SetType(b string) {
	view.InputAttr.Set("type", b)
}
func (view *Slider) GetType() string {
	return view.InputAttr.GetValue("type")
}

func (view *Slider) SetName(b string) {
	view.InputAttr.Set("name", b)
}
func (view *Slider) GetName() string {
	return view.InputAttr.GetValue("name")
}

func (view *Slider) SetId(b string) {
	view.InputAttr.Set("id", b)
}
func (view *Slider) GetId() string {
	return view.InputAttr.GetValue("id")
}

func (view *Slider) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
