package components

import (
	"strconv"

	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type CheckboxSlice struct {
	mgc.ViewComponent
	base.Node
	base.NodeType
	base.NodeMutlipleValues
	base.NodeWithOptions
	base.NodeSingleError
	InputClasses base.ClassList
	InputAttr    base.AttrList
	ItemBlock    string
}

func NewCheckboxSlice() *CheckboxSlice {
	ret := &CheckboxSlice{}
	ret.SetBlock("mgc/form_checkbox_slice")
	ret.SetItemBlock("mgc/form_checkbox")
	ret.SetType("checkbox")
	return ret
}

func NewRadioSlice() *CheckboxSlice {
	ret := NewCheckboxSlice()
	ret.SetType("radio")
	ret.SetItemBlock("mgc/form_radio")
	return ret
}

func (view *CheckboxSlice) SetErrors(p base.ErrorProvider) {
	err := p.GetError(view.GetName())
	if err != nil {
		view.SetError(err)
	}
}

func (i *CheckboxSlice) SetItemBlock(s string) {
	i.ItemBlock = s
}
func (i *CheckboxSlice) GetItemBlock() string {
	return i.ItemBlock
}

func (view *CheckboxSlice) SetRipple(b bool) {
	if b {
		view.Classes.Add("mdl-js-ripple-effect")
	} else {
		view.Classes.Remove("mdl-js-ripple-effect")
	}
}

func (view *CheckboxSlice) SetDisabled(b bool) {
	if b {
		view.InputAttr.Set("disabled", "disabled")
	} else {
		view.InputAttr.Remove("disabled")
	}
}

func (view *CheckboxSlice) SetValue(s string) {
	view.NodeWithOptions.SetValue(s)
}

func (view *CheckboxSlice) GetItem(index int, option base.NodeOption) *Checkbox {
	sview := NewCheckbox()
	sview.Attr.MergeFrom(view.Attr)
	sview.Classes.MergeFrom(view.Classes)
	sview.InputAttr.MergeFrom(view.InputAttr)
	sview.InputClasses.MergeFrom(view.InputClasses)
	sview.SetId(view.GetId() + "-" + strconv.Itoa(index))
	sview.SetName(view.GetName())
	sview.SetType(view.GetType())
	sview.SetBlock(view.GetItemBlock())
	sview.SetOption(option)
	view.GetRenderContext().SetDefaultTo(sview)
	return sview
}

func (view *CheckboxSlice) RenderItem(index int, option base.NodeOption, args ...interface{}) (string, error) {
	return view.GetItem(index, option).Render(args)
}

func (view *CheckboxSlice) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
