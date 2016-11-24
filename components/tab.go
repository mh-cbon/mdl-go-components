package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Tabs struct {
	mgc.ViewComponent
	base.Node
	Tabs []*Tab

	MenuClasses base.ClassList
	MenuAttr    base.AttrList
}

func NewTabs() *Tabs {
	ret := &Tabs{}
	ret.SetBlock("mgc/tabs")
	return ret
}

func (l *Tabs) AddTab(title string) *Tab {
	tab := NewTab()
	tab.SetTitle(title)
	tab.SetActive(len(l.Tabs) == 0)
	l.Tabs = append(l.Tabs, tab)
	return tab
}
func (l *Tabs) At(index int) *Tab {
	return l.Tabs[index]
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
		view.GetRenderContext().SetDefaultTo(tab.Components)
	}
	return view.GetRenderContext().RenderComponent(view, args)
}
func (view *Tabs) Translate(t base.Translator) {
	for _, tab := range view.Tabs {
		tab.Translate(t)
	}
}
func (view *Tabs) SetErrors(p base.ErrorProvider) {
	for _, tab := range view.Tabs {
		tab.SetErrors(p)
	}
}

type Tab struct {
	base.Node
	*ComponentWithAComponentsSlice

	Title string

	TitleClasses base.ClassList
	TitleAttr    base.AttrList
}

func NewTab() *Tab {
	ret := &Tab{}
	ret.ComponentWithAComponentsSlice = NewComponentWithAComponentsSlice()
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

func (view *Tab) SetErrors(p base.ErrorProvider) {
	view.Components.SetErrors(p)
}
func (view *Tab) Translate(t base.Translator) {
	view.Title = t.T(view.Title)
	view.Components.Translate(t)
}
