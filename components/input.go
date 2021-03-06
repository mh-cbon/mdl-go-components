package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Input struct {
	mgc.ViewComponent
	base.NodeType
	base.NodeLabel
	base.NodePlaceholder
	base.NodeSingleValue
	base.NodeSingleError

	Attr         base.AttrList
	Classes      base.ClassList
	InputAttr    base.AttrList
	InputClasses base.ClassList
}

func NewInput() *Input {
	ret := &Input{}
	ret.SetBlock("mgc/form_input")
	return ret
}

func NewInputText() *Input {
	ret := NewInput()
	ret.SetType("text")
	return ret
}

func NewInputHidden() *Input {
	ret := NewInput()
	ret.SetBlock("mgc/form_hidden")
	ret.SetType("hidden")
	return ret
}

func NewInputPassword() *Input {
	ret := NewInput()
	ret.SetType("password")
	return ret
}

func (view *Input) SetErrors(p base.ErrorProvider) {
	err := p.GetError(view.GetName())
	if err != nil {
		view.SetError(err)
	}
}

func (view *Input) SetName(b string) {
	view.InputAttr.Set("name", b)
}
func (view *Input) GetName() string {
	return view.InputAttr.GetValue("name")
}

func (view *Input) SetId(b string) {
	view.InputAttr.Set("id", b)
}
func (view *Input) GetId() string {
	return view.InputAttr.GetValue("id")
}

func (view *Input) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
