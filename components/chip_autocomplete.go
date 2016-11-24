package components

import (
	"encoding/json"

	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type ChipAutocomplete struct {
	mgc.ViewComponent

	base.Node
	base.NodeWithOptions

	ResultsAttr    base.AttrList
	ResultsClasses base.ClassList

	SelectAttr    base.AttrList
	SelectClasses base.ClassList

	Input *Input
}

func NewChipAutocomplete() *ChipAutocomplete {
	ret := &ChipAutocomplete{
		Input: NewInputText(),
	}
	ret.SetBlock("mgc/chip_autocomplete")
	return ret
}

func (view *ChipAutocomplete) Translate(t base.Translator) {
	view.Input.Translate(t)
	view.NodeWithOptions.Translate(t)
}
func (view *ChipAutocomplete) SetErrors(p base.ErrorProvider) {
	err := p.GetError(view.GetName())
	if err != nil {
		view.SetError(err)
	}
}

func (view *ChipAutocomplete) Render(args ...interface{}) (string, error) {
	view.GetRenderContext().SetDefaultTo(view.Input)
	return view.GetRenderContext().RenderComponent(view, args)
}

func (c *ChipAutocomplete) SetName(b string) {
	c.SetOptionValueName(b)
}
func (c *ChipAutocomplete) GetName() string {
	return c.GetOptionValueName()
}

func (c *ChipAutocomplete) SetError(s interface{}) {
	c.Input.SetError(s)
}
func (c *ChipAutocomplete) GetError() interface{} {
	return c.Input.GetError()
}
func (c *ChipAutocomplete) SetSafeLabel(s string) {
	c.Input.SetSafeLabel(s)
}
func (c *ChipAutocomplete) SetLabel(s interface{}) {
	c.Input.SetLabel(s)
}
func (c *ChipAutocomplete) GetLabel() interface{} {
	return c.Input.GetLabel()
}
func (c *ChipAutocomplete) SetPlaceHolderOnly(b bool) {
	c.Input.SetPlaceHolderOnly(b)
}
func (c *ChipAutocomplete) IsPlaceHolderOnly() bool {
	return c.Input.IsPlaceHolderOnly()
}

func (c *ChipAutocomplete) GetUrlCompleter() string {
	return c.Attr.GetValue("url-completer")
}
func (c *ChipAutocomplete) SetUrlCompleter(b string) {
	c.Attr.Set("url-completer", b)
}
func (c *ChipAutocomplete) GetUrlCompleterArgs() string {
	return c.Attr.GetValue("url-completer-args")
}
func (c *ChipAutocomplete) SetUrlCompleterArgs(b string) {
	c.Attr.Set("url-completer-args", b)
}
func (c *ChipAutocomplete) SetEncodedUrlCompleterArgs(b interface{}) error {
	bolB, err := json.Marshal(b)
	if err == nil {
		c.SetUrlCompleterArgs(string(bolB))
	}
	return err
}

func (c *ChipAutocomplete) GetUrlCreator() string {
	return c.Attr.GetValue("url-creator")
}
func (c *ChipAutocomplete) SetUrlCreator(b string) {
	c.Attr.Set("url-creator", b)
}
func (c *ChipAutocomplete) GetUrlCreatorArgs() string {
	return c.Attr.GetValue("url-creator-args")
}
func (c *ChipAutocomplete) SetUrlCreatorArgs(b string) {
	c.Attr.Set("url-creator-args", b)
}
func (c *ChipAutocomplete) SetEncodedUrlCreatorArgs(b interface{}) error {
	bolB, err := json.Marshal(b)
	if err == nil {
		c.SetUrlCreatorArgs(string(bolB))
	}
	return err
}

func (c *ChipAutocomplete) GetUrlPlaceholder() string {
	return c.Attr.GetValue("url-placeholder")
}
func (c *ChipAutocomplete) SetUrlPlaceholder(b string) {
	c.Attr.Set("url-placeholder", b)
}

func (c *ChipAutocomplete) GetUnreachableRemoteTxt() string {
	return c.Attr.GetValue("txt-remote-unreachable")
}
func (c *ChipAutocomplete) SetUnreachableRemoteTxt(b string) {
	c.Attr.Set("txt-remote-unreachable", b)
}

func (c *ChipAutocomplete) GetOptionValueName() string {
	return c.Attr.GetValue("chip-name")
}
func (c *ChipAutocomplete) SetOptionValueName(b string) {
	c.Attr.Set("chip-name", b)
}
