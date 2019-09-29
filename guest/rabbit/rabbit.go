package rabbit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/internal"
	uuid "github.com/satori/go.uuid"
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
	internal.HandleError(err, fmt.Sprintf("failed to declare the queue %s"), internal.Error)
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
	internal.HandleError(err, "failed to publish a message", internal.Error)
}

// Subscribe listens to 'subq' (guest)
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

				var alert     dto.Alert
				var available dto.Available
				var safety    dto.Safety

				cannotParseAlert     := mailDecode(d.Body, &alert)
				cannotParseAvailable := mailDecode(d.Body, &available)
				cannotParseSafety    := mailDecode(d.Body, &safety)

				if cannotParseAlert == nil {
					internal.Debug("guest just receive an Alert mail")
				} else if cannotParseAvailable  == nil {
					internal.Debug("guest just receive an Available mail")
					err = startVisitHogwarts(available)
				} else if cannotParseSafety == nil {
					internal.Debug("guest just receive a Safety mail")

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

func startVisitHogwarts(available dto.Available)(err error){
	hogwartsURL := internal.GetEnvOrElse("HOGWARTS_URL", "http://localhost:9091")

	visitEndpoint := "/actions/visit"

	newVisit, err := json.Marshal(dao.Action{
		ID       : uuid.Must(uuid.NewV4()),
		WizardID : available.GuestID,
		Action   : "visit",
	})

	req, err := http.NewRequest("POST", hogwartsURL+visitEndpoint, bytes.NewBuffer(newVisit))
	if err != nil {
		internal.Warn("cannot create a new http request")
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		internal.Warn("cannot reach hogwarts")
		return err
	}

	_ = resp.Body.Close()
	return err
}
//// AttendHogwarts evaluates the emergency
//// and helps Hogwarts
//func AttendHogwarts(eligible dto.Eligible) {

//}

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