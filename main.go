package main

import (
	"fmt"
	hgFrontHandler "hellogolang/application/handler/front"
	hgHelper "hellogolang/system/helper"
	"net/http"
	"os"
)

var handlers = map[string]func(http.ResponseWriter, *http.Request){
	"/": hgFrontHandler.Index,
	/*post*/
	"/post/":         hgFrontHandler.PostPage,
	"/post/create/":  hgFrontHandler.PostCreate,
	"/post/item/":    hgFrontHandler.PostItem,
	"/post/comment/": hgFrontHandler.CommentCreate,

	/*people*/
	"/login/":          hgFrontHandler.Login,
	"/logout/":         hgFrontHandler.Logout,
	"/regist/":         hgFrontHandler.Regist,
	"/people/isexist/": hgFrontHandler.PeopleAjaxIsExist,
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
		//hgHelper.LogMessage("server start error")
		fmt.Println("server start error")
		os.Exit(1)
	}

}
