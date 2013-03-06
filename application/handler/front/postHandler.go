package front

import (
	//"code.google.com/p/gorilla/sessions"
	"fmt"
	"hellogolang/system/tmplfunc"
	"net/http"
	"text/template"
)

/*
 *	文章分页列表
 */
func PostPage(rw http.ResponseWriter, req *http.Request) {

}

/*
 *	查看单个文章页
 */
func PostItem(rw http.ResponseWriter, req *http.Request) {

}

/*
 *	创建文章
 */
func PostCreate(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		fmt.Println("post create")

		tmpl := template.New("post-createView")
		tmpl.Funcs(template.FuncMap{"StringEqual": tmplfunc.StringEqual, "Int64Equal": tmplfunc.Int64Equal})
		tmpl.ParseFiles(
			"template/front/header.tmpl",
			"template/front/post-create.tmpl",
			"template/front/footer.tmpl")

		siteInfo.Js = []string{
			"kindeditor/kindeditor-min.js",
			"kindeditor/lang/zh_CN.js",
			"js/front/post/post-create.js"}
		siteInfo.CurrentNav = "article"

		tmpl.ExecuteTemplate(rw, "post-create", map[string]interface{}{"siteInfo": siteInfo})
		tmpl.Execute(rw, nil)
	} else if req.Method == "POST" {

	}
}
