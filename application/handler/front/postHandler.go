package front

import (
	//"code.google.com/p/gorilla/sessions"
	"fmt"
	hgForm "hellogolang/HooGL/form"
	hgPageination "hellogolang/HooGL/pageination"
	hgTemplate "hellogolang/HooGL/template"
	"hellogolang/application/model"
	"net/http"
	"strconv"
	"text/template"
)

/*
 *首页
 */

func Index(rw http.ResponseWriter, req *http.Request) {
	postClasses := postClassModel.FindAll()
	people := isLogin(req)
	req.ParseForm()

	pageSize := 10
	page := hgForm.GetInt(req, "page", 1)

	posts, count := postModel.FindAll(page, pageSize, map[string]string{})

	pageHelper := hgPageination.Page
	pageHelper.Count = count
	pageHelper.PageSize = pageSize
	pageHelper.PageNum = page
	pageHelper.BaseUrl = "/?page="
	pageHelper.Compute()

	tmpl := template.New("indexView")
	tmpl.Funcs(template.FuncMap{"StringEqual": hgTemplate.StringEqual, "Int64Equal": hgTemplate.Int64Equal, "IntEqual": hgTemplate.IntEqual, "RemoveHtmlTag": hgTemplate.RemoveHtmlTag})
	tmpl.ParseFiles(
		"template/front/header.tmpl",
		"template/front/index.tmpl",
		"template/front/footer.tmpl",
		"template/front/page.tmpl",
		"template/front/sidebar.tmpl")

	siteInfo.ExtraJs = []string{
		"http://jzaefferer.github.com/jquery-validation/jquery.validate.js"}
	siteInfo.CurrentNav = "index"

	tmpl.ExecuteTemplate(rw, "index", map[string]interface{}{"people": people, "siteInfo": siteInfo, "posts": posts, "pageHelper": pageHelper, "postClasses": postClasses})
}

/*
 *	文章分页列表
 */
func PostPage(rw http.ResponseWriter, req *http.Request) {
	postClasses := postClassModel.FindAll()

	people := isLogin(req)
	siteInfo.BackUrl = req.RequestURI

	req.ParseForm()
	pageSize := 10
	page := hgForm.GetInt(req, "page", 1)

	conditions := make(map[string]string)
	pageHelper := hgPageination.Page

	//IdpostClass, err := strconv.ParseInt(req.FormValue("cat"), 10, 64)
	IdpostClass := hgForm.GetInt64(req, "cat", 1)

	if err == nil {
		conditions["post.idpost_class ="] = req.FormValue("cat")
		pageHelper.BaseUrl = "/post/?cat=" + req.FormValue("cat") + "&page="
	} else {
		pageHelper.BaseUrl = "/post/?page="
	}

	postClass := postClassModel.Find(IdpostClass)

	if postClass == nil {
		//出错处理
	}

	posts, count := postModel.FindAll(page, pageSize, conditions)

	pageHelper.Count = count
	pageHelper.PageSize = pageSize
	pageHelper.PageNum = page
	pageHelper.Compute()

	tmpl := template.New("post-pageView")
	tmpl.Funcs(template.FuncMap{"StringEqual": hgTemplate.StringEqual, "Int64Equal": hgTemplate.Int64Equal, "IntEqual": hgTemplate.IntEqual, "RemoveHtmlTag": hgTemplate.RemoveHtmlTag})
	tmpl.ParseFiles(
		"template/front/header.tmpl",
		"template/front/post-list.tmpl",
		"template/front/footer.tmpl",
		"template/front/page.tmpl",
		"template/front/sidebar.tmpl")

	siteInfo.Js = []string{
		"kindeditor/kindeditor-min.js",
		"kindeditor/lang/zh_CN.js",
		"js/front/post/post-list.js"}

	siteInfo.CurrentNav = "article"

	tmpl.ExecuteTemplate(rw, "post-list", map[string]interface{}{"people": people, "siteInfo": siteInfo, "posts": posts, "postClass": postClass, "pageHelper": pageHelper, "postClasses": postClasses})
	tmpl.Execute(rw, nil)

}

/*
 *	查看单个文章页
 */
func PostItem(rw http.ResponseWriter, req *http.Request) {
	postClasses := postClassModel.FindAll()

	req.ParseForm()
	people := isLogin(req)

	page := hgForm.GetInt(req, "page", 1)

	postId := hgForm.GetInt64(req, "page", 0)
	if postId == 0 {
		//出错
	}

	pageSize := 10
	post := postModel.Find(postId)

	if post == nil {
		fmt.Println("post is nil...")
		//文章不存在
	}

	comments, count := commentModel.FindAllByPostID(postId, page, pageSize)

	pageHelper := hgPageination.Page

	pageHelper.BaseUrl = "/post/item/?postId=" + strconv.FormatInt(postId, 10) + "&page="
	pageHelper.Count = count
	pageHelper.PageSize = pageSize
	pageHelper.PageNum = page
	pageHelper.Compute()

	tmpl := template.New("post-itemView")
	tmpl.Funcs(template.FuncMap{"StringEqual": hgTemplate.StringEqual, "Int64Equal": hgTemplate.Int64Equal, "IntEqual": hgTemplate.IntEqual})
	tmpl.ParseFiles(
		"template/front/header.tmpl",
		"template/front/post-item.tmpl",
		"template/front/footer.tmpl",
		"template/front/page.tmpl",
		"template/front/sidebar.tmpl")

	siteInfo.Js = []string{
		"js/front/post/post-item.js"}
	siteInfo.ExtraJs = []string{
		"http://jzaefferer.github.com/jquery-validation/jquery.validate.js"}

	siteInfo.CurrentNav = "article"

	tmpl.ExecuteTemplate(rw, "post-item", map[string]interface{}{"people": people, "siteInfo": siteInfo, "post": post, "comments": comments, "pageHelper": pageHelper, "postClasses": postClasses})
	tmpl.Execute(rw, nil)

	//更新阅读数
	postModel.UpdateReadNum(*post)
}

/*
 *	创建文章
 */
func PostCreate(rw http.ResponseWriter, req *http.Request) {
	postClasses := postClassModel.FindAll()

	people := isLogin(req)
	if req.Method == "GET" {
		postClass := postClassModel.FindAll()
		tmpl := template.New("post-createView")
		tmpl.Funcs(template.FuncMap{"StringEqual": hgTemplate.StringEqual, "Int64Equal": hgTemplate.Int64Equal})
		tmpl.ParseFiles(
			"template/front/header.tmpl",
			"template/front/post-create.tmpl",
			"template/front/footer.tmpl",
			"template/front/sidebar.tmpl")

		siteInfo.Js = []string{
			"ckeditor/ckeditor.js",
			"js/front/post/post-create.js"}
		siteInfo.CurrentNav = "none"

		tmpl.ExecuteTemplate(rw, "post-create", map[string]interface{}{"siteInfo": siteInfo, "postClass": postClass, "people": people, "postClasses": postClasses})
		tmpl.Execute(rw, nil)
	} else if req.Method == "POST" {
		req.ParseForm()
		people := isLogin(req)

		var err error
		var post model.Post

		post.Class.IdpostClass, err = strconv.ParseInt(req.FormValue("post_class"), 10, 64)

		post.Content = req.FormValue("content")
		post.ReprintFrom = req.FormValue("reprint_from")
		post.ReprintUrl = req.FormValue("reprint_url")
		post.Title = req.FormValue("title")
		post.Author.Idpeople = people.Idpeople

		post.ReadNum = 0
		post.ReplyNum = 0

		postModel.Insert(post)
		http.Redirect(rw, req, "/post/?cat="+req.FormValue("post_class"), http.StatusFound)

		if err != nil {
			fmt.Println(err)
		} else {
			people = peopleModel.Find(people.Idpeople)
			people.Postnum++
			peopleModel.Update(*people)
		}

	}
}

func CommentCreate(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("path", req.URL.Path)

	req.ParseForm()
	postId, _ := strconv.ParseInt(req.FormValue("postId"), 10, 64)
	people := isLogin(req)
	content := req.FormValue("content")
	post := postModel.Find(postId)

	var comment model.Comment
	comment.Idpost = postId
	comment.Content = content
	comment.Author.Idpeople = people.Idpeople
	comment.Parent = 0

	commentModel.Insert(comment)

	http.Redirect(rw, req, "/post/item/?postId="+req.FormValue("postId"), http.StatusFound)

	//更新回复数
	postModel.UpdateReplyNum(*post)
}
