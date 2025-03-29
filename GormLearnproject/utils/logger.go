package utils

import (
    "github.com/sirupsen/logrus"
    "os"
    "time"
)

var Logger *logrus.Logger

func InitLogger() {
    Logger = logrus.New()
    
    // 设置输出格式为 JSON
    Logger.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: time.RFC3339,
    })
    
    // 输出到文件
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err == nil {
        Logger.SetOutput(file)
    }
    
    // 设置日志级别
    Logger.SetLevel(logrus.InfoLevel)
}