package rabbit

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/hogwartsinventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
		"",               // exchange
		Pubq[qname].Name, // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		})

	internal.HandleError(err, "Failed to publish a message", internal.Warn)
}

// Subscribe listens to 'subq' (ministry)
// Each time a message is received
// it is parsed and handled
func Subscribe(db *sql.DB) {
	msgs, err := Chan.Consume(
		Subq.Name, // queue
		"",        // consumer
		false,     // auto-ack (should the message be removed from queue after beind read)
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	internal.HandleError(err, "Failed to register a consumer", internal.Warn)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a mail: %s", d.Body)

			// TODO: check message content, and publish on condition, to the right queue

			if d.Body != nil {

				var slot dto.Slot
				var arrested dto.Arrested
				var born dto.Birth


				cannotParseSlot     := mailDecode(d.Body, &slot)
				cannotParseArrested := mailDecode(d.Body, &arrested)
				cannotParseBorn     := mailDecode(d.Body, &born)

				if cannotParseSlot == nil {
					internal.Debug("new mail is a slot request")

					err, availableSlot := checkSlot(db)
					if err != nil {
						internal.Warn(err.Error())
					}

					err = slotHogwarts(availableSlot)
					if err != nil {
						internal.Warn("cannot inform guest about available slot")
					}

				} else if cannotParseArrested == nil {
					internal.Debug("new mail informing hogwarts that a wizard has been arrested")

					err := safetyHogwarts(arrested)
					if err != nil {
						internal.Warn("cannot inform guest and families about hogwarts status")
					}

				} else if cannotParseBorn == nil {
					internal.Debug("new mail informing hogwarts that a wizard just born")

					student, err := bornStudent(born, db)
					if err != nil {
						internal.Warn("cannot enrolled the new student")
					}
					go isTenYo(student)
				}

				if err != nil {
					internal.Warn(fmt.Sprintf("%s", err))

					err := d.Nack(true, true)
					if err != nil {
						internal.Warn("cannot n.ack current message %s")
					}
				}

				err = d.Ack(false)
				if err != nil {
					internal.Warn("cannot ack the current message : %s")
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
	internal.HandleError(err, "Failed to declare a queue", internal.Warn)

	return q
}

func bornStudent(born dto.Birth, db *sql.DB) (student dao.Student, err error) {
	newStudent := dao.Student{
		ID:       uuid.Must(uuid.NewV4()),
		WizardID: born.WizardID,
		House:    "",
		Status:   "enrolled",
	}

	err = hogwartsinventory.CreateStudent(newStudent, db)
	if err != nil {
		internal.Warn("cannot take a new student")
		return student, err
	}

	return newStudent, err
}

func checkSlot(db *sql.DB) (err error, available int) {

	query := "SELECT * FROM actions WHERE status = 'ongoing' and action = 'visit'"

	ongoing, err := hogwartsinventory.GetActions(db, query)
	if err != nil {
		internal.Warn("cannot get actions in hogwarts inventory")
		return err, 0
	}

	if len(ongoing) > 10 {
		return fmt.Errorf("hogwarts have 10 visit ongoing"), 0
	}
	return err, 9
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

func safetyHogwarts(arrested dto.Arrested) (err error) {
	safety, err := json.Marshal(dto.Safety{
		ID:            uuid.Must(uuid.NewV4()),
		WizardID:      arrested.WizardID,
		SafetyMessage: "Hogwarts is ready to receive new visits",
	})
	if err != nil {
		internal.Warn("cannot serialize Attack to JSON")
		return err
	}

	internal.Debug("inform Guest and Families that Hogwarts is no longer under attack")

	Publish("families", string(safety))
	internal.Debug("Mail (safety) sent to families") //TODO: better message

	Publish("guest", string(safety))
	internal.Debug("Mail (safety) sent to guest") //TODO: better message
	return nil
}

func slotHogwarts(availableSlot int) (err error) {
	available, err := json.Marshal(dto.Available{
		ID:               uuid.Must(uuid.NewV4()),
		AvailableSlot:    availableSlot,
		AvailableMessage: "Hogwarts is ready to receive new visits",
	})
	if err != nil {
		internal.Warn("cannot serialize available message")
		return err
	}

	Publish("guest", string(available))
	internal.Debug("message informing guest that there is available slot at hogwarts sent")
	return nil
}

func isTenYo(student dao.Student) (ok bool, err error) {

	for {
		ok, err = checkWizardAge(student)
		if err != nil {
			internal.Warn("cannot get student age")
			return false, err
		}
		if ok == true {
			internal.Debug("wizard can go to Hogwarts")

			err := sendEligibleMail(student)
			if err != nil {
				internal.Warn("cannot inform families that a new wizard is eligible")
				return true, nil
			}

			return true, nil
		} else {
			internal.Debug("wizard is too young to enter Hogwarts")
			time.Sleep(time.Second * 10)
		}
	}
}

func sendEligibleMail(student dao.Student) (err error) {
	eligible, err := json.Marshal(dto.Eligible{
		ID:              uuid.Must(uuid.NewV4()),
		WizardID:        student.WizardID,
		EligibleMessage: "Wizard is ready to go Hogwarts",
	})
	if err != nil {
		internal.Warn("cannot serialize available message")
		return err
	}

	Publish("families", string(eligible))
	internal.Debug("message informing families that a wizard is eligible at hogwarts sent")
	return err
}

func checkWizardAge(student dao.Student) (ok bool, err error) {
	magicURL := internal.GetEnvOrElse("MAGIC_URL", "http://localhost:9090/")

	wizardEndpoint := "wizards/" + student.WizardID.String()

	resp, err := http.Get(magicURL + wizardEndpoint)
	if err != nil {
		internal.Warn(fmt.Sprintf("cannot reach %s", magicURL))

		return false, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		internal.Warn(fmt.Sprintf("cannot read response from %s", magicURL))
		return false, err

	}

	var wizard dao.Wizard

	err = json.Unmarshal(body, &wizard)
	if err != nil {
		internal.Warn("cannot unserialize body to wizard")
		return false, err
	}
	if wizard.Age == 10 {
		internal.Debug(fmt.Sprintf("wizard %s is %f years old !!", wizard.ID.String(), wizard.Age))
		return true, nil
	}
	_ = resp.Body.Close()
	internal.Debug(fmt.Sprintf("wizard %s is %f years old", wizard.ID.String(), wizard.Age))
	return false, nil
}
