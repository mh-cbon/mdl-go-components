package components

import (
	"math"
	"net/url"
	"strconv"

	mgc "github.com/mh-cbon/mdl-go-components"
	base "github.com/mh-cbon/mdl-go-components/components_common"
)

type Pagination struct {
	mgc.ViewComponent
	base.Node

	BtPrev        *Button
	BtNext        *Button
	Select        *Select
	GoToText      string
	BaseUrl       *url.URL
	PageParamName string
	CurrentPage   uint64
	PageCnt       uint64
}

func NewPagination() *Pagination {
	ret := &Pagination{}
	ret.SetBlock("mgc/pagination")
	ret.BtPrev = NewSubmit()
	ret.BtPrev.SetIcon("keyboard_arrow_left")
	ret.BtNext = NewSubmit()
	ret.BtNext.SetIcon("keyboard_arrow_right")
	ret.Select = NewSelect()
	ret.Select.UpdateWindowUrl(true)
	return ret
}
func (view *Pagination) SetBaseUrl(u *url.URL) {
	view.BaseUrl = u
	view.GuessCurrentPage()
}
func (view *Pagination) SetPageParamName(s string) {
	view.PageParamName = s
}
func (view *Pagination) HasPreviousPage() bool {
	return view.CurrentPage > 1
}
func (view *Pagination) HasNextPage() bool {
	return view.CurrentPage < view.GetPagesCnt()
}
func (view *Pagination) GuessCurrentPage() {
	q := view.BaseUrl.Query()
	view.CurrentPage = uint64(1)
	if x := q.Get(view.PageParamName); x != "" {
		i, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		view.CurrentPage = uint64(i)
	}
}
func (view *Pagination) SetPagesCnt(limit uint64, total uint64) {
	view.PageCnt = 1
	if total >= 2 {
		view.PageCnt = uint64(math.Ceil(float64(total) / float64(limit)))
	}
}
func (view *Pagination) GetPagesCnt() uint64 {
	return view.PageCnt
}
func (view *Pagination) GetUrl(p string) string {
	q := view.BaseUrl.Query()
	q.Set(view.PageParamName, p)
	view.BaseUrl.RawQuery = q.Encode()
	return view.BaseUrl.String()
}
func (view *Pagination) SetTotalItemCount(limit uint64, total uint64) {
	view.SetPagesCnt(limit, total)
	for i := uint64(1); i <= view.PageCnt; i++ {
		p := strconv.FormatUint(i, 10)
		view.Select.AddOption(view.GetUrl(p), p)
	}
	view.Select.SetValue(view.GetUrl(strconv.FormatUint(view.CurrentPage, 10)))
	view.BtPrev.SetDisabled(!view.HasPreviousPage())
	view.BtNext.SetDisabled(!view.HasNextPage())
	view.BtPrev.SetLink("#")
	view.BtNext.SetLink("#")
	if view.HasPreviousPage() {
		view.BtPrev.SetLink(view.GetUrl(strconv.FormatUint(view.CurrentPage-1, 10)))
	}
	if view.HasNextPage() {
		view.BtNext.SetLink(view.GetUrl(strconv.FormatUint(view.CurrentPage+1, 10)))
	}
}

func (view *Pagination) Translate(t base.Translator) {
	view.BtPrev.Translate(t)
	view.BtNext.Translate(t)
	view.Select.Translate(t)
}

func (view *Pagination) Render(args ...interface{}) (string, error) {
	view.GetRenderContext().AttachTo(view.BtPrev)
	view.GetRenderContext().AttachTo(view.BtNext)
	view.GetRenderContext().AttachTo(view.Select)
	return view.GetRenderContext().RenderComponent(view, args)
}
