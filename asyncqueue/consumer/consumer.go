package consumer

import (
    "context"
    "encoding/json"
    "github.com/redis/go-redis/v9"
    "asyncqueue/model"
    "log"
    "time"
)

type Consumer struct {
    rdb *redis.Client
}

func NewConsumer(rdb *redis.Client) *Consumer {
    return &Consumer{rdb: rdb}
}

func (c *Consumer) Start() {
    ctx := context.Background()
    for {
        // 使用 BRPOP 阻塞式获取任务
        result, err := c.rdb.BRPop(ctx, 0, "task_queue").Result()
        if err != nil {
            log.Printf("获取任务失败: %v", err)
            time.Sleep(1 * time.Second)
            continue
        }

        var task model.Task
        if err := json.Unmarshal([]byte(result[1]), &task); err != nil {
            log.Printf("解析任务失败: %v", err)
            continue
        }

        log.Printf("处理任务: %+v", task)
        // 这里添加实际的任务处理逻辑
    }
}