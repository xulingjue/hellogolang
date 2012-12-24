/*
 *文章类
 */
package hghandlers

import (
	//"fmt"
	"net/http"
	"strings"
	"text/template"
)

func ArticlePageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var articleId = r.FormValue("articleid")
	if strings.EqualFold(articleId, "") {
		//fmt.Fprintf(w, "error")
	}

	t, _ := template.ParseFiles("template/header.tmpl",
		"template/right-sidebar-article.tmpl",
		"template/right-sidebar-topic.tmpl",
		"template/article-list.tmpl",
		"template/footer.tmpl")

	t.ExecuteTemplate(w, "article-list", nil)
	t.Execute(w, nil)
}

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
