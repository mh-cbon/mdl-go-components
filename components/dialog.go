package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Dialog struct {
	mgc.ViewComponent
	Node

	Title            string
	BgClasses        ClassList
	ContainerClasses ClassList

	Content mgc.ViewComponentRenderer
	Ok      *Button
	Cancel  *Button
	Close   *Button
}

func NewDialog() *Dialog {
	ret := &Dialog{}
	ret.SetOverlayColor("grey")
	ret.SetBgContainerColor("white")
	ret.Ok = NewButton()
	ret.SetOkText("Ok", "")
	ret.Ok.Classes.Add("custom-dialog-confirm")
	ret.Cancel = NewButton()
	ret.SetCancelText("Cancel", "")
	ret.Cancel.Classes.Add("custom-dialog-cancel")
	ret.Close = NewButton()
	ret.SetCloseText("", "close")
	ret.Close.Classes.Add("custom-dialog-close")
	ret.SetBlock("mgc/component_dialog")
	return ret
}

func (view *Dialog) SetOverlayColor(colorname string) {
	view.BgClasses.RemoveStartingWith("mdl-color--")
	view.BgClasses.Add("mdl-color--" + colorname)
}
func (view *Dialog) SetBgContainerColor(colorname string) {
	view.ContainerClasses.RemoveStartingWith("mdl-color--")
	view.ContainerClasses.Add("mdl-color--" + colorname)
}

func (view *Dialog) Translate(t Translator) {
  view.Ok.Translate(t)
  view.Cancel.Translate(t)
  view.Close.Translate(t)
  if view.Content!=nil {
    if v, ok := view.Content.(ViewTranslator); ok {
      v.Translate(t)
    }
  }
}

func (view *Dialog) SetTitle(t string) {
	view.Title = t
}
func (view *Dialog) GetTitle() string {
	return view.Title
}

func (view *Dialog) SetContent(t string) {
	text := NewText()
	text.SetContent(t)
	view.Content = text
}

func (view *Dialog) SetOkText(t string, icon string) {
	if icon == "" {
		view.Ok.SetLabel(t)
	} else {
		view.Ok.SetSafeLabel(t + "<i class='material-icons'>" + icon + "</i>")
	}
}
func (view *Dialog) SetOkAttr(t string, v string) {
	view.Ok.Attr.Set(t, v)
}

func (view *Dialog) SetCancelText(t string, icon string) {
	if icon == "" {
		view.Cancel.SetLabel(t)
	} else {
		view.Cancel.SetSafeLabel(t + "<i class='material-icons'>" + icon + "</i>")
	}
}
func (view *Dialog) SetCancelAttr(t string, v string) {
	view.Cancel.Attr.Set(t, v)
}

func (view *Dialog) SetCloseText(t string, icon string) {
	if icon == "" {
		view.Close.SetLabel(t)
	} else {
		view.Close.SetSafeLabel(t + "<i class='material-icons'>" + icon + "</i>")
	}
}
func (view *Dialog) SetCloseAttr(t string, v string) {
	view.Close.Attr.Set(t, v)
}

func (view *Dialog) Render(args ...interface{}) (string, error) {
	view.GetRenderContext().SetDefaultTo(view.Ok)
	view.GetRenderContext().SetDefaultTo(view.Cancel)
	view.GetRenderContext().SetDefaultTo(view.Close)
	if view.Content != nil {
		view.GetRenderContext().SetDefaultTo(view.Content)
	}
	return view.GetRenderContext().RenderComponent(view, args)
}
