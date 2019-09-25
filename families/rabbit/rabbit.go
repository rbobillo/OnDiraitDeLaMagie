package rabbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/families/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/families/internal"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

// Conn is the main connection to rabbit
var Conn *amqp.Connection

// Chan is the main rabbit channel
var Chan *amqp.Channel

// Pubq are all the queues
// where Hogwarts should publish in
var Pubq = make(map[string]amqp.Queue)

// Subq is the queue Families listens to
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
	internal.HandleError(err, fmt.Sprintf("failed to declare the queue %s", name), internal.Error)
	return q
}

func DeclareBindQueue(queueName string, routingKey string, exchange string) (err error){
	err = Chan.QueueBind(
		queueName,     // queue name
		routingKey,    // routing key
		exchange,      // exchange
		false,
		nil)
	return err
}
func DeclareExchange(name string, kind string){
	err := Chan.ExchangeDeclare(
		name, // name
		kind,      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	internal.HandleError(err, fmt.Sprintf("failed to declare an exchange %s", name), internal.Error)
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
	internal.HandleError(err, "failed to publish a message", internal.Error)
}

// Subscribe listens to 'subq' (families)
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
	internal.HandleError(err, "Failed to register a consumer", internal.Warn)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a mail: %s", d.Body)

			// TODO: check message content, and publish on condition, to the right queue

			if d.Body != nil {

				var alert    dto.Alert
				var eligible dto.Eligible
				var safety   dto.Safety

				cannotParseAlert    := mailDecode(d.Body, &alert)
				cannotParseEligible := mailDecode(d.Body, &eligible)
				cannotParseSafety   := mailDecode(d.Body, &safety)

				if cannotParseAlert == nil {
					test, err := Chan.QueueInspect("families")
					if err != nil {
						internal.Warn("cannot inspect queue hogwarts")
						log.Println(err)
						return
					}
					log.Println(test.Messages)
					//TODO: inspect queue,
					// get number of Alert mail
					// stop publish in hogwarts queue
					// until number of Alert mail = 0

					internal.Debug("families just receive an Alert mail")
				} else if cannotParseEligible  == nil {
					internal.Debug("families just receive an Eligible mail")
					err = AttendHogwarts(eligible)
				} else if cannotParseSafety == nil {
					internal.Debug("families just receive a Safety mail")
					test, err := Chan.QueueInspect("hogwarts")
					if err != nil {
						internal.Warn("cannot inspect queue hogwarts")
						return
					}
					log.Println(test)
				}

				if err != nil {
					//TODO : set requeue arg to true
					// to test in real condition
					d.Nack(true, true)
				} else {
					d.Ack(false)
				}
			}
		}
	}()

	log.Printf("Waiting for mails...")

	<-forever
}

func mailDecode(payload []byte, dtoFormat interface{}) (err error){

	dec := json.NewDecoder(bytes.NewReader(payload))
	dec.DisallowUnknownFields()

	err = dec.Decode(&dtoFormat)
	if err != nil {
		return err
	}
	return err
}

// AttendHogwarts evaluates the emergency
// and helps Hogwarts
func AttendHogwarts(eligible dto.Eligible) (err error) {
	hogwartsURL := internal.GetEnvOrElse("HOGWARTS_URL", "http://localhost:9091")

	attendEndpoint := "/actions/" + eligible.WizardID.String() + "/attend"

	eligibleWizard, err := json.Marshal(eligible)

	req, err := http.NewRequest("PATCH", hogwartsURL+attendEndpoint, bytes.NewBuffer(eligibleWizard))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		internal.Warn("hogwarts is not available")
		return err
	}

	defer resp.Body.Close()
	return nil
}

//func AlertHogwarts(alert dto.Alert, Conn amqp.Connection) {
//	Conn.addB
//		addBlockedListener(new BlockedListener() {
//		public void handleBlocked(String reason) throws IOException {
//			// Connection is now blocked
//		}
//
//		public void handleUnblocked() throws IOException {
//			// Connection is now unblocked
//		}
//	});
//}