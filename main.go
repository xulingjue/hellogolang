package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var config map[string]string

func init() {
	configFile, configFileErr := os.Open("config.json")
	if configFileErr != nil {
		logMessage("read config.json error")
		panic(configFileErr)
		os.Exit(1)
	}
	defer configFile.Close()

	configFileDec := json.NewDecoder(configFile)
	configFileErr = configFileDec.Decode(&config)
	if configFileErr != nil {
		logMessage("config.json decode error")
		panic(configFileErr)
		os.Exit(1)
	}
}

func main() {
	for url, handler := range handlers {
		http.HandleFunc(url, handler)
	}

	err := http.ListenAndServe(":"+config["port"], nil) //设置监听的端口
	if err != nil {
		logMessage("server start error")
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello-golang")
}
