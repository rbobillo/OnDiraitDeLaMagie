package internal

import (
	"github.com/streadway/amqp"
)

// DeclareBasicQueue is used to declare once
// a RabbitMQ queue, with default parameters
func DeclareBasicQueue(ch *amqp.Channel, name string) amqp.Queue {
	q, err := ch.QueueDeclare(name,
		false, // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	FailOnError(err, "Failed to declare a queue")

	return q
}
