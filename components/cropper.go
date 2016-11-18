package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
)

type Cropper struct {
	mgc.ViewComponent

	Node
	CurrentImgAttr    AttrList
	CurrentImgClasses ClassList
	CurrentImg        string

	InputHidden *Input
	InputFile   *InputFile
	Dialog      *Dialog
}

type GeometryResult struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Rotate float64 `json:"rotate"`
	ScaleX float64 `json:"scaleX"`
	ScaleY float64 `json:"scaleY"`
}

func NewCropper() *Cropper {
	ret := &Cropper{
		InputHidden: NewInputHidden(),
		InputFile:   NewInputFile(),
		Dialog:      NewDialog(),
	}
	ret.SetBlock("cropper")
	ret.SetResultMode("base64")
	ret.Dialog.Classes.Add("custom-cropper-dialog")
	Content := NewText()
	Content.SetHTMLContent(`
    <div class="custom-cropper-component-container">
      <img src="" class="custom-cropper-dialog-img" />
    </div>
    <div class="custom-cropper-preview">
      <img />
    </div>
  `)
	ret.Dialog.Content = Content
	return ret
}

func (view *Cropper) Render(args ...interface{}) (string, error) {
  view.GetRenderContext().SetDefaultTo(view.InputHidden)
  view.GetRenderContext().SetDefaultTo(view.InputFile)
  view.GetRenderContext().SetDefaultTo(view.Dialog)
	return view.GetRenderContext().RenderComponent(view, args)
}

func (c *Cropper) GetCurrentImg() string {
	return c.CurrentImg
}
func (c *Cropper) SetCurrentImg(b string) {
	c.CurrentImg = b
}

func (c *Cropper) SetDialogTitle(b string) {
	c.Dialog.SetTitle(b)
}
func (c *Cropper) SetDialogOkText(b string, icon string) {
	c.Dialog.SetOkText(b, icon)
}
func (c *Cropper) SetDialogCancelText(b string, icon string) {
	c.Dialog.SetCancelText(b, icon)
}
func (c *Cropper) SetDialogCloseText(b string, icon string) {
	c.Dialog.SetCloseText(b, icon)
}

func (c *Cropper) SetError(s string) {
	c.InputFile.SetError(s)
}
func (c *Cropper) GetError() string {
	return c.InputFile.GetError()
}
func (c *Cropper) SetSafeLabel(s string) {
	c.InputFile.SetSafeLabel(s)
}
func (c *Cropper) SetLabel(s interface{}) {
	c.InputFile.SetLabel(s)
}
func (c *Cropper) GetLabel() interface{} {
	return c.InputFile.GetLabel()
}
func (c *Cropper) SetPlaceHolderOnly(b bool) {
	c.InputFile.SetPlaceHolderOnly(b)
}
func (c *Cropper) IsPlaceHolderOnly() bool {
	return c.InputFile.IsPlaceHolderOnly()
}
func (c *Cropper) SetClearIcon(b string) {
	c.InputFile.SetClearIcon(b)
}
func (c *Cropper) GetClearIcon() string {
	return c.InputFile.GetClearIcon()
}
func (c *Cropper) SetAttachIcon(b string) {
	c.InputFile.SetAttachIcon(b)
}
func (c *Cropper) GetAttachIcon() string {
	return c.InputFile.GetAttachIcon()
}

func (c *Cropper) SetName(b string) {
	c.SetResultInputName(b)
}
func (c *Cropper) GetName() string {
	return c.GetResultInputName()
}
func (c *Cropper) SetResultInputName(b string) {
	c.InputHidden.SetName(b)
}
func (c *Cropper) GetResultInputName() string {
	return c.InputHidden.GetName()
}
func (c *Cropper) SetResultMode(b string) {
	c.InputHidden.Classes.Remove("custom-cropper-b64-result")
	c.InputHidden.Classes.Remove("custom-cropper-data-result")
	if b == "base64" {
		c.InputHidden.Classes.Add("custom-cropper-b64-result")
	} else if b == "geometry" {
		c.InputHidden.Classes.Add("custom-cropper-data-result")
	}
}
func (c *Cropper) GetResultMode() string {
	if c.InputHidden.Classes.Contains("custom-cropper-b64-result") {
		return "base64"
	} else if c.InputHidden.Classes.Contains("custom-cropper-b64-result") {
		return "geometry"
	}
	return ""
}
