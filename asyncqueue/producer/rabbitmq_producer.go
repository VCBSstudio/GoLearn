package producer

import (
	"context"
	// "log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQProducer(url string) (*RabbitMQProducer, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQProducer{conn: conn, ch: ch}, nil
}

func (p *RabbitMQProducer) Publish(queueName string, body []byte) error {
	q, err := p.ch.QueueDeclare(
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		})
}

func (p *RabbitMQProducer) Close() {
	p.ch.Close()
	p.conn.Close()
}
