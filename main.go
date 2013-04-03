package main

import (
	"fmt"
	hgConfig "hellogolang/HooGL/config"
	hgFrontHandler "hellogolang/application/handler/front"
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
	"/ajaxlogin/":      hgFrontHandler.AjaxLogin,
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
	startErr := http.ListenAndServe(":"+hgConfig.GetConfig("port"), nil) //设置监听的端口
	if startErr != nil {
		//hgHelper.LogMessage("server start error")
		fmt.Println(startErr)
		fmt.Println("server start error")
		os.Exit(1)
	}

}
