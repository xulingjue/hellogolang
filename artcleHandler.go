package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

/*
	文章分页列表
*/
func articlePageHandler(w http.ResponseWriter, r *http.Request) {
	//findArticleList()
	fmt.Println("test")
	r.ParseForm()
	var articleId = r.FormValue("articleid")
	if strings.EqualFold(articleId, "") {
		//fmt.Fprintf(w, "error")
	}

	//查询文章列表

	t, _ := template.ParseFiles("template/header.tmpl",
		"template/right-sidebar-article.tmpl",
		"template/right-sidebar-topic.tmpl",
		"template/article-list.tmpl",
		"template/footer.tmpl")

	t.ExecuteTemplate(w, "article-list", nil)
	t.Execute(w, nil)
}

/*
	查看单个文章页
*/

func articleItemHandler(w http.ResponseWriter, r *http.Request) {
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
