package main

import (
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/ministry/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/ministry/rabbit"
	"strings"

	"github.com/streadway/amqp" // go get github.com/streadway/amqp
)

// initMinistry sets up Owls with ministry related stuffs
// it creates 'ministry' queue, and 'hogwarts' queue
// then it listens to 'ministry' queue
func initMinistry(url string) (err error) {
	internal.Debug("Listening OWL service...")

	rabbit.Conn, err = amqp.Dial(url)
	if err != nil {
		internal.Warn("Failed to connect to RabbitMQ")
		return err
	}

	rabbit.Chan, err = rabbit.Conn.Channel()
	if err != nil {
		internal.Warn("Failed to open a channel")
		return err
	}

	// subscribe to the ministry queue
	// if it doesn't exist, it creates it
	rabbit.Subq = rabbit.DeclareBasicQueue(internal.GetEnvOrElse("SUBSCRIBE_QUEUE", "ministry"))

	// set up queues to publish in
	// if they dont exist, it creates them
	for _, q := range strings.Split(internal.GetEnvOrElse("PUBLISH_QUEUES", "hogwarts"), ",") {
		rabbit.Pubq[q] = rabbit.DeclareBasicQueue(q)
	}

	rabbit.Subscribe()

	return err
}

func main() {
	host := internal.GetEnvOrElse("RABBITMQ_HOST", "localhost")
	port := internal.GetEnvOrElse("RABBITMQ_PORT", "5672")
	user := internal.GetEnvOrElse("RABBITMQ_USER", "magic")
	pass := internal.GetEnvOrElse("RABBITMQ_PASSWORD", "magic")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)
	internal.Info("Starting ministry service...")

	err := initMinistry(url)

	if err != nil {
		internal.Error(err.Error())
	}

	defer rabbit.Chan.Close()
	defer rabbit.Conn.Close()
}