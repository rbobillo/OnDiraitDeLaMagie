package rabbit

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/rbobillo/OnDiraitDeLaMagie/azkaban/azkabaninventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/azkaban/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/azkaban/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/azkaban/internal"
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
	err := Chan.Publish(
		    "",    		// exchange
		Pubq[qname].Name,			// routing key
		false,			// mandatory
		false,			// immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		})

	internal.FailOnError(err, "Failed to publish a message")
}

// Subscribe listens to 'subq' (azkaban)
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
	internal.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a mail: %s", d.Body)

			// TODO: check message content, and publish on condition, to the right queue

			if d.Body != nil {

				var arrest   dto.Arrest

				dec := json.NewDecoder(bytes.NewReader(d.Body))
				dec.DisallowUnknownFields()

				err := dec.Decode(&arrest)
				if err != nil {
					internal.Warn("bad format in message : cannot use data")
					d.Ack(false)
				}
				err = arrestWizard(arrest, db)
				if err != nil {
					internal.Warn("cannot arrest wizard")
					d.Ack(false)

				}
				d.Ack(false)
			}
		}
	}()

	log.Printf("Waiting for mails...")

	<-forever
}

func arrestWizard(arrest dto.Arrest, db *sql.DB)(err error){
	prisoner := dao.Prisoner{
		ID:       arrest.ID,
		WizardID: arrest.WizardID,
	}
	err = azkabaninventory.CreatePrisoners(prisoner, db)
	if err != nil {
		internal.Warn(fmt.Sprintf("cannot put wizard %s in prison", arrest.WizardID))
		return err
	}
	arrested, err := json.Marshal(dto.Arrested{
		ID: 			uuid.Must(uuid.NewV4()),
		WizardID:		arrest.WizardID,
		Message: 		fmt.Sprintf("Wizard %s has been put in jail !"),
	})
	if err != nil {
		internal.Warn("fail to format message to json")
		return err
	}
	internal.Debug(fmt.Sprintf("informing ministry that wizard %s has been arrested", arrest.WizardID))
	Publish("ministry", string(arrested))
	return err
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
	internal.FailOnError(err, "Failed to declare a queue")

	return q
}