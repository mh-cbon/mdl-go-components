package components_test

import (
	"testing"

	"github.com/mh-cbon/mdl-go-components/components"
)

func TestSlice(t *testing.T) {

	slice := components.NewSlice()

	text := components.NewText()
	text.SetContent("hello content!")
	slice.AddComponent(text)

	text = components.NewText()
	text.SetContent("hello other content!")
	slice.AddComponent(text)

	expectations := []string{
		`\s+hello content!\s+hello other content!\s+`,
	}
	validateComponent(t, slice, expectations)
}
