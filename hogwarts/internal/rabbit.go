package internal

import (
	"fmt"
	"github.com/streadway/amqp"
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

// DeclareBasicQueue is used to declare once a
// RabbitMQ queue, with default parameters
func DeclareBasicQueue(name string) amqp.Queue{
	q, err := Chan.QueueDeclare(name,
		false,
		false,
		false,
		false,
		nil,
		)
	HandleError(err, fmt.Sprintf("failed to declare the queue %s"), Error)
	return q
}

// Publish sends payload to 'pubq'
func Publish(qname string, payload string){
	err := Chan.Publish(
		"",
		Pubq[qname].Name,
		false,
		false,
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(payload),
		})
	HandleError(err, "failed to publish a message", Error)
}

// Subscribe listens to 'subq' (hogwarts)
// Each time a message is received
// it is parsed and handled
func Subscribe(){
	msgs, err := Chan.Consume (
		Subq.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
		)
	HandleError(err, "failed to register a consumer", Error)

	forever := make(chan bool)

	go func(){
		for d := range msgs {
			Info(fmt.Sprintf("Received a message: %s", d.Body))

			// TODO: check message content, and publish on condition, on the right queue
			if d.Body != nil{
					d.Ack(false)
			}
		}
	}()

	Debug("waiting for mails...")
	<-forever
}