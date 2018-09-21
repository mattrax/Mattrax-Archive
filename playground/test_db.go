package main

import (
	"log"
	"time"

	"upper.io/db.v3/postgresql"
)

// Connection settings.
var settings = postgresql.ConnectionURL{
	Host:     "127.0.0.1",
	Database: "oscar.beaumont",
	User:     "oscar.beaumont",
	Password: "",
}

type Configuration struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}

func main() {
	start := time.Now()

	sess, err := postgresql.Open(settings) // Open a connection.
	if err != nil {
		panic(err)
	}

	log.Printf("Connected To Database In %s", time.Since(start))
	queryStart := time.Now()

	config := sess.Collection("configuration")

	var Config []Configuration
	res := config.Find()
	err = res.All(&Config)
	if err != nil {
		panic(err)
	}
	log.Println(Config)
	log.Println(Config["OrganisationName"])

	log.Printf("Queried DB In %s", time.Since(queryStart))
}
