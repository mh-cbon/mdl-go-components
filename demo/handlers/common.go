package handlers

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"html/template"
	"net/http"
)

func SetRoutes() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/textfield", Textfield)
	http.HandleFunc("/button", Button)
	http.HandleFunc("/checkbox", Checkbox)
	http.HandleFunc("/checkbox_slice", Checkbox_slice)
	http.HandleFunc("/select", SelectField)
	http.HandleFunc("/textarea", Textarea)
	http.HandleFunc("/input_confirm", InputConfirm)
	http.HandleFunc("/input_date", InputDate)
	http.HandleFunc("/input_file", InputFile)
	http.HandleFunc("/slider", Slider)
	http.HandleFunc("/expand", Expand)
	http.HandleFunc("/dialog", Dialog)
	http.HandleFunc("/data_table", DataTable)
	http.HandleFunc("/tab", Tab)
	http.HandleFunc("/cropper", Cropper)
	http.HandleFunc("/cropper_post", CropperPost)
	http.HandleFunc("/chip_autocomplete", ChipAutocomplete)
	http.HandleFunc("/chip_autocomplete_list", ChipAutocompleteList)
	http.HandleFunc("/chip_autocomplete_create", ChipAutocompleteCreate)
	http.HandleFunc("/dup", Dup)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := `<html>
  <body>
    <h4>Choose a component</h4>
    <ul>
      <li><a href="/textfield">Textfield</a></li>
      <li><a href="/button">Button</a></li>
      <li><a href="/checkbox">Checkbox</a></li>
      <li><a href="/checkbox_slice">Checbkox slices</a></li>
      <li><a href="/select">Select</a></li>
      <li><a href="/textarea">Textarea</a></li>
      <li><a href="/input_confirm">Input confirm</a></li>
      <li><a href="/input_date">Input date</a></li>
      <li><a href="/input_file">Input file</a></li>
      <li><a href="/expand">Expand</a></li>
      <li><a href="/dialog">Dialog</a></li>
      <li><a href="/slider">Slider</a></li>
      <li><a href="/data_table">Data table</a></li>
      <li><a href="/tab">Tabs</a></li>
      <li><a href="/cropper">Cropper</a></li>
      <li><a href="/chip_autocomplete">Chip autocomplete</a></li>
			<li><a href="/dup">Dup</a></li>

      <li><a href="/snackbar">TBD Snackbar</a></li>
    </ul>
  </body>
  </html>`

	t := template.New("")
	t.Parse(tmpl)
	t.Execute(w, struct{}{})
}

type viewData struct {
	Title      string
	Components []mgc.ViewComponentRenderer
}

func renderComponents(w http.ResponseWriter, data *viewData) {

	tmpl := `<html>
  <head>
  <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
  <link rel="stylesheet" href="/static/node_modules/@mh-cbon/material-design-lite/build/material.min.css">
  </head>
  <body>
    <h4>{{.Title}}</h4>
    {{range .Components}}
      {{.Render}}
      <hr/>
    {{end}}
    <script src="/static/node_modules/@mh-cbon/material-design-lite/build/tinymce/tinymce.min.js"></script>
    <script src="/static/node_modules/@mh-cbon/material-design-lite/build/material.js"></script>
  </body>
  </html>`

	t := mgc.MustTemplate()
  t.Parse(tmpl)

	ctx := mgc.NewRenderContext(t, w)

	for _, v := range data.Components {
		v.SetRenderContext(ctx)
	}

	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}
