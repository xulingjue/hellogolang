package front

import (
	//"code.google.com/p/gorilla/sessions"
	//"fmt"
	"hellogolang/system/tmplfunc"
	"net/http"
	"strconv"
	"text/template"
)

/*
 *	文章分页列表
 */
func PostPage(rw http.ResponseWriter, req *http.Request) {
	tmpl := template.New("post-pageView")
	tmpl.Funcs(template.FuncMap{"StringEqual": tmplfunc.StringEqual, "Int64Equal": tmplfunc.Int64Equal})
	tmpl.ParseFiles(
		"template/front/header.tmpl",
		"template/front/post-list.tmpl",
		"template/front/footer.tmpl")

	siteInfo.Js = []string{
		"kindeditor/kindeditor-min.js",
		"kindeditor/lang/zh_CN.js",
		"js/front/post/post-list.js"}
	siteInfo.CurrentNav = "article"

	tmpl.ExecuteTemplate(rw, "post-list", map[string]interface{}{"siteInfo": siteInfo})
	tmpl.Execute(rw, nil)
}

/*
 *	查看单个文章页
 */
func PostItem(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	pageSize, err := strconv.Atoi(req.FormValue("pageSize"))
	page, err := strconv.Atoi(req.FormValue("page"))
	postId, err := strconv.ParseInt(req.FormValue("postId"), 10, 64)

	if err != nil {

	}

	post, _ := postModel.Find(postId)
	postReplies, _ := postModel.FindAllReply(postId, page, pageSize)

	tmpl := template.New("post-itemView")
	tmpl.Funcs(template.FuncMap{"StringEqual": tmplfunc.StringEqual, "Int64Equal": tmplfunc.Int64Equal})
	tmpl.ParseFiles(
		"template/front/header.tmpl",
		"template/front/post-item.tmpl",
		"template/front/footer.tmpl")

	siteInfo.Js = []string{
		"kindeditor/kindeditor-min.js",
		"kindeditor/lang/zh_CN.js",
		"js/front/post/post-list.js"}
	siteInfo.CurrentNav = "article"

	tmpl.ExecuteTemplate(rw, "post-item", map[string]interface{}{"siteInfo": siteInfo, "post": post, "postReplies": postReplies})
	tmpl.Execute(rw, nil)
}

/*
 *	创建文章
 */
func PostCreate(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		postTypes := postTypeModel.FindAll()

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

		tmpl.ExecuteTemplate(rw, "post-create", map[string]interface{}{"siteInfo": siteInfo, "postTypes": postTypes})
		tmpl.Execute(rw, nil)
	} else if req.Method == "POST" {

	}
}
