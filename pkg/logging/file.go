package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	// LogSavePath 日志保存路径
	LogSavePath = "runtime/logs/"
	// LogSaveName 日志保存名称
	LogSaveName = "log"
	// LogFileExt 日志保存后缀
	LogFileExt = "log"
	// TimeFormat 时间
	TimeFormat = "20060102"
)

// getLogFilePath 获取日志文件路径
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

// getLogFileFullPath 获取日志文件路径（包含文件名与后缀）
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// openLogFile 打开日志文件
func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)

	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

// mkDir 创建目录
func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
