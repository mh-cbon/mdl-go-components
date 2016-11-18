package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Textarea struct {
	mgc.ViewComponent
	Node
	NodeLabel
	NodePlaceholder
	NodeSingleValue
	NodeSingleError
	Rows int
}

func NewTextarea() *Textarea {
	ret := &Textarea{}
	ret.SetBlock("mgc/form_textarea")
	ret.SetRows(3)
	return ret
}

func NewWysiwyg() *Textarea {
	ret := NewTextarea()
	ret.SetBlock("mgc/form_wysiwyg")
	return ret
}

func (i *Textarea) SetRows(b int) {
	i.Rows = b
}
func (i *Textarea) GetRows() int {
	return i.Rows
}

func (view *Textarea) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
