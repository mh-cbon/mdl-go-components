package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"github.com/vincent-petithory/dataurl"
	"image/png"
	"io/ioutil"
	"net/http"
)

func Cropper(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Cropper",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var line *components.Slice
	var t1 *components.Cropper
	var form *components.Form

	// -
	line = components.NewSlice()

	t1 = components.NewCropper()
	t1.SetCurrentImg("/static/avatar-1.png")
	t1.SetName("result")
	t1.SetLabel("base64 result")

	form = components.NewForm()
	form.Add(t1)

	form.SetMethod("POST")
	form.SetAction("/cropper_post")
	form.AddHidden("result_type").SetValue("base64")
	form.AddSubmit("vv").SetLabel("Submit")

	line.Add(form)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()

	t1 = components.NewCropper()
	t1.SetCurrentImg("/static/avatar-1.png")
	t1.SetName("result")
	t1.SetResultMode("geometry")
	t1.SetLabel("geometry result")

	form = components.NewForm()
	form.Add(t1)

	form.SetMethod("POST")
	form.SetAction("/cropper_post")
	form.AddHidden("result_type").SetValue("geometry")
	form.AddSubmit("vv").SetLabel("Submit")

	line.Add(form)
	data.Components = append(data.Components, line)

	//- add an example to show how to manage the response on the server side

	// -
	renderComponents(w, data)
}

func CropperPost(w http.ResponseWriter, r *http.Request) {
	// we will just print out the interesting data
	// and demo the interesting code
	// then return to /cropper

	r.ParseForm()
	result := r.Form.Get("result")
	result_type := r.Form.Get("result_type")
	fmt.Println(result_type)

	if result_type == "base64" {
		fmt.Println(result[0:50])
		dataURL, err := dataurl.DecodeString(result)
		if err != nil {
			panic(err)
		}
		fmt.Println(dataURL.ContentType())
		var b bytes.Buffer
		b.Write(dataURL.Data)
		config, err := png.DecodeConfig(&b)
		if err != nil {
			panic(err)
		}
		fmt.Println(config)
		b.Truncate(0)
		b.Write(dataURL.Data)
		img, err := png.Decode(&b)
		if err != nil {
			panic(err)
		}
		// write the file
		png.Encode(ioutil.Discard, img)

	} else if result_type == "geometry" {
		fmt.Println(result)
		geo := &components.GeometryResult{}
		if err := json.Unmarshal([]byte(result), geo); err != nil {
			panic(err)
		}
		fmt.Println(geo)
		// rest is up to you !
	}
	http.Redirect(w, r, "/cropper", http.StatusSeeOther)
}
