package components_test

import (
	"bytes"
	mgc "github.com/mh-cbon/mdl-go-components"
	"regexp"
	"testing"
	"html/template"
)

var tmpl *template.Template

func init() {
	tmpl = mgc.MustTemplate()
}

func validateComponent(t *testing.T, component mgc.ViewComponentRenderer, expectations []string) {

	var wr bytes.Buffer
	ctx := mgc.NewRenderContext(tmpl, &wr)
	ctx.SetDefaultTo(component)
	if _, err := component.Render(); err != nil {
		t.Error(err)
	}

	got := wr.String()

	if len(expectations) == 0 {
		t.Errorf("You must provide expectations to validate this output:\n'%v'\n---", got)
	}

	for _, expectation := range expectations {
		matched, err := regexp.MatchString(expectation, got)
		if err != nil {
			t.Error(err)
		} else if matched == false {
			t.Errorf("Expectation\n'%v'\ndid not match\n'%v'\n---", expectation, got)
		}
	}
}
