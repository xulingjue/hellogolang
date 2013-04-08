package front

import (
	hgTemplate "hellogolang/HooGL/template"
	"net/http"
	"text/template"
)

func errorMessage(rw http.ResponseWriter, req *http.Request) {
	tmpl := template.New("error-message")
	tmpl.Funcs(template.FuncMap{"StringEqual": hgTemplate.StringEqual, "Int64Equal": hgTemplate.Int64Equal})

	tmpl.ParseFiles(
		"template/front/header.tmpl",
		"template/front/error-message.tmpl",
		"template/front/footer.tmpl")

	tmplInfo := hgTemplate.TmplInfo{}
	tmplInfo.Js = []string{
		"js/jquery.validate.js"}

	tmpl.ExecuteTemplate(rw, "error-message", map[string]interface{}{"tmplInfo": tmplInfo})
}
