package helper

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	config map[string]string
)

func init() {
	//读取配置文件
	fmt.Println("read config file")

	configFile, configFileErr := os.Open("config.json")
	if configFileErr != nil {
		fmt.Println("read config.json error")
		os.Exit(1)
	}
	defer configFile.Close()
	//解析配置文件
	configFileDec := json.NewDecoder(configFile)
	configFileErr = configFileDec.Decode(&config)
	if configFileErr != nil {
		fmt.Println("config.json decode error")
		os.Exit(1)
	}
}

func GetConfig(key string) string {
	return config[key]
}
