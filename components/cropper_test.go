package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestCropper(t *testing.T) {

	input := components.NewCropper()
	input.SetCurrentImg("/static/avatar-1.png")
	input.SetName("result")
	input.SetLabel("base64 result")
	input.SetError("beeep boop nop")

	expectations := []string{
		`<div\s+id="rnd-0"\s+class="custom-cropper custom-js-cropper\s+"\s+>`,
		`<div\s+class="custom-cropper-current-img\s+">\s+<img\s+src="/static/avatar-1.png"\s+/>\s+</div>`,
		`<input\s+class="custom-cropper-b64-result"\s+type="hidden"\s+name="result"\s+value=""\s+/>`,
		`<label\s+class="mdl-textfield__label"\s+for="rnd-2">base64 result</label>`,
		`<span\s+class="mdl-textfield__error">beeep boop nop</span>`,
		`<div\s+class="custom-js-dialog custom-dialog custom-cropper-dialog"\s+id="rnd-3"\s+>`,
	}
	validateComponent(t, input, expectations)
}

func TestCropperGeometry(t *testing.T) {

	input := components.NewCropper()
	input.SetResultMode("geometry")

	expectations := []string{
		`<input\s+class="custom-cropper-data-result"\s+type="hidden"\s+name=""\s+value=""\s+/>`,
	}
	validateComponent(t, input, expectations)
}
