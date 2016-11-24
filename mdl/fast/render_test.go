package fast_test

import (
	"bytes"
	"os"
	"testing"

	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"github.com/mh-cbon/mdl-go-components/mdl/fast"
)

var tplBuffer bytes.Buffer
var templateRenderer *mgc.RenderContext
var buttonA *components.Button

var fastBuffer bytes.Buffer
var fastRenderer *fast.RenderContext
var buttonB *components.Button

func init() {
	templateRenderer = mgc.NewRenderContext(mgc.MustTemplate(), &tplBuffer)
	buttonA = components.NewButton()
	buttonA.SetLabel("hello")
	templateRenderer.AttachTo(buttonA)

	fastRenderer = fast.NewRenderContext(mgc.MustTemplate(), &fastBuffer)
	fastRenderer.Register(fast.NewButtonRenderer())
	buttonB = components.NewButton()
	buttonB.SetLabel("hello")
	fastRenderer.AttachTo(buttonB)
}

func BenchmarkRenderWithTemplate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buttonA.Render()
		(tplBuffer).Reset()
	}
}

func BenchmarkRenderWithFast(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buttonB.Render()
		(fastBuffer).Reset()
	}
}

func TestButtonHref(t *testing.T) {
	input := components.NewButton()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetLink("http://whatever")

	// expectations := []string{
	// 	`<a[^>]+>label</a>`,
	// 	`<a[^>]+name="name"[^>]+>label</a>`,
	// 	`<a[^>]+href="http://whatever"[^>]+>label</a>`,
	// 	`<a[^>]+class="mdl-button mdl-js-button"[^>]+>label</a>`,
	// 	`<a[^>]+id="id"[^>]+>label</a>`,
	// }
	fastRenderer := fast.NewRenderContext(mgc.MustTemplate(), os.Stdout)
	fastRenderer.Register(fast.NewButtonRenderer())
	fastRenderer.AttachTo(buttonB)
	buttonB.Render()
}
