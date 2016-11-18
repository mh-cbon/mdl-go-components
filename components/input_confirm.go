package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type InputConfirm struct {
	mgc.ViewComponent
	NodeType
	NodePlaceholder
	NodeSingleError

	Attr    AttrList
	Classes ClassList

	InputLeft  *PartialInputConfirm
	InputRight *PartialInputConfirm
}

func NewInputConfirm() *InputConfirm {
	ret := &InputConfirm{}
	ret.SetBlock("form_input_confirm")
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
	view.InputLeft.SetId("left-" + view.GetRenderContext().GetId())
	view.InputRight.SetId("right-" + view.GetRenderContext().GetId())
	return view.GetRenderContext().RenderComponent(view, args)
}

func (view *InputConfirm) SetName(b string) {
	view.InputLeft.SetName(b)
	view.InputRight.SetName("confirm-" + b)
}

type PartialInputConfirm struct {
	Node
	NodeLabel
	NodeSingleValue

	InputAttr    AttrList
	InputClasses ClassList
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
