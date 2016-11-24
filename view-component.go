package mdlgocomponents

import (
	"io"
	"strconv"
)

type Templater interface {
	ExecuteTemplate(wr io.Writer, name string, args interface{}) error
}

type Ider interface {
	SetId(s string)
	GetId() string
}

type RenderContext struct {
	t  Templater
	wr io.Writer
	i  uint64
}

func NewRenderContext(t Templater, wr io.Writer) *RenderContext {
	return &RenderContext{
		t:  t,
		wr: wr,
	}
}
func (ctx *RenderContext) GetId() string {
	ret := strconv.FormatUint(ctx.i, 10)
	ctx.i++
	return ret
}
func (ctx *RenderContext) SetTemplate(s Templater) {
	ctx.t = s
}
func (ctx *RenderContext) GetTemplate() Templater {
	return ctx.t
}
func (ctx *RenderContext) SetWriter(s io.Writer) {
	ctx.wr = s
}
func (ctx *RenderContext) GetWriter() io.Writer {
	return ctx.wr
}
func (ctx *RenderContext) RenderComponent(view ViewComponentRenderer, args ...interface{}) (string, error) {
	block := view.GetBlock()
	if block == "" {
		panic("Did you forget to set the block name with SetBlock() ?")
	}
	if v, ok := view.(Ider); ok {
		if v.GetId() == "" {
			v.SetId("rnd-" + ctx.GetId())
		}
	}
	viewdata := NewViewData(args)
	viewdata.Component = view
	tpl := ctx.GetTemplate()
	wr := ctx.GetWriter()
	return "", tpl.ExecuteTemplate(wr, block, viewdata)
}
func (ctx *RenderContext) SetDefaultTo(view ViewComponentContextSetter) {
	view.SetDefaultRenderContext(ctx)
}
func (ctx *RenderContext) AttachTo(view ViewComponentContextSetter) {
	view.SetRenderContext(ctx)
}

type ContextRenderer interface {
	GetId() string
	RenderComponent(view ViewComponentRenderer, args ...interface{}) (string, error)
	SetDefaultTo(view ViewComponentContextSetter)
	AttachTo(view ViewComponentContextSetter)
}

type ViewComponentContextSetter interface {
	SetDefaultRenderContext(r ContextRenderer)
	SetRenderContext(r ContextRenderer)
}

type ViewComponentRenderer interface {
	SetDefaultRenderContext(r ContextRenderer)
	SetRenderContext(r ContextRenderer)
	GetRenderContext() ContextRenderer

	SetBlock(s string)
	GetBlock() string

	Render(args ...interface{}) (string, error)
}

type ViewData struct {
	Component interface{}
	Data      interface{}
}

func NewViewData(datas ...interface{}) *ViewData {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	return &ViewData{
		Data: data,
	}
}

type ViewComponent struct {
	ctx   ContextRenderer
	block string
}

func (view *ViewComponent) SetBlock(s string) {
	view.block = s
}
func (view *ViewComponent) GetBlock() string {
	return view.block
}

func (view *ViewComponent) SetRenderContext(s ContextRenderer) {
	view.ctx = s
}
func (view *ViewComponent) GetRenderContext() ContextRenderer {
	return view.ctx
}

func (view *ViewComponent) SetDefaultRenderContext(ctx ContextRenderer) {
	if view.GetRenderContext() == nil {
		view.SetRenderContext(ctx)
	}
}
