package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/FranciscoHonorat/movies/shared"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	channel   *amqp.Channel
	queueName string
}

func NewRabbitMQPublisher(url string, queueName string) (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	_, err = channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQPublisher{
		channel:   channel,
		queueName: queueName,
	}, nil
}

func (p *RabbitMQPublisher) Publish(ctx context.Context, message shared.MoviePublisherMessage) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return p.channel.Publish(
		"",
		p.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
