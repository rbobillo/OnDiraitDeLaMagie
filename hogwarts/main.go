package main

import (
	"database/sql"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/api"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/hogwartsinventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/rabbit"
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

	rabbit.Subq = rabbit.DeclareBasicQueue(internal.GetEnvOrElse("SUBSCRIBE_QUEUE", "hogwarts"))

	for _, q := range strings.Split(internal.GetEnvOrElse("PUBLISH_QUEUES","ministry,families,guest"), ","){
		rabbit.Pubq[q] = rabbit.DeclareBasicQueue(q)
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
	if err != nil {
		internal.Error(err.Error())
	}

	err = api.InitHogwarts(db)
	if err != nil {
		internal.Error(err.Error())
	}


	err = setupOwls()
	if err != nil {
		internal.Error(err.Error())
	}


	go rabbit.Subscribe(db)

	//// Todo : Handle defer errors
	defer rabbit.Chan.Close()
	defer rabbit.Conn.Close()
	defer db.Close()


	log.Fatal(http.ListenAndServe(":9092", nil))
}