package components_test

import (
	"github.com/mh-cbon/mdl-go-components/components"
	"net/url"
	"testing"
)

func TestPagination(t *testing.T) {

	paginer := components.NewPagination()
  u, _ := url.Parse("http://what.ever/some")
	paginer.SetBaseUrl(u)
	paginer.SetPageParamName("Page")
	paginer.SetTotalItemCount(10, 100)

	expectations := []string{
		`<span\s+disabled="disabled"\s+href="#"\s+id="rnd-1"\s+class="mdl-button mdl-js-button mdl-button--icon"\s+><i class='material-icons'>keyboard_arrow_left</i></span>`,
		`<option\s+value="http://what.ever/some\?Page=1"\s+selected="selected"\s+>1</option>`,
		`<option\s+value="http://what.ever/some\?Page=10"\s+>10</option>`,
		`\s+/\s+10`,
		`<a\s+href="#"\s+id="rnd-3"\s+class="mdl-button mdl-js-button mdl-button--icon"\s+><i class='material-icons'>keyboard_arrow_right</i></a>`,
	}
	validateComponent(t, paginer, expectations)
}
