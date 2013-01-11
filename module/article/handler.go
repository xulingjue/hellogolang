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
func ArticlePageHandler(w http.ResponseWriter, r *http.Request) {
	findArticleList()

	fmt.Println("path", r.URL.Path)

	r.ParseForm()
	var articleId = r.FormValue("articleid")
	if strings.EqualFold(articleId, "") {
		//fmt.Fprintf(w, "error")
	}

	//查询文章列表

	t, _ := template.ParseFiles("template/front/header.tmpl",
		"template/front/right-sidebar-article.tmpl",
		"template/front/right-sidebar-topic.tmpl",
		"template/front/article-list.tmpl",
		"template/front/footer.tmpl")

	t.ExecuteTemplate(w, "article-list", nil)
	t.Execute(w, nil)
}

/*
	查看单个文章页
*/

func ArticleItemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var articleId = r.FormValue("articleid")
	if strings.EqualFold(articleId, "") {
	}

	t, _ := template.ParseFiles("static/index.html")

	//t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "index", nil)
	//t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, nil)
}
