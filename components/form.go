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
	ret.SetBlock("mgc/form")
	return ret
}

func (view *Form) Render(args ...interface{}) (string, error) {
	for _, c := range view.Components {
		view.GetRenderContext().SetDefaultTo(c)
	}
	return view.GetRenderContext().RenderComponent(view, args)
}
func (view *Form) Translate(t Translator) {
	for _, c := range view.Components {
		if v, ok := c.(NodeTranslator); ok {
			v.Translate(t)
		}
	}
}
func (view *Form) SetErrors(p ErrorProvider) {
	for _, c := range view.Components {
		if v, ok := c.(NodeErrorsSetter); ok {
			v.SetErrors(p)
		}
	}
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
func (view *Form) Get(name string) mgc.ViewComponentRenderer {
	for _, c := range view.Components {
		if v, ok := c.(Namer); ok {
			if v.GetName() == name {
				r := view.GetRenderContext()
				if r != nil {
					r.SetDefaultTo(c)
				}
				return c
			}
		}
	}
	return nil
}
func (view *Form) GetText(name string) *Input {
	c := view.Get(name)
	if c != nil {
		if v, ok := c.(*Input); ok {
			return v
		}
	}
	return nil
}
func (view *Form) GetHidden(name string) *Input {
	return view.GetText(name)
}
func (view *Form) GetPassword(name string) *Input {
	return view.GetText(name)
}
func (view *Form) GetSelect(name string) *Select {
	c := view.Get(name)
	if c != nil {
		if v, ok := c.(*Select); ok {
			return v
		}
	}
	return nil
}
func (view *Form) GetValueSetter(name string) ValueSetter {
	c := view.Get(name)
	if c != nil {
		if v, ok := c.(ValueSetter); ok {
			return v
		}
	}
	return nil
}
func (view *Form) GetValueDateSetter(name string) ValueDateSetter {
	c := view.Get(name)
	if c != nil {
		if v, ok := c.(ValueDateSetter); ok {
			return v
		}
	}
	return nil
}
func (view *Form) GetValueSliceSetter(name string) ValueSliceSetter {
	c := view.Get(name)
	if c != nil {
		if v, ok := c.(ValueSliceSetter); ok {
			return v
		}
	}
	return nil
}
func (view *Form) GetButton(name string) *Button {
	c := view.Get(name)
	if c != nil {
		if v, ok := c.(*Button); ok {
			return v
		}
	}
	return nil
}
func (view *Form) GetChipAutocomplete(name string) *ChipAutocomplete {
	c := view.Get(name)
	if c != nil {
		if v, ok := c.(*ChipAutocomplete); ok {
			return v
		}
	}
	return nil
}
func (view *Form) GetSubmit(name string) *Button {
	return view.GetButton(name)
}

func (view *Form) Add(some mgc.ViewComponentRenderer) {
	view.Components = append(view.Components, some)
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
