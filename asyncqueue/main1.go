// package main

// import (
// 	"asyncqueue/consumer"
// 	"asyncqueue/model"
// 	"asyncqueue/producer"
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/redis/go-redis/v9"
// )

// func main() {
// 	// 初始化 Redis 客户端
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "your_strong_password",
// 		DB:       0,
// 	})

// 	// 测试连接
// 	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
// 		log.Fatalf("连接Redis失败: %v", err)
// 	}

// 	// 启动消费者
// 	c := consumer.NewConsumer(rdb)
// 	go c.Start()

// 	// 创建生产者并发送测试任务
// 	p := producer.NewProducer(rdb)
// 	for i := 0; i < 5; i++ {
// 		task := &model.Task{
// 			ID:      time.Now().String(),
// 			Type:    "test_task",
// 			Payload: "测试任务内容",
// 		}
// 		if err := p.Enqueue(task); err != nil {
// 			log.Printf("任务入队失败: %v", err)
// 		} else {
// 			log.Printf("任务已入队: %+v", task)
// 		}
// 		time.Sleep(1 * time.Second)
// 	}

// 	// 保持程序运行
// 	select {}
// }
