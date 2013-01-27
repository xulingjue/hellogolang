package helper

import (
	"encoding/json"
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
}

func ConfigValue(key string) string {
	return config[key]
}

func Test() string {
	return "hello world!"
}
