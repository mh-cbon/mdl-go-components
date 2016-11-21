package mdlgocomponents

import (
	"github.com/mh-cbon/guess-path"
	"html/template"
)

var Tplpath = ""

func GuessGlob() string {
	p := guesspath.Glob(Tplpath, "mdl/templates/", "*.tmpl")
	if p == "" {
		panic("templates not found")
	}
	return p
}

func Template() (*template.Template, error) {
	return template.New("").ParseGlob(GuessGlob())
}

func MustTemplate() *template.Template {
	t, err := Template()
	if err != nil {
		panic(err)
	}
	return t
}
