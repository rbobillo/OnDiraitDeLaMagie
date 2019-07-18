package main

import (
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/ministry/internal"
	"log"

	"github.com/streadway/amqp"
)

var pubq amqp.Queue
var subq amqp.Queue

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
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
	failOnError(err, "Failed to declare a queue")

	return q
}

// Publish sends messages to 'pubq' (hogwarts)
func Publish(ch *amqp.Channel, payload string) {
	err := ch.Publish(
		"",        // exchange
		pubq.Name, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		})

	failOnError(err, "Failed to publish a message")
}

// Subscribe listens to 'subq' (ministry)
// Each time a message is received
// it is parsed and handled
// TODO: better handling; ack/nack ?
func Subscribe(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		subq.Name, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// TODO: check message content, and publish on condition
			if d.Body != nil {
				Publish(ch, "{}")
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}

// InitMinistry sets up Owls with ministry related stuffs
// it creates 'ministry' queue, and 'hogwarts' queue
// then it listens to 'ministry' queue
func InitMinistry(url string) {
	log.Println("Listening OWL service...")

	conn, err := amqp.Dial(url)

	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	pubq = DeclareBasicQueue(ch, internal.GetEnvOrElse("PUBLISH_QUEUE", "hogwarts"))
	subq = DeclareBasicQueue(ch, internal.GetEnvOrElse("SUBSCRIBE_QUEUE", "ministry"))

	Subscribe(ch)

	defer ch.Close()
	defer conn.Close()
}

func main() {
	host := internal.GetEnvOrElse("RABBIT_HOST", "localhost")
	port := internal.GetEnvOrElse("RABBIT_PORT", "5672")
	user := internal.GetEnvOrElse("RABBIT_USER", "guest")
	pass := internal.GetEnvOrElse("RABBIT_PASS", "guest")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)

	log.Println("Starting ministry service...")

	InitMinistry(url)
}
