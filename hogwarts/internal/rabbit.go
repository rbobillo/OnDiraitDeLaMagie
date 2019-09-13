package internal

import (
	"database/sql"
	"encoding/json"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/dto"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"log"
	"fmt"
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
	log.Println("coucou")
	log.Println(payload)
	err := Chan.Publish(
		"",    		// exchange
		Pubq[qname].Name,			// routing key
		false,			// mandatory
		false,			// immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		})

	HandleError(err, "Failed to publish a message", Warn)
}

// Subscribe listens to 'subq' (ministry)
// Each time a message is received
// it is parsed and handled
func Subscribe(db *sql.DB) {
	msgs, err := Chan.Consume(
		Subq.Name,			// queue
		"",		// consumer
		false,		// auto-ack (should the message be removed from queue after beind read)
		false,		// exclusive
		false,		// no-local
		false,		// no-wait
		nil,			// args
	)
	HandleError(err, "Failed to register a consumer", Warn)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a mail: %s", d.Body)

			// TODO: check message content, and publish on condition, to the right queue

			if d.Body != nil {

				var slot dto.Slot
				var arrested dto.Arrested

				cannotParseSlot     := json.Unmarshal(d.Body, &slot) // check if 'help' is well created ?
				cannotParseArrested := json.Unmarshal(d.Body, &arrested) // check if 'help' is well created ?

				if cannotParseSlot == nil {

					//TODO : resolve cycle import error
					//err, availableSlot := calldbfromrabbit.CheckSlot(slot, db)
					//if err != nil {
					//	Warn(fmt.Sprintf("%s", err))
					//
					//	err := d.Nack(true, true)
					//	if  err != nil {
					//		Warn(fmt.Sprintf("cannot n.ack current message %s", slot.ID))
					//		return
					//	}
					//}
					//
					//err = d.Ack(false)
					//if err != nil {
					//	Warn(fmt.Sprintf("cannot ack the current message : %s", slot.ID))
					//	return
					//}
					//
					//available, err := json.Marshal(dto.Available{
					//	ID: 			uuid.Must(uuid.NewV4()),
					//	AvailableSlot:  availableSlot,
					//	Message: 		"Hogwarts is ready to receive new visits",
					//})
					//Publish("guest", string(available))
				} else  if cannotParseArrested == nil {

					err = d.Ack(false)
					if err != nil {
						Warn(fmt.Sprintf("cannot ack the current message : %s", arrested.ID))
						return
					}

					Debug("inform Guest and Families that Hogwarts is no longer under attack")

					safety, err := json.Marshal(dto.Safety{
						ID: 			uuid.Must(uuid.NewV4()),
						WizardID: 		arrested.WizardID,
						Message: 		"Hogwarts is ready to receive new visits",
					})
					if err != nil {
						Warn("cannot serialize Attack to JSON")
						return
					}

					Publish("families", string(safety))
					Debug("Mail (safety) sent to families") //TODO: better message

					Publish("guest", string(safety))
					Debug("Mail (safety) sent to guest") //TODO: better message


				}
			}
		}
	}()

	log.Printf("Waiting for mails...")

	<-forever
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
	HandleError(err, "Failed to declare a queue", Warn)

	return q
}
