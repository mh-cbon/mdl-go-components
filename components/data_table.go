package components

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"time"
)

type DataTable struct {
	mgc.ViewComponent
	Headers []*DataTableHeader
	Rows    []*DataTableRow

	Classes ClassList
	Attr    AttrList

	Empty mgc.ViewComponentRenderer

	SortIcon string
}

func NewDataTable() *DataTable {
	ret := &DataTable{}
	ret.SetBlock("data_table")
	return ret
}
func (view *DataTable) Render(args ...interface{}) (string, error) {
	if k := view.GetSortIcon(); k != "" {
		for _, header := range view.Headers {
			header.SetSortIcon(k)
		}
	}
	for _, row := range view.Rows {
		for _, cell := range row.Cells {
			cell.SetHidePhone(view.IsHidePhone(cell.GetCellName()))
			cell.SetHideTablet(view.IsHideTablet(cell.GetCellName()))
			cell.SetNumeric(view.IsNumeric(cell.GetCellName()))
			if view.IsLink(cell.GetCellName()) {
				cell.SetLinkIcon(view.GetLinkIcon(cell.GetCellName()))
			}
		}
	}
	return view.GetRenderContext().RenderComponent(view, args)
}
func (l *DataTable) AddRow() *DataTableRow {
	ret := &DataTableRow{}
	l.Rows = append(l.Rows, ret)
	return ret
}
func (l *DataTable) GetBtSelector() string {
	return l.Attr.GetValue("bt-el")
}
func (l *DataTable) SetBtSelector(some string) {
	if some == "" {
		l.Attr.Remove("bt-el")
	} else {
		l.Attr.Set("bt-el", some)
	}
}
func (l *DataTable) SetHeader(name string, txt string) *DataTableHeader {
	v := l.GetHeader(name)
	if v == nil {
		v = l.AddHeader(name, txt)
	} else {
		v.CellName = name
		v.CellTxt = txt
	}
	return v
}
func (l *DataTable) AddHeader(name string, txt string) *DataTableHeader {
	v := &DataTableHeader{}
	l.Headers = append(l.Headers, v)
	v.CellName = name
	v.CellTxt = txt
	return v
}
func (l *DataTable) GetHeader(name string) *DataTableHeader {
	for _, header := range l.Headers {
		if header.GetCellName() == name {
			return header
		}
	}
	return nil
}

func (l *DataTable) SetSelectable(b bool) *DataTable {
	if b {
		l.Classes.Add("mdl-data-table--selectable")
	} else {
		l.Classes.Remove("mdl-data-table--selectable")
	}
	return l
}
func (l *DataTable) IsSelectable() bool {
	return l.Classes.Contains("mdl-data-table--selectable")
}

func (l *DataTable) SetSortIcon(icon string) *DataTable {
	l.SortIcon = icon
	return l
}
func (l *DataTable) GetSortIcon() string {
	return l.SortIcon
}

func (l *DataTable) IsNumeric(name string) bool {
	h := l.GetHeader(name)
	if h == nil {
		return false
	}
	return h.IsNumeric()
}
func (l *DataTable) IsLink(name string) bool {
	h := l.GetHeader(name)
	if h == nil {
		return false
	}
	return h.GetLinkIcon() != ""
}
func (l *DataTable) IsHidePhone(name string) bool {
	h := l.GetHeader(name)
	if h == nil {
		return false
	}
	return h.IsHidePhone()
}
func (l *DataTable) IsHideTablet(name string) bool {
	h := l.GetHeader(name)
	if h == nil {
		return false
	}
	return h.IsHideTablet()
}
func (l *DataTable) GetLinkIcon(name string) string {
	h := l.GetHeader(name)
	if h == nil {
		return ""
	}
	return h.GetLinkIcon()
}

type DataTableHeader struct {
	Classes ClassList
	Attr    AttrList

	CellName string
	CellTxt  string

	LinkIcon string
	SortDir  string
	SortIcon string
}

func (l *DataTableHeader) SetCellTxt(b string) *DataTableHeader {
	l.CellTxt = b
	return l
}
func (l *DataTableHeader) GetCellTxt() string {
	return l.CellTxt
}

func (l *DataTableHeader) SetCellName(b string) *DataTableHeader {
	l.CellName = b
	return l
}
func (l *DataTableHeader) GetCellName() string {
	return l.CellName
}

func (l *DataTableHeader) SetHidePhone(b bool) *DataTableHeader {
	if b {
		l.Classes.Add("mdl-cell--hide-phone")
	} else {
		l.Classes.Remove("mdl-cell--hide-phone")
	}
	return l
}
func (l *DataTableHeader) IsHidePhone() bool {
	return l.Classes.Contains("mdl-cell--hide-phone")
}

func (l *DataTableHeader) SetHideTablet(b bool) *DataTableHeader {
	if b {
		l.Classes.Add("mdl-cell--hide-phone")
	} else {
		l.Classes.Remove("mdl-cell--hide-phone")
	}
	return l
}
func (l *DataTableHeader) IsHideTablet() bool {
	return l.Classes.Contains("mdl-cell--hide-tablet")
}

func (l *DataTableHeader) SetSortable(sortable bool) *DataTableHeader {
	if sortable {
		l.Classes.Add("mdl-data-table__header--sorted")
	} else {
		l.Classes.Remove("mdl-data-table__header--sorted")
		l.Classes.Remove("mdl-data-table__header--sorted-descending")
		l.Classes.Remove("mdl-data-table__header--sorted-descending")
	}
	return l
}
func (l *DataTableHeader) IsSortable() bool {
	return l.Classes.Contains("mdl-data-table__header--sorted")
}
func (l *DataTableHeader) SetSortdir(s string) *DataTableHeader {
	if s == "" {
		l.Classes.Add("mdl-data-table__header--sorted")
		l.Classes.Remove("mdl-data-table__header--sorted-ascending")
		l.Classes.Remove("mdl-data-table__header--sorted-descending")
	} else {
		l.Classes.Remove("mdl-data-table__header--sorted-ascending")
		l.Classes.Remove("mdl-data-table__header--sorted-descending")
		l.Classes.Add("mdl-data-table__header--sorted-" + s + "ending")
	}
	return l
}
func (l *DataTableHeader) GetSortdir() string {
	if l.Classes.Contains("mdl-data-table__header--sorted-ascending") {
		return "asc"
	} else if l.Classes.Contains("mdl-data-table__header--sorted-descending") {
		return "desc"
	}
	return ""
}
func (l *DataTableHeader) GetNextSortdir() string {
	k := l.GetSortdir()
	if k == "" {
		return "asc"
	} else if k == "asc" {
		return "desc"
	}
	return ""
}
func (l *DataTableHeader) GetNextSortdirParam(name string) string {
	k := l.GetNextSortdir()
	if k == "" {
		return ""
	}
	return name + "=" + k
}

func (l *DataTableHeader) SetSortIcon(icon string) *DataTableHeader {
	l.SortIcon = icon
	return l
}
func (l *DataTableHeader) GetSortIcon() string {
	return l.SortIcon
}

func (l *DataTableHeader) SetLinkIcon(icon string) *DataTableHeader {
	l.LinkIcon = icon
	return l
}
func (l *DataTableHeader) GetLinkIcon() string {
	return l.LinkIcon
}

func (l *DataTableHeader) SetNumeric(numeric bool) *DataTableHeader {
	if numeric {
		l.Classes.Add("mdl-data-table__cell--non-numeric")
	} else {
		l.Classes.Remove("mdl-data-table__cell--non-numeric")
	}
	return l
}
func (l *DataTableHeader) IsNumeric() bool {
	return l.Classes.Contains("mdl-data-table__cell--non-numeric")
}

type DataTableRow struct {
	Classes ClassList
	Attr    AttrList

	NodeSingleValue
	Cells []*DataTableCell
}

func (l *DataTableRow) GetCell(name string) *DataTableCell {
	for _, cell := range l.Cells {
		if cell.GetCellName() == name {
			return cell
		}
	}
	return nil
}
func (l *DataTableRow) SetCell(name string, value string) *DataTableRow {
	cell := l.GetCell(name)
	if cell == nil {
		cell = NewDataTableCell()
		l.Cells = append(l.Cells, cell)
	}
	cell.SetCellName(name)
	cell.SetCellTxt(value)
	return l
}

type DataTableCell struct {
	Classes ClassList
	Attr    AttrList

	CellName string
	CellTxt  string
	LinkIcon string
}

func NewDataTableCell() *DataTableCell {
	return &DataTableCell{}
}

func (l *DataTableCell) SetCellTxt(b string) *DataTableCell {
	l.CellTxt = b
	return l
}
func (l *DataTableCell) GetCellTxt() string {
	return l.CellTxt
}

func (l *DataTableCell) SetCellName(b string) *DataTableCell {
	l.CellName = b
	return l
}
func (l *DataTableCell) GetCellName() string {
	return l.CellName
}

func (l *DataTableCell) SetTime(value time.Time) {
	l.SetCellTxt(value.Format("Mon Jan _2 15:04:05 2006"))
}
func (l *DataTableCell) SetDate(value time.Time) {
	l.SetCellTxt(value.Format("Mon Jan _2"))
}

func (l *DataTableCell) SetHidePhone(b bool) *DataTableCell {
	if b {
		l.Classes.Add("mdl-cell--hide-phone")
	} else {
		l.Classes.Remove("mdl-cell--hide-phone")
	}
	return l
}
func (l *DataTableCell) IsHidePhone() bool {
	return l.Classes.Contains("mdl-cell--hide-phone")
}

func (l *DataTableCell) SetHideTablet(b bool) *DataTableCell {
	if b {
		l.Classes.Add("mdl-cell--hide-phone")
	} else {
		l.Classes.Remove("mdl-cell--hide-phone")
	}
	return l
}
func (l *DataTableCell) IsHideTablet() bool {
	return l.Classes.Contains("mdl-cell--hide-tablet")
}

func (l *DataTableCell) SetLinkIcon(icon string) *DataTableCell {
	l.LinkIcon = icon
	return l
}
func (l *DataTableCell) GetLinkIcon() string {
	return l.LinkIcon
}

func (l *DataTableCell) SetNumeric(numeric bool) *DataTableCell {
	if numeric {
		l.Classes.Add("mdl-data-table__cell--non-numeric")
	} else {
		l.Classes.Remove("mdl-data-table__cell--non-numeric")
	}
	return l
}
func (l *DataTableCell) IsNumeric() bool {
	return l.Classes.Contains("mdl-data-table__cell--non-numeric")
}
