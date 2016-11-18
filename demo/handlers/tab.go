package handlers

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"net/http"
)

func Tab(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Tabs",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var line *components.Slice
	var t1 *components.Tabs
	var dt *components.DataTable
	var header *components.DataTableHeader
	var row *components.DataTableRow
	var tx *components.Text

	// -
	line = components.NewSlice()

	t1 = components.NewTabs()

	tx = components.NewText()
	tx.SetContent("Hello from tab1!")

	t1.Add("Hello tab", tx)

	dt = components.NewDataTable()
	header = dt.SetHeader("id", "id")
	header.SetNumeric(true)
	header = dt.SetHeader("name", "name")
	header = dt.SetHeader("sku", "sku")
	header.SetHidePhone(true)
	header = dt.SetHeader("edit", "")
	header.SetLinkIcon("edit")

	row = dt.AddRow()
	row.SetValue("1")
	row.SetCell("id", "1")
	row.SetCell("name", "name1")
	row.SetCell("sku", "sku1")
	row.SetCell("edit", "http://google.com")

	row = dt.AddRow()
	row.SetValue("2")
	row.SetCell("id", "2")
	row.SetCell("name", "name2")
	row.SetCell("sku", "sku2")
	row.SetCell("edit", "http://google.com")

	t1.Add("Data table tab", dt)

	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
