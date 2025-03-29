package main

import (
	"GormLearnproject/api/user"
	"GormLearnproject/config"
	"GormLearnproject/model"
	"GormLearnproject/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志
	utils.InitLogger()

	// 初始化数据库连接
	db, err := config.InitDB()
	if err != nil {
		utils.Logger.Fatal(err)
	}

	// 初始化 Redis 连接
	rdb := config.InitRedis()

	// 自动迁移数据库表
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	r := gin.Default()

	// 注册用户路由
	user.RegisterRoutes(r, db, rdb)

	r.Run(":8080")
}
