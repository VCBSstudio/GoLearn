package main

import (
	"GormLearnproject/api/user"
	"GormLearnproject/config"
	"GormLearnproject/model"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	r := gin.Default()

	// 注册用户路由
	user.RegisterRoutes(r, db)

	r.Run(":8080")
}
