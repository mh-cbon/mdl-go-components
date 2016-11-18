package handlers

import (
	mgc "github.com/mh-cbon/mdl-go-components"
	"github.com/mh-cbon/mdl-go-components/components"
	"net/http"
	"time"
)

func InputDate(w http.ResponseWriter, r *http.Request) {

	data := &viewData{
		Title:      "Input date",
		Components: make([]mgc.ViewComponentRenderer, 0),
	}

	var t1 *components.InputDate
	var line *components.Slice

	// -
	line = components.NewSlice()
	t1 = components.NewInputDate()
	t1.SetName("the_date")
	t1.SetLabel("Choisissez une date")

	t1.SetDisplayLocale("fr")
	t1.SetDisplayMode("24H")
	t1.SetDisplayFormat("ddd DD MMM YYYY")
	t1.SetFormat("YYYY-MM-DDThh:mm:ssZ")
	t1.SetOkText("Confirmer")
	t1.SetCancelText("Annuler")
	t1.SetPresent(time.Now().UTC())
	t1.SetPast(time.Now().Add(-8 * 24 * 365 * time.Hour).UTC())
	t1.SetFuture(time.Now().Add(8 * 24 * 365 * time.Hour).UTC())

	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	line = components.NewSlice()
	t1 = components.NewInputDate()
	t1.SetName("the_date")
	t1.SetLabel("Pick a date")
	t1.SetError("Choose another date")

	t1.SetDisplayLocale("en")
	t1.SetDisplayFormat("ddd DD MMM YYYY")
	t1.SetFormat("YYYY-MM-DDThh:mm:ssZ")
	t1.SetOkText("Confirm")
	t1.SetCancelText("Cancel")
	t1.SetPresent(time.Now().UTC())
	t1.SetPast(time.Now().Add(-8 * 24 * 365 * time.Hour).UTC())
	t1.SetFuture(time.Now().Add(8 * 24 * 365 * time.Hour).UTC())

	line.Add(t1)
	data.Components = append(data.Components, line)

	// -
	renderComponents(w, data)
}
