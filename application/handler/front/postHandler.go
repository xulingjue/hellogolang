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
 *首页
 */

func Index(rw http.ResponseWriter, req *http.Request) {
	people := isLogin(req)

	req.ParseForm()
	pageSize := 20
	page, err := strconv.Atoi(req.FormValue("page"))
	if err != nil {
		page = 1
	}

	posts, _ := postModel.FindAll(page, pageSize, map[string]string{"post.parentid =": "'0'"})

	tmpl := template.New("indexView")
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

	tmpl.ExecuteTemplate(rw, "index", map[string]interface{}{"people": people, "siteInfo": siteInfo, "posts": posts})
}

/*
 *	文章分页列表
 */
func PostPage(rw http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	pageSize := 20
	page, err := strconv.Atoi(req.FormValue("page"))
	if err != nil {
		page = 1
	}

	postType := req.FormValue("postType")

	posts, _ := postModel.FindAll(page, pageSize, map[string]string{"post.parentid =": "'0'", "post_type.idpost_type": postType})

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

	tmpl.ExecuteTemplate(rw, "post-list", map[string]interface{}{"siteInfo": siteInfo, "posts": posts})
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

func Test(rw http.ResponseWriter, req *http.Request) {
	postModel.FindAll(1, 1, map[string]string{})
}
