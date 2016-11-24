package components_common

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type ComponentsSlice struct {
	items []mgc.ViewComponentRenderer
}

func NewComponentsSlice() *ComponentsSlice {
	return &ComponentsSlice{}
}

func (slice *ComponentsSlice) SetDefaultRenderContext(r mgc.ContextRenderer) {
	for _, c := range slice.items {
		r.SetDefaultTo(c)
	}
}
func (slice *ComponentsSlice) SetRenderContext(r mgc.ContextRenderer) {
	for _, c := range slice.items {
		r.AttachTo(c)
	}
}
func (slice *ComponentsSlice) Render(args ...interface{}) (string, error) {
	for _, c := range slice.items {
		if _, err := c.Render(args); err != nil {
			return "", err
		}
	}
	return "", nil
}

func (slice *ComponentsSlice) GetComponents() []mgc.ViewComponentRenderer {
	return slice.items
}

func (slice *ComponentsSlice) AddComponent(some mgc.ViewComponentRenderer) {
	slice.items = append(slice.items, some)
}
func (slice *ComponentsSlice) Translate(t Translator) {
	for _, c := range slice.items {
		if v, ok := c.(NodeTranslator); ok {
			v.Translate(t)
		}
	}
}
func (slice *ComponentsSlice) SetErrors(p ErrorProvider) {
	for _, c := range slice.items {
		if v, ok := c.(NodeErrorsSetter); ok {
			v.SetErrors(p)
		}
	}
}
func (slice *ComponentsSlice) GetComponent(name string) mgc.ViewComponentRenderer {
	for _, c := range slice.items {
		if v, ok := c.(Namer); ok {
			if v.GetName() == name {
				return c
			}
		}
	}
	return nil
}
func (slice *ComponentsSlice) GetOptioner(name string) Optionner {
	c := slice.GetComponent(name)
	if c != nil {
		if v, ok := c.(Optionner); ok {
			return v
		}
	}
	return nil
}
func (slice *ComponentsSlice) GetValueSetter(name string) ValueSetter {
	c := slice.GetComponent(name)
	if c != nil {
		if v, ok := c.(ValueSetter); ok {
			return v
		}
	}
	return nil
}
func (slice *ComponentsSlice) GetValueDateSetter(name string) ValueDateSetter {
	c := slice.GetComponent(name)
	if c != nil {
		if v, ok := c.(ValueDateSetter); ok {
			return v
		}
	}
	return nil
}
func (slice *ComponentsSlice) GetValueSliceSetter(name string) ValueSliceSetter {
	c := slice.GetComponent(name)
	if c != nil {
		if v, ok := c.(ValueSliceSetter); ok {
			return v
		}
	}
	return nil
}
