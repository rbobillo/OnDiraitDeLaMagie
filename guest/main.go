package main

import (
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/api"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/rabbit"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"strings"
)

func setupOwls() (err error) {
	host := internal.GetEnvOrElse("RABBIT_HOST", "localhost")
	port := internal.GetEnvOrElse("RABBIT_PORT", "5672")
	user := internal.GetEnvOrElse("RABBIT_USER", "magic")
	pass := internal.GetEnvOrElse("RABBIT_PASS", "magic")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)

	rabbit.Conn, err = amqp.Dial(url)
	if err != nil {
		internal.Error("failed to connect to RabbitMQ")
		return err
	}

	rabbit.Chan, err = rabbit.Conn.Channel()
	if err != nil {
		internal.Error(err.Error())
		return err
	}

	internal.Info("listening OWL service...")

	rabbit.Subq = rabbit.DeclareBasicQueue(internal.GetEnvOrElse("SUBSCRIBE_QUEUE", "guest"))

	for _, q := range strings.Split(internal.GetEnvOrElse("PUBLISH_QUEUES", "hogwarts"), ",") {
		rabbit.Pubq[q] = rabbit.DeclareBasicQueue(q)
	}
	return err
}

func main() {
	internal.Debug("starting guest micro-service. Waiting for event...")
	err := api.InitGuest()

	if err != nil {
		internal.Warn(err.Error())
	}
	err = setupOwls()

	if err != nil {
		panic(err)
	}

	go rabbit.Subscribe()

	defer rabbit.Chan.Close()
	defer rabbit.Conn.Close()

	log.Fatal(http.ListenAndServe(":9093", nil))
}
