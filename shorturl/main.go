package main

import (
	"log"
	"shorturl/api/url"
	"shorturl/config"
	"shorturl/middleware"
	"shorturl/model"
	"shorturl/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitLogger()

	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// 初始化Redis客户端
	rdb := config.InitRedis()

	r := gin.Default()

	// 使用我们自定义的Redis限流中间件 (每分钟60次请求)
	r.Use(middleware.RateLimiter(rdb, 60, time.Minute))

	if err := db.AutoMigrate(&model.URL{}); err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	url.RegisterRoutes(r, db, rdb)

	// // 内存限流 (每秒100个请求)
	// r.Use(limit.NewRateLimiter(time.Second, 100, func(c *gin.Context) (string, error) {
	// 	return c.ClientIP(), nil
	// }))
	r.Run(":8080")
}
