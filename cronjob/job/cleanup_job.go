package job

import (
	"context"
	"log"
	"time"

	"cronjob/model" // 添加这行导入

	"gorm.io/gorm"
)

type CleanupJob struct {
	db *gorm.DB
}

func NewCleanupJob(db *gorm.DB) *CleanupJob {
	return &CleanupJob{db: db}
}

func (j *CleanupJob) Run() {
	log.Println("开始执行数据清理任务...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// 清理7天前的数据
	result := j.db.WithContext(ctx).
		Where("created_at < ?", time.Now().AddDate(0, 0, -7)).
		Delete(&model.URL{}) // 修改为实际的模型

	if result.Error != nil {
		log.Printf("数据清理失败: %v", result.Error)
	} else {
		log.Printf("数据清理完成, 删除了 %d 条记录", result.RowsAffected)
	}
}
