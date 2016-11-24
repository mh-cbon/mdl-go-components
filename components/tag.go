package components

import (
	"html/template"

	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Tag struct {
	mgc.ViewComponent
	base.Node
	*ComponentWithAComponentsSlice
	TagName string
}

func NewTag() *Tag {
	ret := &Tag{}
	ret.ComponentWithAComponentsSlice = NewComponentWithAComponentsSlice()
	ret.SetBlock("mgc/tag")
	return ret
}

func NewDiv() *Tag {
	ret := NewTag()
	ret.SetTagName("div")
	return ret
}

func (view *Tag) SetTagName(s string) {
	view.TagName = s
}

func (view *Tag) GetTagName() string {
	return view.TagName
}

func (view *Tag) OpenTag() template.HTML {
	return template.HTML("<" + view.TagName)
}
func (view *Tag) OpenTagClose() template.HTML {
	return template.HTML(">")
}
func (view *Tag) CloseTag() template.HTML {
	return template.HTML("</" + view.TagName + ">")
}

func (view *Tag) AddComponent(s mgc.ViewComponentRenderer) {
	view.Components.AddComponent(s)
}

// func (view *Tag) Translate(t base.Translator) {
// 	view.Components.Translate(t)
// }

func (view *Tag) Render(args ...interface{}) (string, error) {
	view.GetRenderContext().SetDefaultTo(view.Components)
	return view.GetRenderContext().RenderComponent(view, args)
}
