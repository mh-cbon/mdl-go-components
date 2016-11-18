package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestButton(t *testing.T) {
	input := components.NewButton()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetValue("plop")

	expectations := []string{
		`<button[^>]+>label</button>`,
		`<button[^>]+type="button"[^>]+>label</button>`,
		`<button[^>]+name="name"[^>]+>label</button>`,
		`<button[^>]+id="id"[^>]+>label</button>`,
	}
	validateComponent(t, input, expectations)
}

func TestButtonHref(t *testing.T) {
	input := components.NewButton()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetLink("http://whatever")

	expectations := []string{
		`<a[^>]+>label</a>`,
		`<a[^>]+name="name"[^>]+>label</a>`,
		`<a[^>]+href="http://whatever"[^>]+>label</a>`,
		`<a[^>]+class="mdl-button mdl-js-button"[^>]+>label</a>`,
		`<a[^>]+id="id"[^>]+>label</a>`,
	}
	validateComponent(t, input, expectations)
}

func TestSubmit(t *testing.T) {
	input := components.NewSubmit()
	input.SetLabel("label")
	input.SetName("name")
	input.SetId("id")
	input.SetValue("plop")

	expectations := []string{
		`<button[^>]+>label</button>`,
		`<button[^>]+type="submit"[^>]+>label</button>`,
		`<button[^>]+name="name"[^>]+>label</button>`,
		`<button[^>]+id="id"[^>]+>label</button>`,
	}
	validateComponent(t, input, expectations)
}
