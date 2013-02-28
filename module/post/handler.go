package post

import (
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

	tmpl.ExecuteTemplate(rw, "index", nil)
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
