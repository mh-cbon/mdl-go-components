package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Dup struct {
	mgc.ViewComponent
	Node

	Duped mgc.ViewComponentRenderer
	Items []mgc.ViewComponentRenderer

	BtRemoveAttr    AttrList
	BtRemoveClasses ClassList
	BtRemoveText    string

	BtAddAttr    AttrList
	BtAddClasses ClassList
	BtAddText    string
}

func NewDup() *Dup {
	ret := &Dup{}
	ret.SetBlock("mgc/dup")
	return ret
}

func (view *Dup) SetBtRemoveAttr(p string, v string) {
	view.BtRemoveAttr.Set(p, v)
}
func (view *Dup) AddBtRemoveClass(p string) {
	view.BtRemoveClasses.Add(p)
}
func (view *Dup) SetBtRemoveText(p string) {
	view.BtRemoveText = p
}

func (view *Dup) SetBtAddAttr(p string, v string) {
	view.BtAddAttr.Set(p, v)
}
func (view *Dup) AddBtAddClass(p string) {
	view.BtAddClasses.Add(p)
}
func (view *Dup) SetBtAddText(p string) {
	view.BtAddText = p
}

func (view *Dup) SetDup(duped mgc.ViewComponentRenderer) {
	view.Duped = duped
}

func (view *Dup) Add(some mgc.ViewComponentRenderer) {
	view.Items = append(view.Items, some)
}

func (view *Dup) Render(args ...interface{}) (string, error) {
	for _, v := range view.Items {
		view.GetRenderContext().SetDefaultTo(v)
	}
	if view.Duped != nil {
		ctx := &DupRenderContext{
			b:               view.GetId(),
			ContextRenderer: view.GetRenderContext(),
		}
		ctx.SetDefaultTo(view.Duped)
	}
	return view.GetRenderContext().RenderComponent(view, args)
}

type DupRenderContext struct {
	mgc.ContextRenderer
	b string
}

func (ctx *DupRenderContext) RenderComponent(view mgc.ViewComponentRenderer, args ...interface{}) (string, error) {
	if v, ok := view.(mgc.Ider); ok {
		if v.GetId() == "" {
			v.SetId(ctx.b + "-" + ctx.ContextRenderer.GetId())
		}
	}
	return ctx.ContextRenderer.RenderComponent(view, args)
}
func (ctx *DupRenderContext) SetDefaultTo(view mgc.ViewComponentRenderer) {
	if v, ok := view.(mgc.Ider); ok {
		if v.GetId() == "" {
			v.SetId(ctx.b + "-" + ctx.ContextRenderer.GetId() + "-$incrIndex$")
		}
	}
	view.SetDefaultRenderContext(ctx)
}
func (ctx *DupRenderContext) AttachTo(view mgc.ViewComponentRenderer) {
	if v, ok := view.(mgc.Ider); ok {
		if v.GetId() == "" {
			v.SetId(ctx.b + "-" + ctx.ContextRenderer.GetId() + "-$incrIndex$")
		}
	}
	view.SetRenderContext(ctx)
}
