package main

import (
    "asyncqueue/consumer"
    "asyncqueue/model"
    "asyncqueue/producer"
    "encoding/json"
    "log"
    "time"
)

func main() {
    // 初始化 RabbitMQ 生产者
    p, err := producer.NewRabbitMQProducer("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("创建生产者失败: %v", err)
    }
    defer p.Close()

    // 初始化 RabbitMQ 消费者
    c, err := consumer.NewRabbitMQConsumer("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("创建消费者失败: %v", err)
    }
    defer c.Close()

    // 启动消费者
    go func() {
        if err := c.Consume("task_queue", func(body []byte) error {
            var task model.Task
            if err := json.Unmarshal(body, &task); err != nil {
                return err
            }
            log.Printf("处理任务: %+v", task)
            return nil
        }); err != nil {
            log.Fatalf("启动消费者失败: %v", err)
        }
    }()

    // 发送测试任务
    for i := 0; i < 5; i++ {
        task := &model.Task{
            ID:      time.Now().String(),
            Type:    "test_task",
            Payload: "测试任务内容",
        }
        taskJSON, _ := json.Marshal(task)
        if err := p.Publish("task_queue", taskJSON); err != nil {
            log.Printf("任务发布失败: %v", err)
        } else {
            log.Printf("任务已发布: %+v", task)
        }
        time.Sleep(1 * time.Second)
    }

    // 保持程序运行
    select {}
}