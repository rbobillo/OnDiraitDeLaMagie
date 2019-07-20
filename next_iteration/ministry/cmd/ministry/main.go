package main

import (
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/ministry/internal"
	"log"

	"github.com/streadway/amqp"
)

var pubq amqp.Queue
var subq amqp.Queue

// publish sends messages to 'pubq' (hogwarts)
func publish(ch *amqp.Channel, payload string) {
	err := ch.Publish(
		"",        // exchange
		pubq.Name, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		})

	internal.FailOnError(err, "Failed to publish a message")
}

// subscribe listens to 'subq' (ministry)
// Each time a message is received
// it is parsed and handled
// TODO: better handling; ack/nack ?
func subscribe(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		subq.Name, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	internal.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// TODO: check message content, and publish on condition
			if d.Body != nil {
				publish(ch, "{}")
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}

// initMinistry sets up Owls with ministry related stuffs
// it creates 'ministry' queue, and 'hogwarts' queue
// then it listens to 'ministry' queue
func initMinistry(url string) {
	log.Println("Listening OWL service...")

	conn, err := amqp.Dial(url)

	internal.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	internal.FailOnError(err, "Failed to open a channel")

	subq = internal.DeclareBasicQueue(ch, internal.GetEnvOrElse("SUBSCRIBE_QUEUE", "ministry"))
	pubq = internal.DeclareBasicQueue(ch, internal.GetEnvOrElse("PUBLISH_QUEUE", "hogwarts"))

	subscribe(ch)

	defer ch.Close()
	defer conn.Close()
}

func main() {
	host := internal.GetEnvOrElse("RABBIT_HOST", "localhost")
	port := internal.GetEnvOrElse("RABBIT_PORT", "5672")
	user := internal.GetEnvOrElse("RABBIT_USER", "magic")
	pass := internal.GetEnvOrElse("RABBIT_PASS", "magic")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)

	log.Println("Starting ministry service...")

	initMinistry(url)
}
