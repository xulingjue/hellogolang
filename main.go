package main

import (
	hgPeople "hellogolang/module/people"
	hgPost "hellogolang/module/post"
	hgHelper "hellogolang/system/helper"
	"net/http"
	"os"
)

var handlers = map[string]func(http.ResponseWriter, *http.Request){
	/*post*/
	"/":     hgPost.Index,
	"/post": hgPost.Item,

	/*people*/
	"/login":          hgPeople.Login,
	"/regist":         hgPeople.Regist,
	"/people/isexist": hgPeople.AjaxIsExist,
}

func main() {
	http.Handle("/assets/", http.FileServer(http.Dir("static")))

	//初始化URL
	for url, handler := range handlers {
		http.HandleFunc(url, handler)
	}

	//启动服务器
	startErr := http.ListenAndServe(":"+hgHelper.GetConfig("port"), nil) //设置监听的端口
	if startErr != nil {
		hgHelper.LogMessage("server start error")
		os.Exit(1)
	}
}
