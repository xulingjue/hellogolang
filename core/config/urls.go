package config

import (
	hgArticle "hellogolang/module/article"
	hgPeople "hellogolang/module/people"
	"net/http"
)

var handlers = map[string]func(http.ResponseWriter, *http.Request){
	/*soruce*/
	"/static/": sourceHandler,

	/*article*/
	"/":        hgArticle.Page,
	"/article": hgArticle.Item,

	/*people*/
	"/login":    hgPeople.LoginView,
	"/doLogin":  hgPeople.Login,
	"/regist":   hgPeople.RegistView,
	"/doRegist": hgPeople.Regist,
}

/*静态文件加载函数*/
func sourceHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
