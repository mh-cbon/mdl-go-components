package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Slice struct {
	mgc.ViewComponent
	Components []mgc.ViewComponentRenderer
}

func NewSlice() *Slice {
	ret := &Slice{}
	return ret
}

func (view *Slice) Add(s mgc.ViewComponentRenderer) {
	view.Components = append(view.Components, s)
}

func (view *Slice) Render(args ...interface{}) (string, error) {
	for _, c := range view.Components {
		view.GetRenderContext().SetDefaultTo(c)
		if _, err := c.Render(args); err != nil {
			return "", err
		}
	}
	return "", nil
}
