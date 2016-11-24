package fast

import (
	"errors"
	"html"
	"html/template"
	"io"

	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
)

type ComponentRenderer interface {
	GetBlock() string
	RenderComponent(wr io.Writer, view mgc.ViewComponentRenderer, args ...interface{}) (string, error)
}

type RenderContext struct {
	*mgc.RenderContext
	fastRenderers []ComponentRenderer
}

func NewRenderContext(t mgc.Templater, wr io.Writer) *RenderContext {
	ret := &RenderContext{
		RenderContext: mgc.NewRenderContext(t, wr),
	}
	return ret
}
func (r *RenderContext) Register(c ComponentRenderer) {
	r.fastRenderers = append(r.fastRenderers, c)
}
func (r *RenderContext) GetRenderer(block string) ComponentRenderer {
	for _, v := range r.fastRenderers {
		if v.GetBlock() == block {
			return v
		}
	}
	return nil
}
func (ctx *RenderContext) SetDefaultTo(view mgc.ViewComponentContextSetter) {
	view.SetDefaultRenderContext(ctx)
}
func (ctx *RenderContext) AttachTo(view mgc.ViewComponentContextSetter) {
	view.SetRenderContext(ctx)
}
func (ctx *RenderContext) RenderComponent(view mgc.ViewComponentRenderer, args ...interface{}) (string, error) {
	r := ctx.GetRenderer(view.GetBlock())
	if r == nil {
		return ctx.RenderContext.RenderComponent(view, args)
	}
	return r.RenderComponent(ctx.GetWriter(), view, args)
}

type ButtonRenderer struct {
	Block string
}

func NewButtonRenderer() *ButtonRenderer {
	ret := &ButtonRenderer{}
	ret.Block = "mgc/button"
	return ret
}
func (b *ButtonRenderer) GetBlock() string {
	return b.Block
}
func (b *ButtonRenderer) RenderComponent(wr io.Writer, view mgc.ViewComponentRenderer, args ...interface{}) (string, error) {
	if button, ok := view.(*components.Button); !ok {
		return "", errors.New("wrong type expected components.Button")
	} else {

		tag := "button"
		if button.Attr.Has("href") {
			if button.Attr.Has("disabled") {
				tag = "span"
			} else {
				tag = "a"
			}
		}

		io.WriteString(wr, "<"+tag+" ")

		for _, v := range button.Attr {
			io.WriteString(wr, html.EscapeString(v.Name)+"='"+html.EscapeString(v.Value)+"' ")
		}

		io.WriteString(wr, "class='"+html.EscapeString(button.Classes.Render())+"'")

		if button.GetValue() != "" {
			io.WriteString(wr, "value='"+html.EscapeString(button.GetValue())+"'")
		}
		io.WriteString(wr, ">")

		if s, ok := button.GetLabel().(string); ok {
			io.WriteString(wr, html.EscapeString(s))
		} else if s, ok := button.GetLabel().(template.HTML); ok {
			io.WriteString(wr, string(s))
		}
		io.WriteString(wr, "</"+tag+">")
	}
	return "", nil
}
