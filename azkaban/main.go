package main

import (
	"database/sql"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/azkaban/azkabaninventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/azkaban/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/azkaban/rabbit"
	"github.com/streadway/amqp"
	"strings"
	_ "github.com/lib/pq"

)

// initAzkabanOwls sets up Owls with azkaban related stuffs
// it creates 'ministry' queue
// then it listens to 'azkaban' queue
func initAzkabanOwls() (err error) {
	host := internal.GetEnvOrElse("RABBITMQ_HOST", "localhost")
	port := internal.GetEnvOrElse("RABBITMQ_PORT", "5672")
	user := internal.GetEnvOrElse("RABBITMQ_USER", "magic")
	pass := internal.GetEnvOrElse("RABBITMQ_PASSWORD", "magic")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)
	internal.Info("Starting azkaban service...")

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

	// subscribe to the azkaban queue
	// if it doesn't exist, it creates it
	rabbit.Subq = rabbit.DeclareBasicQueue(internal.GetEnvOrElse("SUBSCRIBE_QUEUE", "azkaban"))

	// set up queues to publish in
	// if they dont exist, it creates them
	for _, q := range strings.Split(internal.GetEnvOrElse("PUBLISH_QUEUES", "ministry"), ",") {
		rabbit.Pubq[q] = rabbit.DeclareBasicQueue(q)
	}


	return err
}

func setUpAzkabanInventory()(db *sql.DB, err error){
	hostname := internal.GetEnvOrElse("PG_HOST", "localhost")
	portaddr := internal.GetEnvOrElse("PG_PORT", "5434")
	username := internal.GetEnvOrElse("POSTGRES_USER", "azkaban")
	password := internal.GetEnvOrElse("POSTGRES_PASSWORD", "azkaban")
	database := internal.GetEnvOrElse("POSTGRES_DB", "azkabaninventory")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		hostname, portaddr, username, password, database)

	return azkabaninventory.InitAzkabanInventory(psqlInfo)
}



func main() {
	db, err := setUpAzkabanInventory()
	if err != nil {
		internal.Error(err.Error())
	}

	err = initAzkabanOwls()
	if err != nil {
		internal.Error(err.Error())
	}

	rabbit.Subscribe(db)

	defer db.Close()
	defer rabbit.Chan.Close()
	defer rabbit.Conn.Close()
}