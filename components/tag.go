package components

import (
  "html/template"
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Tag struct {
	mgc.ViewComponent
  Node
	Components []mgc.ViewComponentRenderer
	TagName string
}

func NewTag() *Tag {
	ret := &Tag{}
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
  return template.HTML("<"+view.TagName)
}
func (view *Tag) OpenTagClose() template.HTML {
  return template.HTML(">")
}
func (view *Tag) CloseTag() template.HTML {
  return template.HTML("</"+view.TagName+">")
}

func (view *Tag) Add(s mgc.ViewComponentRenderer) {
	view.Components = append(view.Components, s)
}

func (view *Tag) Translate(t Translator) {
	for _, c := range view.Components {
		if v, ok := c.(NodeTranslator); ok {
			v.Translate(t)
		}
	}
}

func (view *Tag) Render(args ...interface{}) (string, error) {
	for _, c := range view.Components {
		view.GetRenderContext().SetDefaultTo(c)
	}
	return view.GetRenderContext().RenderComponent(view, args)
}
