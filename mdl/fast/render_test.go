package fast_test

import (
	"bytes"
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
