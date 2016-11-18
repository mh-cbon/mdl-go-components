package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Form struct {
	mgc.ViewComponent
	Node
	Error      string
	Components []mgc.ViewComponentRenderer
}

func NewForm() *Form {
	ret := &Form{}
	ret.SetBlock("form")
	return ret
}

func (view *Form) Render(args ...interface{}) (string, error) {
	for _, c := range view.Components {
		view.GetRenderContext().SetDefaultTo(c)
	}
	return view.GetRenderContext().RenderComponent(view, args)
}

func (view *Form) SetMethod(b string) {
	view.Attr.Set("method", b)
}
func (view *Form) GetMethod() string {
	return view.Attr.GetValue("method")
}
func (view *Form) SetAction(b string) {
	view.Attr.Set("action", b)
}
func (view *Form) GetAction() string {
	return view.Attr.GetValue("action")
}
func (view *Form) SetEncType(b string) {
	view.Attr.Set("enctype", b)
}
func (view *Form) GetEncType() string {
	return view.Attr.GetValue("enctype")
}
func (view *Form) SetTarget(b string) {
	view.Attr.Set("target", b)
}
func (view *Form) GetTarget() string {
	return view.Attr.GetValue("target")
}

func (view *Form) Add(some mgc.ViewComponentRenderer) {
	view.Components = append(view.Components, some)
	// if v, ok := some.(mgc.Ider); ok && v.GetId()=="" {
	//   v.SetId("rnd-"+name+"-"+view.GetRenderContext().GetId())
	// }
}

func (view *Form) AddText(name string) *Input {
	component := NewInputText()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddHidden(name string) *Input {
	component := NewInputHidden()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddPassword(name string) *Input {
	component := NewInputPassword()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddFile(name string) *InputFile {
	component := NewInputFile()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddTextarea(name string) *Textarea {
	component := NewTextarea()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddWysiwyg(name string) *Textarea {
	component := NewWysiwyg()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddDate(name string) *InputDate {
	component := NewInputDate()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddSubmit(name string) *Button {
	component := NewSubmit()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddCheckboxSlice(name string) *CheckboxSlice {
	component := NewCheckboxSlice()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddRadioSlice(name string) *CheckboxSlice {
	component := NewRadioSlice()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddTextSlice(name string) *InputSlice {
	component := NewInputTextSlice()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddHiddenSlice(name string) *InputSlice {
	component := NewInputHiddenSlice()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddSelect(name string) *Select {
	component := NewSelect()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddCropper(name string) *Cropper {
	component := NewCropper()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddChipAutocomplete(name string) *ChipAutocomplete {
	component := NewChipAutocomplete()
	component.SetName(name)
	view.Add(component)
	return component
}

func (view *Form) AddInputConfirm(name string) *InputConfirm {
	component := NewInputConfirmText()
	component.SetName(name)
	view.Add(component)
	return component
}
