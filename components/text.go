package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"html/template"
)

type Text struct {
	mgc.ViewComponent
	Content interface{}
}

func NewText() *Text {
	ret := &Text{}
	ret.SetBlock("mgc/text_node")
	return ret
}

func (t *Text) SetHTMLContent(some string) {
	t.Content = template.HTML(some)
}

func (t *Text) SetContent(some string) {
	t.Content = some
}

func (t *Text) GetContent() interface{} {
	return t.Content
}

func (view *Text) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
