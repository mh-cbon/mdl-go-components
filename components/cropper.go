package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
  "html/template"
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
	ret.SetBlock("mgc/cropper")
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
func (view *Cropper) Translate(t Translator) {
	view.InputFile.Translate(t)
	view.Dialog.Translate(t)
}
func (view *Cropper) SetErrors(p ErrorProvider) {
  err := p.GetError(view.GetName())
  if err!=nil {
    view.SetError(err)
  }
	view.Dialog.SetErrors(p)
}
func (view *Cropper) Render(args ...interface{}) (string, error) {
	view.GetRenderContext().SetDefaultTo(view.InputHidden)
	view.GetRenderContext().SetDefaultTo(view.InputFile)
	view.GetRenderContext().SetDefaultTo(view.Dialog)
	return view.GetRenderContext().RenderComponent(view, args)
}

func (c *Cropper) GetCurrentImgAttr() template.HTMLAttr {
	return template.HTMLAttr("src='"+c.CurrentImg+"'")
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

func (c *Cropper) SetError(s interface{}) {
	c.InputFile.SetError(s)
}
func (c *Cropper) GetError() interface{} {
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
func (view *Cropper) SetInputTextValue(b string) {
	view.InputFile.SetInputTextValue(b)
}
func (view *Cropper) SetInputTextName(b string) {
	view.InputFile.SetInputTextName(b)
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
func (c *Cropper) SetValue(b string) {
	c.InputHidden.SetValue(b)
}
func (c *Cropper) GetValue() string {
	return c.InputHidden.GetValue()
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
func (c *Cropper) RoundedPreview(b bool) {
  if b {
    c.CurrentImgClasses.Add("custom-cropper-current-img--rounded")
  } else {
    c.CurrentImgClasses.Remove("custom-cropper-current-img--rounded")
  }
}
func (c *Cropper) SetB64ExportWidth(b string) {
  if b == "" {
    c.Attr.Remove("b64-export-width")
  } else {
    c.Attr.Set("b64-export-width", b)
  }
}
func (c *Cropper) SetB64ExportHeight(b string) {
  if b == "" {
    c.Attr.Remove("b64-export-height")
  } else {
    c.Attr.Set("b64-export-height", b)
  }
}
func (c *Cropper) SetAspectRatio(b string) {
  if b == "" {
    c.Attr.Remove("aspect-ratio")
  } else {
    c.Attr.Set("aspect-ratio", b)
  }
}
func (c *Cropper) SetMovable(b bool) {
  if b {
    c.Attr.Set("movable", "true")
  } else {
    c.Attr.Set("movable", "false")
  }
}
func (c *Cropper) SetScalable(b bool) {
  if b {
    c.Attr.Set("scalable", "true")
  } else {
    c.Attr.Set("scalable", "false")
  }
}
func (c *Cropper) SetRotatable(b bool) {
  if b {
    c.Attr.Set("rotatable", "true")
  } else {
    c.Attr.Set("rotatable", "false")
  }
}
func (c *Cropper) SetDragMode(b string) {
  if b == "" {
    c.Attr.Remove("drag-mode")
  } else {
    c.Attr.Set("drag-mode", b)
  }
}
// many more like this,
// see https://github.com/mh-cbon/material-design-lite/blob/mdl-1.x/src/custom-cropper/cropper.js#L223
