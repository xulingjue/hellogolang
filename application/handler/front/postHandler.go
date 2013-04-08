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

	pageHelper := hgPageination.Page{}
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

	tmplInfo := hgTemplate.TmplInfo{}

	tmplInfo.Js = []string{
		"js/jquery.validate.js"}

	tmplInfo.CurrentNav = "index"
	tmplInfo.Title = "Hello Golang -首页"
	tmplInfo.Description = "开源Go语言爱好者交流平台"

	tmplInfo.AddData("people", people)
	tmplInfo.AddData("posts", posts)
	tmplInfo.AddData("pageHelper", pageHelper)
	tmplInfo.AddData("postClasses", postClasses)

	tmpl.ExecuteTemplate(rw, "index", map[string]interface{}{"tmplInfo": tmplInfo})
}

/*
 *	文章分页列表
 */
func PostPage(rw http.ResponseWriter, req *http.Request) {
	postClasses := postClassModel.FindAll()
	people := isLogin(req)

	req.ParseForm()
	pageSize := 10
	page := hgForm.GetInt(req, "page", 1)

	conditions := make(map[string]string)
	pageHelper := hgPageination.Page{}

	//IdpostClass, err := strconv.ParseInt(req.FormValue("cat"), 10, 64)
	IdpostClass := hgForm.GetInt64(req, "cat", 0)

	if IdpostClass != 0 {
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

	tmplInfo := hgTemplate.TmplInfo{}
	tmplInfo.Js = []string{
		"kindeditor/kindeditor-min.js",
		"kindeditor/lang/zh_CN.js"}
	tmplInfo.CurrentNav = "article"
	tmplInfo.Title = "Hello Golang -文章"
	tmplInfo.Description = "全新的go语言资讯！精选的go语言教程！经典的go语言代码！"

	tmplInfo.AddData("people", people)
	tmplInfo.AddData("posts", posts)
	tmplInfo.AddData("pageHelper", pageHelper)
	tmplInfo.AddData("postClasses", postClasses)
	tmplInfo.AddData("postClass", postClass)

	tmpl.ExecuteTemplate(rw, "post-list", map[string]interface{}{"tmplInfo": tmplInfo})
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

	postId := hgForm.GetInt64(req, "postId", 0)
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

	pageHelper := hgPageination.Page{}

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

	tmplInfo := hgTemplate.TmplInfo{}
	tmplInfo.Js = []string{
		"js/jquery.validate.js"}

	tmplInfo.CurrentNav = "article"

	tmplInfo.Title = "Hello Golang -" + post.Title
	tmplInfo.Description = post.Title

	tmplInfo.AddData("people", people)
	tmplInfo.AddData("post", post)
	tmplInfo.AddData("pageHelper", pageHelper)
	tmplInfo.AddData("postClasses", postClasses)
	tmplInfo.AddData("comments", comments)

	tmpl.ExecuteTemplate(rw, "post-item", map[string]interface{}{"tmplInfo": tmplInfo})
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

		tmplInfo := hgTemplate.TmplInfo{}

		tmplInfo.Js = []string{
			"ckeditor/ckeditor.js"}
		tmplInfo.CurrentNav = "none"
		tmplInfo.Title = "Hello Golang -新建文章"
		tmplInfo.Description = "开源Go语言爱好者交流平台"

		tmplInfo.AddData("people", people)
		tmplInfo.AddData("postClass", postClass)
		tmplInfo.AddData("postClasses", postClasses)

		tmpl.ExecuteTemplate(rw, "post-create", map[string]interface{}{"tmplInfo": tmplInfo})
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
	postId := hgForm.GetInt64(req, "postId", 0)
	people := isLogin(req)

	content := req.FormValue("content")
	post := postModel.Find(postId)

	if people == nil || post == nil {
		http.Redirect(rw, req, "/post/item/?postId="+req.FormValue("postId"), http.StatusFound)
		return
	}

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
