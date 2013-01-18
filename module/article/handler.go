package article

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

/*
	文章分页列表
*/
func Page(rw http.ResponseWriter, req *http.Request) {
	findArticleList()

	fmt.Println("path", req.URL.Path)

	req.ParseForm()
	var articleId = req.FormValue("articleid")
	if strings.EqualFold(articleId, "") {
		//fmt.Fprintf(w, "error")
	}

	//查询文章列表
	t, _ := template.ParseFiles("template/front/header.tmpl",
		"template/front/right-sidebar-article.tmpl",
		"template/front/right-sidebar-topic.tmpl",
		"template/front/article-list.tmpl",
		"template/front/footer.tmpl")

	t.ExecuteTemplate(rw, "article-list", nil)
	t.Execute(rw, nil)
}

/*
	查看单个文章页
*/

func Item(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var articleId = req.FormValue("articleid")
	if strings.EqualFold(articleId, "") {
	}

	t, _ := template.ParseFiles("static/index.html")

	//t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(rw, "index", nil)
	//t.ExecuteTemplate(w, "footer", nil)
	t.Execute(rw, nil)
}
