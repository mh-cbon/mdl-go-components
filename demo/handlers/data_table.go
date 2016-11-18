package handlers

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"net/http"
)

func DataTable(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Data table",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var line *components.Slice
	var t1 *components.DataTable
	var header *components.DataTableHeader
	var row *components.DataTableRow
	var bt *components.Button
	var dialog *components.Dialog

	// -
	line = components.NewSlice()

	t1 = components.NewDataTable()
	header = t1.SetHeader("id", "id")
	header.SetNumeric(true)
	header = t1.SetHeader("name", "name")
	header = t1.SetHeader("sku", "sku")
	header.SetHidePhone(true)
	header = t1.SetHeader("edit", "")
	header.SetLinkIcon("edit")

	row = t1.AddRow()
	row.SetValue("1")
	row.SetCell("id", "1")
	row.SetCell("name", "name1")
	row.SetCell("sku", "sku1")
	row.SetCell("edit", "http://google.com")

	row = t1.AddRow()
	row.SetValue("2")
	row.SetCell("id", "2")
	row.SetCell("name", "name2")
	row.SetCell("sku", "sku2")
	row.SetCell("edit", "http://google.com")

	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewDataTable()
	t1.SetSelectable(true)
	header = t1.SetHeader("id", "id")
	header.SetNumeric(true)
	header = t1.SetHeader("name", "name")
	header = t1.SetHeader("sku", "sku")
	header.SetHidePhone(true)
	header = t1.SetHeader("edit", "")
	header.SetLinkIcon("edit")

	row = t1.AddRow()
	row.SetValue("1")
	row.SetCell("id", "1")
	row.SetCell("name", "name1")
	row.SetCell("sku", "sku1")
	row.SetCell("edit", "http://google.com")

	row = t1.AddRow()
	row.SetValue("2")
	row.SetCell("id", "2")
	row.SetCell("name", "name2")
	row.SetCell("sku", "sku2")
	row.SetCell("edit", "http://google.com")

	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewDataTable()
	t1.SetSelectable(true)
	t1.SetSortIcon("sort_by_alpha")
	t1.SetBtSelector("#button1")

	header = t1.SetHeader("id", "id")
	header.SetNumeric(true)
	header.SetSortable(true)
	header.SetSortdir("")
	header = t1.SetHeader("name", "name")
	header.SetSortable(true)
	header.SetSortdir("desc")
	header = t1.SetHeader("sku", "sku")
	header.SetHidePhone(true)
	header.SetSortable(true)
	header.SetSortdir("asc")
	header = t1.SetHeader("edit", "")
	header.SetLinkIcon("edit")

	row = t1.AddRow()
	row.SetValue("1")
	row.SetCell("id", "1")
	row.SetCell("name", "name1")
	row.SetCell("sku", "sku1")
	row.SetCell("edit", "http://google.com")

	row = t1.AddRow()
	row.SetValue("2")
	row.SetCell("id", "2")
	row.SetCell("name", "name2")
	row.SetCell("sku", "sku2")
	row.SetCell("edit", "http://google.com")

	line.Add(t1)

	bt = components.NewButton()
	bt.SetLabel("button")
	bt.SetName("button1")
	bt.SetId("button1")
	bt.Classes.Add("custom-js-confirm-button")
	bt.Attr.Set("confirm", "#dialog")
	line.Add(bt)

	dialog = components.NewDialog()
	dialog.SetTitle("The title")
	dialog.SetId("dialog")

	dialog.Ok.SetType("submit")
	dialog.Ok.Attr.Set("form", "form1")

	text := components.NewText()
	text.SetContent("hello content!")
	dialog.Content = text

	line.Add(dialog)

	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
