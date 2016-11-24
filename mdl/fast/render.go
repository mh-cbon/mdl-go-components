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
func (ctx *RenderContext) SetDefaultTo(view mgc.ViewComponentRenderer) {
	view.SetDefaultRenderContext(ctx)
}
func (ctx *RenderContext) AttachTo(view mgc.ViewComponentRenderer) {
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
		//EscapeString
		if button.Attr.Has("href") {
			if button.Attr.Has("disabled") {
				wr.Write([]byte("<span "))
			} else {
				wr.Write([]byte("<a "))
			}
		} else {
			wr.Write([]byte("<button "))
		}

		for _, v := range button.Attr {
			wr.Write([]byte(html.EscapeString(v.Name) + "='" + html.EscapeString(v.Value) + "'"))
		}
		wr.Write([]byte("class='" + html.EscapeString(button.Classes.Render()) + "'"))
		if button.Attr.Has("value") {
			wr.Write([]byte("value='" + html.EscapeString(button.GetValue()) + "'"))
		}
		wr.Write([]byte(">"))

		if s, ok := button.GetLabel().(string); ok {
			wr.Write([]byte(html.EscapeString(s)))
		} else if s, ok := button.GetLabel().(template.HTML); ok {
			wr.Write([]byte(string(s)))
		}

		if button.Attr.Has("href") {
			if button.Attr.Has("disabled") {
				wr.Write([]byte("</span>"))
			} else {
				wr.Write([]byte("</a>"))
			}
		} else {
			wr.Write([]byte("</button>"))
		}
	}
	return "", nil
}
