package main

import (
	"io"
	"myginproject/api/ping"
	"myginproject/api/user"
	"myginproject/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置日志输出为标准输出（默认）
	gin.DisableConsoleColor()
	// 创建日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default() // gin.Default() 已经包含了 Logger 中间件

	// 注册中间件
	r.Use(middleware.Logger()) // 添加日志中间件
	r.Use(middleware.SimpleRecovery())

	// 注册各个模块的路由
	ping.RegisterRoutes(r)
	user.RegisterRoutes(r)

	r.Run(":9000")
}
