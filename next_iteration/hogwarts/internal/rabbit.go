package internal

import (
	"github.com/streadway/amqp"
	"log"
)

// Conn is the main connection to rabbit
var Conn *amqp.Connection

// Chan is the main rabbit channel
var Chan *amqp.Channel

// Pubq are all the queues
// where Hogwarts should publish in
var Pubq = make(map[string]amqp.Queue)

// Subq is the queue Hogwarts listens to
var Subq amqp.Queue

// DeclareBasicQueue is used to declare once
// a RabbitMQ queue, with default parameters
func DeclareBasicQueue(name string) amqp.Queue {
	q, err := Chan.QueueDeclare(name,
		false, // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)

	FailOnError(err, "Failed to declare a queue")

	return q
}

// Publish sends messages to 'pubq'
func Publish(qname string, payload string) {
	err := Chan.Publish(
		"",               // exchange
		Pubq[qname].Name, // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		})

	FailOnError(err, "Failed to publish a message")
}

// Subscribe listens to 'subq' (hogwarts)
// Each time a message is received
// it is parsed and handled
func Subscribe() {
	msgs, err := Chan.Consume(
		Subq.Name, // queue
		"",        // consumer
		false,     // auto-ack (should the message be removed from queue after being read)
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)

	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// TODO: check message content, and publish on condition, to the right queue
			if d.Body != nil {
				d.Ack(false)
			}
		}
	}()

	log.Printf("Waiting for mails...")

	<-forever
}
