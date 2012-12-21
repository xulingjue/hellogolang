/*
 *文章类
 */
package hgHandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func ArticlePageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "articlePage") //这个写入到w的是输出到客户端的
}

func ArticleItemHandler(w http.ResponseWriter, r *http.Request) {
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

	//t.ExecuteTemplate(w, "header", nil)
	t.ExecuteTemplate(w, "article-list", nil)
	//t.ExecuteTemplate(w, "footer", nil)
	t.Execute(w, nil)

}
