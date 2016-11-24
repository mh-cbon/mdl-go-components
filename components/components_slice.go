package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type ComponentWithAComponentsSlice struct {
	Components *base.ComponentsSlice
}

func NewComponentWithAComponentsSlice() *ComponentWithAComponentsSlice {
	return &ComponentWithAComponentsSlice{
		Components: base.NewComponentsSlice(),
	}
}

func (view *ComponentWithAComponentsSlice) GetComponents() []mgc.ViewComponentRenderer {
	return view.Components.GetComponents()
}

func (view *ComponentWithAComponentsSlice) AddComponent(some mgc.ViewComponentRenderer) {
	view.Components.AddComponent(some)
}

func (view *ComponentWithAComponentsSlice) AddText(name string) *Input {
	component := NewInputText()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddHidden(name string) *Input {
	component := NewInputHidden()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddPassword(name string) *Input {
	component := NewInputPassword()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddFile(name string) *InputFile {
	component := NewInputFile()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddTextarea(name string) *Textarea {
	component := NewTextarea()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddWysiwyg(name string) *Textarea {
	component := NewWysiwyg()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddDate(name string) *InputDate {
	component := NewInputDate()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddSubmit(name string) *Button {
	component := NewSubmit()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddCheckboxSlice(name string) *CheckboxSlice {
	component := NewCheckboxSlice()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddRadioSlice(name string) *CheckboxSlice {
	component := NewRadioSlice()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddTextSlice(name string) *InputSlice {
	component := NewInputTextSlice()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddHiddenSlice(name string) *InputSlice {
	component := NewInputHiddenSlice()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddSelect(name string) *Select {
	component := NewSelect()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddCropper(name string) *Cropper {
	component := NewCropper()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddChipAutocomplete(name string) *ChipAutocomplete {
	component := NewChipAutocomplete()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) AddInputConfirm(name string) *InputConfirm {
	component := NewInputConfirmText()
	component.SetName(name)
	view.AddComponent(component)
	return component
}

func (view *ComponentWithAComponentsSlice) GetComponent(name string) mgc.ViewComponentRenderer {
	return view.Components.GetComponent(name)
}
func (view *ComponentWithAComponentsSlice) GetValueSetter(name string) base.ValueSetter {
	return view.Components.GetValueSetter(name)
}
func (view *ComponentWithAComponentsSlice) GetValueDateSetter(name string) base.ValueDateSetter {
	return view.Components.GetValueDateSetter(name)
}
func (view *ComponentWithAComponentsSlice) GetValueSliceSetter(name string) base.ValueSliceSetter {
	return view.Components.GetValueSliceSetter(name)
}
func (view *ComponentWithAComponentsSlice) GetOptioner(name string) base.Optionner {
	return view.Components.GetOptioner(name)
}
func (view *ComponentWithAComponentsSlice) GetButton(name string) *Button {
	c := view.GetComponent(name)
	if c != nil {
		if v, ok := c.(*Button); ok {
			return v
		}
	}
	return nil
}
func (view *ComponentWithAComponentsSlice) GetChipAutocomplete(name string) *ChipAutocomplete {
	c := view.GetComponent(name)
	if c != nil {
		if v, ok := c.(*ChipAutocomplete); ok {
			return v
		}
	}
	return nil
}
func (view *ComponentWithAComponentsSlice) GetCropper(name string) *Cropper {
	c := view.GetComponent(name)
	if c != nil {
		if v, ok := c.(*Cropper); ok {
			return v
		}
	}
	return nil
}
func (view *ComponentWithAComponentsSlice) GetInputConfirm(name string) *InputConfirm {
	c := view.GetComponent(name)
	if c != nil {
		if v, ok := c.(*InputConfirm); ok {
			return v
		}
	}
	return nil
}
func (view *ComponentWithAComponentsSlice) GetExpander() *Expand {
	for _, c := range view.GetComponents() {
		if v, ok := c.(*Expand); ok {
			return v
		}
	}
	return nil
}
func (view *ComponentWithAComponentsSlice) GetSubmit(name string) *Button {
	return view.GetButton(name)
}
func (view *ComponentWithAComponentsSlice) GetText(name string) *Input {
	c := view.GetComponent(name)
	if c != nil {
		if v, ok := c.(*Input); ok {
			return v
		}
	}
	return nil
}
func (view *ComponentWithAComponentsSlice) GetHidden(name string) *Input {
	return view.GetText(name)
}
func (view *ComponentWithAComponentsSlice) GetPassword(name string) *Input {
	return view.GetText(name)
}
func (view *ComponentWithAComponentsSlice) GetSelect(name string) *Select {
	c := view.GetComponent(name)
	if c != nil {
		if v, ok := c.(*Select); ok {
			return v
		}
	}
	return nil
}
