package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"testing"
)

func TestDup(t *testing.T) {

	dup := components.NewDup()
	dup.SetId("id")
	dup.SetBtAddText("more")
	dup.SetBtRemoveText("less")

	bt := components.NewButton()
	bt.SetLabel("hello bt!")
	dup.Add(bt)
	bt = components.NewButton()
	bt.SetLabel("hello bt2!")
	dup.Add(bt)
	bt = components.NewButton()
	bt.SetLabel("hello bt3!")
	dup.Add(bt)

	bt = components.NewButton()
	bt.SetLabel("hello duped bt!")
	dup.SetDup(bt)

	expectations := []string{
		`<div\s+id="id"\s+class="custom-dup custom-js-dup\s+"\s+>`,
		`<div class="custom-dup-template">\s+<div class="custom-dup-item">\s+<div class="custom-dup-component">\s+<button[^>]+id="id-0-[$]incrIndex[$]"`,
		`<div class="custom-dup-container">\s+<div class="custom-dup-item">\s+<div class="custom-dup-component">\s+<button[^>]+id="rnd-1"`,
		`<button type="button"\s+id="id-0-[$]incrIndex[$]"[^>]+>hello duped bt!</button>`,
		`<button[^>]+>\s+less\s+</button>`,
		`<div class="custom-dup-component">\s+<button[^>]+id="rnd-1"[^>]+>hello bt!</button>\s+</div>\s+<button[^>]+>\s+less\s+</button>`,
		`<div class="custom-dup-component">\s+<button[^>]+id="rnd-2"[^>]+>hello bt2!</button>\s+</div>\s+<button[^>]+>\s+less\s+</button>`,
		`<button[^>]+id="rnd-3"[^>]+>hello bt3!</button>`,
	}
	validateComponent(t, dup, expectations)
}

func TestDupSliceDuped(t *testing.T) {

	dup := components.NewDup()
	dup.SetId("id")
	dup.SetBtAddText("more")
	dup.SetBtRemoveText("less")

	bt := components.NewButton()
	bt.SetLabel("hello bt!")
	dup.Add(bt)
	bt = components.NewButton()
	bt.SetLabel("hello bt2!")
	dup.Add(bt)
	bt = components.NewButton()
	bt.SetLabel("hello bt3!")
	dup.Add(bt)

	sl := components.NewSlice()
	txt := components.NewInputText()
	txt.SetLabel("hello duped bt!")
	sl.Add(txt)
	txt = components.NewInputText()
	txt.SetLabel("hello duped bt!2")
	sl.Add(txt)
	dup.SetDup(sl)

	expectations := []string{
		`<input\s+id="id-0-[$]incrIndex[$]"\s+class="mdl-textfield__input "[^>]+>`,
		`<label class="mdl-textfield__label" for="id-0-[$]incrIndex[$]">hello duped bt!</label>`,
		`<input\s+id="id-1-[$]incrIndex[$]"\s+class="mdl-textfield__input "[^>]+>`,
		`<label class="mdl-textfield__label" for="id-1-[$]incrIndex[$]">hello duped bt!2</label>`,
	}
	validateComponent(t, dup, expectations)
}
