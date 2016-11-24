package components

import mgc "github.com/mh-cbon/mdl-go-components"

type Slice struct {
	mgc.ViewComponent
	*ComponentWithAComponentsSlice
}

func NewSlice() *Slice {
	ret := &Slice{}
	ret.ComponentWithAComponentsSlice = NewComponentWithAComponentsSlice()
	return ret
}

// func (view *Slice) Add(s mgc.ViewComponentRenderer) {
// 	view.Components = append(view.Components, s)
// }

func (view *Slice) Render(args ...interface{}) (string, error) {
	view.GetRenderContext().SetDefaultTo(view.Components)
	return view.Components.Render(args)
}

// func (view *Slice) Translate(t base.Translator) {
// 	for _, c := range view.Components {
// 		if v, ok := c.(base.NodeTranslator); ok {
// 			v.Translate(t)
// 		}
// 	}
// }
// func (view *Slice) SetErrors(p base.ErrorProvider) {
// 	for _, c := range view.Components {
// 		if v, ok := c.(base.NodeErrorsSetter); ok {
// 			v.SetErrors(p)
// 		}
// 	}
// }
