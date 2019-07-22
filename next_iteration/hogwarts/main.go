package main

import (
	"database/sql"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts/api"
	"github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts/hogwartsinventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts/internal"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"strings"
)

// setupOwls configures rabbit with mandatory queues
func setupOwls() (err error) {
	host := internal.GetEnvOrElse("RABBIT_HOST", "localhost")
	port := internal.GetEnvOrElse("RABBIT_PORT", "5672")
	user := internal.GetEnvOrElse("RABBIT_USER", "magic")
	pass := internal.GetEnvOrElse("RABBIT_PASS", "magic")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)

	internal.Conn, err = amqp.Dial(url)

	internal.FailOnError(err, "Failed to connect to RabbitMQ")

	internal.Chan, err = internal.Conn.Channel()

	internal.FailOnError(err, "Failed to open a channel")

	log.Println("Listening OWL service...")

	// subscribe to the hogwarts queue
	// if it doesn't exist, it creates it
	internal.Subq = internal.DeclareBasicQueue(internal.GetEnvOrElse("SUBSCRIBE_QUEUE", "hogwarts"))

	// set up queues to publish in
	// if they dont exist, it creates them
	for _, q := range strings.Split(internal.GetEnvOrElse("PUBLISH_QUEUES", "ministry,families,guests"), ",") {
		internal.Pubq[q] = internal.DeclareBasicQueue(q)
	}

	return err
}

func setupHogwartsInventory() (*sql.DB, error) {
	hostname := internal.GetEnvOrElse("PG_HOST", "localhost")
	portaddr := internal.GetEnvOrElse("PG_PORT", "5433")
	username := internal.GetEnvOrElse("POSTGRES_USER", "hogwarts")
	password := internal.GetEnvOrElse("POSTGRES_PASSWORD", "hogwarts")
	database := internal.GetEnvOrElse("POSTGRES_DB", "hogwartsinventory")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		hostname, portaddr, username, password, database)

	return hogwartsinventory.InitHogwartsInventory(psqlInfo)
}

func main() {
	db, err := setupHogwartsInventory()

	if err != nil {
		panic(err)
	}

	err = api.InitHogwarts(db)

	if err != nil {
		panic(err)
	}

	err = setupOwls()

	if err != nil {
		panic(err)
	}

	go internal.Subscribe()

	defer internal.Chan.Close()
	defer internal.Conn.Close()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":9091", nil))
}
