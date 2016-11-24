# mdl-go-components

[![GoDoc](https://godoc.org/github.com/mh-cbon/mdl-go-components?status.svg)](https://godoc.org/github.com/mh-cbon/mdl-go-components)

Material Design Lite components intergration for Go.

This integration is particularly suitable with [my fork of mdl](https://github.com/mh-cbon/material-design-lite).

## Install

```sh
glide get github.com/mh-cbon/mdl-go-components
```

## Build

To build a specific resources location,

```sh
go build -ldflags="-X github.com/mh-cbon/mdl-go-components.Tplpath=YOURPATH" your-main.go
```

## Usage

```go
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
	t := mgc.MustTemplate()

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
```

## Performance

troubled by `html/template` performance ? me too! [Check this out](https://github.com/mh-cbon/mdl-go-components/tree/master/mdl/fast) :D
