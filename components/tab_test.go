package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestTab(t *testing.T) {
	tabs := components.NewTabs()

  text := components.NewText()
  text.SetContent("hello content!")
  tabs.Add("tab1", text)

  text = components.NewText()
  text.SetContent("hello content!")
  tabs.Add("tab2", text)

	expectations := []string{
		`<div\s+id="rnd-2"\s+class="mdl-tabs mdl-js-tabs mdl-js-ripple-effect\s+"\s+>`,
		`<a\s+href="#rnd-0"\s+class="mdl-tabs__tab is-active"\s+>tab1</a>`,
		`<a\s+href="#rnd-1"\s+class="mdl-tabs__tab "\s+>tab2</a>`,
		`<div\s+id="rnd-0"\s+class="mdl-tabs__panel is-active">\s+hello content!\s+</div>`,
		`<div\s+id="rnd-1"\s+class="mdl-tabs__panel ">\s+hello content!\s+</div>`,
	}
	validateComponent(t, tabs, expectations)
}

func TestTabSetActive(t *testing.T) {
	tabs := components.NewTabs()

  text := components.NewText()
  text.SetContent("hello content!")
  tabs.Add("tab1", text)

  text = components.NewText()
  text.SetContent("hello content!")
  tabs.Add("tab2", text)

  tabs.SetActive(1)

	expectations := []string{
		`<div\s+id="rnd-2"\s+class="mdl-tabs mdl-js-tabs mdl-js-ripple-effect\s+"\s+>`,
		`<a\s+href="#rnd-0"\s+class="mdl-tabs__tab "\s+>tab1</a>`,
		`<a\s+href="#rnd-1"\s+class="mdl-tabs__tab is-active"\s+>tab2</a>`,
		`<div\s+id="rnd-0"\s+class="mdl-tabs__panel\s+">\s+hello content!\s+</div>`,
		`<div\s+id="rnd-1"\s+class="mdl-tabs__panel is-active">\s+hello content!\s+</div>`,
	}
	validateComponent(t, tabs, expectations)
}
