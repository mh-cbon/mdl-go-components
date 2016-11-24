package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type InputConfirm struct {
	mgc.ViewComponent
	base.NodeType
	base.NodePlaceholder
	base.NodeSingleError

	Attr    base.AttrList
	Classes base.ClassList

	InputLeft  *PartialInputConfirm
	InputRight *PartialInputConfirm
}

func NewInputConfirm() *InputConfirm {
	ret := &InputConfirm{}
	ret.SetBlock("mgc/form_input_confirm")
	ret.InputLeft = &PartialInputConfirm{}
	ret.InputRight = &PartialInputConfirm{}
	return ret
}

func NewInputConfirmText() *InputConfirm {
	ret := NewInputConfirm()
	ret.SetType("text")
	return ret
}

func NewInputConfirmPassword() *InputConfirm {
	ret := NewInputConfirm()
	ret.SetType("password")
	return ret
}

func (view *InputConfirm) Render(args ...interface{}) (string, error) {
	if view.InputLeft.GetId() == "" {
		view.InputLeft.SetId("left-" + view.GetRenderContext().GetId())
	}
	if view.InputRight.GetId() == "" {
		view.InputRight.SetId("right-" + view.GetRenderContext().GetId())
	}
	return view.GetRenderContext().RenderComponent(view, args)
}

func (view *InputConfirm) Translate(t base.Translator) {
	view.InputLeft.Translate(t)
	view.InputRight.Translate(t)
}
func (view *InputConfirm) SetErrors(p base.ErrorProvider) {
	err := p.GetError(view.InputLeft.GetName())
	if err != nil {
		view.SetError(err)
	}
}

func (view *InputConfirm) SetName(b string) {
	view.InputLeft.SetName(b)
	view.InputRight.SetName("Confirm" + b)
}

func (view *InputConfirm) GetName() string {
	return view.InputLeft.GetName()
}

type PartialInputConfirm struct {
	base.Node
	base.NodeLabel
	base.NodeSingleValue

	InputAttr    base.AttrList
	InputClasses base.ClassList
}

func (view *PartialInputConfirm) SetName(b string) {
	view.InputAttr.Set("name", b)
}
func (view *PartialInputConfirm) GetName() string {
	return view.InputAttr.GetValue("name")
}

func (view *PartialInputConfirm) SetId(b string) {
	view.InputAttr.Set("id", b)
}
func (view *PartialInputConfirm) GetId() string {
	return view.InputAttr.GetValue("id")
}
