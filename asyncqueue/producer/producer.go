package producer

import (
    "context"
    "encoding/json"
    "github.com/redis/go-redis/v9"
    "asyncqueue/model"
)

type Producer struct {
    rdb *redis.Client
}

func NewProducer(rdb *redis.Client) *Producer {
    return &Producer{rdb: rdb}
}

func (p *Producer) Enqueue(task *model.Task) error {
    ctx := context.Background()
    taskJSON, err := json.Marshal(task)
    if err != nil {
        return err
    }
    
    // 使用 LPUSH 将任务添加到队列头部
    return p.rdb.LPush(ctx, "task_queue", taskJSON).Err()
}