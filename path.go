package mdlgocomponents

import (
	"github.com/mh-cbon/this-path"
)

func TemplatesPath() string {
	return thispath.Dir() + "/mdl/templates/"
}

func TemplatesGlob() string {
	return thispath.Dir() + "/mdl/templates/*.tmpl"
}
