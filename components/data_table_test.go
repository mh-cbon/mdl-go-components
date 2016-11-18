package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestDataTable(t *testing.T) {

  var header *components.DataTableHeader
  var row *components.DataTableRow
	input := components.NewDataTable()
	header = input.SetHeader("id", "id")
	header.SetNumeric(true)
	header = input.SetHeader("name", "name")
	header = input.SetHeader("sku", "sku")
	header.SetHidePhone(true)
	header = input.SetHeader("edit", "")
	header.SetLinkIcon("edit")

	row = input.AddRow()
	row.SetValue("1")
	row.SetCell("id", "1")
	row.SetCell("name", "name1")
	row.SetCell("sku", "sku1")
	row.SetCell("edit", "http://google.com")

	row = input.AddRow()
	row.SetValue("2")
	row.SetCell("id", "2")
	row.SetCell("name", "name2")
	row.SetCell("sku", "sku2")
	row.SetCell("edit", "http://google.com")

	expectations := []string{
		`<table\s+class="mdl-data-table custom-js-data-table custom-data-table mdl-shadow--2dp\s+"\s+>`,
		`<th\s+class="mdl-data-table__cell--non-numeric"\s+>\s+id\s+</th>`,
		`<th\s+class=""\s+>\s+name\s+</th>`,
		`<th\s+class="mdl-cell--hide-phone"\s+>\s+sku\s+</th>`,
		`<tr\s+class=""\s+value="1">`,
		`<td\s+class="mdl-data-table__cell--non-numeric"\s+>\s+1\s+</td>`,
		`<td\s+class=""\s+>\s+name1\s+</td>`,
		`<td\s+class=""\s+>\s+sku1\s+</td>`,
		`<td\s+class=""\s+>\s+<a href="http://google.com" target="_blank">\s+<i class="material-icons">edit</i>\s+</a>\s+</td>`,
	}
	validateComponent(t, input, expectations)
}
