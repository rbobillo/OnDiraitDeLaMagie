package internal

import (
	"bytes"
	"encoding/json"
	"github.com/rbobillo/OnDiraitDeLaMagie/ministry/dto"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

// Conn is the main connection to rabbit
var Conn *amqp.Connection

// Chan is the main rabbit channel
var Chan *amqp.Channel

// Pubq are all the queues
// where Ministry should publish in
var Pubq = make(map[string]amqp.Queue)

// Subq is the queue Ministry listens to
var Subq amqp.Queue

// Publish sends messages to 'pubq'
func Publish(qname string, payload string) {
	err := Chan.Publish(
		    "",    		// exchange
		Pubq[qname].Name,			// routing key
		false,			// mandatory
		false,			// immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		})

	FailOnError(err, "Failed to publish a message")
}

// Subscribe listens to 'subq' (ministry)
// Each time a message is received
// it is parsed and handled
func Subscribe() {
	msgs, err := Chan.Consume(
		Subq.Name,			// queue
		"",		// consumer
		false,		// auto-ack (should the message be removed from queue after beind read)
		false,		// exclusive
		false,		// no-local
		false,		// no-wait
		nil,			// args
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a mail: %s", d.Body)

			// TODO: check message content, and publish on condition, to the right queue
			if d.Body != nil {

				var help dto.Help

				cannotParseHelp := json.Unmarshal(d.Body, &help) // check if 'help' is well created ?

				if cannotParseHelp == nil {
					ProtectHogwarts(help)
					d.Ack(false)
				} else {
					// try to parse another type of message, or fail
				}
			}
		}
	}()

	log.Printf("Waiting for mails...")

	<-forever
}

// ProtectHogwarts evaluates the emergency
// and helps Hogwarts
func ProtectHogwarts(help dto.Help) {
	hogwartsURL := GetEnvOrElse("HOGWARTS_URL", "http://localhost:9091")

	protection, err := json.Marshal(dto.Protection{
		Quick:  help.Emergency.Quick,
		Strong: help.Emergency.Strong,
	})

	protectEndpoint := "/actions/" + help.AttackID.String() + "/protect"

	req, err := http.NewRequest("POST", hogwartsURL+protectEndpoint, bytes.NewBuffer(protection))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	// TODO: help logic (delay before sending help ? ...)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}

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