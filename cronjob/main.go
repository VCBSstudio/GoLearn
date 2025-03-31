package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cronjob/job"
	"cronjob/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open("root:Hly@1234@tcp(127.0.0.1:3306)/crondb?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 初始化任务管理器
	jobManager := job.NewJobManager()
	defer jobManager.Stop()

	// 添加每日凌晨3点执行的数据清理任务
	cleanupJob := job.NewCleanupJob(db)
	if _, err := jobManager.AddJob("*/30 * * * * *", cleanupJob.Run); err != nil {
		log.Fatal("添加清理任务失败:", err)
	}

	// 添加每30秒执行一次的测试任务
	if _, err := jobManager.AddJob("*/30 * * * * *", func() {
		log.Println("测试任务执行:", time.Now().Format("2006-01-02 15:04:05"))
	}); err != nil {
		log.Fatal("添加测试任务失败:", err)
	}

	// 启动任务管理器
	jobManager.Start()
	log.Println("定时任务服务已启动...")

	// 等待中断信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("正在关闭定时任务服务...")

	// 在数据库连接后添加自动迁移
	if err := db.AutoMigrate(&model.URL{}); err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
}
