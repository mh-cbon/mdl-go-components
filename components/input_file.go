package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type InputFile struct {
	mgc.ViewComponent
	NodeLabel
	NodePlaceholder
	NodeSingleError

	Attr          AttrList
	Classes       ClassList
	InputAttr     AttrList
	InputClasses  ClassList
	InputFileAttr AttrList

	InputTextValue string
	AttachIcon string
	ClearIcon  string
}

func NewInputFile() *InputFile {
	view := &InputFile{}
	view.SetBlock("mgc/form_input_file")
	view.SetAttachIcon("attach_file")
	view.SetClearIcon("clear")
	view.InputAttr.Set("type", "text")
	view.InputAttr.Set("readonly", "readonly")
	return view
}

func (view *InputFile) SetErrors(p ErrorProvider) {
	err := p.GetError(view.GetName())
	if err!=nil {
		view.SetError(err)
	}
}

func (view *InputFile) SetClearIcon(b string) {
	view.ClearIcon = b
}
func (view *InputFile) GetClearIcon() string {
	return view.ClearIcon
}

func (view *InputFile) SetAttachIcon(b string) {
	view.AttachIcon = b
}
func (view *InputFile) GetAttachIcon() string {
	return view.AttachIcon
}

func (view *InputFile) SetName(b string) {
	view.InputFileAttr.Set("name", b)
}
func (view *InputFile) GetName() string {
	return view.InputFileAttr.GetValue("name")
}

func (view *InputFile) SetId(b string) {
	view.InputAttr.Set("id", b)
}
func (view *InputFile) GetId() string {
	return view.InputAttr.GetValue("id")
}
func (view *InputFile) SetInputTextValue(b string) {
	view.InputTextValue = b
}
func (view *InputFile) SetInputTextName(b string) {
  view.InputAttr.Set("name", b)
}

func (view *InputFile) SetMultiple(b bool) {
	if b {
		view.InputFileAttr.Set("multiple", "multiple")
	} else {
		view.InputFileAttr.Remove("multiple")
	}
}
func (view *InputFile) IsMultiple() bool {
	return view.InputFileAttr.GetValue("multiple") == "multiple"
}

func (view *InputFile) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
