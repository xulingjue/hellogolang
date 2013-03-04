package post

import (
	hgHelper "hellogolang/system/helper"
	"net/http"
	"text/template"
)

/*
 *首页
 */
func Index(rw http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.New("registView").ParseFiles(
		"template/front/header.tmpl",
		"template/front/index.tmpl",
		"template/front/footer.tmpl")

	js := []string{
		"front/people/index.js"}
	extra_js := []string{
		"http://jzaefferer.github.com/jquery-validation/jquery.validate.js"}

	tmpl.ExecuteTemplate(rw, "index", map[string]interface{}{"baseUrl": hgHelper.GetConfig("base_url"), "js": js, "extra_js": extra_js})
}

/*
 *	文章分页列表
 */
func Page(rw http.ResponseWriter, req *http.Request) {

}

/*
 *	查看单个文章页
 */
func Item(rw http.ResponseWriter, req *http.Request) {

}

/*
 *	创建文章
 */
func Create(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		tmpl, _ := template.ParseFiles("template/front/header.tmpl",
			"template/front/post-create.tmpl",
			"template/front/footer.tmpl")

		tmpl.ExecuteTemplate(rw, "post-create", nil)
		tmpl.Execute(rw, nil)
	} else if req.Method == "POST" {

	}
}
