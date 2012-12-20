package main

/*
 *服务器log类，日志按天分
 */

import (
	"fmt"
	"log"
	"os"
	"time"
)

func logMessage(message string) {
	//获得当前时间 构建日志路径
	var logFileName = "logs/hg-log-" + time.Now().Format("2006-01-02") + ".txt"

	_, logFileOpenErr := os.Open(logFileName)
	if logFileOpenErr != nil && os.IsNotExist(logFileOpenErr) {
		fmt.Println("file not exist!")
		logFile, logFileOpenErr := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE, 0755)
		if logFileOpenErr != nil {
			fmt.Println("open logfile error")
		}
		log.SetOutput(logFile)
	}

	log.Println(message)
}
