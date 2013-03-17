package front

import (
	//"code.google.com/p/gorilla/sessions"
	"fmt"
	"hellogolang/application/model"
	"hellogolang/system/library"
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
	pageSize := 2
	page, err := strconv.Atoi(req.FormValue("page"))
	if err != nil {
		page = 1
	}

	posts, _ := postModel.FindAll(page, pageSize, map[string]string{})
	count, _ := postModel.FindAllCount(map[string]string{})

	var pageHelper library.Page
	pageHelper.Count = count
	pageHelper.PageSize = pageSize
	pageHelper.PageNum = page
	pageHelper.BaseUrl = "/?page="
	pageHelper.Compute()

	tmpl := template.New("indexView")
	tmpl.Funcs(template.FuncMap{"StringEqual": tmplfunc.StringEqual, "Int64Equal": tmplfunc.Int64Equal, "IntEqual": tmplfunc.IntEqual, "RemoveHtmlTag": tmplfunc.RemoveHtmlTag})
	tmpl.ParseFiles(
		"template/front/header.tmpl",
		"template/front/index.tmpl",
		"template/front/footer.tmpl",
		"template/front/page.tmpl")

	siteInfo.Js = []string{
		"js/front/people/index.js"}
	siteInfo.ExtraJs = []string{
		"http://jzaefferer.github.com/jquery-validation/jquery.validate.js"}
	siteInfo.CurrentNav = "index"

	tmpl.ExecuteTemplate(rw, "index", map[string]interface{}{"people": people, "siteInfo": siteInfo, "posts": posts, "pageHelper": pageHelper})
}

/*
 *	文章分页列表
 */
func PostPage(rw http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	pageSize := 2
	page, err := strconv.Atoi(req.FormValue("page"))
	if err != nil {
		page = 1
	}

	conditions := make(map[string]string)
	var pageHelper library.Page

	postClass := req.FormValue("cat")
	_, err = strconv.Atoi(req.FormValue("cat"))
	if err == nil {
		conditions["post.idpost_class ="] = postClass
		pageHelper.BaseUrl = "/post?cat=" + postClass + "&page="
	} else {
		pageHelper.BaseUrl = "/post?page="
	}

	posts, _ := postModel.FindAll(page, pageSize, conditions)
	count, _ := postModel.FindAllCount(conditions)

	pageHelper.Count = count
	pageHelper.PageSize = pageSize
	pageHelper.PageNum = page
	pageHelper.Compute()

	tmpl := template.New("post-pageView")
	tmpl.Funcs(template.FuncMap{"StringEqual": tmplfunc.StringEqual, "Int64Equal": tmplfunc.Int64Equal, "IntEqual": tmplfunc.IntEqual, "RemoveHtmlTag": tmplfunc.RemoveHtmlTag})
	tmpl.ParseFiles(
		"template/front/header.tmpl",
		"template/front/post-list.tmpl",
		"template/front/footer.tmpl",
		"template/front/page.tmpl")

	siteInfo.Js = []string{
		"kindeditor/kindeditor-min.js",
		"kindeditor/lang/zh_CN.js",
		"js/front/post/post-list.js"}
	siteInfo.CurrentNav = "article"

	tmpl.ExecuteTemplate(rw, "post-list", map[string]interface{}{"siteInfo": siteInfo, "posts": posts, "postClass": postClass, "pageHelper": pageHelper})
	tmpl.Execute(rw, nil)
}

/*
 *	查看单个文章页
 */
func PostItem(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	page, err := strconv.Atoi(req.FormValue("page"))
	if err != nil {
		page = 1
	}

	postIdStr := req.FormValue("postId")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {

	}

	pageSize := 10
	post, _ := postModel.Find(postId)
	comments, _ := commentModel.FindAll(page, pageSize, map[string]string{"comment.idpost=": postIdStr})

	tmpl := template.New("post-itemView")
	tmpl.Funcs(template.FuncMap{"StringEqual": tmplfunc.StringEqual, "Int64Equal": tmplfunc.Int64Equal})
	tmpl.ParseFiles(
		"template/front/header.tmpl",
		"template/front/post-item.tmpl",
		"template/front/footer.tmpl")

	siteInfo.Js = []string{
		"kindeditor/kindeditor-min.js",
		"kindeditor/lang/zh_CN.js",
		"js/front/post/post-item.js"}
	siteInfo.CurrentNav = "article"

	//tmpl.ExecuteTemplate(rw, "post-item", map[string]interface{}{"siteInfo": siteInfo, "post": post, "postReplies": postReplies})
	tmpl.ExecuteTemplate(rw, "post-item", map[string]interface{}{"siteInfo": siteInfo, "post": post, "comments": comments})
	tmpl.Execute(rw, nil)
}

/*
 *	创建文章
 */
func PostCreate(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		postType, err := strconv.Atoi(req.FormValue("postType"))
		if err != nil {

		}
		postClass := postClassModel.FindAll(postType)

		tmpl := template.New("post-createView")
		tmpl.Funcs(template.FuncMap{"StringEqual": tmplfunc.StringEqual, "Int64Equal": tmplfunc.Int64Equal})
		tmpl.ParseFiles(
			"template/front/header.tmpl",
			"template/front/post-create.tmpl",
			"template/front/footer.tmpl")

		siteInfo.Js = []string{
			"kindeditor/kindeditor-all.js",
			"kindeditor/lang/zh_CN.js",
			"js/front/post/post-create.js"}
		siteInfo.CurrentNav = "article"

		tmpl.ExecuteTemplate(rw, "post-create", map[string]interface{}{"siteInfo": siteInfo, "postClass": postClass})
		tmpl.Execute(rw, nil)
	} else if req.Method == "POST" {
		req.ParseForm()
		people := isLogin(req)
		var post model.Post
		var err error

		post.IdpostClass, err = strconv.ParseInt(req.FormValue("post_class"), 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		post.Content = req.FormValue("content")
		post.ReprintFrom = req.FormValue("reprint_from")
		post.ReprintUrl = req.FormValue("reprint_url")
		post.Title = req.FormValue("title")
		post.Idpeople = people.Idpeople

		post.Idpeople = 1
		post.ReadNum = 0
		post.ReplyNum = 0

		postModel.Insert(post)
	}
}

func Test(rw http.ResponseWriter, req *http.Request) {
	postModel.FindAll(1, 1, map[string]string{})
}
