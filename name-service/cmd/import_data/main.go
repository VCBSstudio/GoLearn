package main

import (
	"log"
	"name-service/config"
	"name-service/tools"
	"path/filepath"
)

func main() {
	// 初始化数据库连接
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 获取数据文件的绝对路径
	dataFile := filepath.Join("data", "characters.json")

	// 导入数据
	if err := tools.ImportCharacters(db, dataFile); err != nil {
		log.Fatalf("导入数据失败: %v", err)
	}

	log.Println("数据导入成功！")
}
