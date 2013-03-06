package front

import (
	"hellogolang/system/tmplfunc"
	"net/http"
	"text/template"
)

/*
 *首页
 */

func Index(rw http.ResponseWriter, req *http.Request) {
	people := isLogin(req)

	tmpl := template.New("registView")
	tmpl.Funcs(template.FuncMap{"StringEqual": tmplfunc.StringEqual, "Int64Equal": tmplfunc.Int64Equal})
	tmpl.ParseFiles(
		"template/front/header.tmpl",
		"template/front/index.tmpl",
		"template/front/footer.tmpl")

	siteInfo.Js = []string{
		"js/front/people/index.js"}
	siteInfo.ExtraJs = []string{
		"http://jzaefferer.github.com/jquery-validation/jquery.validate.js"}
	siteInfo.CurrentNav = "index"

	tmpl.ExecuteTemplate(rw, "index", map[string]interface{}{"people": people, "siteInfo": siteInfo})
}
