package main

import (
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/ministry/debug"
	"github.com/rbobillo/OnDiraitDeLaMagie/ministry/internal"
	"strings"

	"github.com/streadway/amqp" // go get github.com/streadway/amqp
)

// initMinistry sets up Owls with ministry related stuffs
// it creates 'ministry' queue, and 'hogwarts' queue
// then it listens to 'ministry' queue
func initMinistry(url string) (err error) {
	internal.Debug("Listening OWL service...")

	internal.Conn, err = amqp.Dial(url)
	if err != nil {
		internal.Warn("Failed to connect to RabbitMQ")
		return err
	}

	internal.Chan, err = internal.Conn.Channel()
	if err != nil {
		internal.Warn("Failed to open a channel")
		return err
	}

	// subscribe to the ministry queue
	// if it doesn't exist, it creates it
	internal.Subq = internal.DeclareBasicQueue(internal.GetEnvOrElse("SUBSCRIBE_QUEUE", "ministry"))

	// set up queues to publish in
	// if they dont exist, it creates them
	for _, q := range strings.Split(internal.GetEnvOrElse("PUBLISH_QUEUES", "hogwarts"), ",") {
		internal.Pubq[q] = internal.DeclareBasicQueue(q)
	}

	internal.Subscribe()

	return err
}

func main() {
	host := internal.GetEnvOrElse("RABBITMQ_HOST", "localhost")
	port := internal.GetEnvOrElse("RABBITMQ_PORT", "5672")
	user := internal.GetEnvOrElse("RABBITMQ_USER", "magic")
	pass := internal.GetEnvOrElse("RABBITMQ_PASSWORD", "magic")

	debug.PrintEnv()
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)
	internal.Info("Starting ministry service...")

	err := initMinistry(url)

	if err != nil {
		internal.Error(err.Error())
	}

	defer internal.Chan.Close()
	defer internal.Conn.Close()
}