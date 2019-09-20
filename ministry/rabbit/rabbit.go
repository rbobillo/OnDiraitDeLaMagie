package rabbit

import (
	"bytes"
	"encoding/json"
	"github.com/rbobillo/OnDiraitDeLaMagie/ministry/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/ministry/internal"
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

	internal.FailOnError(err, "Failed to publish a message")
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
	internal.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a mail: %s", d.Body)

			// TODO: check message content, and publish on condition, to the right queue

			if d.Body != nil {

				var help     dto.Help
				var born     dto.Born
				var arrested dto.Arrested

				dec := json.NewDecoder(bytes.NewReader(d.Body))
				dec.DisallowUnknownFields()
				cannotParseHelp := dec.Decode(&help) // check if 'help' is well created ?

				dec = json.NewDecoder(bytes.NewReader(d.Body))
				dec.DisallowUnknownFields()
				cannotParseBorn := dec.Decode(&born)

				dec = json.NewDecoder(bytes.NewReader(d.Body))
				dec.DisallowUnknownFields()
				cannotParseArrested := dec.Decode(&arrested)

				if cannotParseHelp == nil {
					err := protectHogwarts(help)
					if err != nil {
						internal.Warn("cannot protect hogwarts")
					} else {
						d.Ack(false)
					}
				} else if cannotParseBorn == nil {
					err := bornWizard(born)
					if err != nil {
						internal.Warn("cannot inform hogwarts about new born wizard")
					} else {
						d.Ack(false)
					}
				} else if cannotParseArrested == nil {
					err := arrestWizard(arrested)
					if err != nil {
						internal.Warn("cannot inform hogwarts about arrested wizard")
					} else {
						d.Ack(false)
					}
				} else {
					internal.Warn("cannot read message : bad format")
				}
			}
		}
	}()

	log.Printf("Waiting for mails...")

	<-forever
}

func protectHogwarts(help dto.Help) (err error) {
	hogwartsURL := internal.GetEnvOrElse("HOGWARTS_URL", "http://localhost:9091")

	protection, err := json.Marshal(dto.Protection{
		Quick:  help.Emergency.Quick,
		Strong: help.Emergency.Strong,
	})

	protectEndpoint := "/actions/" + help.AttackID.String() + "/protect"

	req, err := http.NewRequest("POST", hogwartsURL+protectEndpoint, bytes.NewBuffer(protection))
	if err != nil {
		internal.Warn("cannot create a new request")
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	// TODO: help logic (delay before sending help ? ...)
	resp, err := client.Do(req)

	if err != nil {
		internal.Warn("hogwarts is not reachable")
		return err
	}

	defer resp.Body.Close()
	return err
}
func bornWizard(born dto.Born) (err error){
	payload, err := json.Marshal(born)
	Publish("hogwarts", string(payload))
	internal.Debug("hogwarts receive birth announce")
	return  err
}
func arrestWizard(arrest dto.Arrested) (err error){
	payload, err := json.Marshal(arrest)
	Publish("hogwarts", string(payload))
	internal.Debug("hogwarts receive arrest announce")
	return  err
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