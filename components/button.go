package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Button struct {
	mgc.ViewComponent
	Node
	NodeType
	NodeLabel
	NodeSingleValue

	Link string
}

func NewButton() *Button {
	ret := &Button{}
	ret.SetBlock("button")
	ret.SetType("button")
	ret.Classes.Add("mdl-button mdl-js-button")
	return ret
}

func NewSubmit() *Button {
	ret := NewButton()
	ret.SetType("submit")
	return ret
}

func (view *Button) SetLink(b string) {
	if b == "" {
		view.Attr.Remove("href")
	} else {
		view.Attr.Set("href", b)
	}
}

func (view *Button) SetIcon(b string) {
	view.Classes.Add("mdl-button--fab")
	view.SetSafeLabel("<i class='material-icons'>" + b + "</i>")
}

func (view *Button) SetRipple(b bool) {
	if b {
		view.Classes.Add("mdl-js-ripple-effect")
	} else {
		view.Classes.Remove("mdl-js-ripple-effect")
	}
}

func (view *Button) SetDisabled(b bool) {
	if b {
		view.Attr.Set("disabled", "disabled")
	} else {
		view.Attr.Remove("disabled")
	}
}

func (view *Button) SetRaised(b bool) {
	if b {
		view.Classes.Add("mdl-button--raised")
	} else {
		view.Classes.Remove("mdl-button--raised")
	}
}

func (view *Button) SetColored(b bool) {
	if b {
		view.Classes.Add("mdl-button--colored")
	} else {
		view.Classes.Remove("mdl-button--colored")
	}
}

func (view *Button) SetAccentColored(b bool) {
	if b {
		view.Classes.Add("mdl-button--accent")
	} else {
		view.Classes.Remove("mdl-button--accent")
	}
}

func (view *Button) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
