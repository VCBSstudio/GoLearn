package job

import (
	"context"
	"log"
	// "time"
)

func (j *CleanupJob) RunWithLock(lock *DistributedLock) {
	ctx := context.Background()

	// 尝试获取锁
	acquired, err := lock.Acquire(ctx)
	if err != nil {
		log.Printf("获取分布式锁失败: %v", err)
		return
	}
	if !acquired {
		log.Println("已有其他实例在执行清理任务")
		return
	}
	defer lock.Release(ctx)

	// 执行清理逻辑
	j.Run()
}
