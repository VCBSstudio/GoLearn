package main

import (
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

    // 添加用户路由
    r.POST("/users", func(c *gin.Context) {
        var user model.User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        result := db.Create(&user)
        if result.Error != nil {
            c.JSON(500, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(200, user)
    })

    // 获取用户列表
    r.GET("/users", func(c *gin.Context) {
        var users []model.User
        result := db.Find(&users)
        if result.Error != nil {
            c.JSON(500, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(200, users)
    })

    r.Run(":8080")
}