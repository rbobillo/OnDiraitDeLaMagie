package main

import (
	"database/sql"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/api"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/hogwartsinventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	log "github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	"github.com/streadway/amqp"
	"net/http"
	"strings"
)

func setupOwls() (err error) {
	host := internal.GetEnvOrElse("RABBIT_HOST", "localhost")
	port := internal.GetEnvOrElse("RABBIT_PORT", "5672")
	user := internal.GetEnvOrElse("RABBIT_USER", "magic")
	pass := internal.GetEnvOrElse("RABBIT_PASS", "magic")

	url := fmt.Sprintf("amqp://%s:%s@%s:%S/", user, pass, host, port)

	internal.Conn, err = amqp.Dial(url)
	internal.HandleError(err,"failed to connect to RabbitMQ", log.Error)

	internal.Chan, err = internal.Conn.Channel()
	internal.HandleError(err, "failed to open a channel", log.Error)

	internal.Info("listening OWL service...")

	internal.Subq = internal.DeclareBasicQueue(internal.GetEnvOrElse("SUBSCRIBE_QUEUE", "hogwarts"))

	for _, q := range strings.Split(internal.GetEnvOrElse("PUBLISH_QUEUES","ministery,families,guest"), ","){
		internal.Pubq[q] = internal.DeclareBasicQueue(q)
	}
	return err
}

func setupHogwartsInventory() (*sql.DB, error){
	hostname := internal.GetEnvOrElse("PG_HOST", "localhost")
	portaddr := internal.GetEnvOrElse("PG_PORT", "5433")
	username := internal.GetEnvOrElse("POSTGRES_USER","hogwarts")
	password := internal.GetEnvOrElse("POSTGRES_PASSWORD", "hogwarts")
	database := internal.GetEnvOrElse("POSTGRES_DB", "hogwartsinventory")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		hostname, portaddr, username, password, database)

	return hogwartsinventory.InitHogwartsInventory(psqlInfo)
}

func main() {
	db, err := setupHogwartsInventory()
	internal.HandleError(err, "failed to setup hogwarts database", log.Error)

	err = api.InitHogwarts(db)
	internal.HandleError(err, "failed to init hogwarts database", log.Error)

	err = setupOwls()
	internal.HandleError(err, "failed to setup hogwarts Owls service", log.Error)

	go internal.Subscribe()

	// Todo : Handle defer errors
	defer internal.Chan.Close()
	defer internal.Conn.Close()
	defer db.Close()

	internal.HandleError(http.ListenAndServe(":9091", nil), "nya", log.Info)
}