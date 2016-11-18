package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestSlice(t *testing.T) {

	slice := components.NewSlice()

	text := components.NewText()
	text.SetContent("hello content!")
	slice.Add(text)

	text = components.NewText()
	text.SetContent("hello other content!")
	slice.Add(text)

	expectations := []string{
		`\s+hello content!\s+hello other content!\s+`,
	}
	validateComponent(t, slice, expectations)
}
