package main

import (
	hgHelper "hellogolang/core/helper"
	hgArticle "hellogolang/module/article"
	hgPeople "hellogolang/module/people"
	"net/http"
	"os"
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

func main() {
	//初始化URL
	for url, handler := range handlers {
		http.HandleFunc(url, handler)
	}

	//启动服务器
	startErr := http.ListenAndServe(":"+hgHelper.ConfigValue("port"), nil) //设置监听的端口
	if startErr != nil {
		hgHelper.LogMessage("server start error")
		os.Exit(1)
	}
}
