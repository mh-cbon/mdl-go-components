package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"time"
)

type InputDate struct {
	mgc.ViewComponent
	NodeLabel
	NodePlaceholder
	NodeSingleError

	Attr         AttrList
	Classes      ClassList
	InputAttr    AttrList
	InputClasses ClassList
	HiddenAttr   AttrList

	GoFormat string
}

func NewInputDate() *InputDate {
	ret := &InputDate{}
	ret.SetGoFormat(time.RFC3339)
	ret.SetBlock("mgc/form_date")
	return ret
}

func (view *InputDate) SetErrors(p ErrorProvider) {
	err := p.GetError(view.GetName())
	if err!=nil {
		view.SetError(err)
	}
}

func (i *InputDate) SetGoFormat(b string) {
	i.GoFormat = b
}
func (i *InputDate) GetGoFormat() string {
	return i.GoFormat
}

func (view *InputDate) SetName(b string) {
	view.HiddenAttr.Set("name", b)
}
func (view *InputDate) GetName() string {
	return view.HiddenAttr.GetValue("name")
}

func (view *InputDate) SetId(b string) {
	view.InputAttr.Set("id", b)
}
func (view *InputDate) GetId() string {
	return view.InputAttr.GetValue("id")
}

func (i *InputDate) SetDate(b time.Time) {
	i.SetPresent(b)
}
func (i *InputDate) GetDate() (time.Time, error) {
  d := i.GetPresent()
	return *d, nil
}
func (i *InputDate) MustGetDate() time.Time {
  d := i.GetPresent()
	return *d
}
func (i *InputDate) SetPresent(b time.Time) {
	i.HiddenAttr.Set("value", b.Format(i.GoFormat))
}
func (i *InputDate) GetPresent() *time.Time {
	v := i.HiddenAttr.GetValue("value")
	if v == "" {
		return nil
	}
	d, _ := time.Parse(i.GoFormat, v)
	return &d
}

func (i *InputDate) SetFuture(b time.Time) {
	i.Attr.Set("future", b.Format(i.GoFormat))
}
func (i *InputDate) GetFuture() *time.Time {
	v := i.Attr.GetValue("future")
	if v == "" {
		return nil
	}
	d, _ := time.Parse(i.GoFormat, v)
	return &d
}

func (i *InputDate) SetPast(b time.Time) {
	i.Attr.Set("past", b.Format(i.GoFormat))
}
func (i *InputDate) GetPast() *time.Time {
	v := i.Attr.GetValue("past")
	if v == "" {
		return nil
	}
	d, _ := time.Parse(i.GoFormat, v)
	return &d
}

func (i *InputDate) SetColon(b bool) {
	if b {
		i.Attr.Set("colon", "true")
	} else {
		i.Attr.Remove("colon")
	}
}
func (i *InputDate) IsColon() bool {
	return i.Attr.GetValue("colon") == "true"
}

func (i *InputDate) SetCancelText(b string) {
	i.Attr.Set("cancel", b)
}
func (i *InputDate) GetCancelText() string {
	return i.Attr.GetValue("cancel")
}

func (i *InputDate) SetOkText(b string) {
	i.Attr.Set("ok", b)
}
func (i *InputDate) GetOkText() string {
	return i.Attr.GetValue("ok")
}

func (i *InputDate) SetOrientation(b string) {
	i.Attr.Set("orientation", b)
}
func (i *InputDate) GetOrientation() string {
	return i.Attr.GetValue("orientation")
}

func (i *InputDate) SetDisplayMode(b string) {
	i.Attr.Set("mode", b)
}
func (i *InputDate) GetDisplayMode() string {
	return i.Attr.GetValue("mode")
}

func (i *InputDate) SetDisplayFormat(b string) {
	i.Attr.Set("displayformat", b)
}
func (i *InputDate) GetDisplayFormat() string {
	return i.Attr.GetValue("displayformat")
}

func (i *InputDate) SetFormat(b string) {
	i.Attr.Set("format", b)
}
func (i *InputDate) GetFormat() string {
	return i.Attr.GetValue("format")
}

func (i *InputDate) SetDisplayLocale(b string) {
	i.Attr.Set("displaylocale", b)
}
func (i *InputDate) GetDisplayLocale() string {
	return i.Attr.GetValue("displaylocale")
}

func (view *InputDate) Render(args ...interface{}) (string, error) {
	return view.GetRenderContext().RenderComponent(view, args)
}
