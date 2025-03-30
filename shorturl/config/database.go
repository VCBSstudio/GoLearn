package config

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:Hly@1234@tcp(127.0.0.1:3306)/shorturl?charset=utf8mb4&parseTime=True&loc=Local"

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// 配置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
}
