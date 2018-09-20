package main

import (
	"log"

	"upper.io/db.v3/postgresql"
)

// Connection settings.
var settings = postgresql.ConnectionURL{
	Host:     "127.0.0.1",
	Database: "oscar.beaumont",
	User:     "oscar.beaumont",
	Password: "",
}

func main() {
	sess, err := postgresql.Open(settings) // Open a connection.
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	var config struct {
		OrganisationName  string `db:"OrganisationName"`
		OrganisationEmail string `db:"OrganisationEmail"`
		OrganisationPhone string `db:"OrganisationPhone"`
	}
	if err := sess.Collection("configuration").Find("id", "1").One(&config); err != nil {
		log.Fatal(err)
	}

	log.Println(config.OrganisationName)
}
