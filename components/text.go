package components

import (
	"html/template"

	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Text struct {
	mgc.ViewComponent
	Content         interface{}
	TranslationArgs interface{}
}

func NewText() *Text {
	ret := &Text{}
	ret.SetBlock("mgc/text_node")
	return ret
}

func (view *Text) SetTranslationArgs(some interface{}) {
	view.TranslationArgs = some
}

func (view *Text) SetHTMLContent(some string) {
	view.Content = template.HTML(some)
}

func (view *Text) SetContent(some string) {
	view.Content = some
}

func (view *Text) GetContent() interface{} {
	return view.Content
}

func (view *Text) Translate(t base.Translator) {
	if x, ok := view.Content.(template.HTML); ok {
		view.SetHTMLContent(t.T(string(x), view.TranslationArgs))
	} else if x, ok := view.Content.(string); ok {
		view.SetContent(t.T(x, view.TranslationArgs))
	}
}

func (view *Text) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
