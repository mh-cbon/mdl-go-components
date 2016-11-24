package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Dup struct {
	mgc.ViewComponent
	base.Node
	*ComponentWithAComponentsSlice

	Duped mgc.ViewComponentRenderer

	BtRemoveAttr    base.AttrList
	BtRemoveClasses base.ClassList
	BtRemoveText    string

	BtAddAttr    base.AttrList
	BtAddClasses base.ClassList
	BtAddText    string
}

func NewDup() *Dup {
	ret := &Dup{}
	ret.ComponentWithAComponentsSlice = NewComponentWithAComponentsSlice()
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

func (view *Dup) AddComponent(some mgc.ViewComponentRenderer) {
	view.Components.AddComponent(some)
}
func (view *Dup) Translate(t base.Translator) {
	view.Components.Translate(t)
	if view.Duped != nil {
		if v, ok := view.Duped.(base.NodeTranslator); ok {
			v.Translate(t)
		}
	}
}
func (view *Dup) SetErrors(p base.ErrorProvider) {
	view.Components.SetErrors(p)
}

func (view *Dup) Render(args ...interface{}) (string, error) {
	view.GetRenderContext().SetDefaultTo(view.Components)
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
func (ctx *DupRenderContext) SetDefaultTo(view mgc.ViewComponentContextSetter) {
	if v, ok := view.(mgc.Ider); ok {
		if v.GetId() == "" {
			v.SetId(ctx.b + "-" + ctx.ContextRenderer.GetId() + "-$incrIndex$")
		}
	}
	view.SetDefaultRenderContext(ctx)
}
func (ctx *DupRenderContext) AttachTo(view mgc.ViewComponentContextSetter) {
	if v, ok := view.(mgc.Ider); ok {
		if v.GetId() == "" {
			v.SetId(ctx.b + "-" + ctx.ContextRenderer.GetId() + "-$incrIndex$")
		}
	}
	view.SetRenderContext(ctx)
}
