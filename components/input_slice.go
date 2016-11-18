package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"strconv"
)

type InputSlice struct {
	mgc.ViewComponent
	Node
	NodeType
	NodeLabel
	NodePlaceholder
	NodeMutlipleValues
	NodeSingleError
	ItemBlock string
}

func (i *InputSlice) SetItemBlock(s string) {
	i.ItemBlock = s
}
func (i *InputSlice) GetItemBlock() string {
	return i.ItemBlock
}

func NewInputSlice() *InputSlice {
	ret := &InputSlice{}
	ret.SetBlock("form_input_slice")
	ret.SetItemBlock("form_input")
	return ret
}

func NewInputTextSlice() *InputSlice {
	ret := NewInputSlice()
	ret.SetBlock("form_input_slice")
	ret.SetItemBlock("form_input")
	ret.SetType("text")
	return ret
}

func NewInputPasswordSlice() *InputSlice {
	ret := NewInputSlice()
	ret.SetBlock("form_input_slice")
	ret.SetItemBlock("form_input")
	ret.SetType("password")
	return ret
}

func NewInputHiddenSlice() *InputSlice {
	ret := NewInputSlice()
	ret.SetBlock("form_input_slice")
	ret.SetItemBlock("form_hidden")
	ret.SetType("hidden")
	return ret
}

func (view *InputSlice) GetItem(index int, value string) *Input {
	sview := NewInput()
	sview.SetId(view.GetId() + "-" + strconv.Itoa(index))
	sview.SetType(view.GetType())
	sview.SetLabel(view.GetLabel())
	sview.SetError(view.GetError())
	sview.SetName(view.GetName())
	sview.SetPlaceHolderOnly(view.IsPlaceHolderOnly())
	sview.SetBlock(view.GetItemBlock())
	sview.SetValue(value)
	view.GetRenderContext().SetDefaultTo(sview)
	return sview
}

func (view *InputSlice) RenderItem(index int, value string, args ...interface{}) (string, error) {
	return view.GetItem(index, value).Render(args)
}

func (view *InputSlice) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
