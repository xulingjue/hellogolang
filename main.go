package main

import (
	hgConfig "hellogolang/core/config"
	hgHelper "hellogolang/core/helper"
	"net/http"
	"os"
)

func main() {

	//启动服务器
	startErr := http.ListenAndServe(":"+hgConfig.GetValue("port"), nil) //设置监听的端口
	if startErr != nil {
		hgHelper.LogMessage("server start error")
		os.Exit(1)
	}
}
