package main

import (
	_ "code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

/*
	全局变量
*/
var (
	config map[string]string
	hgDb   *sql.DB
)

func init() {
	//读取配置文件
	configFile, configFileErr := os.Open("config.json")
	if configFileErr != nil {
		logMessage("read config.json error")
		os.Exit(1)
	}
	defer configFile.Close()
	//解析配置文件
	configFileDec := json.NewDecoder(configFile)
	configFileErr = configFileDec.Decode(&config)
	if configFileErr != nil {
		logMessage("config.json decode error")
		os.Exit(1)
	}

	//初始化数据库
	hgDb, dbErr := sql.Open("mysql", "root:adzure1105@tcp(192.168.1.151:3306)/hellogolang?charset=utf8")
	if dbErr != nil {
		logMessage("db connect error")
		os.Exit(1)
	}
	//检查数据库连接
	_, dbErr = hgDb.Query("select 1")
	if dbErr != nil {
		logMessage("db connect error")
		os.Exit(1)
	}

	//初始化URL
	for url, handler := range handlers {
		fmt.Println(url)
		http.HandleFunc(url, handler)
	}
}

func main() {
	//启动服务器
	startErr := http.ListenAndServe(":"+config["port"], nil) //设置监听的端口
	if startErr != nil {
		logMessage("server start error")
		os.Exit(1)
	}

	logMessage("server start success!")
}
