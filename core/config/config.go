package config

import (
	"encoding/json"
	"net/http"
	"os"
)

var (
	config map[string]string
)

func init() {
	//读取配置文件
	configFile, configFileErr := os.Open("config.json")
	if configFileErr != nil {
		//logMessage("read config.json error")
		os.Exit(1)
	}
	defer configFile.Close()
	//解析配置文件
	configFileDec := json.NewDecoder(configFile)
	configFileErr = configFileDec.Decode(&config)
	if configFileErr != nil {
		//logMessage("config.json decode error")
		os.Exit(1)
	}
	//初始化URL
	for url, handler := range handlers {
		http.HandleFunc(url, handler)
	}
}

func GetValue(key string) string {
	return config[key]
}
