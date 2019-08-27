package main

import (
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/ministry/internal"
	"log"
	"strings"

	"github.com/streadway/amqp" // go get github.com/streadway/amqp
)

// initMinistry sets up Owls with ministry related stuffs
// it creates 'ministry' queue, and 'hogwarts' queue
// then it listens to 'ministry' queue
func initMinistry(url string) (err error) {
	log.Println("Listening OWL service...")

	internal.Conn, err = amqp.Dial(url)

	internal.FailOnError(err, "Failed to connect to RabbitMQ")

	internal.Chan, err = internal.Conn.Channel()

	internal.FailOnError(err, "Failed to open a channel")

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
	host := internal.GetEnvOrElse("RABBIT_HOST", "localhost")
	port := internal.GetEnvOrElse("RABBIT_PORT", "5672")
	user := internal.GetEnvOrElse("RABBIT_USER", "magic")
	pass := internal.GetEnvOrElse("RABBIT_PASS", "magic")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)

	log.Println("Starting ministry service...")

	err := initMinistry(url)

	if err != nil {
		panic(err)
	}

	defer internal.Chan.Close()
	defer internal.Conn.Close()
}