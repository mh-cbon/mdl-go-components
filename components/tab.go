package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Tabs struct {
	mgc.ViewComponent
	Node
	Tabs []*Tab

	MenuClasses ClassList
	MenuAttr    AttrList
}

func NewTabs() *Tabs {
	ret := &Tabs{}
	ret.SetBlock("mgc/tabs")
	return ret
}

func (l *Tabs) Add(title string, content mgc.ViewComponentRenderer) *Tab {
	tab := NewTab()
	tab.SetTitle(title)
	tab.SetContent(content)
	tab.SetActive(len(l.Tabs) == 0)
	l.Tabs = append(l.Tabs, tab)
	return tab
}
func (l *Tabs) AddSlice(title string) *Slice {
  slice := NewSlice()
	l.Add(title, slice)
  return slice
}
func (l *Tabs) SetActive(index int) {
	for i, tab := range l.Tabs {
		tab.SetActive(index == i)
	}
}
func (view *Tabs) Render(args ...interface{}) (string, error) {
	for _, tab := range view.Tabs {
		if tab.GetId() == "" {
			tab.SetId("rnd-" + view.GetRenderContext().GetId())
		}
    view.GetRenderContext().SetDefaultTo(tab.Content)
	}
	return view.GetRenderContext().RenderComponent(view, args)
}
func (view *Tabs) Translate(t Translator) {
	for _, tab := range view.Tabs {
    tab.Translate(t)
	}
}
func (view *Tabs) SetErrors(p ErrorProvider) {
	for _, c := range view.Tabs {
		if v, ok := c.Content.(NodeErrorsSetter); ok {
			v.SetErrors(p)
		}
	}
}

type Tab struct {
	Node
	Title   string
	Content mgc.ViewComponentRenderer

	TitleClasses ClassList
	TitleAttr    AttrList
}

func NewTab() *Tab {
	ret := &Tab{}
	return ret
}

func (t *Tab) IsActive() bool {
	return t.Classes.Contains("is-active")
}
func (t *Tab) SetActive(b bool) {
	if b {
		t.TitleClasses.Add("is-active")
		t.Classes.Add("is-active")
	} else {
		t.TitleClasses.Remove("is-active")
		t.Classes.Remove("is-active")
	}
}

func (t *Tab) GetTitle() string {
	return t.Title
}
func (t *Tab) SetTitle(b string) {
	t.Title = b
}

func (t *Tab) GetContent() mgc.ViewComponentRenderer {
	return t.Content
}
func (t *Tab) SetContent(b mgc.ViewComponentRenderer) {
	t.Content = b
}
func (view *Tab) Translate(t Translator) {
  view.Title = t.T(view.Title)
	if v, ok := view.Content.(NodeTranslator); ok {
		v.Translate(t)
	}
}
