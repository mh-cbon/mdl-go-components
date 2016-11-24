package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Textarea struct {
	mgc.ViewComponent
	base.Node
	base.NodeLabel
	base.NodePlaceholder
	base.NodeSingleValue
	base.NodeSingleError
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

func (view *Textarea) SetErrors(p base.ErrorProvider) {
	err := p.GetError(view.GetName())
	if err != nil {
		view.SetError(err)
	}
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
