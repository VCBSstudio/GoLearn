package config

import (
	"context"
	"fmt"
	"name-service/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
	Ctx = context.Background()
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:Hly@1234@tcp(127.0.0.1:3306)/name_service?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	// 自动迁移数据库表结构
	if err := db.AutoMigrate(&models.Character{}); err != nil {
		return nil, fmt.Errorf("数据库迁移失败: %v", err)
	}

	DB = db
	return db, nil
}

func InitRedis() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 如果没有设置密码，使用空字符串
		DB:       0,  // 使用默认数据库
	})

	// 测试连接
	if err := RDB.Ping(Ctx).Err(); err != nil {
		return fmt.Errorf("连接Redis失败: %v", err)
	}

	return nil
}
