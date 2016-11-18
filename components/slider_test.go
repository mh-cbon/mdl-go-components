package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestSlider(t *testing.T) {
	slider := components.NewSlider()

	slider.SetMin("2")
	slider.SetMax("5000")
	slider.SetStep("85")

	expectations := []string{
		`<div\s+class="mdl-textfield\s+mdl-textfield--floating-label\s+">`,
		`<input\s+type="range"\s+min="2"\s+max="5000"\s+step="85"\s+id="rnd-0"\s+class="mdl-slider mdl-js-slider\s+"\s+value=""\s+>`,
		`<label class="mdl-textfield__label" for="rnd-0"></label>`,
		`<span class="mdl-textfield__error"></span>`,
	}
	validateComponent(t, slider, expectations)
}

func TestSliderWithLabel(t *testing.T) {
	slider := components.NewSlider()

	slider.SetLabel("the label")

	expectations := []string{
		`<label class="mdl-textfield__label" for="rnd-0">the label</label>`,
	}
	validateComponent(t, slider, expectations)
}

func TestSliderWithError(t *testing.T) {
	slider := components.NewSlider()

	slider.SetError("the error")

	expectations := []string{
		`<span class="mdl-textfield__error">the error</span>`,
	}
	validateComponent(t, slider, expectations)
}
