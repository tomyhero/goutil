package goutil_test

import (
	"."
	"html/template"
	"os"
)

func ExamplePager() {

	const tpl = `
{{with .pager }}
{{ if ne .LastPage .FirstPage }}
<div class="container text-center">
{{ if .PreviousPage }}<a href="{{.BuildPath .PreviousPage }}" class="prev">&laquo;&nbsp;Prev</a>{{end}}
<ul class="pagination">
{{ range $i,$p := .PagesInNavigation 5 }}
{{ if eq $p $.pager.CurrentPage }}<li class="active">{{$p}}</li>{{else}}<li><a href="{{$.pager.BuildPath $p}}">{{$p}}</a></li>{{end}}
{{ end }}
</ul>
{{ if .NextPage }}<a href="{{.BuildPath .NextPage }}" class="next">Next&nbsp;&raquo;</a>{{end}}
</div><!-- /pager -->
{{ end }}
{{ end }}
`

	goutil.PAGER_DEFAULT_ENTRY_PER_PAGE = 10
	pager := goutil.NewPager("9")
	pager.TotalEntries = 1000
	pager.RequestURI = "/foo/?bar=bar&p=10"
	t, _ := template.New("paging").Parse(tpl)
	t.Execute(os.Stdout, map[string]interface{}{"pager": pager})

	/*
		    // e.g with squirrel
				pager := utils.NewPager("3")
				b := sq.Select("*").From("foo")
				t := sq.Select("count(*)").From("foo")
		        t.RunWith(db).QueryRow().Scan(&pager.TotalEntries)
				b = b.Limit(pager.EntryPerPage)
				b = b.Offset(pager.Offset())
				rows, _ := b.RunWith(db).Query()

	*/

	// Output:
	// <div class="container text-center">
	// <a href="/foo/?bar=bar&amp;p=8" class="prev">&laquo;&nbsp;Prev</a>
	// <ul class="pagination">
	//
	// <li><a href="/foo/?bar=bar&amp;p=7">7</a></li>
	//
	// <li><a href="/foo/?bar=bar&amp;p=8">8</a></li>
	//
	// <li class="active">9</li>
	//
	// <li><a href="/foo/?bar=bar&amp;p=10">10</a></li>
	//
	// <li><a href="/foo/?bar=bar&amp;p=11">11</a></li>
	//
	// </ul>
	// <a href="/foo/?bar=bar&amp;p=10" class="next">Next&nbsp;&raquo;</a>
	// </div>
}
