package consumer

import (
    "log"
    "time"

    amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
    conn *amqp.Connection
    ch   *amqp.Channel
}

func NewRabbitMQConsumer(url string) (*RabbitMQConsumer, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    return &RabbitMQConsumer{conn: conn, ch: ch}, nil
}

func (c *RabbitMQConsumer) Consume(queueName string, handler func([]byte) error) error {
    q, err := c.ch.QueueDeclare(
        queueName, // name
        true,      // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return err
    }

    msgs, err := c.ch.Consume(
        q.Name, // queue
        "",     // consumer
        false,  // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        return err
    }

    go func() {
        for d := range msgs {
            if err := handler(d.Body); err != nil {
                log.Printf("处理消息失败: %v", err)
                time.Sleep(1 * time.Second)
                continue
            }
            d.Ack(false)
        }
    }()

    return nil
}

func (c *RabbitMQConsumer) Close() {
    c.ch.Close()
    c.conn.Close()
}