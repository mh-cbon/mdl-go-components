package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestDialog(t *testing.T) {

	dialog := components.NewDialog()
	dialog.SetTitle("The title")
	dialog.SetId("id")

	dialog.SetOkText("OkText", "")
	dialog.SetOkAttr("type", "submit")
	dialog.SetOkAttr("form", "form1")

	dialog.SetCancelText("CancelText", "")
	dialog.SetCancelAttr("a", "b")

	dialog.SetCloseText("CloseText", "")
	dialog.SetCloseAttr("c", "d")

	text := components.NewText()
	text.SetContent("hello content!")
	dialog.Content = text

	expectations := []string{
		`<div\s+class="custom-js-dialog custom-dialog\s+"\s+id="id"\s+>`,
		`<h4 class="mdl-dialog__title">The title</h4>`,
		`<div class="mdl-dialog__content">\s+hello content!\s+</div>`,

		`<button[^>]+>OkText</button>`,
		`<button[^>]+type="submit"[^>]+>OkText</button>`,
		`<button[^>]+form="form1"[^>]+>OkText</button>`,

		`<button[^>]+class="mdl-button mdl-js-button custom-dialog-cancel"[^>]+>CancelText</button>`,
		`<button[^>]+type="button"[^>]+a="b"[^>]+>CancelText</button>`,

		`<button[^>]+class="mdl-button mdl-js-button custom-dialog-close"[^>]+>CloseText</button>`,
		`<button[^>]+type="button"[^>]+c="d"[^>]+>CloseText</button>`,
	}
	validateComponent(t, dialog, expectations)
}
