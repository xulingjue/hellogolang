package main

import (
	"encoding/json"
	hgHelpers "hellogolang/hghelpers"
	"net/http"
	"os"
)

/*
	全局变量
*/
var (
	config map[string]string
)

func init() {
	configFile, configFileErr := os.Open("config.json")
	if configFileErr != nil {
		hgHelpers.LogMessage("read config.json error")
		panic(configFileErr)
		os.Exit(1)
	}
	defer configFile.Close()

	configFileDec := json.NewDecoder(configFile)
	configFileErr = configFileDec.Decode(&config)
	if configFileErr != nil {
		hgHelpers.LogMessage("config.json decode error")
		panic(configFileErr)
		os.Exit(1)
	}
}

func main() {
	//初始化数据库

	//初始化URL
	for url, handler := range handlers {
		http.HandleFunc(url, handler)
	}

	//启动服务器
	startErr := http.ListenAndServe(":"+config["port"], nil) //设置监听的端口
	if startErr != nil {
		hgHelpers.LogMessage("server start error")
	}

	hgHelpers.LogMessage("server start success!")
}
