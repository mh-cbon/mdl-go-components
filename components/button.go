package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Button struct {
	mgc.ViewComponent
	base.Node
	base.NodeType
	base.NodeLabel
	base.NodeSingleValue

	Link string
}

func NewButton() *Button {
	ret := &Button{}
	ret.SetBlock("mgc/button")
	ret.SetType("button")
	ret.Classes.Add("mdl-button mdl-js-button")
	return ret
}

func NewSubmit() *Button {
	ret := NewButton()
	ret.SetType("submit")
	return ret
}

func (view *Button) SetLink(b string) *Button {
	if b == "" {
		view.Attr.Remove("href")
	} else {
		view.Attr.Set("href", b)
	}
	return view
}

func (view *Button) SetTarget(b string) *Button {
	if b == "" {
		view.Attr.Remove("target")
	} else {
		view.Attr.Set("target", b)
	}
	return view
}

func (view *Button) SetFormId(b string) *Button {
	if b == "" {
		view.Attr.Remove("form")
		view.SetType("button")
	} else {
		view.Attr.Set("form", b)
		view.SetType("submit")
	}
	return view
}

func (view *Button) SetConfirmDialogId(b string) *Button {
	if b == "" {
		view.Attr.Remove("confirm")
		view.Classes.Remove("custom-js-confirm-button")
	} else {
		view.Attr.Set("confirm", "#"+b)
		view.Classes.Add("custom-js-confirm-button")
	}
	return view
}

func (view *Button) SetIcon(b string) *Button {
	if b == "" {
		view.SetLabel("")
		view.Classes.Remove("mdl-button--icon")
	} else {
		view.SetSafeLabel("<i class='material-icons'>" + b + "</i>")
		view.Classes.Add("mdl-button--icon")
	}
	return view
}

func (view *Button) SetRounded(b string) *Button {
	if b == "" {
		view.Classes.Remove("mdl-button--fab")
	} else {
		view.Classes.Add("mdl-button--fab")
	}
	return view
}

func (view *Button) SetMiniRounded(b string) *Button {
	if b == "" {
		view.Classes.Remove("mdl-button--mini-fab")
	} else {
		view.Classes.Add("mdl-button--mini-fab")
	}
	return view
}

func (view *Button) SetRipple(b bool) *Button {
	if b {
		view.Classes.Add("mdl-js-ripple-effect")
	} else {
		view.Classes.Remove("mdl-js-ripple-effect")
	}
	return view
}

func (view *Button) SetDisabled(b bool) *Button {
	if b {
		view.Attr.Set("disabled", "disabled")
	} else {
		view.Attr.Remove("disabled")
	}
	return view
}

func (view *Button) IsDisabled() bool {
	return view.Attr.GetValue("disabled") == "disabled"
}

func (view *Button) SetRaised(b bool) *Button {
	if b {
		view.Classes.Add("mdl-button--raised")
	} else {
		view.Classes.Remove("mdl-button--raised")
	}
	return view
}

func (view *Button) SetColored(b bool) *Button {
	if b {
		view.Classes.Add("mdl-button--colored")
	} else {
		view.Classes.Remove("mdl-button--colored")
	}
	return view
}

func (view *Button) SetAccentColored(b bool) *Button {
	if b {
		view.Classes.Add("mdl-button--accent")
	} else {
		view.Classes.Remove("mdl-button--accent")
	}
	return view
}

func (view *Button) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
