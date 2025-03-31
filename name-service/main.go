package main

import (
	"log"
	"name-service/config"
	"name-service/handlers"
	"name-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 初始化 Redis 连接
	if err := config.InitRedis(); err != nil {
		log.Fatalf("Redis 初始化失败: %v", err)
	}

	r := gin.Default()

	nameService := services.NewNameService(db, config.RDB)
	nameHandler := handlers.NewNameHandler(nameService)

	// API 路由
	r.POST("/api/names/generate", nameHandler.GenerateNames)

	r.Run(":8080")
}
