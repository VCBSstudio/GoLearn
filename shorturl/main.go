package main

import (
	"log"
	"shorturl/api/url"
	"shorturl/config"
	"shorturl/model"
	"shorturl/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitLogger()

	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	rdb := config.InitRedis()

	if err := db.AutoMigrate(&model.URL{}); err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	r := gin.Default()
	url.RegisterRoutes(r, db, rdb)
	r.Run(":8080")
}
