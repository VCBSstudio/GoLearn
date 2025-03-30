package utils

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	} else {
		Logger.SetOutput(os.Stdout)
	}

	// 设置日志级别
	Logger.SetLevel(logrus.InfoLevel)
}

func ErrorResponse(c *gin.Context, code int, message string) {
	Logger.WithFields(logrus.Fields{
		"status": code,
		"error":  message,
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	}).Error("API错误")

	c.JSON(code, gin.H{"error": message})
}
