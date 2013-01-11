package config

import (
	hgArticle "hellogolang/module/article"
	"net/http"
)

var handlers = map[string]func(http.ResponseWriter, *http.Request){
	/*soruce*/
	"/static/": sourceHandler,

	/*article*/
	"/":        hgArticle.ArticlePageHandler,
	"/article": hgArticle.ArticleItemHandler,
}

/*静态文件存储位置*/
func sourceHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
