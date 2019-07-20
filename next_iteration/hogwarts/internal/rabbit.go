package internal

import (
	"github.com/streadway/amqp"
	"log"
)

// Pubq are all the queues
// where Hogwarts should publish in
var Pubq = make(map[string]amqp.Queue)

// Subq is the queue Hogwarts listens to
var Subq amqp.Queue

// Publish sends messages to 'pubq'
func Publish(ch *amqp.Channel, qname string, payload string) {
	err := ch.Publish(
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
// TODO: better handling; ack/nack ?
func Subscribe(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		Subq.Name, // queue
		"",        // consumer
		true,      // auto-ack
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
				Publish(ch, "ministry", "{}")
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}

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
