package mdlgocomponents_test

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"html/template"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// demonstrate use of mdl-go-components
func Example_main() {

	// create  new template to reference mdl components blocks
	t := template.New("")
	if _, err := t.ParseGlob(mgc.TemplatesGlob()); err != nil {
		panic(err)
	}

	// create a new render context: (template + writer)
	ctx := mgc.NewRenderContext(t, out)

	// create your components, and configure them
	btn := components.NewButton()
	btn.SetLabel("label")
	btn.Classes.Add("the-class-name")

	// attach components to a render context
	ctx.AttachTo(btn)

	// Call the component Render method to get the content.
	if _, err := btn.Render(); err != nil {
		panic(err)
	}

	// when suitable, components can be nested with the type
	// mgc.ViewComponentRenderer see tabs, dup, form ect.
}
